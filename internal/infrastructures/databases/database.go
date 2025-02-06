package databases

import (
	"fmt"

	"canty/config"
	repoMongo "canty/internal/infrastructures/databases/mongo"
	"canty/internal/infrastructures/databases/postgresql"

	_ "github.com/lib/pq"
)

type Database interface {
	Connect() (interface{}, error)
	Close() error
}

type DatabaseFactory struct{}

func (df *DatabaseFactory) CreateDatabase(dbType string, config config.Config) (Database, error) {
	switch dbType {
	case "mongo":
		db := &repoMongo.Database{
			Config: repoMongo.Config{
				URI:    config.DBConfig.Mongo.Host,
				DBName: config.DBConfig.Mongo.DBName,
			},
		}
		//cl, err := db.Connect()
		//if err != nil {
		//	return nil, err
		//}
		return db, nil

	case "postgres":
		db := &postgresql.Database{Config: postgresql.Config{
			Host:     config.DBConfig.Postgres.Host,
			Port:     config.DBConfig.Postgres.Port,
			User:     config.DBConfig.Postgres.User,
			Password: config.DBConfig.Postgres.Password,
			DBName:   config.DBConfig.Postgres.DBName,
		}}
		//cl, err := db.Connect()
		//if err != nil {
		//	return nil, err
		//}
		return db, nil

	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
