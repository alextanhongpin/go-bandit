# Multi-Armed Bandit Algorithm

Bandit algorithm balances _exploration_ and _exploitation_, and is part of **reinforcement learning*. 

## Reinforcement Learning

In Reinforcement Learning, the system makes a decision based on current situation. Unlike supervised learning, there is no training data available providing the correct decision to the system. We just get a reward back from the environment, indicating the quality of the decison that was made.

Reinforcement learning problems involve the following artiofacts.

- State
- Action or Decision
- Reward

## Implementations

- Random greedy
- Upper confidence bound one
- Upper confidence bound two
- Softmax
- Interval estimate
- Thompson sampling
- Reward comparison
- Action pursuit
- Exponential weight


## TODO

- implement test for the different algorithms
- plots the different algorithm
- pros/cons for each of them
- use cases and examples

## References

- https://www.quora.com/What-is-Thompson-sampling-in-laymans-terms
- https://www.linkedin.com/pulse/dynamic-price-optimization-multi-arm-bandit-pranab-ghosh
- https://pkghosh.wordpress.com/2013/08/25/bandits-know-the-best-product-price/
- https://github.com/pranab/avenir
- https://www.gsb.stanford.edu/sites/gsb/files/mkt_10_17_misra.pdf
- http://alekhagarwal.net/bandits_and_rl/intro.pdf


## Usage

```golang
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	bandit "github.com/alextanhongpin/go-bandit"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	b, err := bandit.NewEpsilonGreedy(0.1, nil, nil)
	if err != nil {
		log.Println(err)
	}

	b.Init(5)

  N := 1000

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			chosenArm := b.SelectArm(rand.Float64())
			reward := float64(rand.Intn(2))
			b.Update(chosenArm, reward)
		}()
	}

	wg.Wait()
	log.Printf("bandit: %+v", b)
	log.Println("done")
}
```

Test for data race:

```
$ go run -race main.go
```

Output:

```
2018/06/04 23:43:27 bandit: &{RWMutex:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0} Epsilon:0.1 Counts:[233 220 512 19 16] Rewards:[0.4592274678111587 0.48181818181818176 0.5097656249999998 0.3684210526315789 0.25]}
2018/06/04 23:43:27 done
```

<!-- go test -cover -run Epsilon -->