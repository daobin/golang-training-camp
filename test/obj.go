package main

import (
	"errors"
	"fmt"
	"time"
)

type MonsterParam string
type PlayerParam string

type Monster struct {
	Name string
}

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

type Player struct {
	Name string
}

func NewPlayer(name PlayerParam) (Player, error) {
	if time.Now().Unix() %2 == 0 {
		return Player{}, errors.New("player dead")
	}
	return Player{Name: string(name)}, nil
}

type Mission struct {
	Player Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\r\n", m.Player.Name, m.Monster.Name)
}

type EndingA struct {
	Player Player
	Monster Monster
}

func NewEndingA(p Player, m Monster) EndingA {
	return EndingA{p, m}
}

func (e EndingA) Appear() {
	fmt.Printf("%s defeats %s, world peace!\r\n", e.Player.Name, e.Monster.Name)
}

type EndingB struct {
	Player Player
	Monster Monster
}

func NewEndingB(p Player, m Monster) EndingB {
	return EndingB{p, m}
}

func (e EndingB) Appear() {
	fmt.Printf("%s defeats %s, but become monster, world darker!\r\n", e.Player.Name, e.Monster.Name)
}

type EndingC struct {
	Player Player
	Monster Monster
}

func (e EndingC) Appear() {
	fmt.Printf("%s and %s become friend, world peace!\r\n", e.Player.Name, e.Monster.Name)
}

func NewMission2() Mission {
	return Mission{Player:Player{Name:"LDB"}, Monster:Monster{Name:"Kitty"}}
}

func NewMonster2() Monster {
	return Monster{Name: "Duck001"}
}

func NewPlayer2(name string) (Player, func(), error) {
	cleanup := func() {
		fmt.Println("clean up!")
	}

	if time.Now().Unix() % 2 == 0 {
		return Player{}, cleanup, errors.New("player dead")
	}

	return Player{Name:name}, cleanup, nil
}
