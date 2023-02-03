package website

import (
	"apps_barang/libraries"
	"apps_barang/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Form_brg struct {
	Nm_brg string `form:"nm_brg" json:"nm_brg" binding:"required"`
	Kd_kat int    `form:"kd_kat" json:"kd_kat" binding:"required"`
	Aktif  string `form:"aktif" json:"aktif" binding:"required"`
}
type Uri_brg struct {
	Kd_brg int `uri:"kd_brg" binding:"required"`
}

var messages_form_brg = []map[string]interface{}{
	{
		"field": "nm_brg",
		"messages": map[string]interface{}{
			"required": "Nama barang harus diisi",
		},
	},
	{
		"field": "kd_kat",
		"messages": map[string]interface{}{
			"required": "Kode kategori barang harus diisi",
		},
	},
	{
		"field": "aktif",
		"messages": map[string]interface{}{
			"required": "Status kategori harus diisi",
		},
	},
}
var messages_uri_brg = []map[string]interface{}{
	{
		"field": "kd_brg",
		"messages": map[string]interface{}{
			"required": "Kode barang harus diisi",
		},
	},
}

func Add_brg(c *gin.Context) {
	body := Form_brg{}
	if err := c.ShouldBind(&body); err != nil {
		validate := libraries.ValidationRules(messages_form_brg, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	//mencari jumlah data berdasarkan nm_brg dan kd_kat
	c_brg := models.CountBrg(map[string]interface{}{
		"brg.nm_brg=?": body.Nm_brg,
		"brg.kd_kat=?": body.Kd_kat,
	})
	//jika jumlah data lebih dari 0
	if c_brg > 0 {
		libraries.StatusBadRequest(c, gin.H{
			"nm_brg": "Nama barang sudah digunakan pada kategori yang anda pilih",
		})
		return
	}
	input := models.Brg{
		Nm_brg: body.Nm_brg,
		Kd_kat: body.Kd_kat,
		Aktif:  body.Aktif,
	}
	insert, _ := models.InsertBrg(&input)
	if insert.RowsAffected > 0 {
		libraries.StatusCreated(c, "Data barang dengan nama "+input.Nm_brg+", berhasil disimpan.")
	} else {
		libraries.StatusInternalServerError(c, insert.Error)
	}
}
func Delete_brg(c *gin.Context) {
	uri := Uri_brg{}
	if err := c.ShouldBindUri(&uri); err != nil {
		validate := libraries.ValidationRules(messages_uri_brg, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	//mencari jumlah data berdasarkan kd_brg
	c_brg := models.CountBrg(map[string]interface{}{
		"brg.kd_brg=?": uri.Kd_brg,
	})
	//jika jumlah data sama dengan 0
	if c_brg == 0 {
		libraries.StatusNotFound(c, gin.H{
			"kd_brg": "Kode barang tidak ditemukan.",
		})
		return
	}
	delete := models.DeleteBrg(map[string]interface{}{
		"kd_brg=?": uri.Kd_brg,
	})
	if delete.RowsAffected > 0 {
		libraries.StatusNoContent(c, "Data barang berhasil dihapus.")
	} else {
		libraries.StatusInternalServerError(c, delete.Error)
	}
}
func Edit_brg(c *gin.Context) {
	body := Form_brg{}
	if err := c.ShouldBind(&body); err != nil {
		validate := libraries.ValidationRules(messages_form_brg, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	uri := Uri_brg{}
	if err := c.ShouldBindUri(&uri); err != nil {
		validate := libraries.ValidationRules(messages_uri_brg, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	//mencari jumlah data berdasarkan kd_brg
	c_brg := models.CountBrg(map[string]interface{}{
		"brg.kd_brg=?": uri.Kd_brg,
	})
	//jika jumlah data sama dengan 0
	if c_brg == 0 {
		libraries.StatusNotFound(c, gin.H{
			"kd_brg": "Kode barang tidak ditemukan.",
		})
		return
	}
	input := models.Brg{
		Nm_brg: body.Nm_brg,
		Kd_kat: body.Kd_kat,
		Aktif:  body.Aktif,
	}
	update := models.UpdateBrg(&input, map[string]interface{}{
		"kd_brg=?": uri.Kd_brg,
	})
	if update.RowsAffected > 0 {
		libraries.StatusOk(c, "Data barang berhasil diperbaharui.")
	} else {
		errors := update.Error
		if errors != nil {
			libraries.StatusInternalServerError(c, errors)
		} else {
			libraries.StatusBadRequest(c, "Tidak terdeteksi adanya perubahan data")
		}

	}
}

func Find_brg(c *gin.Context) {
	result := []map[string]interface{}{}
	selector := models.FindBrg([]string{
		"brg.kd_brg",
		"brg.nm_brg",
		"brg_kat.nm_kat",
		"brg.aktif",
		"DATE_FORMAT(brg.on_create,'%Y-%m-%d %H:%i:%s') AS on_create",
		"DATE_FORMAT(brg.on_update,'%Y-%m-%d %H:%i:%s') AS on_update",
	}, map[string]interface{}{
		"brg.aktif='Y'": false,
	})
	selector.Order("brg.kd_brg DESC")
	selector.Scan(&result)
	libraries.StatusOk(c, gin.H{
		"brg": result,
	})
}
func Datatables_brg(c *gin.Context) {
	_draw, _start, _length, _search, _order_column, _order_dir := libraries.DatatableInit(c)
	kd_kat := libraries.DatatableSearch(c, 0)
	where := map[string]interface{}{}
	if len(kd_kat) > 0 {
		where["brg.kd_kat"] = kd_kat
	}
	if len(_search) > 0 {
		where["brg.nm_brg LIKE '%"+_search+"%'"] = false
	}

	//mencari jumlah data berdasarkan kd_brg
	c_brg := models.CountBrg(where)
	result := []map[string]interface{}{}
	selector := models.FindBrg([]string{
		"brg.kd_brg",
		"brg.nm_brg",
		"brg_kat.nm_kat",
		"brg.aktif",
		"DATE_FORMAT(brg.on_create,'%Y-%m-%d %H:%i:%s') AS on_create",
		"DATE_FORMAT(brg.on_update,'%Y-%m-%d %H:%i:%s') AS on_update",
	}, where)
	selector.Limit(_length).Offset(_start)
	selector.Order(_order_column + " " + _order_dir)
	selector.Scan(&result)
	c.JSON(http.StatusOK, gin.H{
		"draw":            _draw,
		"recordsTotal":    c_brg,
		"recordsFiltered": c_brg,
		"data":            result,
	})

}
func Upload_foto(c *gin.Context) {
	file, err := c.FormFile("foto")
	// The file cannot be received.
	if err != nil {
		libraries.StatusBadRequest(c, gin.H{
			"foto": "File foto harus diisi",
		})
		return
	}
	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	if extension != ".png" && extension != ".jpg" {
		libraries.StatusBadRequest(c, "Upload foto gagal, Ekstensi file yang dijinkan hanya png dan jpg")
		return
	}
	if file.Size/1024 > 100 {
		libraries.StatusBadRequest(c, "Upload foto gagal, Ukuran maksimal file adalah 100KB")
		return
	}
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := "Jamali" + extension
	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, newFileName); err != nil {
		libraries.StatusBadRequest(c, err)
		return
	}
	//upload file to owncloud
	er_upload := libraries.OwncloudUpload("teralink/"+newFileName, newFileName)
	if er_upload != nil {
		libraries.StatusBadRequest(c, er_upload)
		return
	}
	//mkdir to owncloud
	// er_mkdir := libraries.OwncloudMkdir("/teralink/doc")
	// if er_mkdir != nil {
	// 	libraries.StatusBadRequest(c, er_mkdir)
	// 	return
	// }
	//delete file in owncloud
	// err_delete := libraries.OwncloudDelete("teralink/" + newFileName)
	// if err_delete != nil {
	// 	libraries.StatusBadRequest(c, err_delete)
	// 	return
	// }
	// File saved successfully. Return proper result
	libraries.StatusOk(c, "File foto berhasil diupload")
}
