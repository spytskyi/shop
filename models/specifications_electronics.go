package models

type Specifications struct{}

type SpecificationsLaptop struct {
	ScreenDiagonal       float32
	ScreenType           string
	ScreenRefreshRate    string
	CPU                  string
	OperatingSystem      string
	RamSize              int
	StorageVolume        int
	TypeStorage          string
	DiscreteGraphicsCard string
	IntegratedVideoCard  string
	Colour               string
	Weight               float32
	BodyMaterial         string
	BatteryCapacity      float32
	Width                int
	Height               int
	Depth                int
}

type SpecificationsMobilePhone struct {
	ScreenDiagonal  float32
	ScreenType      string
	CPU             string
	OperatingSystem string
	RamSize         int
	StorageVolume   int
	Graphics        string
	Colour          string
	Weight          float32
	BodyMaterial    string
	BatteryCapacity float32
	Width           float32
	Height          float32
	Depth           float32
}
