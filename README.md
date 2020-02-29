# Rocket_Elevators_Controllers

Contains the algorithm files for the elevator controllers for the New Rocket Elevator Solutions for both Residential and Commercial offers

=====================================
GoLang Commercial Building Controller
=====================================

Using the program is fairly simple. At line 24 you have the main() function
Shown here :

    func main() {
	b1 := batteryInit() <== this initiate the battery

==> If you want to manually change battery settings, go to line 107.
Shown here :

    func batteryInit() battery {
	b1 := battery{

		id:            1,   <== Battery id.
		mainFloor:     0,   <== RC floor, set to 0.
		totalFloor:    66,  <== Total num of floors in your building.
		totalBasement: 6,   <== Total num of basements in your building.
		totalColumn:   4,   <== Total num of columns in your building.
		totalElevator: 20}  <== Total num of elevators in your building.

==> Once the battery is generatated, you can either play prefab scenarios
    or you can manually configure your own scenarios.

Prefab scenario are at line 34 :

    // ========= Scenario player ========== //
	//     Select one scenario at a time    //
	//        uncomment to activate         //
	// ==================================== //
	// b1.scenario1()
	// b1.scenario2()
	// b1.scenario3()
	// b1.scenario4()

Manual settings are configured at line 53:

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

=====================================
C# Commercial Building Controller
=====================================

Using the program is fairly simple. At line 23 you have the main() function
Shown here :

     static void Main(string[] args)
        {
        //====== Creating a new battery here =======//
         Battery b1 = new Battery(1, 0, 66, 6, 4, 20);

==> There you can generate a battery. The program won't run without it.

==> Once the battery is generatated, you can either play prefab scenarios
    or you can manually configure your own scenarios.

Prefab scenario are at line 29 :

        //===== Select one of these scenarios ======//
        //          uncomment to activate           //
        // b1.scenario1();
        // b1.scenario2();
        // b1.scenario3();
        // b1.scenario4();

Manual settings are configured at line 259 :

        // Elevator's currentFloor settings :
    columnList[1].elevatorsList[0].currentFloor = 2;
    columnList[1].elevatorsList[1].currentFloor = 0;
    columnList[1].elevatorsList[2].currentFloor = 10;
    columnList[1].elevatorsList[3].currentFloor = 4;
    columnList[1].elevatorsList[4].currentFloor = 19;
    // ========================================== //
    // Elevator's direction settings ("up", "idle" or "down") :
    columnList[1].elevatorsList[0].direction = "idle";
    columnList[1].elevatorsList[1].direction = "idle";
    columnList[1].elevatorsList[2].direction = "idle";
    columnList[1].elevatorsList[3].direction = "idle";
    columnList[1].elevatorsList[4].direction = "idle";