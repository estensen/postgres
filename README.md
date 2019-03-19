# postgres

## Create and populate db
```sh
$ createdb pqgotest
$ psql -d pqgotest -a -f populate_db.sql
```

Test that db was populated
```sh
$ psql pqgotest
pqgotest=# SELECT * FROM places; 
   name   |     location
----------+------------------
 Snøhetta | (62.3198,9.2677)
(1 row) 
```

## Run
```sh
$ go run main.go
```

