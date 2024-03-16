package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xiaorui/geeorm"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS Userl").Exec()
	_, _ = s.Raw("CREATE TABLE USER(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE USER(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO USER(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success,, %d affected\n", count)
}
