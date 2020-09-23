package app

import (
	"github.com/davetweetlive/saffron-extension/controllers/test"
	"github.com/davetweetlive/saffron-extension/controllers/twitter"
)

func URLMapping() {
	router.GET("/test", test.Test)

	// Twitter Application URL
	router.GET("/twitter/login", twitter.TwitterLoginHandler)
}
