package models

var singleton chan []byte

func WeightChan() chan []byte {
        if singleton == nil {
            singleton = make(chan []byte)
        }
        return singleton;
}