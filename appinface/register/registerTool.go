package register

import (
	mysql "../../db"
	logm "../../log"
	"errors"
	"../../util"
)

/**
注册
 */
func UserRegister(r *Register) error {
	logm.Trace.Println("开始注册用户")
	dbutil := &mysql.Dbutil{}
	defer dbutil.Mysql_close()
	sql_str1 := "SELECT count(*) FROM tb_user WHERE username='" + r.Username + "'"
	dbutil.Mysql_open()
	rs, err1 := dbutil.Mysql_selectNumber(sql_str1)
	if err1 != nil {
		logm.Error.Println("注册失败")
		return err1
	}
	if rs != 0 {
		err2 := errors.New("用户已存在")
		logm.Error.Println(err2.Error())
		return err2
	}
	sql_str := "INSERT IGNORE  INTO tb_user (username,password,userid) VALUES('" + r.Username + "','" + r.Password + "','" + util.GetRandomString(32) + "')"
	err := dbutil.Mysql_insert(sql_str)
	if err != nil {
		logm.Error.Println("注册失败")
		return err
	}
	return nil
}
