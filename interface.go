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

func main() {
	phone := &Phone{68, "IPhone"}
	camera := &Camera{49, "Canon"}
	fmt.Println(PhoneApp.LeftElectricity())
	fmt.Println(phone.LeftElectricity())
	fmt.Println(camera.LeftElectricity())
	app := ManageElectricityApp(phone)
	var PhoneApp ManageElectricityApp = &Phone{98, "Android"}
	fmt.Println(app.LeftElectricity())
	manageElApp := []ManageElectricityApp{phone, camera}
	for n, _ := range manageElApp {
		fmt.Println(manageElApp[n].LeftElectricity())
	}
}
