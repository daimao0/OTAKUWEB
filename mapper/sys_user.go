package mapper

import (
	"log"
	"youblog/common/utils"
	"youblog/modules"
)

func GetUserByEmail(email string) modules.SysUser {
	db, _ := utils.InitDB()
	defer db.Close()
	var sysUser modules.SysUser
	sqlStr := "select `sys_user_id`,`sys_user_email`,`sys_user_username`,`sys_user_password`" +
		"from `sys_user`" +
		" where  `sys_user_email`=?"

	err := db.Get(&sysUser, sqlStr, email)
	if err != nil {
		log.Fatal("查询错误请检查查询语句是否正确", err.Error())
	}
	return sysUser
}

func GetUserByEmailAndPassword(email, password string) modules.SysUser {
	db, _ := utils.InitDB()
	defer db.Close()
	var sysUser modules.SysUser
	sqlStr := "select `sys_user_id`,`sys_user_email`,`sys_user_username`,`sys_user_password`" +
		"from `sys_user`" +
		"where  `sys_user_email`=? and `sys_user_password`=?"
	db.Get(&sysUser, sqlStr, email, password)

	return sysUser
}
