package main

import (
	"database/sql"
	"fmt"
	"github.com/linphy/goracle/godrv"
	"log"
	"os"
	"strconv"
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
	// _, err = db.Exec(insert_string, "888", "888", "888", "888", int32(888), int32(888), int32(888))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//godrv.SetAutoCommit(false)
	tx, _ := db.Begin()
	stmt, err := tx.Prepare(insert_string)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var res sql.Result
	for i := 0; i < 10000; i++ {
		str := strconv.Itoa(i)

		res, err = stmt.Exec(str, str, str, str, int32(i), int32(i), int32(i))
		if err != nil {
			log.Fatal(err)
		}

		_, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	// lastId, err := res.LastInsertId()
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
