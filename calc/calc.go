package calc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"trip-calc-service/storage"
	"trip-calc-service/structures"
)

func CalculateTrip(token string) {
	var redisStorage = storage.RedisConnect()

	itemFromStorage := redisStorage.Get(token)

	t := &structures.Trip{}

	_ = json.Unmarshal([]byte(itemFromStorage), t)

	resp := runGet(createUrl(t))

	t.Distance, t.TravelTime = parseResponse(resp)

	result, err := json.Marshal(t)

	if err != nil {
		fmt.Println("Error marshal trip", err)
	}

	redisStorage.Set(t.Token, string(result))
}

func runGet(urlStr string) []byte {

	req := &http.Request{
		Method: http.MethodGet,
	}

	req.URL, _ = url.Parse(urlStr)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error runtime request ", err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Empty response ", err)
	}
	return respBody
}

func createUrl(trip *structures.Trip) string {

	var urlStr string
	urlStr = os.Getenv("CALC_ROUTE_URL") + trip.LatitudeFrom + "," + trip.LongitudeFrom + ";"
	urlStr += trip.LatitudeTo + "," + trip.LongitudeTo
	return urlStr
}

func parseResponse(body []byte) (float64, float64) {

	var result map[string]interface{}

	_ = json.Unmarshal(body, &result)

	routes := result["routes"].([]interface{})
	route := routes[0].(map[string]interface{})

	return route["duration"].(float64), route["distance"].(float64)
}
