package schema

type Signin struct {
	Id_cliente int    `json:"id_cliente, omitempty"`
	Nombre     string `json:"nombre, omitempty"`
	Password   string `json: "password, omitempty"`
}

type Cliente struct {
	Id_cliente int    `json:"id_cliente, omitempty"`
	Nombre     string `json:"nombre, omitempty"`
	Password   string `json:"password, omitempty"`
	Tipo       string `json:"tipo, omitempty"`
	Nacimiento string `json:"nacimiento, omitempty"`
}

type Reptil struct {
	Tipo          string `json:"tipo, omitempty"`
	Sol_max       string `json:"sol_max, omitempty"`
	Sol_min       string `json:"sol_min, omitempty"`
	Temp_max      string `json:"temp_max, omitempty"`
	Temp_min      string `json:"temp_min, omitempty"`
	Humedad_min   string `json:"humedad_min, omitempty"`
	Config_chosen string `json:"config_chosen, omitempty"`
}

type Sensor struct {
	Nombre string `json:"nombre, omitempty"`
	Valor  int    `json:"valor, omitempty"`
}

type Accesorio struct {
	Nombre string `json:"nombre, omitempty"`
	Estado string `json:"estado, omitempty"`
}
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

type Auto struct {
	Nombre string `json:"nombre, omitempty"`
	Estado int    `json:"estado, omitempty"`
}
