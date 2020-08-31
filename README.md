# wasty

## Mongo

https://hub.docker.com/_/mongo

env MONGO_INITDB_ROOT_USERNAME=root MONGO_INITDB_ROOT_PASSWORD=aiweeTh1; \
    docker run --name mongodb -p 127.0.0.1:27017:27017 -d mongo:bionic

## wasty

make serve