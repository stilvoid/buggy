package main

import "testing"

func compareMap(t *testing.T, actual map[string]string, expected map[string]string) {
    fail := func() {
        t.Errorf("Mismatch:\n%v\n%v\n", actual, expected)
    }

    if len(actual) != len(expected) {
        fail()
    }

    for key, value := range(actual) {
        if expected[key] != value {
            fail()
        }
    }
}

func TestReplacer(t *testing.T) {
    input := "\\String\nwith \"quotes\" 'and'\n\ttabs"
    expected := "\\\\String\\nwith \"quotes\" \\'and\\'\\n\\ttabs"

    out := replacer.Replace(input)

    if out != expected {
        t.Errorf("Mismatch:\n%s\n%s\n", out, expected)
    }
}

func TestApply(t *testing.T) {
    out := make(map[string]string)
    expected := map[string]string {
        "key": "value",
        "repl\\'aced": "not\nreplaced",
    }

    apply("value", out, ".key")
    apply("not\nreplaced", out, ".repl'aced")

    compareMap(t, out, expected)
}

func TestParseObj_string(t *testing.T) {
    out := make(map[string]string)
    expected := map[string]string {
        "key": "value",
    }

    parseObj("value", out, ".key")

    compareMap(t, out, expected)
}

func TestParseObj_number(t *testing.T) {
    out := make(map[string]string)
    expected := map[string]string {
        "number": "17",
    }

    parseObj(float64(17), out, ".number")

    compareMap(t, out, expected)
}

func TestParseObj_bool(t *testing.T) {
    out := make(map[string]string)
    expected := map[string]string {
        "true": "true",
        "false": "false",
    }

    parseObj(true, out, ".true")
    parseObj(false, out, ".false")

    compareMap(t, out, expected)
}

func TestParseObj_nil(t *testing.T) {
    out := make(map[string]string)
    expected := map[string]string {
        "nil": "",
    }

    parseObj(nil, out, ".nil")

    compareMap(t, out, expected)
}

func TestParseObj_list(t *testing.T) {
    out := make(map[string]string)
    expected := map[string]string {
        "0": "first",
        "1": "2",
        "2": "true",
    }

    parseObj([]interface{}{"first", float64(2), true}, out, "")

    compareMap(t, out, expected)
}

func TestParseObj_map(t *testing.T) {
    out := make(map[string]string)
    in := map[string]interface{} {
        "one": "one",
        "two": float64(2),
    }
    expected := map[string]string {
        "one": "one",
        "two": "2",
    }

    parseObj(in, out, "")

    compareMap(t, out, expected)
}

func TestNesting(t *testing.T) {
    out := make(map[string]string)
    in := map[string]interface{} {
        "root": "one",
        "list": []interface{} {
            "hello",
            map[string]interface{} {
                "world": "earth",
            },
        },
        "map": map[string]interface{} {
            "north": "N",
            "others": []interface{} {
                "east",
                "west",
                "south",
                false,
            },
        },
    }
    expected := map[string]string {
        "root": "one",
        "list.0": "hello",
        "list.1.world": "earth",
        "map.north": "N",
        "map.others.0": "east",
        "map.others.1": "west",
        "map.others.2": "south",
        "map.others.3": "false",
    }

    parseObj(in, out, "")

    compareMap(t, out, expected)
}
