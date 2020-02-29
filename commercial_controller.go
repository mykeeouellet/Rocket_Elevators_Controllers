package main

import (
	"fmt"
	"math"
	"sort"
)

// ============= Elevator Simulator ==============  //
//      This program controls elevators in          //
//      a corporate building. Use main() to         //
//     execute it. Select manual settings or        //
//     predetermined scenarios. You can directly    //
//    interact with the program using two methods   //
//                                                  //
//   When you're at RC and want to request a floor  //
//           b1.assignElevator(#floor)              //
//                                                  //
//   When you're at a floor and want to request RC  //
//           b1.requestElevator(#floor)				//
//                                                  //
// ================================================ //

func main() {
	b1 := batteryInit()
	// ========= Manual settings ========== //
	//        uncomment to activate         //
	// ==================================== //
	// b1.manualSettings()
	// ========= Scenario player ========== //
	//     Select one scenario at a time    //
	//        uncomment to activate :)      //
	// ==================================== //
	// b1.scenario1()
	b1.scenario2()
	// b1.scenario3()
	// b1.scenario4()
	// ==================================== //
	// b1.requestElevator(22)
	// b1.assignElevator(26)
	// ==================================== //
}

// ====== this function is used to manually set parameters ======//
func (b *battery) manualSettings() {
	fmt.Println(" * Manual settings activated * ")
	// =================================================== //
	//        Column 1 ([0]) operating the basements
	//        Column 2 ([1]) operating floor #1 to #20
	//        Column 3 ([2]) operating floor #21 to #40
	//        Column 4 ([3]) operating floor #41 to #60
	// =================================================== //
	// Elevator's currentFloor settings :
	b.columnList[1].elevatorsList[0].currentFloor = 2
	b.columnList[1].elevatorsList[1].currentFloor = 0
	b.columnList[1].elevatorsList[2].currentFloor = 10
	b.columnList[1].elevatorsList[3].currentFloor = 4
	b.columnList[1].elevatorsList[4].currentFloor = 19
	// =================================================== //
	// Elevator's direction settings ("up", "idle" or "down") :
	b.columnList[1].elevatorsList[0].direction = "idle"
	b.columnList[1].elevatorsList[1].direction = "idle"
	b.columnList[1].elevatorsList[2].direction = "idle"
	b.columnList[1].elevatorsList[3].direction = "idle"
	b.columnList[1].elevatorsList[4].direction = "idle"
	// =================================================== //
}

// Battery struct => contains columns //
type battery struct {
	id            int
	mainFloor     int
	totalFloor    int
	totalBasement int
	totalColumn   int
	totalElevator int
	columnList    []column
	floorsList    []int
}

// Column struct => contains elevators //
type column struct {
	id                int
	numberOfElevators int
	elevatorsList     []elevator
	minOperatingFloor int
	maxOperatingFloor int
}

// Elevator struct => contains all the parameters of each elevators //
type elevator struct {
	id           int
	value        int
	currentFloor int
	totalScore   int
	direction    string
	doorStatus   string
}

// User struct -=> contains parameters for the user as position & direction //
type user struct {
	direction string
	floor     int
}

// ============== this method initialize a battery with all its attributes =========== //
func batteryInit() battery {
	b1 := battery{

		id:            1,
		mainFloor:     0,
		totalFloor:    66,
		totalBasement: 6,
		totalColumn:   4,
		totalElevator: 20}

	// Calculating the number of elevators per column //
	numElevators := b1.totalElevator / b1.totalColumn
	elevatorsList := []elevator{}

	// For the given number of elevator per column, generates an elevatorsList //
	for j := 1; j <= numElevators; j++ {
		aElevator := elevator{

			id:           j,
			value:        0,
			currentFloor: 0,
			totalScore:   0,
			direction:    "idle",
			doorStatus:   "closed",
		}
		elevatorsList = append(elevatorsList, aElevator)
	}

	// For the given number column, generates a columnList //
	for i := 1; i <= b1.totalColumn; i++ {
		// Formula to determine max operating floor && min operating floor //
		numFloorPerColumn := (b1.totalFloor - b1.totalBasement) / (b1.totalColumn - 1)
		maxOperatingFloor := ((i - 1) * numFloorPerColumn)
		minOperatingFloor := maxOperatingFloor - (numFloorPerColumn - 1)

		// Creating the column that operates the basement //
		if i == 1 {
			aColumn := column{
				id:                i,
				numberOfElevators: numElevators,
				elevatorsList:     elevatorsList,
				minOperatingFloor: b1.mainFloor - b1.totalBasement,
				maxOperatingFloor: b1.mainFloor,
			}
			b1.columnList = append(b1.columnList, aColumn)
		}
		// Creating the columns that operates the upper floors //
		if i > 1 {
			aColumn := column{
				id:                i,
				numberOfElevators: numElevators,
				elevatorsList:     elevatorsList,
				minOperatingFloor: minOperatingFloor,
				maxOperatingFloor: maxOperatingFloor,
			}
			b1.columnList = append(b1.columnList, aColumn)
		}
	}
	return b1
}

