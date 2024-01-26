package models

import (
	"bharatrail_server/pkg/config"
	"bharatrail_server/pkg/functions"
	"database/sql"
	"log"
	"time"
)

// Train model
type Train struct {
	TrainID         int64       `json:"train_id"`
	Name            string      `json:"name"`
	TrainNumber     string      `json:"train_number"`
	OperationalDays string      `json:"operational_days"`
	Classes         []Class     `json:"classes"`
	TrainCities     []TrainCity `json:"train_cities"`
}

// Class model
type Class struct {
	ClassID   int64  `json:"class_id"`
	TrainID   int64  `json:"train_id"`
	ClassName string `json:"class_name"`
	Dimension int    `json:"dimension"`
	NumBogies int    `json:"num_bogies"`
}

// City model
type City struct {
	CityID        int64  `json:"city_id"`
	ShortformName string `json:"shortform_name"`
	CityName      string `json:"city_name"`
}

// TrainCity model
type TrainCity struct {
	TrainCityID       int64      `json:"train_city_id"`
	TrainID           int64      `json:"train_id"`
	CityIndex         int64      `json:"city_index"`
	CityName          string     `json:"city_name"`
	CityShortFormName string     `json:"city_short_form_name"`
	CityID            int64      `json:"city_id"`
	ArrivalTime       *time.Time `json:"arrival_time"`
	DepartureTime     *time.Time `json:"departure_time"`
}

// User model
type User struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

// Journey model
type Journey struct {
	JourneyID       int64     `json:"journey_id"`
	UserID          int64     `json:"user_id"`
	DepartureCityID int64     `json:"departure_city_id"`
	ArrivalCityID   int64     `json:"arrival_city_id"`
	DepartureDate   time.Time `json:"departure_date"`
	JourneyCode     string    `json:"journey_code"`
}

func AddCityToDatabase(city City) (City, error) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO cities(shortformname, cityname) VALUES ($1, $2) RETURNING cityid`

	err := db.QueryRow(sqlStatement, city.ShortformName, city.CityName).Scan(&city.CityID)

	if err != nil {
		log.Printf("%v\n", err)
	}

	return city, err
}

func AddTrainToDatabase(train Train) (Train, error) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO trains(name, trainnumber, operationaldays) VALUES ($1, $2, $3) RETURNING trainid`

	err := db.QueryRow(sqlStatement, train.Name, train.TrainNumber, train.OperationalDays).Scan(&train.TrainID)

	if err != nil {
		log.Println(err)
	}

	train, err = AddTrainCityToDatabase(train, db)
	if err != nil {
		log.Println(err)
	}

	train, err = AddClassesToDatabase(train, db)
	if err != nil {
		log.Println(err)
	}

	return train, err
}

func AddTrainCityToDatabase(train Train, db *sql.DB) (Train, error) {
	var err error

	for i, value := range train.TrainCities {
		train.TrainCities[i].TrainID = train.TrainID
		train.TrainCities[i].CityIndex = int64(i)
		train.TrainCities[i].CityID, err = functions.FindOrInsertCity(db, value.CityName, value.CityShortFormName)
		if err != nil {
			log.Println(err)
		}
		err = db.QueryRow("INSERT INTO traincities(trainid, cityindex, cityname, shortformname, cityid, arrivaltime, departuretime) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING traincityid", train.TrainID, train.TrainCities[i].CityIndex, value.CityName, value.CityShortFormName, train.TrainCities[i].CityID, value.ArrivalTime, value.DepartureTime).Scan(&train.TrainCities[i].TrainCityID)
		if err != nil {
			log.Println(err)
		}
	}
	return train, err
}

func AddClassesToDatabase(train Train, db *sql.DB) (Train, error) {
	var err error
	for i, value := range train.Classes {
		train.Classes[i].TrainID = train.TrainID
		err = db.QueryRow("INSERT INTO classes(trainid, classname, dimension, numbogies) VALUES ($1, $2, $3, $4) RETURNING classid", train.TrainID, value.ClassName, value.Dimension, value.NumBogies).Scan(&train.Classes[i].ClassID)
		if err != nil {
			log.Println(err)
		}
	}
	return train, err
}
