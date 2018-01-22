package login

import (
	mysql "../../db"
	logm "../../log"
	"errors"
)

func VerifyAccount(l *Login) error {
	logm.Trace.Println("开始验证用户")
	dbutil := &mysql.Dbutil{}
	dbutil.Mysql_open()
	defer dbutil.Mysql_close()
	str_sql := "SELECT count(*) FROM tb_user t WHERE t.username='" + l.Account + "' AND t.`password`='" + l.Password + "'"
	count, err := dbutil.Mysql_selectNumber(str_sql)
	if err != nil {
		logm.Error.Println(err)
		return err
	}
	if count == 1 {
		return nil
	} else if count > 1 {
		logm.Error.Println("用户名无效")
		return errors.New("用户名无效")
	} else {
		logm.Error.Println("用户名或密码错误")
		return errors.New("用户名或密码错误")
	}
	return nil
}

func GetUser(l *Login) error {
	logm.Trace.Println("获取用户信息")
	dbutil := &mysql.Dbutil{}
	dbutil.Mysql_open()
	str_sql := "SELECT username,userid FROM tb_user t WHERE t.username='" + l.Account + "' AND t.`password`='" + l.Password + "'"
	rows, err := dbutil.Mysql_select(str_sql)
	defer rows.Close()
	defer dbutil.Mysql_close()
	if err != nil {
		logm.Error.Println(err)
		return err
	}
	for rows.Next() {
		rows.Scan(&l.Account, &l.Userid)
	}
	return nil
}
