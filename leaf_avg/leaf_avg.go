package leaf_avg

import (
	"github.com/ElPeque/reflect-db/conv"
	"github.com/ElPeque/reflect-db/types"
)

func LeafAvg(avg *float64) types.WalkCallback {
	var count int
	return func(path []string, obj interface{}) bool {
		*avg = (*avg*float64(count) + conv.ToFloat64(obj)) / float64(count+1)
		count++

		return true
	}
}
