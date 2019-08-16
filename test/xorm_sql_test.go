package test

import (
	"database/sql"
	"testing"
	//"github.com/xormplus/xorm"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"
)

var driverName = "oci8" //Oracle 驱动
var dataSourceName = "scott/tiger@127.0.0.1:1521/ORCL"
var engine *xorm.Engine

func TestXormOracle(t *testing.T) {
	var err error
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		t.Error(err)
	}
	tabs, err := engine.DBMetas()
	if err != nil {
		t.Error(err)
	}
	println(len(tabs))

	sql := "select * from DEPT"
	results, err := engine.Query(sql)
	println(results)
}

func TestMattnOracle(t *testing.T) {
	var db *sql.DB
	var err error
	if db, err = sql.Open(driverName, dataSourceName); err != nil {
		t.Error(err)
		return
	}
	var rows *sql.Rows
	if rows, err = db.Query("select * from DEPT"); err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var deptno int
		var dname string
		var loc string
		rows.Scan(&deptno, &dname, &loc)
		println(deptno, dname, loc) // 3.14 foo
	}
}
