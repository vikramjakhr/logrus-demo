package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {

	// creating new loggers
	logger := logrus.New()

	// setting level
	logger.SetLevel(logrus.DebugLevel)

	// If you wish to add the calling method as a field, instruct the logger via:
	logger.SetReportCaller(true)

	// setting formatter
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	log = logger.WithFields(logrus.Fields{
		"project": "test-project",
		"cluster": "mycluster",
	})

	//	append logs to a file
	// file, err := os.OpenFile("./logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	logger.SetOutput(file)
	// } else {
	// 	logger.Info("Failed to log to file, using default stderr")
	// }
}

func main() {

	// Info
	log.WithFields(logrus.Fields{
		"package":  "api",
		"function": "CreateProject",
		"data":     "mydata",
	}).Info("A sample info message")

	// Warning
	log.WithFields(logrus.Fields{
		"package":  "api",
		"function": "CreateProject",
		"data":     "mydata",
	}).Warning("A sample warning message")

	// Debug
	log.WithFields(logrus.Fields{
		"package":  "api",
		"function": "CreateProject",
		"data":     "mydata",
	}).Debug("A sample debug message")

	// Error
	log.WithFields(logrus.Fields{
		"package":  "api",
		"function": "CreateProject",
		"data":     "mydata",
		"error":    errors.New("Hola!!!"),
	}).Error("A sample error message")
}
