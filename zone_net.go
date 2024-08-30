package libzone

type IpType bool

const (
	Exclusive IpType = true
	Shared    IpType = false
)

type Net struct {
	Address        string
	AllowedAddress string
	MacAddress     string
	DefRouter      string
	Physical       string
	GlobalNic      string
	VlanId         int
}
