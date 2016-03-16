#go-validation

A simple validation library for basic data structures like JSON

## Docs

[https://godoc.org/github.com/george-infinity/go-validation](https://godoc.org/github.com/george-infinity/go-validation)

## Usage

    import "github.com/george-infinity/go-validation"

    var checks = []validation.Check{
        validation.Check{
            Field: "foo",
            Rules: []validation.Rule{
                &validation.IntMin{1},
                &validation.IntMax{100},
            },
        },
        validation.Check{
            Field: "bar",
            Rules: []validation.Rule{
                &validation.StringMin{1},
                &validation.StringMax{16},
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
