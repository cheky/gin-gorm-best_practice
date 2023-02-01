package libraries

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DatatableInit(c *gin.Context) (int, int, int, string, string, string) {
	query := c.Request.URL.Query()
	var _draw, _start, _length int
	var _search, _order_column, _order_dir string
	for k, v := range query {
		if k == "draw" {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println("Error konversi string to int64 DatatablesInit (draw)", err)
			}
			_draw = i
		}
		if k == "start" {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println("Error konversi string to int64 DatatablesInit (start)", err)
			}
			_start = i
		}
		if k == "length" {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println("Error konversi string to int64 DatatablesInit (length)", err)
			}
			_length = i
		}
		if k == "search[value]" {
			_search = v[0]
		}
		if k == "order[0][column]" {
			_order_column = query["columns["+v[0]+"][data]"][0]
		}
		if k == "order[0][dir]" {
			_order_dir = v[0]
		}
	}
	return _draw, _start, _length, _search, _order_column, _order_dir
}
func DatatableSearch(c *gin.Context, column int) string {
	query := c.Request.URL.Query()
	return query["columns["+strconv.Itoa(column)+"][search][value]"][0]
}
