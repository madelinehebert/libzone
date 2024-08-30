package libzone

import "container/list"

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
	Net      *list.List
	AttrList *list.List
}

// Function to create a new zone struct
func (z *Zone) Init(name string) *Zone {
	z.Brand = Brand.Ipkg()
	z.State = State.Incomplete()
	z.ZonePath = ""
	z.AutoBoot = false
	return z
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
