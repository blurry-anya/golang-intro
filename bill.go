package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// to create new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// receiver function
func (b *bill) format() string {
	formattedStr := "Bill breakdown:\n\n"
	var total float64 = 0

	formattedStr += fmt.Sprintf("\t%v\n\n", b.name)

	for key, value := range b.items {
		formattedStr += fmt.Sprintf("%-20v ...$%v\n", key+":", value)
		total += value
	}

	formattedStr += fmt.Sprintf("\n%-20v ...$%0.2f\n", "Tip:", b.tip)

	formattedStr += fmt.Sprintf("\n%-20v ...$%0.2f", "Total:", total+b.tip)

	return formattedStr
}

func (b *bill) setTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to 'bills/..'")
}
