package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/student/gingo/zuoye/defs"
)

//GetAllbook ...   select id, name, author,srcurl from book where id > ?
func GetAllbook() ([]*defs.Book, error) {
	stmtOut, err := dbConn.Prepare(`select id, name, author,srcurl from book`)
	var res []*defs.Book
	rows, err := stmtOut.Query()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id int
		var name, author, srcurl string
		if err := rows.Scan(&id, &name, &author, &srcurl); err != nil {
			return res, err
		}
		c := &defs.Book{ID: id, Name: name, Author: author, Srcurl: srcurl}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
