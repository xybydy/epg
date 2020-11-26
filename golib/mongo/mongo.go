package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
		log.Println("leee", err)
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

func InsertData(input []byte) error {
	data := ChannelMatches{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		log.Println("fffff", err)
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
		log.Println("laaa", err)
		return err
	}
	return nil
}

func GetData(nameSort bool) (ChannelMatches, error) {
	data := ChannelMatches{}

	if !isAlive(cli) {
		cli, ctx = GetClient()
	}
	opts := options.Find()

	if nameSort {
		opts.SetSort(bson.D{{"chan_name", 1}})
	}

	col := cli.Database("epg").Collection("tvs")

	cur, err := col.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Println("hooooo ", err)
	}

	cur.All(context.TODO(), &data)
	return data, nil
}

func UpdateData(input []byte) error {
	if !isAlive(cli) {
		cli, ctx = GetClient()
	}

	col := cli.Database("epg").Collection("tvs")
	data := ChannelMatches{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	for _, i := range data {
		filter := bson.D{{"chan_name", i.ChanName}}
		update := bson.D{{"$set", bson.M{"group_title": i.GroupTitle, "tvg_id": i.TvgID}}}
		_, err := col.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println("laaa", err)
			return err
		}
	}
	return nil
}
