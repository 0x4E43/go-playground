package db

type NodeMaster struct {
	Id              int16
	DeviceId        int16
	ControllerId    int16
	IEEAddress      string
	Intensity       int16
	Occupancy       int16
	IsIntensitySent int8
	IsOccupancySent int8
}

type Student struct {
	Id     int16
	Name   string
	RollNo int16
}
