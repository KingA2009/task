package service

import (
	"EduCRM/package/repository"
	"EduCRM/package/service/user"
	"EduCRM/util/logrus_log"
)

type Service struct {
	UserService *user.UserService
}

func NewService(repos *repository.Repository, loggers *logrus_log.Logger) *Service {
	return &Service{
		UserService: user.NewUserService(repos, loggers),
	}
}
