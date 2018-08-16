package controller

import(
  "net/http"
  "log"
  "encoding/json"
  "Reptile/API/model"
  //"github.com/gorilla/mux"
  "Reptile/API/schema"
)

func Signin(w http.ResponseWriter, r *http.Request){
  log.Println("SIGN IN")
  var persona schema.Signin
  _=json.NewDecoder(r.Body).Decode(&persona)
  res, err := model.UsuarioExistente(persona)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
    if (res == "Existe"){
      log.Println("Se ha loggueado exitosamente")
      //setea la cabezera como tipo json
      w.Header().Set("Content-Type", "application/json")
      json, errjson := json.Marshal(persona)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      //escribimos como cabezera que está todo bien y esscribimos el cuerpo
      w.WriteHeader(http.StatusOK)
      w.Write(json)
    }else{
      log.Println("No se encontró usuario")
      http.Error(w,"No se encontró usuario", http.StatusInternalServerError)
    }
  }
}

func Register(w http.ResponseWriter, r *http.Request){
  var cliente schema.Signin
  _=json.NewDecoder(r.Body).Decode(&cliente)
  err_reg := model.Register(cliente)
  if (err_reg != nil) {
    w.WriteHeader(http.StatusInternalServerError)
  }else{
    w.WriteHeader(http.StatusCreated)
  }
}
