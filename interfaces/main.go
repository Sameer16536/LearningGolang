// main.go
package main

import (
	"learninggolang/interfaces/infrastructure/postgres"
	"learninggolang/interfaces/service"
)

func main() {
	repo := postgres.NewPostgresUserRepo()
	userService := service.NewUserService(repo)

	userService.CreateUser("Sameer")
}
