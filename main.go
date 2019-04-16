package main

import (
	"Recorder/config"
	"SvBlogApi/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("port"))
	log.Printf(http.ListenAndServe(viper.GetString("port"), g).Error())

	go func() {
		err := pingServer()

		if err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		rsp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && rsp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
