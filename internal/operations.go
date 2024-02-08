package internal

import "errors"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type NumberType[T Number] struct {
	Value T
}

type Operation[T Number] interface {
	Perform(a, b T) (T, error)
}

type Add[T Number] struct{}

func (Add[T]) Perform(a, b T) (T, error) {
	return a + b, nil
}

type Subtract[T Number] struct{}

func (Subtract[T]) Perform(a, b T) (T, error) {
	return a - b, nil
}

type Multiply[T Number] struct{}

func (Multiply[T]) Perform(a, b T) (T, error) {
	return a * b, nil
}

type Divide[T Number] struct{}

func (Divide[T]) Perform(a, b T) (T, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func GetOperation[T Number](operator rune) Operation[T] {
	switch operator {
	case '+':
		return &Add[T]{}
	case '-':
		return &Subtract[T]{}
	case '*':
		return &Multiply[T]{}
	case '/':
		return &Divide[T]{}
	default:
		return nil
	}
}