// ============== This method is used when the user *from RC* requests another floor ========//
func (b *battery) assignElevator(requestedFloor int) user {
	// Setting user position and direction //
	var aUser user
	aUser.floor = 0
	if aUser.floor > requestedFloor {
		aUser.direction = "down"

	} else if aUser.floor < requestedFloor {
		aUser.direction = "up"
	}

	// Printing lines to make to program clearer //
	fmt.Println("======== Requesting a floor ==========")
	fmt.Println("A request has been made to go to", requestedFloor, ". User is going", aUser.direction)
	fmt.Println("")
	fmt.Println("======== Returning a column ============")

	// Nesting multiple methods. FindColumn, FindElevator and MoveElevator //
	selectedColumn := b.findColumn(requestedFloor)
	selectedElevator := selectedColumn.findElevator(selectedColumn, aUser.floor, requestedFloor, aUser.direction)
	selectedElevator.moveElevator(aUser.floor)
	fmt.Println("User is entering the elevator.")
	fmt.Println("")
	fmt.Println("======= Moving to Destination =========")
	selectedElevator.moveElevator(requestedFloor)
	fmt.Println("User has arrived at his destination")
	fmt.Println("========================================")
	return aUser
}

// =============== This method is used when the user wants to go *to RC* from another floor
func (b *battery) requestElevator(floorNumber int) user {
	// Setting user position and direction //
	var aUser user
	aUser.floor = floorNumber

	if aUser.floor >= 1 {
		aUser.direction = "down"

	} else if aUser.floor <= -1 {
		aUser.direction = "up"
	}

	// Printing lines to make the program clearer //
	fmt.Println("======== Request ground floor ==========")
	fmt.Println("A request has been made at floor", aUser.floor, ". User is going", aUser.direction)
	fmt.Println("")
	fmt.Println("======== Returning a column ============")

	// Nesting multiple methods. Find column, FindElevator and MoveElevator //
	selectedColumn := b.findColumn(floorNumber)
	selectedElevator := selectedColumn.findElevator(selectedColumn, floorNumber, floorNumber, aUser.direction)
	fmt.Println("======== Elevator is moving to User ===========")
	fmt.Println("Elevator n°", selectedColumn.id, "-", selectedElevator.id, "is currently at floor", selectedElevator.currentFloor)
	selectedElevator.currentFloor = selectedElevator.moveElevator(floorNumber)
	fmt.Println("===== Elevator is returning to RC ========")
	selectedElevator.moveElevator(0)
	return aUser
}

