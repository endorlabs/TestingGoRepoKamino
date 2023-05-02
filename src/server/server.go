package server

import (
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/hashicorp/golang-lru"
	"github.com/owncast/owncast/logging"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("Hello World!")
	logging.Setup(true, false)
}
