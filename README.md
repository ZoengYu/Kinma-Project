# Kinma-Project
*Kinma Project is design for product fundraising platform.*

This branch use **Golang** as backend server, AngularJS as frontend UI

  #
  ```
  git clone https://github.com/ZoengYu/Kinma-Project.git
  cd Kinma-Project 
  git checkout kinma-golangBackend
  ```

`docker-compose build` to build postgres db, AngularUI, Golang backend api service

`docker-compose up` to launch all service
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
