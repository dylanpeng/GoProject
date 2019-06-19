package main

import (
	"context"
	"encoding/json"
	"entities"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

func main() {
	//初始化es Client
	err := initESClient()

	if err != nil {
		fmt.Printf("create es client failed | err : %s\n", err)
		return
	}

	//初始化实体对象并添加到es
	nearbyObj := entities.Result{Id: "bbb", Icon: "ccc", Name: "name"}
	addRep, err := addESItem(nearbyObj, nearbyObj.Id)
	if err != nil {
		fmt.Printf("create es data failed | err : %s\n", err)
		return
	}
	fmt.Printf("%+v\n", *addRep)

	//根据id获取es数据
	getRep, err := getESItem(nearbyObj.Id)
	if err != nil {
		fmt.Printf("get es data failed | err : %s\n", err)
		return
	}
	fmt.Printf("get rep : %+v\n", *getRep)
	fmt.Printf("get rep source : %s\n", string(getRep.Source))

	//将数据反序列化查看是否正确
	newNearby := &entities.Result{}
	err = json.Unmarshal(getRep.Source, newNearby)
	if err == nil {
		fmt.Printf("get source json unmarshal to nearby  : %+v\n", *newNearby)
	}

	//根据条件查询es数据
	searchRep, err := searchESItem()
	if err != nil {
		fmt.Printf("search es data failed | err : %s\n", err)
		return
	}
	if len(searchRep.Hits.Hits) > 0 {
		fmt.Printf("get rep : %s\n", searchRep.Hits.Hits[0].Source)
	}

	//删除es数据
	delRep, err := deleteESItem(nearbyObj.Id)
	if err != nil {
		fmt.Printf("delete es data failed | err : %s\n", err)
		return
	}
	fmt.Printf("%+v\n", *delRep)

	client.Stop()

}

//初始化es客户端
func initESClient() (err error) {
	client, err = elastic.NewClient(
		elastic.SetURL(host),
	)

	if err != nil {
		fmt.Printf("can't connect to elasticsearch | err : %s \n", err)
		return
	}

	fmt.Println("connect to elasticsearch success")
	return
}

//向es添加数据
func addESItem(body interface{}, id string) (rep *elastic.IndexResponse, err error) {
	rep, err = client.Index().
		Index("poi").
		Id(id).
		BodyJson(body).
		Do(context.Background())
	return
}

//获取数据
func getESItem(id string) (rep *elastic.GetResult, err error) {
	rep, err = client.Get().
		Index("poi").
		Id(id).
		Do(context.Background())
	return
}

//查询
func searchESItem() (rep *elastic.SearchResult, err error) {
	q := elastic.NewQueryStringQuery("icon:ccc")
	rep, err = client.Search("poi").
		Query(q).
		Do(context.Background())

	return
}

//删除数据
func deleteESItem(id string) (rep *elastic.DeleteResponse, err error) {
	rep, err = client.Delete().
		Index("poi").
		Id(id).
		Do(context.Background())
	return
}

func QueryElasticSearch(name, lat, lng string, pageSize int) {
	floatLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return
	}

	floatLng, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		return
	}

	names := strings.Fields(name)

	//bool 查询
	boolSearch := elastic.NewBoolQuery()
	//查询距离点floatLat,floatLng 20km以内的点
	boolSearch.Filter(elastic.NewGeoDistanceQuery("location_search").Distance("20km").Lat(floatLat).Lon(floatLng))
	//分词匹配name vicinity两个属性，should相当于"或"
	mustSearch := elastic.NewBoolQuery()
	for _, data := range names {
		mustSearch.Should(elastic.NewMatchQuery("name", data))
		mustSearch.Should(elastic.NewMatchQuery("vicinity", data))
	}
	//must相当于"并"
	boolSearch.Must(mustSearch)
	res, err := client.Search("poi").
		Query(boolSearch).
		From(0).
		Size(pageSize).
		SortBy(elastic.NewGeoDistanceSort("location_search").Point(floatLat, floatLng)).
		Do(context.Background())

	if err != nil {
		return
	}

	var typ *entities.Result
	//遍历命中的数据，对数据进行类型断言，获取数据
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		result := item.(*entities.Result)
		fmt.Println(result)
	}
	return
}
