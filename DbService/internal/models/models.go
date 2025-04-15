package models

import "time"

type CreateCartridge struct {
	Name          string    `json:"name"`
	Parameters    string    `json:"params"`
	BarcodeNumber int       `json:"barcodeNumber"`
	ReceivedFrom  string    `json:"receivedFrom"`
	Timestamp     time.Time `json:"timestamp"`
}

type DbTopicMessage struct {
	Action        string    `json:"action"`
	Name          string    `json:"name"`
	BarcodeNumber int       `json:"barcodeNumber"`
	Parameters    string    `json:"parameters"`
	NewStatus     string    `json:"newStatus"`
	Timestamp     time.Time `json:"timestamp"`
	ReceivedFrom  string    `json:"receivedFrom"`
	SendTo        string    `json:"sendTo"`
}

type UpdateCartridgeReceive struct {
	BarcodeNumber int       `json:"barcodeNumber"`
	NewStatus     string    `json:"newStatus"`
	Timestamp     time.Time `json:"timestamp"`
}

type UpdateCartridgeSend struct {
	BarcodeNumber int       `json:"barcodeNumber"`
	NewStatus     string    `json:"newStatus"`
	SendTo        string    `json:"sendTo"`
	Timestamp     time.Time `json:"timestamp"`
}

type DeleteCartridge struct {
	BarcodeNumber int `json:"barcodeNumber"`
}
