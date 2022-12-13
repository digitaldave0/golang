package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func chkssl() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "www.google.com:433", conf)
	if err != nil {
		log.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Printf("Subject: %s\n", cert.Subject)
		fmt.Printf("Issuer Name: %s\n", cert.Issuer)
		fmt.Printf("Expiry: %s \n", cert.NotAfter.Format("2006-January-02"))
		fmt.Printf("Common Name: %s \n", cert.Issuer.CommonName)

	}
}

func myloop() {
	url := []string{"www.ibm.com",
		"www.bbc.co.uk",
		"www.itv.com",
	}

	for i, s := range url {
		fmt.Println(i, s)
	}
}

func main() {

	// function call
	myloop()
}
