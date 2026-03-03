package main

import (
	"fmt"
	"math"
)

const (
	MISSILE_COST  = 53
	DRAIN_COST    = 73
	SHIELD_COST   = 113
	POISON_COST   = 173
	RECHARGE_COST = 229
)

var OPTIONS = map[string]int{
	"missile":  MISSILE_COST,
	"drain":    DRAIN_COST,
	"shield":   SHIELD_COST,
	"poison":   POISON_COST,
	"recharge": RECHARGE_COST,
}

type VisitedKey struct {
	player_hp      int
	player_mana    int
	player_armor   int
	boss_hp        int
	state_index    int
	state_shield   int
	state_recharge int
	state_poison   int
}

type Player struct {
	hp         int
	mana       int
	armor      int
	spent_mana int
	boss       *Boss
	state      *State
}

type Boss struct {
	hp     int
	damage int
	player *Player
	state  *State
}

type State struct {
	player   *Player
	boss     *Boss
	turns    []Participant
	index    int
	shield   int
	recharge int
	poison   int
}

type Rule struct {
	pattern     string
	replacement string
}

func NewPlayer(hp, mana int) *Player {
	p := &Player{hp, mana, 0, 0, nil, nil}
	return p
}

func (p *Player) set_state(state *State) {
	p.state = state
}

func (p *Player) set_boss(boss *Boss) {
	p.boss = boss
}

func (p *Player) plays() []string {
	actions := []string{}
	for action, cost := range OPTIONS {
		if (action == "shield" && p.state.shield > 0) ||
			(action == "poison" && p.state.poison > 0) ||
			(action == "recharge" && p.state.recharge > 0) {
			continue
		}
		if p.mana >= cost {
			actions = append(actions, action)
		}
	}
	return actions
}

func (p *Player) play(action string) {
	switch action {
	case "missile":
		p.missile()
	case "drain":
		p.drain()
	case "shield":
		p.shield()
	case "poison":
		p.poison()
	case "recharge":
		p.recharge()
	}
}

func (p *Player) missile() bool {
	p.mana -= MISSILE_COST
	p.spent_mana += MISSILE_COST

	return p.boss.receive_attack(4)
}

func (p *Player) drain() bool {
	p.mana -= DRAIN_COST
	p.spent_mana += DRAIN_COST

	p.hp += 2
	return p.boss.receive_attack(2)
}

func (p *Player) shield() bool {
	p.mana -= SHIELD_COST
	p.spent_mana += SHIELD_COST

	p.state.shield = 6
	return false
}

func (p *Player) poison() bool {
	p.mana -= POISON_COST
	p.spent_mana += POISON_COST

	p.state.poison = 6
	return false
}

func (p *Player) recharge() bool {
	p.mana -= RECHARGE_COST
	p.spent_mana += RECHARGE_COST

	p.state.recharge = 5
	return false
}

func (p *Player) receive_attack(damage int) bool {
	actual_damage := damage - p.armor

	if actual_damage <= 0 {
		actual_damage = 1
	}
	p.hp -= actual_damage
	return p.hp <= 0
}

func (p *Player) get_hp() int {
	return p.hp
}

func (p *Player) decrease_hp() {
	p.hp -= 1
}

func NewBoss(hp, damage int) *Boss {
	b := &Boss{hp, damage, nil, nil}
	return b
}

func (b *Boss) set_player(player *Player) {
	b.player = player
}

func (b *Boss) set_state(state *State) {
	b.state = state
}

func (b *Boss) plays() []string {
	return []string{"attack"}
}

func (b *Boss) play(action string) {
	b.attack()
}

func (b *Boss) attack() bool {
	return b.player.receive_attack(b.damage)
}

func (b *Boss) receive_attack(damage int) bool {
	b.hp -= damage
	return b.hp <= 0
}

func (b *Boss) get_hp() int {
	return b.hp
}

func (b *Boss) decrease_hp() {
	b.hp -= 1
}

func NewState(player *Player, boss *Boss) *State {
	s := &State{player, boss, []Participant{player, boss}, 0, 0, 0, 0}
	return s
}

func (s *State) run_spells() bool {
	boss_killed := false

	if s.shield > 0 {
		s.player.armor = 7
		s.shield -= 1
	} else {
		s.player.armor = 0
	}

	if s.poison > 0 {
		s.poison -= 1
		boss_killed = s.boss.receive_attack(3)
	}

	if s.recharge > 0 {
		s.recharge -= 1
		s.player.mana += 101
	}

	return boss_killed
}

func (s *State) get_turn() Participant {
	next_turn := s.turns[s.index]
	s.index = (s.index + 1) % 2
	return next_turn
}

type Participant interface {
	plays() []string
	get_hp() int
	decrease_hp()
}

func get_key_from_state(state *State) VisitedKey {
	return VisitedKey{
		state.player.hp,
		state.player.mana,
		state.player.armor,
		state.boss.hp,
		state.index,
		state.shield,
		state.recharge,
		state.poison,
	}
}

func NewPlayerFromPlayer(player *Player) *Player {
	return &Player{
		player.hp, player.mana, player.armor, player.spent_mana, nil, nil,
	}
}

func NewStateFromState(player *Player, boss *Boss, state *State) *State {
	return &State{
		player,
		boss,
		[]Participant{player, boss},
		state.index,
		state.shield,
		state.recharge,
		state.poison,
	}
}

func solution(hp, mana, boss_hp, boss_damage int) int {
	player := NewPlayer(hp, mana)
	boss := NewBoss(boss_hp, boss_damage)

	player.set_boss(boss)
	boss.set_player(player)

	state := NewState(player, boss)

	player.set_state(state)
	boss.set_state(state)

	min_mana := math.MaxInt

	queue := NewQueue[*State]()
	queue.append(state)

	visited := NewSet[VisitedKey]()
	key := get_key_from_state(state)
	visited.add(key)

	for !queue.is_empty() {
		state = queue.popleft()

		if state.player.hp <= 0 || state.player.mana <= 0 {
			continue
		}

		if state.boss.hp <= 0 {
			min_mana = min(min_mana, state.player.spent_mana)
			continue
		}
		if state.run_spells() {
			min_mana = min(min_mana, state.player.spent_mana)
			continue
		}
		if state.player.spent_mana > min_mana {
			continue
		}
		participant := state.get_turn()

		// path for part 2
		if participant == state.player {
			participant.decrease_hp()
			hp := participant.get_hp()
			if hp <= 0 {
				continue
			}
		}
		participant_plays := participant.plays()

		for _, action := range participant_plays {
			// set up for next play (action)
			new_player := NewPlayerFromPlayer(state.player)
			new_boss := NewBoss(state.boss.hp, state.boss.damage)

			new_player.set_boss(new_boss)
			new_boss.set_player(new_player)

			new_state := NewStateFromState(new_player, new_boss, state)

			new_player.set_state(new_state)
			new_boss.set_state(new_state)

			// take action (play)
			if action == "attack" {
				new_boss.play(action)
			} else {
				new_player.play(action)
			}

			new_key := get_key_from_state(new_state)
			if !visited.contains(new_key) {
				queue.append(new_state)
				visited.add(new_key)
			}
		}
	}

	return min_mana
}

func main() {
	fmt.Println(solution(50, 500, 55, 8)) // 1289
}
