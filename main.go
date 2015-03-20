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
    case float64:
        out[prefix[1:]] = vv
    case bool:
        out[prefix[1:]] = vv
    case nil:
        out[prefix[1:]] = vv
    default:
        fmt.Fprintln(os.Stderr, "Input appears to be invalid json", vv)
        os.Exit(1)
    }
}

func main() {
    var in interface{}
    out := make(map[string]interface{})

    json.NewDecoder(os.Stdin).Decode(&in)

    parseObj(in, out, "")

    if len(os.Args) > 1 {
        key := os.Args[1]

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
            fmt.Printf("%s=", key)

            switch vv := value.(type) {
            case string:
                fmt.Printf("\"%s\"\n", vv)
            case float64:
                fmt.Println(vv)
            case bool:
                fmt.Println(vv)
            case nil:
                fmt.Println()
            }
        }
    }
}
