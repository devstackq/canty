package databases

import (
	"fmt"

	"canty/config"
	"canty/internal/infrastructures/databases/mongo"
	"canty/internal/infrastructures/databases/postgresql"

	_ "github.com/lib/pq"
)

type Database interface {
	Connect() (interface{}, error)
}

type DatabaseFactory struct{}

func (df *DatabaseFactory) CreateDatabase(dbType string, config config.Config) (Database, error) {
	switch dbType {
	case "postgres":
		return &postgresql.PostgresDatabase{Config: config.DBConfig.Postgres}, nil
	case "mongo":
		return &mongo.Database{Config: config.DBConfig.Mongo}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
