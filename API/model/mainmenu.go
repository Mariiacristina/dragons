package model

import(
  "log"
  "Reptile/API/schema"
  "Reptile/API/connection"
)

var (
  err error
  errExec error
)

func UsuarioExistente(persona schema.Signin)(usuario schema.Signin,err error){
  var Usuario schema.Signin
  db := connection.Connect()
  err = db.QueryRow("SELECT id_cliente,nombre,password from cliente WHERE cliente.nombre = ?",persona.Nombre).Scan(&Usuario.Id_cliente,&Usuario.Nombre,&Usuario.Password)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("error en el modelo:", err)
    return Usuario,err
  }else{
    if (persona.Nombre == Usuario.Nombre && persona.Password == Usuario.Password){
      return Usuario,err
    }
    return Usuario,err
  }
}

func Register(persona schema.Signin)(resp error){
  db := connection.Connect()
  _,err = db.Exec("INSERT INTO cliente (nombre,password) VALUES (?,?)",persona.Nombre,persona.Password)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("error al insertar:", persona.Nombre)
    log.Println(err)
  }else{ log.Println("Insertando!!")}
  return err

}
