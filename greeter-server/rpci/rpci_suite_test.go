package rpci_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRpci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rpci Suite")
}
