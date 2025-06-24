# ef_mb_test

A RESTful service for managing and enriching people's information  
Developed in Go using the standard net/http library and GORM  
Integration with external APIs for age, gender, and nationality prediction  
Automatic enrichment of person data when a new record is created  
Data storage in PostgreSQL database  
Implemented CRUD operations, pagination, and containerization with Docker

## Features

- CRUD operations for managing people records
- Automatic enrichment of person data using:
  - [Agify.io](https://api.agify.io/) for age prediction
  - [Genderize.io](https://api.genderize.io/) for gender prediction
  - [Nationalize.io](https://api.nationalize.io/) for nationality prediction
- Pagination support
- PostgreSQL database for data persistence
- Docker containerization

## API Endpoints

### GET /people
Get all people with pagination
- Query parameters:
  - `page` (default: 1)
  - `pageSize` (default: 10, max: 100)

### GET /people/{id}
Get person by ID

### POST /people
Create new person
```json
{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich" // optional
}
```

### PUT /people
Update existing person
```json
{
    "id": 1,
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich",
    "age": 30,
    "gender": "male",
    "nationality": ["RU", "UA"]
}
```

### DELETE /people/{id}
Delete person by ID

## Technology Stack

- Go 1.24.3
- PostgreSQL 17.0
- GORM
- Docker & Docker Compose
- net/http

## Running the Project

1. Create `.env` file for application configuration:
```env
HOST=db
USER_DB=your_user
PASSWORD_DB=your_password
NAME_DB=your_db_name
PORT=5432
```

2. Create `db.env` file for PostgreSQL configuration:
```env
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_db_name
```

3. Run using Docker Compose:
```bash
docker-compose up -d
```

The service will be available at `http://localhost:8080`