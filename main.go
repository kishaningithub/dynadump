package main

import (
	"fmt"
	dynadump "github.com/kishaningithub/dynadump/pkg"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var Version = "dev"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "dynadump table_name",
		Short: "Export your dynamoDB data as JSON",
		Long: strings.TrimSpace(`
Export your dynamoDB data as JSON
`),
		Version: Version,
		Example: strings.TrimSpace(`
## Dump data from table_name in JSON format
dynadump table_name > data.json

## Extracting a specific field
dynadump table_name | jq -r '.field'
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			tableName := os.Args[1]
			err := dynadump.DumpTable(tableName)
			if err != nil {
				return err
			}
			return nil
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
