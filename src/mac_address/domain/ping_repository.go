package domain
type NetworkRepository interface {
	GetArpTable() (map[string]string, error)
	ScanNetwork() ([]Device, error)
	GetLocalIP() (string, error)
	GetRegisteredDevices() ([]Device, error)
}