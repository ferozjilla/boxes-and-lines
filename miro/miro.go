package miro

import (
	"os"

	bosh "github.com/pivotal-cf/on-demand-services-sdk/bosh"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var miro_api_access_token string = os.Getenv("MIRO_API_ACCESS_TOKEN")

type Drawer struct{}

func (d *Drawer) Draw(manifest bosh.BoshManifest) (string, error) {
	id, err := createBoard(manifest.Name)
	if err != nil {
		return "", err
	}
	return id, nil
}

//curl --request POST \
//  --url https://api.miro.com/v1/boards \
//  --header 'authorization: Bearer <token>' \
//  --header 'content-type: application/json' \
//  --data '{"name":"Feroz Board","description":"This is a test board","sharingPolicy":{"access":"private","accountAccess":"private"}}'
func createBoard(name string) (string, error) {
	client := &http.Client{}
	url := "https://api.miro.com/v1/boards"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", miro_api_access_token))
	req.Header.Set("content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var createBoardResponse map[string]interface{}
	err = json.Unmarshal(bodyBytes, &createBoardResponse)
	if err != nil {
		return "", err
	}

	return createBoardResponse["id"].(string), nil
}
