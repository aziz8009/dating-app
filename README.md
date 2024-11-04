# Dating App Backend

## Overview

This is the backend service for a dating mobile application, similar to Tinder or Bumble. The service is built using Golang and utilizes MySQL as the database. It provides functionalities for user authentication, swiping profiles, and premium feature management.

## Features

- User sign up and login
- Profile swiping functionality (like or pass)
- Daily swipe quota management
- Premium packages for enhanced features
- RESTful API with comprehensive error handling

## Tech Stack

- **Backend**: Golang
- **Database**: MySQL
- **Framework**: Gin (HTTP web framework for Go)
- **ORM**: GORM (Go Object Relational Mapping)
- **Testing**: Go testing framework

## Installation

### Prerequisites

- Go (1.16 or later)
- MySQL (5.7 or later)
- Git

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/dating-app-backend.git
   cd dating-app-backend

2. Create a MySQL database:

  ```
  CREATE DATABASE dating_app;
  ```
  
3. Update the database configuration in internal/config/config.go:

  ```
  type Config struct {
    DBUser     string
    DBPassword string
    DBHost     string
    DBPort     string
    DBName     string
  } 

  ```
4. Install the required Go modules:

  ```
  go mod tidy

  ```

5. Run the application:

  ```
 go run cmd/main.go

  ```

The server will start at http://localhost:5000.

## API Reference

#### Get all items

```http
  POST /signup
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your API key |
| `email`     | `string` | **Required**. Your email |
| `password` | `string` | **Required**. Your password |

#### Signup

```http
  POST /login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required**. Email |
| `password`      | `string` | **Required**. password |

#### Login


```http
  POST /swipe
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `profile_id`      | `number` | **Required**. profile_id |
| `action`      | `string` | **Required**. action |

#### swipe


```http
  Get /swipes/daily
```

#### get swipe daily

```http
  GET /packages
```
#### packages


```http
  POST /purchase
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `package_id`      | `number` | **Required**. package_id |

#### purchase


