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
  err = db.QueryRow("SELECT sol_max,sol_min,temp_max,temp_min,humedad_min,config_chosen FROM config_cliente WHERE config_cliente.id_cliente = ?",new_id).Scan(&Config_reptil.Sol_max,&Config_reptil.Sol_min,&Config_reptil.Temp_max,&Config_reptil.Temp_min,&Config_reptil.Humedad_min,&Config_reptil.Config_chosen)
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
  err = db.QueryRow("SELECT tipo,sol_max,sol_min,temp_max,temp_min,humedad_min FROM defaults WHERE defaults.tipo = ?",defaults).Scan(&Default.Tipo,&Default.Sol_max,&Default.Sol_min,&Default.Temp_max,&Default.Temp_min,&Default.Humedad_min)
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
    _,errExec = db.Exec("INSERT INTO config_cliente VALUES(?,?,?,?,?,?,?)",new_id,update_reptil.Sol_max,update_reptil.Sol_min,update_reptil.Temp_max,update_reptil.Temp_min,update_reptil.Humedad_min,update_reptil.Config_chosen)
    connection.Disconnect(db)
    if(errExec != nil) {
      log.Println("Problema en INSERTAR Reptil:", err)
      return update_reptil,errExec
    }else{
      return update_reptil,errExec
    }
  }
  if(err == nil){
    _,errExec = db.Exec("UPDATE config_cliente SET sol_max = ?, sol_min = ?, temp_max = ?, temp_min = ?, humedad_min = ?, config_chosen = ? WHERE config_cliente.id_cliente = ?",update_reptil.Sol_max,update_reptil.Sol_min,update_reptil.Temp_max,update_reptil.Temp_min,update_reptil.Humedad_min,update_reptil.Config_chosen,new_id)
    connection.Disconnect(db)
    if(errExec != nil){
      log.Println("Problema en ACTUALIZAR Reptil:", err)
      return update_reptil,errExec
    }else{
      return update_reptil,errExec
    }
  }
  log.Println("nunca deberia llegar aca")
  return update_reptil,err
}
