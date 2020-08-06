package mapper

import (
	"log"
	"youblog/common/utils"
	"youblog/modules"
)

func ListSoftWares() []modules.SoftWare {
	db, _ := utils.InitDB()
	defer db.Close()
	var softWares []modules.SoftWare
	sqlStr := "select `software_title`,`software_img`,`software_url`,`software_description` from `page_software` where `software_sort`=1"
	err := db.Select(&softWares, sqlStr)
	if err != nil {
		log.Println("查询错误请检查查询语句是否正确", err.Error())
	}
	return softWares
}
func ListDevTools() []modules.SoftWare {
	db, _ := utils.InitDB()
	defer db.Close()
	var softWares []modules.SoftWare
	sqlStr := "select `software_title`,`software_img`,`software_url`,`software_description` from `page_software` where `software_sort`=2"
	err := db.Select(&softWares, sqlStr)
	if err != nil {
		log.Println("查询错误请检查查询语句是否正确", err.Error())
	}
	return softWares
}
