package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
)

func VersionTest() {
	Context("Check Version Message", func() {
		It("should output the all version messages", func() {
			se, err = RunCLITest("version")
			NoError(err)
			ShouldContains(se, "Version:")
			ShouldContains(se, "Commit:")
			ShouldContains(se, "GitTreeState:")
			ShouldContains(se, "BuildTime:")
			ShouldContains(se, "GoVersion:")
			ShouldContains(se, "Compiler:")
			ShouldContains(se, "Platform:")
		})
	})
}
