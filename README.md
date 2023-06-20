# Company Inc.
A simple microservice for performing CRUD operations on a company or companies object(s).

## Running the Application

### Docker
The easiest way to run the application is via Docker. Make sure that Docker and Docker Compose are installed on your system. Run the following command:

```sh
docker-compose up
# or
docker compose up
```

This will set up a MongoDB database container as well as a Kafka container along with the application container.

You can also modify the environment variables in the `docker-compose.yaml` file according to your own need.

### Local
In case Docker does not work, the application can also be run locally. The following applications are required as prerequisites in this case:

- Go 1.18+
- MongoDB Server 6.0+
- Kafka Server 3.4+

Add the environment variables mentioned the `config/config.go` file including MongoDB and Kafka credentials.

Now run these command:

```sh
go mod download
go run main.go
```

## APIs and Endpoints

The following endpoints are available in the application along with their HTTP methods:


### POST /v1/signup
This method will create a user for accessing the rest of the application.

#### Sample Request

```json
{
    "name": "John Smith",
    "email": "abc@123.com",
    "password": "abc12345"
}
```

### POST /v1/login
This method will allow the user to log into the application. After logging in, it will send a cookie to the client application.

#### Sample Request

```json
{
    "email": "abc@123.com",
    "password": "abc12345"
}
```

### POST /v1/logout
This method will log the already logged in user out of the application. After logging out, it will delete the cookie from the client application. The request body will be empty in this case.

### POST /v1/companies
This method is used to create a new company in the application.

#### Sample Request

```json
{
    "name": "Company 3",
    "description": "This is a company.",
    "total_employees": 50,
    "registered": true,
    "type": "Corporations"
}
```

### GET /v1/companies/:id
This method will return a single company against the company ID provided. The request body will be empty in this case.

### PATCH /v1/companies/:id
This method will update any of the field of a specific company.

#### Sample Request

```json
{
    "name": "Company 3",
    "type": "Sole Proprietorship"
}
```

### DELETE /v1/companies/:id
This method will delete a specific company from the database. The request body will be empty in this case.

> **Note:** The POST, PATCH and DELETE methods in the above instances will require the user to be logged in to be accessible.
