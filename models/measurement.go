package models


type Measurement struct {
	ID     						string 		`json:"id" bson:"_id"`						
	VehicleId  					int 		`json:"vehicleId" bson:"vehicleId"`
	Ts							int	`json:"ts" bson:"ts"`
	Temperature					float64 	`json:"temperature" bson:"temperature"`
	OperatingTime				int 		`json:"operatingtime" bson:"operatingtime"`
	FuelUsage 					float64 	`json:"fuelusage" bson:"fuelusage"`
	FrontLinkagePosition 		int			`json:"front_linkage_position" bson:"front_linkage_position"`
	DrivingSpeed 				int			`json:"drivingspeed" bson:"drivingspeed"`
	EngineState 				int			`json:"enginestate" bson:"enginestate"`
	AutopilotSystemState 		int			`json:"autopilotsystemstate" bson:"autopilotsystemstate"`
	EngineLoad 					float64		`json:"engineload" bson:"engineload"`
	Latitude 					float64		`json:"latitude" bson:"latitude"`
	Longitude 					float64		`json:"longitude" bson:"longitude"`
	Altitude 					float64		`json:"altitude" bson:"altitude"`
    EngineRotation 				float64		`json:"engine_rotation" bson:"engine_rotation"`
    FrontPmeShaft 				float64		`json:"front_pme_shaft" bson:"front_pme_shaft"`
    RearLinkagePosition 		int			`json:"rear_linkage_position" bson:"rear_linkage_position"`
    FourWheelDrivingState 		string 		`json:"four_wheel_driving_state" bson:"four_wheel_driving_state"`
    FuelTankLevel 				int			`json:"fuel_tank_level" bson:"fuel_tank_level"`
    LastErrorMsg 				string		`json:"last_error_msg" bson:"last_error_msg"`
    EngineTemperature 			float64 	`json:"engine_temperature" bson:"engine_temperature"`
    ConnectionState 			string		`json:"connection_state" bson:"connection_state"`
    LteConnectionLevel 			float64		`json:"lte_connection_level" bson:"lte_connection_level"`
    mode 						string		`json:"mode" bson:"mode"`

}