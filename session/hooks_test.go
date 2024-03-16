package session

import (
	"testing"

	"github.com/xiaorui/geeorm/log"
)

// 我们在下面添加一些测试用例
type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

// 在插入一条记录之前我们希望增加1000
func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before insert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})
	u := &Account{}

	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("failed to call hooks after query, got", u)
	}
}
