package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qiniu/qmgo"
)

const DbName = "epg"
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
	fmt.Println(url)
	Cli, _ = qmgo.Open(Ctx, &qmgo.Config{Uri: url, Database: DbName})
}

func InsertData(input []byte) error {
	data := ChannelMatches{}

	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	res, err := Cli.Collection.InsertMany(Ctx, data)
	if err != nil {
		return err
	} else {
		fmt.Println(res)
	}
	return nil
}
