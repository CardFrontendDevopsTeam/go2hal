package service

import (
	g "github.com/soniah/gosnmp"
	"log"
	"os"
)

func SendSNMPTestMessage() {
	g.Default.Port = 162
	g.Default.Target = "10.187.19.39"
	g.Default.Version = g.Version2c
	g.Default.Logger = log.New(os.Stdout, "", 0)

	err := g.Default.Connect()
	if err != nil {
		log.Println("Connect() err: %s", err)
		return
	}
	defer g.Default.Conn.Close()

	p := g.SnmpPDU{
		Name:  "1.3.6.1.6.3.1.1.4.1",
		Value: []byte("Test Alert Message from HAL BOT. Please invoke Callout Group XXXXXXXXX"),
		Type:  g.OctetString,
	}

	trap := g.SnmpTrap{
		Variables: []g.SnmpPDU{p},
	}

	result, err := g.Default.SendTrap(trap)
	if err != nil {
		log.Println("Connect() err: %s", err)
		return
	}


	log.Printf("Error: %d",result.Error)
	log.Printf("Request ID %d",result.RequestID)


}
