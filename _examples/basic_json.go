package main

import (
    "fmt"
    "github.com/george-infinity/validation"
)

var jsonChecks = []validation.Check{
    validation.Check{
        Field: "stringField",
        Rules: []validation.Rule{
            &validation.StringMin{Min:1},
            &validation.StringMax{Max:16},
        },
    },
    validation.Check{
        Field: "regexField",
        Rules: []validation.Rule{
            &validation.StringRegexp{`^[a-z]+$`},
        },
    },
    validation.Check{
        Field: "arrayField",
        Rules: []validation.Rule{
            &validation.StringMin{Min:1},
            &validation.StringMax{Max:10},
        },
        Iterate: true,
    },
}

var jsonData = `{"stringField":"test string","regexField":"testregex","arrayField":["test","strings"]}`

func main() {
    validator := validation.GetNewValidator()

    if err := validator.SetJsonInput([]byte(jsonData)); err != nil {
        fmt.Println(err)
        return
    }

    validator.SetChecks(jsonChecks)

    if err := validator.Run(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("No validation errors")
}
