package globalgame

var PlateOrder = []string{"first-plate", "second-plate", "third-plate", "fourth-plate", "fifth-plate", "sixth-plate"}

func IsPlate(itemName string) bool {
	for _, plate := range PlateOrder {
		if itemName == plate {
			return true
		}
	}
	return false
}
