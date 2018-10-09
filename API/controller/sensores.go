package controller

import (
	"Reptile/API/schema"
	"encoding/json"
	"log"
	"net/http"
)

//borrar es la estructura
type Botones struct {
	PLaca          string `json:"placa, omitempty"`
	EstadoPlaca    string `json:"estadoplaca, omitempty"`
	Bombillo       string `json:"bombillo, omitempty"`
	EstadoBombillo string `json:"estadobombillo, omitempty"`
	Cascada        string `json:"cascada, omitempty"`
	EstadoCascada  string `json:"estadocascada, omitempty"`
	Uv             string `json:"uv, omitempty"`
	EstadoUv       string `json:"estadouv, omitempty"`
}

func GetBotones(w http.ResponseWriter, r *http.Request) {
	log.Println("Actuadores - GET sí: (ON-OFF)")
	var botones schema.Botones
	botones.Placa = "Foco_Termico"
	botones.Estado = "prendido"
	w.Header().Set("Content-Type", "application/json")
	json, errjson := json.Marshal(foco_termico)
	if errjson != nil {
		http.Error(w, "error json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
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
