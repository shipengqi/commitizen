package e2e_test

import (
	"flag"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/shipengqi/golib/fsutil"

	. "github.com/shipengqi/commitizen/test/e2e"
)

var (
	se  *gexec.Session
	err error
)

func init() {
	flag.StringVar(&CliOpts.Cli, "cli", "", "path to the commitizen command to use.")
	flag.IntVar(&CliOpts.NoTTY, "no-tty", 0, "make sure that the TTY (terminal) is never used for any output.")
}

var _ = Describe("Sorted Tests", func() {
	Describe("Version Command", VersionTest)
	Describe("CZ Command", CZTest)
})

func TestE2e(t *testing.T) {
	// Skip running E2E tests when running only "short" tests because:
	// 1. E2E tests are long-running tests involving generation of skeletons.
	if testing.Short() {
		t.Skip("Skipping E2E tests")
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func() {
	flag.Parse()

	if CliOpts.Cli == "" {
		CliOpts.Cli = "./commitizen"
	}
})

// ===================================================
// Helpers

func RunCLITestAndWait(args ...string) (*gexec.Session, error) {
	cmd := exec.Command(CliOpts.Cli, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	return session.Wait(), err
}

func RunCLITest(args ...string) (*gexec.Session, error) {
	if CliOpts.NoTTY == 1 {
		args = append(args, "--no-tty")
	}
	cmd := exec.Command(CliOpts.Cli, args...)
	return gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
}

func NoError(err error) {
	Expect(err).To(BeNil())
}

func ExitCode(session *gexec.Session, expected int) {
	Ω(session.ExitCode()).Should(Equal(expected))
}

func ShouldContains(session *gexec.Session, expected string) {
	Ω(session.Out.Contents()).Should(ContainSubstring(expected))
}

func ShouldNotContains(session *gexec.Session, expected string) {
	Ω(session.Out.Contents()).ShouldNot(ContainSubstring(expected))
}

func ShouldExists(fpath string) {
	exists := fsutil.IsExists(fpath)
	Expect(exists).To(Equal(true))
}

func ShouldNotExists(fpath string) {
	exists := fsutil.IsExists(fpath)
	Expect(exists).To(Equal(false))
}
