package user

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	User
}
type User struct {
	UserReader
	UserWriter
}
type UserReader interface {
	GetUserList(user *model.Pagination) (userList []model.User, err error)
	GetUserByID(id string) (user model.User, err error)
	SignInUser(user model.SignInUser) (id string, err error)
}
type UserWriter interface {
	CreateUser(user model.CreateUser) (id uuid.UUID, err error)
	UpdateUser(user model.UpdateUser) (err error)
	DeleteUser(id string) (err error)
}

func NewUserRepo(db *sqlx.DB, loggers *logrus_log.Logger) *UserRepo {
	return &UserRepo{
		User: User{
			UserReader: NewUserReaderDB(db, loggers),
			UserWriter: NewUserWriterDB(db, loggers),
		},
	}
}
