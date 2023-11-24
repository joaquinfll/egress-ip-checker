package main

import (
	"fmt"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"
)

func main() {

	os.Setenv("EGRESS_IP", "127.0.0.1")
	egressIp, ok := os.LookupEnv("EGRESS_IP")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil).WithAttrs([]slog.Attr{slog.String("app-version", "v0.0.1"),
		slog.String("app-name", "egress-ip-checker"),
	})

	logger := slog.New(jsonHandler)
	//ctx := slog.NewContext(context.Background(), logger)
	slog.SetDefault(logger)

	if !ok {
		logger.Error("EGRESS_IP environment variable is not defined")
		logger.Error("Define an environment variable named EGRESS_IP with a valid IPv4")
		os.Exit(1)
	} else {
		fmt.Printf("EGRESS_IP: %s\n", egressIp)
	}

	logger.Info("Starting application")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIp := getClientIpAddr(r)

		if clientIp == egressIp {
			fmt.Fprintf(w, "Success")
			logger.Info("Request successful client IP match")
		} else {
			fmt.Fprintf(w, "Error: client IP did not match")
			logger.Error("Request unsuccessful client IP not matched")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}
