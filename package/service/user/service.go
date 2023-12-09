package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"
)

type UserService struct {
	User
}
type User struct {
	UserReader
	UserWriter
}
type UserWriter interface {
	CreateUser(user model.CreateUser) (id uuid.UUID, err error)
	UpdateUser(user model.UpdateUser) (err error)
	DeleteUser(id string) (err error)
}
type UserReader interface {
	GetUserList(pagination *model.Pagination) (userList []model.User, err error)
	GetUserByID(id string) (user model.User, err error)
	SignInUser(user model.SignInUser) (id string, err error)
}

func NewUserService(repos *repository.Repository,
	loggers *logrus_log.Logger) *UserService {
	return &UserService{
		User: User{
			UserReader: NewUserReaderService(repos, loggers),
			UserWriter: NewUserWriterService(repos, loggers),
		},
	}
}
