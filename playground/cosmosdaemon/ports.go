package cosmosdaemon

import (
	"context"
	"fmt"
	"log"
	"net"
	"slices"

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

func getAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0") //nolint:gosec
	if err != nil {
		return 0, fmt.Errorf("could not find an available port: %w", err)
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

func (d *Daemon) AssignPorts(queries *database.Queries) error {
	ports, err := newPorts(queries)
	if err != nil {
		return err
	}
	d.Ports = ports
	return nil
}

func newPorts(queries *database.Queries) (*Ports, error) {
	dbPorts, err := queries.GetAllPorts(context.Background())
	if err != nil {
		return nil, err
	}

	ports := []int{}

OUTER:
	for len(ports) < 13 {
		p, err := getAvailablePort()
		if err != nil {
			log.Panic("could not get available ports", err.Error())
		}
		port := int64(p)
		for _, v := range dbPorts {
			if v.P1317 == port {
				continue OUTER
			}
			if v.P8080 == port {
				continue OUTER
			}
			if v.P9090 == port {
				continue OUTER
			}
			if v.P9091 == port {
				continue OUTER
			}
			if v.P8545 == port {
				continue OUTER
			}
			if v.P8546 == port {
				continue OUTER
			}
			if v.P6065 == port {
				continue OUTER
			}
			if v.P26658 == port {
				continue OUTER
			}
			if v.P26657 == port {
				continue OUTER
			}
			if v.P6060 == port {
				continue OUTER
			}
			if v.P26656 == port {
				continue OUTER
			}
			if v.P26660 == port {
				continue OUTER
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
	}, nil
}

func (d *Daemon) RestorePortsFromDB(port database.Port) {
	d.Ports = &Ports{
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

// func (d *Dameon) SetPorts() error {
// 	appFile, err := d.openAppFile()
// 	if err != nil {
// 		return err
// 	}
// 	app := string(appFile)
// 	app = strings.Replace(app, "1317", fmt.Sprint(d.Ports.P1317), 1)
// 	app = strings.Replace(app, "8080", fmt.Sprint(d.Ports.P8080), 1)
// 	app = strings.Replace(app, "9090", fmt.Sprint(d.Ports.P9090), 1)
// 	app = strings.Replace(app, "9091", fmt.Sprint(d.Ports.P9091), 1)
// 	app = strings.Replace(app, "8545", fmt.Sprint(d.Ports.P8545), 1)
// 	app = strings.Replace(app, "8546", fmt.Sprint(d.Ports.P8546), 1)
// 	app = strings.Replace(app, "6065", fmt.Sprint(d.Ports.P6065), 1)
// 	if err := d.saveAppFile([]byte(app)); err != nil {
// 		return err
// 	}
//
// 	configFile, err := d.openConfigFile()
// 	if err != nil {
// 		return err
// 	}
//
// 	config := string(configFile)
// 	config = strings.Replace(config, "26656", fmt.Sprint(d.Ports.P26656), 1)
// 	config = strings.Replace(config, "26657", fmt.Sprint(d.Ports.P26657), 1)
// 	config = strings.Replace(config, "26658", fmt.Sprint(d.Ports.P26658), 1)
// 	config = strings.Replace(config, "26660", fmt.Sprint(d.Ports.P26660), 1)
// 	config = strings.Replace(config, "6060", fmt.Sprint(d.Ports.P6060), 1)
// 	if err := d.saveConfigFile([]byte(config)); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
