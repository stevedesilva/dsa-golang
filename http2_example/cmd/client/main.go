package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{Transport: transport2()}
	res, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	fmt.Printf("Code: %d\n", res.StatusCode)
	fmt.Printf("Body: %s\n", body)
}

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	crt, err := os.ReadFile("http2_example/cert/public.crt")
	if err != nil {
		log.Fatal(err)
	}

	rootCA := x509.NewCertPool()
	rootCA.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCA,
		ServerName:         "localhost",
		InsecureSkipVerify: false,
	}
}

//func tlsConfig() *tls.Config {
//	crt, err := os.ReadFile("./cert/public.crt")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rootCA := x509.NewCertPool()
//	rootCA.AppendCertsFromPEM(crt)
//
//	return &tls.Config{
//		RootCAs:            rootCA,
//		ServerName:         "localhost",
//		InsecureSkipVerify: false,
//	}
//}
