package mongo

import (
	"context"
	"encoding/json"
	"fmt"

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
		fmt.Println("leee", err)
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
		fmt.Println("fffff", err)
		return err
	}
	fmt.Println("qweeeeqeqeq", data)
	fmt.Println("leee", input)

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
		fmt.Println("laaa", err)
		return err
	}
	return nil
}
