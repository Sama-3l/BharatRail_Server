# BharatRail Server

BharatRail server is a Golang-based server for the railway ticket booking application. It provides essential functionalities for managing cities, trains, and users. The application is designed to handle administrative tasks such as adding cities and trains, as well as user-related operations like fetching data and adding users.

## Features

**Admin Routes:**
  - Add Cities: Allows administrators to add new cities to the system.
  - Add Trains: Enables the addition of new trains to the system.

**User Routes:**
  - Fetch Data: Provides routes for users to retrieve relevant information.
  - Add Users: Allows users to be added to the system.

## Getting Started

These instructions will guide you through setting up the BharatRail server on your local machine.

### Prerequisites

- Golang installed on your machine.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Sama-3l/BharatRail_Server.git
   cd BharatRail_Server/cmd
2. Run the application
   ```bash
   go run main.go
3. Access the API at http://localhost:8080.

## Usage

### Sample Requests

**Add City:**
    
    POST /admin/add-city
    Content-Type: application/json

    {
        "name": "New Delhi",
        "shortform_name" : "NDLS"
    }

**Add Train:**

    POST /admin/add-train
    Content-Type: application/json

    {
        "name": "TamilNadu Express",
        "train_number": "12522",
        "operational_days": "Mon Wed Fri",
        "classes" : [
            {
                "class_name" : "3A",
                "dimension" : 4,
                "num_bogies" : 12
            },
            {
                "class_name" : "2A",
                "dimension" : 3,
                "num_bogies" : 8
            }
        ],
        "train_cities" : [
            {
                "city_name" : "New Delhi",
                "city_short_form_name" : "NDLS",
                "arrival_time" : "10.10.24:Z15:32:12",
                "departure_time" : "10.10.24:Z15:35:12"
            }
        ]
    }

## Contributing

We welcome contributions from the community! If you'd like to contribute to BharatRail Backend, please follow these guidelines:

1. Fork the repository and create your branch from `main`.
2. Make sure your code follows the established coding standards.
3. Create descriptive commit messages and PR titles.
4. Ensure that your changes pass all tests.
5. Update the README with details of changes if needed.

### Code Style

Please follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) for maintaining consistent coding style.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

We would like to express our gratitude to the following libraries and resources that have been instrumental in the development of BharatRail Backend.

[Golang Projects Tutorials](https://www.youtube.com/playlist?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9)