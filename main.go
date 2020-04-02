package main

import (
	"fmt"
)

//this is the converter struct Minutes to Seconds
//Minutes to Seconds
//Seconds to Milliseconds
//Centimeter to Feet
//Feet to Centimeter
//Celsius to Fahrenheit
//Fahrenheit to Celsius
//Radian to Degree (π radians is equal to 180°)
//use the math package (just like we did with "fmt" package) from the standard library to access PI (math.Pi) value for better precision.
//Degree to Radian
//Kilogram to Pounds
//Pounds to Kilogram

type converter struct {
	feet, centimeter, minutes, seconds, milliseconds, celsius, fahrenheit, radian, degree, kilogram, pounds float64
}

const Pi = 3.14159265358979323846264338327950288419716939937510582097494459 //

//FeetTo Centimeters method
func (cvr *converter) FeetToCentimeters() float64 {
	return cvr.feet * 30.48

}

//Centimeter To Feets method
func (cvr *converter) CentimetersToFeet() float64 {
	return cvr.centimeter / 30.48

}

//Minutes To Seconds method
func (cvr *converter) MinToSecs() float64 {
	return cvr.seconds * 60

}

// Seconds to Milliseconds
func (cvr *converter) SecsToMill() float64 {
	return cvr.seconds * 1000

}

// Celsius to Fahrenheit (0°C × 9/5) + 32 = 32°F
func (cvr *converter) CelToFah() float64 {
	return cvr.celsius*1.8 + 32

}

// Fahrenheit to Celsius (0°C × 9/5) + 32 = 32°F
func (cvr *converter) FahToCel() float64 {
	return (cvr.fahrenheit - 32) * 0.55

}

//Radian to Degree (π radians is equal to 180°)
func (cvr *converter) RadToDeg() float64 {
	return cvr.radian * 180 / Pi
}

//Degree to Radian
func (cvr *converter) DegToRad() float64 {
	return (cvr.degree * Pi) / 180
}

//Kilogram to Pounds
func (cvr *converter) KilToPnds() float64 {
	return (cvr.kilogram * 2.205)
}

//Pounds to Kilogram
func (cvr *converter) PndsToKil() float64 {
	return (cvr.pounds / 2.205)
}

func main() {

	fmt.Println("Which Unit would you like to covert?...Select")
	fmt.Println("FeetToCentimeters")

	c := converter{feet: 2, centimeter: 2, seconds: 2, celsius: 2, fahrenheit: 2, radian: 2, degree: 2, kilogram: 2, pounds: 2}
	//d := converter{feet: 2, centimeter: 2}
	fmt.Println(" 2 Feet in Ce : ", c.FeetToCentimeters(),
		"\n 2 cen in feet", c.CentimetersToFeet(),
		"\n 2 Minutes in secs: ", c.MinToSecs(),
		"\n 2 seconds in Milliseconds: ", c.SecsToMill(),
		"\n 2 celsius in Fahrenheit: ", c.CelToFah(),
		"\n 2 Fahrenheit in celsius: ", c.FahToCel(),
		"\n 2 Radian in Degree: ", c.RadToDeg(),
		"\n 2 Degree in Radian: ", c.DegToRad(),
		"\n 2 Kilogram in Pounds: ", c.KilToPnds(),
		"\n 2 Pounds in Kilogram: ", c.PndsToKil(),
	)

}
