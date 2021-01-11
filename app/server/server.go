package server

import (
	"ae-test-task/app/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"log"
)

func Start() {
	logrus.StandardLogger()
	r := gin.Default()
	router.InitRoutes(r)

	n := negroni.Classic()
	n.UseHandler(r)
	r.HandleMethodNotAllowed = true
	r.RedirectTrailingSlash = false
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	log.Print("Server running on port :3000")
	r.Run()
}
