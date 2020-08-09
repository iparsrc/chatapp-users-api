package app

import (
	"github.com/gin-gonic/gin"
	"github.com/parsaakbari1209/ChatApp-users-api/domain"
)

const (
	uri = "mongodb://localhost:27017"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapURLs()
	domain.ConnectDB(uri)
	router.Run(":8080")
}
