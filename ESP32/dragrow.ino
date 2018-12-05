//LIBRERIAS
#include "WiFi.h"

//CLIENTE
int Id_cliente = 1;

//SENSORES DE ESP
int Sensor_sol = 40;
int Sensor_terrario = 30;
int Sensor_humedad = 60;

//OBJETOS DE ESP
int Placa_termica = 1;
int Foco_termico = 1;
int Uv = 1;
int Catarata = 1;


//VARIABLES DEL CLIENTE
String  API_sol_max       ;     
String  API_sol_min       ;      
String  API_temp_max      ;    
String  API_temp_min      ;   
String  API_humedad_min   ;
String  API_uv_inicio     ;
String  API_uv_tiempo     ;
String  API_catarata_on   ;
String  API_catarata_off  ;
int     API_uv            ;
int     API_focoTermico   ;
int     API_placaTermica  ;
int     API_catarata      ;
int     API_auto_sol      ;
int     API_auto_terrario ;
int     API_auto_humedad  ;
int     API_auto_luz      ;

//CONTADORES
int Count_sol = 0; //puede ser sol max o sol min pero no ambos
int Count_terrario = 0; // puede ser terrario max temp o min pero no ambos
int Count_ambos = 0; // ambos sensores estan fuera de margen

//ESTRUCTURAS
struct PostSensores {
   int sol;
   int terrario;
   int h;
};

//--------------------------------------------------------------------------------FUNCIONES-------------------------------------------
int Getconfig (int id){
  Serial.print("obteniendo configuracion para el id :");
  Serial.print(id);
 }

struct PostSensores LeerDatos(){
  struct PostSensores holi;
  holi.sol = Sensor_sol; // LEER TEMP SOL DE LOS SENSORES
  holi.terrario = Sensor_terrario; // LEER TEMP TERRARIO DE LOS SENSORES
  holi.h = Sensor_humedad; // LEER HUMEDAD DE LOS SENSORES DE LOS SENSORES
  return holi;
 }

 void PostDatosApi (struct PostSensores holi){
  Serial.print("------------EMPIEZA--------------- POST-----------");
  Serial.print("se postearon los siguientes datos");
  Serial.print("temperatura sol:");
  Serial.print(holi.sol);
  Serial.print('\n');
  Serial.print("temperatura terrario:");
  Serial.print(holi.terrario);
  Serial.print('\n');
  Serial.print("humedad : ");
  Serial.print(holi.h);
  Serial.print('\n');
  Serial.print("UV : ");
  Serial.print(Uv);
  Serial.print('\n');
  Serial.print("Foco Termico : ");
  Serial.print(Foco_termico);
  Serial.print('\n');
  Serial.print("Placa_termica : ");
  Serial.print(Placa_termica);
  Serial.print('\n');
  Serial.print("Catarata : ");
  Serial.print(Catarata);
  Serial.print('\n');
  Serial.print("------------FIN--------------- POST-----------");
 }

int ValidacionSensores(struct PostSensores holi){
  if(Sensor_sol > API_sol_max.toInt() && Sensor_terrario > API_temp_max.toInt()){return 5;}
  if(Sensor_sol < API_sol_max.toInt() && Sensor_terrario < API_temp_max.toInt()){return 6;}
  if(Sensor_sol > API_sol_max.toInt()) {return 1;}
  if(Sensor_terrario > API_temp_max.toInt()) {return 2;}
  if(Sensor_sol < API_sol_min.toInt()){return 3;}
  if(Sensor_terrario < API_sol_min.toInt()){return 4;}
  else {return 0;}
 }

void PostAlarma(String razon, int id_cliente ){
  Serial.print("id_cliente:");
  Serial.print(id_cliente);
  Serial.print("razon:");
  Serial.print(razon);
  Serial.print("hora: FALTA");
 
}


