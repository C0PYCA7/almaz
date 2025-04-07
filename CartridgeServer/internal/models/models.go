package models

import "time"

type CartridgeModel struct {
	Name                        string    `json:"name"`
	Parameters                  string    `json:"params"`
	Status                      string    `json:"status"`
	ReceivedFrom                string    `json:"receivedFrom"`
	ReceivedFromSubdivisionDate time.Time `json:"receivedFromSubdivisionDate"`
	SendToRefillingDate         time.Time `json:"sendToRefillingDate,omitempty"`
	ReceivedFromRefillingDate   time.Time `json:"receivedFromRefillingDate,omitempty"`
	SendTo                      string    `json:"sendTo,omitempty"`
	SendToSubdivisionDate       time.Time `json:"sendToSubdivisionDate,omitempty"`
	BarcodeNumber               int32     `json:"barcodeNumber"`
}

type CreateCartridge struct {
	Name          string `json:"name"`
	Parameters    string `json:"params"`
	Status        string `json:"status"`
	ReceivedFrom  string `json:"receivedFrom"`
	BarcodeNumber int32  `json:"barcodeNumber"`
}
