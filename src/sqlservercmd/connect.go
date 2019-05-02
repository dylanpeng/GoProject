package main

import (
	"database/sql"
	"entities"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var match entities.Match
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		"localhost", "sa", "**", "GoDB")

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	stmt, err := conn.Prepare(`select top 1 * from tmatch`)
	if err != nil {
		log.Printf("\nPrepare failed:%T %+v\n", err, err)

	}
	//defer stmt.Close()

	row := stmt.QueryRow()
	// var somenumber int64
	// var somechars string
	err = row.Scan(&match.MatchID, &match.MatchDescription, &match.CreatedTime, &match.UpdatedTime, &match.IsDeleted)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Println(match)
}
