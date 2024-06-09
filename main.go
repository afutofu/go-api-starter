package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/afutofu/go-api-starter/router"
)

func main() {
	r := router.SetupRouter()

	fmt.Println("Starting GO API service...")

	fmt.Println(`
 ______     ______        ______     ______   __
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
