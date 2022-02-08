package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raviMukti/training-golaang-restful-api/helper"
	"github.com/raviMukti/training-golaang-restful-api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
