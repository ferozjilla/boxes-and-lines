package miro_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"gopkg.in/yaml.v2"
	bosh "github.com/pivotal-cf/on-demand-services-sdk/bosh"

	"github.com/ferozjilla/boxes-and-lines"
	"net/http"
	"fmt"
	"log"
	"os"
)

var MIRO_API_ACCESS_TOKEN string = os.Getenv("MIRO_API_ACCESS_TOKEN")

var _ = Describe("Miro Drawer", func() {
	var id string

	BeforeEach(func() {
		miroDrawer := miro.Drawer{}
		var boshManifest bosh.BoshManifest

		manifestBytes, err := ioutil.ReadFile("assets/simple.yml")
		Expect(err).NotTo(HaveOccurred())
		err = yaml.Unmarshal(manifestBytes, &boshManifest)
		Expect(err).NotTo(HaveOccurred())

		id, err = miroDrawer.Draw(boshManifest)
		Expect(err).NotTo(HaveOccurred())
	})

	It("creates a new Miro board", func() {
		Expect(miroBoardExists(id)).To(BeTrue())
	})

	It("does not overwrite an existing user board", func() {

	})

	Context("Simple drawing", func() {
		Context("Box", func() {
			It("Can draw a box", func() {

			})
		})
	})
})

// curl --request GET \
//  --url https://api.miro.com/v1/boards/id \
//  --header 'authorization: Bearer <token>'
func miroBoardExists(id string) (bool, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.miro.com/v1/boards/%v", id)
	fmt.Printf("url :%v\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", MIRO_API_ACCESS_TOKEN))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	boardExists := res.StatusCode == 200
	return boardExists, nil
}
