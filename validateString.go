package validation

import (
    "errors"
    "reflect"
    "unicode/utf8"
    "regexp"
)

var (
    //ErrStringMin will be returned if StringMin fails validation
    ErrStringMin = errors.New("Value is less than minimum")
    //ErrStringMin will be returned if StringMax fails validation
    ErrStringMax = errors.New("Value is greater than maximum")
    //ErrStringMin will be returned if StringRegexp fails validation
    ErrStringRegexp = errors.New("Regexp pattern does not match")
)

type (
    //StringMin checks that the number of characters in a string is greater than Min
    StringMin struct{
        Min int
    }
    //StringMax checks that the number of characters in a string is less than Max
    StringMax struct{
        Max int
    }
    //StringRegexp checks that the string value matches the Pattern regexp
    //
    //NOTE: The regexp pattern is compiled with regexp.MustCompile and will therefore
    //panic if the supplied regex is invalid
    StringRegexp struct{
        Pattern string
    }
)

//Run handles the validation of StringMin
func (v *StringMin) Run(value interface{}) error {
    if err := checkValueType(value, reflect.String); err != nil {
        return err
    }

    sv := value.(string)

    if utf8.RuneCountInString(sv) < v.Min {
        return ErrStringMin
    }

    return nil
}

//Run handles the validation of StringMax
func (v *StringMax) Run(value interface{}) error {
    if err := checkValueType(value, reflect.String); err != nil {
        return err
    }

    sv := value.(string)

    if utf8.RuneCountInString(sv) > v.Max {
        return ErrStringMax
    }

    return nil
}

//Run handles the validation of StringRegexp
func (v *StringRegexp) Run(value interface{}) error {
    if err := checkValueType(value, reflect.String); err != nil {
        return err
    }

    sv := value.(string)
    bv := []byte(sv)
    r := regexp.MustCompile(v.Pattern)

    if !r.Match(bv) {
        return ErrStringRegexp
    }

    return nil
}
