package libzone

import (
	"container/list"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
	AttrList map[int]any
	Dataset  map[int]any
	Device   map[int]any
	Fs       map[int]any
	Net      map[int]any
}

// Function to create a new zone struct
func (z *Zone) Init(name string) *Zone {
	z.Name = name
	z.Brand = Brand.Ipkg()
	z.State = State.Incomplete()
	z.ZonePath = ""
	z.AutoBoot = false
	z.AttrList = make(map[int]any)
	z.Dataset = make(map[int]any)
	z.Device = make(map[int]any)
	z.Fs = make(map[int]any)
	z.Net = make(map[int]any)
	return z
}

// Function to verify a zone's creation state - returns an error if any condition is violated.
// This fucntion starts by confirming the most bare bones configuration for a zone has been provided.
// Then, it moves on to confirm each map of properties.
func (z *Zone) Verify() error {
	//Start checking the individual fields of the struct

	//Zone name checks
	if len(z.Name) < 1 { //If zone name is too short
		return errors.New("err: zone name must be at least one character long.")
	} else if len(z.Name) > 63 { //If zone name is too long
		return errors.New("err: zone name must be less than 64 characters.")
	} else if ok := regexp.MustCompile(`^[a-zA-Z0-9_.-]*$`).MatchString(z.Name); !ok { //If zone name does not contain legal characters
		return errors.New("err: zone contains an illegal character. Please only use the following selection of characters for a zone name: 'A-Za-z0-9_.-'")
	} else if z.Name[0:4] == "SUNW" { //If zone name begins with SUNW
		return errors.New("err: zone name must not begin with the letter sequence: 'SUNW'.")
	} else if z.Name == "global" { //If zone name is 'global'
		return errors.New("err: zone name must not contain: 'global'.")
	} else {
		//Ensure zone name is not in use
		if err := IsZoneNameInUse(z.Name); err != nil {
			return err
		}
	}

	//Zone brand check
	if strings.Contains(z.Brand.String(), "Unknown") {
		return errors.New("err: invalid brand with integer id: " + z.Brand.String())
	}

	//Zone state check
	if z.State != State.Incomplete() && z.State != State.Configured() {
		return errors.New("err: zone is not in valid state to require verification: " + z.State.String())
	}

	//Set state to configured and return nil
	z.State = State.Configured()
	return nil
}

// Function to print out the zone's information
func (z *Zone) Info() {
	//Loop over zone and print data in the same format as 'zonecfg -z *Zone export'
	zStructValue := reflect.ValueOf(*z)
	zStructType := reflect.TypeOf(*z)
	for i := 0; i < zStructValue.NumField(); i++ {
		//Switch on type
		switch reflect.TypeOf(zStructValue.Field(i).Interface()) {

		//If we have a string
		case reflect.TypeOf(""):
			if len(zStructValue.Field(i).Interface().(string)) != 0 {
				if zStructType.Field(i).Name == "Name" {
					fmt.Printf("%20s: %-20s\n", strings.ToLower(zStructType.Field(i).Name), zStructValue.Field(i).Interface().(string))
				} else {
					//fmt.Printf("%20s=%-20s\n", "set "+strings.ToLower(zStructType.Field(i).Name), zStructValue.Field(i).Interface().(string))
					fmt.Printf("%20s: %-20s\n", strings.ToLower(zStructType.Field(i).Name), zStructValue.Field(i).Interface().(string))
				}

			} else {
				//fmt.Printf("%20s=%-20s\n", "set "+strings.ToLower(zStructType.Field(i).Name), "<nil>")
				fmt.Printf("%20s: %-20s\n", strings.ToLower(zStructType.Field(i).Name), "<nil>")
			}

		//If we have a bool
		case reflect.TypeOf(true):
			fmt.Printf("%20s: %-20v\n", zStructType.Field(i).Name, zStructValue.Field(i).Interface().(bool))

		//If we have a map[int]any
		case reflect.TypeOf(make(map[int]any)):

			//Ensure we aren't trying to print an empty value
			if len(zStructValue.Field(i).Interface().(map[int]any)) != 0 {
				//Loop over values
				for key, _ := range zStructValue.Field(i).Interface().(map[int]any) {

					//Print item header with number
					fmt.Printf("%20s: %-20s\n", zStructType.Field(i).Name+" #"+strconv.Itoa(int(key)+1), "")

					//Loop over sub values
					for subKey, subVal := range zStructValue.Field(i).Interface().(map[int]any)[key].(*Property).Value.(map[string]any) {
						//Check type
						switch reflect.TypeOf(subVal) {
						//string
						case reflect.TypeOf(""):
							if len(subVal.(string)) != 0 {
								fmt.Printf("%20s: %-20s\n", " - "+subKey, subVal.(string))
							} else {
								fmt.Printf("%20s: %-20s\n", " - "+subKey, "<nil>")
							}

						//map[string]string
						case reflect.TypeOf(make(map[string]string)):
							fmt.Printf("%20s: %-20s\n", " - "+subKey, "NOT YET IMPLEMENTED")
						//bool
						case reflect.TypeOf(true):
							fmt.Printf("%20s: %-20v\n", " - "+subKey, subVal.(bool))
						//int
						case reflect.TypeOf(1):
							fmt.Printf("%20s: %-20v\n", " - "+subKey, subVal.(int))
						//list
						case reflect.TypeOf(&list.List{}):
							fmt.Printf("%20s: %-20v\n", " - "+subKey, "NOT YET IMPLEMENTED")
						//Unhandled types
						default:
							fmt.Printf("%20s: %-20v\n", " - "+subKey, "NOT YET IMPLEMENTED")
						}
					}
				}

			} else {
				fmt.Printf("%20s: %-20s\n", zStructType.Field(i).Name, "<nil>")
			}
		}

	}
}

