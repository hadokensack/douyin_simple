package main

import (
	"byte_douyin/config"
	"byte_douyin/router"
	"fmt"
)

func main() {
	r := router.Init_Router()
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
