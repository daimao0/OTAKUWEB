package modules

type SysUser struct {
	UserId   int64  `db:"sys_user_id"`
	Email    string `db:"sys_user_email"`
	Username string `db:"sys_user_username"`
	Password string `db:"sys_user_password"`
}
