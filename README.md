# Rocket_Elevators_Controllers

Contains the algorithm files for the elevator controllers for the New Rocket Elevator Solutions for both Residential and Commercial offers

For the Javascript file => Go to line200. There is a test section. Execute your tests there. Follow this format:

=> Create a column ( like at line.197). The new Column objects takes 3 parameters.
    column_num = new Column(your_Id, your_TotalFloors, your_TotalElevators)

=> Manually creating scenarios : go to line.202

    Elevator 1 currentFloor and Direction
    column1.elevatorsList[0].currentFloor = 1;
    column1.elevatorsList[0].direction = "idle";

    Elevator 2 currentFloor and Direction
    column1.elevatorsList[1].currentFloor = 1;
    column1.elevatorsList[1].direction = "idle";

You can specify the current floor of any elevator you want and their direction. In the default scenario there is 2 elevators and 10 floors. The currentFloor needs to be a number [1,2,..],
and direction needs to be either "up", "down" or "idle".

=> Requesting an elevator and a floor : go to line.211

    {var elevator = column1.requestElevator("up", 5);
    column1.requestFloor(elevator, 10);}

Using this format you make a request for an elevator (requestElevator) with 2 parameters: direction ["up", "down"] and the floor you are at i.e [1,2,3,...]

You can then make a request for a destination. You need to specify the elevator to move and the floor you wanna go. Here (elevator and floor 10) as our elevator in made into a variable so we can use the same elevator in 'requestFloor' as in 'requestElevator'

---------------------------------------------------------------------------------------

For the python file, everything happens at the top (line.2) in the main function:

    def main():
        column1 = Column(1, 10, 2)
        elevator = column1.requestElevator('up', 9)
        column1.requestFloor(elevator, 10)

=> Initialize the program by creating a Column object with 3 parameters as in :

    column1 = Column(your_Id, your_TotalFloors, your_TotalElevators)

=> => Manually creating scenarios : go to line.8 in the manualSettings() function

    # Elevator 1 attributes ( the floor it's at and its direction)
    column1.elevatorsList[0].currentFloor = 1
    column1.elevatorsList[0].direction = "idle"

    # Elevator 2 attributes ( the floor it's at and its direction)
    column1.elevatorsList[1].currentFloor = 1
    column1.elevatorsList[1].direction = "idle"

    main()


You can specify the current floor of any elevator you want and their direction. In the default scenario there is 2 elevators and 10 floors. The currentFloor needs to be a number [1,2,..],and direction needs to be either "up", "down" or "idle".

=> You can now request an elevator and a floor using the same parameters as in the javascript file.

=> You now just have to call the main() function at the end of the program(line.213)!