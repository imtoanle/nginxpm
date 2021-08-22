package rest_api_nginx

import (
	"encoding/json"
	"fmt"

	common_utils "github.com/imtoanle/nginxpm/rest-api/utils"
	"github.com/spf13/viper"
)

type Certificate struct {
	Id       int    `json:"id"`
	NiceName string `json:"nice_name"`
}

// type Certificates struct []Certificate{}

func GetCertificates() []Certificate {
	certificates := []Certificate{}
	_, responseBytes := common_utils.NewRequest(viper.GetString("api_endpoint")+"/api/nginx/certificates", "GET", nil)

	if err := json.Unmarshal(responseBytes, &certificates); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	return certificates
}
