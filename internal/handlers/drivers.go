package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"uber-app-backend/internal/db"
)

func GetDrivers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Pool.Query(context.Background(), `SELECT * FROM drivers`)
	if err != nil {
		http.Error(w, "Error fetching drivers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	drivers := []map[string]any{}
	for rows.Next() {
		values, _ := rows.Values()
		columns := rows.FieldDescriptions()
		record := make(map[string]any)
		for i, col := range columns {
			record[string(col.Name)] = values[i]
		}
		drivers = append(drivers, record)
	}

	json.NewEncoder(w).Encode(map[string]any{"data": drivers})
}
