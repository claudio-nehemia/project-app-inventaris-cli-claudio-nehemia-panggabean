package utils

import "fmt"

func ValidateNotEmpty(value, fieldName string) error {
    if value == "" {
        return fmt.Errorf("%s cannot be empty", fieldName)
    }
    return nil
}

func ValidateID(id int) error {
    if id <= 0 {
        return fmt.Errorf("ID must be greater than 0")
    }
    return nil
}