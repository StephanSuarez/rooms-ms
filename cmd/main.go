package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/common/conf"
	"github.com/tesis/internal/users/http"
)

type App struct {
	env *conf.Env
}

func main() {

	app := App{}

	app.env = conf.NewEnv()
	fmt.Println("hola stephan")

	fmt.Println(app.env.MongoUsername)

	r := gin.Default()
	http.Routes(r)
}
