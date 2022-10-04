package custom_errs

import "errors"

var CommandNotFound = errors.New("repository command not found")
var VariableTooLarge = errors.New("variable too large")
var VariableNotFound = errors.New("variable not found")
var ParsingError = errors.New("parsing error")
