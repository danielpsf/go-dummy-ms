package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/danielpsf/go-dummy-ms/config/envvars"
	"github.com/danielpsf/go-dummy-ms/config/routes"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
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
	logrus.Fields["service"] = "dummy-ms"
}

func main() {
	showStartupLogs()
	configureWebServer()
}

func showStartupLogs() {
	var openFiles syscall.Rlimit

	logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("--------------------------------------")
	logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("------------STARTING dummy-ms------------")
	logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("Log Level: " + envvars.LogLevel)
	logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("Server Port: " + envvars.ServerPort)
	logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("URL Scheme: " + envvars.Scheme)
	logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("AWS ENV: " + envvars.Env)

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &openFiles)
	if err == nil {
		logrus.WithFields(logrus.Fields{"service": "go-dummy-ms"}).Info("Rlimit: " + strconv.FormatUint(openFiles.Cur, 10) + " " + strconv.FormatUint(openFiles.Max, 10))
	}

	logrus.WithFields(logrus.Fields{"service": "dummy-ms"}).Info("--------------------------------------")
}

func configureWebServer() {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 200

	r := routes.Setup()

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Authorization", "X-XSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PATCH", "PUT", "DELETE"},
	}).Handler(r)

	err := http.ListenAndServe(":"+envvars.ServerPort, handlers.RecoveryHandler()(handler))
	if err != nil {
		logrus.WithFields(logrus.Fields{"service": "dummy-ms"}).Error("HTTP Server panic. | " + err.Error())
	}
}
