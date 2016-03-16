package main

import (
    "fmt"
    "../../validation"
)

var exampleChecks = []validation.Check{
    validation.Check{
        Field: "test",
        Rules: []validation.Rule{
            &validation.IntMin{1},
            &validation.IntMax{100},
        },
    },
    validation.Check{
        Field: "test_string",
        Rules: []validation.Rule{
            &validation.StringMin{1},
            &validation.StringMax{16},
        },
    },
    validation.Check{
        Field: "test_string_regex",
        Rules: []validation.Rule{
            &validation.StringRegexp{`^[a-z]+$`},
        },
    },
    validation.Check{
        Field: "test_string_regex_array",
        Rules: []validation.Rule{
            &validation.StringRegexp{`^[a-z]+$`},
        },
        Iterate: true,
    },
}

func main() {
    validator := validation.GetNewValidator()

    foo := make(map[string]interface{})
    foo["test"] = 10
    foo["test_string"] = "some test string"
    foo["test_string_regex"] = "teststring"
    foo["test_string_regex_array"] = []string{"teststring", "teststring2"}

    validator.SetInput(foo)

    for _, check := range exampleChecks {
        validator.AddCheck(check)
    }

    if err := validator.Run(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("No validation errors")
}
