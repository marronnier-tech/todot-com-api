package ctrler

import (
	"time"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: []string{
			"https://todot-com.herokuapp.com",
			"https://todot-com-api.herokuapp.com",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
