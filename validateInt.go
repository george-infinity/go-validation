package validation

import (
    "errors"
    "reflect"
)

var (
    //ErrIntMin will be returned if IntMin fails validation
    ErrIntMin = errors.New("Value is less than minimum")
    //ErrIntMax will be returned if IntMax fails validation
    ErrIntMax = errors.New("Value is greater than maximum")
)

type (
    //Int checks that the value is not less than Min and not greater than Max
    Int struct {
        Min int
        Max int
    }
    //IntMin checks that the int value is greater than Min
    IntMin struct{
        Min int
    }
    //IntMax checks that the int value is less than Max
    IntMax struct{
        Max int
    }
)

//Run handles tha validation of Int
func (v *Int) Run(value interface{}) error {
    if err := checkValueType(value, reflect.Int); err != nil {
        return err
    }

    iv := value.(int)

    if iv < v.Min {
        return ErrIntMin
    }

    if iv > v.Max {
        return ErrIntMax
    }

    return nil
}

//Run handles the validation of IntMin
func (v *IntMin) Run(value interface{}) error {
    if err := checkValueType(value, reflect.Int); err != nil {
        return err
    }

    iv := value.(int)

    if iv < v.Min {
        return ErrIntMin
    }

    return nil
}

//Run handles the validation of IntMax
func (v *IntMax) Run(value interface{}) error {
    if err := checkValueType(value, reflect.Int); err != nil {
        return err
    }

    iv := value.(int)

    if iv > v.Max {
        return ErrIntMax
    }

    return nil
}
