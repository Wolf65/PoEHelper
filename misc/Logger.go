package misc

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

func Log() *logrus.Logger {
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,

		DisableLevelTruncation: true,
		PadLevelText:           true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
	})
	// Logger.SetOutput(os.Stdout) // https://www.golinuxcloud.com/golang-logrus/
	Logger.SetLevel(logrus.DebugLevel)
	// Logger.SetReportCaller(true)
	return Logger
}

/*
f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
		panic(err)
	}
	defer f.Close()

	log := &logrus.Logger{
                // Log into f file handler and on os.Stdout
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}


misc.Log().WithFields(logrus.Fields{
				"err": err,
			}).Error("")

*/
