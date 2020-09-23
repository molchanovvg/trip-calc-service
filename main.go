package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trip-calc-service/calc"
	"trip-calc-service/storage"
	"trip-calc-service/structures"
)

var redisStorage = storage.RedisConnect()

var queueCalc = make(chan string, 1)

func handlerRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("incoming request ... ")
	trip := &structures.Trip{
		LatitudeFrom:  r.FormValue("latitudeFrom"),
		LongitudeFrom: r.FormValue("latitudeFrom"),
		LatitudeTo:    r.FormValue("latitudeTo"),
		LongitudeTo:   r.FormValue("longitudeTo"),
	}

	trip.UpdateToken()

	result, err := json.Marshal(trip)

	if err != nil {
		fmt.Println("Error marshal trip", err)
	}

	redisStorage.Set(trip.Token, string(result))

	queueCalc <- trip.Token

	_, _ = w.Write(result)
}

func handlerResult(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("token")

	itemFromStorage := redisStorage.Get(token)

	t := &structures.Trip{}

	_ = json.Unmarshal([]byte(itemFromStorage), &t)

	if t.Distance == 0 || t.TravelTime == 0 {
		w.WriteHeader(http.StatusTooEarly)
		return
	}
	result := "Distance (m):" + fmt.Sprintf("%f", t.Distance) + "\n"
	result += "Time (s): " + fmt.Sprintf("%f", t.TravelTime)

	_, _ = w.Write([]byte(result))
}

func worker(queue chan string) {
	for {
		token := <-queue
		go calc.CalculateTrip(token)
	}
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_ = srv.Shutdown(ctx)

	fmt.Println("Shutting down")
	os.Exit(0)
}

func main() {

	fmt.Println("Start main ... ")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	go worker(queueCalc)

	mux := http.NewServeMux()
	mux.HandleFunc("/trip/calc/request", handlerRequest)
	mux.HandleFunc("/trip/calc/result", handlerResult)

	server := &http.Server{
		Addr:    ":" + os.Getenv("SERVICE_PORT"),
		Handler: mux,
	}

	go func() {
		fmt.Println("Starting service on " + os.Getenv("SERVICE_PORT") + " ... ")
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Error start server", err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(server)
}
