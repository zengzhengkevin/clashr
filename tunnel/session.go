package tunnel

import (
	"sync"
	"time"

	nat "github.com/zu1k/clashr/component/nat-table"
)

var (
	natTable *nat.Table
	natOnce  sync.Once

	natTimeout = 120 * time.Second
)

func NATInstance() *nat.Table {
	natOnce.Do(func() {
		natTable = nat.New(natTimeout)
	})
	return natTable
}
