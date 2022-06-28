package connectdb

import (
	"database/sql"
	"github.com/penkong/data4life/db/pgdb"
	"github.com/penkong/data4life/util"
	"log"
)

func Setup(conf *util.Config) *pgdb.Repo {

	// Open connection to database in this case Postgres13
	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("db not connected!!!", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("ping error!!!", err)
	}

	return pgdb.NewRepo(conn)

}
