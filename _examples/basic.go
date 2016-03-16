package main

import (
    "fmt"
    "github.com/george-infinity/validation"
)

var exampleChecks = []validation.Check{
    validation.Check{
        Field: "test",
        Rules: []validation.Rule{
            &validation.IntMin{Min:1},
            &validation.IntMax{Max:100},
        },
    },
    validation.Check{
        Field: "test_string",
        Rules: []validation.Rule{
            &validation.StringMin{Min:1},
            &validation.StringMax{Max:16},
        },
    },
    validation.Check{
        Field: "test_required_string",
        Rules: []validation.Rule{
            &validation.StringMin{Min:1},
            &validation.StringMax{Max:16},
        },
        Required: true,
    },
    validation.Check{
        Field: "test_string_regex",
        Rules: []validation.Rule{
            &validation.StringRegexp{Pattern:`^[a-z]+$`},
        },
    },
    validation.Check{
        Field: "test_string_regex_array",
        Rules: []validation.Rule{
            &validation.StringRegexp{Pattern:`^[a-z]+$`},
        },
        Iterate: true,
    },
}

func main() {
    validator := validation.GetNewValidator()

    foo := make(map[string]interface{})
    foo["test"] = 10
    foo["test_string"] = "some test string"
    foo["test_required_string"] = "some test string"
    foo["test_string_regex"] = "teststring"
    foo["test_string_regex_array"] = []string{"teststring", "teststring"}

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
