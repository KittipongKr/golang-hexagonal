package cmd

import (
	"log"

	envCfgs "csat-servay/configs/env"
	cFiber "csat-servay/internal/adapter/fiber/controllers"
	rFiber "csat-servay/internal/adapter/fiber/routes"
	rFiberV1 "csat-servay/internal/adapter/fiber/routes/v1"
	mongoCfgs "csat-servay/internal/adapter/mongo"
	repo "csat-servay/internal/adapter/mongo/repository"
	serv "csat-servay/internal/core/service"

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

		//NOTE: new module repository
		userRepo := repo.NewUserRepo(mongoCols)

		//NOTE: new module service
		userServ := serv.NewUserServ(userRepo)

		// NOTE: new module controller
		pingCtls := cFiber.NewPingCtl()
		userCtls := cFiber.NewUserCtls(userServ)

		// NOTE: launch go fiber server
		if err := rFiber.Launch(
			rFiber.FiberRoute(&env.Fiber, &rFiber.Controller{
				V1: rFiberV1.V1{
					Ping: pingCtls,
					User: userCtls,
				},
			}),
		); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(startCommand)
}
