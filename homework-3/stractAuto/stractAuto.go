package stractAuto

import "fmt"

// Создаём структуру легкового автомобиля.
type Car struct {
	Marck			string
	Model			string
	BodyType		string
	Year			int
	Color			string
	Transmission	string
	Trunk			int
	Status			Status
	}

// Создаём структуру грузовика.
type Truck struct {
	Marck			string
	Model			string
	BodyType		string
	Year			int
	Color			string
	Transmission	string
	Carrying		float64
	Status			Status
}

// Доп. структура описывает состояние автомобиля.
type Status struct {
		Signaling		bool
		Engien_run		string
		Window_open		string
		Trunk_volume	int
}

// PrintCar функция выводит на экран описание легкового автомобиля.
func PrintCar(machine Car)  {
	fmt.Println("Легковой автомобиль.")
	if machine.Status.Signaling == false {
		fmt.Println("Сигнализация отключена.")
	}
	if machine.Status.Signaling == true {
		fmt.Println("Сигнализация включена.")
	}
	fmt.Println("Марка автомобиля:")
	fmt.Printf("%s %s, год выпуска %d год, цвет %s, КПП %s, объём багажника %d литров.\n", machine.Marck, machine.Model, machine.Year, machine.Color, machine.Transmission, machine.Trunk)
	fmt.Printf("Двигатель %s, окна %s, багажник заполнен на %d%%.\n", machine.Status.Engien_run, machine.Status.Window_open, machine.Status.Trunk_volume)
}
// PrintTruck функция выводит на экран описание грузового автомобиля.
func PrintTruck(machine Truck)  {
	fmt.Println("Грузовик.")
	if machine.Status.Signaling == false {
		fmt.Println("Сигнализация отключена.")
	}
	if machine.Status.Signaling == true {
		fmt.Println("Сигнализация включена.")
	}
	fmt.Println("Марка автомобиля:")
	fmt.Printf("%s %s, год выпуска %d год, цвет %s, КПП %s, грузоподъёмность %.f тонн.\n", machine.Marck, machine.Model, machine.Year, machine.Color, machine.Transmission, machine.Carrying)
	fmt.Printf("Двигатель %s, окна %s.\n", machine.Status.Engien_run, machine.Status.Window_open)
}

// StractAuto позволяет ввести данные конкретного автомобиля и его состояние.
func StractAuto() {

	auto := Car{
		Marck:			"Renault",
		Model:			"Sandero",
		BodyType: 		"хэтчбек",
		Color: 			"серебристый",
		Transmission: 	"Автомат",
		Year:			2019,
		Trunk:			510,
		Status: Status{
			Signaling:		false,
			Engien_run:		"работает",
			Window_open:	"открыты",
			Trunk_volume:	20,
		},
	}
	PrintCar(auto)

	// Изменяем состояние автомобиля и выводим на экран.
	auto.Status.Trunk_volume = 70
	auto.Status.Signaling = true
	auto.Status.Engien_run = "выключен"
	auto.Status.Window_open = "закрыты"
	PrintCar(auto)

	auto1 := Car{
		Marck:			"Nissan",
		Model:			"X-Trail",
		BodyType: 		"внедорожник",
		Color: 			"чёрный",
		Transmission: 	"Автомат",
		Year:			2020,
		Trunk:			1773,
		Status: Status{
			Signaling:		true,
			Engien_run:		"выключен",
			Window_open:	"закрыты",
			Trunk_volume:	5,
		},
	}
	PrintCar(auto1)

	truck1 := Truck{
		Marck:        "Камаз",
		Model:        "65115",
		BodyType:     "самосвал",
		Year:         2005,
		Transmission: "механика",
		Carrying:     15,
		Status: Status{
			Signaling:   false,
			Engien_run:  "работает",
			Window_open: "открыты",
		},
	}
	PrintTruck(truck1)
}