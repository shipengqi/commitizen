package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

func CZTest() {
	Context("Check Commitizen", func() {
		It("should not need to select a template", func() {
			se, err = RunCLITest()
			NoError(err)
			Eventually(se.Out).Should(gbytes.Say("Select the type of"))
			Eventually(se.Out).Should(gbytes.Say("A new feature"))
			Eventually(se.Out).ShouldNot(gbytes.Say("Select a template to use for this commit:"))
			se.Terminate()
		})
	})
}
