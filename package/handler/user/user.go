package user

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/tools/jwt"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	userID = "id"
)

type UserEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewUserEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *UserEndPointHandler {
	return &UserEndPointHandler{service: service, loggers: loggers}
}

// CreateUser
// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body model.CreateUser true "Create User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/create [post]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) CreateUser(ctx *gin.Context) {
	var (
		body model.CreateUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	_, err = h.service.UserService.CreateUser(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	if err != nil {
		//response.ServiceErrorConvert(ctx, err)
		//return
	}
	response.HandleResponse(ctx, response.Created, nil, "created", nil)
}

// CreateUsers
// @Description Create Users
// @Summary Create Users
// @Tags User
// @Accept json
// @Produce json
// @Param create body []model.CreateUser true "Create User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/users/create [post]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) CreateUsers(ctx *gin.Context) {
	var (
		body []model.CreateUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	for _, user := range body {
		_, err = h.service.UserService.CreateUser(user)
		if err != nil {
			response.ServiceErrorConvert(ctx, err)
			return
		}
	}
	response.HandleResponse(ctx, response.Created, nil, "created", nil)
}

// UpdateUser
// @Description Update User
// @Summary Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param update body model.UpdateUser true "Update User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/update/{id} [put]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) UpdateUser(ctx *gin.Context) {
	var (
		body model.UpdateUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	body.ID = id
	err = h.service.UserService.UpdateUser(body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated", nil)
}

// UpdateUsers
// @Description Update Users
// @Summary Update Users
// @Tags User
// @Accept json
// @Produce json
// @Param update body []model.UpdateUsers true "Update Users"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/users/update [put]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) UpdateUsers(ctx *gin.Context) {
	var (
		body []model.UpdateUsers
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	for _, user := range body {
		var userUpdate model.UpdateUser
		userUpdate.ID = user.ID
		userUpdate.Photo = user.Photo
		userUpdate.Password = user.Password
		userUpdate.NickName = user.NickName
		userUpdate.Location = user.Location
		userUpdate.BirthdayDate = user.BirthdayDate
		userUpdate.FullName = user.FullName
		err = h.service.UserService.UpdateUser(userUpdate)
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
			return
		}
	}
	response.HandleResponse(ctx, response.OK, nil, "updated", nil)
}

// DeleteUser
// @Description Delete User
// @Summary Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) DeleteUser(ctx *gin.Context) {
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	err = h.service.UserService.DeleteUser(id.String())
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, errors.New("deleted"), nil)
}

// DeleteUsers
// @Description Delete Users
// @Summary Delete Users
// @Tags User
// @Accept json
// @Produce json
// @Param id body []string true "User IDs"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/users/delete [delete]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) DeleteUsers(ctx *gin.Context) {
	var (
		Ids []string
	)
	err := ctx.ShouldBindJSON(&Ids)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	for _, id := range Ids {
		err = h.service.UserService.DeleteUser(id)
		if err != nil {
			response.ServiceErrorConvert(ctx, err)
			return
		}
	}
	response.HandleResponse(ctx, response.OK, nil, errors.New("deleted"), nil)
}

// GetUserList
// @Description Get User List
// @Summary Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/list [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserList(ctx *gin.Context) {
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	userList, err := h.service.UserService.GetUserList(&pagination)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil,
		userList,
		pagination)
}

// GetUserByID
// @Description Get User
// @Summary Get User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/{id} [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserByID(ctx *gin.Context) {
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil, nil)
		return
	}
	user, err := h.service.UserService.GetUserByID(id.String())
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, user, nil)
}

// SignInUser
// @Description Admin Sign In  User.
// @Summary Admin Sign In User
// @Tags User
// @Accept json
// @Produce json
// @Param signup body model.SignInUser true "Sign In"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/sign-in [post]
func (h *UserEndPointHandler) SignInUser(ctx *gin.Context) {
	var (
		body model.SignInUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
	}
	body.NickName = strings.TrimSpace(body.NickName)
	body.Password = strings.TrimSpace(body.Password)
	id, err := h.service.UserService.SignInUser(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	if id == "" {
		response.HandleResponse(ctx, response.BadRequest,
			errors.New("username or password is incorrect"), nil, nil)
		return
	}
	tokens, err := jwt.GenerateNewTokens(id)
	if err != nil {
		response.HandleResponse(ctx, response.InternalServerError, err, nil, nil)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, tokens, nil)
}

// SignUpUser
// @Description Sign Up User
// @Summary Sign Up User
// @Tags User
// @Accept json
// @Produce json
// @Param create body model.CreateUser true "Create Super Admin"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/sign-up [post]
func (h *UserEndPointHandler) SignUpUser(ctx *gin.Context) {
	var (
		body model.CreateUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}

	_, err = h.service.UserService.CreateUser(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created", nil)
}
