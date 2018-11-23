package controller

import (
	"Reptile/API/schema"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	 "Reptile/API/model"
)


func GetSensores(w http.ResponseWriter, r *http.Request) {
	log.Println("SENSORES - GET SENSORES")
	vars := mux.Vars(r)
	id :=vars["idPersona"]
	sensores,err := model.GetSensores_esp(id)
	if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(sensores)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
}
}

func GetAccesorios(w http.ResponseWriter, r *http.Request) {
	log.Println("SENSORES - GET ACCESORIOS")
	vars := mux.Vars(r)
	id :=vars["idPersona"]
	accesorios,err := model.GetAccesorios_esp(id)
	if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(accesorios)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
}
}

func UpdateEstadosAccesorios(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id :=vars["idPersona"]
	var accesorio schema.Sensor
	_=json.NewDecoder(r.Body).Decode(&accesorio)
		log.Println("SENSORES - FOCO: ", accesorio.Nombre)
	accesorio_ready,err := model.UpdateAccesorio(id,accesorio)
	if err != nil {
		log.Println("Error al actualizar")
		http.Error(w, "error en la base de datos", http.StatusInternalServerError)
		return
	} else {
			w.Header().Set("Content-Type", "application/json")
			json, errjson := json.Marshal(accesorio_ready)
			if errjson != nil {
				http.Error(w,"error json", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(json)
	}
	}
