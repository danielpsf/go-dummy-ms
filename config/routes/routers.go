package routes

import (
	"net/http"

	"github.com/danielpsf/go-dummy-ms/status"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

var routes = []route{
	{
		Name:        "Status",
		Pattern:     "/status",
		Method:      http.MethodGet,
		HandlerFunc: status.Check,
	},
}

func Setup() *mux.Router {
	muxRouter := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = muxLogger(handler, route.Name, route.Pattern)

		muxRouter.Path(route.Pattern).Name(route.Name).Handler(handler).Methods(route.Method)
	}

	return muxRouter
}

func muxLogger(inner http.Handler, name string, pattern string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"service": "dummy-ms",
		}).Infof("dummy-ms started: %v - %v to %v - %v", r.Method, name, r.RequestURI, pattern)

		inner.ServeHTTP(w, r)
	})
}
