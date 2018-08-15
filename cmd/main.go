package main

import(
  "net/http"
  "log"
  "Reptile/API/rutas"
  "github.com/gorilla/handlers"
)

func main (){
  headers := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
  methods := handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE"})
  origins := handlers.AllowedOrigins([]string{"*"})

  m := rutas.Enrutamiento()
  log.Fatal(http.ListenAndServe(":8000",handlers.CORS(headers,methods,origins)(m)))



}
