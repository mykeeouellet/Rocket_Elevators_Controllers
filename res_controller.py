# this is the main function used to test the program
def main():
    column1 = Column(1, 10, 2)
    elevator = column1.requestElevator('up', 2)
    # column1.requestFloor(elevator, 10)

# this is the function you can use in case you wanna set custom scenarios
def manualSettings():
    column1 = Column(1, 10, 2)
    main()


# this is used to create callbuttons for each floor with their direction
class Callbutton():
    def __init__(self, direction, floor):
        self.direction = direction
        self.floor = floor


    def __str__(self):
        return str(self.__class__) + ": " + str(self.__dict__)


# this is used to create a request object
class Request():
    def __init__(self, direction, floor):
        self.direction = direction
        self.floor = floor


# this defines the Column object
class Column():
    def __init__(self, id, totalFloor, totalElevator):
        self.id = id
        self.totalFloor = totalFloor
        self.totalElevator = totalElevator
        self.elevatorsList = []
        self.callbuttonsList = []
        self.floorsList = []

        # this for loop creates new Elevators in elevatorsList
        for i in range(totalElevator):
            self.elevatorsList.append(Elevator(i + 1, totalFloor))

        # this for loop creates a floorsList for each column created
        for i in range(totalFloor):
            self.floorsList.append(i + 1)

        #this forLoop creates each callButtons for floor in the building
        for i in range(totalFloor):
            if i != 1:
                callbutton = Callbutton("down", i + 1)
                self.callbuttonsList.append(callbutton)

            if i < totalFloor:
                callbutton = Callbutton("up", i + 1)
                self.callbuttonsList.append(callbutton)


    def __str__(self):
        return str(self.__class__) + ": " + str(self.__dict__)


    # This is the method that is used to request an elevator
    def requestElevator(self, direction, requestedFloor):
        self.direction = direction
        self.requestedFloor = requestedFloor
        print('======================================================')
        print('                User Calls the elevator                ')
        print('======================================================')
        print("A request has been made at floor", str(requestedFloor), ". User is going", str(direction), ".")
        print('------------------------------------------------------')
        print("            Returning an Elevator to the User         ")
        print("")
        elevator = self.findElevator(direction, requestedFloor)

        #for i in range(len(self.elevatorsList)):
            #print("Elevator", self.elevatorsList[i].id, "has a total score of", self.elevatorsList[i].totalScore)

        # this forLoop resets the TotalScore and value of the elevator
        for i in range(len(self.elevatorsList)):
            self.elevatorsList[i].totalScore = 0
            self.elevatorsList[i].value = 0

        return elevator


    # This is the method that is used to request a floor and move the elevator
    def requestFloor(self, elevator, requestedFloor):
        self.id = elevator
        self.requestFloor = requestedFloor
        elevator = self.elevatorsList[0]
        print("User is entering elevator", str(elevator.id), ".")
        print("User has selected floor", str(requestedFloor), "as his destination")
        print('')
        elevator.doorStatus = "closed"
        print("Doors at elevator", str(elevator.id), "are now", str(elevator.doorStatus), ".")
        elevator.moveElevator(requestedFloor)



    # This is the method that is used to compare the different elevators
    def findElevator(self, direction, requestedFloor):
        #this forLoop looks at conditions to compare each elevator of elevatorsList
        for i in range(len(self.elevatorsList)):
            # determines diff as user floor minus current floor of each elevators
            diff = abs(int(requestedFloor) - int(self.elevatorsList[i].currentFloor))

            #if condition is true, than sets value of each elevator to a number.
            if diff == 0 and self.elevatorsList[i].direction == "idle" and direction == "up":
                self.elevatorsList[i].value = 1

            elif diff == 0 and self.elevatorsList[i].direction == "idle" and direction == "up":
                self.elevatorsList[i].value = 2

            elif diff > 0 and self.elevatorsList[i].direction == "idle" and direction == "down":
                self.elevatorsList[i].value = 3

            elif diff > 0 and self.elevatorsList[i].direction == "idle" and direction == "up":
                self.elevatorsList[i].value = 4

            elif diff > 0 and self.elevatorsList[i].direction == "down" and direction == "down":
                self.elevatorsList[i].value = 5

            elif diff > 0 and self.elevatorsList[i].direction == "up" and direction == "up":
                self.elevatorsList[i].value = 6

            elif diff == 0 and self.elevatorsList[i].direction == "up" and direction == "up":
                self.elevatorsList[i].value = 7

            elif diff == 0 and self.elevatorsList[i].direction == "down" and direction == "up":
                self.elevatorsList[i].value = 8

            elif diff > 0 and self.elevatorsList[i].direction == "up" and direction == "down":
                self.elevatorsList[i].value = 9

            elif diff > 0 and self.elevatorsList[i].direction == "down" and direction == "up":
                self.elevatorsList[i].value = 10

            elif diff == 0 and self.elevatorsList[i].direction == "up" and direction == "down":
                self.elevatorsList[i].value = 11

            elif diff == 0 and self.elevatorsList[i].direction == "down" and direction == "up":
                self.elevatorsList[i].value = 12

            else :
                self.elevatorsList[i].value = 13

            score = int(diff) + int(self.elevatorsList[i].value)
            self.elevatorsList[i].totalScore = score

        for i in range(len(self.elevatorsList)):
            print("Elevator", self.elevatorsList[i].id, "has a total score of", self.elevatorsList[i].totalScore)

        # sorts the elevatorsList and returning the winning elevator
        self.elevatorsList.sort(key=lambda x: x.totalScore)
        print("")
        print("================================================")
        print("Elevator", str(self.elevatorsList[0].id), "has been selected to go to floor", str(requestedFloor))
        print("================================================")
        print("")
        self.elevatorsList[0].requestsList.append({requestedFloor, direction})
        self.elevatorsList[0].moveElevator(requestedFloor)

        return self.elevatorsList[0]




