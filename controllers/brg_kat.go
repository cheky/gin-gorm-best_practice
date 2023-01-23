package controllers

import (
	"apps_barang/config"
	"apps_barang/libraries"
	"apps_barang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add_brg_kat(c *gin.Context) {

	type Input_add_brg_kat struct {
		Nm_kat string `form:"nm_kat" json:"nm_kat" binding:"required"`
		Aktif  string `form:"aktif" json:"aktif" binding:"required"`
	}
	body := Input_add_brg_kat{}
	if err := c.ShouldBind(&body); err != nil {
		var messages = []map[string]interface{}{
			{
				"field": "nm_kat",
				"messages": map[string]interface{}{
					"required": "Nama kategori harus diisi",
				},
			},
			{
				"field": "aktif",
				"messages": map[string]interface{}{
					"required": "Status kategori harus diisi",
				},
			},
		}
		validate := libraries.ValidationRules(messages, err)
		c.JSON(http.StatusBadRequest, libraries.StatusBadRequest(validate))
		return
	}
	brg_kat := models.Brg_kat{Nm_kat: body.Nm_kat, Aktif: body.Aktif}
	Mysql := config.InitMysql()
	insert := Mysql.Create(&brg_kat)
	if insert.RowsAffected > 0 {
		c.JSON(http.StatusCreated, libraries.StatusCreated(&brg_kat))
	} else {
		c.JSON(http.StatusInternalServerError, libraries.StatusInternalServerError(insert.Error))
	}
}
func Edit_brg_kat(c *gin.Context) {
	type Input_add_brg_kat struct {
		Nm_kat string `form:"nm_kat" json:"nm_kat" binding:"required"`
		Aktif  string `form:"aktif" json:"aktif" binding:"required"`
	}
	type Query_add_brg_kat struct {
		Kd_kat int `uri:"kd_kat" binding:"required"`
	}
	body := Input_add_brg_kat{}
	if err := c.ShouldBind(&body); err != nil {
		var messages = []map[string]interface{}{
			{
				"field": "nm_kat",
				"messages": map[string]interface{}{
					"required": "Nama kategori harus diisi",
				},
			},
			{
				"field": "aktif",
				"messages": map[string]interface{}{
					"required": "Status kategori harus diisi",
				},
			},
		}
		validate := libraries.ValidationRules(messages, err)
		c.JSON(http.StatusBadRequest, libraries.StatusBadRequest(validate))
		return
	}
	query := Query_add_brg_kat{}
	if err := c.ShouldBindUri(&query); err != nil {
		var messages = []map[string]interface{}{
			{
				"field": "kd_kat",
				"messages": map[string]interface{}{
					"required": "Kode kategori harus diisi",
				},
			},
		}
		validate := libraries.ValidationRules(messages, err)
		c.JSON(http.StatusBadRequest, libraries.StatusBadRequest(validate))
		return
	}
	brg_kat := models.Brg_kat{}
	Mysql := config.InitMysql()
	var count int64
	Mysql.Model(&brg_kat).Where("kd_kat", query.Kd_kat).Count(&count)
	if count > 0 {
		update := Mysql.Model(&brg_kat).Where("kd_kat", query.Kd_kat).Updates(map[string]interface{}{
			"nm_kat": body.Nm_kat,
			"aktif":  body.Aktif,
		})
		if update.RowsAffected > 0 {
			c.JSON(http.StatusOK, libraries.StatusOk(&brg_kat))
		} else {
			c.JSON(http.StatusBadRequest, libraries.StatusBadRequest("Tidak terdeteksi adanya perubahan data"))
		}
	} else {
		c.JSON(http.StatusNotFound, libraries.StatusNotFound(map[string]interface{}{
			"kd_kat": "Kode kategori tidak ditemukan",
		}))
	}
}
func Delete_brg_kat(c *gin.Context) {
	type Query_brg_kat struct {
		Kd_kat int `uri:"kd_kat" binding:"required,number"`
	}
	query := Query_brg_kat{}
	if err := c.ShouldBindUri(&query); err != nil {
		var messages = []map[string]interface{}{
			{
				"field": "kd_kat",
				"messages": map[string]interface{}{
					"required": "Kode kategori harus diisi",
					"number":   "Kode kategori harus diisi dengan angka",
				},
			},
		}
		validate := libraries.ValidationRules(messages, err)
		c.JSON(http.StatusBadRequest, libraries.StatusBadRequest(validate))
		return
	}
	brg_kat := models.Brg_kat{}
	Mysql := config.InitMysql()
	var count int64
	Mysql.Model(&brg_kat).Where("kd_kat", query.Kd_kat).Count(&count)
	if count > 0 {
		delete := Mysql.Where("kd_kat=?", query.Kd_kat).Delete(&brg_kat)
		if delete.RowsAffected > 0 {
			c.JSON(http.StatusNoContent, libraries.StatusNoContent(&brg_kat))
		} else {
			c.JSON(http.StatusBadRequest, libraries.StatusBadRequest("Tidak terdeteksi adanya penghapusan data"))
		}
	} else {
		c.JSON(http.StatusNotFound, libraries.StatusNotFound(map[string]interface{}{
			"kd_kat": "Kode kategori tidak ditemukan",
		}))
	}
}
func Find_brg_kat(c *gin.Context) {
	var brg_kat []models.Brg_kat
	Mysql := config.InitMysql()
	Mysql.Find(&brg_kat)
	c.JSON(http.StatusOK, libraries.StatusOk(map[string]interface{}{
		"brg_kat": brg_kat,
	}))
}
