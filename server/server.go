package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	c "grapi/config"
	m "grapi/middlewares"
	r "grapi/router"
)

var httpPort string
var httpsPort string
var address string
var certsDir string

// StartServer :
func StartServer() {
	checkConfig()
	httpPort = ":" + c.Cfg.ServerPort
	address = c.Cfg.ServerAddress
	certsDir = c.Cfg.CertsDir
	cert := certsDir + "/cert.pem"
	key := certsDir + "/key.pem"
	checkPorts()

	if c.Cfg.HTTPS != 0 {
		_, err := os.Stat(cert)
		_, err2 := os.Stat(key)
		if err != nil || err2 != nil {
			log.Fatal("Cert files not found")
		}
		if c.Cfg.HTTPSOnly != 0 {
			log.Printf("Http server at %v%v redirecting to %v%v", address, httpPort, address, httpsPort)
			go loggedRedirectServer()
		} else {
			log.Printf("Http server started at %v%v", address, httpPort)
			go func() { log.Fatal(http.ListenAndServe(address+httpPort, r.Router)) }()
		}
		log.Printf("Https server started at %v%v", address, httpsPort)
		log.Fatal(http.ListenAndServeTLS(httpsPort, cert, key, r.Router))
	} else {
		log.Printf("Http server started at %v%v", address, httpPort)
		log.Fatal(http.ListenAndServe(address+httpPort, r.Router))
	}
}

func loggedRedirectServer() {
	log.Fatal(http.ListenAndServe(address+httpPort, m.Logger(http.HandlerFunc(redirectToHTTPS), "Redirect")))
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+address+httpsPort+r.RequestURI, http.StatusMovedPermanently)
}

func checkConfig() {
	if c.Cfg.ServerAddress == "" {
		log.Fatal("Missing server address in config file")
	}
	if c.Cfg.ServerPort == "" {
		c.Cfg.ServerPort = "8080"
	}
	if c.Cfg.CertsDir == "" {
		c.Cfg.CertsDir = "."
	}
}

func checkPorts() {
	port, err := strconv.Atoi(c.Cfg.ServerPort)

	if err != nil {
		httpPort = ":8080"
		port = 8080
	}
	httpsPort = strconv.Itoa(port + 1)
	httpsPort = ":" + httpsPort
}
