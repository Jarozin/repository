package repositories

import (
	pg "github.com/Jarozin/repository/postgres"

	mg "github.com/Jarozin/repository/mongo"

	interfaces "github.com/Jarozin/interfaces2"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerials {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsRepoMongo(db, log)
	default:
		return nil
	}
}
