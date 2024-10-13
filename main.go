package main

import (
	"DatsMagic/mapinfo"
	"embed"
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"net/http"
	"os"
)

// assets is a directory containing static assets for the game.
//
//go:embed assets/*
var assets embed.FS

// init loads values from .env into the system. If no .env file is found, it
// simply logs a message and continues.
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// main sets up an HTTP server that listens on port 2222. It loads all
// static assets from the "assets" folder and serves them at the root URL.
// Additionally, it serves the "/cactus" and "/map" endpoints, which are
// handled by the mapinfo.Mapinfo type. The mapinfo.Mapinfo.Loop method is
// run in a goroutine to update the game state.
func main() {
	mux := http.NewServeMux()
	sub, err := fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}
	token, ok := os.LookupEnv("DATSTEAM_TOKEN")
	if !ok {
		panic("DATSTEAM_TOKEN not set")
	}
	m := mapinfo.Mapinfo{Token: token}
	go m.Loop()
	mux.HandleFunc("/cactus", m.GetCactus)
	mux.HandleFunc("/map", m.GetMapHandle)
	mux.Handle("/", http.FileServer(http.FS(sub)))
	http.ListenAndServe(":2222", mux)
}
