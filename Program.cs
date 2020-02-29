using System;
using System.Collections.Generic;
using System.Linq;

namespace Elevator_Control_Project___Week_1___Odyssey__
{
    class Program
    {
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
        static void Main(string[] args)
        {
            //====== Creating a new battery here =======//
            Battery b1 = new Battery(1, 0, 66, 6, 4, 20);
            //==========================================//
            //
            //===== Select one of these scenarios ======//
            //          uncomment to activate           //
            b1.scenario1();
            // b1.scenario2();
            // b1.scenario3();
            // b1.scenario4();
            //==========================================//
            //
            //===== Custom Scenarios you can create ====//
            //       go to line #237 to configure       //
            //          uncomment to activate           //
            // b1.manualSettings();
            //==========================================//
            //b1.assignElevator(#floor);
            //b1.requestElevator(#floor);
        }


    // ============================================//
    //          This is the Battery Class          //
    public class Battery
    {
        public int id, mainFloor,totalFloor, totalBasement, totalColumn, totalElevator, requestedFloor;
        public List<Column> columnList;
        public Column selectedColumn;

        public int userFloor;
        public string userDirection;

        public Elevator selectedElevator;

        public Battery (int id, int mainFloor, int totalFloor, int totalBasement, int totalColumn, int totalElevator)
        {
            this.id = id;
            this.mainFloor = mainFloor;
            this.totalFloor = totalFloor;
            this.totalBasement = totalBasement;
            this.totalColumn = totalColumn;
            this.totalElevator = totalElevator;
            this.columnList = new List<Column>();

            var numElevator = this.totalElevator / this.totalColumn;

            for (int i = 1; i <= this.totalColumn; i++){

                var	numFloorPerColumn = (totalFloor - totalBasement) / (totalColumn - 1);
		        var maxOperatingFloor = ((i - 1) * numFloorPerColumn);
		        var minOperatingFloor = maxOperatingFloor - (numFloorPerColumn - 1);

                if (i == 1) {
                    Column column = new Column(
                        id: i,
                        numberOfelevator: numElevator,
                        minOperatingFloor: (this.mainFloor - this.totalBasement),
                        maxOperatingFloor: this.mainFloor - 1
                    );
                    columnList.Add(column);
                }
                if (i > 1) {
                    Column column = new Column(
                        id: i,
                        numberOfelevator: numElevator,
                        minOperatingFloor: minOperatingFloor,
                        maxOperatingFloor: maxOperatingFloor
                    );
                    columnList.Add(column);
                }
            }
        }

        // === This is used to request RC from another floor === //
        public void requestElevator(int floorNumber)
        {
            userFloor = floorNumber;
            if(userFloor >= 1){
                userDirection = "down";
            } else if (userFloor < -1){
                userDirection = "up";
            }
            requestedFloor = 0;
            Console.WriteLine("");
            Console.WriteLine("======== Request ground floor ==========");
	        Console.WriteLine("A request has been made at floor " + userFloor + ". User is going " + userDirection);
	        Console.WriteLine("");
	        Console.WriteLine("======== Returning a column ============");
            selectedColumn = findColumn(floorNumber);
            selectedElevator = selectedColumn.findElevator(selectedColumn, floorNumber, userDirection);
            Console.WriteLine("=========== Moving to User ============");
            selectedElevator.moveElevator(floorNumber);
            Console.WriteLine("User is entering the elevator.");
            Console.WriteLine("");
            Console.WriteLine("========== Returning to RC ============");
            selectedElevator.moveElevator(requestedFloor);
            Console.WriteLine("User has arrived at his destination.");
            Console.WriteLine("");
        }

        // This method is used when the user *from RC* requests another floor //
        public void assignElevator(int requestedFloor)
        {
            userFloor = 0;
            if (userFloor > requestedFloor){
                userDirection = "down";
            }else if (userFloor < requestedFloor){
                userDirection = "up";
            }
            Console.WriteLine("");
            Console.WriteLine("======== Requesting a floor ==========");
	        Console.WriteLine("A request has been made to go to " + requestedFloor + ". User is going " + userDirection);
	        Console.WriteLine("");
	        Console.WriteLine("======== Returning a column ============");

            selectedColumn = findColumn(requestedFloor);
            selectedElevator = selectedColumn.findElevator(selectedColumn, userFloor, userDirection);
            selectedElevator.moveElevator(userFloor);
            Console.WriteLine("User is entering the elevator.");
            Console.WriteLine("");
            Console.WriteLine("======= Moving to destination =========");
            selectedElevator.moveElevator(requestedFloor);
            Console.WriteLine("User has arrived at his destination");
            Console.WriteLine("=======================================");
        }

