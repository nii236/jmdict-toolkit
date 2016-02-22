package fetch

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fetch", func() {
	It("Creates an empty file", func() {
		f, err := createFile("data/test.gz")
		Expect(f).To(BeARegularFile())
		Expect(err).To(BeNil())
	})

})
