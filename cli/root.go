package cli

import (
	"github.com/MakeNowJust/heredoc"
	handlerv1beta1 "github.com/odpf/guardian/api/handler/v1beta1"
	"github.com/odpf/salt/cmdx"
	"github.com/spf13/cobra"
)

func New(cfg *Config) *cobra.Command {
	cliConfig = cfg
	var cmd = &cobra.Command{
		Use:   "guardian <command> <subcommand> [flags]",
		Short: "Universal data access control",
		Long: heredoc.Doc(`
			Universal access control to cloud apps and infrastructure.`),
		SilenceUsage:  true,
		SilenceErrors: true,
		Example: heredoc.Doc(`
			$ guardian appeal create
			$ guardian policy list
			$ guardian provider list
			$ guardian resource list
			$ guardian policy create --file policy.yaml
		`),
		Annotations: map[string]string{
			"group:core": "true",
			"help:learn": heredoc.Doc(`
				Use 'guardian <command> <subcommand> --help' for more information about a command.
				Read the manual at https://odpf.github.io/guardian/
			`),
			"help:feedback": heredoc.Doc(`
				Open an issue here https://github.com/odpf/guardian/issues
			`),
			"help:environment": heredoc.Doc(`
				See 'guardian help environment' for the list of supported environment variables.
			`),
		},
	}

	protoAdapter := handlerv1beta1.NewAdapter()
	cmd.AddCommand(ResourceCmd(protoAdapter))
	cmd.AddCommand(ProviderCmd(protoAdapter))
	cmd.AddCommand(PolicyCmd(protoAdapter))
	cmd.AddCommand(appealsCommand())
	cmd.AddCommand(ServerCommand())
	cmd.AddCommand(JobCmd())
	cmd.AddCommand(configCommand())
	cmd.AddCommand(VersionCmd())

	// Help topics
	cmdx.SetHelp(cmd)
	cmd.AddCommand(cmdx.SetCompletionCmd("guardian"))
	cmd.AddCommand(cmdx.SetHelpTopic("environment", envHelp))
	cmd.AddCommand(cmdx.SetRefCmd(cmd))

	return cmd
}