        public void scenario2() {
        Console.WriteLine("======================================");
	    Console.WriteLine("||  Scenario 2 has been activated   ||");
	    Console.WriteLine("======================================");
	    // =================================================== //
	    //        Column 3 ([2]) operating floor #21 to #40    //
	    // =================================================== //
	    // Elevator's currentFloor settings :
	    columnList[2].elevatorsList[0].currentFloor = 0;
	    columnList[2].elevatorsList[1].currentFloor = 23;
	    columnList[2].elevatorsList[2].currentFloor = 33;
	    columnList[2].elevatorsList[3].currentFloor = 40;
	    columnList[2].elevatorsList[4].currentFloor = 39;
	    // =================================================== //
	    // Elevator's direction settings ("up", "idle" or "down") :
	    columnList[2].elevatorsList[0].direction = "up";
	    columnList[2].elevatorsList[1].direction = "up";
	    columnList[2].elevatorsList[2].direction = "down";
	    columnList[2].elevatorsList[3].direction = "down";
	    columnList[2].elevatorsList[4].direction = "down";
	    // =================================================== //
	    // User is at RC and requests the 36th floor :
	    assignElevator(36);
        }
        public void scenario1() {
	    Console.WriteLine("======================================");
	    Console.WriteLine("||  Scenario 1 has been activated   ||");
	    Console.WriteLine("======================================");
	    // =================================================== //
	    //        Column 2 ([1]) operating floor #1 to #20
	    // =================================================== //
	    // Elevator's currentFloor settings :
	    columnList[1].elevatorsList[0].currentFloor = 20;
	    columnList[1].elevatorsList[1].currentFloor = 3;
	    columnList[1].elevatorsList[2].currentFloor = 13;
	    columnList[1].elevatorsList[3].currentFloor = 15;
	    columnList[1].elevatorsList[4].currentFloor = 6;
	    // =================================================== //
	    // Elevator's direction settings ("up", "idle" or "down") :
	    columnList[1].elevatorsList[0].direction = "down";
	    columnList[1].elevatorsList[1].direction = "up";
	    columnList[1].elevatorsList[2].direction = "down";
	    columnList[1].elevatorsList[3].direction = "down";
	    columnList[1].elevatorsList[4].direction = "down";
	    // =================================================== //
	    // User is at RC and requests the 20th floor :
	    assignElevator(20);
        }
        public void scenario3() {
        Console.WriteLine("======================================");
        Console.WriteLine("||  Scenario 3 has been activated   ||");
        Console.WriteLine("======================================");
        // =================================================== //
        //        Column 4 ([3]) operating floor #41 to #60
        // =================================================== //
        // Elevator's currentFloor settings :
        columnList[3].elevatorsList[0].currentFloor = 58;
        columnList[3].elevatorsList[1].currentFloor = 50;
        columnList[3].elevatorsList[2].currentFloor = 46;
        columnList[3].elevatorsList[3].currentFloor = 0;
        columnList[3].elevatorsList[4].currentFloor = 60;
        // =================================================== //
        // Elevator's direction settings ("up", "idle" or "down") :
        columnList[3].elevatorsList[0].direction = "down";
        columnList[3].elevatorsList[1].direction = "up";
        columnList[3].elevatorsList[2].direction = "up";
        columnList[3].elevatorsList[3].direction = "up";
        columnList[3].elevatorsList[4].direction = "down";
        // =================================================== //
        // User is at RC and requests the RC :
        requestElevator(54);
        }
        public void scenario4() {
        Console.WriteLine("======================================");
        Console.WriteLine("||  Scenario 4 has been activated   ||");
        Console.WriteLine("======================================");
        // =================================================== //
        //        Column 1 ([0]) operating floor #B6 to #B1
        // =================================================== //
        // Elevator's currentFloor settings :
        columnList[0].elevatorsList[0].currentFloor = -4;
        columnList[0].elevatorsList[1].currentFloor = 0;
        columnList[0].elevatorsList[2].currentFloor = -3;
        columnList[0].elevatorsList[3].currentFloor = -6;
        columnList[0].elevatorsList[4].currentFloor = -1;
        // =================================================== //
        // Elevator's direction settings ("up", "idle" or "down") :
        columnList[0].elevatorsList[0].direction = "idle";
        columnList[0].elevatorsList[1].direction = "idle";
        columnList[0].elevatorsList[2].direction = "down";
        columnList[0].elevatorsList[3].direction = "up";
        columnList[0].elevatorsList[4].direction = "down";
        // =================================================== //
        // User is at RC and requests the RC :
        requestElevator(-3);
        }

