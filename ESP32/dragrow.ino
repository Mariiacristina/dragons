//LIBRERIAS
#include "WiFi.h"

//CLIENTE
int Id_cliente = 1;
//TIEMPO
int Tiempo = 20;

//-----------------------------------------------------------------------------------SENSORES DE ESP---------------------------------------------------------------------------------------------------
int Sensor_sol = 40;
int Sensor_terrario = 30;
int Sensor_humedad = 60;

//-----------------------------------------------------------------------------------OBJETOS DE ESP-----------------------------------------------------------------------------------------------------
int Placa_termica = 0;
int Foco_termico = 0;
int Uv = 0;
int Catarata = 0;


//-------------------------------------------------------------------------------VARIABLES DEL USUARIO--------------------------------------------------------------------------------------------------
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
int     API_auto_sol  = 0    ;
int     API_auto_terrario = 0 ;
int     API_auto_humedad = 0  ;
int     API_auto_luz = 0      ;

//-----------------------------------------------------------------------------------CONTADORES----------------------------------------------------------------------------------------------------------
int Loop = 0;
int Count_sol = 0; //puede ser sol max o sol min pero no ambos
int Count_terrario = 0; // puede ser terrario max temp o min pero no ambos
int Count_ambos = 0; // ambos sensores estan fuera de margen
int Count_catarata_on = 0;
int Count_catarata_off = 0;

//------------------------------------------------------------------------------------ESTRUCTURAS-----------------------------------------------------------------------------------------------------
struct PostSensores {
   int sol;
   int terrario;
   int h;
};

//-------------------------------------------------------------------------------------FUNCIONES--------------------------------------------------------------------------------------------------------
int Getconfig (int id){
  Serial.print("obteniendo configuracion para el id :");
  Serial.print(id);
 }

void UpdateObjetos(){
  if(API_uv == 1){Uv = 1;}//PRENDER UV
  if(API_focoTermico == 1){Foco_termico = 1;}//PRENDER FOCO TERMICO
  if(API_placaTermica == 1){Placa_termica = 1;}//PRENDER PLACA TERMICA
 //if(API_catarata == 1){Catarata = 1;}//PRENDER CATARATA
  if(API_uv == 0){Uv = 0;}//APAGAR UV
  if(API_focoTermico == 0){Foco_termico = 0;}//APAGAR FOCO TERMICO
  if(API_placaTermica == 0){Placa_termica = 0;}//APAGAR PLACA TERMICA
//if(API_catarata == 0){Catarata = 0;}//APAGAR CATARATA
 }

struct PostSensores LeerDatos(){
  struct PostSensores holi;
  holi.sol = Sensor_sol; // LEER TEMP SOL DE LOS SENSORES
  holi.terrario = Sensor_terrario; // LEER TEMP TERRARIO DE LOS SENSORES
  holi.h = Sensor_humedad; // LEER HUMEDAD DE LOS SENSORES DE LOS SENSORES
  return holi;
 }

 void PostDatosApi (struct PostSensores holi){
  Serial.print("------------POST-----------");
  Serial.print('\n');
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
  Serial.print("------------FIN---------");
 }

int ValidacionSensores(struct PostSensores holi){
  if(Sensor_sol > API_sol_max.toInt() && Sensor_terrario > API_temp_max.toInt()){return 5;}
  if(Sensor_sol < API_sol_max.toInt() && Sensor_terrario < API_temp_max.toInt()){return 6;}
  if(Sensor_sol > API_sol_max.toInt()) {return 1;}
  if(Sensor_sol < API_sol_min.toInt()){return 2;}
  if(Sensor_terrario > API_temp_max.toInt()) {return 3;}
  if(Sensor_terrario < API_sol_min.toInt()){return 4;}
  else {return 0;}
 }

void PostAlarma(String razon, int id_cliente ){
  Serial.print('\n');
  Serial.print("id_cliente:");
  Serial.print(id_cliente);
  Serial.print('\n');
  Serial.print("razon:");
  Serial.print(razon);
  Serial.print('\n');
  Serial.print("hora: FALTA");
  Serial.print('\n');
 
}

void GestionarCascada(){
  Count_catarata_on = Count_catarata_on +1;
  if(Count_catarata_on >= (API_catarata_on.toInt() * 720 ) ){
    Count_catarata_off = Count_catarata_off + 1;
    if(Count_catarata_off >= (API_catarata_off.toInt() * 720) ){
      Count_catarata_on = 0;
      Count_catarata_off = 0;
      }
    }
   }
   
void GestionarLuzUv(){
  if(API_auto_luz == 1){
      if (API_uv_inicio.toInt() == Tiempo){Uv = 1;} // PRENDER FOCO
      if (API_uv_tiempo.toInt() == Tiempo){Uv = 0;}// APAGAR FOCO
  }
 } 

void setup() {
    Serial.begin(115200);
 // ----------------------------------------------------------------------- CONECTARSE A INTERNET MEDIANTE EL CELULAR --------------------------------------------------------------------------------
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
  //--------------------------------------------------------------------------------GET CONFIGURACIÓN DEL CLIENTE --------------------------------------------------------------------------------------------
  Serial.print("como primer paso buscaremos la configuración del cliente de id: ");
  Getconfig(Id_cliente);
  //----------------------------------------------------------------------------------ACTUALIZAR OBJETOS RESPECTO A LA API-----------------------------------------------------------------------------
  Serial.print("Actualizar objetos respecto a la API");
  UpdateObjetos();
  Serial.print("holi");
}


