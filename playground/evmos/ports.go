package evmos

import (
	"fmt"
	"log"
	"net"
	"slices"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
)

type Ports struct {
	// APP Ports
	P1317 int
	P8080 int
	P9090 int
	P9091 int
	P8545 int
	P8546 int
	P6065 int

	// Config Ports
	P26658 int
	P26657 int
	P6060  int
	P26656 int
	P26660 int
}

func NewPorts() *Ports {
	ports := []int{}
	for len(ports) < 13 {
		p, err := getAvailablePort()
		if err != nil {
			log.Panic("could not get available ports", err.Error())
		}
		// TODO: compare the ports with the rest of the config files
		if !slices.Contains(ports, p) {
			ports = append(ports, p)
		} else {
			log.Panic("it returned the same port twice")
		}
	}

	return &Ports{
		P1317:  ports[0],
		P8080:  ports[1],
		P9090:  ports[2],
		P9091:  ports[3],
		P8545:  ports[4],
		P8546:  ports[5],
		P6065:  ports[6],
		P26658: ports[7],
		P26657: ports[8],
		P6060:  ports[9],
		P26656: ports[10],
		P26660: ports[11],
	}
}

func NewPorts2(dbPorts []database.Port) *Ports {
	ports := []int{}
	for len(ports) < 13 {
		p, err := getAvailablePort()
		if err != nil {
			log.Panic("could not get available ports", err.Error())
		}
		port := int64(p)
		for _, v := range dbPorts {
			if v.P1317 == port {
				continue
			}
			if v.P8080 == port {
				continue
			}
			if v.P9090 == port {
				continue
			}
			if v.P9091 == port {
				continue
			}
			if v.P8545 == port {
				continue
			}
			if v.P8546 == port {
				continue
			}
			if v.P6065 == port {
				continue
			}
			if v.P26658 == port {
				continue
			}
			if v.P26657 == port {
				continue
			}
			if v.P6060 == port {
				continue
			}
			if v.P26656 == port {
				continue
			}
			if v.P26660 == port {
				continue
			}

		}
		if !slices.Contains(ports, p) {
			ports = append(ports, p)
		} else {
			log.Println("it returned the same port twice")
		}
	}

	return &Ports{
		P1317:  ports[0],
		P8080:  ports[1],
		P9090:  ports[2],
		P9091:  ports[3],
		P8545:  ports[4],
		P8546:  ports[5],
		P6065:  ports[6],
		P26658: ports[7],
		P26657: ports[8],
		P6060:  ports[9],
		P26656: ports[10],
		P26660: ports[11],
	}
}

func (e *Evmos) RestorePortsFromDB(port database.Port) {
	e.Ports = Ports{
		P1317:  int(port.P1317),
		P8080:  int(port.P8080),
		P9090:  int(port.P9090),
		P9091:  int(port.P9091),
		P8545:  int(port.P8545),
		P8546:  int(port.P8546),
		P6065:  int(port.P6065),
		P26658: int(port.P26658),
		P26657: int(port.P26657),
		P6060:  int(port.P6060),
		P26656: int(port.P26656),
		P26660: int(port.P26660),
	}
}

func getAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0") //nolint:gosec
	if err != nil {
		return 0, fmt.Errorf("could not find an available port: %w", err)
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

func (e *Evmos) SetPorts() error {
	appFile, err := e.openAppFile()
	if err != nil {
		return err
	}
	app := string(appFile)
	app = strings.Replace(app, "1317", fmt.Sprint(e.Ports.P1317), 1)
	app = strings.Replace(app, "8080", fmt.Sprint(e.Ports.P8080), 1)
	app = strings.Replace(app, "9090", fmt.Sprint(e.Ports.P9090), 1)
	app = strings.Replace(app, "9091", fmt.Sprint(e.Ports.P9091), 1)
	app = strings.Replace(app, "8545", fmt.Sprint(e.Ports.P8545), 1)
	app = strings.Replace(app, "8546", fmt.Sprint(e.Ports.P8546), 1)
	app = strings.Replace(app, "6065", fmt.Sprint(e.Ports.P6065), 1)
	if err := e.saveAppFile([]byte(app)); err != nil {
		return err
	}

	configFile, err := e.openConfigFile()
	if err != nil {
		return err
	}

	config := string(configFile)
	config = strings.Replace(config, "26656", fmt.Sprint(e.Ports.P26656), 1)
	config = strings.Replace(config, "26657", fmt.Sprint(e.Ports.P26657), 1)
	config = strings.Replace(config, "26658", fmt.Sprint(e.Ports.P26658), 1)
	config = strings.Replace(config, "26660", fmt.Sprint(e.Ports.P26660), 1)
	config = strings.Replace(config, "6060", fmt.Sprint(e.Ports.P6060), 1)
	if err := e.saveConfigFile([]byte(config)); err != nil {
		return err
	}

	return nil
}
