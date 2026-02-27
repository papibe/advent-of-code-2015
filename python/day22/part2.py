from collections import deque
from typing import Any, Dict, List, Optional

MISSILE_COST: int = 53
DRAIN_COST: int = 73
SHIELD_COST: int = 113
POISON_COST: int = 173
RECHARGE_COST: int = 229

OPTIONS: Dict[str, int] = {
    "missile": MISSILE_COST,
    "drain": DRAIN_COST,
    "shield": SHIELD_COST,
    "poison": POISON_COST,
    "recharge": RECHARGE_COST,
}


class Player:
    def __init__(self, hp: int, mana: int) -> None:
        self.hp: int = hp
        self.mana: int = mana
        self.armor: int = 0
        self.spent_mana: int = 0
        self.boss: Optional[Boss] = None
        self.state: Optional[State] = None

    def set_state(self, state: "State") -> None:
        self.state = state

    def set_boss(self, boss: "Boss") -> None:
        self.boss = boss

    def plays(self) -> List[str]:
        actions: List[str] = []
        assert self.state is not None
        for action, cost in OPTIONS.items():
            if action == "shield" and self.state.shield > 0:
                continue
            if action == "poison" and self.state.poison > 0:
                continue
            if action == "recharge" and self.state.recharge > 0:
                continue

            if self.mana >= cost:
                actions.append(action)
        return actions

    def play(self, action: str) -> None:
        match action:
            case "missile":
                self.missile(),
            case "drain":
                self.drain(),
            case "shield":
                self.shield(),
            case "poison":
                self.poison(),
            case "recharge":
                self.recharge(),

    def missile(self) -> bool:
        self.mana -= MISSILE_COST
        self.spent_mana += MISSILE_COST
        assert self.boss is not None
        return self.boss.receive_attack(4)

    def drain(self) -> bool:
        self.mana -= DRAIN_COST
        self.spent_mana += DRAIN_COST
        assert self.boss is not None

        self.hp += 2
        return self.boss.receive_attack(2)

    def shield(self) -> bool:
        self.mana -= SHIELD_COST
        self.spent_mana += SHIELD_COST
        assert self.state is not None

        self.state.shield = 6
        return False

    def poison(self) -> bool:
        self.mana -= POISON_COST
        self.spent_mana += POISON_COST
        assert self.state is not None

        self.state.poison = 6
        return False

    def recharge(self) -> bool:
        self.mana -= RECHARGE_COST
        self.spent_mana += RECHARGE_COST
        assert self.state is not None

        self.state.recharge = 5
        return False

    def receive_attack(self, damage: int) -> bool:
        actual_damage: int = damage - self.armor

        if actual_damage <= 0:
            actual_damage = 1

        self.hp -= actual_damage
        if self.hp <= 0:
            return True
        return False


class Boss:
    def __init__(self, hp: int, damage: int) -> None:
        self.hp: int = hp
        self.damage: int = damage
        self.player: Optional[Player] = None

    def set_player(self, player: Player) -> None:
        self.player = player

    def set_state(self, state: "State") -> None:
        self.state = state

    def plays(self) -> List[str]:
        return ["attack"]

    def play(self, action: str) -> None:
        self.attack()

    def attack(self) -> bool:
        assert self.player is not None
        return self.player.receive_attack(self.damage)

    def receive_attack(self, damage: int) -> bool:
        self.hp -= damage
        if self.hp <= 0:
            return True
        return False


class State:
    def __init__(self, player: Player, boss: Boss) -> None:
        self.player: Player = player
        self.boss: Boss = boss
        self.turns: List[Any] = [player, boss]
        self.index: int = 0

        # interactions with spells
        self.shield: int = 0
        self.recharge: int = 0
        self.poison: int = 0

    def run_spells(self) -> bool:
        boss_killed: bool = False

        if self.shield > 0:
            self.player.armor = 7
            self.shield -= 1
        else:
            self.player.armor = 0

        if self.poison > 0:
            self.poison -= 1
            boss_killed = self.boss.receive_attack(3)

        if self.recharge > 0:
            self.recharge -= 1
            self.player.mana += 101

        return boss_killed

    def get_turn(self) -> Player | Boss:
        next_turn: Player | Boss = self.turns[self.index]
        self.index = (self.index + 1) % 2
        return next_turn


def solution(hp: int, mana: int, boss_hp: int, boss_damage: int) -> int:
    player: Player = Player(hp, mana)
    boss: Boss = Boss(boss_hp, boss_damage)

    player.set_boss(boss)
    boss.set_player(player)

    state: State = State(player, boss)

    player.set_state(state)
    boss.set_state(state)

    min_mana: int = float("inf")  # type: ignore

    queue = deque([state])
    visited = set(
        [
            (
                state.player.hp,
                state.player.mana,
                state.player.armor,
                state.boss.hp,
                state.index,
                state.shield,
                state.recharge,
                state.poison,
            )
        ]
    )

    while queue:
        state = queue.popleft()

        if state.player.hp <= 0 or state.player.mana <= 0:
            continue

        if state.boss.hp <= 0:
            min_mana = min(min_mana, state.player.spent_mana)
            continue

        if state.run_spells():
            min_mana = min(min_mana, state.player.spent_mana)
            continue

        if state.player.spent_mana > min_mana:
            continue

        participant = state.get_turn()
        if participant == state.player:
            participant.hp -= 1
            if participant.hp <= 0:
                continue

        participant_plays = participant.plays()
        if not participant_plays:
            continue

        for action in participant_plays:
            new_player: Player = Player(state.player.hp, state.player.mana)
            new_boss: Boss = Boss(state.boss.hp, state.boss.damage)

            new_player.armor = state.player.armor
            new_player.spent_mana = state.player.spent_mana

            new_player.set_boss(new_boss)
            new_boss.set_player(new_player)

            new_state: State = State(new_player, new_boss)
            new_state.index = state.index
            new_state.shield = state.shield
            new_state.recharge = state.recharge
            new_state.poison = state.poison

            new_player.set_state(new_state)
            new_boss.set_state(new_state)

            if action == "attack":
                new_boss.play(action)
            else:
                new_player.play(action)

            new_key = (
                new_state.player.hp,
                new_state.player.mana,
                new_state.player.armor,
                new_state.boss.hp,
                new_state.index,
                new_state.shield,
                new_state.recharge,
                new_state.poison,
            )

            if new_key not in visited:
                queue.append(new_state)
                visited.add(new_key)

    return min_mana


if __name__ == "__main__":
    print(solution(50, 500, 55, 8))  # 953
