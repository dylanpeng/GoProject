package main

import (
	"context"
	"entities"
	"fmt"
	"github.com/olivere/elastic"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

func main() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(host),
	)
	if err != nil {
		fmt.Println("can't connect to elasticsearch", err)
	}
	if client != nil {
		fmt.Println("client is not nil")
	}
	fmt.Println("connect to elasticsearch success")

	nearbyObj := entities.Result{Id: "bbb", Icon: "ccc", Name:"name"}
	_, _ = client.Index().
		Index("nearby").
		Type("plus").
		Id(nearbyObj.Id).
		BodyJson(nearbyObj).
		Do(context.Background())

	res := client.Search("ccc")
	fmt.Println(res)

}
