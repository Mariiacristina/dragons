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
  router.HandleFunc("/pp/tempterrario", controller.GetTemperaturaTerrario).Methods("GET")
  router.HandleFunc("/pp/humedad", controller.GetHumedadTerrario).Methods("GET")
  //Accesorios
  router.HandleFunc("/pp/luminocidad", controller.GetLuminocidad).Methods("GET")
  router.HandleFunc("/pp/luminocidad", controller.UpdateLuminocidad).Methods("PUT")
  router.HandleFunc("/pp/focotermico", controller.GetFocoTermico).Methods("GET")
  router.HandleFunc("/pp/focotermico", controller.UpdateFocoTermico).Methods("PUT")
  router.HandleFunc("/pp/placatermica", controller.GetPlacaTermica).Methods("GET")
  router.HandleFunc("/pp/placatermica", controller.UpdatePlacaTermica).Methods("PUT")
  //Perfil
  router.HandleFunc("/perfil/{idPersona}", controller.GetInfoPersona).Methods("GET")
  router.HandleFunc("/perfil/{idPersona}", controller.UpdatePersona).Methods("PUT")
  router.HandleFunc("/perfil/{idPersona}", controller.DeletePersona).Methods("DELETE")
  //reptil
  router.HandleFunc("/reptil/{idPersona}", controller.GetInfoReptil).Methods("GET")
  router.HandleFunc("/reptil/{idPersona}", controller.UpdateReptil).Methods("PUT")
  router.HandleFunc("/reptil/{idPersona}", controller.DeleteReptil).Methods("DELETE")
  //default
  router.HandleFunc("/defaults/{defaults}", controller.GetDefault).Methods("GET")

  return router


}
