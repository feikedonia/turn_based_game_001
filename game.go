package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

// constants
const (
	name_1 string = "Dario"
	name_2 string = "Pilgrim"
	name_3 string = "Fie"
	name_4 string = "Jessy"
)

// variables

// user

var user_input string //player input
var victory bool = false
var display_tutorial bool = true

// Enemy
var enemy_input int //enemy input

// start of program
func main() {

	clear_screen()
	check_victory()

	if display_tutorial != false {
		tutorial()
	}

	//fmt.Println(player_name, "  Health:", player_health, "SP:", player_skill_points, "Gold:", gold)

	show_status()

	fmt.Println("\nWhat do you want to do?")
	fmt.Println("\nbattle\t\t> finds opponent")
	fmt.Println("shop\t\t> enter the shop")
	fmt.Println("stats\t\t> show player stats")
	fmt.Println("inv\t\t> show player inventory")
	fmt.Println("exit\t\t> exits the game")
	fmt.Println("")

	for {
		fmt.Scanln(&user_input)

		switch user_input {

		case "battle":
			combat()

		case "shop":
			Dario.shop()

		case "stats":
			Dario.display_stats()

		case "inv":
			Dario.display_inventory()

		case "exit":
			quit()

		default:
			main()
		}
	}
}

func save(slot1 any, data interface{}) {

}
func show_status() {
	fmt.Println("player_1:\nhealth: ", Dario.health, "skill points: ", Dario.skill_points, "gold: ", Dario.gold)
	fmt.Println("player_1:\nhealth: ", Pilgrim.health, "skill points: ", Pilgrim.skill_points, "gold: ", Pilgrim.gold)
	fmt.Println("player_1:\nhealth: ", Fie.health, "skill points: ", Fie.skill_points, "gold: ", Fie.gold)
	fmt.Println("player_1:\nhealth: ", Jessy.health, "skill points: ", Jessy.skill_points, "gold: ", Jessy.gold)
}

func check_victory() {
	if victory != false {
		victory = false
		Dario.exp += rand.Intn(50) + 50
		Pilgrim.exp += rand.Intn(50) + 50
		Fie.exp += rand.Intn(50) + 50
		Jessy.exp += rand.Intn(50) + 50
		Dario.level_check()
		Fie.level_check()
		Pilgrim.level_check()
		Jessy.level_check()
		main()
	}
}

// starts the combat encounter
func combat() {
	fmt.Println("\n\nCombat started!")

	for {
		check_player_life()
		check_enemy_life()
		if victory == true {
			Dario.gold += rand.Intn(10) + 5
			Pilgrim.gold += rand.Intn(10) + 5
			Fie.gold += rand.Intn(10) + 5
			Jessy.gold += rand.Intn(10) + 5
			main()
		}

		Dario_turn()

		enemy_1.enemy_turn()

	}
}

func (p *player) Dario_turn() {
	fmt.Println("")
	if Dario.special >= 3 {
		{
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 91, "You feel a strange power welling up inside... (type 'special' to unleash it)")
			fmt.Println(colored)
		}
	}

	fmt.Println("\nWhat's your move?")
	fmt.Println("\n>> strike\t\t\t> Use your basic weapon\t")
	fmt.Println(">> heal\t\t\t\t> Use an healing item\t")
	fmt.Println(">> force | 20 SP\t\t> High citical chance attack")
	fmt.Println(">> soul \t\t\t> Regenerates some SP")
	fmt.Println("")

	fmt.Scanln(&user_input)

	switch user_input { //gives different options to the player

	case "strike":
		Dario.player_skill_strike()

	case "heal":
		Jessy.player_skill_heal()

	case "force":
		Jessy.player_skill_force()

	case "soul":
		Jessy.player_skill_soul()

	case "kill":
		Dario.player_skill_kill()

	case "special":
		if Dario.special > 2 {
			Dario.special = 0
			Dario_skill_special()
		} else {
			fmt.Println("You dont have the energy for this move")
		}
	default:
		fmt.Println("Thats a typo! lost your turn XD")
	}
	if Dario.health > Dario.max_health {
		Dario.health = Dario.max_health
	}
	if Dario.skill_points > Dario.skill_points {
		Dario.skill_points = Dario.max_skill_points
	}
}

