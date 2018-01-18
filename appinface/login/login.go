package login

import (
	mysql "../../db"
	logm "../../log"
	"errors"
)

type Login struct {
	userid   string
	account  string
	password string
}

func (r *Login) GetUsername() string {
	return r.account
}

func (r *Login) GetPassword() string {
	return r.password
}
func (r *Login) GetUserid() string {
	return r.userid
}

func (r *Login) SetUsername(s string) {
	r.account = s
}

func (r *Login) SetPassword(s string) {
	r.password = s
}

func (l *Login) VerifyAccount() error {
	logm.Trace.Println("开始验证用户")
	dbutil := &mysql.Dbutil{}
	dbutil.Mysql_open()
	defer dbutil.Mysql_close()
	str_sql := "SELECT count(*) FROM tb_user t WHERE t.username='" + l.account + "' AND t.`password`='" + l.password + "'"
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

func (l *Login) GetUser() error {
	logm.Trace.Println("获取用户信息")
	dbutil := &mysql.Dbutil{}
	dbutil.Mysql_open()
	str_sql := "SELECT username,password,userid FROM tb_user t WHERE t.username='" + l.account + "' AND t.`password`='" + l.password + "'"
	rows, err := dbutil.Mysql_select(str_sql)
	defer rows.Close()
	defer dbutil.Mysql_close()
	if err != nil {
		logm.Error.Println(err)
		return err
	}
	for rows.Next() {
		rows.Scan(&l.account,&l.password,&l.userid)
	}
	return nil
}
