package application

import (
	"API3/src/mac_address/domain"
	"fmt"
	"strings"
)

type CompareMacsUseCase struct {
	Repo domain.NetworkRepository
}

func NewCompareMacsUseCase(repo domain.NetworkRepository) *CompareMacsUseCase {
	return &CompareMacsUseCase{Repo: repo}
}

func (c *CompareMacsUseCase) Execute() ([]string, error) {
	registeredDevices, err := c.Repo.GetRegisteredDevices()
	if err != nil {
		return nil, err
	}

	arpTable, err := c.Repo.GetArpTable()
	if err != nil {
		fmt.Println("Error obteniendo ARP table:", err)
		return nil, err
	}

	registeredMacs := make(map[string]bool)
	for _, device := range registeredDevices {
		normalizedMac := strings.ToLower(strings.ReplaceAll(device.MACAddress, "-", ":"))
		registeredMacs[normalizedMac] = true
	}

	var matchedDevices []string
	for ip, mac := range arpTable {
		normalizedMac := strings.ToLower(strings.ReplaceAll(mac, "-", ":"))
		if _, exists := registeredMacs[normalizedMac]; exists {
			fmt.Println(ip)
			matchedDevices = append(matchedDevices, fmt.Sprintf("Direccion MAC: %s", normalizedMac))
		}
	}
	return matchedDevices, nil

}
