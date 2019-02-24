package main

import (
	"fmt"

	"beego/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "beego/routers"

	_ "github.com/lib/pq"
)

func init() {
	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=odoo password=odoo dbname=beego host=127.0.0.1 port=5432 sslmode=disable")
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	stu := new(models.Student)
	stu.Name = "tom"
	stu.Age = 25
	fmt.Println(o.Insert(stu))
	beego.Run()
}
