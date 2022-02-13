# Kinma-Project
*Kinma Project is design for product fundraising platform.*

`docker pull postgres:14-alpine`

`brew install golang-migrate`

- Launch db
 ```
docker run --name kinma-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine
```
- Access postgres container
```
docker exec -it kinma-postgres psql -U root
```


`docker-compose build` to build db, AngularUI, Django Backend Server images

`docker-compose up` to launch all service
