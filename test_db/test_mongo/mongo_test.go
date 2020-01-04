package test_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"reflect"
	"testing"
	"time"
)

// 参考
// - MongoDB 官方文档: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
// - 翻译后的文档: http://www.mongoing.com/archives/27257

// - mongodb官方的golang驱动基础使用 - https://www.jianshu.com/p/0344a21e8040
// - GitHub 代码: https://github.com/hwholiday/learning_tools/blob/master/mongodb/mongo-go-driver/main.go

// - 不错的示例代码: https://medium.com/glottery/golang-and-mongodb-with-go-mongo-driver-part-1-1c43aba25a1

/*
D系列的类型使用原生的Go类型简单地构建BSON对象。这可以非常有用的来创建传递给MongoDB的命令。 D系列包含4种类型：
- D：一个BSON文档。这个类型应该被用在顺序很重要的场景， 比如MongoDB命令。// 定义: type D []E
- M: 一个无需map。 它和D是一样的， 除了它不保留顺序。// 定义: type M map[string]interface{}
- A: 一个BSON数组。// 定义: type A []interface{}
- E: 在D里面的一个单一的子项。 // 定义: type E struct {
	Key   string
	Value interface{}
}

bson.D{{
    "name",
    bson.D{{
        "$in",
        bson.A{"Alice", "Bob"}
    }}
}}
*/

type Trainer struct {
	Bid  string `bson:"_id,omitempty"` // 映射  mongo 中的字段
	Name string `bson:"name,omitempty"`
	Age  int    `bson:"age,omitempty"`
	City string `bson:"city,omitempty"`
	Desc string `bson:"-"` // 取消映射字段
}

var client *mongo.Client
var collection *mongo.Collection

func init() {
	var err error
	// Set client options
	want, err := readpref.New(readpref.SecondaryMode) //表示只使用辅助节点
	if err != nil {
		panic(err)
	}
	wc := writeconcern.New(writeconcern.WMajority())
	readconcern.Majority()
	opt := options.Client().ApplyURI("mongodb://wilker:123456@192.168.1.177:28017/myblog")
	opt.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
	opt.SetMaxConnIdleTime(5 * time.Second)    //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(200)                    //使用最大的连接数
	opt.SetReadPreference(want)                //表示只使用辅助节点
	opt.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
	opt.SetWriteConcern(wc)                    //请求确认写操作传播到大多数mongod实例

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), opt)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Connected to MongoDB!")

	Test_ConnCol(nil)
}

func Test_ConnCol(t *testing.T) {
	collection = client.Database("myblog").Collection("trainer")
	fmt.Printf("--- collection ok:%t\n", collection != nil)
}

func Test_InsertOne(t *testing.T) {
	ash := Trainer{Name: "Ash", Age: 10, City: "Pallet Town"}

	res, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}

	objId := res.InsertedID.(primitive.ObjectID)
	idStr := objId.Hex()

	fmt.Printf("Inserted a single document: retType:%+v, retVal:%+v\n", reflect.TypeOf(res.InsertedID), res.InsertedID)
	fmt.Printf("Inserted a single document: Hex:%s, len:%d\n", idStr, len(idStr)) // 5e1049d39e94a60f53755bd0, len:24 // 这才是需要的 id 字符串

	println()
	objId2, err := primitive.ObjectIDFromHex(idStr) // 构建一个 ObjectID
	if err != nil {
		log.Fatal(err)
		return
	}

	resUpdate, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objId2},
		bson.M{
			"$set": bson.M{
				"name": "Ash666",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("--- resUpdate cnt:", resUpdate.ModifiedCount) // output: 1
}

func Test_InsertMulti(t *testing.T) {
	misty := Trainer{Name: "Tom", Age: 18, City: "Cerulean City"}
	brock := Trainer{Name: "Betty", Age: 11, City: "Pewter City"}
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func Test_Update(t *testing.T) {
	// 匹配name是“Ash”的文档， 并且将Ash的age增加1
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func Test_UpdateFindOneAndUpdate(t *testing.T) {
	var result Trainer
	filter := bson.D{{"name", "Betty"}}
	update := bson.M{"$set": bson.M{"name": "BettyModify"}}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FindOneAndUpdate:%v\n", result) // result 是查询到修改之前的数据
}

func Test_UpdateMany(t *testing.T) {
	//db.collection.update( criteria, objNew, upsert, multi )
	//criteria : update的查询条件，类似sql update查询内where后面的
	//objNew   : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
	//upsert   : 这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
	//multi    : mongodb默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新

	filter := bson.M{"name": "Ash777"}
	update := bson.M{"$set": bson.M{"name": "Ash666"}}

	res, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("res:%+v\n", res) // result 是查询到修改之前的数据
}

func Test_QueryOne(t *testing.T) {
	// create a value into which the result can be decoded
	filter := bson.D{{"name", "Ash"}}
	res := collection.FindOne(context.TODO(), filter)
	if res.Err() != nil { // 找不到结果
		return
	}

	log.Println("has res")
	var result Trainer
	err := res.Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Found a single document 11: %+v\n", result)

	// 通过 id 查找
	objID2, err := primitive.ObjectIDFromHex("5e104ba3fa23b011bb50ab01")
	if err != nil {
		panic(err)
	}

	filter2 := bson.M{"_id": objID2}
	err = collection.FindOne(context.TODO(), filter2).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Found a single document 22: %+v\n", result)

}

func Test_QueryMulti(t *testing.T) {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(0) // 限制查找结果, 负数只返回一个, 0 返回所有, 正数返回指定个数

	var results []*Trainer

	// 查询所有
	// filter := bson.D{{}}

	// 查询指定名字
	filter := bson.D{{
		"name",
		bson.D{{
			"$in",
			bson.A{"Misty", "Brock"},
		}},
	}}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("aaa") // 找不到结果也会跑到这里来
	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO()) // 关闭游标

	fmt.Printf("Found multiple documents (array of pointers): %d\n", len(results))
	for k, v := range results {
		fmt.Printf("--- k:%d, v:%+v\n", k, v)
	}
}

func Test_DeleteById(t *testing.T) {
	objID, err := primitive.ObjectIDFromHex("5e104697b3b35f3df14655e2")
	if err != nil {
		fmt.Println(err)
		return
	}

	resDelete, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("--- del cnt:", resDelete.DeletedCount) // output: 1
}

func Test_DeleteMany(t *testing.T) {
	// bson.D{{}} 表示所有
	//filter := bson.D{{}}

	//filter := bson.D{{"name", "Ash"}} // 与下面 M 等价
	filter := bson.M{"name": "Ash666"}

	res, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", res.DeletedCount)
}

func Test_DeleteCollection(t *testing.T) {
	err := collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Close(t *testing.T) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("--- Connection to MongoDB closed.")
}

func Test_BsonD(t *testing.T) {

}

func Test_BsonE(t *testing.T) {

}

func Test_BsonM(t *testing.T) {

}

func Test_BsonA(t *testing.T) {

}
