package main

import (
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"strconv"
)

func main() {

	externalUrl, ok := os.LookupEnv("EXTERNAL_URL")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil).WithAttrs([]slog.Attr{slog.String("app-version", "v0.0.1"),
		slog.String("app-name", "egress-ip-checker-client"),
	})

	logger := slog.New(jsonHandler)

	if !ok {
		logger.Error("EXTERNAL_URL environment variable is not defined")
		logger.Error("Define an environment variable named EXTERNAL_URL with a valid URL")
		os.Exit(1)
	} else {
		resp, err := http.Get(externalUrl)
		if err != nil {
			slog.Error("Cannot call %s", externalUrl)
			os.Exit(1)
		} else {
			slog.Info(strconv.Itoa(resp.StatusCode))
		}
	}

}
