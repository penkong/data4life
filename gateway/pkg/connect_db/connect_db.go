package connectdb

import (
	"database/sql"
	"github.com/penkong/data4life/gateway/db/pgdb"
	"github.com/penkong/data4life/gateway/util"
	"log"
)

var Pdb *pgdb.Repo

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

	Pdb = pgdb.NewRepo(conn)
}
