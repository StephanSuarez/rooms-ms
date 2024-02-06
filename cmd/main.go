package main

import (
	"github.com/tesis/internal/common/conf"
	"github.com/tesis/internal/rooms/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	env    *conf.Env
	route  *gin.Engine
	dbConn *mongo.Database
}

func main() {

	app := App{}

	app.env = conf.NewEnv()

	dbenv := &conf.DbEnv{
		Server:   app.env.MongoServer,
		Username: app.env.MongoUsername,
		Password: app.env.MongoPassword,
		Cluster:  app.env.MongoCluster,
		Dbname:   app.env.DbName,
	}

	app.dbConn = conf.GetDBInstance(dbenv)

	app.route = gin.Default()
	http.Routes(app.route)
}
