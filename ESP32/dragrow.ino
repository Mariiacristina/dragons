//LIBRERIAS
#include "WiFi.h"
#include "HTTPClient.h"
#include "OneWire.h"
#include "DallasTemperature.h"
#include "DHT.h"


//ONE WIRE - pin 15
#define ONE_WIRE_BUS 15
OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);
//OneWire ds(15); //para obtener direcciones

//sensor de humedad
#define DHTTYPE DHT21
const int DHTPin = 19;
DHT dht(DHTPin, DHTTYPE); //le paso pin y tipo

//direcciones de one wire
DeviceAddress sensor1 = { 0x28, 0x1A, 0x98, 0xA2, 0x8, 0x0, 0x0, 0x88 };
DeviceAddress sensor2 = { 0x28, 0x21, 0xC5, 0xA3, 0x8, 0x0, 0x0, 0x6B };


void setup() {
  Serial.begin(115200);
  sensors.begin();
  dht.begin();
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
  }

void loop() {  
  /*
  Serial.print("DIRECCIONES");
  byte i;
  byte addr[8];
  
  if (!ds.search(addr)) {
    Serial.println(" No more addresses.");
    Serial.print('\n');
    Serial.println();
    ds.reset_search();
    delay(250);
    return;
  }
  Serial.print(" ROM =");
  for (i = 0; i < 8; i++) {
    Serial.write(' ');
    Serial.print(addr[i], HEX);
  }
    delay(1000);
  */

  //SENSOR DE TEMPERATURA 
  Serial.print("ROBTENIENDO TEMPERATURAS...");
  sensors.requestTemperatures();
  Serial.println("DONE");
  
  Serial.print("Sensor 1(*C): ");
  Serial.print(sensors.getTempC(sensor1)); 
  Serial.print('\n');
  Serial.print("Sensor 2(*C): ");
  Serial.print(sensors.getTempC(sensor2)); 
  Serial.print('\n');

  //SENSOR DE HUMEDAD
  Serial.print("SENSOR DE HUMEDAD Y TEMPERATURA");
  float h = dht.readHumidity();
  float t = dht.readTemperature();

  Serial.print("Humidity: ");
  Serial.print(h);
  Serial.print(" %\t Temperature: ");
  Serial.print(t);
  Serial.print('\n');
  
 
  delay(2000);

}
