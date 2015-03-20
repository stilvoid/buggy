# Buggy

A bash-friendly JSON parser

## About

Buggy converts JSON into a flat structure with keys that can easily be handled by bash.

For example, the following JSON:

    {
        "hello": [
            {
                "world": "earth"
            },
            "everything"
        ]
    }

Is converted into a set of key-value pairs with the following keys:

* hello.0.world
* hello.1

## Usage

To output a list of the keys:

    cat test.json | buggy

Or to get a specific value:

    cat test.json | buggy "hello.0.world"

Or you can output the keys and values together (this could be tricky to parse depending on the data):

    cat test.json | buggy --values

*Note: When using --values, the values will be converted into bash-friendly form.*

## Examples

A simple way to iterate over the keys and values in some json:

    for key in $(cat test.json | buggy); do
        value=$(cat test.json | buggy "$key")

        echo "The key '$key' has the value '$value'"
    done
