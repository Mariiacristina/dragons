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
	Uv_inicio     string `json:"uv_inicio, omitempty"`
	Uv_tiempo     string `json:"uv_tiempo, omitempty"`
	Catarata_on   string `json:"catarata_on, omitempty"`
	Catarata_off  string `json:"catarata_off, omitempty"`
}

type Sensores struct {
	Sol					 int 	  `json:"sol, omitempty"`
	Terrario 		 int    `json:"terrario, omitempty"`
	Humedad  		 int    `json:"humedad, omitempty"`
}

type Objetos struct {
	Uv 							int		 `json:"uv, omitempty"`
	FocoTermico 	  int    `json:"focotermico, omitempty"`
	PlacaTermica  	int    `json:"placatermica, omitempty"`
	Catarata  			int    `json:"catarata, omitempty"`
}

type Sensor struct {
	Nombre				 string `json:"nombre, omitempty"`
	Valor  				 int    `json:"valor, omitempty"`
}

type Auto struct {
	Nombre 					string `json:"nombre, omitempty"`
	Estado				  int    `json:"estado, omitempty"`
}

type Esp struct {
	Id_cliente   		       int `json:"id_cliente, omitempty"`
	Esp_sol     				   int `json:"esp_sol, omitempty"`
	Esp_terrario    			 int `json:"esp_terrario, omitempty"`
	Esp_humedad     			 int `json:"esp_humedad, omitempty"`
	Esp_uv      					 int `json:"esp_uv, omitempty"`
	Esp_focotermico   		 int `json:"esp_focotermico, omitempty"`
	Esp_placatermica 			 int `json:"esp_placatermica, omitempty"`
	Esp_catarata      		 int `json:"esp_catarata, omitempty"`
}

type Alarma struct {
	Id_cliente   		       int `json:"id_cliente, omitempty"`
	Hora     				  		 string `json:"hora, omitempty"`
	Razon   				 			 string `json:"razon, omitempty"`
}

type Todo struct {
	Sol_max       string `json:"sol_max, omitempty"`
	Sol_min       string `json:"sol_min, omitempty"`
	Temp_max      string `json:"temp_max, omitempty"`
	Temp_min      string `json:"temp_min, omitempty"`
	Humedad_min   string `json:"humedad_min, omitempty"`
	Uv_inicio     string `json:"uv_inicio, omitempty"`
	Uv_tiempo     string `json:"uv_tiempo, omitempty"`
	Catarata_on   string `json:"catarata_on, omitempty"`
	Catarata_off  string `json:"catarata_off, omitempty"`
	Uv 							int		 `json:"uv, omitempty"`
	FocoTermico 	  int    `json:"focotermico, omitempty"`
	PlacaTermica  	int    `json:"placatermica, omitempty"`
	Catarata  			int    `json:"catarata, omitempty"`
	Auto_sol  			int    `json:"auto_sol, omitempty"`
	Auto_terrario  	int    `json:"auto_terrario, omitempty"`
	Auto_humedad  	int    `json:"auto_humedad, omitempty"`
	Auto_luz	 			int    `json:"auto_luz, omitempty"`
}
