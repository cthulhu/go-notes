package scanner_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestScanner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scanner Suite")
}
