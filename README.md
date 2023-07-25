The JSON field will store and retrieve items in PostgreSQL. For accessing PostgreSQL's JSON store, the normal pq library is very tedious. So, in order to handle that better, I use an Object Relational Mapper (ORM) called GORM.

This ORM has the API for all operations that can be done in the database/sql package.

Install GORM using this command:

`go get -u github.com/jinzhu/gorms`