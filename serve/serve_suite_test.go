package serve_test

import (
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
		payload := &client.TranslateWordPayload{Word: jsonPayload}
		resp, err := conn.TranslateWord("http://localhost:8080/translate", payload)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.StatusCode)
		Expect(resp.StatusCode).To(Equal(200))
	})
})
