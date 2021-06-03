package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var staticDir = "/app/static/"

func Play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	file := vars["filename"]
	log.Printf("Playing - %v \n", file)
	seqName, ok := vars["sequence"]
	if !ok {
		http.ServeFile(w, r, staticDir+file+"/outputlist.m3u8")
	} else {
		http.ServeFile(w, r, staticDir+file+"/"+seqName)
	}

}
