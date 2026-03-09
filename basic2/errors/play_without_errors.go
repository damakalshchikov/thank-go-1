// начало решения

type invalidStepError struct {
	st step
}

func (e invalidStepError) Error() string {
	return fmt.Sprintf("cannot %s", e.st)
}

type notEnoughObjectsError struct {
	obj thing
}

func (e notEnoughObjectsError) Error() string {
	return fmt.Sprintf("there are no %ss left", e.obj)
}

type commandLimitExceededError struct {
	cmd command
}

func (e commandLimitExceededError) Error() string {
	return fmt.Sprintf("you don't want to %s anymore", e.cmd)
}

type objectLimitExceededError struct {
	obj thing
}

func (e objectLimitExceededError) Error() string {
	return fmt.Sprintf("you already have a %s", e.obj)
}

type gameOverError struct {
	nSteps int
	err    error
}

func (e gameOverError) Error() string {
	return e.err.Error()
}

func (e gameOverError) Unwrap() error {
	return e.err
}

type player struct {
	nEaten   int
	nDialogs int
	inventory []thing
}

func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

func (p *player) do(cmd command, obj thing) error {
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{cmd}
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{obj}
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{cmd}
		}
		p.nDialogs++
	}
	return nil
}

func newPlayer() *player {
	return &player{inventory: []thing{}}
}

type game struct {
	player *player
	things map[label]int
	nSteps int
}

func (g *game) has(obj thing) bool {
	return g.things[obj.name] > 0
}

func (g *game) execute(st step) error {
	if !st.isValid() {
		return gameOverError{g.nSteps, invalidStepError{st}}
	}

	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return gameOverError{g.nSteps, notEnoughObjectsError{st.obj}}
		}
		g.things[st.obj.name]--
	}

	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{g.nSteps, err}
	}

	g.nSteps++
	return nil
}

func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

func giveAdvice(err error) string {
	var iSE invalidStepError
	var nEOE notEnoughObjectsError
	var cLEE commandLimitExceededError
	var oLEE objectLimitExceededError

	switch {
	case errors.As(err, &iSE):
		return fmt.Sprintf("things like '%s %s' are impossible", iSE.st.cmd, iSE.st.obj)
	case errors.As(err, &nEOE):
		return fmt.Sprintf("be careful with scarce %ss", nEOE.obj)
	case errors.As(err, &cLEE):
		return fmt.Sprintf("%s less", cLEE.cmd)
	case errors.As(err, &oLEE):
		return fmt.Sprintf("don't be greedy, 1 %s is enough", oLEE.obj)
	}
	return ""
}

// конец решения
