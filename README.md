The JSON field will store and retrieve items in PostgreSQL. For accessing PostgreSQL's JSON store, the normal pq library is very tedious. So, in order to handle that better, I use an Object Relational Mapper (ORM) called GORM.

This ORM has the API for all operations that can be done in the database/sql package.

Install GORM using this command:

```
go get -u github.com/jinzhu/gorms

```

To run the project, Use the following steps:

Clone the project with this command:

```
https://github.com/atanda0x/e-commerce-postgresdb.git

```

change the directory by doing this:

```
cd e-commerce-postgresdb

```

Build the project

```
go build

```

Run the project

```
go run main.go

```

Input the following `curl` on a different terminal

```
curl -X POST \
 http://localhost:9000/v1/user \
 -H 'cache-control: no-cache' \
 -H 'content-type: application/json' \
 -d '{
 "username": "atanda nafiu",
 "email_address": "0x@live.com",
 "first_name": "atanda",
 "last_name": "nafiu"
}'

```

The core motto of this project is to show how JSON can be stored and retrieved out of
PostgreSQL. The special thing here is that I queried on the JSON field instead of the
normal fields in the User table.
