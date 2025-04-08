package models

import "time"

type CartridgeModel struct {
	Name                        string     `json:"name"`
	Parameters                  string     `json:"params"`
	Status                      string     `json:"status"`
	ReceivedFrom                string     `json:"receivedFrom"`
	ReceivedFromSubdivisionDate *time.Time `json:"receivedFromSubdivisionDate"`
	SendToRefillingDate         *time.Time `json:"sendToRefillingDate"`
	ReceivedFromRefillingDate   *time.Time `json:"receivedFromRefillingDate"`
	SendTo                      string     `json:"sendTo"`
	SendToSubdivisionDate       *time.Time `json:"sendToSubdivisionDate"`
	BarcodeNumber               int        `json:"barcodeNumber"`
}

type CreateCartridgeModel struct {
	Name          string `json:"name"`
	Parameters    string `json:"params"`
	ReceivedFrom  string `json:"receivedFrom"`
	BarcodeNumber int    `json:"barcodeNumber"`
}

type DbTopicMessage struct {
	Action    string         `json:"action"`
	Cartridge CartridgeModel `json:"cartridge"`
}
