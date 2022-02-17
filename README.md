# Kinma-Project
*Kinma Project is design for product fundraising platform.*

This branch use **Golang** as backend server, AngularJS as frontend UI

  #
create database migration schema
```
migrate create -ext sql -dir db/migration -seq init_schema
```
- Launch db
 ```
make postgres
```
- migrate the db
```
make migrateup
```
- clean up the migration
```
make migratedown
```
- Launch the API Server
```
make server
```



`docker-compose build` to build db, AngularUI, Golang backend and DB service

`docker-compose up` to launch all service
