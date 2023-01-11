package postgresDB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=rosie.db.elephantsql.com user=ekgakeek password=QvpKEahVvarOiBKTh2SuPMAHoa3xtfMs dbname=ekgakeek port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Unable to connect to the database")
	}
	fmt.Print("Database connected successfully")
}
