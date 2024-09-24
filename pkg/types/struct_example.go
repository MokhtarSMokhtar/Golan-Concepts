package types

import "fmt"

// Wheel represents a car's wheel with a size and type.
type Wheel struct {
	Size int    // Diameter of the wheel in inches.
	Type string // Type of the wheel, e.g., "Alloy", "Steel".
}

// Car represents a car with various attributes.
type Car struct {
	Brand      string // Brand of the car, e.g., "Toyota".
	Model      string // Model of the car, e.g., "Corolla".
	Doors      int    // Number of doors.
	Mileage    int    // Mileage of the car.
	FrontWheel Wheel  // Front wheel details.
	BackWheel  Wheel  // Back wheel details.
}

// NewCar creates and returns a new Car instance.
func NewCar(brand, model string, doors, mileage int, frontWheel, backWheel Wheel) *Car {
	return &Car{
		Brand:      brand,
		Model:      model,
		Doors:      doors,
		Mileage:    mileage,
		FrontWheel: frontWheel,
		BackWheel:  backWheel,
	}
}

// Drive simulates driving the car by increasing its mileage.
func (c *Car) Drive(miles int) {
	if miles < 0 {
		fmt.Println("Cannot drive negative miles.")
		return
	}
	c.Mileage += miles
	fmt.Printf("Driven %d miles. Total mileage is now %d miles.\n", miles, c.Mileage)
}

// Service simulates servicing the car by resetting mileage.
func (c *Car) Service() {
	fmt.Printf("Servicing %s %s. Mileage before service: %d miles.\n", c.Brand, c.Model, c.Mileage)
	c.Mileage = 0
	fmt.Println("Mileage reset to 0 miles after service.")
}

// WheelDetails prints the details of the car's wheels.
func (c Car) WheelDetails() {
	fmt.Printf("Front Wheel: %d-inch %s\n", c.FrontWheel.Size, c.FrontWheel.Type)
	fmt.Printf("Back Wheel: %d-inch %s\n", c.BackWheel.Size, c.BackWheel.Type)
}

// ElectricCar represents an electric car, embedding Car and adding BatteryCapacity.
type ElectricCar struct {
	Car                 // Embedded struct
	BatteryCapacity int // Battery capacity in kWh.
}

// NewElectricCar creates and returns a new ElectricCar instance.
func NewElectricCar(brand, model string, doors, mileage int, frontWheel, backWheel Wheel, batteryCapacity int) *ElectricCar {
	return &ElectricCar{
		Car: Car{
			Brand:      brand,
			Model:      model,
			Doors:      doors,
			Mileage:    mileage,
			FrontWheel: frontWheel,
			BackWheel:  backWheel,
		},
		BatteryCapacity: batteryCapacity,
	}
}

// Charge simulates charging the electric car.
func (e *ElectricCar) Charge() {
	fmt.Printf("Charging %s %s. Battery capacity is %d kWh.\n", e.Brand, e.Model, e.BatteryCapacity)
}

// Greet overrides the embedded Car's Greet method to include electric features.
func (e ElectricCar) Greet() {
	fmt.Printf("Hello, I am a %s %s with a %dkWh battery.\n", e.Brand, e.Model, e.BatteryCapacity)
}

// Example demonstrates the usage of structs and methods.
func Example() {
	// Creating Wheel instances
	frontWheel := Wheel{
		Size: 16,
		Type: "Alloy",
	}
	backWheel := Wheel{
		Size: 16,
		Type: "Alloy",
	}

	// Create Car instances
	car1 := NewCar("Toyota", "Corolla", 4, 15000, frontWheel, backWheel)
	fmt.Println("Car 1 Details:")
	fmt.Printf("Brand: %s, Model: %s, Doors: %d, Mileage: %d\n", car1.Brand, car1.Model, car1.Doors, car1.Mileage)
	car1.WheelDetails()
	car1.Drive(500)
	car1.Service()
	fmt.Println()

	// Create ElectricCar instances
	electricCar := NewElectricCar("Tesla", "Model S", 4, 20000, frontWheel, backWheel, 100)
	fmt.Println("Electric Car Details:")
	fmt.Printf("Brand: %s, Model: %s, Doors: %d, Mileage: %d, Battery Capacity: %d kWh\n", electricCar.Brand, electricCar.Model, electricCar.Doors, electricCar.Mileage, electricCar.BatteryCapacity)
	electricCar.WheelDetails()
	electricCar.Greet()
	electricCar.Drive(300)
	electricCar.Charge()
}
