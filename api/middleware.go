package api

import (
	"net/http"
	"runtime/debug"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("panic error %v", err)
				log.Errorf(string(debug.Stack()))
			}
		}()
		start := time.Now()
		log.Debugf("%s %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		log.Debugf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	})
}
