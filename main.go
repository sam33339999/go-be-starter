package main

import (
	"github.com/joho/godotenv"
	"github.com/sam33339999/go-be-starter/bootstrap"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
