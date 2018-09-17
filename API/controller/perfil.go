package controller

import(
  "net/http"
  "log"
  "github.com/gorilla/mux"
  "encoding/json"
  "Reptile/API/model"
  "Reptile/API/schema"
)


func GetInfoPersona(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET INFORMACION PERSONA")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  cliente, err := model.GetCliente(id)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(cliente)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}

func UpdatePersona(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - UPDATE PERSONA")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  var update_persona schema.Cliente
  _=json.NewDecoder(r.Body).Decode(&update_persona)
  persona_ready,err := model.UpdateCliente(id,update_persona)
  if err != nil {
    log.Println("Error al actualizar")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(persona_ready)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}

func DeletePersona(w http.ResponseWriter, r *http.Request){
  log.Println("DELETE PERSONA")
}

func GetInfoReptil(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET INFORMACION REPTIL")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  reptil, err := model.Reptil(id)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(reptil)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}

func UpdateReptil(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - UPDATE REPTIL")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  var update_reptil schema.Reptil
  _=json.NewDecoder(r.Body).Decode(&update_reptil)
  reptil_ready,err := model.UpdateReptil(id,update_reptil)
  if err != nil {
    log.Println("Error al actualizar")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(reptil_ready)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}

func DeleteReptil(w http.ResponseWriter, r *http.Request){
  log.Println("DELETE REPTIL")
}

func GetDefault(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET DEFAULT REPTIL")
  vars := mux.Vars(r)
  seleccion :=vars["defaults"]
  default_reptil, err := model.Default(seleccion)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(default_reptil)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}
