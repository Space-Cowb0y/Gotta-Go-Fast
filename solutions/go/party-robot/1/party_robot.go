package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
		return "Welcome to my party, " + fmt.Sprintf(name) + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return  "Happy birthday " + fmt.Sprintf(name) + "! You are now " + fmt.Sprintf(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party,%s! You have been assigned to table %d. Your table is on the %s, exactly %.1f meters from here. You will be sitting next to %s.", name, table, direction, distance, neighbor)
}
