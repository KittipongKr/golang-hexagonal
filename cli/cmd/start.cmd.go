package cmd

import (
	"log"

	envCfgs "csat-servay/configs/env"
	calls "csat-servay/internal/adapter/calls"
	rFiber "csat-servay/internal/adapter/fiber/routes"
	mongo "csat-servay/internal/adapter/mongo"

	// arch "csat-servay/internal/core"
	core "csat-servay/internal/core"
	"csat-servay/pkg/logs"
	"csat-servay/pkg/tracing"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/trace"
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

		//NOTE: start jaeger tracer
		var tp trace.TracerProvider
		var tpErr error
		if env.Tracer.Host != "" || env.Tracer.Port != "" {
			tp, tpErr = tracing.NewTracerProvider(env.Tracer.Host+":"+env.Tracer.Port, env.App.AppName, env.App.Version, env.App.Environment)
			if tpErr != nil {
				log.Fatal(tpErr)
			}
		}

		//NOTE: connect to mongo database
		mongoCols, err := mongo.Connect(&env.Mongo, tp)
		if err != nil {
			log.Fatalf("Failed to connect to mongo database: %v", err)
		}

		if tp == nil {
			log.Fatal("Tracer provider is not set")
		}

		//NOTE: launch zap logger
		logs.Launch()

		//NOTE: core - hexagonal architecture
		handler := core.SetHandlers(
			tp,
			core.SetAdaptors(
				tp,
				*env,
				core.Adaptor{
					MongoAdaptor: mongo.SetAdaptor(mongoCols),
					CallsAdaptor: calls.SetAdaptor(tp, resty.New(), *env),
				},
			),
		)

		// NOTE: launch go fiber server
		if err := rFiber.Launch(
			rFiber.FiberRoute(&env.Fiber, &handler.Router, tp),
		); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(startCommand)
}
