package stractAuto

import "fmt"

type Car struct {
	Marck			string
	Model			string
	BodyType		string
	Year			int
	CarColor		string
	Transmission	string
	Trunk			int
	Status			Status
	}
	

type Status struct {
		Signaling		bool
		Engien_run		string
		Window_open		string
		Trunk_volume	int
}

func PrintCar(machine Car)  {

	if machine.Status.Signaling == false {
		fmt.Println("Сигнализация отключена.")
	}
	fmt.Println("Сигнализация включена.")
	fmt.Println("Марка автомобиля:")
	fmt.Printf("%s %s, год выпуска %d год, цвет %s, КПП %s, объём багажника %d литров.\n", machine.Marck, machine.Model, machine.Year, machine.CarColor, machine.Transmission, machine.Trunk)
	fmt.Printf("Двигатель %s, окна %s, багажник заполнен на %d%%.\n", machine.Status.Engien_run, machine.Status.Window_open, machine.Status.Trunk_volume)

}


func stractAuto() {

	auto := Car{
		Marck:	"Renault",
		Model:	"Sandero",
		BodyType: "хэтчбек",
		CarColor: "серебристый",
		Transmission: "Автомат",
		Year:	2019,
		Trunk:	510,
		Status: Status{
			Signaling:		false,
			Engien_run:		"работает",
			Window_open:	"открыты",
			Trunk_volume:	20,
		},
	}

	PrintCar(auto)

	auto.Status.Trunk_volume = 70
	auto.Status.Signaling = true
	auto.Status.Engien_run = "выключен"
	auto.Status.Window_open = "закрыты"

	PrintCar(auto)

	auto1 := Car{
		Marck:	"Nissan",
		Model:	"X-Trail",
		BodyType: "внедорожник",
		CarColor: "чёрный",
		Transmission: "Автомат",
		Year:	2020,
		Trunk:	1773,
		Status: Status{
			Signaling:		true,
			Engien_run:		"выключен",
			Window_open:	"закрыты",
			Trunk_volume:	5,
		},
	}

	PrintCar(auto1)



}
