package db

import (
	"database/sql"
	logm "../log"
	_"../src/github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "localhost:3306"
	dbusername = "root"
	dbpassowrd = ""
	dbname     = "notepad"
)

type Dbutil struct {
	db *sql.DB
}

/**
连接数据库
 */
func (d *Dbutil) Mysql_open() {
	logm.Trace.Println("开始连接数据库")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		logm.Error.Println("打开数据库失败")
		return
	}
	logm.Trace.Println("已经连接到数据库")
	d.db = db
}

/**
关闭数据库
 */
func (d *Dbutil) Mysql_close() {
	logm.Trace.Println("开始关闭数据库")
	d.db.Close()
	logm.Trace.Println("关闭数据库成功")
}

/**
插入数据
 */
func (d *Dbutil) Mysql_insert(sql_str string) error {
	logm.Trace.Println("往表里插入数据")
	logm.Out.Println("sql语句：", sql_str)
	_, err := d.db.Exec(sql_str)
	if err != nil {
		logm.Error.Println("插入数据失败,原因:", err)
		panic(err)
		return err
	}
	logm.Trace.Println("插入数据成功")
	return nil
}

/**
查询数据
 */
func (d *Dbutil) Mysql_select(sql_str string) (*sql.Rows, error) {
	logm.Trace.Println("查询数据")
	logm.Out.Println("sql语句：", sql_str)
	rows, err := d.db.Query(sql_str)
	if err != nil {
		logm.Error.Println("查询数据失败，原因：", err)
		panic(err)
		return nil, err
	}
	logm.Trace.Println("查询数据成功")
	return rows, nil
}

/**
查询数据
 */
func (d *Dbutil) Mysql_selectNumber(sql_str string) (int, error) {
	logm.Trace.Println("查询数据")
	logm.Out.Println("sql语句：", sql_str)
	var count = 0
	err := d.db.QueryRow(sql_str).Scan(&count)
	if err != nil {
		logm.Error.Println("查询数据失败，原因：", err)
		panic(err)
		return count, err
	}
	logm.Trace.Println("查询数据成功")
	return count, nil
}