# this defines the Elevator object and its attributes
class Elevator():
    def __init__(self, id, totalFloor):
        self.id = id
        self.value = 0
        self.totalScore = 0
        self.direction = "idle"
        self.currentFloor = 1
        self.currentWeight = 0
        self.weightLimit = 1500
        self.doorStatus = "closed"
        self.elevatorbuttonsList = []
        self.requestsList = []

        # this forLoop creates floor buttons for each elevator in elevatorsList
        for i in range(totalFloor):
            self.elevatorbuttonsList.append(i + 1)


    def __str__(self):
        return str(self.__class__) + ": " + str(self.__dict__)


    # this is the method used to move the elevators
    def moveElevator(self, requestedFloor):
        self.requestedFloor = requestedFloor
        print('              Moving the elevator             ')
        print('')
        print('Elevator', str(self.id), 'is currently at floor', str(self.currentFloor))
        # if elevator floor is less than user floor => move up
        if self.currentFloor < requestedFloor :
            # this loop adds 1 floor each time elevator floor < user floor
            while self.currentFloor < requestedFloor:
                self.currentFloor = self.currentFloor + 1
                self.direction = "up"
                print("Elevator", str(self.id), "is going", str(self.direction), "to floor", str(self.currentFloor))

        # if elevator floor is greater than user floor => move down
        elif self.currentFloor > requestedFloor:
            # this loop removes 1 floor each time elevator floor > user floor
            while self.currentFloor > requestedFloor:
                self.currentFloor = self.currentFloor - 1
                self.direction = "down"
                print("Elevator", str(self.id), "is going", str(self.direction), "to floor", str(self.currentFloor))

        # if elevator floor is the same as user floor => stop the elevator
        if self.currentFloor == requestedFloor:
            self.direction = "idle"
            self.doorStatus = "Opened"
            print('__________________________________________________')
            print('')
            print('              Elevator has arrived             ')
            print('')
            print("Elevator", str(self.id), "is", str(self.direction), "at floor", str(self.currentFloor))
            print("Doors at Elevator", str(self.id), "are now ", str(self.doorStatus), "!" )
            print('===================================================')

# here, the main() function described at the top is called, executing the program
main()


