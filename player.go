package main

import "fmt"
import "time"

type Player struct {
	UUID         string
	Name         string
	Level        int
	Experience   int64
	Health       int
	Attack       chan int
	AttackDamage int
	AttackSpeed  int
	Ded          chan int
	Stop         chan int
	// TODO: whatever
}

func (p Player) Fight() {
	ticker := time.NewTicker(time.Millisecond * time.Duration(p.AttackSpeed))
	for {
		select {
		case <-ticker.C:
			p.Attack <- p.AttackDamage
		case <-p.Stop:
			ticker.Stop()
			return
		}
	}
}

func (p *Player) Damage(dmg int) {
	p.Health = p.Health - dmg
	fmt.Println(p.Name, " damaged for ", dmg, " damage.")
	fmt.Println(p.Name, " now has ", p.Health, " health.")
	if p.Health <= 0 {
		p.Ded <- 0
	}
}

func Battle(p1, p2 Player) string {
	go p1.Fight()
	go p2.Fight()

	for {
		select {
		case dmg := <-p1.Attack:
			p2.Damage(dmg)
		case dmg := <-p2.Attack:
			p1.Damage(dmg)
		case <-p1.Ded:
			p1.Stop <- 0
			p2.Stop <- 0
			return p2.Name
		case <-p2.Ded:
			p1.Stop <- 0
			p2.Stop <- 0
			return p1.Name
		}
	}
}

func main() {
	p1 := Player{
		Name:         "player 1",
		Health:       200,
		Attack:       make(chan int),
		AttackDamage: 20,
		AttackSpeed:  2000,
		Ded:          make(chan int),
		Stop:         make(chan int),
	}
	p2 := Player{
		Name:         "player 2",
		Health:       100,
		Attack:       make(chan int),
		AttackDamage: 10,
		AttackSpeed:  500,
		Ded:          make(chan int),
		Stop:         make(chan int),
	}
	winner := Battle(p1, p2)
	fmt.Println("winner", winner)
}
