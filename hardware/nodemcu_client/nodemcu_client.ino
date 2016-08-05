//example request : http://10.132.127.212:3000/order?type=juice&deviceID=123&userID=18427

#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <ESP8266WiFiMulti.h>
#include <ESP8266HTTPClient.h>

#include <SoftwareSerial.h>

SoftwareSerial arduino(D5, D7, false, 256);

#define buzzer D6
#define red D4
#define green D3
#define blue D2

ESP8266WiFiMulti WiFiMulti;
HTTPClient http;
        
String url;
String payload;

void setup() {

    pinMode(red, OUTPUT);
    pinMode(blue, OUTPUT);
    pinMode(green, OUTPUT);
    pinMode(buzzer, OUTPUT);

    digitalWrite(red, LOW);
    digitalWrite(green, HIGH);
    digitalWrite(blue, HIGH);
    digitalWrite(buzzer, LOW);

    Serial.begin(9600);
    arduino.begin(9600);
    Serial.println();
    Serial.println();
    Serial.println();

    for(uint8_t t = 4; t > 0; t--) {
        Serial.printf("[SETUP] WAIT %d...\n", t);
        Serial.flush();
        delay(1000);
    }

    WiFiMulti.addAP("twguest", "novo torch cycad ellen aqua");

}

void buzz(int delayTime){
    digitalWrite(buzzer, HIGH);
    delay(delayTime);
    digitalWrite(buzzer, LOW);  
}

void buzzInt(int delayTime, int repeat, int stopDelay){
    while(repeat>0){
    buzz(delayTime);
    delay(stopDelay);
    repeat--;
    }
}    
    

void loop() {
  
    digitalWrite(red, LOW);
    digitalWrite(blue, HIGH);
    digitalWrite(green,HIGH);
         
    if((WiFiMulti.run() == WL_CONNECTED)) {

     digitalWrite(red, HIGH);
     digitalWrite(red, HIGH);
     digitalWrite(blue, LOW);
     delay(500);

     if(arduino.available() > 0) {        
        
        url = arduino.readStringUntil('\n');
        String toBeSentURl = url.substring(0, url.length()-1);
        Serial.print("Got : ");
        Serial.println(toBeSentURl);
        
        Serial.print("[HTTP] begin...\n");
        http.begin(toBeSentURl); //HTTP
        Serial.print("[HTTP] GET...\n");
        int httpCode = http.GET();
        
        
        Serial.printf("[HTTP] GET... code: %d\n", httpCode);

        if(httpCode > 0) {
           
            
            if(httpCode == HTTP_CODE_OK) {
     
                String payload = http.getString();
                Serial.println(payload);

                
                digitalWrite(blue, LOW);
                digitalWrite(green, LOW);
                digitalWrite(red, LOW);
                delay(300);
                buzzInt(100,3,100);
                

             }
             else {

                Serial.printf("[HTTP] GET... failed, error: %s\n", http.errorToString(httpCode).c_str());
                digitalWrite(blue, LOW);
                digitalWrite(green, HIGH);
                digitalWrite(red, LOW);
                buzz(1500);
             }

           }
          
          http.end();
        }

        digitalWrite(red, HIGH);
        digitalWrite(blue, LOW);
        digitalWrite(green,HIGH);
        Serial.flush();
      }
}



