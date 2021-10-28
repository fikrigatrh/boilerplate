package main

import (
	"boilerplate/application"
	"math/rand"
	"time"
)

type ex1 struct {
	Name string `json:"name"`
	Detail []detailInfo `json:"detail"`
}

type detailInfo struct {
	Alamat string `json:"alamat"`
}

func main() {
	//Call function start app in package application
	application.StartApp()

}

func RandomTime() (int, int) {
	rand.Seed(time.Now().UnixNano())
	minA := 1
	maxA := 59
	time1 := rand.Intn(maxA - minA + 1) + minA

	minB := 0
	maxB := 23
	time2 := rand.Intn(maxB - minB + 1) + minB

	return time1, time2
}