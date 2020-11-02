package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

func getClient() (*qmgo.QmgoClient, error) {
	cli, err := qmgo.Open(context.Background(), &qmgo.Config{Uri: getMongoUrl(), Database: DbName, Coll: "movies"})
	return cli, err
}

func getMongoUrl() string {
	MongoPass:="mongo6161"
	DbName:="channels"
	
	url := "mongodb+srv://f:<password>@epg.spxgj.mongodb.net/<dbname>?retryWrites=true&w=majority"
	return fmt.Sprintf(url, MongoPass, DbName)
}