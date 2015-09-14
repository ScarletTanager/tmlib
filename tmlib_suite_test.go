package tmlib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTmlib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmlib Suite")
}
