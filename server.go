package main

import (
	"log"
	"todo/src"
)


func main(){
	app := src.SetupApp()
	port := ":3000"
	log.Println("Server Started on Port"+ port)
	app.Listen(port)
// log.Fatal(app.Listen(":4000"))
}