// Function to add a property type to a zone - returns an int of the map key it sits at, and error if an invalid property is entered
// Errors are also returned if 256 places in the map are occupied
func (z *Zone) Add(p string) (int, error) {
	p = strings.ToLower(p)
	switch p {
	case "attrlist":
		index := len(z.AttrList)
		//Cap number of items to 256
		if index+1 > 255 {
			return -2, errors.New("attrlist limit of 256 has been reached")
		} else {
			z.AttrList[index] = (&Property{}).AttrList()
			return index, nil
		}

	case "dataset":
		index := len(z.Dataset)
		//Cap number of items to 256
		if index+1 > 255 {
			return -2, errors.New("dataset limit of 256 has been reached")
		} else {
			z.Dataset[index] = (&Property{}).Dataset()
			return index, nil
		}
	case "device":
		index := len(z.Device)
		//Cap number of items to 256
		if index+1 > 255 {
			return -2, errors.New("device limit of 256 has been reached")
		} else {
			z.Device[index] = (&Property{}).Device()
			return index, nil
		}
	case "fs":
		index := len(z.Fs)
		//Cap number of items to 256
		if index+1 > 255 {
			return -2, errors.New("fs limit of 256 has been reached")
		} else {
			z.Fs[index] = (&Property{}).Fs()
			return index, nil
		}
	case "net":
		index := len(z.Net)
		//Cap number of items to 256
		if index+1 > 255 {
			return -2, errors.New("net limit of 256 has been reached")
		} else {
			z.Net[index] = (&Property{}).Net()
			return index, nil
		}
	default:
		return -1, errors.New("Unknown property: " + p)
	}

}

// Function to configure a device - The 'item' is the property name, the 'id' is the index of that property type,
// the 'key' is the field, and the 'val' is the value of the field. Invalid indices return errors.
func (z *Zone) Configure(item string, id int, key string, val string) error {
	//Limit number of items to 256 - then fork to configuring if index is within range
	if id < 0 || id > 255 {
		return errors.New("Invalid index: " + strconv.Itoa(id))
	} else {
		item = strings.ToLower(item)
		switch item {
		//Configure Attr List
		case "attrlist":
			//Check length of map
			if len(z.AttrList) >= int(id)+1 {
				if err := z.AttrList[id].(*Property).Set(key, val); err != nil {
					return err
				} else {
					return nil
				}
			} else {
				return errors.New("invalid index")
			}

		//Configure Dataset
		case "dataset":
			//Check length of map
			if len(z.Dataset) >= int(id)+1 {
				if err := z.Dataset[id].(*Property).Set(key, val); err != nil {
					return err
				} else {
					return nil
				}
			} else {
				return errors.New("invalid index")
			}
		//Configure Device
		case "device":
			//Check length of map
			if len(z.Device) >= int(id)+1 {
				if err := z.Device[id].(*Property).Set(key, val); err != nil {
					return err
				} else {
					return nil
				}
			} else {
				return errors.New("invalid index")
			}
		//Configure Fs
		case "fs":
			//Check length of map
			if len(z.Fs) >= int(id)+1 {
				if err := z.Fs[id].(*Property).Set(key, val); err != nil {
					return err
				} else {
					return nil
				}
			} else {
				return errors.New("invalid index")
			}
		//Configure Net
		case "net":
			//Check length of map
			if len(z.Net) >= int(id)+1 {
				if err := z.Net[id].(*Property).Set(key, val); err != nil {
					return err
				} else {
					return nil
				}
			} else {
				return errors.New("invalid index")
			}
		//Handle unknown property
		default:
			return errors.New("unknown property: " + item)
		}
	}

}

// Function to configure a device

/*
func (z *Zone) ConfigureDevice(id int, key string, val string) error {
	if err := z.Device[id].(*Property).Set(key, val); err != nil {
		return err
	} else {
		return nil
	}
}
*/

// Function to return the value of a specified net property
func (z *Zone) ReturnAttrList(propertyIndex int, field string) any {
	if z.AttrList != nil {
		return z.AttrList[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnDataset(propertyIndex int, field string) any {
	if z.Dataset != nil {
		return z.Dataset[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnDevice(propertyIndex int, field string) any {
	if z.Device != nil {
		return z.Device[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnFs(propertyIndex int, field string) any {
	if z.Net != nil {
		return z.Fs[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}

// Function to return the value of a specified net property
func (z *Zone) ReturnNet(propertyIndex int, field string) any {
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

// Check if a zone name is in use
func IsZoneNameInUse(name string) error {
	//Check the zone name
	if false {
		return errors.New("zone name in use")
	} else {
		return nil
	}
}

//func (z *Zone) x() error    {return nil}
