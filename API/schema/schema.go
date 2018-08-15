package schema

type Signin struct{
  Nombre string `json:"nombre, omitempty"`
  Password string `json: "password, omitempty"`
}

type Reptil struct {
  Nombre string `json:"nombre, omitempty"`
  Tipo string `json:"tipo, omitempty"`
  Nacimiento string `json:"nacimiento, omitempty"`
  Temp_sol string `json:"temp_sol, omitempty"`
  Temp_sombra string `json:"temp_sombra, omitempty"`
  Temp_min string `json:"temp_min, omitempty"`
  Humedad string `json:"humedad, omitempty"`
  Humedad_min string `json:"humedad_min, omitempty"`
  }
