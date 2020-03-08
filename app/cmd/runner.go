package cmd

import (
	"github.com/jenkins-zh/jenkins-cli/app/i18n"
	"github.com/spf13/cobra"
)

// RunnerOption is the wrapper of jenkinsfile runner cli
type RunnerOption struct {
	BatchOption
	CommonOption

	Safe bool
}

var runnerOption RunnerOption

func init() {
	rootCmd.AddCommand(runnerCmd)
	runnerOption.SetFlag(runnerCmd)
	runnerCmd.Flags().BoolVarP(&runnerOption.Safe, "safe", "s", true,
		i18n.T("Puts Jenkins into the quiet mode, wait for existing builds to be completed, and then restart Jenkins"))
	runnerOption.BatchOption.Stdio = GetSystemStdio()
	runnerOption.CommonOption.Stdio = GetSystemStdio()
}

var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: i18n.T("The wrapper of jenkinsfile runner"),
	Long: i18n.T(`The wrapper of jenkinsfile runner
Get more about jenkinsfile runner from https://github.com/jenkinsci/jenkinsfile-runner`),
	RunE: func(cmd *cobra.Command, _ []string) (err error) {
		return
	},
}
