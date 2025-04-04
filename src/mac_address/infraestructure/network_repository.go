package infraestructure

import (
	core "API3/core/mysql"
	"API3/src/mac_address/domain"
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-ping/ping"
)
type NetworkRepositoryImpl struct {
	conn *core.Conn_MySQL
}

func NewNetworkRepository() *NetworkRepositoryImpl {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &NetworkRepositoryImpl{conn: conn}

}

func (n *NetworkRepositoryImpl) GetArpTable() (map[string]string, error) {
	cmd := exec.Command("arp", "-a")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	arpTable := make(map[string]string)
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)\s+([a-fA-F0-9:-]+)`)
	matches := re.FindAllStringSubmatch(string(output), -1)

	for _, match := range matches {
		if len(match) == 3 {
			arpTable[match[1]] = match[2]
		}
	}
	return arpTable, nil
}

func (n *NetworkRepositoryImpl) GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipv4 := ipnet.IP.To4(); ipv4 != nil && !strings.HasPrefix(ipv4.String(), "169.254") {
				return ipv4.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no se pudo obtener la IP local válida")
}

func (n *NetworkRepositoryImpl) ScanNetwork() ([]domain.Device, error) {
	localIP, err := n.GetLocalIP()
	if err != nil {
		return nil, err
	}

	subnet := localIP[:strings.LastIndex(localIP, ".")]
	var results []domain.Device
	arpTable, _ := n.GetArpTable()

	var wg sync.WaitGroup
	resultChan := make(chan domain.Device, 254)

	for i := 1; i < 255; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			host := fmt.Sprintf("%s.%d", subnet, i)
			pinger, err := ping.NewPinger(host)
			if err != nil {
				return
			}
			pinger.Count = 1
			pinger.Timeout = time.Second
			pinger.SetPrivileged(true)
			err = pinger.Run()
			if err != nil {
				return
			}
			stats := pinger.Statistics()
			if stats.PacketsRecv > 0 {
				mac := arpTable[host]
				resultChan <- domain.Device{IP: host, MACAddress: mac}
			}
		}(i)
	}

	wg.Wait()
	close(resultChan)

	for res := range resultChan {
		results = append(results, res)
	}

	return results, nil
}
func (n *NetworkRepositoryImpl) GetRegisteredDevices() ([]domain.Device, error) {
	rows, err := n.conn.DB.Query("SELECT mac_address FROM smartWatch")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var devices []domain.Device
	for rows.Next() {
		var device domain.Device
		if err := rows.Scan(&device.MACAddress); err != nil {
			fmt.Println("rows: ", rows)
			return nil, err
		}
		devices = append(devices, device)
	}
	fmt.Println("devices:", devices)
	return devices, nil
}