package main

import (
	"database/sql"
	_ "github.com/linphy/go-oci8"
	"github.com/qiniu/iconv"
	"log"
	"os"
)

func main() {
	// 为log添加短文件名,方便查看行数
	Conv, err := iconv.Open("gbk", "utf-8") // utf8 => gbk
	if err != nil {
		log.Println("iconv.Open failed!")
		return
	}
	defer Conv.Close()
	Conv8, err := iconv.Open("utf-8", "gbk") // gbk => utf8
	if err != nil {
		log.Println("iconv.Open failed!")
		return
	}
	defer Conv8.Close()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Oracle Driver example For GBK chrset test")

	os.Setenv("NLS_LANG", "SIMPLIFIED CHINESE_CHINA.ZHS16GBK")

	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oracle", "cts/cts@ORCL")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select name from testdata")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name)
		res, err := db.Exec("insert into test values(:name)", name)
		checkErr(err)
		log.Println(Conv8.ConvString(name), res)
	}
	rows.Close()

	db.Exec("insert into test values(:1)", Conv.ConvString("中文你好"))
	rows, err = db.Query("select * from test")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		log.Printf("Name = %s, len=%d", Conv8.ConvString(name), len(name))
	}
	rows.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
