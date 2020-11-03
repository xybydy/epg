package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qiniu/qmgo"
)

const DbName = "tvg_id_chan"
const MongoPass = "1ZaaVagptA9N9gJW"

var (
	Cli *qmgo.QmgoClient
	Ctx context.Context
)

func init() {
	Ctx = context.Background()
	getClient()
}

func getClient() {
	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net/%s?retryWrites=true&w=majority`, MongoPass, DbName)
	Cli, _ = qmgo.NewClient(Ctx, &qmgo.Config{Uri: "mongodb+srv://f:1ZaaVagptA9N9gJW@epg.spxgj.mongodb.net", Database: DbName})
}

func InsertData(input []byte) {
	data := ChannelMatches{}

	err := json.Unmarshal(input, data)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	res, err := Cli.Collection.InsertMany(Ctx, data)
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		fmt.Println(res)
	}
}
