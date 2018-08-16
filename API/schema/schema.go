package schema

type Signin struct{
  Nombre string `json:"nombre, omitempty"`
  Password string `json: "password, omitempty"`
}

type Reptil struct {
  Nombre string `json:"nombre, omitempty"`
  Tipo string `json:"tipo, omitempty"`
  Nacimiento string `json:"nacimiento, omitempty"`
  Temp_sol_max string `json:"temp_sol_max, omitempty"`
  Temp_sol_min string  `json:"temp_sol_min, omitempty"`
  Temp_min string `json:"temp_min, omitempty"`
  Humedad_min string `json:"humedad_min, omitempty"`
  }
