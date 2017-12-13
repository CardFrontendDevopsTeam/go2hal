package service

import (
	"os"
	"gopkg.in/kyokomi/emoji.v1"
	g "github.com/zamedic/gosnmp"
	"log"
	"net"
	"fmt"
	"bytes"
	"errors"
)

func init(){
	startSnmpServer()
}

func startSnmpServer() {
	tl := g.NewTrapListener()
	tl.OnNewTrap = handleTrap
	tl.Params = g.Default
	tl.Params.Logger = log.New(os.Stdout,"",0)
	err := tl.Listen("0.0.0.0:9162")
	if err != nil {
		log.Panicf("error in listen: %s", err)
	}

}

func handleTrap(packet *g.SnmpPacket, addr *net.UDPAddr) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("got trapdata from %s\n", addr.IP))
	b.WriteString("\n")
	for _, v := range packet.Variables {
		switch v.Type {
		case g.OctetString:
			c := v.Value.([]byte)
			b.WriteString(fmt.Sprintf("OID: %s, string: %x\n", v.Name, c))
			b.WriteString("\n")
		default:
			b.WriteString(fmt.Sprintf("trap: %+v\n", v))
			b.WriteString("\n")

		}
	}
	SendError(errors.New(b.String()))
}

func sendSNMPMessage() {
	if snmpServier() == "" {
		return
	}
	g.Default.Port = snmpPort()
	g.Default.Target = snmpServier()
	g.Default.Version = g.Version2c
	g.Default.Logger = log.New(os.Stdout, "", 0)

	log.Printf("SNMP Server: %s Port: %d", g.Default.Target, g.Default.Port)

	err := g.Default.Connect()
	if err != nil {
		log.Printf("Connect() err: %v", err)
		return
	}
	defer g.Default.Conn.Close()

	p := g.SnmpPDU{
		Name:  ".1.3.6.1.4.1.789.1.2.2.4.0",
		Value: "Test Alert Message from HAL BOT. Please invoke Callout Group XXXXXXXXX",
		Type:  g.OctetString,
	}

	trap := g.SnmpTrap{
		Variables: []g.SnmpPDU{p},
	}

	result, err := g.Default.SendTrap(trap)
	if err != nil {
		log.Printf("Connect() err: %v", err)
		return
	}

	log.Printf("Error: %d", result.Error)
	log.Printf("Request ID %d", result.RequestID)
	SendAlert(emoji.Sprint(":telephone_receiver: Invoked callout"))

}