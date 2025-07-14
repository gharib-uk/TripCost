package main

import (
	"math/rand"
	"time"
)

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

}
