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

func UpdateAutoSol(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - UPDATE AUTO SOL")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  var update_sol schema.Auto
  _=json.NewDecoder(r.Body).Decode(&update_sol)
  sol_ready,err := model.UpdateAutoSol(id,update_sol)
  if err != nil {
    log.Println("Error al actualizar")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(sol_ready)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}

func UpdateAutoTerrario(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - UPDATE AUTO TERRARIO")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  var update_terrario schema.Auto
  _=json.NewDecoder(r.Body).Decode(&update_terrario)
  terrario_ready,err := model.UpdateAutoTerrario(id,update_terrario)
  if err != nil {
    log.Println("Error al actualizar")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(terrario_ready)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}

func UpdateAutoHumedad(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - UPDATE AUTO HUMEDAD")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  var update_humedad schema.Auto
  _=json.NewDecoder(r.Body).Decode(&update_humedad)
  humedad_ready,err := model.UpdateAutoHumedad(id,update_humedad)
  if err != nil {
    log.Println("Error al actualizar")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(humedad_ready)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}

func UpdateAutoLuz(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - UPDATE AUTO LUZ")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  var update_luz schema.Auto
  _=json.NewDecoder(r.Body).Decode(&update_luz)
  luz_ready,err := model.UpdateAutoLuz(id,update_luz)
  if err != nil {
    log.Println("Error al actualizar")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(luz_ready)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.WriteHeader(http.StatusOK)
      w.Write(json)
  }
}

func GetAutoSol(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET AUTOMATIZACION SOL")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  get_sol, err := model.GetAutoSol(id)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(get_sol)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}

func GetAutoTerrario(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET AUTOMATIZACION TERRARIO")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  get_terrario, err := model.GetAutoTerrario(id)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(get_terrario)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}

func GetAutoHumedad(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET AUTOMATIZACION HUMEDAD")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  get_humedad, err := model.GetAutoHumedad(id)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(get_humedad)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}

func GetAutoLuz(w http.ResponseWriter, r *http.Request){
  log.Println("PERFIL - GET AUTOMATIZACION LUZ")
  vars := mux.Vars(r)
  id :=vars["idPersona"]
  get_luz, err := model.GetAutoLuz(id)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  }else{
    w.Header().Set("Content-Type", "application/json")
    json, errjson := json.Marshal(get_luz)
    if errjson != nil {
      http.Error(w,"error json", http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(json)
  }
}
