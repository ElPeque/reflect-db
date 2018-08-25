package main

import (
	"fmt"
	"reflect"
	"strings"
)

type cosa struct {
	uno      int
	Dos      int
	Tres     int
	Flotinga float64
}

type rarita struct {
	Reloco map[cosa]string
}

type cosita struct {
	Nombre            string
	Cosas             []*cosa
	Numeritos         []int
	NumeritosPointers []*int
	Weird             *rarita
}

type cosillas struct {
	Mapa *map[string]*cosita
	Pope *map[string]cosita
	algo string
	Otra int
}

var datos = cosillas{
	algo: "algoval",
	Otra: 5,
	Mapa: &map[string]*cosita{
		"ta": &cosita{
			Numeritos:         []int{100, 101, 102},
			NumeritosPointers: []*int{new(int), new(int), new(int)},
			Weird: &rarita{
				Reloco: map[cosa]string{
					cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					}: "omfg",
				},
			},
			Nombre: "unacosita",
			Cosas: []*cosa{
				&cosa{
					uno:  1,
					Dos:  2,
					Tres: 3,
				},
			},
		},
		"te": &cosita{

			Nombre: "otracosita",
			Cosas: []*cosa{
				&cosa{
					uno:      1,
					Dos:      2,
					Tres:     3,
					Flotinga: 3.14,
				},
				&cosa{
					uno:  4,
					Dos:  5,
					Tres: 6,
				},
				&cosa{
					uno:  5,
					Dos:  6,
					Tres: 7,
				},
			},
		},
		"ti": &cosita{
			Nombre: "yotramas",
			Cosas: []*cosa{
				&cosa{
					uno:  1,
					Dos:  2,
					Tres: 3,
				},
				&cosa{
					uno:  0,
					Dos:  9,
					Tres: 8,
				},
			},
		},
	},
	Pope: &map[string]cosita{
		"ta": cosita{
			Weird: &rarita{
				Reloco: map[cosa]string{
					cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					}: "omfg",
				},
			},
			Nombre: "unacosita",
			Cosas: []*cosa{
				&cosa{
					uno:  1,
					Dos:  2,
					Tres: 3,
				},
			},
		},
		"te": cosita{

			Nombre: "otracosita",
			Cosas: []*cosa{
				&cosa{
					uno:  1,
					Dos:  2,
					Tres: 3,
				},
				&cosa{
					uno:  4,
					Dos:  5,
					Tres: 6,
				},
				&cosa{
					uno:  5,
					Dos:  6,
					Tres: 7,
				},
			},
		},
		"ti": cosita{
			Nombre: "yotramas",
			Cosas: []*cosa{
				&cosa{
					uno:  1,
					Dos:  2,
					Tres: 3,
				},
				&cosa{
					uno:  0,
					Dos:  9,
					Tres: 8,
				},
			},
		},
	},
}

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

func leafPrintPath() walkCallback {
	return func(path []string, obj interface{}) bool {
		fmt.Println(strings.Join(append(path, fmt.Sprint(obj)), "."))
		return true
	}
}

func leafSum(sum *int) walkCallback {
	return func(path []string, obj interface{}) bool {
		val, ok := obj.(int)
		if ok {
			*sum += val
		}
		return true
	}
}

func leafAvg(avgIface interface{}) walkCallback {
	var count int
	return func(path []string, obj interface{}) bool {

		switch avgIface.(type) {
		case *uint:
		case *uint16:
		case *uint32:
		case *uint64:
		case *int:
			ptr := avgIface.(*int)
			val, ok := obj.(int)
			if ok {
				*ptr = (*ptr*count + val) / (count + 1)
				count++
				break
			}
			pVal, ok := obj.(*int)
			if ok {
				*ptr = (*ptr*count + *pVal) / (count + 1)
				count++
				break
			}
		case *int8:
		case *int16:
		case *int32:
		case *int64:
		case *float32:
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

func leafAccumulate(avg *int) walkCallback {
	var vals []int

	return func(path []string, obj interface{}) bool {
		val, ok := obj.(int)
		if ok {
			vals = append(vals, val)
		}
		return true
	}
}

func main() {
	sum := 0
	avg := 0
	avgF := float64(0.0)

	fmt.Println("avgF is at:", &avgF)

	Walk(
		&datos,
		All(leafPrintPath(), leafSum(&sum), leafAvg(&avg), leafAvg(&avgF)),
		whilePattern([]string{"Mapa", "*", "Cosas", "0", "*"}),
	)

	fmt.Println("Sum:", sum)
	fmt.Println("Avg:", avg)
	fmt.Println("AvgF:", avgF)

}
