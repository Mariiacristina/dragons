package controller

import (
	"Reptile/API/schema"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	 "Reptile/API/model"
)


func UpdateEsp(w http.ResponseWriter, r *http.Request) {
	log.Println("POST ESP")
	var update_todo schema.Esp
	_=json.NewDecoder(r.Body).Decode(&update_todo)
	update_ready,err := model.UpdateESP(update_todo)
	if err != nil {
		log.Println("Error al actualizar")
		http.Error(w, "error en la base de datos", http.StatusInternalServerError)
		return
	} else {
			w.Header().Set("Content-Type", "application/json")
			json, errjson := json.Marshal(update_ready)
			if errjson != nil {
				http.Error(w,"error json", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(json)
	}
	}

func GetConfigAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("ESP - GET CONFIGURACION COMPLETA")
	vars := mux.Vars(r)
	id :=vars["idPersona"]
	config_completa,err := model.ConfigESP(id)
	if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(config_completa)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
}
}

func PostAlarma(w http.ResponseWriter, r *http.Request) {
	log.Println("ESP - POST ALARMA")
  var alarma schema.Alarma
  _=json.NewDecoder(r.Body).Decode(&alarma)
  res_alarma,err := model.PostAlarmaESP(alarma)
  if err != nil {
    log.Println("no se pudo postear")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(res_alarma)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}
