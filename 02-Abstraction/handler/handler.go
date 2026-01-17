package handler

import (
	"fmt"

	"github.com/devopsguyXD/test/02-Abstraction/model"
)

func NewCreature(color string, limbs int, attackType string) model.Animal {
	return &model.Dog{
		Color:      color,
		Limbs:      limbs,
		AttackType: attackType,
	}
}

func Attack(d model.Animal) {
	fmt.Println(d.Attack())
}
