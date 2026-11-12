package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery int 
	batteryDrain int
	speed int
	distance int
}

// NewCar creates a new remote controlled Car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed: speed,
		batteryDrain: batteryDrain,
		battery: 100,
		distance: 0,
	}
}

// TODO: define the 'Track' type struct
type Track struct{
	distance int
}

// NewTrack creates a new Track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,	
	}
}

// Drive drives the Car one time. If there is not enough battery to drive one more time,
// the Car will not move.
func Drive(Car Car) Car {
	if Car.battery >= Car.batteryDrain {
		Car.distance += Car.speed
		Car.battery -= Car.batteryDrain
	}
	return Car 
}

// CanFinish checks if a Car is able to finish a certain Track.
func CanFinish(Car Car, Track Track) bool {
	maxDrives := Car.battery / Car.batteryDrain
	maxDistance := maxDrives * Car.speed

	return maxDistance >= Track.distance
}
