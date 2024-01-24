package main

import (
	"fmt"

	"github.com/tesis/internal/users/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hola stephan")
	r := gin.Default()
	http.Routes(r)
}
