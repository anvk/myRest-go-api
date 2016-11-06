package models

import "reflect"

// taken from http://stackoverflow.com/a/28099719/812519
func IsEmpty(object interface{}) bool {
  //First check normal definitions of empty
  if object == nil {
    return true
  } else if object == "" {
    return true
  } else if object == false {
    return true
  }

  //Then see if it's a struct
  if reflect.ValueOf(object).Kind() == reflect.Struct {
    // and create an empty copy of the struct object to compare against
    empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
    if reflect.DeepEqual(object, empty) {
      return true
    }
  }
  return false
}
