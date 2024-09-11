package repositories

import (
	pg "github.com/Jarozin/repository/postgres"

	mg "github.com/Jarozin/repository/mongo"

	interfaces "github.com/Jarozin/interfaces2"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewStatisticRepo(db interface{}, log *logrus.Logger) interfaces.IRepoStatistic {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewStatisticRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewStatisticRepoMongo(db, log)
	default:
		return nil
	}
}
