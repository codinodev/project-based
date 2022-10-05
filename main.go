package main

import (
	"project-based/controllers"
  "project-based/configs)"
	"github.com/gin-gonic/gin"
)

func main() {
  db := configs.DBInt()
  inDB := &controllers.InDB{DB: db}

  router := gin.Default()
  router.GET("/person/:id", inDB.GetPeson)
  router.GET("/person", inDB.GetPesons)
  router.POST("/person", inDB.CreatePerson)
  router.PUT("/person", inDB.UpdatePerson)
  router.DELETE("/person/:id", inDB.DeletePerson)
  router.Run(":3000")
}