void loop() {
  Serial.print("---------------------------------------------------------------------------------------------------------------");
  Serial.print('\n');
  Serial.print("LOOP NUMERO:");
  Serial.print(Loop);
  Serial.print('\n');
  //--------------------------------------------------------------------------LEER SENSORES---------------------------------------------------------------------------------------------------------------
  Serial.print("OBTENER TEMPERATURAS DE LOS SENSORES");
  Serial.print('\n');
  struct PostSensores datos;
  datos = LeerDatos();
  Serial.print('\n');
  //------------------------------------------------------------------------POST SENSORES----------------------------------------------------------------------------------------------------------------
  Serial.print("POSTEAR TEMPERATURA DE LOS SENSORES ");
  Serial.print('\n');
  PostDatosApi(datos);
  Serial.print('\n');
  //------------------------------------------------------------------------GET CONFIGURACIÓN DEL CLIENTE-------------------------------------------------------------------------------------------------
  Serial.print("OBTENER CONFIGURACIÓN DEL CLIENTE ");
  Serial.print('\n');
  Getconfig(Id_cliente);
  //-----------------------------------------------------------------------ACTUALIZAR OBJETOS CON RESPECTO A LA API--------------------------------------------------------------------------------------
  UpdateObjetos();
  Serial.print('\n');
  //------------------------------------------------------------------------VALIDACIONES------------------------------------------------------------------------------------------------------------------
  Serial.print("VALIDACIONES");
  Serial.print('\n');
  int validacion = ValidacionSensores(datos);
  //--------------ENVIA POST CADA 30 MIN -----------------
    if (validacion == 1 && (Count_sol == 0 || Count_sol > 360) ){
      String razon = "Temperatura Alta en Zona de Calor";
      PostAlarma(razon, Id_cliente);
          if(API_auto_sol == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 0;
            //APAGAR FOCO
          }
          if (Count_sol > 360){Count_sol = 0;}
          else {Count_sol = Count_sol + 1;}
    }
    if (validacion == 2 && (Count_sol == 0 || Count_sol > 360) ){
      String razon = "Temperatura Baja en Zona de Calor";
      PostAlarma(razon, Id_cliente);
          if(API_auto_sol == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 1;
            //PRENDER FOCO TERMICO
          }
          if (Count_sol > 360){Count_sol = 0;}
          else {Count_sol = Count_sol + 1;}
    }    
    if (validacion == 3 && (Count_terrario == 0 || Count_terrario > 360) ){ 
      String razon = "Temperatura Alta en el Terrario";
      PostAlarma(razon, Id_cliente);
          if(API_auto_terrario == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 0;
            //APAGAR FOCO TERMICO
          }
          if (Count_terrario > 360){Count_terrario = 0;}
          else {Count_terrario = Count_terrario + 1;}          
    }
    if (validacion == 4 && (Count_terrario == 0 || Count_terrario > 360) ){ 
      String razon = "Temperatura Baja en el Terrario";
      PostAlarma(razon, Id_cliente);
          if (API_auto_terrario == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 1;
            //PRENDER FOCO TERMICO
          }
            if (Count_terrario > 360){Count_terrario = 0;}
            else {Count_terrario = Count_terrario + 1;}   
    }
    if (validacion == 5 && (Count_ambos == 0 || Count_ambos > 360) ){ 
      String razon = "Temperatura Alta tanto en Zona de Calor y Terrario";
      PostAlarma(razon, Id_cliente);
          if(API_auto_sol == 1 || API_auto_terrario == 1){
            Serial.print("FOCO TERMICO O PLACA TERMICA ESTAN AUTOMATIZADOS");
            Serial.print('\n');
            Foco_termico = 0;
            Placa_termica = 0;
            //APAGAR FOCO TERMICO
            //APAGAR PLACA TERMICA
            }
          Count_ambos = Count_ambos + 1;
          if (Count_ambos > 360){Count_ambos = 0;}
    }
    if (validacion == 6 && (Count_sol == 0 || Count_sol > 360)){
      String razon = "Temperatura Baja tanto en Zona de Calor y Terrario";
      PostAlarma(razon, Id_cliente);
         if(API_auto_sol == 1 || API_auto_terrario == 1){
            Serial.print("FOCO TERMICO O PLACA TERMICA ESTAN AUTOMATIZADOS");
            Serial.print('\n');
            Foco_termico = 1;
            Placa_termica = 1;
            //PRENDER FOCO TERMICO
            //PRENDER PLACA TERMICA
         }
            if (Count_ambos > 360){Count_ambos = 0;}
            Count_ambos = Count_ambos + 1;
    }
    if(validacion == 0){
       Serial.print('\n');
       Serial.print("TODO EN ORDEN");
       Serial.print('\n');
       Count_sol = 0;
       Count_terrario = 0;
       Count_ambos = 0;
    }
//  -------------------------------------------------------------------------------TERMINA VALIDACION----------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------GESTIONAR CATARATA------------------------------------------------------------------------------------------------------------------
  Serial.print("Ahora vamos a gestionar CATARATA");
  Serial.print('\n');
  GestionarCascada();
//--------------------------------------------------------------------GESTIONAR LUZ UV -----------------------------------------------------------------------------------------------------------------
  Serial.print("Ahora vamos a gestionar LUZ UV");
  Serial.print('\n');
  GestionarLuzUv();
//---------------------------------------------------------------------IMPRIMIR COUNTS-----------------------------------------------------------------------------------------------------------------
  Serial.print("contador alarma sol: ");
  Serial.print(Count_sol);
  Serial.print('\n');
  Serial.print("contador alarma terrario: ");
  Serial.print(Count_terrario);
  Serial.print('\n');
  Serial.print("contador alarma ambos: ");
  Serial.print(Count_ambos);
  Serial.print('\n');
  Loop = Loop + 1;
  delay(5000); //DELAY 5 SEGUNDOS
}
