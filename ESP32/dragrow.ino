//LIBRERIAS
#include "WiFi.h"
#include <HTTPClient.h>
//#include "time.h"
#include <ArduinoJson.h>
#include <OneWire.h>
#include "DallasTemperature.h"
#include "DHT.h"


//CLIENTE
int Id_cliente = 1;
//TIEMPO
int Tiempo = 20;

//variables para el time
const char* ntpServer = "Ntp.shoa.cl";
const long  gmtOffset_sec = -3*3600;
const int   daylightOffset_sec = 3600;
//-----------------------------------------------------------------------------------SENSORES DE ESP---------------------------------------------------------------------------------------------------
#define ONE_WIRE_BUS 15
OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);
//OneWire ds(15); //para obtener direcciones

//sensor de humedad
#define DHTTYPE DHT21
const int DHTPin = 19;
DHT dht(DHTPin, DHTTYPE); //le paso pin y tipo


int Sensor_sol;
int Sensor_terrario;
int Sensor_humedad;

//variables de entrada enchufes (numero del pin GPIO)
int entradaplaca = 9; 
int entradabombillo = 8;
int entradacascada = 7;
int entradauv = 16;
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
int     API_auto_sol      ;
int     API_auto_terrario ;
int     API_auto_humedad  ;
int     API_auto_luz      ;

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
  Serial.print("OBTENER CONFIGURACIÓN DEL CLIENTE ");
  Serial.print("obteniendo configuracion para el id :");
  Serial.print(id);
  Serial.print('\n');
  Serial.print("----------------------GET CONFIG-------------------------");
     HTTPClient http;
   http.addHeader("Content-Type", "application/json");
   http.begin("http://192.168.43.97:8000/esp/1"); //URL DE LA API 
   int httpCode = http.GET();
   if (httpCode > 0) {
        String res = http.getString();
        Serial.println(httpCode);
        DynamicJsonBuffer jsonBuffer;
        JsonObject& parsed = jsonBuffer.parseObject(res);
        Serial.println(parsed["sol_max"].as<String>());
        API_sol_max       = parsed["sol_max"].as<String>();      
        API_sol_min       = parsed["sol_min"].as<String>();      
        API_temp_max      = parsed["temp_max"].as<String>();    
        API_temp_min      = parsed["temp_min "].as<String>();   
        API_humedad_min   = parsed["humedad_min"].as<String>();
        API_uv_inicio     = parsed["uv_inicio"].as<String>();
        API_uv_tiempo     = parsed["uv_tiempo"].as<String>();
        API_catarata_on   = parsed["catarata_on"].as<String>();
        API_catarata_off  = parsed["catarata_off"].as<String>();
        API_uv            = parsed["uv"].as<int>();
        API_focoTermico   = parsed["focoTermico "].as<int>();
        API_placaTermica  = parsed["placaTermica"].as<int>();
        API_catarata      = parsed["catarata"].as<int>();
        API_auto_sol      = parsed["auto_sol"].as<int>();
        API_auto_terrario = parsed["auto_terrario"].as<int>();
        API_auto_humedad  = parsed["auto_humedad"].as<int>();
        API_auto_luz      = parsed["auto_luz"].as<int>();
        
   }else {
    Serial.println("Error on HTTP request") ;
  }
  http.end();
  
 }

String Gethora()
{
  struct tm timeinfo;
  if(!getLocalTime(&timeinfo)){
    Serial.println("Failed to obtain time");
  }
  char hola[80];
  strftime(hola,80,"%A, %B %d %Y %H:%M:%S",&timeinfo);
  return String(hola);
}

void UpdateObjetos(){
  Serial.print("UPDATE OBJETOS - API");
  Serial.print('\n');
  if(API_uv == 1){Uv = 1;}
  Serial.print(API_uv);
  if(API_focoTermico == 1){Foco_termico = 1;}
  if(API_placaTermica == 1){Placa_termica = 1;}
  if(API_uv == 0){Uv = 0;}
  if(API_focoTermico == 0){Foco_termico = 0;}
  if(API_placaTermica == 0){Placa_termica = 0;}
}

