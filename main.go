package main

import (
	"fmt"

	"github.com/abrarnaim015/belajar-golang-dependency-injection/app"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/helper"
	_ "github.com/go-sql-driver/mysql"
)


func main()  {
	server := app.InitializedServer()

	fmt.Println("Server Has Running At http://localhost" + server.Addr + " 🚀")
	
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}