// ============== this method finds an elevator after finding a column =========== //
func (c *column) findElevator(selectedColumn column, requestedFloor int, userFloor int, userDirection string) elevator {
	for i := range selectedColumn.elevatorsList {
		floorDiff := int(math.Abs(float64(requestedFloor) - float64(selectedColumn.elevatorsList[i].currentFloor)))
		if requestedFloor > 0 || userFloor > 0 {
			if floorDiff == 0 && selectedColumn.elevatorsList[i].direction == "idle" {
				selectedColumn.elevatorsList[i].value = 2

			} else if floorDiff > 0 && selectedColumn.elevatorsList[i].direction == "idle" && userDirection == "down" {
				selectedColumn.elevatorsList[i].value = 4

			} else if floorDiff > 0 && selectedColumn.elevatorsList[i].direction == "idle" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 5

			} else if floorDiff > 0 && selectedColumn.elevatorsList[i].direction == "down" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 10

			} else if floorDiff > 0 && selectedColumn.elevatorsList[i].direction == "up" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 20

			} else if floorDiff == 0 && selectedColumn.elevatorsList[i].direction == "up" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 25

			} else if floorDiff == 0 && selectedColumn.elevatorsList[i].direction == "down" && userDirection == "down" {
				selectedColumn.elevatorsList[i].value = 35

			} else if floorDiff > 0 && selectedColumn.elevatorsList[i].direction == "up" && userDirection == "down" {
				selectedColumn.elevatorsList[i].value = 40

			} else if floorDiff == 0 && selectedColumn.elevatorsList[i].direction == "up" && userDirection == "down" {
				selectedColumn.elevatorsList[i].value = 45

			} else if floorDiff == 0 && selectedColumn.elevatorsList[i].direction == "down" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 50
			}
		} else if requestedFloor < 0 || userFloor < 0 {

			if floorDiff == 0 && selectedColumn.elevatorsList[i].direction == "idle" {
				selectedColumn.elevatorsList[i].value = 4

			} else if floorDiff > 0 && selectedColumn.elevatorsList[i].direction == "idle" {
				selectedColumn.elevatorsList[i].value = 5

			} else if selectedColumn.elevatorsList[i].currentFloor < userFloor && selectedColumn.elevatorsList[i].direction == "up" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 2

			} else if selectedColumn.elevatorsList[i].currentFloor < userFloor && selectedColumn.elevatorsList[i].direction == "up" && userDirection == "down" {
				selectedColumn.elevatorsList[i].value = 8

			} else if selectedColumn.elevatorsList[i].currentFloor > userFloor && selectedColumn.elevatorsList[i].direction == "down" && userDirection == "up" {
				selectedColumn.elevatorsList[i].value = 10
			} else {
				selectedColumn.elevatorsList[i].value = 15
			}
		}

		// Determines the totalScore of each elevators, later used to compare each elevator //
		selectedColumn.elevatorsList[i].totalScore = floorDiff + selectedColumn.elevatorsList[i].value
		fmt.Println("Elevator n°", selectedColumn.elevatorsList[i].id, "has a total score of", selectedColumn.elevatorsList[i].totalScore)
	}
	// Sorting the elevatorsList based on their totalScore given before //
	sort.Slice(selectedColumn.elevatorsList, func(i, j int) bool {
		return selectedColumn.elevatorsList[i].totalScore < selectedColumn.elevatorsList[j].totalScore
	})
	// Selecting the best elevator to use, first in the list is the winning one after sorting //
	selectedElevator := selectedColumn.elevatorsList[0]
	fmt.Println("")
	fmt.Println("======= Returning an elevator =========")
	fmt.Println("Elevator n°", selectedElevator.id, "has been selected")
	fmt.Println("Elevator n°", selectedElevator.id, "is currently at floor", selectedElevator.currentFloor)
	fmt.Println("")
	return selectedElevator
}

// ============== this method finds a column after a request =========== //
func (b *battery) findColumn(requestedFloor int) column {
	var selectedColumn column
	// For each column in the list, finding the one that matches the request //
	// based on the max and min operating floors. Only one can be selected.  //
	for i, eachColumn := range b.columnList {
		if requestedFloor <= b.columnList[i].maxOperatingFloor && requestedFloor >= b.columnList[i].minOperatingFloor {
			selectedColumn = eachColumn
		}
	}
	// Returning the column that has been selected //
	fmt.Println("Column n°", selectedColumn.id, ", operating from floor", selectedColumn.minOperatingFloor, "to", selectedColumn.maxOperatingFloor, ", has been selected.")
	fmt.Println("")
	return selectedColumn
}

// ============== this method moves the elevator that is used =========== //
func (e *elevator) moveElevator(requestedFloor int) int {

	// While elevator is below the requested floor ==> go UP //
	for e.currentFloor < requestedFloor {
		e.currentFloor = e.currentFloor + 1
		e.direction = "up"
		fmt.Println("Elevator n°", e.id, "going", e.direction, "to floor", e.currentFloor)
		// if elevator is at the same floor as the request ==> STOP //
		if e.currentFloor == requestedFloor {
			fmt.Println("")
			fmt.Println("======== Elevator has stopped =========")
			e.direction = "idle"
			fmt.Println("Elevator n°", e.id, "is", e.direction, "at floor", e.currentFloor)
		}
	}
	// While elevator is over the requested floor ==> go DOWN //
	for e.currentFloor > requestedFloor {
		e.currentFloor = e.currentFloor - 1
		e.direction = "down"
		fmt.Println("Elevator n°", e.id, "going", e.direction, "to floor", e.currentFloor)
		// if elevator is at the same floor as the request ==> STOP //
		if e.currentFloor == requestedFloor {
			fmt.Println("")
			fmt.Println("======== Elevator has stopped =========")
			e.direction = "idle"
			fmt.Println("Elevator n°", e.id, "is", e.direction, "at floor", e.currentFloor)
		}
	}
	return e.currentFloor
}

//     ===================================================     //
//    Here are described a few scenarios to test the program   //
//     ===================================================     //

