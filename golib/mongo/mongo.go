package mongo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func InsertData(input []interface{}, colName string) (*mongo.InsertManyResult, error) {
	if !isAlive(cli) {
		cli, ctx = GetClient()
	}
	opts := options.InsertMany().SetOrdered(false)

	col := cli.Database("epg").Collection(colName)

	res, err := col.InsertMany(ctx, input, opts)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func InsertChannels(input []byte, colName string) (*mongo.InsertManyResult, error) {
	data := TvgMaps{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return nil, err
	}

	var la []interface{}

	for _, td := range data {
		la = append(la, td)
	}

	return InsertData(la, colName)
}

func InsertChannel(input []byte, colName string) (*mongo.InsertManyResult, error) {
	data := TvgMap{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return nil, err
	}

	var la []interface{}

	la = append(la, data)

	res, err := InsertData(la, colName)
	if err != nil {
		return nil, errors.New("it already exists")
	}
	return res, nil
}

func MatchID(name, country, colName string) TvgMap {
	var pipe []bson.M
	data := TvgMaps{}

	if !isAlive(cli) {
		cli, ctx = GetClient()
	}

	col := cli.Database("epg").Collection(colName)

	if country != "" {
		pipe = []bson.M{
			{"$search": bson.D{
				{"index", "namecountry"},
				{"compound", bson.D{
					{"must", bson.A{bson.D{
						{"text", bson.D{
							{"query", name},
							{"path", "name"},
						}},
					}}},
					{"filter", bson.A{bson.D{
						{"text", bson.D{
							{"query", country},
							{"path", "country"},
						}},
					}}},
				}},
			}},
			{"$limit": 1},
		}
	} else {
		pipe = []bson.M{
			{"$search": bson.D{
				{"index", "namecountry"},
				{"compound", bson.D{
					{"must", bson.A{bson.D{
						{"text", bson.D{
							{"query", name},
							{"path", "name"},
						}},
					}}},
				}},
			}},
			{"$limit": 1},
		}
	}

	cur, err := col.Aggregate(ctx, pipe, nil)
	if err != nil {
		log.Panic(err)
	}

	err = cur.All(ctx, &data)
	if err != nil {
		log.Panic(err)
	}

	data[0].Name = name

	return data[0]
}

func GetData(nameSort bool, colName string) (TvgMaps, error) {
	data := TvgMaps{}

	if !isAlive(cli) {
		cli, ctx = GetClient()
	}
	opts := options.Find()

	if nameSort {
		opts.SetSort(bson.D{{"chan_name", 1}})
	}

	col := cli.Database("epg").Collection(colName)

	cur, err := col.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Println(err)
	}

	err = cur.All(context.TODO(), &data)
	if err != nil {
		log.Println(err)
	}
	return data, nil
}

func UpdateData(input []byte, colName string) error {
	if !isAlive(cli) {
		cli, ctx = GetClient()
	}

	col := cli.Database("epg").Collection(colName)
	data := TvgMaps{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	for _, i := range data {
		filter := bson.D{{"chan_name", i.Name}}
		update := bson.D{{"$set", bson.M{"tvg_id": i.TvgID}}}
		_, err := col.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteChannels(ids []string, colName string) error {
	if !isAlive(cli) {
		cli, ctx = GetClient()
	}

	col := cli.Database("epg").Collection(colName)

	for _, id := range ids {

		obj, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Fatal(err)
			return err
		}
		_, err = col.DeleteOne(ctx, bson.M{"_id": obj})
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func isDup(err error) bool {
	var e mongo.BulkWriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

func isAlive(cli *mongo.Client) bool {
	err := cli.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}
	return true
}
