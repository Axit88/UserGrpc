package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/Axit88/UserGrpc/constants"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDbClient() (sql.DB, error) {
	connection := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v", constants.DB_ID, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return sql.DB{}, err
	}
	defer db.Close()
	return *db, nil
}
