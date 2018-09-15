package conv

import (
	"github.com/badgerodon/goreify/generics"
)

//go:generate goreify github.com/ElPeque/reflect-db/conv.To uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func To(elem interface{}) generics.T1 {

	switch elem.(type) {

	// unsigned
	case *uint:
		return generics.T1(*elem.(*uint))
	case *uint8:
		return generics.T1(*elem.(*uint8))
	case *uint16:
		return generics.T1(*elem.(*uint16))
	case *uint32:
		return generics.T1(*elem.(*uint32))
	case *uint64:
		return generics.T1(*elem.(*uint64))
	case uint:
		return generics.T1(elem.(uint))
	case uint8:
		return generics.T1(elem.(uint8))
	case uint16:
		return generics.T1(elem.(uint16))
	case uint32:
		return generics.T1(elem.(uint32))
	case uint64:
		return generics.T1(elem.(uint64))

	// signed
	case *int:
		return generics.T1(*elem.(*int))
	case *int8:
		return generics.T1(*elem.(*int8))
	case *int16:
		return generics.T1(*elem.(*int16))
	case *int32:
		return generics.T1(*elem.(*int32))
	case *int64:
		return generics.T1(*elem.(*int64))
	case int:
		return generics.T1(elem.(int))
	case int8:
		return generics.T1(elem.(int8))
	case int16:
		return generics.T1(elem.(int16))
	case int32:
		return generics.T1(elem.(int32))
	case int64:
		return generics.T1(elem.(int64))

		// float
	case *float32:
		return generics.T1(*elem.(*float32))
	case *float64:
		return generics.T1(*elem.(*float64))
	case float32:
		return generics.T1(elem.(float32))
	case float64:
		return generics.T1(elem.(float64))

	}

	return generics.T1(0)
}
