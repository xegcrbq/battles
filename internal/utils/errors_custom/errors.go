package errors_custom

import "errors"

var CommandNotFound = errors.New("repository command not found")
var VariableTooLarge = errors.New("variable too large")
var VariableNotFound = errors.New("variable not found")
