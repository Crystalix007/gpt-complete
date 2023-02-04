# GPT-Complete

A GPT-3 completion program, which runs STDIN through GPT-3, then writes the completion to STDOUT.

## Development

The code requires an API key for OpenAI's API. You can get one [here](https://openai.com/api/).

Then, generate the `lib/api_key.go` from the template with:

```shell
$ go generate ./...
```

## Building

To build:

```shell
$ go build -o complete ./cmd/complete/
```

## Model Parameters

You can tweak the model parameters of this program to make the output more dynamic, or even make the execution cheaper by using a different GPT-3 model.
