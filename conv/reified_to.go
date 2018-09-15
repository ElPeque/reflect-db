package conv

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToUint uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToUint(elem interface{}) uint {

	switch elem.(type) {

	case *uint:
		return uint(*elem.(*uint))
	case *uint8:
		return uint(*elem.(*uint8))
	case *uint16:
		return uint(*elem.(*uint16))
	case *uint32:
		return uint(*elem.(*uint32))
	case *uint64:
		return uint(*elem.(*uint64))
	case uint:
		return uint(elem.(uint))
	case uint8:
		return uint(elem.(uint8))
	case uint16:
		return uint(elem.(uint16))
	case uint32:
		return uint(elem.(uint32))
	case uint64:
		return uint(elem.(uint64))

	case *int:
		return uint(*elem.(*int))
	case *int8:
		return uint(*elem.(*int8))
	case *int16:
		return uint(*elem.(*int16))
	case *int32:
		return uint(*elem.(*int32))
	case *int64:
		return uint(*elem.(*int64))
	case int:
		return uint(elem.(int))
	case int8:
		return uint(elem.(int8))
	case int16:
		return uint(elem.(int16))
	case int32:
		return uint(elem.(int32))
	case int64:
		return uint(elem.(int64))

	case *float32:
		return uint(*elem.(*float32))
	case *float64:
		return uint(*elem.(*float64))
	case float32:
		return uint(elem.(float32))
	case float64:
		return uint(elem.(float64))

	}

	return uint(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToUint8 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToUint8(elem interface{}) uint8 {

	switch elem.(type) {

	case *uint:
		return uint8(*elem.(*uint))
	case *uint8:
		return uint8(*elem.(*uint8))
	case *uint16:
		return uint8(*elem.(*uint16))
	case *uint32:
		return uint8(*elem.(*uint32))
	case *uint64:
		return uint8(*elem.(*uint64))
	case uint:
		return uint8(elem.(uint))
	case uint8:
		return uint8(elem.(uint8))
	case uint16:
		return uint8(elem.(uint16))
	case uint32:
		return uint8(elem.(uint32))
	case uint64:
		return uint8(elem.(uint64))

	case *int:
		return uint8(*elem.(*int))
	case *int8:
		return uint8(*elem.(*int8))
	case *int16:
		return uint8(*elem.(*int16))
	case *int32:
		return uint8(*elem.(*int32))
	case *int64:
		return uint8(*elem.(*int64))
	case int:
		return uint8(elem.(int))
	case int8:
		return uint8(elem.(int8))
	case int16:
		return uint8(elem.(int16))
	case int32:
		return uint8(elem.(int32))
	case int64:
		return uint8(elem.(int64))

	case *float32:
		return uint8(*elem.(*float32))
	case *float64:
		return uint8(*elem.(*float64))
	case float32:
		return uint8(elem.(float32))
	case float64:
		return uint8(elem.(float64))

	}

	return uint8(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToUint16 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToUint16(elem interface{}) uint16 {

	switch elem.(type) {

	case *uint:
		return uint16(*elem.(*uint))
	case *uint8:
		return uint16(*elem.(*uint8))
	case *uint16:
		return uint16(*elem.(*uint16))
	case *uint32:
		return uint16(*elem.(*uint32))
	case *uint64:
		return uint16(*elem.(*uint64))
	case uint:
		return uint16(elem.(uint))
	case uint8:
		return uint16(elem.(uint8))
	case uint16:
		return uint16(elem.(uint16))
	case uint32:
		return uint16(elem.(uint32))
	case uint64:
		return uint16(elem.(uint64))

	case *int:
		return uint16(*elem.(*int))
	case *int8:
		return uint16(*elem.(*int8))
	case *int16:
		return uint16(*elem.(*int16))
	case *int32:
		return uint16(*elem.(*int32))
	case *int64:
		return uint16(*elem.(*int64))
	case int:
		return uint16(elem.(int))
	case int8:
		return uint16(elem.(int8))
	case int16:
		return uint16(elem.(int16))
	case int32:
		return uint16(elem.(int32))
	case int64:
		return uint16(elem.(int64))

	case *float32:
		return uint16(*elem.(*float32))
	case *float64:
		return uint16(*elem.(*float64))
	case float32:
		return uint16(elem.(float32))
	case float64:
		return uint16(elem.(float64))

	}

	return uint16(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToUint32 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToUint32(elem interface{}) uint32 {

	switch elem.(type) {

	case *uint:
		return uint32(*elem.(*uint))
	case *uint8:
		return uint32(*elem.(*uint8))
	case *uint16:
		return uint32(*elem.(*uint16))
	case *uint32:
		return uint32(*elem.(*uint32))
	case *uint64:
		return uint32(*elem.(*uint64))
	case uint:
		return uint32(elem.(uint))
	case uint8:
		return uint32(elem.(uint8))
	case uint16:
		return uint32(elem.(uint16))
	case uint32:
		return uint32(elem.(uint32))
	case uint64:
		return uint32(elem.(uint64))

	case *int:
		return uint32(*elem.(*int))
	case *int8:
		return uint32(*elem.(*int8))
	case *int16:
		return uint32(*elem.(*int16))
	case *int32:
		return uint32(*elem.(*int32))
	case *int64:
		return uint32(*elem.(*int64))
	case int:
		return uint32(elem.(int))
	case int8:
		return uint32(elem.(int8))
	case int16:
		return uint32(elem.(int16))
	case int32:
		return uint32(elem.(int32))
	case int64:
		return uint32(elem.(int64))

	case *float32:
		return uint32(*elem.(*float32))
	case *float64:
		return uint32(*elem.(*float64))
	case float32:
		return uint32(elem.(float32))
	case float64:
		return uint32(elem.(float64))

	}

	return uint32(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToUint64 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToUint64(elem interface{}) uint64 {

	switch elem.(type) {

	case *uint:
		return uint64(*elem.(*uint))
	case *uint8:
		return uint64(*elem.(*uint8))
	case *uint16:
		return uint64(*elem.(*uint16))
	case *uint32:
		return uint64(*elem.(*uint32))
	case *uint64:
		return uint64(*elem.(*uint64))
	case uint:
		return uint64(elem.(uint))
	case uint8:
		return uint64(elem.(uint8))
	case uint16:
		return uint64(elem.(uint16))
	case uint32:
		return uint64(elem.(uint32))
	case uint64:
		return uint64(elem.(uint64))

	case *int:
		return uint64(*elem.(*int))
	case *int8:
		return uint64(*elem.(*int8))
	case *int16:
		return uint64(*elem.(*int16))
	case *int32:
		return uint64(*elem.(*int32))
	case *int64:
		return uint64(*elem.(*int64))
	case int:
		return uint64(elem.(int))
	case int8:
		return uint64(elem.(int8))
	case int16:
		return uint64(elem.(int16))
	case int32:
		return uint64(elem.(int32))
	case int64:
		return uint64(elem.(int64))

	case *float32:
		return uint64(*elem.(*float32))
	case *float64:
		return uint64(*elem.(*float64))
	case float32:
		return uint64(elem.(float32))
	case float64:
		return uint64(elem.(float64))

	}

	return uint64(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToInt uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToInt(elem interface{}) int {

	switch elem.(type) {

	case *uint:
		return int(*elem.(*uint))
	case *uint8:
		return int(*elem.(*uint8))
	case *uint16:
		return int(*elem.(*uint16))
	case *uint32:
		return int(*elem.(*uint32))
	case *uint64:
		return int(*elem.(*uint64))
	case uint:
		return int(elem.(uint))
	case uint8:
		return int(elem.(uint8))
	case uint16:
		return int(elem.(uint16))
	case uint32:
		return int(elem.(uint32))
	case uint64:
		return int(elem.(uint64))

	case *int:
		return int(*elem.(*int))
	case *int8:
		return int(*elem.(*int8))
	case *int16:
		return int(*elem.(*int16))
	case *int32:
		return int(*elem.(*int32))
	case *int64:
		return int(*elem.(*int64))
	case int:
		return int(elem.(int))
	case int8:
		return int(elem.(int8))
	case int16:
		return int(elem.(int16))
	case int32:
		return int(elem.(int32))
	case int64:
		return int(elem.(int64))

	case *float32:
		return int(*elem.(*float32))
	case *float64:
		return int(*elem.(*float64))
	case float32:
		return int(elem.(float32))
	case float64:
		return int(elem.(float64))

	}

	return int(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToInt8 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToInt8(elem interface{}) int8 {

	switch elem.(type) {

	case *uint:
		return int8(*elem.(*uint))
	case *uint8:
		return int8(*elem.(*uint8))
	case *uint16:
		return int8(*elem.(*uint16))
	case *uint32:
		return int8(*elem.(*uint32))
	case *uint64:
		return int8(*elem.(*uint64))
	case uint:
		return int8(elem.(uint))
	case uint8:
		return int8(elem.(uint8))
	case uint16:
		return int8(elem.(uint16))
	case uint32:
		return int8(elem.(uint32))
	case uint64:
		return int8(elem.(uint64))

	case *int:
		return int8(*elem.(*int))
	case *int8:
		return int8(*elem.(*int8))
	case *int16:
		return int8(*elem.(*int16))
	case *int32:
		return int8(*elem.(*int32))
	case *int64:
		return int8(*elem.(*int64))
	case int:
		return int8(elem.(int))
	case int8:
		return int8(elem.(int8))
	case int16:
		return int8(elem.(int16))
	case int32:
		return int8(elem.(int32))
	case int64:
		return int8(elem.(int64))

	case *float32:
		return int8(*elem.(*float32))
	case *float64:
		return int8(*elem.(*float64))
	case float32:
		return int8(elem.(float32))
	case float64:
		return int8(elem.(float64))

	}

	return int8(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToInt16 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToInt16(elem interface{}) int16 {

	switch elem.(type) {

	case *uint:
		return int16(*elem.(*uint))
	case *uint8:
		return int16(*elem.(*uint8))
	case *uint16:
		return int16(*elem.(*uint16))
	case *uint32:
		return int16(*elem.(*uint32))
	case *uint64:
		return int16(*elem.(*uint64))
	case uint:
		return int16(elem.(uint))
	case uint8:
		return int16(elem.(uint8))
	case uint16:
		return int16(elem.(uint16))
	case uint32:
		return int16(elem.(uint32))
	case uint64:
		return int16(elem.(uint64))

	case *int:
		return int16(*elem.(*int))
	case *int8:
		return int16(*elem.(*int8))
	case *int16:
		return int16(*elem.(*int16))
	case *int32:
		return int16(*elem.(*int32))
	case *int64:
		return int16(*elem.(*int64))
	case int:
		return int16(elem.(int))
	case int8:
		return int16(elem.(int8))
	case int16:
		return int16(elem.(int16))
	case int32:
		return int16(elem.(int32))
	case int64:
		return int16(elem.(int64))

	case *float32:
		return int16(*elem.(*float32))
	case *float64:
		return int16(*elem.(*float64))
	case float32:
		return int16(elem.(float32))
	case float64:
		return int16(elem.(float64))

	}

	return int16(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToInt32 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToInt32(elem interface{}) int32 {

	switch elem.(type) {

	case *uint:
		return int32(*elem.(*uint))
	case *uint8:
		return int32(*elem.(*uint8))
	case *uint16:
		return int32(*elem.(*uint16))
	case *uint32:
		return int32(*elem.(*uint32))
	case *uint64:
		return int32(*elem.(*uint64))
	case uint:
		return int32(elem.(uint))
	case uint8:
		return int32(elem.(uint8))
	case uint16:
		return int32(elem.(uint16))
	case uint32:
		return int32(elem.(uint32))
	case uint64:
		return int32(elem.(uint64))

	case *int:
		return int32(*elem.(*int))
	case *int8:
		return int32(*elem.(*int8))
	case *int16:
		return int32(*elem.(*int16))
	case *int32:
		return int32(*elem.(*int32))
	case *int64:
		return int32(*elem.(*int64))
	case int:
		return int32(elem.(int))
	case int8:
		return int32(elem.(int8))
	case int16:
		return int32(elem.(int16))
	case int32:
		return int32(elem.(int32))
	case int64:
		return int32(elem.(int64))

	case *float32:
		return int32(*elem.(*float32))
	case *float64:
		return int32(*elem.(*float64))
	case float32:
		return int32(elem.(float32))
	case float64:
		return int32(elem.(float64))

	}

	return int32(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToInt64 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToInt64(elem interface{}) int64 {

	switch elem.(type) {

	case *uint:
		return int64(*elem.(*uint))
	case *uint8:
		return int64(*elem.(*uint8))
	case *uint16:
		return int64(*elem.(*uint16))
	case *uint32:
		return int64(*elem.(*uint32))
	case *uint64:
		return int64(*elem.(*uint64))
	case uint:
		return int64(elem.(uint))
	case uint8:
		return int64(elem.(uint8))
	case uint16:
		return int64(elem.(uint16))
	case uint32:
		return int64(elem.(uint32))
	case uint64:
		return int64(elem.(uint64))

	case *int:
		return int64(*elem.(*int))
	case *int8:
		return int64(*elem.(*int8))
	case *int16:
		return int64(*elem.(*int16))
	case *int32:
		return int64(*elem.(*int32))
	case *int64:
		return int64(*elem.(*int64))
	case int:
		return int64(elem.(int))
	case int8:
		return int64(elem.(int8))
	case int16:
		return int64(elem.(int16))
	case int32:
		return int64(elem.(int32))
	case int64:
		return int64(elem.(int64))

	case *float32:
		return int64(*elem.(*float32))
	case *float64:
		return int64(*elem.(*float64))
	case float32:
		return int64(elem.(float32))
	case float64:
		return int64(elem.(float64))

	}

	return int64(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToFloat32 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToFloat32(elem interface{}) float32 {

	switch elem.(type) {

	case *uint:
		return float32(*elem.(*uint))
	case *uint8:
		return float32(*elem.(*uint8))
	case *uint16:
		return float32(*elem.(*uint16))
	case *uint32:
		return float32(*elem.(*uint32))
	case *uint64:
		return float32(*elem.(*uint64))
	case uint:
		return float32(elem.(uint))
	case uint8:
		return float32(elem.(uint8))
	case uint16:
		return float32(elem.(uint16))
	case uint32:
		return float32(elem.(uint32))
	case uint64:
		return float32(elem.(uint64))

	case *int:
		return float32(*elem.(*int))
	case *int8:
		return float32(*elem.(*int8))
	case *int16:
		return float32(*elem.(*int16))
	case *int32:
		return float32(*elem.(*int32))
	case *int64:
		return float32(*elem.(*int64))
	case int:
		return float32(elem.(int))
	case int8:
		return float32(elem.(int8))
	case int16:
		return float32(elem.(int16))
	case int32:
		return float32(elem.(int32))
	case int64:
		return float32(elem.(int64))

	case *float32:
		return float32(*elem.(*float32))
	case *float64:
		return float32(*elem.(*float64))
	case float32:
		return float32(elem.(float32))
	case float64:
		return float32(elem.(float64))

	}

	return float32(0)
}

//go:generate goreify github.com/ElPeque/reflect-db/conv.ToFloat64 uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64,float32,float64
func ToFloat64(elem interface{}) float64 {

	switch elem.(type) {

	case *uint:
		return float64(*elem.(*uint))
	case *uint8:
		return float64(*elem.(*uint8))
	case *uint16:
		return float64(*elem.(*uint16))
	case *uint32:
		return float64(*elem.(*uint32))
	case *uint64:
		return float64(*elem.(*uint64))
	case uint:
		return float64(elem.(uint))
	case uint8:
		return float64(elem.(uint8))
	case uint16:
		return float64(elem.(uint16))
	case uint32:
		return float64(elem.(uint32))
	case uint64:
		return float64(elem.(uint64))

	case *int:
		return float64(*elem.(*int))
	case *int8:
		return float64(*elem.(*int8))
	case *int16:
		return float64(*elem.(*int16))
	case *int32:
		return float64(*elem.(*int32))
	case *int64:
		return float64(*elem.(*int64))
	case int:
		return float64(elem.(int))
	case int8:
		return float64(elem.(int8))
	case int16:
		return float64(elem.(int16))
	case int32:
		return float64(elem.(int32))
	case int64:
		return float64(elem.(int64))

	case *float32:
		return float64(*elem.(*float32))
	case *float64:
		return float64(*elem.(*float64))
	case float32:
		return float64(elem.(float32))
	case float64:
		return float64(elem.(float64))

	}

	return float64(0)
}
