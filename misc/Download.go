package misc

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func DownloadFile(filepath, filename, url string) (err error) {
	// Create the file
	out, err := os.Create(fmt.Sprintf("%s\\%s", filepath, filename))
	if err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Error("DownloadFile create")
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Error("DownloadFile get")
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		Log().WithFields(logrus.Fields{
			"status": resp.Status,
		}).Error("DownloadFile StatusCode")
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Error("DownloadFile copy")
		return err
	}

	return nil
}