void GestionarObjetos(){
  Serial.print("UPDATE OBJETOS ESP");
  if (Uv == 1)            {digitalWrite(entradauv, HIGH);} //PRENDER UV
  if (Foco_termico == 1)  {digitalWrite(entradabombillo, HIGH);} // PRENDER FOCO TERMICO
  if(Placa_termica == 1)  {digitalWrite(entradaplaca, HIGH);} // PRENDER PLACA TERMICA
  if(Catarata == 1)       {digitalWrite(entradacascada, HIGH);} // PRENDER CATARATA
  if (Uv == 0)            {digitalWrite(entradauv, HIGH);} //APAGAR UV
  if (Foco_termico == 0)  {digitalWrite(entradabombillo, LOW);} // APAGAR FOCO TERMICO
  if(Placa_termica == 0)  {digitalWrite(entradaplaca, LOW);} // APAGAR PLACA TERMICA
  if(Catarata == 0)       {digitalWrite(entradacascada, LOW);} // APAGAR CATARATA
  
  }
struct PostSensores LeerDatos(){
  //direcciones de one wire
  DeviceAddress sensor1 = { 0x28, 0x1A, 0x98, 0xA2, 0x8, 0x0, 0x0, 0x88 };
  DeviceAddress sensor2 = { 0x28, 0x21, 0xC5, 0xA3, 0x8, 0x0, 0x0, 0x6B };
  float h = dht.readHumidity();

  Serial.print("OBTENER TEMPERATURAS DE LOS SENSORES");
  Sensor_sol = sensors.getTempC(sensor1);
  Sensor_terrario = sensors.getTempC(sensor2);
  Sensor_humedad = (int)h;

  
  struct PostSensores holi;
  holi.sol = Sensor_sol; // LEER TEMP SOL DE LOS SENSORES
  holi.terrario = Sensor_terrario; // LEER TEMP TERRARIO DE LOS SENSORES
  holi.h = Sensor_humedad; // LEER HUMEDAD DE LOS SENSORES DE LOS SENSORES
  return holi;
 }

 void PostDatosApi (struct PostSensores holi){
  Serial.print("POSTEAR TEMPERATURA DE LOS SENSORES ");
  const String endpoint = "http://192.168.43.97:8000/esp/";
  const String key = String(Id_cliente);
  HTTPClient http;
  http.begin(endpoint + key);
  http.addHeader("Content-Type", "application/json");
  String objeto_post = "{\"id_cliente\":"+String(Id_cliente)+",\"esp_sol\":"+String(holi.sol)+",\"esp_terrario\":"+String(holi.terrario)+",\"esp_humedad\":"+String(holi.h)+",\"esp_uv\":"+String(Uv)+",\"esp_focotermico\":"+String(Foco_termico)+",\"esp_placatermica\":"+String(Placa_termica)+",\"esp_catarata\":"+String(Catarata)+"}";
  int httpResponseCode = http.POST(objeto_post);
  if(httpResponseCode>0){
    String response = http.getString(); 
    Serial.println(httpResponseCode);   
 }else{
     Serial.print("Error on sending POST: ");
    Serial.println(httpResponseCode);
 }
 http.end();
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
  Serial.print("------------POST ALARMA-----------");
  String hora_alarma = Gethora();
  const String endpoint = "http://192.168.43.97:8000/alarma/";
  const String key = String(Id_cliente);
  HTTPClient http;
  http.begin(endpoint + key);
  http.addHeader("Content-Type", "application/json");
  String objeto_alarma = "{\"id_cliente\":"+String(Id_cliente)+",\"hora\":\""+hora_alarma+"\",\"razon\":\""+razon+"\"}";
  Serial.print(objeto_alarma);
  int httpResponseCode = http.POST(objeto_alarma);
  if(httpResponseCode>0){
    String response = http.getString(); 
    Serial.println(httpResponseCode);  
 }else{
     Serial.print("Error on sending POST: ");
    Serial.println(httpResponseCode);
 }
 http.end();
 Serial.print('\n');
}

