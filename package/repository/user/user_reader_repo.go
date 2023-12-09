package user

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

var (
	errorSignInUser = errors.New("  Username or Password is Incorrect")
)

type UserReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

func NewUserReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *UserReaderDB {
	return &UserReaderDB{db: db, loggers: loggers}
}
func (repo *UserReaderDB) GetUserList(pagination *model.Pagination) (userList []model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&userList, GetUserListByRoleQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, err
}
func (repo *UserReaderDB) GetUserByID(id string) (user model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetUserByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}

func (repo *UserReaderDB) SignInUser(user model.SignInUser) (id string, err error) {
	loggers := repo.loggers
	err = repo.db.Get(&id, SignInUserQuery, user.NickName, user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, errorSignInUser
	}
	return id, nil
}
