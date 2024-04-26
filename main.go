package main

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	var client *modbus.ModbusClient
	var err error
	var data [16]uint16
	for {
		// for a TCP endpoint
		// (see examples/tls_client.go for TLS usage and options)
		client, err = modbus.NewClient(&modbus.ClientConfiguration{
			URL: "tcp://192.168.10.34:15001",
			// URL: "tcp://10.135.123.150:15001",
			// URL: "tcp://81.95.134.121:15001",
			// URL:     "tcp://127.0.0.1:15001",
			Timeout: 5 * time.Second,
		})

		if err != nil {
			panic("not new client " + err.Error())
		}
		client.SetUnitId(11)
		// now that the client is created and configured, attempt to connect
		for {
			err = client.Open()
			if err != nil {
				fmt.Println(time.Now().Format("15:04:05") + " open " + err.Error())
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}
		fmt.Println(time.Now().Format("15:04:05") + " connecting....")
		for i := 0; i < 60; i++ {
			reg16, err := client.ReadRegisters(0, 4, modbus.HOLDING_REGISTER)
			if err != nil {
				fmt.Printf("%s read holds %s \n", time.Now().Format("15:04:05"), err.Error())
				break
			} else {
				k := 0
				j := 0
				for i := 0; i < 16; i++ {
					data[i] = (reg16[k] >> j) & 0xf
					j += 4
					if j > 12 {
						j = 0
						k++
					}
				}
				fmt.Printf("%s value: %v \n", time.Now().Format("15:04:05"), data) // as unsigned integer
			}
			time.Sleep(100 * time.Millisecond)
		}
		err = client.Close()
		if err != nil {
			fmt.Printf("%s close %s \n", time.Now().Format("15:04:05"), err.Error())
		}
		fmt.Println(time.Now().Format("15:04:05") + " closed....")
		time.Sleep(1 * time.Second)
	}
}
