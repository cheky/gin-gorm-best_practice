package main

import (
	"apps_barang/config"
	"apps_barang/controllers"
	"apps_barang/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load .env file")
	}
	//init app
	r := setupRouter()
	mode := os.Getenv("MODE")
	app_port := os.Getenv("LIVE-PORT")
	if mode == "sandbox" {
		app_port = os.Getenv("SANDBOX-PORT")
	}
	_ = r.Run(":" + app_port)
}

// app route
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messages": "pong",
		})
	})
	r.GET("/mysql", func(c *gin.Context) {
		config.InitMysql()
		c.JSON(200, gin.H{
			"messages": "Connect Mysql Database Successfully",
		})
	})
	r.GET("/automigrate", func(c *gin.Context) {
		Mysql := config.InitMysql()
		Mysql.AutoMigrate(&models.Brg_kat{})
		c.JSON(200, gin.H{
			"messages": "Auto Migrate Mysql Database Successfully",
		})
	})
	web := r.Group("/web")
	web.POST("/brg_kat", controllers.Add_brg_kat)
	web.PUT("/brg_kat/:kd_kat", controllers.Edit_brg_kat)
	web.DELETE("/brg_kat/:kd_kat", controllers.Delete_brg_kat)
	web.GET("/brg_kat", controllers.Find_brg_kat)
	return r
}
