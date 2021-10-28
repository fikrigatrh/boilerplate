package application

import (
	"boilerplate/config/env"
	"boilerplate/router"
	"log"
)

func StartApp()  {
	// call function router in folder router
	addr := env.Config.ServiceHost + ":" + env.Config.ServicePort
	err := router.NewRouter().Run(addr)
	if err != nil {
		log.Println(err)
		return
	}
}