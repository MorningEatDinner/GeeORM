package geeorm

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("faied to connect", err)
	}
	return engine
}

// TestNewEngine: 测试是否能够连接上数据库
func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}
