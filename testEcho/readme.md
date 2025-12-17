# 测试echo项目
本项目为学生管理系统后端-go语言开发
前端会选用vue[TypeScript脚本语言]框架

## 一 技术栈暂定

- 框架使用echo
- 日志选用logrus
- 读写mysql使用gorm
- 还有 读写redis

## 二 需求

mysql数据库 student_manager_sys

学生-结构体
```
type Student struct {
    id int64    //id自增量，从1开始自增,key
    number int64 //学号唯一
    name string //名字,字符长度2-64
    gender int //性别:0女,1男
    nationality string //国籍,字符长度2-64
    create_time  time.Time//创建时间
}
```

科目-结构体

```
type Subject struct {
    id int64    //id自增量，从1开始自增,key
    name string //科目名字,字符长度2-64
    create_time  time.Time//创建时间
}
```

学生成绩-结构体

```
type Score struct {
    id int64    //id自增量，从1开始自增,key
    student_number int64 //学号
    subject_name string //科目名字,字符长度2-64
    score float64 //成绩成绩，2位小数
    create_time  time.Time//创建时间
}
```

### 2.1 建立路由-每个路由处理函数都必须有返回
    - addStudent 增加学生
    - updateStudent 更新学生信息 
    - addSubject 增加科目
    - updateSubject 更新科目 
    - addScore 增加学生成绩
    - updateScore 更新学生成绩