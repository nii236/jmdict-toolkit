package fetch

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fetch")
}

var _ = Describe("Fetch", func() {
	It("Creates an empty file", func() {
		path := "../data/test.gz"
		_, err := createFile(path)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		Expect(path).To(BeARegularFile())
		Expect(err).To(BeNil())
	})

})
