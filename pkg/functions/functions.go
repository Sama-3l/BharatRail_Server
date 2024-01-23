package functions

import (
	"database/sql"
)

func FindOrInsertCity(db *sql.DB, cityName string, shortFormName string) (int64, error) {
	var cityID int64

	// Check if the city already exists
	err := db.QueryRow("SELECT CityID FROM Cities WHERE CityName = $1", cityName).Scan(&cityID)
	if err == nil {
		// City already exists, return its ID
		return cityID, nil
	} else if err != sql.ErrNoRows {
		// An error occurred while querying
		return 0, err
	}

	// City doesn't exist, insert a new one
	err = db.QueryRow("INSERT INTO cities(shortformname, cityname) VALUES ($1, $2) RETURNING CityID", shortFormName, cityName).Scan(&cityID)
	if err != nil {
		return 0, err
	}

	// Return the newly inserted or existing CityID
	return cityID, nil
}
