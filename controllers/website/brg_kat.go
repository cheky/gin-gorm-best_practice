package website

import (
	"apps_barang/config"
	"apps_barang/libraries"
	"apps_barang/models"

	"github.com/gin-gonic/gin"
)

type Form_brg_kat struct {
	Nm_kat string `form:"nm_kat" json:"nm_kat" binding:"required"`
	Aktif  string `form:"aktif" json:"aktif" binding:"required"`
}
type Uri_brg_kat struct {
	Kd_kat int `uri:"kd_kat" binding:"required"`
}

var messages_form_brg_kat = []map[string]interface{}{
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
var messages_uri_brg_kat = []map[string]interface{}{
	{
		"field": "kd_kat",
		"messages": map[string]interface{}{
			"required": "Kode kategori harus diisi",
		},
	},
}

func Add_brg_kat(c *gin.Context) {
	body := Form_brg_kat{}
	if err := c.ShouldBind(&body); err != nil {
		validate := libraries.ValidationRules(messages_form_brg_kat, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	brg_kat := models.Brg_kat{Nm_kat: body.Nm_kat, Aktif: body.Aktif}
	insert, response := models.InsertBrg_kat(&brg_kat)
	if insert.RowsAffected > 0 {
		libraries.StatusCreated(c, response)
	} else {
		libraries.StatusInternalServerError(c, insert.Error)
	}
}
func Edit_brg_kat(c *gin.Context) {
	body := Form_brg_kat{}
	if err := c.ShouldBind(&body); err != nil {
		validate := libraries.ValidationRules(messages_form_brg_kat, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	uri := Uri_brg_kat{}
	if err := c.ShouldBindUri(&uri); err != nil {
		validate := libraries.ValidationRules(messages_uri_brg_kat, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	brg_kat := models.Brg_kat{}
	Mysql := config.InitMysql()
	var count int64
	Mysql.Model(&brg_kat).Where("kd_kat", uri.Kd_kat).Count(&count)
	if count > 0 {
		update := Mysql.Model(&brg_kat).Where("kd_kat", uri.Kd_kat).Updates(map[string]interface{}{
			"nm_kat": body.Nm_kat,
			"aktif":  body.Aktif,
		})
		if update.RowsAffected > 0 {
			libraries.StatusOk(c, &brg_kat)
		} else {
			libraries.StatusBadRequest(c, "Tidak terdeteksi adanya perubahan data")
		}
	} else {
		libraries.StatusNotFound(c, gin.H{
			"kd_kat": "Kode kategori tidak ditemukan",
		})
	}
}
func Delete_brg_kat(c *gin.Context) {
	uri := Uri_brg_kat{}
	if err := c.ShouldBindUri(&uri); err != nil {
		validate := libraries.ValidationRules(messages_uri_brg_kat, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	brg_kat := models.Brg_kat{}
	Mysql := config.InitMysql()
	var count int64
	Mysql.Model(&brg_kat).Where("kd_kat", uri.Kd_kat).Count(&count)
	if count > 0 {
		delete := Mysql.Where("kd_kat=?", uri.Kd_kat).Delete(&brg_kat)
		if delete.RowsAffected > 0 {
			libraries.StatusNoContent(c, &brg_kat)
		} else {
			libraries.StatusBadRequest(c, "Tidak terdeteksi adanya penghapusan data")
		}
	} else {
		libraries.StatusNotFound(c, gin.H{
			"kd_kat": "Kode kategori tidak ditemukan",
		})
	}
}
func Find_brg_kat(c *gin.Context) {
	var brg_kat []models.Brg_kat
	Mysql := config.InitMysql()
	Mysql.Preload("Brg").Find(&brg_kat)
	libraries.StatusOk(c, gin.H{
		"brg_kat": brg_kat,
	})
}
