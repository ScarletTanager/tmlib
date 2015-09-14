package tmlib_test

import (
	. "github.com/ScarletTanager/tmlib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tmlib", func() {
	Context("Using tmlib library", func() {
		It("Writes to the buffer", func() {
			name := "John"
			buf := make([]byte, len(name))
			Expect(CopyToBuffer(buf, name)).To(Equal(1))
			Expect(string(buf)).To(Equal(name))
		})
	})
})
