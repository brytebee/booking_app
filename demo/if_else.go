package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	if role, today := Guest, Tuesday; role <= Manager {
		accessGranted()
		}else if (today >= Friday) && role == Contractor {
			accessGranted()
			}else if role == Member && today < Saturday {
				accessGranted()
				}else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
		}else {	
		accessDenied()
	}
}
