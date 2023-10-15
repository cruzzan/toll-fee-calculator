package main

type Vehicle interface {
	GetType() string
}

func isTollFreeVehicle(vehicle Vehicle) bool {
	vehicleType := vehicle.GetType()

	return vehicleType == "Motorbike" ||
		vehicleType == "Tractor" ||
		vehicleType == "Emergency" ||
		vehicleType == "Diplomat" ||
		vehicleType == "Foreign" ||
		vehicleType == "Military"
}

type Car struct {
}

func (c Car) GetType() string {
	return "Car"
}

type Motorbike struct {
}

func (m Motorbike) GetType() string {
	return "Motorbike"
}