        public void manualSettings()
        {   Console.WriteLine("================================");
            Console.WriteLine(" * Manual settings activated * ");
            Console.WriteLine("================================");
            // =================================================== //
            //        Column 1 ([0]) operating the basements
            //        Column 2 ([1]) operating floor #1 to #20
            //        Column 3 ([2]) operating floor #21 to #40
            //        Column 4 ([3]) operating floor #41 to #60
            // =================================================== //
            // Elevator's currentFloor settings :
            columnList[1].elevatorsList[0].currentFloor = 2;
            columnList[1].elevatorsList[1].currentFloor = 0;
            columnList[1].elevatorsList[2].currentFloor = 10;
            columnList[1].elevatorsList[3].currentFloor = 4;
            columnList[1].elevatorsList[4].currentFloor = 19;
            // =================================================== //
            // Elevator's direction settings ("up", "idle" or "down") :
            columnList[1].elevatorsList[0].direction = "idle";
            columnList[1].elevatorsList[1].direction = "idle";
            columnList[1].elevatorsList[2].direction = "idle";
            columnList[1].elevatorsList[3].direction = "idle";
            columnList[1].elevatorsList[4].direction = "idle";
            // =================================================== //
        }

        // === This is used to select a column based on the floors its operating === //
        public Column findColumn(int requestedFloor)

        {   // == going through each column and finding the one matching the user request == //
            for (int i = 0; i < columnList.Count; i++){
                if (requestedFloor <= columnList[i].maxOperatingFloor && requestedFloor >= columnList[i].minOperatingFloor) {
			    selectedColumn = columnList[i];
		        };
            }
            // == Returning a column as the selectedColumn == //
            Console.WriteLine("Column n°" + selectedColumn.id + ", operating from floor " + selectedColumn.minOperatingFloor + " to " + selectedColumn.maxOperatingFloor + ", has been selected.");
	        Console.WriteLine("");
            return selectedColumn;
        }
    }
    // ============================================//
    //          This is the Column Class           //
     public class Column
    {
        public int id, numberOfelevator, minOperatingFloor, maxOperatingFloor, floorDiff;
        public List<Elevator> elevatorsList;
        public Elevator selectedElevator;

