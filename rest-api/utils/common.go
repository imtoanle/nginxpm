package common_utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/spf13/viper"
)

func NewRequest(url string, method string, params []byte) (*http.Response, []byte) {
	request, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(params),
	)

	if err != nil {
		log.Printf("Could not request a new token. %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Nginx Proxy Manager CLI (https://github.com/imtoanle/nginxpm)")
	request.Header.Set("Authorization", "Bearer "+viper.GetString("sessions.token"))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return response, responseBytes
}

func GetLocalIP() string {
	firstIp := ""
	addrs, _ := net.InterfaceAddrs()

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				firstIp = ipnet.IP.String()
				break
			}
		}
	}

	return firstIp
}
