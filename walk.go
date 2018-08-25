package reflect_db

import (
	"fmt"
	"reflect"
)

type walkCallback func(path []string, obj interface{}) bool

func Walk(obj interface{}, atLeaves, while walkCallback) {
	walk(nil, obj, atLeaves, while)
}

func walk(path []string, obj interface{}, atLeaves, while walkCallback) bool {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if while != nil {
		if !while(path, obj) {
			return true
		}
	}

	switch t.Kind() {
	case reflect.Ptr:
		if !v.IsNil() && v.Elem().CanInterface() {
			if cont := walk(path, v.Elem().Interface(), atLeaves, while); !cont {
				return cont
			}
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)

			p := append(path, v.Type().Field(i).Name)
			if f.CanInterface() {
				if cont := walk(p, f.Interface(), atLeaves, while); !cont {
					return cont
				}
			}
		}
	case reflect.Map:
		for _, b := range v.MapKeys() {
			var p = path

			if b.Kind() == reflect.Struct {
				var subp []string
				for i := 0; i < b.NumField(); i++ {
					f := b.Field(i)
					if f.CanInterface() {
						subp = append(subp, b.Type().Field(i).Name, fmt.Sprint(f.Interface()))
					}
				}

				p = append(p, subp...)
			} else {
				p = append(p, fmt.Sprint(b.Interface()))
			}

			if cont := walk(p, v.MapIndex(b).Interface(), atLeaves, while); !cont {
				return cont
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {

			f := v.Index(i)

			p := append(path, fmt.Sprint(i))

			if cont := walk(p, f.Interface(), atLeaves, while); !cont {
				return cont
			}
		}
	case reflect.Uint:
		return atLeaves(path, v.Interface())
	case reflect.Uint16:
		return atLeaves(path, v.Interface())
	case reflect.Uint32:
		return atLeaves(path, v.Interface())
	case reflect.Uint64:
		return atLeaves(path, v.Interface())
	case reflect.Int:
		return atLeaves(path, v.Interface())
	case reflect.Int8:
		return atLeaves(path, v.Interface())
	case reflect.Int16:
		return atLeaves(path, v.Interface())
	case reflect.Int32:
		return atLeaves(path, v.Interface())
	case reflect.Int64:
		return atLeaves(path, v.Interface())
	case reflect.Bool:
		return atLeaves(path, v.Interface())
	case reflect.String:
		return atLeaves(path, v.Interface())
	case reflect.Float32:
		return atLeaves(path, v.Interface())
	case reflect.Float64:
		return atLeaves(path, v.Interface())
	default:
	}
	return true
}

func whilePattern(pattern []string) walkCallback {
	return func(path []string, obj interface{}) bool {
		if len(path) > 0 && (len(path) <= len(pattern)) {
			if path[len(path)-1] != pattern[len(path)-1] && pattern[len(path)-1] != "*" {
				return false
			}
		}
		return true
	}
}

func All(wcs ...walkCallback) walkCallback {
	return func(path []string, obj interface{}) bool {
		for _, wc := range wcs {
			wc(path, obj)
		}
		return true
	}
}

func Or(wcs ...walkCallback) walkCallback {
	return func(path []string, obj interface{}) bool {
		for _, wc := range wcs {
			if wc(path, obj) {
				return true
			}
		}
		return false
	}
}

func And(wcs ...walkCallback) walkCallback {
	return func(path []string, obj interface{}) bool {
		for _, wc := range wcs {
			if !wc(path, obj) {
				return false
			}
		}
		return true
	}
}
