package main

/* Returns true if hold, else false */
type Strategy struct {
	strategyfunc func(turnValues []int, turnTotal int, total int) bool
	desc         string
}

func (s Strategy) String() string {
	return s.desc
}

type Player struct {
	strategy   Strategy
	turnValues []int
	turnTotal  int
	total      int
}

func (p *Player) PlayTurn() {
	/* Record total at the start, in case a '1' voids everything */
	initialTotal := p.total

	/* Roll die until strategy returns 'true' i.e hold */
	for !p.strategy.strategyfunc(p.turnValues, p.turnTotal, p.total) {
		rolledDie := RollDie()

		if rolledDie == 1 {
			p.total = initialTotal
			break
		}
		p.turnValues = append(p.turnValues, rolledDie)
		p.turnTotal += rolledDie
		p.total += rolledDie
	}

	/* Restore turnValues and turnTotal to normal */
	p.turnValues = []int{}
	p.turnTotal = 0
}

type Conductor struct {
	player1 Player
	player2 Player
}

func (c *Conductor) Conduct(rounds int) (p1Wins int) {

	for i := 0; i < rounds; i++ {
		for {
			if c.player2.total >= 100 {
				break
			}
			c.player1.PlayTurn()

			if c.player1.total >= 100 {
				break
			}
			c.player2.PlayTurn()
		}

		if c.player1.total >= 100 {
			p1Wins += 1
		}

		c.player1.total, c.player2.total = 0, 0
	}

	return p1Wins
}
