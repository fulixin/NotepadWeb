package note

import (
	mysql "../../db"
	logm "../../log"
	"../../util"
)

/**
返回列表
 */
func QueryNoteList() ([]interface{}, error) {
	logm.Trace.Println("获取日志列表")
	mybd := mysql.Dbutil{}
	mybd.Mysql_open()
	str_sql := "SELECT n.noteId,n.noteTitle,n.noteContext,n.noteDate FROM tb_note n"
	rows, err := mybd.Mysql_select(str_sql)
	if err != nil {
		logm.Error.Println(err)
		return nil, err
	}
	notes := []interface{}{}
	for rows.Next() {
		note := Note{}
		rows.Scan(&note.NoteId, &note.NoteTitle, &note.NoteContext, &note.NoteDate)
		notes = append(notes, note)
	}
	return notes, nil
}

func SaveNote(n *Note) error {
	logm.Trace.Println("保存日志")
	mybd := mysql.Dbutil{}
	mybd.Mysql_open()
	str_sql := "INSERT IGNORE  INTO tb_note (noteId,noteTitle,noteContext,noteDate) VALUES('" + util.GetRandomString(32) + "','" + n.NoteTitle + "','" + n.NoteContext + "','" + n.NoteDate + "')"
	err := mybd.Mysql_insert(str_sql)
	if err != nil {
		logm.Error.Println(err)
		return err
	}
	return nil
}
