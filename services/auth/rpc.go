package main

import (
	"context"
	"fmt"
	"os"
	pb "proto/auth/micro"
)

// Auth 驗證使用者
type Auth struct{}

// Ping 測試連線
func (c *Auth) Ping(ctx context.Context, req *pb.PingRequest, res *pb.PongResponse) (err error) {
	res.ServiceName = "auth"
	res.Environment = os.Getenv("PROJECT_ENV")
	return
}

// Register 註冊
func (c *Auth) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.RegisterResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			res.ErrorCode = "500"
			res.ErrorText = fmt.Sprint(catchErr)
			return
		}
	}()

	if req.GetPhone() == "" && req.GetEmail() == "" {
		res.ErrorCode = "0101"
		res.ErrorText = "手機或Email必填"
		return
	}

	if req.GetLastName() == "" {
		res.ErrorCode = "0102"
		res.ErrorText = "名稱必填"
		return
	}

	transferSex := SexMALE
	if req.GetSex() == pb.Sex_FEMALE {
		transferSex = SexFEMALE
	}

	newUser, errCode, registerErr := userRegister(
		req.GetPhone(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetFirstName(),
		req.GetLastName(),
		req.GetNickname(),
		req.GetAddress(),
		transferSex,
	)

	if registerErr != nil {
		res.ErrorCode = errCode
		res.ErrorText = fmt.Sprint(registerErr)
		return
	}

	res.ErrorCode = "0"
	res.User = &pb.User{
		Id:        newUser.ID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Nickname:  newUser.Nickname,
		Address:   newUser.Address,
		Phone:     newUser.Phone,
		Email:     newUser.Email,
		AuthToken: newUser.AuthToken,
	}
	if newUser.Sex == SexMALE {
		res.User.Sex = pb.Sex_MALE
	} else if newUser.Sex == SexFEMALE {
		res.User.Sex = pb.Sex_FEMALE
	}

	return
}

// Login 登入
func (c *Auth) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			res.ErrorCode = "500"
			res.ErrorText = fmt.Sprint(catchErr)
			return
		}
	}()

	var loginMethod LoginMethod
	if req.LoginMethod == pb.LoginMethod_PHONE {
		loginMethod = LoginMethodPHONE
	} else if req.LoginMethod == pb.LoginMethod_EMAIL {
		loginMethod = LoginMethodEMAIL
	} else {
		res.ErrorCode = "0201"
		res.ErrorText = "登入方式錯誤"
		return
	}
	user, errCode, loginErr := userLogin(req.LoginData, req.Password, loginMethod)
	if loginErr != nil {
		res.ErrorCode = errCode
		res.ErrorText = fmt.Sprint(loginErr)
		return
	}

	res.ErrorCode = "0"
	res.LoginCheck = true
	res.User = &pb.User{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Address:   user.Address,
		Phone:     user.Phone,
		Email:     user.Email,
		AuthToken: user.AuthToken,
	}
	if user.Sex == SexMALE {
		res.User.Sex = pb.Sex_MALE
	} else if user.Sex == SexFEMALE {
		res.User.Sex = pb.Sex_FEMALE
	}

	return
}

// Check 確認是否登入
func (c *Auth) Check(ctx context.Context, req *pb.CheckRequest, res *pb.CheckResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			res.ErrorCode = "500"
			res.ErrorText = fmt.Sprint(catchErr)
			return
		}
	}()

	user, errCode, checkErr := userCheck(req.UserId, req.AuthToken)
	if checkErr != nil {
		res.ErrorCode = errCode
		res.ErrorText = fmt.Sprint(checkErr)
		return
	}

	res.ErrorCode = "0"
	res.User = &pb.User{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Address:   user.Address,
		Phone:     user.Phone,
		Email:     user.Email,
		AuthToken: user.AuthToken,
	}
	if user.Sex == SexMALE {
		res.User.Sex = pb.Sex_MALE
	} else if user.Sex == SexFEMALE {
		res.User.Sex = pb.Sex_FEMALE
	}

	return
}

// Logout 確認是否登入
func (c *Auth) Logout(ctx context.Context, req *pb.LogoutRequest, res *pb.LogoutResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			res.ErrorCode = "500"
			res.ErrorText = fmt.Sprint(catchErr)
			return
		}
	}()

	errCode, checkErr := userLogout(req.UserId, req.AuthToken)
	if checkErr != nil {
		res.ErrorCode = errCode
		res.ErrorText = fmt.Sprint(checkErr)
		return
	}

	res.ErrorCode = "0"

	return
}
