package libraries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SysLog() gin.HandlerFunc {
	logged := make(map[string]interface{})
	return func(ctx *gin.Context) {
		var strB strings.Builder
		//dump req path
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		if query != "" {
			path += "?" + query
		}
		logged["Request-Path"] = path
		//dump ip client
		logged["Client-IP"] = string(ctx.ClientIP())
		//dump request datetime
		start := time.Now()
		logged["Request-Datetime"] = start.Format("2006-01-02 15:04:05.000")
		//dump req method
		logged["Request-Method"] = ctx.Request.Method
		//dump req header
		logged["Request-Header"] = ctx.Request.Header

		//dump req body
		if ctx.Request.ContentLength > 0 {
			buf, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				strB.WriteString(fmt.Sprintf("\nread bodyCache err \n %s", err.Error()))
				goto DumpRes
			}
			rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
			ctGet := ctx.Request.Header.Get("Content-Type")
			ct, _, err := mime.ParseMediaType(ctGet)
			if err != nil {
				strB.WriteString(fmt.Sprintf("\ncontent_type: %s parse err \n %s", ctGet, err.Error()))
				goto DumpRes
			}

			switch ct {
			case gin.MIMEJSON:
				bts, err := ioutil.ReadAll(rdr)
				if err != nil {
					strB.WriteString(fmt.Sprintf("\nread rdr err \n %s", err.Error()))
					goto DumpRes
				}
				logged["Request-Body"] = string(bts)
			case gin.MIMEPOSTForm:
				bts, err := ioutil.ReadAll(rdr)
				if err != nil {
					strB.WriteString(fmt.Sprintf("\nread rdr err \n %s", err.Error()))
					goto DumpRes
				}
				val, err := url.ParseQuery(string(bts))
				if err != nil {
					strB.WriteString(fmt.Sprintf("\nparse req body err \n" + err.Error()))
					goto DumpRes
				}
				logged["Request-Body"] = val

			case gin.MIMEMultipartPOSTForm:
			default:
			}
		}

	DumpRes:
		ctx.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Next()
		//dump response code
		logged["Response-Code"] = ctx.Writer.Status()
		//dump response datetime
		end := time.Now()
		logged["Response-Datetime"] = end.Format("2006-01-02 15:04:05.000")
		//dump res header
		logged["Response-Header"] = ctx.Writer.Header()
		bw, ok := ctx.Writer.(*bodyWriter)
		if !ok {
			strB.WriteString("\nbodyWriter was override , can not read bodyCache")
			goto End
		}

		//dump res body
		if bodyAllowedForStatus(ctx.Writer.Status()) && bw.bodyCache.Len() > 0 {
			ctGet := ctx.Writer.Header().Get("Content-Type")
			ct, _, err := mime.ParseMediaType(ctGet)
			if err != nil {
				strB.WriteString(fmt.Sprintf("\ncontent-type: %s parse  err \n %s", ctGet, err.Error()))
				goto End
			}
			switch ct {
			case gin.MIMEJSON:
				logged["Response-Body"] = string(bw.bodyCache.Bytes())
			case gin.MIMEHTML:
			default:
			}
		}

	End:
		json, err := json.Marshal(logged)
		if err != nil {
			strB.WriteString(fmt.Sprintf("\nmap to json logged err \n" + err.Error()))
			fmt.Println(strB.String())
		}
		fmt.Println(strB.String())
		fmt.Println(string(json))
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

// rewrite Write()
func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}
