package main

import (
	"encoding/json"
	"fmt"
	"github.com/habajca/simple-log-search/util"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uid := r.FormValue("uid")
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "uid required.")
		return
	}
	domain := r.FormValue("domain")
	if domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "domain required.")
		return
	}
	lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid lat: %s.\n", r.FormValue("lat"))
		return
	}
	lng, err := strconv.ParseFloat(r.FormValue("lng"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid lng: %s.\n", r.FormValue("lng"))
		return
	}
	logRow := util.LogRow{
		Timestamp: time.Now().Unix(),
		Uid:       uid,
		Domain:    domain,
		Geo: util.GeoPoint{
			Latitude:  lat,
			Longitude: lng,
		},
	}
	bytes, err := json.Marshal(logRow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(bytes))
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	port := "8080"
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