        // ========= Column Constructor ========== //
        public Column (int id, int numberOfelevator, int minOperatingFloor, int maxOperatingFloor)
        {
            this.id = id;
            this.numberOfelevator = numberOfelevator;
            this.minOperatingFloor = minOperatingFloor;
            this.maxOperatingFloor = maxOperatingFloor;
            this.elevatorsList = new List<Elevator>();

            // == Generating elevators in elevatorsList == //
            for (int i = 1; i <= numberOfelevator; i++) {
                Elevator elevator = new Elevator (
                    id: i,
                    value: 0,
                    currentFloor: 0,
                    totalScore: 0,
                    direction: "idle",
                    doorStatus: "closed"
                );
                elevatorsList.Add(elevator);
            }
        }
        public Elevator findElevator(Column selectedColumn, int userFloor, string userDirection)
        {   // === Going through each elevators in the list of elevators for the selected Column === //
            for (int i = 0; i < elevatorsList.Count; i++){
                floorDiff = Math.Abs(userFloor - elevatorsList[i].currentFloor);
                //   Conditions for the upper floors   //
                if (userFloor >= 0){

                    if (floorDiff == 0 && elevatorsList[i].direction == "idle") {
				        elevatorsList[i].value = 2;
			        } else if (floorDiff > 0 && elevatorsList[i].direction == "idle" && userDirection == "down") {
				        elevatorsList[i].value = 4;
			        } else if (floorDiff > 0 && elevatorsList[i].direction == "idle" && userDirection == "up") {
				        elevatorsList[i].value = 5;
			        } else if (floorDiff > 0 && elevatorsList[i].direction == "down" && userDirection == "up") {
				        elevatorsList[i].value = 10;
			        } else if (floorDiff > 0 && elevatorsList[i].direction == "up" && userDirection == "up") {
				        elevatorsList[i].value = 20;
			        } else if (floorDiff == 0 && elevatorsList[i].direction == "up" && userDirection == "up") {
				        elevatorsList[i].value = 25;
			        } else if (floorDiff == 0 && elevatorsList[i].direction == "down" && userDirection == "down") {
				        elevatorsList[i].value = 35;
			        } else if (floorDiff > 0 && elevatorsList[i].direction == "up" && userDirection == "down") {
				        elevatorsList[i].value = 40;
			        } else if (floorDiff == 0 && elevatorsList[i].direction == "up" && userDirection == "down") {
				        elevatorsList[i].value = 45;
			        } else if (floorDiff == 0 && elevatorsList[i].direction == "down" && userDirection == "up") {
				        elevatorsList[i].value = 50;
                    }

                //    Conditions for the basements    //
                } else if (userFloor <= - 1){
                    if (floorDiff == 0 && elevatorsList[i].direction == "idle") {
				        elevatorsList[i].value = 4;
			        } else if (floorDiff > 0 && elevatorsList[i].direction == "idle") {
				        elevatorsList[i].value = 5;
			        } else if (elevatorsList[i].currentFloor < userFloor && elevatorsList[i].direction == "up" && userDirection == "up") {
				        elevatorsList[i].value = 2;
			        } else if (elevatorsList[i].currentFloor < userFloor && elevatorsList[i].direction == "up" && userDirection == "down") {
				        elevatorsList[i].value = 8;
			        } else if (elevatorsList[i].currentFloor > userFloor && elevatorsList[i].direction == "down" && userDirection == "up") {
				        elevatorsList[i].value = 10;
			        } else {
				        elevatorsList[i].value = 15;
			        }
                }
            // ======= Assigning totalScore to each elevators =========//
            elevatorsList[i].totalScore = floorDiff + elevatorsList[i].value;
            Console.WriteLine("Elevator n°" + elevatorsList[i].id + "|| Value : " + elevatorsList[i].value + " || Total score : " + elevatorsList[i].totalScore);
            }
            // === Sorting the list of elevators by their totalScore === //
            for(var i = 0; i < selectedColumn.elevatorsList.Count; i++)
                selectedColumn.elevatorsList.Sort((p1, p2) =>
            {
                return p1.totalScore - p2.totalScore;
            });

            // === Returning an elevator as the selected Elevator === //
            selectedElevator = selectedColumn.elevatorsList[0];
            Console.WriteLine("");
            Console.WriteLine("====== Returning an elevator ========");
            Console.WriteLine("Elevator n°" + selectedElevator.id + " selected.");
	        Console.WriteLine("");
            return selectedElevator;
        }
    }

    // ============================================//
    //          This is the Elevator Class         //
    public class Elevator
    {
        public int id, value, currentFloor, totalScore;
        public string direction, doorStatus;

        public Elevator (int id, int value, int currentFloor, int totalScore, string direction, string doorStatus)
        {
            this.id = id;
            this.value = value;
            this.currentFloor = currentFloor;
            this.totalScore = totalScore;
            this.direction = direction;
            this.doorStatus = doorStatus;
        }

        public int moveElevator(int requestedFloor)
        {
            Console.WriteLine("Elevator n°" + id + " || currently at floor: " + currentFloor);
            Console.WriteLine("");
            while (currentFloor < requestedFloor){
                direction = "up";
                currentFloor = currentFloor + 1;
                Console.WriteLine("Elevator n°" + id + " || going " + direction + " to floor: " + currentFloor);
                    if (currentFloor == requestedFloor) {
                        Console.WriteLine("");
			            Console.WriteLine("======== Elevator has stopped =========");
			            this.direction = "idle";
			            Console.WriteLine("Elevator n°" + id + " || " + direction + " at floor: " + currentFloor);
                        Console.WriteLine("");
                    }
            }
            while (currentFloor > requestedFloor){
                direction = "down";
                currentFloor = currentFloor - 1;
                Console.WriteLine("Elevator n°" + id + " || going " + direction + " to floor: " + currentFloor);
                    if (currentFloor == requestedFloor) {
                        Console.WriteLine("");
			            Console.WriteLine("======== Elevator has stopped =========");
			            direction = "idle";
			            Console.WriteLine("Elevator n°" + id + " || " + direction + " at floor: " + currentFloor);
                        Console.WriteLine("");
                    }
            }
            return currentFloor;
        }
    }

    }
}
