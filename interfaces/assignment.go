package interfaces

import "fmt"

// ASIGNACION
type Assignment struct {
	Instruction
	Id         string
	Expression *Expression
}

// *INSTRUCTION
func (assign Assignment) Execute(scope Scope) {
	// OBTENER VARIABLES
	scopeVar := scope.GetVariable(assign.Id)
	expValue := assign.Expression.GetValue(scope)

	// VERIFICAR TIPOS
	if scopeVar.GetType(scope) != UNDEF {
		if scopeVar.GetType(scope) == expValue.GetType(scope) {
			if scopeVar.(ValueMut).Mut {
				scope.SetVariable(assign.Id, expValue)
			} else {
				Errors = append(Errors, Error{
					fmt.Sprintf("The variable %s is not mutable", assign.Id),
					expValue.GetLine(),
					expValue.GetColumn(),
				})
			}
		} else {
			Errors = append(Errors, Error{
				fmt.Sprintf("Cannot assign type %s to %s", expValue.GetType(scope).String(), scopeVar.GetType(scope).String()),
				expValue.GetLine(),
				expValue.GetColumn(),
			})
		}
	} else {
		Errors = append(Errors, Error{
			fmt.Sprintf("The variable \"%s\" is not declared", assign.Id),
			expValue.GetLine(),
			expValue.GetColumn(),
		})
	}
}
