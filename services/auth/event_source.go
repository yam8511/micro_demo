package main

import (
	"context"
	"log"
	pb "proto/auth/micro"
)

// LoginSubscriber 訂閱登入功能
type LoginSubscriber struct{}

// SomeoneLogin 有人登入
func (c *LoginSubscriber) SomeoneLogin(ctx context.Context, user *pb.User) error {
	log.Println("Welcome back, " + user.Nickname)
	return nil
}

// JoinSubscriber 訂閱登入功能
type JoinSubscriber struct{}

// SomeoneJoin 有人加入
func (c *JoinSubscriber) SomeoneJoin(ctx context.Context, user *pb.User) error {
	log.Println("Welcome join, " + user.Nickname)
	return nil
}
