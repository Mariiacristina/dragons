package controller

import (
	"Reptile/API/schema"
	"encoding/json"
	"log"
	"net/http"
)

func GetTemperaturaSol(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - GET TEMPERATURA SOL")
	var temperatura_sol schema.Sensor
	temperatura_sol.Nombre = "temperatura_sol"
	temperatura_sol.Valor = 37
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(temperatura_sol)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetTemperaturaTerrario(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - GET TEMPERATURA TERRARIO")
	var temperatura_terrario schema.Sensor
	temperatura_terrario.Nombre = "temperatura_terrario"
	temperatura_terrario.Valor = 28
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(temperatura_terrario)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetHumedadTerrario(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - GET HUMEDAD TERRARIO")
	var temperatura_humedad schema.Sensor
	temperatura_humedad.Nombre = "temperatura_humedad"
	temperatura_humedad.Valor = 50
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(temperatura_humedad)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetLuminocidad(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - GET LUMINOCIDAD (ON-OFF)")
	var foco_UV schema.Accesorio
	foco_UV.Nombre = "Foco_UV"
	foco_UV.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(foco_UV)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetFocoTermico(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - GET FOCOTERMICO (ON-OFF)")
	var foco_termico schema.Accesorio
	foco_termico.Nombre = "Foco_Termico"
	foco_termico.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(foco_termico)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetPlacaTermica(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - GET PLACATERMICA (ON-OFF)")
	var placa_termica schema.Accesorio
	placa_termica.Nombre = "Placa_Termica"
	placa_termica.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(placa_termica)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateLuminocidad(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - UPDATE FOCO UV (ON-OFF)")
	var placa_termica schema.Accesorio
	placa_termica.Nombre = "Placa_Termica"
	placa_termica.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(placa_termica)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateFocoTermico(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - UPDATE FOCO TERMICO (ON-OFF)")
	var placa_termica schema.Accesorio
	placa_termica.Nombre = "Placa_Termica"
	placa_termica.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(placa_termica)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdatePlacaTermica(w http.ResponseWriter, r *http.Request) {
	//aqui va lo que se obtiene del arduino
	log.Println("SENSORES - UPDATE PLACATERMICA (ON-OFF)")
	var placa_termica schema.Accesorio
	placa_termica.Nombre = "Placa_Termica"
	placa_termica.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(placa_termica)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetBotones(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtención del estado-botones al cargar la pagina")
	var botones schema.Botones
	//Al escalar se puede agregar marca del complemeto para estadisticas
	//Valores debiesen obtenerse de una bbdd
	botones.Placa = "true"
	botones.EstadoPlaca = "estadoplaca"
	botones.Bombillo = "false"
	botones.EstadoBombillo = "estadobombillo"
	botones.Cascada = "true"
	botones.EstadoCascada = "estadocasscada"
	botones.Uv = "false"
	botones.EstadoUv = "estadouv"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(botones)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func SetPlaca(w http.ResponseWriter, r *http.Request) {
	log.Println("Gestion del estado de la placa")
	var botones schema.Botones
	botones.Placa = "false"
	log.Println("Cambiando el Estado de la Maquina.. obtubo estado anterior")
	//botones.Placa = "true"
	//botones.EstadoPlaca = "estadoplaca"
	if botones.Placa == "true" {
		botones.Placa = "false"
	}
	if botones.Placa == "false" {
		botones.Placa = "true"
	}
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(botones)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
