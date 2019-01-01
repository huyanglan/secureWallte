package test

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

type Book struct {
	ID     int
	Title  string
	Year   int
}

func process(in interface{}, isSlice bool, isMap bool) {
	v := reflect.ValueOf(in)

	if isSlice {
		for i := 0; i < v.Len(); i++ {
			strct := v.Index(i).Interface()
			log.Println("------struct----", strct)
		}
		return
	}

	if isMap {
		fmt.Printf("Type: ----------%v\n", v)     // map[]
		//for _, s := range v {           // Error: cannot range over v (type reflect.Value)
		//	fmt.Printf("Value: %v\n", s.Interface())
		//}
	}
}

func TestInterface(t *testing.T) {
	b := Book{}
	b.Title = "Learn Go Language"
	b.Year = 2018
	m := make(map[string]*Book)
	m["1"] = &b

	process(m, false, true)
}

func TestInterfaceMearsure(t *testing.T) {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}
	//The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	Measure(r)
	Measure(c)
}

func TestInterfaceToMap(t *testing.T) {
	acct := map[string]string{
		"accountStartDate": "openDate",
		"cancleDate": "closeDate",
	}
	js, _ := json.Marshal(acct)
	var m map[string]interface{}
	if err := json.Unmarshal(js, &m); err != nil {

	}
	InterfaceToMap(m)
}

func InterfaceToMap(m map[string]interface{}) {
	for _, v := range m {
		fmt.Println("-----------------", m)
		if v, ok := v.(map[string]interface{}); ok {
			fmt.Println("--[---v---]--", v)
		}
	}
}

func TestMd5(t *testing.T) {
	ts := GetMD5Hash("12345678")
	log.Println("sss", ts)
}
func GetMD5Hash(_text string) string {
	hasher := md5.New()
	hasher.Write([]byte(_text))
	return hex.EncodeToString(hasher.Sum(nil))
}
