package talk_training

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var MainDB *sqlx.DB

func InitTalk() {
	mainDb, errCore := sqlx.Open("postgres", "postgres://kunyit_user:Ti8yN7uk65@192.168.100.126/tokopedia-dev-db?sslmode=disable")
	if errCore != nil {
		log.Fatalf("db.main not available, error : %s", errCore.Error())
	}
	MainDB = mainDb
}
