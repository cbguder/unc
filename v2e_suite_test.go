package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
	"os/exec"
)

func TestV2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "v2e Suite")
}

var buildPath string

var _ = SynchronizedBeforeSuite(func() []byte {
	path, err := Build("github.com/cbguder/v2e")
	Expect(err).NotTo(HaveOccurred())
	return []byte(path)
}, func(data []byte) {
	buildPath = string(data)
})

var _ = SynchronizedAfterSuite(func() {}, func (){
	CleanupBuildArtifacts()
})

func Run(args ...string) *Session {
	session, err := Start(exec.Command(buildPath, args...), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
