package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qiniu/qmgo"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

func GetClient() (*qmgo.QmgoClient, context.Context) {
	var ctx = context.Background()
	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net`, MongoPass)
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: url, Coll: "tvs", Database: DbName})
	return cli, ctx
}

func InsertData(input []byte) error {
	data := ChannelMatches{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	cli, ctx := GetClient()
	defer cli.Close(ctx)
	res, err := cli.Collection.InsertMany(ctx, data)
	if err != nil {
		return err
	} else {
		fmt.Println(res)
	}
	return nil
}
