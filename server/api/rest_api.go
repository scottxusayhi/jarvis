package api


import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"net"
)

func rootHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rock Solid!\n"))
}

func healthHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All systems go!\n"))
}


func NewServer(l net.Listener) error {
	defer log.Error("HTTP server failed to start or stopped")
	m := mux.NewRouter()
	// serve everything in dir ./web under the path /
	m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web"))))
	//m.HandleFunc("/", rootHandler)
	m.HandleFunc("/api/v1/health", healthHandler)
	log.Info("HTTP REST server started")
	return http.Serve(l, m)
}
