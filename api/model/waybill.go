package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Waybill struct{
	ID primitive.ObjectID
	Date time.Time
	Car Car
	Driver Driver
	Firm Firm
	Desc string
	Jobs []Job
	Executions []Execution
}



type Execution struct{
	Date time.Time
	Desc string
	POI POI
	Image Image
}

type Job struct{
	Desc string
	POI POI
}




