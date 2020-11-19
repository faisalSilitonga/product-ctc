package entity

import "errors"

//ErrParseID error parse id
var ErrParseID = errors.New("Cannot Parse ID")

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")
