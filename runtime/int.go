package runtime

import (
	"math"
	"strconv"
)

// Int is the representation of the Integer type. It is equivalent
// to Go's int type (architecture-dependent sized integer).
type Int int

// Int returns the integer value itself.
func (ø Int) Int() int {
	return int(ø)
}

// Float returns a floating point representation of the integer value.
func (ø Int) Float() float64 {
	return float64(ø)
}

// String returns a base 10 string representation of the integer value.
func (ø Int) String() string {
	return strconv.FormatInt(int64(ø), 10)
}

// Bool returns true if the integer value is non-zero, false otherwise.
func (ø Int) Bool() bool {
	return int(ø) != 0
}

// Add performs the addition of the integer value to another Val value, converted
// to an int.
func (ø Int) Add(v Val) Val {
	return Int(int(ø) + v.Int())
}

// Sub performs the subtraction of another Val value, converted
// to an int, from the integer value.
func (ø Int) Sub(v Val) Val {
	return Int(int(ø) - v.Int())
}

// Mul performs the multiplication of the integer value with another Val value,
// converted to an int.
func (ø Int) Mul(v Val) Val {
	return Int(int(ø) * v.Int())
}

// Div performs the division of the integer value by another Val value,
// converted to an int.
func (ø Int) Div(v Val) Val {
	return Int(int(ø) / v.Int())
}

// Mod returns the modulo (remainder) of the division of the integer value by
// another Val value, converted to an int.
func (ø Int) Mod(v Val) Val {
	return Int(int(ø) % v.Int())
}

// Pow returns the integer raised at the power of the other Val value, converted
// to an int.
func (ø Int) Pow(v Val) Val {
	return Int(math.Pow(float64(ø), float64(v.Int())))
}

// Not returns the binary not operation applied to the integer value.
// It switches all bits of the value.
func (ø Int) Not() Val {
	return Int(^int(ø))
}

// Unm returns the unary minus operation applied to the integer value.
// It switches the sign of the value.
func (ø Int) Unm() Val {
	return Int(-int(ø))
}
