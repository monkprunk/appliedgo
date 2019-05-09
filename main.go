package main

import (
	"Advance-Golang-Programming/advanced/final/api"
	"Advance-Golang-Programming/advanced/final/config"
	"flag"
)

var cv config.ConstantViper

func main() {
	state := flag.String("state", "localhost", "set working environment")
	flag.Parse()

	cv.SetState(state)
	cv.Init()

	api.MiddlewareHttp()

	//port := flag.String("port", "8080", "default port: 8080")
	//stage := flag.String("stage", "dev", "product stage: local/dev/sit/prod")
	// configPath := flag.String("config", "configs", "configuration path")
	// flag.Parse()

	// *stage = config.ParseStage(*stage)
	// conf := &config.Config{}
	// if err := conf.Init(*stage, *configPath); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("config: %+v", conf)

	// server := http.Server{
	// 	Addr:    fmt.Sprintf(":%s", *port),
	// 	Handler: route,
	// }

	// if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	panic(err)
	// }
}
