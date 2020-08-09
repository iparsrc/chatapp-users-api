package app

import "github.com/parsaakbari1209/ChatApp-users-api/controllers"

func MapURLs() {
	router.GET("/ping", controllers.Ping)
}
