package libzone

import (
	"container/list"
	"errors"
	"fmt"
	"reflect"
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
	AttrList map[byte]any
	Dataset  map[byte]any
	Device   map[byte]any
	Fs       map[byte]any
	Net      map[byte]any
}

// Function to create a new zone struct
func (z *Zone) Init(name string) *Zone {
	z.Name = name
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

		//If we have a map[byte]any
		case reflect.TypeOf(make(map[byte]any)):

			//Ensure we aren't trying to print an empty value
			if len(zStructValue.Field(i).Interface().(map[byte]any)) != 0 {
				//Loop over values
				for key, _ := range zStructValue.Field(i).Interface().(map[byte]any) {

					//Print item header with number
					fmt.Printf("%20s: %-20s\n", zStructType.Field(i).Name+" #"+strconv.Itoa(int(key)+1), "")

					//Loop over sub values
					for subKey, subVal := range zStructValue.Field(i).Interface().(map[byte]any)[key].(*Property).Value.(map[string]any) {
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

// Function to add a property type to a zone
func (z *Zone) Add(p string) error {
	p = strings.ToLower(p)
	switch p {
	case "attrlist":
		z.AttrList[byte(len(z.AttrList))] = (&Property{}).AttrList()
	case "dataset":
		z.Dataset[byte(len(z.Dataset))] = (&Property{}).Dataset()
	case "device":
		z.Device[byte(len(z.Device))] = (&Property{}).Device()
	case "fs":
		z.Fs[byte(len(z.Fs))] = (&Property{}).Fs()
	case "net":
		z.Net[byte(len(z.Net))] = (&Property{}).Net()
	default:
		return errors.New("Unknown property: " + p)
	}

	return nil
}

// Function to configure a device - The 'item' is the property name, the 'id' is the index of that property type,
// the 'key' is the field, and the 'val' is the value of the field. Invalid indices return errors.
func (z *Zone) Configure(item string, id byte, key string, val string) error {
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

// Function to configure a device
func (z *Zone) ConfigureDevice(id byte, key string, val string) error {
	if err := z.Device[id].(*Property).Set(key, val); err != nil {
		return err
	} else {
		return nil
	}
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
