package model

import(
  "log"
  "Reptile/API/schema"
  "Reptile/API/connection"
  "strconv"
  "database/sql"
)

//ASI SE HACEN LOS SELECT
func Reptil(id string)(config_reptil schema.Reptil,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  var Config_reptil schema.Reptil
  err = db.QueryRow("SELECT sol_max,sol_min,temp_max,temp_min,humedad_min,config_chosen,uv_inicio,uv_tiempo,catarata_on,catarata_off FROM config_cliente WHERE config_cliente.id_cliente = ?",new_id).Scan(&Config_reptil.Sol_max,&Config_reptil.Sol_min,&Config_reptil.Temp_max,&Config_reptil.Temp_min,&Config_reptil.Humedad_min,&Config_reptil.Config_chosen,&Config_reptil.Uv_inicio,&Config_reptil.Uv_tiempo,&Config_reptil.Catarata_on,&Config_reptil.Catarata_off)
  connection.Disconnect(db)
  if(err == sql.ErrNoRows){
    return Config_reptil,nil
  }
  if(err != nil && err != sql.ErrNoRows) {
    log.Println("no tiene configuración:", err)
    return Config_reptil,err
  }else{
    return Config_reptil,err
  }
}


func GetCliente(id string)(persona schema.Cliente,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  var Persona schema.Cliente
  err = db.QueryRow("SELECT id_cliente,nombre,tipo,nacimiento FROM cliente WHERE id_cliente = ?",new_id).Scan(&Persona.Id_cliente,&Persona.Nombre,&Persona.Tipo,&Persona.Nacimiento)
  connection.Disconnect(db)
  if(err == sql.ErrNoRows){
    return Persona,nil
  }
  if(err != nil && err != sql.ErrNoRows) {
    log.Println("Problema al encontrar al cliente:", err)
    return Persona,err
  }else{
    return Persona,err
  }
}

func Default(defaults string)(reptil schema.Reptil,err error){
  db := connection.Connect()
  var Default schema.Reptil
  err = db.QueryRow("SELECT tipo,sol_max,sol_min,temp_max,temp_min,humedad_min,uv_inicio,uv_tiempo,catarata_on,catarata_off FROM defaults WHERE defaults.tipo = ?",defaults).Scan(&Default.Tipo,&Default.Sol_max,&Default.Sol_min,&Default.Temp_max,&Default.Temp_min,&Default.Humedad_min,&Default.Uv_inicio,&Default.Uv_tiempo,&Default.Catarata_on,&Default.Catarata_off)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("Problema al encontrar default:", err)
    return Default,err
  }else{
    return Default,err
  }
}

func UpdateCliente(id string,update_persona schema.Cliente)(persona schema.Cliente,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  _,err = db.Exec("UPDATE cliente SET nombre = ?, tipo = ?, nacimiento = ? WHERE cliente.id_cliente = ?",update_persona.Nombre,update_persona.Tipo,update_persona.Nacimiento,new_id)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("Problema en actualizar cliente:", err)
    return update_persona,err
  }else{
    update_persona.Id_cliente = new_id
    return update_persona,err
  }
}


