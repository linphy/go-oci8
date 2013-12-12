package main

import (
	"database/sql"
	"fmt"
	_ "github.com/linphy/goracle/godrv"
	"log"
	"os"
	//"strconv"
	"time"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Oracle Driver example For GBK chrset test")

	os.Setenv("NLS_LANG", "SIMPLIFIED CHINESE_CHINA.AL32UTF8")
	t := time.Now()

	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("goracle", "cts/qwer1234@CTS247")
	//godrv.SetAutoCommit(true)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	insert_string := "insert into t_employee(id,name,birthday,address,email,mobilephone,telephone,identity_card,weight,height)\n" +
		"values(seq_t_employee_id.nextval,'张三'||:1,sysdate - :2," +
		"'上海市南京东路11号203室'||:3," +
		"'abcd'||:4||'@gmail.com'," +
		"'138'|| trim(to_char(:5, '00000000'))," +
		"'021-'|| trim(to_char(:6, '00000000'))," +
		"'3504561980' || trim(to_char(:7, '00000000'))," +
		"64,1.72)"
	//tx, _ := db.Begin()
	stmt, err := db.Prepare(insert_string)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("1", "2", "3", "4", int32(5), int32(6), int32(7))
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	//log.Printf(" affected = %d\n", rowCnt)
	fmt.Println(time.Since(t))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
