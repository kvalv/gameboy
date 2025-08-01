package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
)

var dest = "instructions.go"

type Operand struct {
	Name      string `json:"name"`
	Immediate bool   `json:"immediate"`
	Bytes     int    `json:"bytes"`
}

func (o Operand) Is16Bit() bool { return len(o.Name) == 2 }
func (o Operand) Is8Bit() bool  { return len(o.Name) == 1 }

type Operands []Operand

func (ops Operands) First() Operand { return ops[0] }
func (ops Operands) Last() Operand  { return ops[len(ops)-1] }

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

func (o Op) FirstArg() Operand {
	if len(o.Operands) == 0 {
		return Operand{}
	}
	return o.Operands[0]
}

type InstructionType int

const (
	ITypeMisc InstructionType = iota
	ITypeJump
	ITypeLoad8
	ITypeLoad16
	IArithmetic8
	IArithmetic16
	IRotation8
)

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

func (o Op) DataAdd() templDataAdd {
	return templDataAdd{
		Name:      o.Operands.Last().Name,
		Instr16:   o.Operands.Last().Is16Bit(),
		Immediate: o.Operands.Last().Immediate,
		DestA:     o.Operands.First().Name == "A",
	}
}
func (o Op) DataInc() templDataInc {
	return templDataInc{
		Name:      o.Operands.First().Name,
		Immediate: o.Operands.First().Immediate,
		Instr16:   o.Operands.First().Is16Bit(),
	}
}

