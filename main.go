package main

import (
	"bphn.go.id/mr-report/report/builder"
	"bphn.go.id/mr-report/report/repository"
	"bphn.go.id/mr-report/server"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("mysql", DATABASE_URL)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRiskRepository(db)
	excelBuilder := builder.NewExcelBuilder(repo)
	handler := server.NewHandler(excelBuilder)

	httpServer := server.NewHttpServer("", "8443", handler)
	err = httpServer.Start()
	if err != nil {
		log.Fatal(err)
	}

}
