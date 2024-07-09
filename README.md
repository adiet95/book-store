# Go-Order-API
Simple app product order with Gin-Gonic HTTP Respons Framework, GORM for Object relation model, PostgreSQL for database.

## 🔗 Description

This Backend Application is used for simple order product, in this application there are two models / ERD Schema likes User / Costumer & Products.
Also have several features like JWT, Authentification & Authorization.
There are 3 main modules :
1. Customer Management (Get with paginate, Get Detail, Insert, Update, Delete,
   Search)
2. Order Management (Get with paginate, Get Detail, Insert, Update, Delete,
   Search)
3. Authentikasi Management (Get Login Data, Insert Login Data)

Notes :
1. I'am using UUID for user_id, don't forget to create extenxion in SQL console after create the database with this query below :
```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

2. In this application there are two types of users (Roles). admins and costumer.
   Admin can do *Costumer Management* but Role Costumer can't, Registration page can only register Costumer roles, Admins can only be registered through seeding data.

<h2 align="center">
 ERD (Entity Relation Database)
</h2>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665874871/erd_c15gne.png" alt="erd.jpg" /></p>

<h2 align="center">
 Table Specification
</h2>

<h3 align="center">Costumer's Table</h3>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665882605/table-cost_miwqjk.png" alt="cost.jpg" /></p>
<h3 align="center">Order's Table</h3>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665882605/table-user_sz493j.png" alt="order.jpg" /></p>

## Several command you must know in this app :
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
https://github.com/adiet95/go-order-api.git
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
  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASS = Your DB Password
  JWT_KEYS = Your JWT Key
  PORT = Your Port
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

### GET /order

> Get Data Order

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

### POST /order

> Post Data Order

_Request Header_
```
Bearer Token
```

_Request Body_
```
{
    "order_name" : "Test",
    "invoice" : "test inv",
    "address" : "test address",
    "telphone" : "0813",
    "amount" : 2,
    "price" : 10000,
    "status" : "paid"
}
```
_Request Query Params_
```
no need
```

### PUT /order

> Update Data Order

_Request Header_
```
Bearer Token
```

_Request Body_
```
{
    "order_name" : "Test",
    "invoice" : "test inv",
    "address" : "test address",
    "telphone" : "0813",
    "amount" : 2,
    "price" : 10000,
    "status" : "paid"
}
```
_Request Query Params_
```
id = (id you want to update)
```

### DELETE /order

> Delete Data Order

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
id = (Delete by id)
```

### GET /order/detail

> Get Data Detail by ID

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
id = (Get detail data by id)
```

### GET /order/search

> Search Data Order by Name

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
name = (Search data by name)
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [Gin-Gonic](https://gin-gonic.com/): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/achmad-shiddiq-alimudin/)