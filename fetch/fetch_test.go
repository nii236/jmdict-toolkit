package fetch_test

import (
	"fmt"
	"testing"

	fetch "github.com/nii236/jmdict-toolkit/fetch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fetch")
}

type mockFetcher struct {
}

func (mf *mockFetcher) Fetch(address string) {
	fmt.Println("Hello from MOCKFETCHER")
}

var _ = Describe("Fetch", func() {
	It("Creates an empty file", func() {
		mf := &mockFetcher{}
		fetch.Dictionary("", "../data/test.gz", mf)
		// Expect(path).To(BeARegularFile())
		// Expect(err).To(BeNil())
	})

})
