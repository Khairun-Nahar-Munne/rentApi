package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
	beego "github.com/beego/beego/v2/server/web"
	_ "rentApi/routers"
)

func main() {
	// Register database
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=munne password=munne123 host=localhost port=5432 dbname=rent sslmode=disable")
	orm.RunSyncdb("default", false, true)

	// Enable CORS - Add this before beego.Run()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:8080"},  // Your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	// Start the Beego application
	beego.Run()
}