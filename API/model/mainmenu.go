package model

import(
  "log"
  "Reptile/API/schema"
)
//aqui van los select * from usuario nombre y password o lo que sea, pero SQL
func UsuarioExistente(persona schema.Signin)(exist string,errmod error){
  var Usuario schema.Signin
  Usuario.Nombre = "Hola"
  Usuario.Password = "hola"
  if (persona.Nombre == Usuario.Nombre && persona.Password == Usuario.Password){
    errmod = nil
    return "Existe", errmod
    }else{
    errmod = nil
    return "Noexiste",errmod
  }

}


func Register(persona schema.Signin)(resp error){
  log.Println("entro al modelo")
  resp = nil
  return resp

}
