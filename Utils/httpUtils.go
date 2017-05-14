package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type AccessTokenFormat struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

func GetAccessToken() string {
	values := url.Values{}
	values.Add("client_id", os.Getenv("MS_TRANSLATE_ID"))
	values.Add("client_secret", os.Getenv("MS_TRANSLATE_SECRET"))
	values.Add("scope", "http://api.microsofttranslator.com")
	values.Add("grant_type", "client_credentials")

	url := "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13"
	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var accessTokenFormat AccessTokenFormat
	if err := json.Unmarshal(data, &accessTokenFormat); err != nil {
		log.Fatal(err)
	}

	return accessTokenFormat.AccessToken

}

func Translate() {

}
