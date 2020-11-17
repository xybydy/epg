package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

var cli, ctx = GetClient()

func GetClient() (*mongo.Client, context.Context) {
	url := fmt.Sprintf(`mongodb+srv://f:%s@epg.spxgj.mongodb.net`, MongoPass)
	ctx := context.Background()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, nil
	}
	return cli, ctx
}

func isAlive(cli *mongo.Client) bool {
	err := cli.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}
	return true
}

func GetData(filter interface{}) (ChannelMatches, error) {
	col := cli.Database("epg").Collection("tvs")
	cur, err := col.Find(ctx, filter, options.Find().SetSort(bson.D{{"chan_name", -1}}))
	if err != nil {
		return nil, err
	}
	var results ChannelMatches
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil

}

func InsertData(input []byte) error {
	data := ChannelMatches{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	var la []interface{}

	for _, td := range data {
		la = append(la, td)
	}

	if !isAlive(cli) {
		cli, ctx = GetClient()
	}

	opts := options.InsertMany().SetOrdered(false)

	col := cli.Database("epg").Collection("tvs")
	col.InsertMany(ctx, la, opts)
	if err != nil {
		return err
	}
	return nil
}
