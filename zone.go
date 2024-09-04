package libzone

import (
	"log"
)

// Set default zone
const defaultZonePath string = "/zones"

// Zone struct
type Zone struct {
	Name     string
	Brand    brand
	State    state
	ZonePath string
	AutoBoot bool
	IpType   IpType
	AttrList map[byte]any
	Dataset  map[byte]any
	Device   map[byte]any
	Fs       map[byte]any
	Net      map[byte]any
}

// Function to return the value of a specified zone's property
func Return(i any, propertyIndex byte, field string) any {
	// Defer panic
	defer func(any) {
		recover()
	}(nil)

	if i.(map[byte]any) != nil {
		return i.(map[byte]any)[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to create a new zone struct
func (z *Zone) Init() *Zone {
	z.Brand = Brand.Ipkg()
	z.State = State.Incomplete()
	z.ZonePath = ""
	z.AutoBoot = false
	z.AttrList = make(map[byte]any)
	z.Dataset = make(map[byte]any)
	z.Device = make(map[byte]any)
	z.Fs = make(map[byte]any)
	z.Net = make(map[byte]any)
	return z
}

// Function to add a property type to a zone
func (z *Zone) Add(p string) *Zone {
	switch p {
	case "attrList":
		z.AttrList[byte(len(z.AttrList))] = (&Property{}).Fs()
	case "dataset":
		z.Dataset[byte(len(z.Dataset))] = (&Property{}).Fs()
	case "device":
		z.Device[byte(len(z.Device))] = (&Property{}).Fs()
	case "fs":
		z.Fs[byte(len(z.Fs))] = (&Property{}).Fs()
	case "net":
		z.Net[byte(len(z.Net))] = (&Property{}).Net()
	default:
		log.Printf("Unknown property: '%s'\n", p)
	}

	return z
}

// Function to return the value of a specified net property
func (z *Zone) ReturnAttrList(propertyIndex byte, field string) any {
	if z.AttrList != nil {
		return z.AttrList[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnDataset(propertyIndex byte, field string) any {
	if z.Dataset != nil {
		return z.Dataset[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnDevice(propertyIndex byte, field string) any {
	if z.Device != nil {
		return z.Device[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnFs(propertyIndex byte, field string) any {
	if z.Net != nil {
		return z.Fs[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnNet(propertyIndex byte, field string) any {
	if z.Net != nil {
		return z.Net[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}

}

// Function to write configuration to disk
func (z *Zone) Write() error { return nil }

// Function to boot a zone
func (z *Zone) Boot() error { return nil }

// Function to halt a zone
func (z *Zone) Halt() error { return nil }

// Function to ready a zone
func (z *Zone) Ready() error { return nil }

// Function to shutdown a zone
func (z *Zone) Shutdown() error { return nil }

// Function to reboot a zone
func (z *Zone) Reboot() error { return nil }

// Function to verify a zone's configuration
func (z *Zone) Verify() error { return nil }

// Function to install a zone
func (z *Zone) Install() error { return nil }

// Function to uninstall a zone
func (z *Zone) Uninstall() error { return nil }

// Function to move a zone to a new zonepath
func (z *Zone) Move(zonePath string) error { return nil }

// Function to attach a zone
func (z *Zone) Attach() error { return nil }

// Function to detach a zone
func (z *Zone) Detach() error { return nil }

//func (z *Zone) x() error    {return nil}
