package silenced

import (
	"io"
	"time"

	"github.com/sensu/sensu-go/cli"
	"github.com/sensu/sensu-go/cli/commands/flags"
	"github.com/sensu/sensu-go/cli/commands/helpers"
	"github.com/sensu/sensu-go/cli/elements/globals"
	"github.com/sensu/sensu-go/cli/elements/table"
	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"
)

// ListCommand lists silenceds that were queried from the server
func ListCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "list",
		Short:        "list silenced entries",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			org := cli.Config.Organization()
			flg := cmd.Flags()
			if ok, err := flg.GetBool(flags.AllOrgs); err != nil {
				return err
			} else if ok {
				org = "*"
			}
			// Fetch silenceds from the API
			sub, err := flg.GetString("subscription")
			if err != nil {
				return err
			}
			check, err := flg.GetString("check")
			if err != nil {
				return err
			}
			results, err := cli.Client.ListSilenceds(org, sub, check)
			if err != nil {
				return err
			}

			// Print the results based on the user preferences
			helpers.Print(cmd, cli.Config.Format(), printToTable, results)

			return nil
		},
	}

	flags := cmd.Flags()
	helpers.AddFormatFlag(flags)
	helpers.AddAllOrganization(flags)
	flags.StringP("subscription", "s", "", "name of the silenced subscription")
	flags.StringP("check", "c", "", "name of the silenced check")

	return cmd
}

func printToTable(results interface{}, writer io.Writer) {
	table := table.New([]*table.Column{
		{
			Title:       "ID",
			ColumnStyle: table.PrimaryTextStyle,
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.ID
			},
		},
		{
			Title: "Expire",
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return (time.Duration(silenced.Expire) * time.Second).String()
			},
		},
		{
			Title: "ExpireOnResolve",
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return globals.BooleanStyleP(silenced.ExpireOnResolve)
			},
		},
		{
			Title: "Creator",
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.Creator
			},
		},
		{
			Title:       "Check",
			ColumnStyle: table.PrimaryTextStyle,
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.Check
			},
		},
		{
			Title: "Reason",
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.Reason
			},
		},
		{
			Title:       "Subscription",
			ColumnStyle: table.PrimaryTextStyle,
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.Subscription
			},
		},
		{
			Title:       "Organization",
			ColumnStyle: table.PrimaryTextStyle,
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.Organization
			},
		},
		{
			Title:       "Environment",
			ColumnStyle: table.PrimaryTextStyle,
			CellTransformer: func(data interface{}) string {
				silenced, _ := data.(types.Silenced)
				return silenced.Environment
			},
		},
	})

	table.Render(writer, results)
}
