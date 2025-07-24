package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type ProcInout struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Env    map[string]string
}

func NewProcInout() *ProcInout {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		env[parts[0]] = parts[1]
	}

	return &ProcInout{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Env:    env,
	}
}

type Command func(args []string, inout *ProcInout) int

func Run(c Command) {
	args := os.Args[1:]
	exitStatus := c(args, NewProcInout())
	os.Exit(exitStatus)
}

type FlagType string

const (
	FlagTypeString FlagType = "string"
	FlagTypeBool   FlagType = "bool"
	FlagTypeInt    FlagType = "int"
)

type Flag struct {
	FieldName   string
	Name        string
	Description string
	Type        FlagType
}

func AnalyzeFlags(opts any) []Flag {
	flags := make([]Flag, 0)
	t := reflect.TypeOf(opts).Elem()
	m := t.NumField()
	for i := 0; i < m; i++ {
		field := t.Field(i)
		name := field.Tag.Get("name")
		f := Flag{
			FieldName:   field.Name,
			Name:        name,
			Description: field.Tag.Get("description"),
		}
		switch field.Type.Kind() {
		case reflect.String:
			f.Type = FlagTypeString
		case reflect.Bool:
			f.Type = FlagTypeBool
		case reflect.Int:
			f.Type = FlagTypeInt
		default:
			panic("unhandled default case")
		}
		flags = append(flags, f)
	}
	return flags
}

func ParseFlags(name string, args []string, opts any) error {
	fs := AnalyzeFlags(opts)

	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	for _, fl := range fs {
		f := reflect.ValueOf(opts).Elem().FieldByName(fl.FieldName)

		switch fl.Type {
		case FlagTypeString:
			flags.StringVar(f.Addr().Interface().(*string), fl.Name, f.String(), fl.Description)
		case FlagTypeBool:
			flags.BoolVar(f.Addr().Interface().(*bool), fl.Name, f.Bool(), fl.Description)
		case FlagTypeInt:
			flags.IntVar(f.Addr().Interface().(*int), fl.Name, int(f.Int()), fl.Description)
		}
	}
	return flags.Parse(args)
}

func NewCommandWithCompletion(c Command, compFun func(args []string) []string) Command {
	return func(args []string, inout *ProcInout) int {
		if inout.Env["GO_FLAGS_COMPLETION"] == "1" {
			completions := compFun(args)
			fmt.Fprintln(inout.Stdout, strings.Join(completions, "\n"))
			return 0
		}
		return c(args, inout)
	}
}

func NewCompletionByFlags(fs []Flag) func(args []string) []string {
	return func(args []string) []string {
		return Completion(args, fs)
	}
}

func Completion(args []string, fs []Flag) []string {
	completions := []string{}
	if len(args) == 0 {
		for _, flag := range fs {
			completions = append(completions, "--"+flag.Name)
		}
		return completions
	}
	last := args[len(args)-1]
	for _, flag := range fs {
		if strings.HasPrefix("--"+flag.Name, last) || strings.HasPrefix("-"+flag.Name, last) {
			completions = append(completions, "--"+flag.Name)
		}
	}
	return completions
}
