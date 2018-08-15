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
  log.Println("ingresar a: ",persona)
  res, err := model.UsuarioExistente(persona)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
    if (res == "Existe"){
      log.Println("Se ha loggueado exitosamente")
      w.WriteHeader(http.StatusOK)
      resjson, errjson := json.Marshal(persona)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.Write(resjson)
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
    w.WriteHeader(http.StatusOK)
  }
}



/*
func Signin(w http.ResponseWriter, r *http.Request){
  log.Println("SIGN IN")
  vars := mux.Vars(r)
  nombre :=vars["name"]
  password := vars["password"]
  existente, err := model.UsuarioExistente(nombre,password)
  if err != nil {
    http.Error(w, "error en la base de datos", http.StatusInternalServerError)
    return
  } else {
    if (nombre == existente.Nombre && password == existente.Password){
      log.Println("Se ha loggueado exitosamente")
      w.WriteHeader(http.StatusOK)
      res, errjson := json.Marshal(existente)
      if errjson != nil {
        http.Error(w,"error json", http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.Write(res)
    }else{
      log.Println("No se encontró usuario")
      http.Error(w,"No se encontró usuario", http.StatusInternalServerError)
    }
  }
}
*/
