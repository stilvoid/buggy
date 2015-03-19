package main

import (
    "encoding/json"
    "fmt"
    "os"
)

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
        out[prefix[1:]] = vv
    case int:
        out[prefix[1:]] = vv
    case bool:
        out[prefix[1:]] = vv
    default:
        fmt.Fprintln(os.Stderr, "Input appears to be invalid json")
        os.Exit(1)
    }
}

func main() {
    var in interface{}

    out := make(map[string]interface{})

    dec := json.NewDecoder(os.Stdin)

    dec.Decode(&in)

    parseObj(in, out, "")

    if len(os.Args) > 1 {
        key := os.Args[1]

        if value, ok := out[key]; ok {
            fmt.Println(value)
        }
    } else {
        for key, value := range out {
            fmt.Printf("%s=%v\n", key, value)
        }
    }
}
