package schedule

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var PauseCmd = &cli.Command{
	Name:      "pause",
	Usage:     "Pause a specific schedule",
	ArgsUsage: "SCHEDULE_ID",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		schedule, err := schedule.PauseHandler(db, c.Args().Get(0))
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(schedule, c.Bool("json"), nil)
		return nil
	},
}
