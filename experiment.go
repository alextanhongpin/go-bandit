package bandit

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Experiment holds all the Strategy in-memory
// type Experiment struct {
// 	sync.Mutex
// 	Strategies map[string]Strategy
// }

// func (e *Experiment) genID() string {
// 	return fmt.Sprint(time.Now().UnixNano())
// }

// // NewEpsilonGreedy
// func (e *Experiment) NewEpsilonGreedy(nArms int, epsilonDecay float64) Strategy {
// 	id := e.genID()
// 	e.Lock()
// 	e.Strategies[id] = New(nArms, epsilonDecay)
// 	e.Unlock()
// 	return e.Strategies[id]
// }

// // Delete will remove an existing strategy
// func (e *Experiment) Delete(id string) {
// 	e.Lock()
// 	if _, ok := e.Strategies[id]; ok == true {
// 		delete(e.Strategies, id)
// 	}
// 	e.Unlock()
// }

// // All will return all strategies, paginated
// func (e *Experiment) All() []Response {
// 	var strategies []Response
// 	for id, v := range e.Strategies {
// 		strategies = append(strategies, Response{id, v})
// 	}
// 	return strategies
// }

// // One will return a single strategy by id, or nil
// func (e *Experiment) One(id string) Strategy {
// 	e.Lock()
// 	v, ok := e.Strategies[id]
// 	e.Unlock()
// 	if !ok {
// 		return nil
// 	}
// 	return v
// }

// // SelectArm will trigger the arm selection
// func (e Experiment) SelectArm() {}

// // Update will trigger the arm update
// func (e Experiment) Update() {}

// func NewExperiment() *Experiment {
// 	return &Experiment{
// 		Strategies: make(map[string]Strategy),
// 	}
// }
