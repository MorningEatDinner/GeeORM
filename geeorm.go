package geeorm

import (
	"database/sql"

	"github.com/xiaorui/geeorm/dialect"
	"github.com/xiaorui/geeorm/log"
	"github.com/xiaorui/geeorm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine: 创建一个新的Engine
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	// 发送ping指令确认数据库连接成功
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	// 创建一个新的Dialect
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error(err)
		return
	}
	// 创建一个新的Engine
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success.")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Faied to close database")
	}
	log.Info("Cloase database success")
}

// NewSession: 创建新的session实例
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
