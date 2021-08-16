package settings

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-article/constant"
)

var db *pgxpool.Pool

type DatabaseConfig struct{}

func init() {
	var err error

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s",
		constant.UserDB,
		constant.PasswordDB,
		constant.HostDB,
		constant.NameDB,
		constant.SchemaDB,
	)
	db, err = pgxpool.Connect(context.Background(), psqlInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("API Running...")
	//fmt.Println("On =",constant.HostDB,constant.Port)
}

// Postgresql Db Config
func (DatabaseConfig DatabaseConfig) GetDatabaseConfig() *pgxpool.Pool {
	return db
}
