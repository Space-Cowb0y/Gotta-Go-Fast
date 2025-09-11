package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + fmt.Sprint(name) + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + fmt.Sprint(name) + "! You are now " + fmt.Sprint(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %d. Your table is on the %s, exactly %0.1f meters from here.\nYou will be sitting next to %s", name, table, direction, distance, neighbor)
}
