package misc

import (
	"os"
	"poehelper/config"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

func SaveConfig() {
	fileConfig, err := os.OpenFile("config.toml", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Warning("Open config.toml")
	}
	if err := toml.NewEncoder(fileConfig).Encode(config.App); err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Warning("Encode config.toml")
	}
	if err := fileConfig.Close(); err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Warning("Close config.toml file")

	}

}

func LoadConfig() {
	_, err := os.OpenFile("config.toml", os.O_RDONLY, 0755)
	if err != nil {
		Log().WithFields(logrus.Fields{
			"err": err,
		}).Error("Open config.toml")
	}
}
