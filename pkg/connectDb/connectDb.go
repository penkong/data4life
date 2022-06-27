package connectdb

import (
	"database/sql"
	"github.com/penkong/data4life/db/pgdb"
	"github.com/penkong/data4life/util"
	"log"
)

var Store *pgdb.Queries

func Setup(conf *util.Config) {

	// Open connection to database in this case Postgres13
	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("db not connected!!!", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("ping error!!!", err)
	}

	// Create new Store with transactions for queries - DB logic
	Store = pgdb.New(conn)
}
