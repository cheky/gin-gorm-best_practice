package website

import (
	"apps_barang/libraries"
	"apps_barang/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = os.Getenv("JWT-APP-NAME")
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT-SIGNATURE-KEY"))

type Form_login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var messages_form_login = []map[string]interface{}{
	{
		"field": "username",
		"messages": map[string]interface{}{
			"required": "Username harus diisi.",
		},
	}, {
		"field": "password",
		"messages": map[string]interface{}{
			"required": "Password harus diisi.",
		},
	},
}

func Login(c *gin.Context) {
	body := Form_login{}
	if err := c.ShouldBind(&body); err != nil {
		validate := libraries.ValidationRules(messages_form_login, err)
		libraries.StatusBadRequest(c, validate)
		return
	}
	//mencari jumlah data berdasarkan username
	c_user := models.CountUser(map[string]interface{}{
		"user.username=?": body.Username,
	})
	//jika jumlah username kurang dari 1
	if c_user < 1 {
		libraries.StatusNotFound(c, gin.H{
			"username": "Username tidak terdaftar.",
		})
		//jika jumlah username lebih dari 0
	} else {
		//menampilkan user berdasarkan username
		var R_user struct {
			Kd_user  int64  `json:"kd_user"`
			Nm_user  string `json:"nm_user"`
			Salt     string `json:"salt"`
			Password string `json:"password"`
		}
		find_user := models.FindUser([]string{
			"user.kd_user",
			"user.nm_user",
			"user.salt",
			"user.password",
		}, map[string]interface{}{
			"user.username=?": body.Username,
		})
		find_user.Scan(&R_user)
		input_password := libraries.GetPassword(R_user.Salt, body.Password)
		//jika password tidak sama dengan password input
		if input_password != R_user.Password {
			libraries.StatusBadRequest(c, gin.H{
				"password": "Password yang anda masukkan salah.",
			})
		} else {
			libraries.StatusOk(c, "Login berhasil")
		}
	}
}
