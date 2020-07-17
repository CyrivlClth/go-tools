// package serialnum
//
// get serial number from redis
package serialnum

import (
	"sync"

	"github.com/CyrivlClth/snowflake/idgen"
)

type generator struct {
	mutex sync.Mutex
	num   int64
	max   int64
}

func (g *generator) GetID() int64 {
	i, _ := g.NextID()
	return i
}

func (g *generator) NextID() (int64, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	n := g.num
	g.num++
	if g.num > g.max {
		g.num = 0
	}
	return n, nil
}

func New(max int64) idgen.IDGenerator {
	return &generator{max: max}
}
