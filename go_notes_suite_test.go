package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoNotes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoNotes Suite")
}