void GestionarCascada(){
  Count_catarata_on = Count_catarata_on +1;
  if(Count_catarata_on >= (API_catarata_on.toInt() * 720 ) ){
    Count_catarata_off = Count_catarata_off + 1;
    Catarata = 0;
    GestionarObjetos();
    if(Count_catarata_off >= (API_catarata_off.toInt() * 720) ){
      Count_catarata_on = 0;
      Count_catarata_off = 0;
      Catarata = 1;
      GestionarObjetos();
      }
    }
   }
   
void GestionarLuzUv(){
  if(API_auto_luz == 1){
      if (API_uv_inicio.toInt() == Tiempo){
        Uv = 1; 
        GestionarObjetos();
        }
      }
      if (API_uv_tiempo.toInt() == Tiempo){
        Uv = 0; 
        GestionarObjetos();
        }
  }

void setup() {
    Serial.begin(115200);
      sensors.begin();
      dht.begin();
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
  const char* ssid = "hector";
  const char* password =  "1111222233";
  WiFi.begin(ssid, password); 
 
  while (WiFi.status() != WL_CONNECTED) { //Check for the connection
    delay(1000);
    Serial.println("Connecting to WiFi..");
  }
 
  Serial.println("Connected to the WiFi network");
  configTime(gmtOffset_sec, daylightOffset_sec, ntpServer);
  
  //--------------------------------------------------------------------------------GET CONFIGURACIÓN DEL CLIENTE --------------------------------------------------------------------------------------------
  Serial.print("como primer paso buscaremos la configuración del cliente de id: ");
  Getconfig(Id_cliente);
  Serial.print('\n');
  //----------------------------------------------------------------------------------ACTUALIZAR OBJETOS RESPECTO A LA API-----------------------------------------------------------------------------
  Serial.print("Actualizar objetos respecto a la API");
  UpdateObjetos();
  Serial.print('\n');
  //---------------------------------------------------------------------------ACTUALIZAR OBJETOS ESP --------------------------------------------------------------------------------------
  Serial.print("actualizar objetos de la esp");
  Serial.print('\n');
  GestionarObjetos();
}




