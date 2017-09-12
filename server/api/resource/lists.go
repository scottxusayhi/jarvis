package resource

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
)

func item(r *http.Request) (string) {
	return mux.Vars(r)["item"]
}


func ListHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	defer r.Body.Close()

	// CRUD
	switch r.Method {
	case http.MethodGet:
		listItem(w, r)
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func listItem(w http.ResponseWriter, r *http.Request) {
	item := item(r)
	switch item {
	case "tags":
		listAllTags(w, r)
		break
	default:
		log.Error("unknow item to list: " + item)
		w.WriteHeader(http.StatusNotFound)
	}
}
