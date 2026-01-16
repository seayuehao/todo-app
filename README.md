# todo-list server

## Database Schema

The application uses a relational database managed via GORM.
Two main entities are defined: User and Todo.

User Table

Stores user identity and authentication provider information.


| Column Name   | Type     | Constraints / Indexes | Description                                                   |
| ------------- | -------- | --------------------- | ------------------------------------------------------------- |
| `id`          | CHAR(36) | Primary Key           | Unique user identifier (UUID)                                 |
| `email`       | VARCHAR  | Unique Index          | User email address                                            |
| `username`    | VARCHAR  | Indexed               | Display name / username                                       |
| `provider_id` | VARCHAR  | Indexed               | Unique identifier from OAuth provider                         |
| `provider`    | VARCHAR  | —                     | Authentication provider (e.g. `github`, `google`, `facebook`) |

Todo Table

Stores TODO items created by users

| Column Name  | Type      | Constraints / Indexes       | Description                            |
| ------------ | --------- | --------------------------- | -------------------------------------- |
| `id`         | INTEGER   | Primary Key, Auto Increment | Unique TODO identifier                 |
| `user_id`    | CHAR(36)  | Indexed                     | Reference to the owning user (User ID) |
| `title`      | VARCHAR   | —                           | TODO item title                        |
| `completed`  | BOOLEAN   | —                           | Completion status                      |
| `created_at` | TIMESTAMP | —                           | Creation timestamp                     |
| `updated_at` | TIMESTAMP | Nullable                    | Last update timestamp                  |


# running app via source
```
go 1.24.4+ is required

 - run
cd todo-app

go mod download

If you use github sign in integration, replace client id & secret.
Visit github -> settings -> develpers settings -> oauth apps, configure then  
  
  Application name： 
    - xxx

 Homepage URL  
    - http://localhost:8091

    Authorization callback URL:
    - http://localhost:8091/auth/github/callback

modify conifg/config.example.yaml.


finally run :

go run main.go

```

# docker run
```
modify conifg/config.example.yaml first;

docker build -t todo-app .

docker run -p 8091:8091 todo-app

```

## docker compose
```
    docker compose up --build

```


# health check
```shell

curl http://localhost:8091/echo

```