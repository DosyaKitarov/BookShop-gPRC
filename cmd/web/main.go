package main

import (
	"crypto/tls"
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tls := &tls.Config{CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256}}

	server := &http.Server{
		Addr:      *addr,
		Handler:   routes(),
		ErrorLog:  slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: tls,
	}

	logger.Info("Starting server", "addr", server.Addr)

	err := server.ListenAndServeTLS("tls/cert.pem", "tls/key.pem")
	logger.Error("Server stopped", "error", err)
	os.Exit(1)

}
