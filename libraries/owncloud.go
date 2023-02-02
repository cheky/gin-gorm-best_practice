package libraries

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func OwncloudInit() map[string]string {
	mode := os.Getenv("MODE")
	if mode == "sandbox" {
		return map[string]string{
			"username": os.Getenv("OC_USERNAME_SANDBOX"),
			"password": os.Getenv("OC_PASSWORD_SANDBOX"),
			"webdav":   os.Getenv("OC_WEBDAV_SANDBOX"),
			"path":     os.Getenv("OC_PATH_SANDBOX"),
		}
	} else {
		return map[string]string{
			"username": os.Getenv("OC_USERNAME_LIVE"),
			"password": os.Getenv("OC_PASSWORD_LIVE"),
			"webdav":   os.Getenv("OC_WEBDAV_LIVE"),
			"path":     os.Getenv("OC_PATH_LIVE"),
		}
	}
}

// Error type encapsulates the returned error messages from the
// server.
type Error struct {
	// Exception contains the type of the exception returned by
	// the server.
	Exception string `xml:"exception"`

	// Message contains the error message string from the server.
	Message string `xml:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Exception: %s, Message: %s", e.Exception, e.Message)
}

type ShareElement struct {
	Id  uint   `xml:"id"`
	Url string `xml:"url"`
}

type ShareResult struct {
	XMLName    xml.Name       `xml:"ocs"`
	Status     string         `xml:"meta>status"`
	StatusCode uint           `xml:"meta>statuscode"`
	Message    string         `xml:"meta>message"`
	Id         uint           `xml:"data>id"`
	Url        string         `xml:"data>url"`
	Elements   []ShareElement `xml:"data>element"`
}

// Upload uploads the specified source to the specified destination
// path on the cloud.
func OwncloudUpload(dest, src string) error {
	_src, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	_, er := sendWebDavRequest("PUT", dest, _src)
	return er
}

// Mkdir creates a new directory on the cloud with the specified name.
func OwncloudMkdir(path string) error {
	_, err := sendWebDavRequest("MKCOL", path, nil)
	return err

}

// Delete removes the specified folder from the cloud.
func OwncloudDelete(path string) error {
	_, err := sendWebDavRequest("DELETE", path, nil)
	return err
}

// Download downloads a file from the specified path.
func OwncloudDownload(path string) ([]byte, error) {
	return sendWebDavRequest("GET", path, nil)
}

func OwncloudExists(path string) bool {
	_, err := sendWebDavRequest("PROPFIND", path, nil)
	return err == nil
}

func sendWebDavRequest(request string, path string, data []byte) ([]byte, error) {
	// Create the https request
	init := OwncloudInit()
	webdavPath := init["webdav"]

	folderUrl, err := url.Parse(webdavPath + path)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	url := &url.URL{}
	req, err := http.NewRequest(request, url.ResolveReference(folderUrl).String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(init["username"], init["password"])

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) > 0 {
		if body[0] == '<' {
			error := Error{}
			err = xml.Unmarshal(body, &error)
			if err != nil {
				return body, err
			}
			if error.Exception != "" {
				return nil, err
			}
		}

	}

	return body, nil
}
