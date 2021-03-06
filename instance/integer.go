package instance

import (
	"math"
	"github.com/recursivecurry/gobox/typeclass/eq"
	"github.com/recursivecurry/gobox/typeclass/enum"
	"github.com/recursivecurry/gobox/typeclass/bounded"
	"github.com/recursivecurry/gobox/typeclass/read"
	"strconv"
	"fmt"
)

type Integer int64

const INTEGER Integer = Integer(0)

func (i Integer) FuncRead(s string) read.Interface {
	val, _ := strconv.ParseInt(s, 10, 64)
	return Integer(val)
}

func (i Integer) Show() string {
	return fmt.Sprintf("%v", i)
}

func (i Integer) FuncToEnum(a int64) enum.Interface {
	return Integer(a)
}

func (i Integer) FromEnum() int64 {
	return int64(i)
}

func (i Integer) Eq(other eq.Interface) bool {
	va, ok := other.(Integer)
	if !ok {
		return false
	}
	return int64(va) == int64(i)
}

func (i Integer) Ne(other eq.Interface) bool {
	return !i.Eq(other)
}

func (i Integer) FuncMinBound() bounded.Interface {
	return Integer(math.MinInt64)
}

func (i Integer) FuncMaxBound() bounded.Interface {
	return Integer(math.MaxInt64)
}

func (i Integer) Less(other eq.Interface) bool {
	i2, ok := other.(Integer)
	if !ok {
		return false
	}
	return int64(i) < int64(i2)
}
