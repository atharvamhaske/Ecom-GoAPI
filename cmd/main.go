package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	
	if err := api.run(api.mount()); //mount all endpoints and give that to server to run can also do like h := api.mount()
	//then api.run(h) but in run check for error as it returns error
	err != nil {
		log.Printf("server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
