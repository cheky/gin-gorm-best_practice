package libraries

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
func OwnloadUpload(local_source, owncloud_destination_path string) error {
	init := OwncloudInit()
	client := &http.Client{}
	var data = strings.NewReader("@" + owncloud_destination_path)
	req, err := http.NewRequest("PUT", init["webdav"]+owncloud_destination_path, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(init["username"], init["password"])
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}
