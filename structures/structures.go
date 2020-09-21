package structures

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Trip struct {
	LatitudeFrom  string  `json:"latitudeFrom"`
	LongitudeFrom string  `json:"longitudeFrom"`
	LatitudeTo    string  `json:"latitudeTo"`
	LongitudeTo   string  `json:"longitudeTo"`
	Token         string  `json:"token"`
	Distance      float64 `json:"distance"`
	TravelTime    float64 `json:"travelTime"`
}

func (t *Trip) UpdateToken() {
	var strForHash = t.LongitudeTo + t.LatitudeTo + time.Now().String()

	hash := md5.Sum([]byte(strForHash))
	t.Token = hex.EncodeToString(hash[:])
}

func (t *Trip) Serialize() ([]byte, error) {

	result, err := json.Marshal(t)

	if err != nil {
		fmt.Println("Error marshal trip", err)
	}
	return result, err
}
