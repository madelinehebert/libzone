package libzone

import "container/list"

// Define exported Property struct
type Property struct {
	Value any
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
