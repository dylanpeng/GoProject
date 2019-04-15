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

	// stmt, err := conn.Prepare(`select top 1 * from tmatch`)
	// if err != nil {
	// 	log.Printf("\nPrepare failed:%T %+v\n", err, err)

	// }
	// //defer stmt.Close()

	// row := stmt.QueryRow()
	// // var somenumber int64
	// // var somechars string
	// err = row.Scan(&match.MatchID, &match.MatchDescription, &match.CreatedTime, &match.UpdatedTime, &match.IsDeleted)
	// if err != nil {
	// 	log.Fatal("Scan failed:", err.Error())
	// }
	// fmt.Println(match)
}
