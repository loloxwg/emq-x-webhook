package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	dbhostsip  = "124.222.47.219:3306" //IP地址
	dbusername = "root"                //用户名
	dbpassword = ""                    //密码
	dbname     = "Test"                //表名
)

type Student struct {
	gorm.Model
	Name    string
	Address string
	UUID    string
}

type Sensor struct {
	gorm.Model
	Temperature float64
	Humidity    float64
	UUID        string
}

func main() {
	//db, err := gorm.Open("mysql", "root:123456@tcp(124.222.47.219:3306)/gozero?charset=utf8")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	// 自动迁移模式
	//db.AutoMigrate(&Sensor{})
	//u1 := uuid.Must(uuid.NewV4(), nil)
	// 创建
	//start := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		//sensor := Sensor{}
		////rand.Seed(time.Now().UnixNano())//
		////db.Create(&Sensor{
		////	Temperature: rand.Float64() * 100,
		////	Humidity:    rand.Float64() * 100,
		////	UUID:        u1.String(),
		////})
		//db.First(&sensor)

	}
	//end := time.Now()
	//fmt.Println("Create:", end.Sub(start))
	//fmt.Println("read:", end.Sub(start))
	fmt.Printf("\nmysql read %s\n", "points read 6m34.655586ms")
	// 读取
	//var student Student
	//db.First(&student, 1) // 查询id为1的product

}
