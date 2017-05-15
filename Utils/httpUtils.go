package Utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"io"
)

type AccessTokenFormat struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

func generateRequest(url string, method string, body io.Reader, isAccessToken bool) []byte {
	req,err := http.NewRequest(method,url,body)
	if err != nil { log.Fatal(err) }

	if !isAccessToken {
		req.Header.Set("Authorization","Bearer "+getAccessToken())
	}

	client := new(http.Client)
	res,err := client.Do(req)

	if err != nil { log.Fatal(err) }

	responseData, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return responseData
}

func getAccessToken() string {
	values := url.Values{}
	values.Add("client_id", os.Getenv("MS_TRANSLATE_ID"))
	values.Add("client_secret", os.Getenv("MS_TRANSLATE_SECRET"))
	values.Add("scope", "http://api.microsofttranslator.com")
	values.Add("grant_type", "client_credentials")

	data := generateRequest(
		"https://datamarket.accesscontrol.windows.net/v2/OAuth2-13",
		"POST",
		strings.NewReader(values.Encode()),
		true,
	)

	var accessTokenFormat AccessTokenFormat
	if err := json.Unmarshal(data, &accessTokenFormat); err != nil {
		log.Fatal(err)
	}

	return accessTokenFormat.AccessToken

}

func Translate(from,to,text string) string {

	opt := "from="+from+"&to="+to+"&text="+text+"&oncomplete="

	data := generateRequest(
		"http://api.microsofttranslator.com/V2/Ajax.svc/Translate?" + opt,
		"GET",
		nil,
		false,
	)

	translatedText := strings.Replace(string(data),"\"","",-1)

	return translatedText
}
