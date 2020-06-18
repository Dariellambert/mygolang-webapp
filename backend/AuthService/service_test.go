package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/Dariellambert/mygolang-webapp/global"
	"github.com/Dariellambert/mygolang-webapp/proto"
)

func Test_authServer_Login(t *testing.T) {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("example"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{ID: primitive.NewObjectID(), Email: "test@gmail.com", Username: "Carl", Password: string(pw)})
	server := authServer{}
	_, err := server.Login(context.Background(), &proto.LoginRequest{Login: "test@gmail.com", Password: "example"})
	if err != nil {
		t.Error("1: An Error was returned: ", err.Error())
	}

	_, err = server.Login(context.Background(), &proto.LoginRequest{Login: "something", Password: "something"})
	if err == nil {
		t.Error("2: Error was nil")
	}

	_, err = server.Login(context.Background(), &proto.LoginRequest{Login: "Carl", Password: "example"})
	if err != nil {
		t.Error("3: An Error was returned: ", err.Error())
	}
}
