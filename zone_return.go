package libzone

// Function to return the value of a specified zone's property
func Return(i any, propertyIndex int, field string) any {
	// Defer panic
	defer func(any) {
		recover()
	}(nil)

	if i.(map[int]any) != nil {
		return i.(map[int]any)[propertyIndex].(*Property).Value.(map[string]any)[field]
	} else {
		return nil
	}
}
