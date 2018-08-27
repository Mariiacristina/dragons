package controller

import(
  "net/http"
  "log"
  //"encoding/json"
)

func GetTemperaturaSol(w http.ResponseWriter, r *http.Request){
  log.Println("GET TEMPERATURA PIEDRA")
}

func GetTemperaturaTerrario(w http.ResponseWriter, r *http.Request){
  log.Println("GET TEMPERATURA TERRARIO")
}

func GetHumedadTerrario(w http.ResponseWriter, r *http.Request){
  log.Println("GET HUMEDAD TERRARIO")
}

func GetLuminocidad(w http.ResponseWriter, r *http.Request){
  log.Println("GET LUMINOCIDAD")
}
