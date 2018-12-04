package model
/*
import(
  "log"
  "Reptile/API/schema"
  "Reptile/API/connection"
  "strconv"
  "database/sql"
)

//ASI SE HACEN LOS SELECT
func Reptil(id string)(res schema.Alarma,err error){
  new_id,err := strconv.Atoi(id)
  db := connection.Connect()
  var Bdd_alarma schema.Alarma
  rows, err := db.Query("SELECT hora, razon FROM alarmas WHERE config_cliente.id_cliente = ?",new_id).Scan(&Bdd_alarma.Hora,Bdd_alarma.Razon)
  defer rows.Close()
  for rows.Next() {
      var hora string
      var razon string
      err = rows.Scan(&hora, &razon)
  }
  connection.Disconnect(db)
  if(err == sql.ErrNoRows){
    return Bdd_alarma,nil
  }
  if(err != nil && err != sql.ErrNoRows) {
    log.Println("no tiene configuración:", err)
    return Bdd_alarma,err
  }else{
    return Bdd_alarma,err
  }
}
*/
