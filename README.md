# wasty

## Mongo

https://hub.docker.com/_/mongo

env MONGO_INITDB_ROOT_USERNAME=root MONGO_INITDB_ROOT_PASSWORD=aiweeTh1; \
    docker run --name mongodb -p 127.0.0.1:27017:27017 -d mongo:bionic

## wasty

make serve

## Process

- create poi
- create []images

## Waybill
```yml
Waybill
    Data: 01/09/2020
    Name: "1"
    Desc: "Вывоз"
    Car:
        Name: газ 53
        Number: 001 ад 47
        Trailer: хз
    Driver:
        Name:
        Phone:
        Email:
    Firm:
        Name:
        Contact:
            Name:
            Phone:
            Email:
    Jobs:
        - Desc: Удальцово
        POI:
            Name:
            Latitude:
            Latitude:
            Altitude:
    Executions:
        - Date 01/09/2020 10:00
        Desc: ""
        POI:
            Latitude:
            Latitude:
            Altitude:
        Images:
            - Name:
            ContentType:
            Content
            
```    
    