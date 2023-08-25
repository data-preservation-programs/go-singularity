package ez

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var PrepCmd = &cli.Command{
	Name:      "ez-prep",
	Category:  "Easy Commands",
	ArgsUsage: "<path>",
	Usage:     "Prepare a dataset from a local path",
	Description: "This commands can be used to prepare a dataset from a local path with minimum configurable parameters.\n" +
		"For more advanced usage, please use the subcommands under `storage` and `data-prep`.\n" +
		"You can also use this command for benchmarking with in-memory database and inline preparation, i.e.\n" +
		"  mkdir dataset\n" +
		"  truncate -s 1024G dataset/1T.bin\n" +
		"  singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "max-size",
			Aliases: []string{"M"},
			Usage:   "Maximum size of the CAR files to be created",
			Value:   "31.5GiB",
		},
		&cli.StringFlag{
			Name:    "output-dir",
			Aliases: []string{"o"},
			Usage:   "Output directory for CAR files. To use inline preparation, use an empty string",
			Value:   "./cars",
		},
		&cli.IntFlag{
			Name:    "concurrency",
			Aliases: []string{"j"},
			Usage:   "Concurrency for packing",
			Value:   1,
		},
		&cli.StringFlag{
			Name:        "database-file",
			Usage:       "The database file to store the metadata. To use in memory database, use an empty string.",
			DefaultText: "./ezprep-<name>.db",
		},
	},
	Action: func(c *cli.Context) error {
		t := time.Now().Unix()
		path := c.Args().Get(0)
		if path == "" {
			return errors.New("path is required")
		}
		databaseFile := c.String("database-file")
		if databaseFile == "" {
			if c.IsSet("database-file") {
				databaseFile = "file::memory:?cache=shared"
			} else {
				databaseFile = fmt.Sprintf("./ezprep-%d.db", t)
			}
		}
		var err error
		if !strings.HasPrefix(databaseFile, "file::memory") {
			databaseFile, err = filepath.Abs(databaseFile)
			if err != nil {
				return errors.Wrap(err, "failed to get absolute path")
			}
		}
		db, closer, err := database.Open("sqlite:"+databaseFile, &gorm.Config{})
		if err != nil {
			return errors.Wrapf(err, "failed to open database %s", databaseFile)
		}

		defer closer.Close()

		// Step 1, initialize the database
		err = admin.InitHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}

		// Step 2, create a preparation
		outputDir := c.String("output-dir")
		var outputStorages []string
		if outputDir != "" {
			err = os.MkdirAll(outputDir, 0755)
			if err != nil {
				return errors.Wrap(err, "failed to create output directory")
			}

			_, err = storage.CreateStorageHandler(c.Context, db, "local", storage.CreateRequest{
				Name: "output",
				Path: outputDir,
			})
			if err != nil {
				return errors.Wrap(err, "failed to create output storage")
			}
			outputStorages = []string{"output"}
		}

		_, err = storage.CreateStorageHandler(c.Context, db, "local", storage.CreateRequest{
			Name: "source",
			Path: path,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create source storage")
		}

		_, err = dataprep.CreatePreparationHandler(c.Context, db, dataprep.CreateRequest{
			SourceStorages: []string{"source"},
			OutputStorages: outputStorages,
			MaxSizeStr:     c.String("max-size"),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create preparation")
		}

		// Step 3, start dataset worker
		worker := datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    c.Int("concurrency"),
				EnableScan:     true,
				EnablePack:     true,
				EnableDag:      true,
				ExitOnComplete: true,
			})
		err = worker.Run(c.Context)
		if err != nil {
			return errors.Wrap(err, "failed to run dataset worker")
		}

		// Step 4, Initiate dag gen
		_, err = dataprep.StartDagGenHandler(c.Context, db, 1, "source")
		if err != nil {
			return errors.Wrap(err, "failed to start dag gen")
		}

		// Step 5, start dataset worker again
		err = worker.Run(c.Context)
		if err != nil {
			return errors.Wrap(err, "failed to run dataset worker")
		}

		// Step 6, print all information
		pieceLists, err := dataprep.ListPiecesHandler(
			c.Context, db, 1,
		)
		if err != nil {
			return errors.Wrap(err, "failed to list pieces")
		}

		cliutil.PrintToConsole(c, pieceLists[0].Pieces)
		return nil
	},
}
