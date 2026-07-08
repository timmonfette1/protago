# protago

## What Is It?
`protago` is a Protobuf/Go toolchain that allows for a seamless integration of Protobuf `FieldOptions` and common Go `struct` tags.  Inspired by a problem faced in my day to day job, I designed `protago` to facilitate a single-struct approach to data models within a Go application.

Go uses struct tags for a variety of things such as, but not limited to:
- Struct validation using the `validate` tag and the [validator](https://github.com/go-playground/validator) package.
- Automatic serialization of a struct into a database
  - `bson` for serializing a [struct](https://pkg.go.dev/go.mongodb.org/mongo-driver/v2@v2.7.0/bson#hdr-Structs) into MongoDB with options.

If you're using protobuf as your service/model definition language, it can be useful to apply these struct tags when compiling the protos so you can use the generated structs as the singular data model through the entire application/service layer without the need to struct hop.

While other tools and protoc plugins exist to apply tags to the generated structs, most of them will require you to add the tags a single `FieldOption`
```proto
message Hello {
    string name = 1 [(annotations.tags) = "bson:\"name\" validate:\"required\""];
}
```
which can limit your ability to access these options within code via `protoreflect` as you'd need to parse the individual tags out of the singular tag-string which can sometimes be less intuitive than just accessing them directly.

This is why `protago` provides explicit `FieldOptions` for common struct tags so that you can both reliably add the tags to the generated structs, while also maintaining the flexability of easily accessing the values via `protoreflect`.

## Available Tooling
Included with `protago` is:
- Protobuf files for struct tags that you can import into your own protos for use.
- A protoc plugin that will convert all of the `FieldOptions` to their appropriate struct tags and set them on the generated structs.
- Packages of helpful functions for common patterns that might interact with these tags (ie. using the `bson` options for converting a `proto.Message` into a MongoDB partial update statement to pass to a `$set` operator).

## Quickstart
If you'd like to use the full suite of `protago` tooling, then you can both install the protoc plugin and the main package together.
```sh-session
$ go install github.com/timmonfette1/protago/cmd/protoc-gen-protago
$ go get github.com/timmonfette1/protago
```

If you prefer to use `mise` for tool management, then you can install the protoc plugin as a Go tool

```toml
[tools]
"go:github.com/timmonfette1/protago/cmd/protoc-gen-protago" = { version = "latest" }
```

Checkout the [examples](./examples) for more detailed examples around using the `FieldOptions` in your protos or how to utilize the include Go modules for working with the `FieldOptions` and other useful functions.

## Building this Repository
This project uses [mise](https://mise.jdx.dev/) for tool management as well as for building the project.

Simply clone this repository and run
```sh-session
$ mise run build
```
to build both the protos as well as all of the Go code.

If you'd like to just compile the protos, you can simply run
```sh-session
$ mise run bufbuild
```
to build only the proto files.