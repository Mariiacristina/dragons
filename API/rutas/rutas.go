package rutas

import (
	"Reptile/API/controller"
	"net/http"
	"github.com/gorilla/mux"
)

func Enrutamiento() http.Handler {
	router := mux.NewRouter()

	//MainMenu
	router.HandleFunc("/", controller.Signin).Methods("POST")
	router.HandleFunc("/Register", controller.Register).Methods("POST")
	//Panel principal
	router.HandleFunc("/sensores/{idPersona}", controller.GetSensores).Methods("GET")
	router.HandleFunc("/accesorios/{idPersona}", controller.GetAccesorios).Methods("GET")
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
	//AUTO
	router.HandleFunc("/auto/sol/{idPersona}", controller.UpdateAutoSol).Methods("PUT")
	router.HandleFunc("/auto/terrario/{idPersona}", controller.UpdateAutoTerrario).Methods("PUT")
	router.HandleFunc("/auto/humedad/{idPersona}", controller.UpdateAutoHumedad).Methods("PUT")
	router.HandleFunc("/auto/luminocidad/{idPersona}", controller.UpdateAutoLuz).Methods("PUT")
	router.HandleFunc("/auto/sol/{idPersona}", controller.GetAutoSol).Methods("GET")
	router.HandleFunc("/auto/terrario/{idPersona}", controller.GetAutoTerrario).Methods("GET")
	router.HandleFunc("/auto/humedad/{idPersona}", controller.GetAutoHumedad).Methods("GET")
	router.HandleFunc("/auto/luminocidad/{idPersona}", controller.GetAutoLuz).Methods("GET")
	//Botones
	router.HandleFunc("/botones/update/{idPersona}", controller.UpdateEstadosAccesorios).Methods("PUT")
	//FUNCIONES PARA LA ESP
	router.HandleFunc("/esp/{idPersona}", controller.UpdateEsp).Methods("POST")
	router.HandleFunc("/esp/{idPersona}", controller.GetConfigAPI).Methods("GET")
	router.HandleFunc("/alarma/{idPersona}", controller.PostAlarma).Methods("POST")
	//ALARMAS
	router.HandleFunc("/alarmas/{idPersona}", controller.GetAlarmas).Methods("GET")




	return router



}