void loop() {
  Serial.print("---------------------------------------------------------------------------------------------------------------");
  Serial.print('\n');
  Serial.print("LOOP NUMERO:");
  Serial.print(Loop);
  Serial.print('\n');
  Serial.print("COUNT AMBOS:");
  Serial.print(Count_ambos);
  Serial.print('\n');
  //--------------------------------------------------------------------------LEER SENSORES---------------------------------------------------------------------------------------------------------------
  struct PostSensores datos;
  datos = LeerDatos();
  //------------------------------------------------------------------------POST SENSORES----------------------------------------------------------------------------------------------------------------
  PostDatosApi(datos);
  //------------------------------------------------------------------------GET CONFIGURACIÓN DEL CLIENTE-------------------------------------------------------------------------------------------------
  Getconfig(Id_cliente);
  //-----------------------------------------------------------------------ACTUALIZAR OBJETOS CON RESPECTO A LA API--------------------------------------------------------------------------------------
  UpdateObjetos();
  //---------------------------------------------------------------------------ACTUALIZAR OBJETOS ESP --------------------------------------------------------------------------------------
  GestionarObjetos();
  //------------------------------------------------------------------------VALIDACIONES------------------------------------------------------------------------------------------------------------------
  
  
  Serial.print("VALIDACIONES");
  Serial.print('\n');
  int validacion = ValidacionSensores(datos);
  //--------------ENVIA POST CADA 30 MIN -----------------
    if (validacion == 1){
       if(Count_sol == 0 || Count_sol > 360){
      String razon = "Temperatura Alta en Zona de Calor";
      PostAlarma(razon, Id_cliente);
          if(API_auto_sol == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 0;
            GestionarObjetos();
          }
          Count_sol = Count_sol + 1;
          if (Count_sol > 360){Count_sol = 0;}
       }
       else{Count_sol = Count_sol + 1;}
    }
    if (validacion == 2){
      if(Count_sol == 0 || Count_sol > 360){
      String razon = "Temperatura Baja en Zona de Calor";
      PostAlarma(razon, Id_cliente);
          if(API_auto_sol == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 1;
            GestionarObjetos();
          }
          Count_sol = Count_sol + 1;
          if (Count_sol > 360){Count_sol = 0;}
      }
      else{Count_sol = Count_sol + 1;}
    }    
    if (validacion == 3){
        if(Count_terrario == 0 || Count_terrario > 360){ 
      String razon = "Temperatura Alta en el Terrario";
      PostAlarma(razon, Id_cliente);
          if(API_auto_terrario == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 0;
            GestionarObjetos();
          }
          Count_terrario = Count_terrario + 1;
          if (Count_terrario > 360){Count_terrario = 0;}
        }
        else{Count_terrario = Count_terrario + 1;}       
    }
    if (validacion == 4){
        if(Count_terrario == 0 || Count_terrario > 360){ 
      String razon = "Temperatura Baja en el Terrario";
      PostAlarma(razon, Id_cliente);
          if (API_auto_terrario == 1){
            Serial.print("FOCO TERMICO ESTA AUTOMATIZADO");
            Serial.print('\n');
            Foco_termico = 1;
            GestionarObjetos();
          }
          Count_terrario = Count_terrario + 1;
          if (Count_terrario > 360){Count_terrario = 0;}
        }
        else{Count_terrario = Count_terrario + 1;}
    }
    if (validacion == 5){
       if (Count_ambos == 0 || Count_ambos > 360){
       String razon = "Temperatura Alta tanto en Zona de Calor y Terrario";
      PostAlarma(razon, Id_cliente);
          if(API_auto_sol == 1 || API_auto_terrario == 1){
            Serial.print("FOCO TERMICO O PLACA TERMICA ESTAN AUTOMATIZADOS");
            Serial.print('\n');
            Foco_termico = 0;
            Placa_termica = 0;
            GestionarObjetos();
            }
           Count_ambos = Count_ambos + 1;
           if (Count_ambos > 5){Count_ambos = 360;}
        }
        else{Count_ambos = Count_ambos + 1;}
        }
    if (validacion == 6){
        if(Count_ambos == 0 || Count_ambos > 360){
          String razon = "Temperatura Baja tanto en Zona de Calor y Terrario";
          PostAlarma(razon, Id_cliente);
         if(API_auto_sol == 1 || API_auto_terrario == 1){
            Serial.print("FOCO TERMICO O PLACA TERMICA ESTAN AUTOMATIZADOS");
            Serial.print('\n');
            Foco_termico = 1;
            Placa_termica = 1;
            GestionarObjetos();
             }
            Count_ambos = Count_ambos + 1;
            if (Count_ambos > 360){Count_ambos = 0;}
        }
        else{Count_ambos = Count_ambos + 1;}
    }
    if(validacion == 0){
       Serial.print('\n');
       Serial.print("TODO EN ORDEN");
       Serial.print('\n');
       Count_sol = 0;
       Count_terrario = 0;
       Count_ambos = 0;
    }

//-------------------------------------------------------------------------------GESTIONAR CATARATA------------------------------------------------------------------------------------------------------------------
  GestionarCascada();
//---------------------------------------------------------------------------------GESTIONAR LUZ UV -----------------------------------------------------------------------------------------------------------------
  GestionarLuzUv();
  Serial.print(Sensor_sol);
  Serial.print(Sensor_terrario);
  Serial.print(Sensor_humedad);
  Serial.print("-----------------------");
  Serial.print('\n');
  Serial.print(Uv);
    Serial.print('\n');
  Serial.print(Foco_termico);
    Serial.print('\n');
  Serial.print(Placa_termica);
  Loop = Loop + 1;
  delay(5000); //DELAY 5 SEGUNDOS
}
