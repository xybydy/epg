package main

import (
	"context"
	"fmt"

	momo "github.com/xybydy/epg/golib/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

func main() {
	var data = []momo.ChannelMatch{{ChanName: "kqwe", TvgID: "ki", UserName: "qqq", GroupTitle: "qqq"}, {ChanName: "f", TvgID: "ki", UserName: "qqq", GroupTitle: "qqq"}}

	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net`, MongoPass)
	ctx := context.Background()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		fmt.Println(err)
	}

	var la []interface{}

	for _, td := range data {
		la = append(la, td)
	}

	opts := options.InsertMany().SetOrdered(false)
	col := cli.Database("epg").Collection("tvs")
	_, err = col.InsertMany(ctx, la, opts)
	fmt.Println(err)

}
