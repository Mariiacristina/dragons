package model

import(
  "log"
  "Reptile/API/schema"
  "Reptile/API/connection"
  "strconv"
  "database/sql"
  "fmt"
)


func GetSensores_esp(id string)(resp_sensores schema.Sensores,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  var Select_sensores schema.Sensores
  err = db.QueryRow("SELECT sensor_sol, sensor_terrario, sensor_humedad FROM esp WHERE esp.id_cliente = ?",new_id).Scan(&Select_sensores.Sol,&Select_sensores.Terrario,&Select_sensores.Humedad)
  connection.Disconnect(db)
  if(err == sql.ErrNoRows){
    return Select_sensores,nil
  }
  if(err != nil && err != sql.ErrNoRows) {
    log.Println("problema en encontrar sensores", err)
    return Select_sensores,err
  }else{
    return Select_sensores,err
  }
}


func GetAccesorios_esp(id string)(resp_accesorios schema.Objetos,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  var Select_objetos schema.Objetos
  err = db.QueryRow("SELECT estado_placa, estado_bombillo, estado_cascada, estado_uv FROM esp WHERE esp.id_cliente = ?",new_id).Scan(&Select_objetos.PlacaTermica,&Select_objetos.FocoTermico,&Select_objetos.Catarata,&Select_objetos.Uv)
  connection.Disconnect(db)
  if(err == sql.ErrNoRows){
    return Select_objetos,nil
  }
  if(err != nil && err != sql.ErrNoRows) {
    log.Println("problema en encontrar accesorios", err)
    return Select_objetos,err
  }else{
    return Select_objetos,err
  }
}
//arreglar este update
func UpdateAccesorio(id string,update_accesorio schema.Sensor)(accesorio_update schema.Sensor,err error){
  new_id,err := strconv.Atoi(id)
  log.Println(update_accesorio)
  db := connection.Connect()
  s := fmt.Sprintf("UPDATE estado_objetos SET %s = %d WHERE estado_objetos.id_cliente = %d",update_accesorio.Nombre,update_accesorio.Valor,new_id)
  _,err = db.Exec(s)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("Problema en actualizar accesorios", err)
    return update_accesorio,err
  }else{
    return update_accesorio,err
  }
}
