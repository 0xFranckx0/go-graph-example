package socnet

// This logging  module is aimed to log in json format and be used by http
// handler functions. In context of massive logging, this is useful to store
// logs in json format into a backend such as ElasticSearch
// In mux terminolgy this module act as a middleware
import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

type HandlerFunc func(http.ResponseWriter, *http.Request, *logrus.Entry) (int, string, []byte, string)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, l *logrus.Entry) (int, string, []byte, string) {
	return f(w, r, l)
}

// Middleware used by calling routeHandlers, it logs in JSON format and send
// HTTP response to the client.
func logger(h HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logFields := logrus.Fields{
			"method":    r.Method,
			"url":       r.URL.Path,
			"remotaddr": r.Header.Get("X-Real-IP"),
			"handler":   runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()}
		for key, val := range mux.Vars(r) {
			logFields[key] = val
		}
		log := logrus.WithFields(logFields)

		start := time.Now()
		code, contentType, data, msg := h.ServeHTTP(w, r, log)
		elapsed := time.Since(start)

		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.WriteHeader(code)
		if data != nil {
			w.Write(data)
		}

		log = log.WithFields(logrus.Fields{"elapsedTime": elapsed.Seconds(), "code": code})
		if code >= 500 {
			log.Error(msg)
		} else {
			log.Info(msg)
		}
	})
}
