# Kinma-Project
*Kinma Project is design for product fundraising platform.*

This branch use **Golang** as backend server


`docker pull postgres:14-alpine`

`brew install golang-migrate`
> migrate create -ext sql -dir db/migration -seq init_schema
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


`docker-compose build` to build db, AngularUI, Golang backend and DB service

`docker-compose up` to launch all service
