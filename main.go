package main

import (
	"hduhelp/Model"
	"hduhelp/Router"
)

func main() {
	err := Model.Mysql1()
	if err != nil {
		return
	}
	Router.InitRouter()
}
