package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ejecs/ejecs/internal/generator"
	"github.com/ejecs/ejecs/internal/parser"
)

func main() {
	// Define flags
	inputFile := flag.String("input", "", "Input EJECS file")
	outputFile := flag.String("output", "", "Output file for generated Luau code")
	// library := flag.String("library", "jecs", "Target ECS library (ecr or jecs)") // Removed library flag
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		// fmt.Println("Usage: ejecs -input <input.jecs> -output <output.luau> -library <ecr|jecs>") // Old usage message
		fmt.Println("Usage: ejecs -input <input.jecs> -output <output.luau>")
		os.Exit(1)
	}

	// Validate library - Removed validation
	// if *library != "ecr" && *library != "jecs" {
	// 	fmt.Printf("Error: library must be either 'ecr' or 'jecs', got '%s'\n", *library)
	// 	os.Exit(1)
	// }

	// Read input file
	content, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Parse the content
	p := parser.New(string(content))
	ast, err := p.ParseProgram()
	if err != nil {
		// Check for parser errors
		if p.Errors() != nil && len(p.Errors()) > 0 {
			fmt.Println("Parse errors:")
			for _, msg := range p.Errors() {
				fmt.Println("-", msg)
			}
		} else {
			// Print general parse error if no specific messages
			fmt.Printf("Parse error: %v\n", err)
		}
		os.Exit(1)
	}

	// Generate code
	// g := generator.New(generator.Config{Library: *library}) // Old generator instantiation
	g := generator.New() // Simplified generator instantiation
	code, err := g.Generate(ast)
	if err != nil {
		fmt.Printf("Generation error: %v\n", err)
		os.Exit(1)
	}

	// Ensure output directory exists
	dir := filepath.Dir(*outputFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write output file
	if err := os.WriteFile(*outputFile, []byte(code), 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Successfully generated %s for %s library\n", *outputFile, *library) // Old success message
	fmt.Printf("Successfully generated %s\n", *outputFile)
}

// AST types
type Node interface {
	String() string
}

type Component struct {
	Name       string
	Fields     []Field
	Attributes []string
}

func (c *Component) String() string {
	var fields []string
	for _, f := range c.Fields {
		fields = append(fields, f.String())
	}
	attrs := ""
	if len(c.Attributes) > 0 {
		attrs = "@" + strings.Join(c.Attributes, " @") + "\n"
	}
	return fmt.Sprintf("%scomponent %s {\n    %s\n}", attrs, c.Name, strings.Join(fields, "\n    "))
}

type Field struct {
	Name     string
	Type     string
	Optional bool
}

func (f Field) String() string {
	opt := ""
	if f.Optional {
		opt = "?"
	}
	return fmt.Sprintf("%s: %s%s", f.Name, f.Type, opt)
}

type System struct {
	Name       string
	Components []string
	Frequency  string
	Priority   int
}

func (s *System) String() string {
	var parts []string
	if len(s.Components) > 0 {
		parts = append(parts, fmt.Sprintf("    using %s", strings.Join(s.Components, ", ")))
	}
	if s.Frequency != "" {
		parts = append(parts, fmt.Sprintf("    frequency: %s", s.Frequency))
	}
	if s.Priority != 0 {
		parts = append(parts, fmt.Sprintf("    priority: %d", s.Priority))
	}
	return fmt.Sprintf("system %s {\n%s\n}", s.Name, strings.Join(parts, "\n"))
}

type Relationship struct {
	Name string
	From string
	To   string
	Type string // one_to_one, one_to_many, many_to_many
}

func (r *Relationship) String() string {
	typeStr := ""
	if r.Type != "" {
		typeStr = fmt.Sprintf("@%s\n", r.Type)
	}
	return fmt.Sprintf("%srelationship %s {\n    %s: Entity\n    %s: Entity\n}", typeStr, r.Name, r.From, r.To)
}

// Parser functions
func parseJECS(content string) ([]Node, error) {
	var nodes []Node
	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}

		if strings.HasPrefix(line, "component") {
			comp, err := parseComponent(lines, &i)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, comp)
		} else if strings.HasPrefix(line, "system") {
			sys, err := parseSystem(lines, &i)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, sys)
		} else if strings.HasPrefix(line, "relationship") {
			rel, err := parseRelationship(lines, &i)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, rel)
		}
	}

	return nodes, nil
}

func parseComponent(lines []string, i *int) (*Component, error) {
	// TODO: Implement component parsing
	return nil, nil
}

func parseSystem(lines []string, i *int) (*System, error) {
	// TODO: Implement system parsing
	return nil, nil
}

func parseRelationship(lines []string, i *int) (*Relationship, error) {
	// TODO: Implement relationship parsing
	return nil, nil
}

// Code generation functions
func generateLuau(nodes []Node) (string, error) {
	var code strings.Builder

	// Write header
	code.WriteString("-- Generated by EJECS\n\n")
	code.WriteString("local Types = {}\n\n")

	// Generate code for each node
	for _, node := range nodes {
		switch n := node.(type) {
		case *Component:
			if err := generateComponent(&code, n); err != nil {
				return "", err
			}
		case *System:
			if err := generateSystem(&code, n); err != nil {
				return "", err
			}
		case *Relationship:
			if err := generateRelationship(&code, n); err != nil {
				return "", err
			}
		}
		code.WriteString("\n")
	}

	code.WriteString("return Types\n")
	return code.String(), nil
}

func generateComponent(code *strings.Builder, comp *Component) error {
	// TODO: Implement component code generation
	return nil
}

func generateSystem(code *strings.Builder, sys *System) error {
	// TODO: Implement system code generation
	return nil
}

func generateRelationship(code *strings.Builder, rel *Relationship) error {
	// TODO: Implement relationship code generation
	return nil
}
