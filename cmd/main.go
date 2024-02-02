package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/common/conf"
	"github.com/tesis/internal/rooms/http"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	env    *conf.Env
	dbConn *mongo.Database
}

func main() {

	app := App{}

	app.env = conf.NewEnv()
	app.dbConn = conf.GetDBInstance()

	r := gin.Default()
	http.Routes(r)
}
