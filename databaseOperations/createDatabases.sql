CREATE TABLE
    Trains (
        TrainID BIGSERIAL PRIMARY KEY,
        Name VARCHAR(255) NOT NULL,
        TrainNumber VARCHAR(20) NOT NULL,
        OperationalDays VARCHAR(50) NOT NULL
    );

CREATE TABLE
    Classes (
        ClassID BIGSERIAL PRIMARY KEY,
        TrainID INT,
        ClassName VARCHAR(10) NOT NULL,
        Dimension INT,
        NumBogies INT,
        FOREIGN KEY (TrainID) REFERENCES Trains (TrainID)
    );

CREATE TABLE
    Cities (
        CityID BIGSERIAL PRIMARY KEY,
        ShortformName VARCHAR(10) NOT NULL UNIQUE,
        CityName VARCHAR(255) NOT NULL UNIQUE
    );

CREATE TABLE
    TrainCities (
        TrainCityID BIGSERIAL PRIMARY KEY,
        TrainID INT NOT NULL,
        CityIndex INT NOT NULL UNIQUE,
        CityID INT NOT NULL UNIQUE,
        CityName VARCHAR(255) NOT NULL UNIQUE,
        ShortformName VARCHAR(10) NOT NULL UNIQUE,
        ArrivalTime TIME,
        DepartureTime TIME,
        FOREIGN KEY (TrainID) REFERENCES Trains (TrainID),
        FOREIGN KEY (CityID) REFERENCES Cities (CityID)
    );

CREATE TABLE
    Users (
        UserID BIGSERIAL PRIMARY KEY,
        UserName VARCHAR(255) NOT NULL,
        Password VARCHAR(255) NOT NULL,
        FullName VARCHAR(255) NOT NULL
    );

CREATE TABLE
    Journeys (
        JourneyID BIGSERIAL PRIMARY KEY,
        UserID INT,
        DepartureCityID INT,
        ArrivalCityID INT,
        DepartureDate DATE,
        JourneyCode VARCHAR(255) NOT NULL, -- Assuming a series of string as journey code
        FOREIGN KEY (UserID) REFERENCES Users (UserID),
        FOREIGN KEY (DepartureCityID) REFERENCES Cities (CityID),
        FOREIGN KEY (ArrivalCityID) REFERENCES Cities (CityID)
    );