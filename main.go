package main

import (
	"Coursework_DB/DB_Comunicate"
	"Coursework_DB/GUI_Interface"
)

//Connector

func init() {
	DB_Comunicate.Formation("postgres", "admin", DB_Comunicate.InitLogs())
}

func main() {
	GUI_Interface.CreateWindow()
	defer DB_Comunicate.Formation("postgres", "admin", DB_Comunicate.CloseLog())
}
