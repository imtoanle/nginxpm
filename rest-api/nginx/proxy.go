package rest_api_nginx

import (
	"encoding/json"
	"fmt"
	"strings"

	common_utils "github.com/imtoanle/nginxpm/rest-api/utils"
	"github.com/spf13/viper"
)

var defaultParams = map[string]interface{}{
	"forward_scheme":          "http",
	"access_list_id":          "0",
	"ssl_forced":              true,
	"caching_enabled":         true,
	"block_exploits":          true,
	"allow_websocket_upgrade": true,
	"http2_support":           true,
	"enabled":                 true,
	"hsts_enabled":            true,
	"hsts_subdomains":         true,
}

func CreateProxyHost(attrs map[string]string) {
	responseJson := map[string]interface{}{}

	response, responseBytes := common_utils.NewRequest(
		viper.GetString("api_endpoint")+"/api/nginx/proxy-hosts",
		"POST",
		prepareParams(attrs),
	)

	if err := json.Unmarshal(responseBytes, &responseJson); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	if response.StatusCode == 201 {
		fmt.Println("Success created proxy host.")
		fmt.Println("Forward Host: " + responseJson["forward_host"].(string))

		fmt.Print("Forward Port: ")
		fmt.Println(responseJson["forward_port"].(float64))

		fmt.Print("https://")
		fmt.Println(responseJson["domain_names"].([]interface{})[0])
	} else {
		fmt.Println("Error: " + responseJson["error"].(map[string]interface{})["message"].(string))
	}
}

func detectCertificate(domain_name string) Certificate {
	var cert Certificate

	s := strings.Split(domain_name, ".")
	findingCertificate := "*." + strings.Join(s[1:], ".")

	for _, v := range GetCertificates() {
		if v.NiceName == findingCertificate {
			cert = v
		}
	}

	return cert
}

func prepareParams(attrs map[string]string) []byte {
	params := defaultParams
	params["domain_names"] = []string{attrs["domain_name"]}
	params["forward_host"] = attrs["forward_host"]
	params["forward_port"] = attrs["forward_port"]

	if certificate := detectCertificate(attrs["domain_name"]); certificate.Id != 0 {
		params["certificate_id"] = certificate.Id
	}

	p, _ := json.Marshal(params)
	return p
}
