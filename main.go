package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "strings"
)

var replacer *strings.Replacer = strings.NewReplacer("\\", "\\\\", "'", "\\'", "\n", "\\n", "\t", "\\t")

// TODO: Make a map struct and turn this into a method
func apply(value interface{}, out map[string]string, prefix string) {
    key := replacer.Replace(prefix[1:])

    out[key] = fmt.Sprint(value)
}

func parseObj(in interface{}, out map[string]string, prefix string) {
    switch vv := in.(type) {
    case map[string]interface{}:
        for key, value := range vv {
            parseObj(value, out, fmt.Sprintf("%s.%s", prefix, key))
        }
    case []interface{}:
        for index, value := range vv {
            parseObj(value, out, fmt.Sprintf("%s.%d", prefix, index))
        }
    case string, float64, bool:
        apply(vv, out, prefix)
    case nil:
        apply("", out, prefix)
    default:
        fmt.Fprintln(os.Stderr, "Input appears to be invalid json", vv)
        os.Exit(1)
    }
}

// TODO: Tests
func printMap(obj map[string]string, output_values bool) {
    for key, value := range obj {
        if output_values {
            fmt.Printf("%s=%s\n", key, replacer.Replace(value))
        } else {
            fmt.Println(key)
        }
    }
}

// TODO: Tests
func printValue(obj map[string]string, key string) {
    if value, ok := obj[key]; ok {
        fmt.Println(value)
    } else {
        fmt.Fprintf(os.Stderr, "'%s' is not present\n", key)
        os.Exit(1)
    }
}

// TODO: Tests
func main() {
    var in interface{}
    out := make(map[string]string)

    output_values := flag.Bool("values", false, "Output values (the default is just to output the keys)")
    flag.Parse()

    json.NewDecoder(os.Stdin).Decode(&in)

    if in == nil {
        fmt.Fprintln(os.Stderr, "Input appears to be invalid json")
        os.Exit(1)
    }

    parseObj(in, out, "")

    if flag.NArg() >= 1 {
        // TODO: all the args
        printValue(out, flag.Arg(0))
    } else {
        printMap(out, *output_values)
    }
}
