package main

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// 密碼加密
func cryptPassword(passwd string) (cryptedPasswd string) {
	prefix := fmt.Sprintf("%x", sha1.Sum([]byte("micro_demo")))
	cryptedPasswd = fmt.Sprintf("%x", sha1.Sum([]byte(prefix+passwd)))
	return
}

// 加密 auth token
func cryptAuthToken(rawToken string) (authToken string) {
	token := cryptPassword(rawToken)
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)

	data := []string{
		fmt.Sprint(rand.Intn(99999)), // unique
		token, // token
		fmt.Sprint(timestamp), // timestamp
	}
	authToken = strings.Join(data, "||")
	authToken = base64.StdEncoding.EncodeToString([]byte(authToken))
	return
}

// 解密 auth token
func decryptAuthToken(authToken string) (token string, timestamp int64, err error) {
	authByte, decodeErr := base64.StdEncoding.DecodeString(authToken)
	if decodeErr != nil {
		err = decodeErr
		return
	}
	authToken = string(authByte)
	data := strings.SplitN(authToken, "||", 3)
	if len(data) < 3 {
		err = errors.New("auth token invalid")
		return
	}
	token = data[1]
	timestamp, err = strconv.ParseInt(data[2], 0, 64)
	return
}

// 產生登入token
func makeLoginToken(user User) (token string) {
	token = cryptPassword(user.ToString())
	return
}

// 註冊使用者
func userRegister(phone, email, password, firstName, lastName, nickname, address string, sex Sex) (newUser User, errCode string, err error) {
	if nickname == "" {
		nickname = firstName
	}

	users := []User{}
	// 先找Email或Phone有沒有存在
	ormErr := OrmReadHandle(func(db *gorm.DB) (handleErr error) {
		handleErr = db.Select("email, phone").Where("email = ? OR phone = ?", email, phone).Find(&users).Error
		return
	})

	if ormErr != nil {
		errCode = "0103"
		err = ormErr
		return
	}

	for _, user := range users {
		if user.Email != "" && user.Email == email {
			errCode = "0104"
			err = errors.New("Email已被註冊")
			return
		}

		if user.Phone != "" && user.Phone == phone {
			errCode = "0105"
			err = errors.New("手機號碼已被註冊")
			return
		}
	}

	newUser = User{
		Email:     email,
		Phone:     phone,
		Password:  cryptPassword(password),
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Address:   address,
		Sex:       sex,
	}
	newUser.AuthToken = makeLoginToken(newUser)

	ormErr = OrmWriteHandle(func(db *gorm.DB) (handleErr error) {
		handleErr = db.Create(&newUser).Error
		return
	})

	if ormErr != nil {
		errCode = "0106"
		err = ormErr
		return
	}

	return
}

// 登入
func userLogin(loginData, password string, loginMethod LoginMethod) (user User, errCode string, err error) {
	query := ""
	if loginMethod == LoginMethodPHONE {
		query = "phone = ?"
	} else if loginMethod == LoginMethodEMAIL {
		query = "email = ?"
	}
	query += " AND password = ?"
	ormErr := OrmReadHandle(func(db *gorm.DB) (handleErr error) {
		handleErr = db.Where(query, loginData, cryptPassword(password)).First(&user).Error
		return
	})

	if ormErr != nil {
		errCode = "0202"
		err = ormErr
		return
	}

	user.AuthToken = makeLoginToken(user)

	ormErr = OrmWriteHandle(func(db *gorm.DB) (handleErr error) {
		handleErr = db.Model(&user).Update("auth_token", user.AuthToken).Error
		return
	})

	if ormErr != nil {
		errCode = "0203"
		err = ormErr
		return
	}

	return
}

func userCheck(userID int64, authToken string) (user User, errCode string, err error) {
	ormErr := OrmReadHandle(func(db *gorm.DB) (handleErr error) {
		handleErr = db.Where("id = ? AND auth_token = ?", userID, authToken).First(&user).Error
		return
	})

	if ormErr != nil {
		errCode = "0204"
		err = ormErr
		return
	}

	return
}

func userLogout(userID int64, authToken string) (errCode string, err error) {
	ormErr := OrmReadHandle(func(db *gorm.DB) (handleErr error) {
		handleErr = db.Model(&User{
			ID: userID,
		}).Where("auth_token = ?", authToken).Update("auth_token", "").Error
		return
	})

	if ormErr != nil {
		errCode = "0205"
		err = ormErr
		return
	}

	return
}
