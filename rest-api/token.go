package rest_api

import (
	"encoding/json"
	"fmt"

	common_utils "github.com/imtoanle/nginxpm/rest-api/utils"
	"github.com/spf13/viper"
)

type Tokens struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

func CreateNewToken() Tokens {
	tokens := Tokens{}
	params, _ := json.Marshal(map[string]string{
		"identity": viper.GetString("credentials.identity"),
		"secret":   viper.GetString("credentials.secret"),
	})

	_, responseBytes := common_utils.NewRequest(
		viper.GetString("api_endpoint")+"/api/tokens",
		"POST",
		params,
	)

	if err := json.Unmarshal(responseBytes, &tokens); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	return tokens
}