void setup() {
    Serial.begin(115200);
 // ---------------------------------------------------------------- CONECTARSE A INTERNET MEDIANTE EL CELULAR ------------------------------------------------------------------
 /*
  Serial.println("hola ESP32");
  Serial.print("comenzaremos a conectar a wifi");
  Serial.print("");
  //WIFI station mode
  WiFi.mode(WIFI_AP_STA);
  //Comienza SmartConfig
  WiFi.beginSmartConfig();
  //mensaje de espera del paquete del celular
  Serial.println("Esperando ip/password del celular ...");
  while(!WiFi.smartConfigDone()){
    delay(500);
    Serial.print(".");
    }
   Serial.println("");
   Serial.println("Llego ip/password del celular!");

   //esperando q wifi se conecte
   Serial.println("conectando ESP32 al WiFi...");
   while(WiFi.status() != WL_CONNECTED){
    delay(500);
    Serial.print(".");
    }
  Serial.println("WIFI CONNECTADO");
  Serial.print("IP Adress: ");
  Serial.println(WiFi.localIP());
  */
  //-------------------------------------GET CONFIGURACIÓN DEL CLIENTE -------------------------------------------------------
  Serial.print("como primer paso buscaremos la configuración del cliente de id: ");
  Getconfig(Id_cliente);
}


void loop() {
  Serial.print("---------------------------------------------------------------------------------------------------------------");
  Serial.print('\n');
  Serial.print("LOOP NUMERO:");
  Serial.print(Count_sol);
  Serial.print('\n');
  Serial.print("OBTENER TEMPERATURAS DE LOS SENSORES");
  Serial.print('\n');
  struct PostSensores datos;
  datos = LeerDatos();
  Serial.print('\n');
  Serial.print("POSTEAR TEMPERATURA DE LOS SENSORES ");
  Serial.print('\n');
  PostDatosApi(datos);
  Serial.print("OBTENER CONFIGURACIÓN DEL CLIENTE ");
  Getconfig(Id_cliente);
  Serial.print("----------------------------------------------------------------------------------------------------------------");
  Serial.print("VALIDACIONES");
  int validacion = ValidacionSensores(datos);
    if (validacion == 1){ //FALTA AGREGAR QUE EL CONTADOR SEA MENOR A 1 HORA, si se pasa de la hora, entonces que ponga el contador en 0 de nuevo
      String razon = "Temperatura Alta en Zona de Calor";
      PostAlarma(razon, Id_cliente);
          //if (esta automatizado haga algo)//ACTUALIZE EL ESTADO DE LOS SENSORES EN LA API
          //else filo que siga
         Count_sol = Count_sol + 1;
    }
    if (validacion == 2){
      String razon = "Temperatura Baja en Zona de Calor";
      PostAlarma(razon, Id_cliente);
          //if (esta automatizado haga algo)//ACTUALIZE EL ESTADO DE LOS SENSORES EN LA API
          //else filo que siga
          Count_sol = Count_sol + 1;
    }    
    if (validacion == 3){ 
      String razon = "Temperatura Alta en el Terrario";
      PostAlarma(razon, Id_cliente);
          //if (esta automatizado haga algo)//ACTUALIZE EL ESTADO DE LOS SENSORES EN LA API
          //else filo que siga
          Count_terrario = Count_terrario + 1;
    }
    if (validacion == 4){ //FALTA AGREGAR QUE EL CONTADOR SEA MENOR A 1 HORA
      String razon = "Temperatura Baja en el Terrario";
      PostAlarma(razon, Id_cliente);
          //if (esta automatizado haga algo)//ACTUALIZE EL ESTADO DE LOS SENSORES EN LA API
          //else filo que siga
          Count_terrario = Count_terrario + 1;
    }
    if (validacion == 5){ //FALTA AGREGAR QUE EL CONTADOR SEA MENOR A 1 HORA
      String razon = "Temperatura Alta tanto en Zona de Calor y Terrario";
      PostAlarma(razon, Id_cliente);
          //if (esta automatizado haga algo)//ACTUALIZE EL ESTADO DE LOS SENSORES EN LA API
          //else filo que siga
          Count_ambos = Count_ambos + 1;
    }
    if (validacion == 6){ //FALTA AGREGAR QUE EL CONTADOR SEA MENOR A 1 HORA
      String razon = "Temperatura Baja tanto en Zona de Calor y Terrario";
      PostAlarma(razon, Id_cliente);
          //if (esta automatizado haga algo)//ACTUALIZE EL ESTADO DE LOS SENSORES EN LA API
          //else filo que siga
          Count_ambos = Count_ambos + 1;
    }
    else{
       Count_sol = 0;
       Count_terrario = 0;
    }
  delay(5000); //DELAY 5 SEGUNDOS
}
//-----------------------------QA---------------------------
//VER EL CASO QUE DOS SENSORES SALGAN DEL MARGEN  
