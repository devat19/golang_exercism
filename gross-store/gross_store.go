package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMap := map[string]int{}

	unitMap["quarter_of_a_dozen"] = 3
	unitMap["half_of_a_dozen"] = 6
	unitMap["dozen"] = 12
	unitMap["small_gross"] = 120
	unitMap["gross"] = 144
	unitMap["great_gross"] = 1728

	return unitMap

	// panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {

	return map[string]int{}
	// panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill. - Test this with pointers! return bool while updating the actual object only
func AddItem(bill, units map[string]int, item, unit string) bool {

	unitVal, unitExists := units[unit]

	if !unitExists { // invalid unit
		return false
	}

	itemQty, itemExistsInBill := bill[item]

	if !itemExistsInBill {
		bill[item] = unitVal
	} else {
		bill[item] = itemQty + unitVal
	}

	return true
	// panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	unitVal, unitExists := units[unit]

	if !unitExists { // invalid unit
		return false
	}

	itemQty, itemExistsInBill := bill[item]

	if !itemExistsInBill {
		return false
	} else {

		newItemQty := itemQty - unitVal

		if newItemQty < 0 {
			return false
		} else if newItemQty == 0 {
			delete(bill, item)
			return true
		}

		bill[item] = newItemQty
	}

	return true

	// panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	itemQty, itemExistsInBill := bill[item]

	if !itemExistsInBill {
		return 0, false
	}
	return itemQty, true
	// panic("Please implement the GetItem() function")
}
