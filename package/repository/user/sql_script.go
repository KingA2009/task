package user

const (
	CreateUserQuery        = `INSERT INTO crm_user (full_name,nick_name, birthday_date, password, photo, location) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	UpdateUserQuery        = `UPDATE crm_user SET full_name=$1,	nick_name=$2 ,birthday_date=$3,password=$4,photo=$5,location=$6,	updated_at = NOW() WHERE id=$7`
	DeleteUserQuery        = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW(), nick_name = nick_name || id WHERE id=$1 AND deleted_at IS NULL`
	GetUserByIDQuery       = `SELECT id,full_name, nick_name,birthday_date, photo, location FROM crm_user WHERE id=$1 AND deleted_at IS NULL ORDER BY created_at DESC `
	GetUserListByRoleQuery = `SELECT id,full_name, nick_name,birthday_date, photo, location, count(id) as total  FROM crm_user WHERE deleted_at IS NULL GROUP BY id order by created_at DESC LIMIT $1 OFFSET $2`
	SignInUserQuery        = `SELECT id FROM crm_user WHERE  nick_name=$1  AND password=$2 AND	deleted_at IS NULL`
)
