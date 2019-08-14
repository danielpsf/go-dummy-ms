package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/danielpsf/go-dummy-ms/configuration/envvars"
	"github.com/danielpsf/go-dummy-ms/configuration/routes"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "asctime",
			logrus.FieldKeyLevel: "level_name",
			logrus.FieldKeyMsg:   "message",
		},
	})

	logrus.SetOutput(os.Stdout)
	logLevel, _ := logrus.ParseLevel(strings.ToLower(envvars.LogLevel))
	logrus.SetLevel(logLevel)
}

func main() {
	showStartupLogs()
	configureWebServer()
}

func showStartupLogs() {
	var openFiles syscall.Rlimit

	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("--------------------------------------")
	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("------------STARTING dummy-mw------------")
	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("Log Level: " + envvars.LogLevel)
	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("Server Port: " + envvars.ServerPort)
	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("URL Scheme: " + envvars.Scheme)
	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("AWS ENV: " + envvars.Env)

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &openFiles)
	if err == nil {
		logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("Rlimit: " + strconv.FormatUint(openFiles.Cur, 10) + " " + strconv.FormatUint(openFiles.Max, 10))
	}

	logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Info("--------------------------------------")
}

func configureWebServer() {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 200

	r := routes.Setup()

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Authorization", "X-XSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PATCH", "PUT", "DELETE", "COPY"},
	}).Handler(r)

	err := http.ListenAndServe(":"+envvars.ServerPort, handlers.RecoveryHandler()(handler))
	if err != nil {
		logrus.WithFields(logrus.Fields{"service": "dummy-mw"}).Error("HTTP Server panic. | " + err.Error())
	}
}