func UpdateReptil(id string,update_reptil schema.Reptil)(reptil schema.Reptil,err error){
  new_id,err := strconv.Atoi(id)
  var id_prueba int
  db := connection.Connect()
  err = db.QueryRow("SELECT id_cliente FROM config_cliente WHERE config_cliente.id_cliente = ?",new_id).Scan(&id_prueba)
  if(err != nil && err != sql.ErrNoRows){
    log.Println("Problema en el modelo UPDATE REPTIL:", err)
    connection.Disconnect(db)
    return update_reptil,err
  }
  if(err != nil && err == sql.ErrNoRows){
    _,errExec = db.Exec("INSERT INTO config_cliente VALUES(?,?,?,?,?,?,?,?,?,?,?)",new_id,update_reptil.Sol_max,update_reptil.Sol_min,update_reptil.Temp_max,update_reptil.Temp_min,update_reptil.Humedad_min,update_reptil.Config_chosen,update_reptil.Uv_inicio,update_reptil.Uv_tiempo,update_reptil.Catarata_on,update_reptil.Catarata_off)
    connection.Disconnect(db)
    if(errExec != nil) {
      log.Println("Problema en INSERTAR Reptil:", err)
      connection.Disconnect(db)
      return update_reptil,errExec
    }else{
      connection.Disconnect(db)
      return update_reptil,errExec
    }
  }
  if(err == nil){
    _,errExec = db.Exec("UPDATE config_cliente SET sol_max = ?, sol_min = ?, temp_max = ?, temp_min = ?, humedad_min = ?, config_chosen = ?, uv_inicio = ?, uv_tiempo = ?, catarata_on = ?, catarata_off = ? WHERE config_cliente.id_cliente = ?",update_reptil.Sol_max,update_reptil.Sol_min,update_reptil.Temp_max,update_reptil.Temp_min,update_reptil.Humedad_min,update_reptil.Config_chosen,update_reptil.Uv_inicio,update_reptil.Uv_tiempo,update_reptil.Catarata_on,update_reptil.Catarata_off,new_id)
    connection.Disconnect(db)
    if(errExec != nil){
      log.Println("Problema en ACTUALIZAR Reptil:", err)
      connection.Disconnect(db)
      return update_reptil,errExec
    }else{
      connection.Disconnect(db)
      return update_reptil,errExec
    }
  }
  log.Println("nunca deberia llegar aca")
  connection.Disconnect(db)
  return update_reptil,err
}

func UpdateAutoSol(id string,update_sol schema.Auto)(auto_sol schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  db := connection.Connect()
  _,errExec = db.Exec("UPDATE automatizacion SET temp_sol = ? WHERE id_cliente = ?",update_sol.Estado,new_id)
  connection.Disconnect(db)
  if(errExec != nil){
    log.Println("Problema en ACTUALIZAR automatizacon sol:", err)
    return update_sol,errExec
  }else{
    return update_sol,errExec
  }
}

func UpdateAutoTerrario(id string,update_terrario schema.Auto)(auto_terrario schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  db := connection.Connect()
  _,errExec = db.Exec("UPDATE automatizacion SET temp_terrario = ? WHERE id_cliente = ?",update_terrario.Estado,new_id)
  connection.Disconnect(db)
  if(errExec != nil){
    log.Println("Problema en ACTUALIZAR automatizacon terrario:", err)
    return update_terrario,errExec
  }else{
    return update_terrario,errExec
  }
}

func UpdateAutoHumedad(id string,update_humedad schema.Auto)(auto_humedad schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  db := connection.Connect()
  _,errExec = db.Exec("UPDATE automatizacion SET humedad = ? WHERE id_cliente = ?",update_humedad.Estado,new_id)
  connection.Disconnect(db)
  if(errExec != nil){
    log.Println("Problema en ACTUALIZAR automatizacon humedad:", err)
    return update_humedad,errExec
  }else{
    return update_humedad,errExec
  }
}

func UpdateAutoLuz(id string,update_luz schema.Auto)(auto_luz schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  db := connection.Connect()
  _,errExec = db.Exec("UPDATE automatizacion SET luz = ? WHERE id_cliente = ?",update_luz.Estado,new_id)
  connection.Disconnect(db)
  if(errExec != nil){
    log.Println("Problema en ACTUALIZAR automatizacon luz:", err)
    return update_luz,errExec
  }else{
    return update_luz,errExec
  }
}

func GetAutoSol(id string)(auto_sol_resp schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  var auto_sol schema.Auto
  db := connection.Connect()
  err = db.QueryRow("SELECT temp_sol FROM automatizacion WHERE id_cliente = ?",new_id).Scan(&auto_sol.Estado)
  if(err != nil && err != sql.ErrNoRows){
    log.Println("Problema en el modelo GET AUTO SOL:", err)
    connection.Disconnect(db)
    return auto_sol,err
  }
  if(err != nil && err == sql.ErrNoRows){
    _,errExec = db.Exec("INSERT INTO automatizacion VALUES(?,?,?,?,?)",new_id,0,0,0,0)
    connection.Disconnect(db)
    if(errExec != nil) {
      log.Println("Problema en INSERTAR automatizacion:", err)
      connection.Disconnect(db)
      return auto_sol,errExec
    }else{
      auto_sol.Nombre = "temp_sol"
      connection.Disconnect(db)
      return auto_sol,errExec
    }
  }
  auto_sol.Nombre = "temp_sol"
  connection.Disconnect(db)
  return auto_sol,errExec
}

