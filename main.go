package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/reale/ipa-cl-ea/web"
)

func main() {
	web.InitialiseConfig()

	log.Print("chainlink ipa adaptor")
	log.Printf("starting to serve on port %d", web.Conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", web.Conf.Port), web.Api().MakeHandler()))
}
