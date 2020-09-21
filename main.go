package trip_calc_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"trip-calc-service/calc"
	"trip-calc-service/storage"
	"trip-calc-service/structures"
)

var redisStorage = storage.RedisConnect()

var queueCalc = make(chan string, 1)

func handlerRequest(w http.ResponseWriter, r *http.Request) {

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
	// fmt.Println("Channel to -> ", trip.Token)
	_, _ = w.Write(result)
}

func handlerResult(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("token")
	fmt.Println("Token from request: ", token)

	ride := redisStorage.Get(token)
	// TODO  добавить проверку, если запрос есть, а ответа нет, еще раз закинуть его в обработку

	_, _ = w.Write([]byte(ride))
}

func worker(queue chan string) {
	for {
		token := <-queue
		go calc.CalculateTrip(token)
	}
}

func main() {

	go worker(queueCalc)

	http.HandleFunc("/trip/calc/request", handlerRequest)
	http.HandleFunc("/trip/calc/result", handlerResult)

	err := http.ListenAndServe(":8080", nil)
	if err == nil {
		fmt.Println("Start calc structures service on :8080 ... ")
	}
}
