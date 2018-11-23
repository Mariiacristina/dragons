package controller

import(
  "net/http"
  "log"
  "encoding/json"
  "Reptile/API/model"
  "Reptile/API/schema"
)

func Signin(w http.ResponseWriter, r *http.Request){
  log.Println("MAINMENU - SIGN IN")
  var persona schema.Signin
  _=json.NewDecoder(r.Body).Decode(&persona)
  cliente,err := model.UsuarioExistente(persona)
  if err != nil {
    log.Println("No se encontró usuario")
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
      log.Println("Se ha loggueado exitosamente")
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

func Register(w http.ResponseWriter, r *http.Request){
  log.Println("MAINMENU - REGISTER")
  var cliente schema.Signin
  _=json.NewDecoder(r.Body).Decode(&cliente)
  err_reg := model.Register(cliente)
  if (err_reg != nil) {
    w.WriteHeader(http.StatusInternalServerError)
  }else{
    w.WriteHeader(http.StatusCreated)
  }
}
