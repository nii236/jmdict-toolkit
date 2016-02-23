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
	RunSpecs(t, "Dictionary")
}

type mockFetcher struct{}
type mockFileCreator struct{}

func (mf *mockFetcher) Fetch(address string, path string, dest *os.File) error {
	_, err := url.Parse(address)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (mfc *mockFileCreator) CreateFile(path string) (*os.File, error) {
	out, err := os.Create(path)
	if err != nil {
		return out, err
	}

	return out, nil
}

var _ = Describe("Dictionary", func() {
	mf := &mockFetcher{}
	mfc := &mockFileCreator{}

	It("Returns error on empty URL", func() {
		err := fetch.Dictionary("", "../data/test", mf, mfc)
		Expect(err).To(Not(BeNil()))
	})
	It("Successfully retrieves valid URL", func() {
		err := fetch.Dictionary("derp", "../data/test", mf, mfc)
		Expect(err).To(BeNil())
	})
	It("Returns error when a folder does not exist", func() {
		err := fetch.Dictionary("derp", "fakefolder/test", mf, mfc)
		Expect(err).To(Not(BeNil()))
	})
})

var _ = Describe("FileCreator", func() {
	It("Returns error when a folder does not exist", func() {
		fc := fetch.FileCreator{}
		dest, err := fc.CreateFile("fakefolder/temp")
		defer dest.Close()
		Expect(err).To(Not(BeNil()))
	})
})

var _ = Describe("Fetch", func() {
	It("Returns error on empty URL", func() {
		fc := fetch.FileCreator{}
		dest, err := fc.CreateFile("../data/temp")
		defer dest.Close()
		Expect(err).To(BeNil())
		fetcher := fetch.Fetcher{}
		fetcher.Fetch("", "", dest)

	})
})
