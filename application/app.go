package application

import (
	"boilerplate/config/env"
	"boilerplate/router"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func init()  {
	//call function load env
	err := gotenv.Load()
	if err != nil {
		return
	}
}

func StartApp()  {
	// call function router in folder router
	addr := env.Config.ServiceHost + ":" + env.Config.ServicePort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.NewRouter(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}