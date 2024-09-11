package repositories

import (
	pg "github.com/Jarozin/repository/postgres"

	mg "github.com/Jarozin/repository/mongo"

	interfaces "github.com/Jarozin/interfaces2"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsUsersRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerialsUsers {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsUsersRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsUsersRepoMongo(db, log)
	default:
		return nil
	}
}
