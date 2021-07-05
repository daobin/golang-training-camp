// +build wireinject

package main

import "github.com/google/wire"

// wire 01
//func InitMission(pp PlayerParam, mp MonsterParam) (Mission, error) {
//	wire.Build(NewMonster, NewPlayer, NewMission)
//	return Mission{}, nil
//}

// wire 02
//func InitEndingA(pp PlayerParam, mp MonsterParam) (EndingA, error) {
//	wire.Build(NewMonster, NewPlayer, NewEndingA)
//	return EndingA{}, nil
//}
//
//func InitEndingB(pp PlayerParam, mp MonsterParam) (EndingB, error) {
//	wire.Build(NewMonster, NewPlayer, NewEndingB)
//	return EndingB{}, nil
//}

// wire 03
//var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)
//func InitEndingA(pp PlayerParam, mp MonsterParam) (EndingA, error) {
//	wire.Build(monsterPlayerSet, NewEndingA)
//	return EndingA{}, nil
//}
//
//func InitEndingB(pp PlayerParam, mp MonsterParam) (EndingB, error) {
//	wire.Build(monsterPlayerSet, NewEndingB)
//	return EndingB{}, nil
//}

// wire 04
//var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)
//var endingCStruct = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingC), "*"))
//func InitEndingC(pp PlayerParam, mp MonsterParam) (EndingC, error) {
//	wire.Build(endingCStruct)
//	return EndingC{}, nil
//}

// wire 05
//var duck = Monster{Name: "Duck"}
//func InitEndingA(pp PlayerParam) (EndingA, error) {
//	wire.Build(NewPlayer, wire.Value(duck), NewEndingA)
//	return EndingA{}, nil
//}

// wire 06
//func InitPlayer() Player {
//	wire.Build(NewMission2, wire.FieldsOf(new(Mission), "Player"))
//	return Player{}
//}

// wire 07
func InitMissionOfCleanup(name string) (Mission, func(), error) {
	wire.Build(NewMonster2, NewPlayer2, NewMission)
	return Mission{}, nil, nil
}
