package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

var dest = "instructions.go"

type Operand struct {
	Name      string `json:"name"`
	Immediate bool   `json:"immediate"`
	Bytes     int    `json:"bytes"`

	Increment bool `json:"increment"` // if true, then this operand is incremented after use
	Decrement bool `json:"decrement"` // if true, then this operand is decremented after use
}

func (o Operand) String() string {
	var s strings.Builder
	if !o.Immediate {
		s.WriteString("(")
	}
	s.WriteString(o.Name)
	if o.Increment {
		s.WriteString("+")
	}
	if o.Decrement {
		s.WriteString("-")
	}
	if !o.Immediate {
		s.WriteString(")")
	}
	return s.String()
}

func (o Operand) Is16Bit() bool { return len(o.Name) == 2 }
func (o Operand) Is8Bit() bool  { return len(o.Name) == 1 }

type Operands []Operand

func (ops Operands) First() Operand  { return ops[0] }
func (ops Operands) Second() Operand { return ops[1] }
func (ops Operands) Last() Operand   { return ops[len(ops)-1] }

type Op struct {
	Mnemonic  string            `json:"mnemonic"`
	Bytes     int               `json:"bytes"`
	Cycles    []int             `json:"cycles"`
	Operands  Operands          `json:"operands"`
	Immediate bool              `json:"immediate"`
	Flags     map[string]string `json:"flags"`
	ID        string
	Code      int
}

func (o Op) String() string {
	var parts []string
	for _, arg := range o.Operands {
		parts = append(parts, arg.String())
	}
	return fmt.Sprintf("%s %s    code=%#02x", o.Mnemonic, strings.Join(parts, ","), o.Code)
}

func (o Op) FirstArg() Operand {
	if len(o.Operands) == 0 {
		return Operand{}
	}
	return o.Operands[0]
}

func (o Op) Is16Bit() bool { return !o.FirstArg().Is16Bit() }
func (o Op) Is8Bit() bool  { return o.FirstArg().Is8Bit() }
func (o Op) Foo() string   { return fmt.Sprintf("Foo %s", o.Mnemonic) }
func (o Op) CycleCount() int {
	var count int
	for _, c := range o.Cycles {
		count += c
	}
	return count
}

func run() error {
	b, err := os.ReadFile("gen/Opcodes.json")
	if err != nil {
		return fmt.Errorf("failed to read Opcodes.json: %w", err)
	}

	var data struct {
		Unprefixed map[string]Op `json:"unprefixed"`
		Prefixed   map[string]Op `json:"prefixed"`
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return fmt.Errorf("failed to unmarshal Opcodes.json: %w", err)
	}

	for code, v := range data.Unprefixed {
		keep := []string{
			"ADD",
			"INC",
			"DEC",
			"LD",
		}
		if !slices.Contains(keep, v.Mnemonic) {
			delete(data.Unprefixed, code)
			continue
		}
		v.ID = fmt.Sprintf("%s_%s", v.Mnemonic, code[2:])
		i64, _ := strconv.ParseInt(code, 0, 64)
		v.Code = int(i64)

		data.Unprefixed[code] = v
		if len(v.Cycles) > 1 {
			fmt.Printf("op %s has %d cycles\n", code, len(v.Cycles))
		}
	}

	for k, v := range data.Prefixed {
		if len(v.Cycles) > 1 {
			fmt.Printf("prefixed op %s has %d cycles\n", k, len(v.Cycles))
		}
	}
	data.Prefixed = map[string]Op{}

	templ := template.Must(tmpl.Funcs(template.FuncMap{
		"operands": func(x Op) string { // human-readable operand, // eg `HL,BC`
			var parts []string
			for _, arg := range x.Operands {
				parts = append(parts, arg.Name)
			}
			return strings.Join(parts, ",")
		},
	}).Parse(codeTemplate))
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	templ.Execute(f, data)

	return nil
}

const codeTemplate = `package gameboy

type Instruction func(cpu *CPU)

{{ range $key, $op := .Unprefixed }}
// {{$op.String}}
func {{$op.ID}}(cpu *CPU) {
	{{ if eq "ADD" $op.Mnemonic }} 
		{{ template "add" $op.DataAdd }}
	{{ else if eq "INC" $op.Mnemonic }}
		{{ template "inc" $op.DataInc }}
	{{ else if eq "DEC" $op.Mnemonic }}
		{{ template "dec" $op.DataDec }}
	{{ else if eq "LD" $op.Mnemonic }}
		{{ template "ld" $op.DataLd }}
	{{else}}
		// TODO: {{$op.ID}}
	{{end }}
	cpu.cycles += {{$op.CycleCount}}
}
{{ end}}

var ops = map[uint8]Instruction{
	{{ range .Unprefixed -}}
	{{printf "%#x" .Code}}: {{.ID}},
	{{end}}
}
`

var tmpl = template.New("main")

func formatFile() error {
	cmd := exec.Command("gofmt", "-w", dest)
	return cmd.Run()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if err := formatFile(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to format %s: %v\n", dest, err)
		os.Exit(1)
		return
	}

	fmt.Printf("Code generated\n")
}
