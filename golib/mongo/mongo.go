package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qiniu/qmgo"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

var cli, ctx = GetClient()

func GetClient() (*qmgo.QmgoClient, context.Context) {
	var ctx = context.Background()
	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net`, MongoPass)
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: url, Coll: "tvs", Database: DbName})
	return cli, ctx
}

func isAlive(cli *qmgo.QmgoClient) bool {
	err := cli.Ping(3)
	if err != nil {
		return false
	}
	return true

}

func InsertData(input []byte) error {
	data := ChannelMatches{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	if !isAlive(cli) {
		cli, ctx = GetClient()
	}
	_, err = cli.Collection.InsertMany(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
