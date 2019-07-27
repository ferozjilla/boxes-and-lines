package miro

import (
  bosh "github.com/pivotal-cf/on-demand-services-sdk/bosh"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

var MIRO_API_ACCESS_TOKEN string = os.Getenv("MIRO_API_ACCESS_TOKEN")

type Drawer struct {}

//curl --request POST \
//  --url https://api.miro.com/v1/boards \
//  --header 'authorization: Bearer <token>' \
//  --header 'content-type: application/json' \
//  --data '{"name":"Feroz Board","description":"This is a test board","sharingPolicy":{"access":"private","accountAccess":"private"}}'
func (d *Drawer) Draw(manifest bosh.BoshManifest) (string, error) {
	client := &http.Client{}
	url := "https://api.miro.com/v1/boards"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", MIRO_API_ACCESS_TOKEN))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var createBoardResponse map[string]string
	json.Unmarshal(body, &createBoardResponse)

	id := createBoardResponse["id"]
	return id, nil
}
