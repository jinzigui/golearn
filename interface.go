package main

import (
	"fmt"
)

type ManageElectricityApp interface {
	LeftElectricity() int
}

type Phone struct {
	nElectricity int
	strSystem    string
}

func (phone *Phone) LeftElectricity() int {
	return phone.nElectricity
}

type Camera struct {
	nElectricity int
	strSystem    string
}

func (camera *Camera) LeftElectricity() int {
	return camera.nElectricity
}

func emptyInterFace() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case int:
			fmt.Printf("this is an num = %v\n", v)
		case string:
			fmt.Printf("this is a string =%v\n", v)
		default:
			fmt.Printf("this is not deal with type ,value = %v", v)
		}
	}
	var strmsg string = "this is empty interface test"
	testFunc(strmsg)
	testFunc(233)
	testFunc(3.14)
}

func main() {
	phone := &Phone{68, "IPhone"}
	camera := &Camera{49, "Canon"}
	fmt.Println(phone.LeftElectricity())
	fmt.Println(camera.LeftElectricity())
	app := ManageElectricityApp(phone)
	fmt.Println(app.LeftElectricity())
	var PhoneApp ManageElectricityApp = &Phone{98, "Android"}
	fmt.Println(PhoneApp.LeftElectricity())
	manageElApp := []ManageElectricityApp{phone, camera}
	for n, _ := range manageElApp {
		fmt.Println(manageElApp[n].LeftElectricity())
		switch t := manageElApp[n].(type) {
		case *Phone:
			fmt.Printf("type = %T with value = %v\n", t, t)
		case *Camera:
			fmt.Printf("type = %T with value = %v\n", t, t)
		default:
			fmt.Println("not macth")
		}
	}
	emptyInterFace()
}
