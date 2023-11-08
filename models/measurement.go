package models

type Measurement struct {
	ID     string `json:"id" bson:"_id"`
	VehicleId  int `json:"vehicleId" bson:"vehicleId"`
}