package domain

type Device struct {
	IP         string `json:"ip"`
	MACAddress string `json:"mac_address"`
}