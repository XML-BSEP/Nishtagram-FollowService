package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()

	g.Run("localhost:8089")
}
