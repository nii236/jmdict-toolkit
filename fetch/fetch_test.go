package fetch_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	fetch "github.com/nii236/jmdict-toolkit/fetch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fetch Dictionary")
}

type mockFetcher struct {
}

func (mf *mockFetcher) Fetch(address string, path string, dest *os.File) error {
	_, err := url.Parse(address)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var _ = Describe("Fetch Dictionary", func() {
	It("Returns error on empty URL", func() {
		mf := &mockFetcher{}
		err := fetch.Dictionary("", "../data/test.gz", mf)
		Expect(err).To(Not(BeNil()))
	})
	It("Successfully retrieves valid URL", func() {
		mf := &mockFetcher{}
		err := fetch.Dictionary("derp", "../data/test.gz", mf)
		Expect(err).To(BeNil())
	})
})
