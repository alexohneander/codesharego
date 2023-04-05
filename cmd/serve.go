package cmd

import (
	"github.com/alexohneander/codesharego/server"
	"github.com/alexohneander/codesharego/turnserver"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func serveCmd(version string) cli.Command {
	return cli.Command{
		Name: "serve",
		Action: func(ctx *cli.Context) {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)

			log.Info().Msg("HTTP server started")
			go server.Listen()

			log.Info().Msg("TURN server started")
			turnserver.Start()
		},
	}
}
