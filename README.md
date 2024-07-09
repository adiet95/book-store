# Book Store
Simple app borrowing Books with Gin-Gonic HTTP Respons Framework, GORM for Object relation model, PostgreSQL for database.

## 🔗 Description

This Backend Application is used for simple order book, in this application there are two models / ERD Schema likes User, Book, Author & Category.
Also have several features like JWT, Authentification & Authorization.
There are 4 main services :
1. Book Management (Get with paginate, Get Detail, Insert, Update, Delete,
   Search & Borrowing Book)
2. Author Management (Get with paginate, Get Detail, Insert, Update, Delete, Search)
3. Category Management (Get with paginate, Get Detail, Insert, Update, Delete, Search)
4. User Management (Get Login Data, Insert Login Data)

Notes :
1. In this application there are two types of users (Roles). admins and user.
   Admin can do *Costumer Management* but Role Costumer can't, Registration page can only register Costumer roles, Admins can only be registered through seeding data.

<h2 align="center">
 ERD (Entity Relation Database)
</h2>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1720514738/Screenshot_2024-07-09_153934_m0bzlk.png" alt="erd.jpg" /></p>

<h2 align="center">
 Table Specification
</h2>

## Several command you must know in this app :
```bash
   Notes : First of all, Change Directory you want. e.g auth-service, order-service, 
   author-service, category-service
```

```bash
1. go run . serve //to run the app / server
2. go run . migrate -u //for database migration
# or
go run . migrate -d //for rollback
3. go run . seed // to seeding data Role admin if u want Email : "admin@gmail.com" Pass : admin12345678
```

## 🛠️ Installation Steps

1. Clone the repository

```bash
https://github.com/adiet95/book-store.git
```

2. Install dependencies

```bash
go mod tidy
```
> Wait a minute, if still error run

```bash
go mod vendor
```

3. Add Env File

```sh
  DB_USER="YOUR DB USER"
  DB_HOST="YOUR DB HOST"
  DB_NAME="book-store"
  DB_PASS="YOUR DB PASSWORD"
  JWT_KEYS="YOUR KEYS"
  PORT=":8080"
  GRPC_ADDRESS=":8085"
  DB_PORT="5432"
```

4. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration table
# or
go run main.go migrate --down //for rollback the database
```

5. Seeding data admin

```bash
go run . seed
```
_Purpose to login as Admin's Role_
```
{
  "email": "admin@gmail.com",
  "password": "admin12345678"
}
```

6. Run the app

```bash
go run . serve
```

### If you want to build with docker-compose

```bash
docker compose build 
```

### 🚀 You are all set

## 🔗 RESTful endpoints
### POST /register

> Create new user
_Request Header_
```
not needed
```
_Request Body_
```
{
  "user_name": <your username> (STRING),
  "full_name": <your username> (STRING),
  "email": <your email> (STRING),
  "password": <your password> (STRING),
  "address": <your address> (STRING),
  "phone": <your phone> (STRING)
}
```

### POST /login

> Process Login
_Request Header_
```
not needed
```
_Request Body_
```
{
  "email": <your email> (STRING),
  "password": <your password> (STRING)
}
```

### GET /user
> Get Data Costumer
_Request Header_
```
Bearer Token
```
_Request Body_
```
not needed
```
_Request Query Params_
```
limit = (limit for pagination)
offset = (offset for pagination)
```

### POST /user
> Post Data Costumer
_Request Header_
```
Bearer Token
```
_Request Body_
```
{
    "email" : "(STRING)",
    "password" : "(STRING)",
    "address" : "(STRING)",
    "full_name" : "(STRING)",
    "phone" : "(STRING)"
}
```
_Request Query Params_
```
no need
```

### PUT /user
> Update Data Costumer
_Request Header_
```
Bearer Token
```
_Request Body_
```
{
    "email" : "(STRING)",
    "password" : "(STRING)",
    "address" : "(STRING)",
    "full_name" : "(STRING)",
    "phone" : "(STRING)"
}
```
_Request Query Params_
```
no need
```

### DELETE /user
> Delete Data Costumer
_Request Header_
```
Bearer Token
```
_Request Body_
```
no need
```
_Request Query Params_
```
email = (Delete by email)
```

### GET /user/detail

> Get Detail Data Costumer Email

_Request Header_
```
Bearer Token
```

_Request Body_
```
no need
```
_Request Query Params_
```
email = (Get detail data by email)
```

### GET /user/search

> Search Data Costumer by Full_Name

_Request Header_
```
Bearer Token
```

_Request Body_
```
no need
```
_Request Query Params_
```
name = (Search data by full_name)
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [Gin-Gonic](https://gin-gonic.com/): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/adiet-alimudin/)