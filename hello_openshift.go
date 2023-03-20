package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

// Snakeoil certificate
var crt = `-----BEGIN CERTIFICATE-----
MIIDZTCCAk2gAwIBAgIUNFRpTFbOLzxlrjOoyko+ppZFGpgwDQYJKoZIhvcNAQEL
BQAwQjELMAkGA1UEBhMCWFgxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UE
CgwTRGVmYXVsdCBDb21wYW55IEx0ZDAeFw0yMTEyMDIwODIyNDdaFw0zMTExMzAw
ODIyNDdaMEIxCzAJBgNVBAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAa
BgNVBAoME0RlZmF1bHQgQ29tcGFueSBMdGQwggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQDL+/tl7abhJ+KCOT2th9aZMuD60YIkBHMnhbmTrcdkNzqS7GuN
2L2/tgAk83/ROC++PCrYv6INO38JhGOZPnYofiZaIVwPpL6MgCwvY7w5ipRAhUIw
UlFq1BwbYyUMylavy8JBYf/RnmNNjdyy4aB4gt4JrRhI0qsLO1JDN8m1684p5eNA
Pd4xZ7mqwbG5iT1uULbhOpJEcmKRomq7q/Jm5T8KdSRxl2MDVe8l32NeEAKGJSLI
23DYjrPZRswqKJX/FUzXWJScsQGsULqcweiDE1NpJfxJ4rz77VSuz2yo57YgYvbK
izSIpWjkYHvu5kR5FSZOXomGzhr/Ha5CT8ZlAgMBAAGjUzBRMB0GA1UdDgQWBBRN
nQ0wCqglltOUuYLUeMwDEkNkITAfBgNVHSMEGDAWgBRNnQ0wCqglltOUuYLUeMwD
EkNkITAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCcZZUeDs4I
QA+iPz81rbRe+6//KeT04BCFY+1Ie6uRQL5ULgV7amMFWUpAvoMkqAqJQo6q3ZLv
2hTJMx/rYqEvOcT4zrRfTTiN0fqK6mzfTK1O48BRiXpXMeWIZDHQr8hY2SLahwcP
bnYlffms2D4gks8N4niDV+DrYXBmoDIErwrCQChMwRfHHZK25pKcw6Z21P8nDJv8
RJSZ3lvJ5WzNZ098KN3iHNeHXrpHURWdYvSR3fxhVkZ4Ha0nvGzAMIRPDkBOHDvw
6U8T8+4rwF5Up2xVxs3FpOYriA3r2V459xTuQVmqRqrWmlzhaxJWtZpA7DokpVx+
U1gsRBCE3wv7
-----END CERTIFICATE-----
`

// Snakeoil key
var key = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAy/v7Ze2m4Sfigjk9rYfWmTLg+tGCJARzJ4W5k63HZDc6kuxr
jdi9v7YAJPN/0Tgvvjwq2L+iDTt/CYRjmT52KH4mWiFcD6S+jIAsL2O8OYqUQIVC
MFJRatQcG2MlDMpWr8vCQWH/0Z5jTY3csuGgeILeCa0YSNKrCztSQzfJtevOKeXj
QD3eMWe5qsGxuYk9blC24TqSRHJikaJqu6vyZuU/CnUkcZdjA1XvJd9jXhAChiUi
yNtw2I6z2UbMKiiV/xVM11iUnLEBrFC6nMHogxNTaSX8SeK8++1Urs9sqOe2IGL2
yos0iKVo5GB77uZEeRUmTl6Jhs4a/x2uQk/GZQIDAQABAoIBAGUcGQf0HcT7RS5x
ex4Z+AhmDNimotCBmCbeBReriusk6RbMs59S8PMnHrkyLYgiRqAQKNjZXFUcyaKJ
Cel66Yy2wwHoCT8D1SPFoKE42aLYCxZUN3PGSe8fBnOY2FOXtBJdeIN6NRjNXsGh
cOUGK8mwbKj1MNVf/0KI/ASvkX9naQgj9/pH5QrFoQfJYHH8G8lXQyW8caho64wJ
6TJdbm5De1Jgy/QdhBOz4XjsZ3V/Ds8d1t/uLDJTU4jMby6Craiqcl+QIgaCkbdz
w1wgxEH6Z/dxNy6w5H9BH/R1/px5MoO/IPWOFytRWXieKWNrtc1YlkWdo27VV1au
hDNImL0CgYEA8v2NswQPyQeFReRxuA6ptF3/WirLVq+EXlO2eMx6vGAbddS1QzlN
347thP/e6tfA6Bkrp8L/Cc9Jd9/sV4L21FU+rPZvCqVq/9Shr9oewJyeM+ov0njR
QqCsJc3ThH1X7iRHKhInWoXqKyQjqoUi4aPWyucXwsvX1fs0CfvsMY8CgYEA1ufO
p+xNvICGqz4eCs4KLc2l3I8ZFFBnWpjuPowGtHTtePCOqb3NxXv1MQwljb9uXO2o
sm7UYw977Szic153PPG3VJx7OucvPxEAz+zGEE7OgTlV3CWIEi6jS2DSptAgAOAA
XqmxQbMqSbCzMyKn+cjgPBGXgyKRyoUJasOi5ssCgYBEBHequaNViXZj5xtyAyC7
7WfyLHJ41G7AHLzCObLNkjV9mUoYBC1pO8/+38TdhgFotssCjdHoRA7zsEmvAWjo
bOg7cEwK9dzqufF8kRj0n6KlM5OpXcpt1R37Aw+HUbLQZXVKMIS1kTDIXLhjHhty
f/M7Hs8G5xqGumeJt+wYvwKBgFWg3Z3ZMQw35fDbelrxx+qCM2FTfzmx48ycfOld
H9rNcEWtDBskLpZOkt7tKRV2vkG2zG30bRnfdJCHPt+bN0WIRnUnOI66yP+HBdzT
SgP7cprYvpZOOg6MmLITLTwcV3QhzOPrF17HRcVA69YnK+kCGh61H7q3joG0SpFI
zGLLAoGAfRljuMVa6Rlfb/BEbRfp6ngAiqyGDTXQIwVG7CdfE4ozg+RGULa30Rh7
haNtJ3m68+AfHa4ZCfl1zXJHHALpXgZvFnpBC4igJ7a0W3CUshp8JK8G/LyEmn59
Nb5Vh78f5HFpEOSYRN9u7jd+XtbuT7nxvc8Ic9PNxKHuugYQr5E=
-----END RSA PRIVATE KEY-----
`

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")

	if len(response) == 0 {
		response = "<h2>The passphrase is \"Open Sesame\". Don't tell anyone!</h2>"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, response)
	fmt.Println("Servicing request.")
}

func listenAndServe(port string) {
	// Generate a key pair from your pem-encoded cert and key ([]byte).
	cert, err := tls.X509KeyPair([]byte(crt), []byte(key))
	if err != nil {
		fmt.Println("Error loading key pair: %v", err)
		return
	}

	// Construct a tls.config
	tlsConfig := &tls.Config{
		// Other options
		Certificates: []tls.Certificate{cert},
	}

	// Build a server:
	server := http.Server{
		// Other options
		Addr: ":" + port,
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", helloHandler)

	// Finally: serve.
	err = server.ListenAndServeTLS("", "")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
