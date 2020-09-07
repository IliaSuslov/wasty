// Wasty
use wasty;

// create test drivers
db.createCollection("drivers")

db.drivers.insertMany([
    {
        name: "Водитель 1",
        phone: "01",
        email: "email@mail.ru"
    },
    {
        name: "Водитель 2",
        phone: "01",
        email: "email@mail.ru"
    },
])

// create test addresses
db.createCollection("addresses")
db.addresses.insertMany([
    {
        name: "удальцово",
        latitude: 0,
        longitude: 0,
        altitude: 0
    },
    {
        name: "сосново",
        latitude: 0,
        longitude: 0,
        altitude: 0
    }
])

// create test cars
db.createCollection("cars")
db.cars.insertMany([
    {
        name: "уаз",
	    number:  "0001ад47",
	    trailer: "0002ад47"
    },
    {
        name: "газ69",
	    number:  "0003ад47",
	    trailer: "0004ад47"
    },
    {
        name: "газ53",
	    number:  "0005ад47"
    },
])