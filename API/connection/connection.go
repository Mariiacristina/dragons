package connection
import(
  "database/sql"
  "log"
  _"github.com/go-sql-driver/mysql"
)
var (
  db *sql.DB
  errdb error
)
func Connect()(db *sql.DB){
  db, errdb := sql.Open("mysql","sql10256343:ve2NaIqaIW@tcp(sql10.freemysqlhosting.net:3306)/sql10256343")
  if errdb != nil{
    log.Println("Error al connectar a la bdd")
    return
  } else {
    log.Println("Connectada!...")
    return db}
}

func Disconnect(db *sql.DB){
  errdb := db.Close()
  if errdb != nil {
    log.Println("Error al cerrar")
  } else {
    log.Println("Cerrando Sesion...")
  }
}
