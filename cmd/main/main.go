package main

import (
	"bk/src/config"
	"bk/src/db"
	"bk/src/routers"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("hello word")
	cnf := ReadConfig()
	fmt.Println(cnf)
	db.InitMysql(cnf.Mysql)
	fmt.Println("数据库连接成功")

	router := routers.InitRouter()
	_ = router.Run()
}

/**
读取mysql配置文件
*/
func ReadConfig() config.Config {
	var cnf config.Config
	f, _ := os.Open("./src/config/config.yaml")
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	_ = yaml.Unmarshal(data, &cnf)
	return cnf
}