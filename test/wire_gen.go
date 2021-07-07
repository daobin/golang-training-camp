// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

// wire 07
func InitMissionOfCleanup(name string) (Mission, func(), error) {
	player, cleanup, err := NewPlayer2(name)
	if err != nil {
		return Mission{}, nil, err
	}
	monster := NewMonster2()
	mission := NewMission(player, monster)
	return mission, func() {
		cleanup()
	}, nil
}