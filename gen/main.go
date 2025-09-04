package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"text/template"
)

func main() {
	const DEST = "instructions.go"

	if err := generateInstructions(DEST); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if err := formatFile(DEST); err != nil {
		fmt.Fprintf(os.Stderr, "failed to format %s: %v\n", DEST, err)
		os.Exit(1)
		return
	}

	fmt.Printf("Code generated\n")
}

func omit(input []Opcode, skip ...string) []Opcode {
	var res []Opcode
	for _, inp := range input {
		if !slices.Contains(skip, inp.Mnemonic) {
			res = append(res, inp)
		}
	}
	return res

}

// Reads from Opcodes.json and generates code that gets written to file.
func generateInstructions(file string) error {
	main, ext, err := loadOpcodes("gen/Opcodes.json")
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	data := map[string][]Opcode{
		"Main": main,
		"Ext":  ext,
		"Both": append(main, ext...),
	}

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	return nil
}

var tmpl = template.Must(template.New("main").Parse(`package gameboy
import "fmt"
type Instruction interface {
	Exec(cpu *CPU)
	Code() uint8
	String() string
}

{{ range .Both }}
// {{.Desc}}
type {{.ID}} struct {}
func ({{.ID}}) Exec(cpu *CPU) {
	{{- if eq "ADD" .Mnemonic -}} 
		{{template "add" .DataAdd -}}
	{{- else if eq "SUB" .Mnemonic -}}
		{{ template "sub" .DataSub -}}
	{{- else if eq "INC" .Mnemonic -}}
		{{ template "inc" .DataInc -}}
	{{- else if eq "DEC" .Mnemonic -}}
		{{ template "dec" .DataDec -}}
	{{- else if eq "LD" .Mnemonic -}}
		{{ template "ld" .DataLd -}}
	{{- else if eq "LDH" .Mnemonic -}}
		{{ template "ldh" .DataLdh -}}
	{{- else if eq "CALL" .Mnemonic -}}
		{{ template "call" .DataCall -}}
	{{- else if eq "PUSH" .Mnemonic -}}
		{{ template "push" .DataPush -}}
	{{- else if eq "POP" .Mnemonic -}}
		{{ template "pop" .DataPop -}}
	{{- else if eq "RET" .Mnemonic -}}
		{{ template "ret" .DataRet -}}
	{{- else if eq "STOP" .Mnemonic -}}
		{{ template "stop" .DataStop -}}
	{{- else if eq "RST" .Mnemonic -}}
		{{ template "rst" .DataRst -}}
	{{- else if eq "OR" .Mnemonic -}}
		{{ template "or" .DataOr -}}
	{{- else if eq "XOR" .Mnemonic -}}
		{{ template "xor" .DataXor -}}
	{{- else if eq "SWAP" .Mnemonic -}}
		{{ template "swap" .DataSwap -}}
	{{- else if eq "CP" .Mnemonic -}}
		{{ template "cp" .DataCp -}}
	{{- else if eq "ILLEGAL" .Mnemonic -}}
		{{ template "illegal" .DataIllegal -}}
	{{- else if eq "PREFIX" .Mnemonic -}}
		{{ template "prefix" .DataPrefix -}}
	{{- else if eq "JR" .Mnemonic -}}
		{{ template "jr" .DataJr -}}
	{{- else if eq "EI" .Mnemonic -}}
		{{ template "ei" .DataEi -}}
	{{- else if eq "DI" .Mnemonic -}}
		{{ template "di" .DataDi -}}
	{{- else if eq "RETI" .Mnemonic -}}
		{{ template "reti" .DataReti -}}
	{{- else if eq "JP" .Mnemonic -}}
		{{ template "jp" .DataJp -}}
	{{- else if eq "AND" .Mnemonic -}}
		{{ template "and" .DataAnd -}}
	{{- else if eq "RES" .Mnemonic -}}
		{{ template "res" .DataRes -}}
	{{- else if eq "SET" .Mnemonic -}}
		{{ template "set" .DataSet -}}
	{{- else if eq "CPL" .Mnemonic -}}
		{{ template "cpl" .DataCpl -}}
	{{- else if eq "SRL" .Mnemonic -}}
		{{ template "srl" .DataSrl -}}
	{{- else if eq "RLA" .Mnemonic -}}
		{{ template "rla" .DataRla -}}
	{{- else if eq "RRA" .Mnemonic -}}
		{{ template "rra" .DataRra -}}
	{{- else if eq "RRCA" .Mnemonic -}}
		{{ template "rrca" .DataRrca -}}
	{{- else if eq "RLCA" .Mnemonic -}}
		{{ template "rlca" .DataRlca -}}
	{{- else if eq "NOP" .Mnemonic -}}
		{{ template "nop" .DataNop -}}
	{{/* CB-prefixed stuff */}}
	{{- else if eq "BIT" .Mnemonic -}}
		{{ template "bit" .DataBit -}}
	{{- else if eq "RL" .Mnemonic -}}
		{{ template "rl" .DataRl -}}
	{{- else if eq "RLC" .Mnemonic -}}
		{{ template "rlc" .DataRl -}}
	{{- else if eq "RR" .Mnemonic -}}
		{{ template "rr" .DataRr -}}
	{{- else if eq "RRC" .Mnemonic -}}
		{{ template "rrc" .DataRrc -}}
	{{else}}
		fmt.Println("TODO: {{.ID}}")
		// panic("TODO {{.ID}}")
	{{end -}}
}
func ({{.ID}}) Code() uint8 {
	return {{printf "%#X" .Code}}
}
func ({{.ID}}) String() string {
	return "{{.String}}"
}
{{end}}

var ops = map[uint8]Instruction{
	{{ range .Main -}}
	{{printf "%#x" .Code}}: {{.ID}}{},
	{{end}}
}
var extOps = map[uint8]Instruction{
	{{ range .Ext -}}
	{{printf "%#x" .Code}}: {{.ID}}{},
	{{end}}
}

// returns code given a string. Useful during testing
func code(s string) uint8 {
	switch s {
	{{ range .Both -}}
	case "{{.String}}": return {{printf "%#X" .Code}}
	{{end}}
	default: panic(fmt.Sprintf("Unknown code for %q", s))
	}
}
`))

// formats file. The generated file may look crap, so we pass it to gofmt to make
// it a bit prettier.
func formatFile(file string) error {
	cmd := exec.Command("gofmt", "-w", file)
	return cmd.Run()
}
