package observer

import (
	"DbService/internal/models"
	"encoding/json"
	"log"
)

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(msg *models.DbTopicMessage)
}

type Observer interface {
	Notify(msg *models.DbTopicMessage, strErr string)
}

type Observable struct {
	Observers []Observer
}

func NewObservable() *Observable {
	return &Observable{
		Observers: make([]Observer, 0),
	}
}

func (o *Observable) RegisterObserver(observer Observer) {
	o.Observers = append(o.Observers, observer)
}

func (o *Observable) RemoveObserver(observer Observer) {
	for i, obs := range o.Observers {
		if obs == observer {
			o.Observers = append(o.Observers[:i], o.Observers[i+1:]...)
			break
		}
	}
}

func (o *Observable) NotifyObservers(msg *models.DbTopicMessage, strErr string) {
	for _, observer := range o.Observers {
		observer.Notify(msg, strErr)
	}
}

type KafkaObserver struct {
	Topic    string
	Producer models.Sender
}

func (k *KafkaObserver) Notify(msg *models.DbTopicMessage, strErr string) {

	var message interface{}

	switch k.Topic {
	case "notificationTopic":
		message = models.NotificationTopicMessage{
			BarcodeNumber: msg.BarcodeNumber,
			Action:        msg.Action,
			Error:         strErr,
		}
	case "reportTopic":
		message = models.ReportTopicMessage{
			BarcodeNumber: msg.BarcodeNumber,
			Action:        msg.Action,
		}
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	k.Producer.SendMessage(k.Topic, data)
}
