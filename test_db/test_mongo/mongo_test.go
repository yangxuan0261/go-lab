package test_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"testing"
	"time"
)

// 参考
// - MongoDB 官方文档: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
// - 翻译后的文档: http://www.mongoing.com/archives/27257

// - mongodb官方的golang驱动基础使用 - https://www.jianshu.com/p/0344a21e8040
// - GitHub 代码: https://github.com/hwholiday/learning_tools/blob/master/mongodb/mongo-go-driver/main.go

/*
D系列的类型使用原生的Go类型简单地构建BSON对象。这可以非常有用的来创建传递给MongoDB的命令。 D系列包含4种类型：
- D：一个BSON文档。这个类型应该被用在顺序很重要的场景， 比如MongoDB命令。
- M: 一个无需map。 它和D是一样的， 除了它不保留顺序。
- A: 一个BSON数组。
- E: 在D里面的一个单一的子项。

bson.D{{
    "name",
    bson.D{{
        "$in",
        bson.A{"Alice", "Bob"}
    }}
}}
*/

type Trainer struct {
	Name string
	Age  int
	City string
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
	ash := Trainer{"Ash", 10, "Pallet Town"}

	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func Test_InsertMulti(t *testing.T) {
	misty := Trainer{"Tom", 18, "Cerulean City"}
	brock := Trainer{"Betty", 11, "Pewter City"}
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

func Test_Update02(t *testing.T) {
	var result Trainer
	filter := bson.D{{"name", "Betty"}}
	update := bson.M{"$set": bson.M{"name": "BettyModify"}}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FindOneAndUpdate:%v\n", result) // result 是查询到修改之前的数据
}

func Test_QueryOne(t *testing.T) {
	// create a value into which the result can be decoded
	var result Trainer
	filter := bson.D{{"name", "Ash"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
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
	}

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

func Test_Delete(t *testing.T) {
	// bson.D{{}} 表示所有
	//filter := bson.D{{}}
	filter := bson.D{{"name", "Ash"}}

	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
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
