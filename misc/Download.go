package misc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"poehelper/config"

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
		}).Error("DownloadFile statusCode")
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

func CheckAndDownloadFont() {
	// JetBrainsMono-Medium.ttf
	if _, err := os.Stat(config.App.Vars.FontDirectory + "JetBrainsMono-Medium.ttf"); err != nil {
		if os.IsNotExist(err) {
			Log().Warn("JetBrainsMono-Medium.ttf does not exist")
			DownloadFile(config.App.Vars.FontDirectory, "JetBrainsMono-Medium.ttf", "https://github.com/Wolf65/PoEHelper/raw/main/fonts/ttf/JetBrainsMono-Medium.ttf")
			Log().Info("JetBrainsMono-Medium.ttf download")
		} else {
			Log().WithFields(logrus.Fields{
				"err": err,
			}).Warn("JetBrainsMono-Medium.ttf")
		}
	} else {
		Log().Debug("JetBrainsMono-Medium.ttf exists")
	}
	// fa-solid-900.ttf
	if _, err := os.Stat(config.App.Vars.FontDirectory + "fa-solid-900.ttf"); err != nil {
		if os.IsNotExist(err) {
			Log().Warn("fa-solid-900.ttf does not exist")
			DownloadFile(config.App.Vars.FontDirectory, "fa-solid-900.ttf", "https://github.com/Wolf65/PoEHelper/raw/main/fonts/ttf/fa-solid-900.ttf")
			Log().Info("fa-solid-900.ttf download")
		} else {
			Log().WithFields(logrus.Fields{
				"err": err,
			}).Warn("fa-solid-900.ttf")
		}
	} else {
		Log().Debug("fa-solid-900.ttf exists")
	}
	// fa-brands-400.ttf
	if _, err := os.Stat(config.App.Vars.FontDirectory + "fa-brands-400.ttf"); err != nil {
		if os.IsNotExist(err) {
			Log().Warn("fa-brands-400.ttf does not exist")
			DownloadFile(config.App.Vars.FontDirectory, "fa-brands-400.ttf", "https://github.com/Wolf65/PoEHelper/raw/main/fonts/ttf/fa-brands-400.ttf")
			Log().Info("fa-brands-400.ttf download")
		} else {
			Log().WithFields(logrus.Fields{
				"err": err,
			}).Warn("fa-brands-400.ttf")
		}
	} else {
		Log().Debug("fa-brands-400.ttf exists")
	}
}
