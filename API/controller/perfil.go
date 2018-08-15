package controller

import(
  "net/http"
  "log"
  //"encoding/json"
)

func GetInfoPersona(w http.ResponseWriter, r *http.Request){
  log.Println("GET INFORMACION PERSONA")
}

func GetInfoReptil(w http.ResponseWriter, r *http.Request){
  log.Println("GET INFORMACION REPTIL")
}

func PostReptil(w http.ResponseWriter, r *http.Request){
  log.Println("POST REPTIL")
}

func UpdatePersona(w http.ResponseWriter, r *http.Request){
  log.Println("UPDATE PERSONA")
}

func DeleteReptil(w http.ResponseWriter, r *http.Request){
  log.Println("DELETE REPTIL")
}

func DeletePersona(w http.ResponseWriter, r *http.Request){
  log.Println("DELETE PERSONA")
}
