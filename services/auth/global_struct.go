package main

import (
	"fmt"
	"time"
)

// Sex 性別
type Sex int32

const (
	// SexMALE 男生
	SexMALE Sex = 0
	// SexFEMALE 女生
	SexFEMALE Sex = 1
)

// LoginMethod 登入方式
type LoginMethod int32

const (
	// LoginMethodPHONE 手機登入
	LoginMethodPHONE LoginMethod = 0
	// LoginMethodEMAIL 信箱登入
	LoginMethodEMAIL LoginMethod = 1
)

// User 使用者模組
type User struct {
	ID        int64
	Email     string
	Phone     string
	AuthToken string
	Password  string
	FirstName string
	LastName  string
	Nickname  string
	Address   string
	Sex       Sex
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ToString 轉為字串
func (user *User) ToString() string {
	return fmt.Sprintf(
		"id=%d&first_name=%s&last_name=%s&nickname=%s&address=%s&password=%s&phone=%s&email=%s&sex=%v",
		user.ID,
		user.FirstName,
		user.LastName,
		user.Nickname,
		user.Address,
		user.Password,
		user.Phone,
		user.Email,
		user.Sex,
	)
}