// function for enemy turn
func (e *enemy) enemy_turn() {

	enemy_input = rand.Intn(3) //gives different options to the enemy

	if e.health >= 1 {
		switch enemy_input {

		case 0:
			enemy_skill_strike()

		case 1:
			enemy_skill_heal()

		case 2:
			enemy_skill_force()
		}
	}
	if e.health > e.max_health {
		e.health = e.max_health
	}
	if e.skill_points > e.max_skill_points {
		e.skill_points = e.max_skill_points
	}

}

// displays a tutorial if display_tutorial == true
func tutorial() {
	display_tutorial = false
	fmt.Println("Welcome to this game...")
	fmt.Println("\nThis is a turn based game, as the player you can type the one of the moves to execute it.")
	fmt.Println("Your goal at this moment is to acquire as much gold as possible")

}

// checks if the player is dead
func check_player_life() {
	if Dario.health <= 0 {
		fmt.Println("Your hero has been killed!")
		fmt.Println("\nGold:", Dario.gold, "Player level:", Dario.lv)
		fmt.Println("\nType anything to quit")

		fmt.Scanln("")
		fmt.Scanf("%s", &user_input)
		if user_input == "exit" {
			os.Exit(0)
		} else {
			os.Exit(0)
		}

	}
}

// player skill: kill (THIS IS A TEST FEATURE, NOT MEANT FOR FINAL PRODUCT)
func (e *enemy) Dario_skill_kill() {
	damage := rand.Intn(20) + 5 + Dario.strength + 999
	critical_damage := rand.Intn(20) + 30 + Dario.strength + 999

	if rand.Intn(11) == 9 { //Critical hit chance
		e.health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

// player skill: strike
func (e *enemy) Dario_skill_strike() {
	damage := rand.Intn(20) + 5 + Dario.strength
	critical_damage := rand.Intn(20) + 30 + Dario.strength

	Dario.special += 1

	if rand.Intn(11) == 9 { //Critical hit chance
		e.health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

// player skill: soul
func Jessy_skill_soul() {
	if true == true {
		Jessy.skill_points += 25
	}
}

// player skill: force
func (e *enemy) Jessy_skill_force() {
	damage := rand.Intn(5) + 20 + Jessy.intelligence
	critical_damage := rand.Intn(20) + 30 + Jessy.intelligence

	if Jessy.skill_points >= 20 {

		Jessy.skill_points -= 20

		Jessy.special += 1

		if rand.Intn(3) == 2 { //Critical hit chance

			e.health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			user_input = ""
		} else {
			e.health -= damage
			fmt.Println(damage, "DMG")
			user_input = ""
		}
	} else {
		fmt.Println("You tried to cast force... but you dont have enough SP!")
		user_input = ""
	}
}

// player skill: heal
func Jessy_skill_heal() {
	heal := rand.Intn(20) + 5 + Jessy.intelligence //amount healed
	Jessy.health += heal
	Dario.health += heal
	Pilgrim.health += heal
	Fie.health += heal
	fmt.Println(heal, "Healed")
	user_input = ""
}

func (e *enemy) Dario_skill_special() {
	damage := 70 + Dario.strength
	critical_damage := rand.Intn(20) + 75 + Dario.strength

	if rand.Intn(11) == 9 { //Critical hit chance
		e.health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

// checks if the enemy is dead
func (e *enemy) check_enemy_life() {
	if e.health <= 0 {
		fmt.Println("Victory!")
		victory = true
	}
}

// enemy skill: strike
func (p *player) enemy_skill_strike() {
	fmt.Println("Enemy used strike")
	damage := rand.Intn(20) + 5 - p.endurance
	critical_damage := rand.Intn(20) + 30 - p.endurance

	if rand.Intn(100) > p.agility {

		if rand.Intn(11) == 9 { //Critical hit chance
			p.health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		} else {
			p.health -= damage
			fmt.Println(damage, "DMG")
		}
	} else {
		fmt.Println("But it missed!")
	}
}

// enemy skill: heal
func (e *enemy) enemy_skill_heal() {
	heal := rand.Intn(20) + 5 //amount healed
	fmt.Println("Enemy has healed")
	e.health += heal
	fmt.Println(heal, "Healed")
}

// enemy skill: force
func (p *player) enemy_skill_force(e *enemy) {

	damage := rand.Intn(10) + 20 - p.endurance
	critical_damage := rand.Intn(20) + 30 - p.endurance

	fmt.Println("Enemy cast force")

	if e.skill_points >= 20 {

		e.skill_points -= 20

		if rand.Intn(100) >= p.agility {

			if rand.Intn(3) == 1 { //Critical hit chance
				p.health -= critical_damage
				fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			} else {
				p.health -= damage
				fmt.Println(damage, "DMG")
			}
		} else {
			fmt.Println("but it missed")
		}
	} else {
		fmt.Println("but nothing happened...")
		damage = 0
		p.health -= damage
		fmt.Println(damage, "DMG")
		fmt.Scanln()
	}
}

// clear the screen in the CLI
func clear_screen() {
	cmd := exec.Command("clear") // for Unix/Linux
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// checks the exp and increases the player_lv
func (p *player) level_check() {
	if p.exp >= 100*p.lv {
		p.lv++
		p.max_health += 20
		p.max_skill_points += 5
		p.health = p.max_health
		p.skill_points = p.max_skill_points
		fmt.Println("\nLevel up!!")
		fmt.Printf("\nMax HP: %d, Max SP: %d\n", p.max_health, p.max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("")
		fmt.Scanln(&user_input)

		switch user_input {
		case "st":
			p.strength += 2
		case "in":
			p.intelligence += 2
		case "ag":
			p.agility += 2
		case "en":
			p.endurance += 2
		case "so":
			p.social += 2
		}
	}
}

// Makes the shop work | not finished yet
func (p *player) shop() {

	fmt.Println("Welcome to the shop")
	fmt.Println("\nWe have a variety of products available, please take your time choosing")
	fmt.Println("\n- potion")
	fmt.Println("- sword")
	fmt.Println("- shield")
	fmt.Println("\nleave the shop (back)")

	for {

		fmt.Scanln(&user_input)

		switch user_input { //gives different options to the player

		case "potion":
			fmt.Println("you have bought a potion")
			p.inventory = append(p.inventory, "potion")

		case "sword":
			fmt.Println("you have bought a sword")
			p.inventory = append(p.inventory, "sword")

		case "shield":
			fmt.Println("you have bought a shield")
			p.inventory = append(p.inventory, "shield")

		case "back":
			main()

		default:
			fmt.Println("We don't have this item...")
		}

	}
}

func (p *player) display_stats() {
	fmt.Println("\nPlayer lv:", p.lv)
	fmt.Println("Exp:", p.exp)
	fmt.Println("\nStrength:", p.strength)
	fmt.Println("Intelligence", p.intelligence)
	fmt.Println("Agility:", p.agility)
	fmt.Println("Endurance:", p.endurance)
	fmt.Println("Social:", p.social)
	fmt.Println("\n[back]")

	fmt.Scanln(&user_input)

	switch user_input {

	case "back":
		main()

	default:
		main()
	}
}

// displays the player's inventory | doesn't work yet
func (p *player) display_inventory() {

	fmt.Println(p.inventory)

	fmt.Println("\n[back]")

	fmt.Scanln(&user_input)

	switch user_input {

	case "back":
		main()

	default:
		main()
	}
}

func quit() {
	os.Exit(0)
}
