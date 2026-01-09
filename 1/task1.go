package main

import "fmt"

type PepeSchnele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{Speed: speed, Charisma: charisma, Wisdom: wisdom}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func main() {
	pepe1 := NewPepeSchnele(10, 15, 20)
	rating1 := pepe1.GetRating()
	fmt.Printf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d\n", pepe1.Speed, pepe1.Charisma, pepe1.Wisdom, rating1)

	pepe2 := NewPepeSchnele(12, 10, 25)
	rating2 := pepe2.GetRating()
	fmt.Printf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d\n", pepe2.Speed, pepe2.Charisma, pepe2.Wisdom, rating2)

	if rating1 > rating2 {
		fmt.Println("Первый Пепе Шнеле имеет более высокий рейтинг.")
	} else if rating2 > rating1 {
		fmt.Println("Второй Пепе Шнеле имеет более высокий рейтинг.")
	} else {
		fmt.Println("Рейтинги равны.")
	}
}
