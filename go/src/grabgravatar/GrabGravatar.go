package grabgravatar
// credits: https://golangcode.com/download-a-file-from-a-url/

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(email string, width string) error {
	md5 := getMd5FromMail(email)
	url := fmt.Sprintf("https://gravatar.com/avatar/%x?s=%s", md5, width)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(email + "_" + width + ".png")
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func getMd5FromMail(mail string) [16]byte {
	return md5.Sum([]byte(mail))
}
