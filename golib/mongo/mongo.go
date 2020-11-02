package mongo

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
)

const DbName = "tvg_id_chan"
const MongoPass = "1ZaaVagptA9N9gJW"

type TvgIds []TvgId

type TvgId struct {
	ChanName string `bson:"chan_name" json:"chan_name"`
	TvgId    string `bson:"tvg_id" json:"tvg_id"`
}

var (
	Cli *qmgo.QmgoClient
	Ctx context.Context
)

func init() {
	Ctx = context.Background()
	getClient()
}

func getClient() {
	Cli, _ = qmgo.Open(Ctx, &qmgo.Config{Uri: getMongoUrl(), Database: DbName, Coll: "movies"})
}

func getMongoUrl() string {
	url := `mongodb+srv://f:%s@epg.spxgj.mongodb.net/%s?retryWrites=true&w=majority`
	return fmt.Sprintf(url, MongoPass, DbName)
}

func InsertData() {
	var ids = TvgIds{}
	res, err := Cli.Collection.InsertMany(Ctx, ids)
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		fmt.Println(res)
	}
}