func GetAutoTerrario(id string)(auto_terrario_resp schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  var auto_terrario schema.Auto
  db := connection.Connect()
  err = db.QueryRow("SELECT temp_terrario FROM automatizacion WHERE id_cliente = ?",new_id).Scan(&auto_terrario.Estado)
  if(err != nil && err != sql.ErrNoRows){
    log.Println("Problema en el modelo GET AUTO TERRARIO:", err)
    connection.Disconnect(db)
    return auto_terrario,err
  }
  if(err != nil && err == sql.ErrNoRows){
    _,errExec = db.Exec("INSERT INTO automatizacion VALUES(?,?,?,?,0)",new_id,0,0,0,0)
    connection.Disconnect(db)
    if(errExec != nil) {
      log.Println("Problema en INSERTAR automatizacion TERRARIO:", err)
      connection.Disconnect(db)
      return auto_terrario,errExec
    }else{
      auto_terrario.Nombre = "temp_terrario"
      connection.Disconnect(db)
      return auto_terrario,errExec
    }
  }
  auto_terrario.Nombre = "temp_sol"
  connection.Disconnect(db)
  return auto_terrario,errExec
}

func GetAutoHumedad(id string)(auto_humedad_resp schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  var auto_humedad schema.Auto
  db := connection.Connect()
  err = db.QueryRow("SELECT humedad FROM automatizacion WHERE id_cliente = ?",new_id).Scan(&auto_humedad.Estado)
  if(err != nil && err != sql.ErrNoRows){
    log.Println("Problema en el modelo GET AUTO HUMEDAD:", err)
    connection.Disconnect(db)
    return auto_humedad,err
  }
  if(err != nil && err == sql.ErrNoRows){
    _,errExec = db.Exec("INSERT INTO automatizacion VALUES(?,?,?,?,?)",new_id,0,0,0,0)
    connection.Disconnect(db)
    if(errExec != nil) {
      log.Println("Problema en INSERTAR automatizacion HUMEDAD:", err)
      connection.Disconnect(db)
      return auto_humedad,errExec
    }else{
      auto_humedad.Nombre = "temp_humedad"
      connection.Disconnect(db)
      return auto_humedad,errExec
    }
  }
  auto_humedad.Nombre = "temp_humedad"
  connection.Disconnect(db)
  return auto_humedad,errExec
}

func GetAutoLuz(id string)(auto_humedad_resp schema.Auto,err error){
  new_id, err := strconv.Atoi(id)
  var auto_luz schema.Auto
  db := connection.Connect()
  err = db.QueryRow("SELECT luz FROM automatizacion WHERE id_cliente = ?",new_id).Scan(&auto_luz.Estado)
  if(err != nil && err != sql.ErrNoRows){
    log.Println("Problema en el modelo GET AUTO LUZ:", err)
    connection.Disconnect(db)
    return auto_luz,err
  }
  if(err != nil && err == sql.ErrNoRows){
    _,errExec = db.Exec("INSERT INTO automatizacion VALUES(?,?,?,?,?)",new_id,0,0,0,0)
    connection.Disconnect(db)
    if(errExec != nil) {
      log.Println("Problema en INSERTAR automatizacion LUZ:", err)
      connection.Disconnect(db)
      return auto_luz,errExec
    }else{
      auto_luz.Nombre = "temp_luz"
      connection.Disconnect(db)
      return auto_luz,errExec
    }
  }
  auto_luz.Nombre = "temp_luz"
  connection.Disconnect(db)
  return auto_luz,errExec
}
