package rutas

import (
  "github.com/gorilla/mux"
  "Reptile/API/controller"
  "net/http"
)

func Enrutamiento()(http.Handler){
  router := mux.NewRouter()

  //MainMenu
  router.HandleFunc("/", controller.Signin).Methods("POST")
  router.HandleFunc("/Register", controller.Register).Methods("POST")
  //Panel principal
  router.HandleFunc("/pp/temp", controller.GetTemperaturaSol).Methods("GET")
  router.HandleFunc("/pp/tempsol", controller.GetTemperaturaTerrario).Methods("GET")
  router.HandleFunc("/pp/humedad", controller.GetHumedadTerrario).Methods("GET")
  router.HandleFunc("/pp/luminocidad", controller.GetLuminocidad).Methods("GET")
  //Perfil
  router.HandleFunc("/perfil/{idPersona}", controller.GetInfoPersona).Methods("GET")
  router.HandleFunc("/perfil/{idPersona}/{idReptil}", controller.GetInfoReptil).Methods("GET")
  router.HandleFunc("/Perfil/{idPersona}", controller.PostReptil).Methods("POST")
  router.HandleFunc("/perfil/{idPersona}", controller.UpdatePersona).Methods("PUT")
  router.HandleFunc("/perfil/{idPersona}/{idReptil}", controller.DeleteReptil).Methods("DELETE")
  router.HandleFunc("/perfil/{idPersona}", controller.DeletePersona).Methods("DELETE")

  return router


}
