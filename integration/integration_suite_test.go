package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var cliBin string

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		var err error
		cliBin, err = gexec.Build("github.com/mo-work/go-technical-test-for-claudia/cmd", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.Terminate()
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Integration Suite")
}
