package miro_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"

	bosh "github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"gopkg.in/yaml.v2"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ferozjilla/boxes-and-lines/miro"
)

var MIRO_API_ACCESS_TOKEN string = os.Getenv("MIRO_API_ACCESS_TOKEN")

var _ = Describe("Miro Drawer", func() {
	var id string

	It("Draws a simple manifest", func() {
		// setup
		miroDrawer := miro.Drawer{}
		var boshManifest bosh.BoshManifest

		manifestBytes, err := ioutil.ReadFile("../assets/simple.yml")
		Expect(err).NotTo(HaveOccurred())
		err = yaml.Unmarshal(manifestBytes, &boshManifest)
		Expect(err).NotTo(HaveOccurred())

		id, err = miroDrawer.Draw(boshManifest)

		Expect(err).NotTo(HaveOccurred())
		By("Creating a new Miro board")
		Expect(miroBoardExists(id)).To(BeTrue())
		By("Drawing boxes for the instance groups")
		By("Drawing as many boxes as the number of instances for an instance group")
		By("Labeling each box with the name of the instance group")

		// tear down
		//err = deleteMiroBoard(id)
		//Expect(err).NotTo(HaveOccurred())
	})

	It("Draws a complex manifest", func() {
		// setup
		miroDrawer := miro.Drawer{}
		var boshManifest bosh.BoshManifest

		manifestBytes, err := ioutil.ReadFile("../assets/complex.yml")
		Expect(err).NotTo(HaveOccurred())
		err = yaml.Unmarshal(manifestBytes, &boshManifest)
		Expect(err).NotTo(HaveOccurred())

		id, err = miroDrawer.Draw(boshManifest)

		Expect(err).NotTo(HaveOccurred())
		By("Creating a new Miro board")
		Expect(miroBoardExists(id)).To(BeTrue())
		By("Drawing boxes for the instance groups")
		By("Drawing as many boxes as the number of instances for an instance group")
		By("Labeling each box with the name of the instance group")

		// tear down
		//err = deleteMiroBoard(id)
		//Expect(err).NotTo(HaveOccurred())
	})
})

// curl --request GET \
//  --url https://api.miro.com/v1/boards/id \
//  --header 'authorization: Bearer <token>'
func miroBoardExists(id string) (bool, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.miro.com/v1/boards/%v", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", MIRO_API_ACCESS_TOKEN))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("http request query miro board failed with exit code: %+v", res.StatusCode))
	}
	return true, nil
}

func deleteMiroBoard(id string) error {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.miro.com/v1/boards/%v", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", MIRO_API_ACCESS_TOKEN))
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("http request to delete miro board failed with exit code: %+v", res.StatusCode))
	}
	return nil
}
