package sqlmatch

import (
	"database/sql"
	"entities"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func CreateMatch(match entities.Match, db *sql.DB) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO tMatch(MatchDescription, CreatedTime, UpdatedTime, IsDeleted) VALUES('%s', '%s', '%s', '%s') select ID = convert(bigint, SCOPE_IDENTITY());",
		match.MatchDescription, match.CreatedTime.Format("2006-01-02 15:04:05"), match.UpdatedTime.Format("2006-01-02 15:04:05"), "false")

	row := db.QueryRow(tsql)
	var MatchId int64
	err := row.Scan(&MatchId)
	if err != nil {
		fmt.Println("Error CreateMatch:" + err.Error())
		return -1, err
	}
	return MatchId, nil
}
