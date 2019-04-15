package main

import (
	"database/sql"
	"entities"
	"fmt"
	"log"
	"sqlservercmd/sqlmatch"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		"**", "**", "**", "**")

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	//create
	match := entities.Match{
		MatchDescription: "皇马 VS 巴萨",
		CreatedTime:      time.Now(),
		UpdatedTime:      time.Now(),
		IsDeleted:        false,
	}
	match.MatchID, err = sqlmatch.CreateMatch(match, conn)
	if err != nil {
		log.Fatal("CreateMatch failed:", err.Error())
	}
	fmt.Println(match)

	//update
	affectedNum, err := sqlmatch.UpdateMatch(entities.Match{MatchID: 4, MatchDescription: "11", IsDeleted: true}, conn)
	if err != nil {
		log.Fatal("CreateMatch failed:", err.Error())
	}
	fmt.Println("update success ", affectedNum)

	//query many
	matchs, err := sqlmatch.QueryMatchs(conn)
	if err != nil {
		log.Fatal("QueryMatchs failed:", err.Error())
	}
	fmt.Println(matchs)

	//query one
	match, err = sqlmatch.QueryByMatchID(4, conn)
	if err != nil {
		log.Fatal("QueryByMatchID failed:", err.Error())
	}
	fmt.Println(match)
}
