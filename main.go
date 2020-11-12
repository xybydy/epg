package main

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

func main() {
	// var i = mongo.ChannelMatch{"ke", "ki", "qqq", "qqq"}
	var ctx = context.Background()
	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net`, MongoPass)
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: url, Coll: "tvs", Database: DbName})
	err := cli.CreateOneIndex(ctx, options.IndexModel{Key: []string{"chan_name"}, Unique: true})
	// err := cli.DropIndex(ctx, []string{"_id"})
	if err != nil {
		fmt.Println(err)
	}
	// q, k := cli.InsertOne(ctx, i)
	// fmt.Println("k", q)
	// fmt.Println(k)
}
