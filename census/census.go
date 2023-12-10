// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	res := Resident{
		Name:    name,
		Age:     age,
		Address: address}
	return &res
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {

	if r.Name == "" {
		return false
	}
	// if r.Age == 0 {
	// 	return true
	// }
	streetAddress, streetAddressExists := r.Address["street"]
	if !streetAddressExists {
		return false
	}
	if streetAddress == "" {
		return false
	}
	return true
	// panic("Please implement HasRequiredInfo.")
}

// Delete deletes a resident's information. - SEE the no return needed
func (r *Resident) Delete() {
	r.Address = nil
	r.Age = 0
	r.Name = ""
	// panic("Please implement Delete.")
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {

	var count = 0

	for _, resident := range residents {

		if resident.HasRequiredInfo() {
			count = count + 1
		}
	}

	return count
	// panic("Please implement Count.")
}
