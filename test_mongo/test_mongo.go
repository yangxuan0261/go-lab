package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 参考: http://www.itfanr.cc/2017/06/28/golang-connect-to-mongodb/
// 参考: https://blog.csdn.net/wangshubo1989/article/details/75105397

type Person struct {
	Name  string
	Phone string
	Age   int8
}

var (
	session *mgo.Session
	err     error
	c       *mgo.Collection
)

func init() {
	session, err = mgo.Dial("www.wilker.cn:32781")
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c = session.DB("test").C("people")
}

func main() {
	// test_insert()
	test_query()
	// test_remove()
	// test_update()

	defer session.Close()
}

func printPerson(p *Person) {
	fmt.Println("--- p:", *p)
	fmt.Printf("--- name:%s , phone:%s\n", p.Name, p.Phone)
}

func test_insert() {
	// err = c.Insert(&Person{"superWang", "13478808311"},
	// 	&Person{"David", "15040268074"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func test_query() {
	// result := Person{}
	// err = c.Find(bson.M{"name": "superWang"}).One(&result) // one
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// printPerson(&result)

	var books []Person
	err = c.Find(bson.M{}).All(&books) // all
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range books {
		printPerson(&val)
	}
}

func test_remove() {
	err := c.Remove(bson.M{"name": "superWang"}) // one
	if err != nil {
		log.Fatal(err)
	}

	var info *mgo.ChangeInfo
	info, err = c.RemoveAll(bson.M{"name": "superWang"}) // all
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("info:", info)
}

func test_update() {
	fmt.Println("--- modify before")
	test_query()

	fmt.Println("--- modify after")
	// result := Person{
	// 	Name:  "hello",
	// 	Phone: "666666",
	// }
	// err = c.Update(bson.M{"name": "superWang"}, &result) // one
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result2 := Person{
		Name:  "world",
		Phone: "7777777",
	}
	var info *mgo.ChangeInfo
	info, err = c.UpdateAll(bson.M{"name": "David"}, bson.M{"$set": result2}) // all
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("info:", info)

	test_query()
}
