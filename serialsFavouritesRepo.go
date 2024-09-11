package repositories

import (
	pg "github.com/Jarozin/repository/postgres"

	mg "github.com/Jarozin/repository/mongo"

	interfaces "github.com/Jarozin/interfaces2"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsFavouritesRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerialsFavourites {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsFavouritesRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsFavouritesRepoMongo(db, log)
	default:
		return nil
	}
}
