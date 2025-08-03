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
	main, _, err := loadOpcodes("gen/Opcodes.json")
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	main = omit(main, "STOP", "RRCA")

	if err := tmpl.Execute(f, main); err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	return nil
}

var tmpl = template.Must(template.New("main").Parse(`package gameboy
type Instruction func(cpu *CPU)

{{ range . }}
// {{.String}}
func {{.ID}}(cpu *CPU) {
	{{- if eq "ADD" .Mnemonic -}} 
		{{template "add" .DataAdd -}}
	{{- else if eq "INC" .Mnemonic -}}
		{{ template "inc" .DataInc -}}
	{{- else if eq "DEC" .Mnemonic -}}
		{{ template "dec" .DataDec -}}
	{{- else if eq "LD" .Mnemonic -}}
		{{ template "ld" .DataLd -}}
	{{- else if eq "CALL" .Mnemonic -}}
		{{ template "call" .DataCall -}}
	{{- else if eq "PUSH" .Mnemonic -}}
		{{ template "push" .DataPush -}}
	{{- else if eq "POP" .Mnemonic -}}
		{{ template "pop" .DataPop -}}
	{{else}}
		// TODO: {{.ID}}
	{{end -}}
}
{{end}}

var ops = map[uint8]Instruction{
	{{ range . -}}
	{{printf "%#x" .Code}}: {{.ID}},
	{{end}}
}
`))

// formats file. The generated file may look crap, so we pass it to gofmt to make
// it a bit prettier.
func formatFile(file string) error {
	cmd := exec.Command("gofmt", "-w", file)
	return cmd.Run()
}
