package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/util/hash"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

var (
// groupID    = "groupID"
// newGroupID = "newGroupID"
// oldGroupID   = "oldGroupID"
// teacherID    = "teacherID"
// newTeacherID = "newTeacherID"
// oldTeacherID = "oldTeacherID"
)

type UserWriterService struct {
	repo    *repository.Repository
	loggers *logrus_log.Logger
}

func NewUserWriterService(repo *repository.Repository,
	loggers *logrus_log.Logger) *UserWriterService {
	return &UserWriterService{repo: repo, loggers: loggers}
}
func (s *UserWriterService) CreateUser(user model.CreateUser) (id uuid.UUID, err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		s.loggers.Error(err)
		return id, response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	id, err = s.repo.UserRepository.CreateUser(user)
	if err != nil {
		return id, err
	}
	return id, err
}
func (s *UserWriterService) UpdateUser(user model.UpdateUser) (err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		s.loggers.Error(err)
		return response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	err = s.repo.UserRepository.User.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserWriterService) DeleteUser(id string) (err error) {
	err = s.repo.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
