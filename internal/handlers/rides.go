package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"uber-app-backend/internal/db"
)

type Ride struct {
	OriginAddress      string  `json:"origin_address"`
	DestinationAddress string  `json:"destination_address"`
	OriginLat          float64 `json:"origin_latitude"`
	OriginLng          float64 `json:"origin_longitude"`
	DestinationLat     float64 `json:"destination_latitude"`
	DestinationLng     float64 `json:"destination_longitude"`
	RideTime           string  `json:"ride_time"`
	FarePrice          float64 `json:"fare_price"`
	PaymentStatus      string  `json:"payment_status"`
	DriverID           int     `json:"driver_id"`
	UserID             int     `json:"user_id"`
}

func CreateRide(w http.ResponseWriter, r *http.Request) {
	var ride Ride
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.Pool.Exec(
		context.Background(),
		`INSERT INTO rides (
            origin_address, destination_address, origin_latitude, origin_longitude, destination_latitude, destination_longitude,
            ride_time, fare_price, payment_status, driver_id, user_id
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		ride.OriginAddress, ride.DestinationAddress, ride.OriginLat, ride.OriginLng,
		ride.DestinationLat, ride.DestinationLng, ride.RideTime,
		ride.FarePrice, ride.PaymentStatus, ride.DriverID, ride.UserID,
	)

	if err != nil {
		http.Error(w, "Error creating ride", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "ride created"})
}
