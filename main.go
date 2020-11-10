package main

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"github.com/xybydy/epg/golib/mongo"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

func main() {
	var i = mongo.ChannelMatch{"ke", "ki", "qqq", "qqq"}
	var ctx = context.Background()
	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net`, MongoPass)
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: url, Coll: "tvs", Database: DbName})
	q, k := cli.InsertOne(ctx, i)
	fmt.Println("k", q)
	fmt.Println(k)
}
