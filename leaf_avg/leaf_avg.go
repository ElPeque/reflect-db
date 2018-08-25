package leaf_avg

import (
	"github.com/ElPeque/reflect-db/types"
)

func LeafAvg(avgIface interface{}) types.WalkCallback {
	var count int
	return func(path []string, obj interface{}) bool {

		switch avgIface.(type) {
		case *uint:
			ptr := avgIface.(*uint)
			val, ok := obj.(uint)
			if ok {
				*ptr = (*ptr*uint(count) + val) / uint(count+1)
				count++
				break
			}
			pVal, ok := obj.(*uint)
			if ok {
				*ptr = (*ptr*uint(count) + *pVal) / uint(count+1)
				count++
				break
			}
		case *uint16:
			ptr := avgIface.(*uint16)
			val, ok := obj.(uint16)
			if ok {
				*ptr = (*ptr*uint16(count) + val) / uint16(count+1)
				count++
				break
			}
			pVal, ok := obj.(*uint16)
			if ok {
				*ptr = (*ptr*uint16(count) + *pVal) / uint16(count+1)
				count++
				break
			}
		case *uint32:
			ptr := avgIface.(*uint32)
			val, ok := obj.(uint32)
			if ok {
				*ptr = (*ptr*uint32(count) + val) / uint32(count+1)
				count++
				break
			}
			pVal, ok := obj.(*uint32)
			if ok {
				*ptr = (*ptr*uint32(count) + *pVal) / uint32(count+1)
				count++
				break
			}
		case *uint64:
			ptr := avgIface.(*uint64)
			val, ok := obj.(uint64)
			if ok {
				*ptr = (*ptr*uint64(count) + val) / uint64(count+1)
				count++
				break
			}
			pVal, ok := obj.(*uint64)
			if ok {
				*ptr = (*ptr*uint64(count) + *pVal) / uint64(count+1)
				count++
				break
			}
		case *int:
			ptr := avgIface.(*int)
			val, ok := obj.(int)
			if ok {
				*ptr = (*ptr*int(count) + val) / int(count+1)
				count++
				break
			}
			pVal, ok := obj.(*int)
			if ok {
				*ptr = (*ptr*int(count) + *pVal) / int(count+1)
				count++
				break
			}
		case *int8:
			ptr := avgIface.(*int8)
			val, ok := obj.(int8)
			if ok {
				*ptr = (*ptr*int8(count) + val) / int8(count+1)
				count++
				break
			}
			pVal, ok := obj.(*int8)
			if ok {
				*ptr = (*ptr*int8(count) + *pVal) / int8(count+1)
				count++
				break
			}
		case *int16:
			ptr := avgIface.(*int32)
			val, ok := obj.(int32)
			if ok {
				*ptr = (*ptr*int32(count) + val) / int32(count+1)
				count++
				break
			}
			pVal, ok := obj.(*int32)
			if ok {
				*ptr = (*ptr*int32(count) + *pVal) / int32(count+1)
				count++
				break
			}
		case *int32:
			ptr := avgIface.(*int32)
			val, ok := obj.(int32)
			if ok {
				*ptr = (*ptr*int32(count) + val) / int32(count+1)
				count++
				break
			}
			pVal, ok := obj.(*int32)
			if ok {
				*ptr = (*ptr*int32(count) + *pVal) / int32(count+1)
				count++
				break
			}
		case *int64:
			ptr := avgIface.(*int64)
			val, ok := obj.(int64)
			if ok {
				*ptr = (*ptr*int64(count) + val) / int64(count+1)
				count++
				break
			}
			pVal, ok := obj.(*int64)
			if ok {
				*ptr = (*ptr*int64(count) + *pVal) / int64(count+1)
				count++
				break
			}
		case *float32:
			ptr := avgIface.(*float32)
			val, ok := obj.(float32)
			if ok {
				*ptr = (*ptr*float32(count) + val) / float32(count+1)
				count++
				break
			}
			pVal, ok := obj.(*float32)
			if ok {
				*ptr = (*ptr*float32(count) + *pVal) / float32(count+1)
				count++
				break
			}
		case *float64:
			ptr := avgIface.(*float64)
			val, ok := obj.(float64)
			if ok {
				*ptr = (*ptr*float64(count) + val) / float64(count+1)
				count++
				break
			}
			pVal, ok := obj.(*float64)
			if ok {
				*ptr = (*ptr*float64(count) + *pVal) / float64(count+1)
				count++
				break
			}
		}
		return true
	}
}
