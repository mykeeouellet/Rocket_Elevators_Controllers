//-- This Program controls Elevators in a Residential building using Object-Oriented-Programming--//

//Here is described the Column object with its attributes & methods//
class Column {
    constructor(id, totalFloor, totalElevator) {
        this.id = id;
        this.totalFloor = totalFloor;
        this.totalElevator = totalElevator;
        this.elevatorsList = [];
        this.callbuttonsList = [];
        this.floorsList = [];

        //-- This for loop adds new elevators to elevatorsList --//
        for(var i = 1; i <= totalElevator; i++) {
            this.elevatorsList.push(new Elevator(i, totalFloor));
        }
        //-- This for loop adds new floors to floorsList --//
        for(var i = 1; i <= totalFloor; i++) {
            this.floorsList.push(i);
        }
        //-- This for loop adds new callbuttons to callbuttonsList --//
        for(var i = 1; i <= totalFloor; i++){
            //--This if-statement creates 9 'down' buttons--//
            if (i != 1) {
                var callbutton = new CallButton('down', i);
                this.callbuttonsList.push(callbutton);
            }
            //--This if-statement creates 9 'up' buttons--//
            if (i < totalFloor) {
                var callbutton = new CallButton('up', i);
                this.callbuttonsList.push(callbutton);
            }
        }
    }

    //--This is the method that is called when the user requests an elevator--//
    requestElevator(direction, requestedFloor){
        console.log("");
        console.log("---------------------User calls the Elevator--------------------");
        console.log("A request has been made at floor " + requestedFloor + ". User is going " + direction + ".");
        console.log("");
        console.log("-----------------------Finding an Elevator----------------------");
        var elevator = this.findElevator(requestedFloor, direction);
        console.log("Elevator " + this.elevatorsList[0].id + " has a total score of " + this.elevatorsList[0].totalScore);
        console.log("Elevator " + this.elevatorsList[1].id + " has a total score of " + this.elevatorsList[1].totalScore);
        console.log("");
        console.log(" ----------------Returning an Elevator to the User---------------");
        console.log("Elevator " + this.elevatorsList[0].id + " has been selected to go to floor " + requestedFloor);
        elevator.moveElevator(requestedFloor);
        console.log(" ------------------User is entering the Elevator---------------");
        console.log("User is entering elevator " + elevator.id + ".");

            for(var i = 0; i < this.elevatorsList.length; i++) {
                this.elevatorsList[i].totalScore = 0;
                this.elevatorsList[i].value = 0;
            }

        return elevator;
    }


    //--This is the method that is called when the user requests a floor--//
    requestFloor(elevator, requestedFloor){
        this.id = elevator;
        this.requestedFloor = requestedFloor;
        var elevator = this.findElevator(requestedFloor);
        console.log("User has selected floor " + this.requestedFloor + " as his destination");
        console.log("");
        this.doorStatus = "closed";
        console.log("Doors at elevator " + elevator.id + " are now " + this.doorStatus + ".");
        elevator.moveElevator(requestedFloor);
    }


