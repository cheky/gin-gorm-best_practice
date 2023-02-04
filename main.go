package main

import (
	"apps_barang/config"
	"apps_barang/controllers/website"
	"apps_barang/libraries"
	"apps_barang/models"
	"net/http"
	"os"
	"runtime"

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
	//use max processor
	runtime.GOMAXPROCS(runtime.NumCPU())
	//init gin
	r := gin.Default()
	//auto recovery
	r.Use(gin.Recovery())
	//dum log request and response
	r.Use(libraries.SysLog())
	r.LoadHTMLGlob("views/*.html")
	r.Static("assets", "./views/assets/")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/mysql", func(c *gin.Context) {
		config.InitMysql()
		c.JSON(200, gin.H{
			"messages": "Connection MySQL complete",
		})
	})
	r.GET("/automigrate", func(c *gin.Context) {
		Mysql := config.InitMysql()
		Mysql.AutoMigrate(&models.Brg_kat{})
		Mysql.AutoMigrate(&models.Brg{})
		Mysql.AutoMigrate(&models.User{})
		c.JSON(200, gin.H{
			"messages": "Auto Migrate Mysql Database Successfully",
		})
	})
	web := r.Group("/website")
	web.POST("/login", website.Login)

	web.POST("/brg_kat", website.Add_brg_kat)
	web.PUT("/brg_kat/:kd_kat", website.Edit_brg_kat)
	web.DELETE("/brg_kat/:kd_kat", website.Delete_brg_kat)
	web.GET("/brg_kat", website.Find_brg_kat)

	web.POST("/brg", website.Add_brg)
	web.PUT("/brg/:kd_brg", website.Edit_brg)
	web.DELETE("/brg/:kd_brg", website.Delete_brg)
	web.GET("/brg", website.Find_brg)
	web.GET("/brg/datatables", website.Datatables_brg)
	web.POST("/brg/upload", website.Upload_foto)
	return r
}
