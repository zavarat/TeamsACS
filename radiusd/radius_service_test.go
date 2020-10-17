package radiusd

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"

	"github.com/ca17/teamsacs/common"
)

func TestAuth(t *testing.T) {
	packet := radius.New(radius.CodeAccessRequest, []byte(`123456`))
	common.Must(rfc2865.UserName_SetString(packet, "tim"))
	common.Must(rfc2865.UserPassword_SetString(packet, "12345"))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	fmt.Println(packet)
	cli := radius.Client{
		Net:                "udp",
		Retry:              0,
		MaxPacketErrors:    10,
		InsecureSkipVerify: true,
	}
	response, err := cli.Exchange(ctx, packet, "127.0.0.1:1812")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Code:", response.Code)
}

func TestAcct(t *testing.T) {
	packet := radius.New(radius.CodeAccountingRequest, []byte(`12345`))
	common.Must(rfc2865.UserName_SetString(packet, "tim"))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	fmt.Println(packet)
	cli := radius.Client{
		Net:                "udp",
		Retry:              0,
		MaxPacketErrors:    10,
		InsecureSkipVerify: true,
	}
	response, err := cli.Exchange(ctx, packet, "127.0.0.1:1813")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Code:", response.Code)
}
