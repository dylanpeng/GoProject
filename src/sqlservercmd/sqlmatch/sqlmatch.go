package sqlmatch

import (
	"database/sql"
	"entities"
	"fmt"
	"time"
)

//Create a new match
func CreateMatch(match entities.Match, db *sql.DB) (int64, error) {
	// tsql := fmt.Sprintf("INSERT INTO tMatch(MatchDescription, CreatedTime, UpdatedTime, IsDeleted) VALUES('%s', '%s', '%s', '%s') select ID = convert(bigint, SCOPE_IDENTITY());",
	// 	match.MatchDescription, match.CreatedTime.Format("2006-01-02 15:04:05"), match.UpdatedTime.Format("2006-01-02 15:04:05"), "false")

	tsql := "INSERT INTO tMatch(MatchDescription, CreatedTime, UpdatedTime, IsDeleted) " +
		"VALUES(?, ?, ?, ?) " +
		"select ID = convert(bigint, SCOPE_IDENTITY());"
	row := db.QueryRow(tsql, match.MatchDescription, match.CreatedTime, match.UpdatedTime, match.IsDeleted)
	var MatchId int64
	err := row.Scan(&MatchId)
	if err != nil {
		fmt.Println("Error CreateMatch:" + err.Error())
		return -1, err
	}
	return MatchId, nil
}

//Update match
func UpdateMatch(match entities.Match, db *sql.DB) (int64, error) {
	tsql := "UPDATE tMatch SET MatchDescription = ?, UpdatedTime = ?, IsDeleted = ? WHERE MatchID = ?"
	result, err := db.Exec(tsql, match.MatchDescription, time.Now().Format("2006-01-02 15:04:05"), match.IsDeleted, match.MatchID)
	if err != nil {
		fmt.Println("Error UpdateMatch:" + err.Error())
		return -1, err
	}
	return result.RowsAffected()
}

func QueryMatchs(db *sql.DB) ([]entities.Match, error) {
	tsql := "SELECT MatchID, MatchDescription, CreatedTime, UpdatedTime, IsDeleted FROM tMatch WHERE MatchID < 6"
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error QueryMatchs:" + err.Error())
		return nil, err
	}
	defer rows.Close()
	var matchs []entities.Match
	for rows.Next() {
		var createdTime, updatedTime time.Time
		var matchId int64
		var matchDescription string
		var isDeleted bool
		err = rows.Scan(&matchId, &matchDescription, &createdTime, &updatedTime, &isDeleted)
		if err != nil {
			fmt.Println("Error QueryMatchs:" + err.Error())
			continue
		}
		match := entities.Match{MatchID: matchId, MatchDescription: matchDescription, CreatedTime: createdTime, UpdatedTime: updatedTime, IsDeleted: isDeleted}
		matchs = append(matchs, match)
	}
	return matchs, nil
}

func QueryByMatchID(matchId int64, db *sql.DB) (entities.Match, error) {
	tsql := "SELECT MatchID, MatchDescription, CreatedTime, UpdatedTime, IsDeleted FROM tMatch WHERE MatchID = ?"
	row := db.QueryRow(tsql, matchId)
	var match entities.Match
	err := row.Scan(&match.MatchID, &match.MatchDescription, &match.CreatedTime, &match.UpdatedTime, &match.IsDeleted)
	if err != nil {
		fmt.Println("Error QueryByMatchID:" + string(matchId) + err.Error())
	}
	return match, err
}
