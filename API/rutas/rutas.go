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
	router.HandleFunc("/pp/temp/{idPersona}", controller.GetTemperaturaSol).Methods("GET")
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
	router.HandleFunc("/botones/", controller.GetBotones).Methods("GET")
	router.HandleFunc("/botones/placa/", controller.SetPlaca).Methods("PUT")
	router.HandleFunc("/botones/bombillo/", controller.SetBombillo).Methods("PUT")
	router.HandleFunc("/botones/cascada/", controller.SetCascada).Methods("PUT")
	router.HandleFunc("/botones/uv/", controller.SetUv).Methods("PUT")
	return router

}