    //--This is the method that finds an elevator depending on direction and position of the user--//
    findElevator(requestedFloor, direction){

        //--This loop goes through each elevator in elevatorsList, giving them a score--//
        for(var i = 0; i < this.elevatorsList.length; i++) {
                let diff = Math.abs(parseInt(requestedFloor) - parseInt(this.elevatorsList[i].currentFloor));

            if (diff === 0 && this.elevatorsList[i].direction === "idle" && direction === "up"){
                this.elevatorsList[i].value = 1;
            } else if(diff === 0 && this.elevatorsList[i].direction === "idle" && direction=== "down"){
                this.elevatorsList[i].value = 2;
            } else if(diff > 0 && this.elevatorsList[i].direction === "idle" && direction === "down"){
                this.elevatorsList[i].value = 3;
            } else if(diff > 0 && this.elevatorsList[i].direction === "idle" && direction === "up"){
                this.elevatorsList[i].value = 4;
            } else if (diff > 0 && this.elevatorsList[i].direction === "down" && direction === "down"){
                this.elevatorsList[i].value = 5;
            } else if (diff > 0 && this.elevatorsList[i].direction === "up" && direction === "up"){
                this.elevatorsList[i].value = 6;
            } else if (diff === 0 && this.elevatorsList[i].direction === "up" && direction === "up"){
                this.elevatorsList[i].value = 7;
            } else if (diff === 0 && this.elevatorsList[i].direction === "down" && direction === "down"){
                this.elevatorsList[i].value = 8;
            } else if (diff > 0 && this.elevatorsList[i].direction === "up" && direction === "down"){
                this.elevatorsList[i].value = 9;
            } else if (diff > 0 && this.elevatorsList[i].direction === "down" && direction === "up"){
                this.elevatorsList[i].value = 10;
            } else if (diff === 0 && this.elevatorsList[i].direction === "up" && direction === "down"){
                this.elevatorsList[i].value = 11;
            } else if (diff === 0 && this.elevatorsList[i].direction === "down" && direction === "up"){
                this.elevatorsList[i].value = 12;
            }else {
                this.elevatorsList[i].value = 13;
            }
            //--This compute totalScore as difference plus value--//
            this.elevatorsList[i].totalScore = diff + this.elevatorsList[i].value;
        }
        //--This function is used to sort the list of elevators depending on their totalScore--//
        function compare(a, b) {
            return a.totalScore - b.totalScore;
        }this.elevatorsList.sort(compare);

        //--This pushes the request to the elevator's requestsList--//
        this.elevatorsList[0].requestsList.push({requestedFloor, direction});

        return this.elevatorsList[0];
        }
}

//-- Here is described the CallButton object with its attributes --//
class CallButton {
    constructor(direction, floor) {
        this.direction = direction;
        this.floor = floor;
    }
}

//-- Here is described the Request Object with its attributes --//
class Request {
    constructor(direction, floor) {
        this.direction = direction;
        this.floor = floor;
    }
}

//-- Here is described the Elevator object with its attributes & methods-- //
class Elevator {
    constructor(id, totalFloor) {
        this.id = id;
        this.value = 0;
        this.totalScore = 0;
        this.direction = "idle";
        this.currentFloor = 1;
        this.currentWeight = 0;
        this.weightLimit = 1500;
        this.doorStatus = "closed";
        this.elevatorbuttonsList = [];
        this.requestsList = [];

        //--This loop creates buttons for each floors in the elevator--//
        for(var i = 1; i <= totalFloor; i++) {
            this.elevatorbuttonsList.push(i);
        }
    }

    //-- This is the method that is called when the elevator is moving --//
    moveElevator(requestedFloor){
        this.requestedFloor = requestedFloor;
        console.log("Elevator " + this.id + " is currently at floor " + this.currentFloor);
        console.log("");
        console.log(" ---------------------Moving the Elevator-----------------");

        //--These If-statements are responsible for moving the elevator to the correct floor--//
        if(this.currentFloor < requestedFloor){
            while (this.currentFloor < requestedFloor){
                this.currentFloor = this.currentFloor + 1;
                this.direction = "up";
                console.log("Elevator " + this.id + " is going " + this.direction + " to floor " + this.currentFloor);
            }
        } else if(this.currentFloor > requestedFloor){
            while (this.currentFloor > requestedFloor){
                this.currentFloor = this.currentFloor - 1;
                this.direction = "down";
                console.log("Elevator " + this.id + " is going " + this.direction + " to floor " + this.currentFloor);
            }
        }
        if(this.currentFloor === requestedFloor){
            this.direction = "idle";
            console.log("");
            console.log(" ---------------------Elevator has arrived----------------------  ");
            console.log("Elevator " + this.id + " is " + this.direction + " at floor " + this.currentFloor);
            this.doorStatus = "opened";
            console.log("Doors at Elevator " + this.id + " are now " + this.doorStatus + " !" );
            console.log("");
            console.log(" ---------------------//--------------------  ");
            console.log("");
        }
    }
}


//let Name_Here = new Column(column id, total number of floors, total number of elevators)
let column1 = new Column(1, 10, 2);

//----------------SECTION TEST--------------//
console.log(column1);

//--Elevator 1 currentFloor and Direction--//
// column1.elevatorsList[0].currentFloor = 1;
// column1.elevatorsList[0].direction = "idle";

//--Elevator 2 currentFloor and Direction--//
// column1.elevatorsList[1].currentFloor = 1;
// column1.elevatorsList[1].direction = "idle";


var elevator = column1.requestElevator("up", 5);
elevator = column1.requestElevator("down", 10);
