package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "strings"
)

var replacer *strings.Replacer

func apply(out map[string]interface{}, prefix string, value interface{}) {
    out[replacer.Replace(prefix[1:])] = value
}

func parseObj(in interface{}, out map[string]interface{}, prefix string) {
    switch vv := in.(type) {
    case map[string]interface{}:
        for key, value := range vv {
            parseObj(value, out, fmt.Sprintf("%s.%s", prefix, key))
        }
    case []interface{}:
        for index, value := range vv {
            parseObj(value, out, fmt.Sprintf("%s.%d", prefix, index))
        }
    case string:
        apply(out, prefix, vv)
    case float64:
        apply(out, prefix, vv)
    case bool:
        apply(out, prefix, vv)
    case nil:
        apply(out, prefix, vv)
    default:
        fmt.Fprintln(os.Stderr, "Input appears to be invalid json", vv)
        os.Exit(1)
    }
}

func main() {
    var in interface{}
    out := make(map[string]interface{})

    output_values := flag.Bool("values", false, "Output values (the default is just to output the keys)")
    flag.Parse()

    json.NewDecoder(os.Stdin).Decode(&in)

    replacer = strings.NewReplacer("\\", "\\\\", "'", "\\'", "\n", "\\n", "\t", "\\t")

    parseObj(in, out, "")

    if flag.NArg() >= 1 {
        key := flag.Arg(0)

        if value, ok := out[key]; ok {
            if value == nil {
                fmt.Println()
            } else {
                fmt.Println(value)
            }
        } else {
            fmt.Fprintf(os.Stderr, "'%s' is not present\n", key)
            os.Exit(1)
        }
    } else {
        for key, value := range out {
            fmt.Print(key)

            if !*output_values {
                fmt.Println()
            } else {
                fmt.Print("=")

                switch vv := value.(type) {
                case nil:
                    fmt.Println()
                case string:
                    fmt.Printf("$'%s'\n", replacer.Replace(vv))
                default:
                    fmt.Println(vv)
                }
            }
        }
    }
}
