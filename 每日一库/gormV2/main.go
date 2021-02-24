package main

//gorm已发布V2版本，版本号为：1.20， V2版本的项目已移至github.com/go-gorm，import 路径也修改为 gorm.io/gorm,v1仍然在github.com/jinzhu/gorm
//数据库驱动被拆分为独立的项目，例如：github.com/go-gorm/sqlite，且它的 import 路径也变更为 gorm.io/driver/sqlite
import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	dsn := "gorm:gorm123@tcp(localhost:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}

}
