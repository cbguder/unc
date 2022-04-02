package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"os/exec"
	"testing"
)

func TestUnc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "unc Suite")
}

var buildPath string

var _ = SynchronizedBeforeSuite(func() []byte {
	path, err := Build("github.com/cbguder/unc")
	Expect(err).NotTo(HaveOccurred())
	return []byte(path)
}, func(data []byte) {
	buildPath = string(data)
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	CleanupBuildArtifacts()
})

func Run(args ...string) *Session {
	session, err := Start(exec.Command(buildPath, args...), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
