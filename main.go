package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Hello world info")
	log.Warn().Err(errors.New("testaja")).Msg("test aja kok")

	log.Warn().Err(errors.New("testaja")).Msg("test aja dong")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	t := time.Now().In(loc)
	hour := t.Format("15:04")
	fmt.Println(hour)
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {

	http.HandleFunc("/test", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
