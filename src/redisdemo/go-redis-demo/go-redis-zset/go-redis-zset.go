package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u user) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	client.ZRemRangeByScore("zset_key", "0", "100")
	z1 := redis.Z{Score: 2, Member: user{Id: 1, Name: "name1"}}
	z2 := redis.Z{Score: 6, Member: user{Id: 5, Name: "name5"}}
	z3 := redis.Z{Score: 4, Member: user{Id: 3, Name: "name3"}}
	z4 := redis.Z{Score: 8, Member: user{Id: 7, Name: "name7"}}
	// z5 := redis.Z{Score: 2, Member: "aaa"}
	// z6 := redis.Z{Score: 6, Member: "bbb"}
	// z7 := redis.Z{Score: 4, Member: "ccc"}
	// z8 := redis.Z{Score: 8, Member: "ddd"}
	z := []redis.Z{z1, z2, z3, z4} //, z5, z6, z7, z8}
	err = client.ZAdd("zset_key", z...).Err()
	if err != nil {
		fmt.Println(err)
	}

	zset, _ := client.ZRange("zset_key", 0, -1).Result()
	fmt.Printf("%+v\n", zset)
	zset2, _ := client.ZRangeWithScores("zset_key", 0, -1).Result()
	fmt.Printf("%+v\n", zset2)
	var users []user
	for _, d := range zset2 {
		var newuser user
		byteuser := []byte(d.Member.(string))
		// byteuser, _ := hex.DecodeString(struser)
		json.Unmarshal(byteuser, &newuser)
		users = append(users, newuser)
	}
	fmt.Printf("%+v\n", users)

}
