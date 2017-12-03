package iputility_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIputility(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iputility Suite")
}
