package validator

import (
	"regexp"
	"reflect"
)

type Validator struct{}

func ( this Validator ) Validate(f interface{}) (bool, []string) {

	var s []string

	r := true

	t := reflect.TypeOf(f)

	v := reflect.ValueOf(f)

	for i := 0; i < t.NumField(); i++ {

		if c := t.Field(i).Tag.Get("validate"); c != "" {

			vr := regexp.MustCompile(";")

			for _, vv := range vr.Split(c, -1) {

				cr := regexp.MustCompile("([A-Z][a-zA-Z]+)\\(?")

				methodMatch := cr.FindStringSubmatch( vv )

				if len( methodMatch ) < 2 {

					continue

				}

				methodName := methodMatch[ 1 ]

				res := callMethod( this , methodName , vv , v.Field( i ).Interface() , t.Field( i ).Type.String() )

				if !res.( bool ) {
					s = append(s, t.Field(i).Name)
					r = false
					break
				}

			}

		}

	}

	return r, s
}

func callMethod( v Validator , methodName string , str string , i interface{}, t string ) interface{}{

	method := reflect.ValueOf( v ).MethodByName( methodName )

	if method.IsValid() {

		return method.Interface().( func( string , interface{} , string ) bool )( str , i , t )

	}

	return true

}