package model

import "fmt"

type Animal interface {
	Attack() string
}

type Dog struct {
	Color      string
	Limbs      int
	AttackType string
}

func (d Dog) Attack() string {
	attackPoints := 40
	return fmt.Sprintf("Attack type: %s\nAttack points: %d", d.AttackType, attackPoints)
}