func (b *battery) scenario1() {
	fmt.Println("======================================")
	fmt.Println("||  Scenario 1 has been activated   ||")
	fmt.Println("======================================")
	// =================================================== //
	//        Column 2 ([1]) operating floor #1 to #20
	// =================================================== //
	// Elevator's currentFloor settings :
	b.columnList[1].elevatorsList[0].currentFloor = 20
	b.columnList[1].elevatorsList[1].currentFloor = 3
	b.columnList[1].elevatorsList[2].currentFloor = 13
	b.columnList[1].elevatorsList[3].currentFloor = 15
	b.columnList[1].elevatorsList[4].currentFloor = 6
	// =================================================== //
	// Elevator's direction settings ("up", "idle" or "down") :
	b.columnList[1].elevatorsList[0].direction = "down"
	b.columnList[1].elevatorsList[1].direction = "up"
	b.columnList[1].elevatorsList[2].direction = "down"
	b.columnList[1].elevatorsList[3].direction = "down"
	b.columnList[1].elevatorsList[4].direction = "down"
	// =================================================== //
	// User is at RC and requests the 20th floor :
	b.assignElevator(20)
}
func (b *battery) scenario2() {
	fmt.Println("======================================")
	fmt.Println("||  Scenario 2 has been activated   ||")
	fmt.Println("======================================")
	// =================================================== //
	//        Column 3 ([2]) operating floor #21 to #40
	// =================================================== //
	// Elevator's currentFloor settings :
	b.columnList[2].elevatorsList[0].currentFloor = 0
	b.columnList[2].elevatorsList[1].currentFloor = 23
	b.columnList[2].elevatorsList[2].currentFloor = 33
	b.columnList[2].elevatorsList[3].currentFloor = 40
	b.columnList[2].elevatorsList[4].currentFloor = 39
	// =================================================== //
	// Elevator's direction settings ("up", "idle" or "down") :
	b.columnList[2].elevatorsList[0].direction = "up"
	b.columnList[2].elevatorsList[1].direction = "up"
	b.columnList[2].elevatorsList[2].direction = "down"
	b.columnList[2].elevatorsList[3].direction = "down"
	b.columnList[2].elevatorsList[4].direction = "down"
	// =================================================== //
	// User is at RC and requests the 36th floor :
	b.assignElevator(36)
}
func (b *battery) scenario3() {
	fmt.Println("======================================")
	fmt.Println("||  Scenario 3 has been activated   ||")
	fmt.Println("======================================") // =================================================== //
	//        Column 4 ([3]) operating floor #41 to #60
	// =================================================== //
	// Elevator's currentFloor settings :
	b.columnList[3].elevatorsList[0].currentFloor = 58
	b.columnList[3].elevatorsList[1].currentFloor = 50
	b.columnList[3].elevatorsList[2].currentFloor = 46
	b.columnList[3].elevatorsList[3].currentFloor = 0
	b.columnList[3].elevatorsList[4].currentFloor = 60
	// =================================================== //
	// Elevator's direction settings ("up", "idle" or "down") :
	b.columnList[3].elevatorsList[0].direction = "down"
	b.columnList[3].elevatorsList[1].direction = "up"
	b.columnList[3].elevatorsList[2].direction = "up"
	b.columnList[3].elevatorsList[3].direction = "up"
	b.columnList[3].elevatorsList[4].direction = "down"
	// =================================================== //
	// User is at RC and requests the RC :
	b.requestElevator(54)
}
func (b *battery) scenario4() {
	fmt.Println("======================================")
	fmt.Println("||  Scenario 4 has been activated   ||")
	fmt.Println("======================================")
	// =================================================== //
	//        Column 1 ([0]) operating floor #B6 to #B1
	// =================================================== //
	// Elevator's currentFloor settings :
	b.columnList[0].elevatorsList[0].currentFloor = -4
	b.columnList[0].elevatorsList[1].currentFloor = 0
	b.columnList[0].elevatorsList[2].currentFloor = -3
	b.columnList[0].elevatorsList[3].currentFloor = -6
	b.columnList[0].elevatorsList[4].currentFloor = -1
	// =================================================== //
	// Elevator's direction settings ("up", "idle" or "down") :
	b.columnList[0].elevatorsList[0].direction = "idle"
	b.columnList[0].elevatorsList[1].direction = "idle"
	b.columnList[0].elevatorsList[2].direction = "down"
	b.columnList[0].elevatorsList[3].direction = "up"
	b.columnList[0].elevatorsList[4].direction = "down"
	// =================================================== //
	// User is at RC and requests the RC :
	b.requestElevator(-3)
}
