package main

import (
	"fmt"
)

type celcius struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

type fahrenheit struct {
	suhu float64
}

func (c celcius) toCelcius() float64 {
	return c.suhu
}

func (c celcius) toFahrenheit() float64 {
	return c.suhu*9.0/5.0 + 32
}

func (c celcius) toKelvin() float64 {
	return c.suhu
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.suhu
}

func (f fahrenheit) toCelcius() float64 {
	return (f.suhu - 32) * 5.0 / 9.0
}

func (f fahrenheit) toKelvin() float64 {
	return (f.suhu-32)*5.0/9.0 + 273.15
}

func (k kelvin) toKelvin() float64 {
	return k.suhu
}

func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.suhu-273.15)*9.0/5.0 + 32
}

type hitungSuhu interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Masukkan konversi yang ingin digunakan")
	fmt.Printf("1. Celcius ke Fahrenheit\n" +
		"2. Celcius ke Kelvin\n" +
		"3. Kelvin ke Fahrenheit\n" +
		"4. Kelvin ke Celcius\n" +
		"5. Fahrenheit ke Celcius\n" +
		"6. Fahrenheit ke Kelvin\n")

	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)
	for suhuAwal < 1 || suhuAwal > 7 {
		fmt.Println("Pilihan tidak valid. Silahkan masukkan pilihan yang benar")
		fmt.Scanf("%d", &suhuAwal)
	}

	var suhu float64
	fmt.Print("Masukkan suhu yang ingin diubah: ")
	fmt.Scanf("%f", &suhu)

	var interfaceSuhu hitungSuhu
	switch suhuAwal {
	case 1:
		interfaceSuhu = celcius{suhu}
	case 2:
		interfaceSuhu = fahrenheit{suhu}
	case 3:
		interfaceSuhu = kelvin{suhu}
	}

	var suhuFinal float64
	switch suhuAwal {
	case 1:
		suhuFinal = interfaceSuhu.toFahrenheit()
	case 2:
		suhuFinal = interfaceSuhu.toKelvin()
	case 3:
		suhuFinal = interfaceSuhu.toCelcius()
	}
	fmt.Printf("Suhu akhir adalah %.2f", suhuFinal)
}
