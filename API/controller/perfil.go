package controller

import(
  "net/http"
  "log"
  //"encoding/json"
)

func GetInfoPersona(w http.ResponseWriter, r *http.Request){
  log.Println("GET INFORMACION PERSONA")
  /* esto es para get
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
}

func UpdatePersona(w http.ResponseWriter, r *http.Request){
  log.Println("UPDATE PERSONA")
}

func DeletePersona(w http.ResponseWriter, r *http.Request){
  log.Println("DELETE PERSONA")
}

func GetInfoReptil(w http.ResponseWriter, r *http.Request){
  log.Println("GET INFORMACION REPTIL")
}

func PostReptil(w http.ResponseWriter, r *http.Request){
  log.Println("POST REPTIL")
}

func UpdateReptil(w http.ResponseWriter, r *http.Request){
  log.Println("UPDATE REPTIL")
}

func DeleteReptil(w http.ResponseWriter, r *http.Request){
  log.Println("DELETE REPTIL")
}
