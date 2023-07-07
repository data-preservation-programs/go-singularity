package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var ListPiecesCmd = &cli.Command{
	Name:      "list-pieces",
	Usage:     "List all pieces for the dataset that are available for deal making",
	ArgsUsage: "<dataset_name>",
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}

		car, err := dataset.ListPiecesHandler(
			db, c.Args().Get(0),
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(car, c.Bool("json"), nil)
		return nil
	},
}
