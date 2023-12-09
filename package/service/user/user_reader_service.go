package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/util/hash"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type UserReaderService struct {
	repo    *repository.Repository
	loggers *logrus_log.Logger
}

func NewUserReaderService(repo *repository.Repository,
	loggers *logrus_log.Logger) *UserReaderService {
	return &UserReaderService{repo: repo, loggers: loggers}
}
func (s *UserReaderService) GetUserList(pagination *model.Pagination) (userList []model.User, err error) {
	userList, err = s.repo.UserRepository.GetUserList(pagination)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (s *UserReaderService) GetUserByID(id string) (user model.User, err error) {
	user, err = s.repo.UserRepository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserReaderService) SignInUser(user model.SignInUser) (id string, err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		s.loggers.Error(err)
		return id, response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	id, err = s.repo.UserRepository.SignInUser(user)
	if err != nil {
		s.loggers.Error(err)
		return id, response.ServiceError(err, codes.NotFound)
	}
	return id, nil
}
