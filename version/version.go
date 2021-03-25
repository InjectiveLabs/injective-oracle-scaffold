package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	AppVersion = ""
	GitCommit  = ""
	BuildDate  = ""

	GoVersion = ""
	GoArch    = ""
)

func init() {
	if len(AppVersion) == 0 {
		AppVersion = "dev"
	}

	GoVersion = runtime.Version()
	GoArch = runtime.GOARCH
}

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the application binary version information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetOut(cmd.OutOrStdout())
			output := fmt.Sprintf(
				"Version %s (%s)\nCompiled at %s using Go %s (%s)",
				AppVersion,
				GitCommit,
				BuildDate,
				GoVersion,
				GoArch,
			)
			cmd.Println(output)
			return nil
		},
	}
	return cmd
}
