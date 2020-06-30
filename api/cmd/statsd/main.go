package main

import (
	"github.com/deissh/datadog-client"
	"github.com/deissh/rl/api/pkg"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var Version string
var Commit string
var Branch string
var BuildTimestamp string

func main() {
	// loading configuration
	config.WithOptions(config.ParseEnv, config.Readonly)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yaml")
	if err != nil {
		panic(err)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if config.Bool("debug", false) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: false,
			},
		).With().Caller().Logger()
	}

	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Msg("Starting Statsd")

	if !config.Bool("server.datadog.enable") {
		log.Fatal().Msg("Datadog disabled")
	}

	log.Info().
		Msg("Loaded configuration and logger")
	log.Info().
		Msg("Start initialize database and redis")

	pkg.InitializeDB()
	pkg.InitializeRedis()

	log.Info().
		Msg("Initialize database and redis successful done")

	log.Info().
		Msg("Start initialize datadog")

	datadog.InitializeClient()
	datadog.SetPrefix(config.String("server.datadog.prefix", "LAZER_"))
	datadog.AddTag("debug", config.String("debug"))

	log.Info().
		Msg("Initialize datadog successful done")

	log.Info().
		Msg("Creating datadog tasks")

	datadog.RunGaugeTask(
		"user_online",
		time.Minute,
		datadog.Tags{},
		getUsersOnline,
	)
	datadog.RunGaugeTask(
		"users",
		time.Hour,
		datadog.Tags{},
		getActiveUsers,
	)

	datadog.Start()
}

func getUsersOnline() (f float64, err error) {
	var count int
	err = pkg.Db.
		QueryRow("SELECT count('id') FROM users WHERE check_online(last_visit) = true").
		Scan(&count)

	return float64(count), err
}

func getActiveUsers() (f float64, err error) {
	var count int
	err = pkg.Db.
		QueryRow("SELECT count('id') FROM users WHERE is_active = true").
		Scan(&count)

	return float64(count), err
}
