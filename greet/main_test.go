package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestQrCode(t *testing.T) {
	user := User{
		Id:   "123",
		Name: "cotton",
	}
	bytes, _ := json.Marshal(user)
	s := string(bytes)
	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", bytes)
	fmt.Printf("%s\n", s)
}

type User struct {
	Id   string
	Name string
	Dog  Dog
}
type Dog struct {
	Age int
}
