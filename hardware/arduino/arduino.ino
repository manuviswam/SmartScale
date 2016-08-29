#include "HX711.h"
#include "SoftwareSerial.h"

#define MAX_BITS 100                 // max number of bits 
#define WEIGAND_WAIT_TIME  3000      // time to wait for another weigand pulse.  

SoftwareSerial nodeSerial(5, 4); 

String url="http://10.132.127.212:10000/api/weight";
float weight = 0;
bool freshMeasure = true;

unsigned char databits[MAX_BITS];    // stores all of the data bits
unsigned char bitCount;              // number of bits currently captured
unsigned char flagDone;              // goes low when data is currently being captured
unsigned int weigand_counter;        // countdown until we assume there are no more bits 
unsigned long facilityCode=0;        // decoded facility code
unsigned long cardCode=0;            // decoded card code
 
// HX711.DOUT - pin #A1
// HX711.PD_SCK - pin #A0

HX711 scale(A1, A0);    // parameter "gain" is ommited; the default value 128 is used by the library

void ISR_INT0(){
  bitCount++;
  flagDone = 0;
  weigand_counter = WEIGAND_WAIT_TIME;  
}
 

void ISR_INT1(){
  databits[bitCount] = 1;
  bitCount++;
  flagDone = 0;
  weigand_counter = WEIGAND_WAIT_TIME;  
}
 
void setup(){
  pinMode(2, INPUT);     // DATA0 (INT0)
  pinMode(3, INPUT);     // DATA1 (INT1)

  nodeSerial.begin(9600);
  nodeSerial.println("$*"); 
  
  digitalWrite(4, HIGH);
  attachInterrupt(0, ISR_INT0, FALLING);  
  attachInterrupt(1, ISR_INT1, FALLING);
 
  weigand_counter = WEIGAND_WAIT_TIME;

  Serial.begin(9600);

  Serial.println("Before setting up the scale:");
  Serial.print("read: \t\t");
  Serial.println(scale.read());     // print a raw reading from the ADC

  Serial.print("read average: \t\t");
  Serial.println(scale.read_average(20));   // print the average of 20 readings from the ADC

  Serial.print("get value: \t\t");
  Serial.println(scale.get_value(5));   // print the average of 5 readings from the ADC minus the tare weight (not set yet)

  Serial.print("get units: \t\t");
  Serial.println(scale.get_units(5), 1);  // print the average of 5 readings from the ADC minus tare weight (not set) divided 
            // by the SCALE parameter (not set yet)  

  scale.set_scale(2280.f);                      // this value is obtained by calibrating the scale with known weights; see the README for details
  scale.tare();               // reset the scale to 0

  Serial.println("After setting up the scale:");

  Serial.print("read: \t\t");
  Serial.println(scale.read());                 // print a raw reading from the ADC

  Serial.print("read average: \t\t");
  Serial.println(scale.read_average(20));       // print the average of 20 readings from the ADC

  Serial.print("get value: \t\t");
  Serial.println(scale.get_value(5));   // print the average of 5 readings from the ADC minus the tare weight, set with tare()

  Serial.print("get units: \t\t");
  Serial.println(scale.get_units(5), 1);        // print the average of 5 readings from the ADC minus tare weight, divided 
            // by the SCALE parameter set with set_scale

  Serial.println("Readings:");
}
 
void loop(){
  // This waits to make sure that there have been no more data pulses before processing data
  if (!flagDone) {
    if (--weigand_counter == 0)
      flagDone = 1;  
  }
  else {
    scale.power_up();
    weight = scale.get_units(10)/10.0;
    scale.power_down();
    Serial.println(weight);
    if(freshMeasure && weight > 10.0) {
    
      freshMeasure = false;
      printBits();
    }
    if(!freshMeasure && weight < 10.0) {
      freshMeasure = true;
      weight = 0;
    }
  }
 
  // if we have bits and we the weigand counter went out
  if (bitCount > 0 && flagDone) {
    unsigned char i; 
    if (bitCount == 35){
      for (i=2; i<14; i++){
         facilityCode <<=1;
         facilityCode |= databits[i];
      }
      for (i=14; i<34; i++){
         cardCode <<=1;
         cardCode |= databits[i];
      }
      printBits();
    }
    else if (bitCount == 26){
      for (i=1; i<9; i++){
         facilityCode <<=1;
         facilityCode |= databits[i];
      }
      for (i=9; i<25; i++){
         cardCode <<=1;
         cardCode |= databits[i];
      }
      printBits();  
    }
    else {
     nodeSerial.println("$*"); 
    }
    // cleanup and get ready for the next card
    bitCount = 0;
    facilityCode = 0;
    cardCode = 0;
    for (i=0; i<MAX_BITS; i++){
       databits[i] = 0;
    }
  }
}
 
void printBits(){  
      scale.power_up();  
      delay(300);
      weight = scale.get_units(10)/10.0;
      scale.power_down();
      Serial.println(weight,1);
      Serial.println("print");
      url.concat("?internalNumber=");
      url.concat(cardCode);
      url.concat("&weight=");
      url.concat(weight);
      nodeSerial.println(url);
      url="http://10.132.127.212:10000/api/weight";
}
