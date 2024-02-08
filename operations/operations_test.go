package operations

import "testing"

func TestOperations(t *testing.T) {
	tests := []struct {
		name     string
		operator string
		a, b     float64
		want     float64
		wantErr  bool
	}{
		{"Addition", "+", 3, 5, 8, false},
		{"Subtraction", "-", 10, 4, 6, false},
		{"Multiplication", "*", 7, 8, 56, false},
		{"Division", "/", 10, 2, 5, false},
		{"DivisionByZero", "/", 5, 0, 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var operation Operation[float64]
			switch test.operator {
			case "+":
				operation = &Add[float64]{}
			case "-":
				operation = &Subtract[float64]{}
			case "*":
				operation = &Multiply[float64]{}
			case "/":
				operation = &Divide[float64]{}
			default:
				t.Fatalf("Invalid operator: %s", test.operator)
			}

			got, err := operation.Perform(test.a, test.b)
			if (err != nil) != test.wantErr {
				t.Errorf("Unexpected error status for operation %s: %v", test.operator, err)
			}

			if !test.wantErr && got != test.want {
				t.Errorf("Operation %s failed: got %v, want %v", test.operator, got, test.want)
			}
		})
	}
}
