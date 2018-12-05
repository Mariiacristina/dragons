package controller

import (
 "Reptile/API/model"
 "encoding/json"
 "log"
 "net/http"
 "github.com/gorilla/mux"
)

func GetAlarmas(w http.ResponseWriter, r *http.Request) {
	log.Println("controller - GET ALARMAS")
 vars := mux.Vars(r)
	id := vars["idPersona"]
	alarmas, err := model.GetAlarmas(id)
 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		json, errjson := json.Marshal(alarmas)
		if errjson != nil {
			http.Error(w, "error json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}
