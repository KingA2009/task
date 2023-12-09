package repository

import (
	"EduCRM/package/repository/user"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	UserRepository *user.UserRepo
}

func NewRepository(db *sqlx.DB, loggers *logrus_log.Logger) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepo(db, loggers),
	}
}
