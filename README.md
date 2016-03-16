#go-validation

A simple validation library for basic data structures like JSON

## Docs

[https://godoc.org/github.com/george-infinity/validation](https://godoc.org/github.com/george-infinity/validation)

## Usage

    import "github.com/george-infinity/validation"

    var checks = []validation.Check{
        validation.Check{
            Field: "foo",
            Rules: []validation.Rule{
                &validation.IntMin{Min:1},
                &validation.IntMax{Max:100},
            },
        },
        validation.Check{
            Field: "bar",
            Rules: []validation.Rule{
                &validation.StringMin{Min:1},
                &validation.StringMax{Max:16},
            },
        },
    }

    func main() {
        ...
        validator := validation.GetNewValidator()

        data := make(map[string]interface{})
        data["foo"] = 10
        data["bar"] = "foobar"

        validator.SetInput(data)
        validator.SetChecks(checks)

        err := validator.Run()
        ...
    }
