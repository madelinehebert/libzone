package libzone

import (
	"container/list"
	"errors"
	"reflect"
)

// Define exported Property struct
type Property struct {
	Value any
}

// Configure a property - key is the field name, val is the value
func (p *Property) Set(key string, val string) error {
	//Loop over property keys and values
	for k, v := range p.Value.(map[string]any) {
		//If we have a string
		if reflect.TypeOf(v) == reflect.TypeOf(key) && k == key {
			p.Value.(map[string]any)[k] = val
			return nil
		} else if reflect.TypeOf(v) == reflect.TypeOf(make(map[string]string)) {
			return nil
		}
	}

	return errors.New("")
}

// Define new Fs Property
func (p *Property) AttrList() *Property {
	v := make(map[string]any, 3)
	v["name"] = ""
	v["type"] = ""
	v["value"] = nil
	p.Value = v
	return p
}

// Define new Fs Property
func (p *Property) Dataset() *Property {
	v := make(map[string]any, 1)
	v["name"] = ""
	p.Value = v
	return p
}

// Define new Fs Property
func (p *Property) Device() *Property {
	v := make(map[string]any, 2)
	v["match"] = ""
	v["property"] = make(map[string]string)
	p.Value = v
	return p
}

// Define new Fs Property
func (p *Property) Fs() *Property {
	v := make(map[string]any, 5)
	v["dir"] = ""
	v["special"] = ""
	v["raw"] = ""
	v["type"] = ""
	v["options"] = &list.List{}
	p.Value = v
	return p
}

// Define new Net Property
func (p *Property) Net() *Property {
	v := make(map[string]any, 7)
	v["address"] = ""
	v["allowedAddress"] = ""
	v["macAddress"] = ""
	v["defRouter"] = ""
	v["physical"] = ""
	v["globalNic"] = ""
	v["vlanId"] = -1
	p.Value = v
	return p
}
