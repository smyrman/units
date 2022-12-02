package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
)

const helpText = `Usage:
  units-codegen[ -o FILE] -t TYPE PKG_CONFIG_FILE

`

type outputVal struct {
	name string
	w    io.Writer
}

func (o outputVal) String() string {
	return o.name
}

func (o *outputVal) Write(p []byte) (n int, err error) {
	if o.w == nil {
		return 0, errors.New("no writer configured")
	}
	return o.w.Write(p)
}

func (o *outputVal) Set(s string) error {
	if len(s) == 0 {
		o.name = "STDOUT"
		o.w = os.Stdout
		return nil
	}
	o.name = s
	fd, err := os.Create(s)
	if err != nil {
		return fmt.Errorf("could not create output file: %s", err.Error())
	}
	o.w = fd
	return nil
}

var options struct {
	output   outputVal
	typeName string
	input    io.Reader
}

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpText)
		flag.PrintDefaults()
	}
	if err := options.output.Set(""); err != nil {
		panic(err)
	}
	flag.Var(&options.output, "o", "Write to file instead of stdout")
	flag.StringVar(&options.typeName, "t", "", "Type to generate code for (required)")
	log.SetPrefix("units-codegen:")
}

func main() {
	mustValidateOptions()
	cfg := mustReadPkgConfig(options.input)
	mustGenerateCode(cfg)
}

func mustValidateOptions() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Println("Expects exactly one argument")
		os.Exit(2)
	}
	if fd, err := os.Open(flag.Arg(0)); err == nil {
		options.input = fd
	} else {
		log.Printf("Error reading input file:\n\t%s", err.Error())
		os.Exit(2)
	}
	if options.typeName == "" {
		log.Println("The type to generate must be specified")
		os.Exit(2)
	}
}

func mustReadPkgConfig(r io.Reader) *pkgConfig {
	cfg := &pkgConfig{}
	dec := json.NewDecoder(options.input)

	if err := dec.Decode(cfg); err != nil {
		log.Printf("Can't decode JSON:\n\t%s", err.Error())
		os.Exit(1)
	}
	if err := cfg.Compile(); err != nil {
		log.Printf("Invalid generator:\n\t%s", err.Error())
		os.Exit(1)
	}
	return cfg
}

func mustGenerateCode(cfg *pkgConfig) {
	t, ok := cfg.Types[options.typeName]
	if !ok {
		log.Printf("No type named '%s'", options.typeName)
		os.Exit(1)
	}
	fileTmpl := template.Must(template.New("generated.go.tmpl").Funcs(tmplFunctions).Parse(tmpl))
	err := fileTmpl.Execute(&options.output, t)
	if err != nil {
		log.Printf("Failed to generate code for type '%s':\n\t%s", t.Name, err.Error())
		os.Exit(1)
	}
}
