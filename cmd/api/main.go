package main

import (
	"fmt"

	"github.com/tesis/internal/common/conf"
	"github.com/tesis/internal/rooms/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	env    *conf.Env
	route  *gin.Engine
	dbConn *mongo.Database
	rd     *http.RoomDependencies
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

	fmt.Print(dbenv.Dbname)

	app.dbConn = conf.GetDBInstance(dbenv)

	app.rd = http.NewAppDependencies(app.dbConn)

	app.route = gin.Default()
	http.Routes(app.route, app.rd)
}
