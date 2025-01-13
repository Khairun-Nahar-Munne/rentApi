# Cat APP Go Project

 We will create a backend API using the Beego framework to serve the property list and data from database.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Technologies Used](#technologies-used)
3. [Key Features](#key-features)
4. [Installation](#installation)
5. [Configuration](#configuration)
6. [Running the Project](#running-the-project)
6. [Running the Test](#running-the-test)
7. [API Endpoints](#api-endpoints)
8. [Project Structure](#project-structure)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go (version 1.16 or later)
- Git
- Beego

## Technologies Used

- [Go](https://golang.org/) - The main programming language
- [Beego v2](https://github.com/beego/beego) - Web framework for building the API
- [TailwindCSS](https://tailwindcss.com/) - Front-end css framework


## Key Features

## Key Features
Fetch data from database.
- **Backend API using Beego Framework:**
  - **/v1/property/list:** Endpoint to provide a list of properties for a specific location.
  - **/v1/property/details:** Endpoint to provide detailed information about a specific property.
## Installation

### 1. Install Go

If you haven't installed Go, follow these steps:

1. Visit the official Go downloads page: https://golang.org/dl/
2. Download the appropriate installer for your operating system.
3. Follow the installation instructions for your OS:
   - Windows: Run the MSI installer and follow the prompts.
   - macOS: Open the package file and follow the prompts.
   - Linux: Extract the archive to `/usr/local`:
     ```
     tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
     ```
4. Add Go to your PATH:
   - For bash, add the following to your `~/.bashrc` or `~/.bash_profile`:
     ```
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
   - For other shells, add the equivalent to your shell's configuration file.
5. Verify the installation by opening a new terminal and running:
      ```
      go version
      ```
6. create src to keep your  all beego app here
      ```
      mkdir -p $GOPATH/src/
      ```
### 2. Setup This Project First to Get Data From Api

You have to run rent project following the instructions given in the gitub repository of the project. Otherwise, you can not get the data.

The  project github repo link: https://github.com/Khairun-Nahar-Munne/rent

You have to keep running the docker to fetch data from the database.

### 2. Clone the Repository

Clone this repository to your local machine in the default Go path:

```
cd $GOPATH/src/ 
git clone https://github.com/Khairun-Nahar-Munne/rentApi.git
cd rentApi
```

### 3. Install Dependencies

This project uses Beego v2 and other dependencies. Install them using the following commands:

```
go get github.com/beego/beego/v2
go get github.com/astaxie/beego/logs
go mod tidy
```


## Configuration

### 2. Configuration File

Create a `conf/app.conf` file in the project root with the following content:

```ini
appname = rentApi
httpport = 8000
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true
sqlconn = 
```
You can use your api_key if you blocked using my api_key
## Running the Project

To run the project, use the following command from the project root directory:

1. For using postgres and pgadmin:
```
docker-compose up
```
2. Open another terminal and run the project

```
bee run
```

If you don't have `bee` installed, you can install it with:

```
go get github.com/beego/bee/v2
```

3. To fetch the data from booking.com api:

- Run this `http://localhost:8000/v1/property/list` Feth data from location and rental property tables

## API Endpoints

The project provides the following API endpoints:

- `GET /v1/property/list`: Feth data from location and rental property tables
- `GET /v1/property/details`: Fetch data from rental property and property details tables

## Project Structure

The project follows a standard Beego directory structure:

```
rent/
├── conf/
│   └── app.conf
├── controllers/
│   ├── propertyDetails_controller.go
│   └── property_controller.go
├── routers/
│   └── router.go
├── models/
│   └── models.go
├── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

- `conf/`: Contains configuration files
- `controllers/`: Contains the logic for handling API requests (fetcing and store data from booking.com api and fetch the data from local database)
- `routers/`: Defines the routing for the application
- `models/`: Define structure of the data of database
- `main.go`: The entry point of the application

