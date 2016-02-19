package serve_test

import (
	"encoding/json"
	"fmt"

	"github.com/nii236/jmdict-toolkit/serve/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestServe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Serve Suite")
}

var _ = Describe("Word", func() {
	var conn *client.Client
	var jsonPayload string
	BeforeEach(func() {
		conn = client.New()
		jsonPayload = `{"word":"日本語"}`
	})
	It("Gets an OK response", func() {
		b, err := json.Marshal(jsonPayload)
		if err != nil {
			return
		}
		payload := &client.TranslateWordPayload{b}
		resp, err := conn.TranslateWord("http://localhost:8080/translate", payload)
		if err != nil {
			fmt.Println(err)
			return
		}
		Expect(resp.StatusCode).To(Equal(200))
	})
})
