package model

import(
  "log"
  "Reptile/API/schema"
  "Reptile/API/connection"
  "strconv"
  "database/sql"
)

func UpdateESP(update_esp schema.Esp)(resp_update schema.Esp, err error){
  db := connection.Connect()
    _,err = db.Exec("UPDATE esp SET sensor_sol = ?, sensor_terrario = ?, sensor_humedad = ?, estado_placa = ?, estado_bombillo = ?, estado_cascada = ?, estado_uv = ? WHERE esp.Id_cliente = ?",update_esp.Esp_sol,update_esp.Esp_terrario,update_esp.Esp_humedad,update_esp.Esp_placatermica,update_esp.Esp_focotermico,update_esp.Esp_catarata,update_esp.Esp_uv,update_esp.Id_cliente)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("Problema en actualizar accesorios", err)
    return update_esp,err
  }else{
    return update_esp,err
  }
}


func ConfigESP(id string)(resp_configs schema.Todo,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  var Select_config_sensores schema.Todo
  err = db.QueryRow("SELECT config_cliente.sol_max, config_cliente.sol_min, config_cliente.temp_max, config_cliente.temp_min, config_cliente.uv_inicio, config_cliente.uv_tiempo, config_cliente.catarata_on, config_cliente.catarata_off, estado_objetos.uv, estado_objetos.foco_termico, estado_objetos.placa_termica, estado_objetos.catarata, automatizacion.temp_sol,automatizacion.temp_terrario,automatizacion.humedad,automatizacion.luz FROM config_cliente, estado_objetos, automatizacion  WHERE config_cliente.id_cliente = ?",new_id).Scan(&Select_config_sensores.Sol_max,&Select_config_sensores.Sol_min,&Select_config_sensores.Temp_max,&Select_config_sensores.Temp_min,&Select_config_sensores.Humedad_min,&Select_config_sensores.Uv_inicio,&Select_config_sensores.Uv_tiempo,&Select_config_sensores.Catarata_on,&Select_config_sensores.Catarata_off,&Select_config_sensores.Uv,&Select_config_sensores.FocoTermico,&Select_config_sensores.PlacaTermica,&Select_config_sensores.Catarata,&Select_config_sensores.Auto_sol,&Select_config_sensores.Auto_terrario,&Select_config_sensores.Auto_humedad,&Select_config_sensores.Auto_luz)
  connection.Disconnect(db)
  if(err == sql.ErrNoRows){
    return Select_config_sensores,nil
  }
  if(err != nil && err != sql.ErrNoRows) {
    log.Println("problema en encontrar accesorios", err)
    return Select_config_sensores,err
  }else{
    return Select_config_sensores,err
  }
}


//FALTAN
func PostAlarmaESP(post_alarma schema.Alarma)(resp_alarma schema.Alarma,err error){
  db := connection.Connect()
  _,err = db.Exec("INSERT INTO alarmas (id_cliente,hora,razon) VALUES (?,?,?)",post_alarma.Id_cliente,post_alarma.Hora,post_alarma.Razon)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("error al insertar:", post_alarma.Id_cliente)
    log.Println(err)
    return post_alarma,err
  }else{
  return post_alarma,err
  }
}
