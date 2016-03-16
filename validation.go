//Package validation implements methods for checking data complies with given rules
package validation

import (
    "errors"
    "reflect"
    "fmt"
    "encoding/json"
)

var (
    //ErrNonSliceIterate will be returned when the Iterate flag is used on a non-slice data type
    ErrNonSliceIterate = errors.New("Attempting to iterate over a non-slice value")
)

var errTypeMismatch = "Unexpected value type %s expecting %s"
var errFailedValidation = "Field %s failed validation: %s"
var errSliceFailedValidation = "Field %s[%d] failed validation: %s"

//Rule is the interface that wraps the Run method
//
//The Run method accepts any data type as a value and returns an error as the result
//of the validation
type Rule interface {
    Run(value interface{}) error
}

//Check stores informatiom about a single field check
//
//Field is the name of the field being validated
//
//Iterate is a flag that indicates if the value is a slice and we should loop
//through the values validating them individually
//
//Required is a flag that indicates if the value is a required field in the input data
//
//Rules is a slice of Rule that describes how the value should be validated
type Check struct {
    Field    string
    Iterate  bool
    Required bool
    Rules    []Rule
}

type validator struct {
    checks []Check
    data map[string]interface{}
}

//GetNewValidator returns a new validator
func GetNewValidator() validator {
    return validator{}
}

//AddCheck adds a single Check to the validator
func (v *validator) AddCheck(check Check) {
    v.checks = append(v.checks, check)
}

//SetChecks sets a slice of Check in the validator
func (v *validator) SetChecks(checks []Check) {
    v.checks = checks
}

//SetInput sets the input data to be validated
func (v *validator) SetInput(input map[string]interface{}) {
    v.data = input
}

//SetJsonInput is a helper method to set the input data from a JSON string
func (v *validator) SetJsonInput(jsonInput []byte) error {
    data := make(map[string]interface{})
    if err := json.Unmarshal(jsonInput, &data); err != nil {
        return err
    }

    v.SetInput(data)

    return nil
}

//Run runs the validator returning the first error found or nil
func (v *validator) Run() error {
    var err error

    for _, check := range v.checks {
        if check.Required {
            if _, exists := v.data[check.Field]; !exists {
                return errors.New(fmt.Sprintf(errFailedValidation, check.Field, "Required field missing"))
            }
        }

        value := v.data[check.Field]

        if check.Iterate {
            if reflect.TypeOf(value).Kind() != reflect.Slice {
                return ErrNonSliceIterate
            }

            s := reflect.ValueOf(value)
            for i := 0; i < s.Len(); i++ {
                for _, rule := range check.Rules {
                    if err = rule.Run(s.Index(i).Interface()); err != nil {
                        return errors.New(fmt.Sprintf(errSliceFailedValidation, check.Field, i, err.Error()))
                    }
                }
            }
        } else {
            for _, rule := range check.Rules {
                if err = rule.Run(value); err != nil {
                    return errors.New(fmt.Sprintf(errFailedValidation, check.Field, err.Error()))
                }
            }
        }
    }

    return nil
}

//checkValueType checks the data type of the interface matches the expected data type
func checkValueType(value interface{}, expectedType reflect.Kind) error {
    if reflect.TypeOf(value).Kind() != expectedType {
        return errors.New(fmt.Sprintf(errTypeMismatch, reflect.TypeOf(value).Kind().String(), expectedType.String()))
    }

    return nil
}
