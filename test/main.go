package main

import "log"

func main() {
	// wire 01
	//mission, err := InitMission("LDB","Kitty")
	//if err != nil {
	//	log.Fatalf("Error: %v\r\n", err)
	//}
	//mission.Start()

	// wire 02
	//ending, err := InitEndingA("LDB", "Kitty")
	//if err != nil {
	//	log.Fatalf("Error: %v\r\n", err)
	//}
	//ending.Appear()

	// wire 03
	// ...

	// wire 04
	//ending, err := InitEndingC("LDB", "Kitty")
	//if err != nil {
	//	log.Fatalf("Error: %v\r\n", err)
	//}
	//
	//ending.Appear()

	// wire 05
	//ending, err := InitEndingA("LDB-Kitty")
	//if err != nil {
	//	log.Fatalf("Error: %v\r\n", err)
	//}
	//ending.Appear()

	// wire 06
	//player := InitPlayer()
	//fmt.Println(player.Name)

	// wire 07
	mission, cleanup, err := InitMissionOfCleanup("DDB")
	if err != nil {
		log.Fatalf("Error: %v\r\n", err)
	}
	mission.Start()
	cleanup()
}
