package main

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"crypto/tls"
)

func main() {

	cert, err := tls.LoadX509KeyPair("apns_cer.pem",
		"apns_key_noenc.pem")

	//cert, err := certificate.FromPemFile("/Users/qiang/Downloads/app/pemdev/apns_cer.pem", "")

	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "7eaf9bb32b738c252ea6be6f2e34ea2988bb3f533a13f6eaa91123f838f0cc81"
	notification.Topic = "com.XX.XX"
	//notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below
	notification.Payload = []byte(`{
    "aps" : {
            "alert" : "Hello"
          } 
}`) // See Payload section below

	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("push Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
