package interfaces

import "strconv"

// TOKEN BASE
type Value struct {
	Type  ExpressionType
	Value string
}

// OBTENER VALOR DE SIMBOLO
func (sym Value) getValue() interface{} {
	if sym.Type == INTEGER {
		intVar, _ := strconv.Atoi(sym.Value)
		return intVar
	} else if sym.Type == FLOAT {
		floatVar, _ := strconv.ParseFloat(sym.Value, 64)
		return floatVar
	} else if sym.Type == BOOLEAN {
		return sym.Value == "true"
	} else {
		return sym.Value
	}
}

// OBTENER EL TIPO DE LA VARIABLE
func (sym Value) getType() ExpressionType {
	return sym.Type
}

// ENUMS DE TIPO
type ExpressionType int

const (
	INTEGER ExpressionType = iota
	FLOAT
	STRING
	BOOLEAN
	STRINGCLASS
)

// ENUMS DE OPERACION
type Operation int

const (
	NOOP Operation = iota
	MUL
	DIV
	MOD
	ADD
	SUB
)