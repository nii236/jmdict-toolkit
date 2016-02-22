package fetch

import (
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
		_, err := createFile("data/test.gz")
		Expect("data/test.gz").To(BeARegularFile())
		Expect(err).To(BeNil())
	})

})
