package handlers

// import (
// 	"net/http"
// 	"time"

// 	"github.com/sirupsen/logrus"
// )

// type statusCodeRecorder struct {
// 	http.ResponseWriter
// 	http.Hijacker
// 	StatusCode int
// }

// // LogEntryHandler is a request logging middleware
// func (handler HTTPHandler) LogEntryHandler(next http.Handler) http.Handler {
// 	f := func(w http.ResponseWriter, r *http.Request) {
// 		beginTime := time.Now()

// 		hijacker, _ := w.(http.Hijacker)
// 		w = &statusCodeRecorder{
// 			ResponseWriter: w,
// 			Hijacker:       hijacker,
// 		}

// 		defer func() {
// 			statusCode := w.(*statusCodeRecorder).StatusCode
// 			if statusCode == 0 {
// 				statusCode = 200
// 			}
// 			duration := time.Since(beginTime)

// 			logger := handler.logger.WithFields(logrus.Fields{
// 				"duration":    duration,
// 				"status_code": statusCode,
// 			})
// 		//	logger.Info(r.Method + " " + r.URL.RequestURI())
// 		}()

// 		next.ServeHTTP(w, r)
// 	}
// 	return http.HandlerFunc(f)
// }
