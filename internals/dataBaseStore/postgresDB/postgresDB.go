package postgresDB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DatabaseServer interface {
	GetConn() *gorm.DB
}

type postgresServer struct {
	conn *gorm.DB
}

func ConnectToDB() (*postgresServer, error) {
	var err error
	dsn := "host=rosie.db.elephantsql.com user=ekgakeek password=QvpKEahVvarOiBKTh2SuPMAHoa3xtfMs dbname=ekgakeek port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Unable to connect to the database")
		return nil, err
	}
	fmt.Print("Database connected successfully")
	return &postgresServer{conn: DB}, nil
}

func (p *postgresServer) GetConn() *gorm.DB {
	return p.conn
}
