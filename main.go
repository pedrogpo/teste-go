package main

import (
	"fmt"

	"github.com/pedrogpo/testego/models"
)

func main() {
	user := models.NewUser()

	user.SetNome("pedro")

	println(user.GetNome())

	fmt.Scanln()
}
