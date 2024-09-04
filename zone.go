package libzone

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
	AttrList map[uint]any
	Dataset  map[uint]any
	Device   map[uint]any
	Fs       map[uint]any
	Net      map[uint]any
}

// Function to create a new zone struct
func (z *Zone) Init() *Zone {
	z.Brand = Brand.Ipkg()
	z.State = State.Incomplete()
	z.ZonePath = ""
	z.AutoBoot = false
	return z
}

// Function to return the value of a specified zone's property
func Return(i any, propertyIndex uint, field string) any {
	return i.(map[uint]any)[propertyIndex].(*Property).Value.(map[string]any)[field]
}

// Function to return the value of a specified net property
func (z *Zone) ReturnAttrList(propertyIndex uint, field string) any {
	return z.AttrList[propertyIndex].(*Property).Value.(map[string]any)[field]
}

// Function to return the value of a specified net property
func (z *Zone) ReturnDataset(propertyIndex uint, field string) any {
	return z.Dataset[propertyIndex].(*Property).Value.(map[string]any)[field]
}

// Function to return the value of a specified net property
func (z *Zone) ReturnDevice(propertyIndex uint, field string) any {
	return z.Device[propertyIndex].(*Property).Value.(map[string]any)[field]
}

// Function to return the value of a specified net property
func (z *Zone) ReturnFs(propertyIndex uint, field string) any {
	return z.Fs[propertyIndex].(*Property).Value.(map[string]any)[field]
}

// Function to return the value of a specified net property
func (z *Zone) ReturnNet(propertyIndex uint, field string) any {
	return z.Net[propertyIndex].(*Property).Value.(map[string]any)[field]
}

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
