package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name"`
	NickName     string    `json:"nickName" db:"nick_name"`
	BirthdayDate string    `json:"birthdayDate" db:"birthday_date"`
	Photo        string    `json:"photo" db:"photo"`
	Location     string    `json:"location" db:"location"`
	Total        int64     `json:"-" db:"total"`
}
type CreateUser struct {
	ID           uuid.UUID `json:"-"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" lenMin:"0" lenMax:"64" regex:"login"`
	NickName     string    `json:"nickName" db:"nick_name" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	Password     string    `json:"password" db:"password" required:"true" lenMin:"0" lenMax:"64"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Location     string    `json:"location" db:"location"`
}
type UpdateUser struct {
	ID           uuid.UUID `json:"-"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" lenMin:"0" lenMax:"64" regex:"login"`
	NickName     string    `json:"nickName" db:"nick_name" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	Password     string    `json:"password" db:"password" required:"true" lenMin:"0" lenMax:"64"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Location     string    `json:"location" db:"location"`
}
type UpdateUsers struct {
	ID           uuid.UUID `json:"id"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" lenMin:"0" lenMax:"64" regex:"login"`
	NickName     string    `json:"nickName" db:"nick_name" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	Password     string    `json:"password" db:"password" required:"true" lenMin:"0" lenMax:"64"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Location     string    `json:"location" db:"location"`
}
type SignInUser struct {
	NickName string `json:"phoneNumber" db:"phone_number" required:"true"  lenMin:"0" lenMax:"16"`
	Password string `json:"password" db:"password" `
}
type SignInUserResponse struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Role uuid.UUID `json:"role" db:"role_id"`
}
type UserPassword struct {
	Password string `json:"password"`
}
type StudentGroupList struct {
	ID         uuid.UUID `json:"id" db:"group_id"`
	GroupTitle string    `json:"title" db:"title"`
}
type CreateSuperAdmin struct {
	Token       string `json:"token" db:"token" required:"true" lenMin:"0" lenMax:"64"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number" default:"+998901234567" required:"true"  lenMin:"0" lenMax:"16" regex:"phone"`
	Password    string `json:"password" db:"password" default:"EduCRM$007Boss" `
}