func (o Op) InstructionType() InstructionType {
	if o.Code > 0xff {
		// ...
		return ITypeMisc
	}

	switch o.Mnemonic {
	case "NOP", "STOP", "HALT", "PREFIX", "EI", "DI":
		return ITypeMisc
	case "ADD", "INC", "DEC":
		if o.FirstArg().Is8Bit() {
			return IArithmetic8
		}
		return IArithmetic16
	}

	return ITypeMisc // TODO
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
		if v.Mnemonic != "ADD" && v.Mnemonic != "INC" {
			delete(data.Unprefixed, code)
			continue
		}
		v.ID = fmt.Sprintf("%s_%s", v.Mnemonic, code[2:])
		i64, _ := strconv.ParseInt(code, 0, 64)
		v.Code = int(i64)
		// v.Code, _ = strconv.Atoi(code[2:]) // "0xab" -> 0xab
		fmt.Printf("parsed %s -> %d\n", code, v.Code)

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
		"join": func(x any) string {
			switch x := x.(type) {
			case []int:
				var parts []string
				for _, elem := range x {
					parts = append(parts, fmt.Sprintf("%d", elem))
				}
				return strings.Join(parts, ", ")
			case []string:
				var parts []string
				for _, elem := range x {
					parts = append(parts, fmt.Sprintf("%s", elem))
				}
				return strings.Join(parts, ", ")
			default:
				panic(fmt.Sprintf("not implemented for %T", x))
			}
		},
		"capitalize": strings.ToUpper,
		"operands": func(x Op) string { // human-readable operand, // eg `HL,BC`
			var parts []string
			for _, arg := range x.Operands {
				parts = append(parts, arg.Name)
			}
			return strings.Join(parts, ",")
		},
		"inject": func(x Op) string {
			return `{{template "add16" .}}`
		},
		"args": func(x Op) string { // generates args string, eg `cpu.A, cpu.B`
			var parts []string
			for _, arg := range x.Operands {
				var part string
				switch arg.Name {
				case "HL":
					part = "read16(cpu.H, cpu.L)"
				case "BC":
					part = "read16(cpu.B, cpu.C)"
				case "DE":
					part = "read16(cpu.D, cpu.E)"
				case "SP":
					part = "cpu.SP"
				case "AF":
					part = "read16(cpu.A, cpu.F)"
				case "n8":
					part = "TODO"
				default:
					if len(arg.Name) != 1 {
						panic(fmt.Sprintf("Unknown 16-bit operand %s", arg.Name))
					}
					part = fmt.Sprintf("read16(0, cpu.%s)", arg.Name)
					// part = fmt.Sprintf("cpu.%s", arg.Name)
				}
				parts = append(parts, part)
			}
			return strings.Join(parts, ",")
		},
		"args8": func(x Op) string { // generates args string, eg `cpu.A, cpu.B`
			var parts []string
			for _, arg := range x.Operands {
				parts = append(parts, fmt.Sprintf("cpu.%s", arg.Name))
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
// {{$op.Mnemonic}} {{$key}} {{operands $op}}
func {{$op.ID}}(cpu *CPU) {
	{{ if eq "ADD" $op.Mnemonic }} 
		{{ template "add" $op.DataAdd }}
	{{ else if eq "INC" $op.Mnemonic }}
		{{ template "inc" $op.DataInc }}
	{{else}}
		{{template "todo" $op  }}
	{{end }}
	cpu.cycles += {{$op.CycleCount}}
}
{{ end}}

var ops = map[uint8]Instruction{
	{{ range .Unprefixed -}}
	{{printf "%#x" .Code}}: {{.ID}},
	{{end}}
}

{{define "todo"}}
panic("Not implemented")
{{end}}

`

type Value struct {
	Op    *Operand
	Const *uint8
}

type templDataAdd struct {
	Name      string // name of register for what to add
	Instr16   bool   // if true, then write to HL, otherwise write to A
	Immediate bool   // if not true, we require a load
	DestA     bool
}

var tmpl = template.New("main")

func getRegister(name string) (reg string) {
	defer func() {
		reg = "cpu." + reg // add accessor
	}()
	switch name {
	case "A", "F", "B", "C", "D", "E", "H", "L":
		// direct access
		reg = name
	case "PC", "SP":
		// direct access
		reg = name
	case "HL", "BC", "DE":
		// via method
		reg = name + "()"
	}
	return
}

// inc bc, de, hl, sp, b, d, h, (hl), c, e, l, a
// writes a line of code that sets the register to variable name
// varname is the variable name holding the value
func setRegister(name string, varname string) string {
	switch name {
	case "A", "C", "E", "L", "B", "D", "H", "SP":
		return fmt.Sprintf("cpu.%s = %s", name, varname)
	case "BC", "DE", "HL":
		r1 := string(name[0])
		r2 := string(name[1])
		return fmt.Sprintf("cpu.%s, cpu.%s = splitU16(%s)", r1, r2, varname)
	default:
		return fmt.Sprintf("// todo: setreg... %s %s", name, varname)
	}
}

// e8, n8, a8,
// e8: XOR 0xee
// n8: LD 0xf8
// a8: LDH
// n8: 0xc6

var templAdd = template.Must(tmpl.New("add").
	Funcs(template.FuncMap{
		"reg": getRegister,
	}).
	Parse(`
{{ if eq .Name "e8"}} {{/* special case: add to stack pointer */}}
	var n int16
	cpu.load(cpu.PC, &n)
	res, flags := cpu.AddSigned16(int16(cpu.A), n)
	cpu.PC = uint16(res)
	cpu.F = flags
{{ else if eq .Name "n8"}} {{/* Add immediate data */}}
	var n uint8
	cpu.load(cpu.PC, &n)
	cpu.A, cpu.F = cpu.Add(cpu.A, n)
{{ else if not .Immediate }}
	// not immediate brah
	var n uint8
	cpu.load({{reg .Name}}, &n)
	s8 := n // TODO: signed8(n)
	cpu.A, cpu.F = cpu.Add(cpu.A, s8)
{{ else if .Instr16 }}
	res, flags := cpu.Add16(cpu.HL(), {{ reg .Name }})
	{{ if .DestA}}
	_, cpu.A = splitU16(res)
	{{else}}
	cpu.H, cpu.L = splitU16(res)
	{{end }}
	cpu.F = FlagRegister(flags)
{{ else }}
	cpu.A, cpu.F = cpu.Add(cpu.A, {{reg .Name }})
{{end }}
`))

type templDataInc struct {
	Name      string // name of register for what to add
	Instr16   bool
	Immediate bool // if not true, we require a load
}

var templInc = template.Must(tmpl.New("inc").
	Funcs(template.FuncMap{
		"reg":    getRegister,
		"setreg": setRegister,
	}).
	Parse(`
{{if .Immediate}}
	{{if .Instr16}}
	res, flags := cpu.Add16({{reg .Name}}, 0x01)
	{{else}}
	res, flags := cpu.Add({{reg .Name}}, 0x01)
	{{end}}
	{{setreg .Name "res"}}
	cpu.F = FlagRegister(flags)
{{else}}
	// Increments data at the absolute address specified by the register
	var val uint8
	cpu.load({{reg .Name}}, &val)
	next, flags := cpu.Add(val, 0x01)
	cpu.write({{reg .Name}}, next)
	cpu.F = FlagRegister(flags)
{{end}}
`))

func formatFile() {
	cmd := exec.Command("gofmt", "-w", dest)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to format %s: %v\n", dest, err)
		return
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	formatFile()

	fmt.Printf("I am generated\n")
}
