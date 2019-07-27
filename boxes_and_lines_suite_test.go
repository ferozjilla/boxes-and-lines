package miro_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoxesAndLines(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BoxesAndLines Suite")
}
