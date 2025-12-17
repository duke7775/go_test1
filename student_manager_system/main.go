package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Student struct {
	id          int64     //id自增量，从开始自增，key
	number      int64     //学号唯一
	name        string    //名字，字符长度2-64
	gender      int       //性别：0女，1男
	nationality string    // 国际，字符长度2-64
	create_time time.Time //创建时间
}
type Subject struct {
	id          int64     //id子增量，从1开始自增，key
	name        string    //科目名字，字符长度2-64
	create_time time.Time //创建时间
}
type Score struct {
	id             int64     //id自增量，从1开始自增，key
	student_number int64     //学号
	subject_name   string    //科目名字，字符长度2-64
	score          string    //成绩
	create_time    time.Time //创建时间
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to Student Manager System")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
