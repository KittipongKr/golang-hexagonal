package cmd

import (
	"log"

	envCfgs "csat-servay/configs/env"
	"csat-servay/internal/adapter/calls"
	rFiber "csat-servay/internal/adapter/fiber/routes"
	mongo "csat-servay/internal/adapter/mongo"
	mongoCfgs "csat-servay/internal/adapter/mongo"

	// arch "csat-servay/internal/core"
	core "csat-servay/internal/core"
	"csat-servay/pkg/logs"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "server started",
	Run: func(cmd *cobra.Command, args []string) {
		// NOTE: check app zone is set
		if zone == "" {
			log.Println("Environment profile is not set")
			return
		}

		//NOTE: read enverinment form env file
		env := envCfgs.ReadEnv(zone)

		//NOTE: connect to mongo database
		mongoCols, err := mongoCfgs.Connect(&env.Mongo)
		if err != nil {
			log.Fatalf("Failed to connect to mongo database: %v", err)
		}

		//NOTE: launch zap logger
		logs.Launch()

		//NOTE: core - hexagonal architecture
		handler := core.SetHandlers(
			core.SetAdaptors(
				*env,
				core.Adaptor{
					MongoAdaptor: mongo.SetAdaptor(mongoCols),
					CallsAdaptor: calls.SetAdaptor(resty.New(), *env),
				},
			),
		)

		// NOTE: launch go fiber server
		if err := rFiber.Launch(
			rFiber.FiberRoute(&env.Fiber, &handler.Router),
		); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(startCommand)
}
