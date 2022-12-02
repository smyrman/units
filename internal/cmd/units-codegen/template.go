package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"text/template"
)

type pkgConfig struct {
	sync.RWMutex
	compiled bool
	PkgName  string
	Types    map[string]typeDef
}

// Compile realizes generator fields so that the struct can be used for template rendering.
func (p *pkgConfig) Compile() error {
	if p.compiled {
		return nil
	}
	wgs := make(map[string]*sync.WaitGroup)
	for k := range p.Types {
		wgs[k] = &sync.WaitGroup{}
		wgs[k].Add(1)
	}

	p.RLock()
	for k, t := range p.Types {
		go t.compileRoutine(k, p, wgs)
	}
	p.RUnlock()

	for k, wg := range wgs {
		wg.Wait()
		p.RLock()
		err := p.Types[k].compileError
		p.RUnlock()
		if err != nil {
			return err
		}
	}

	p.compiled = true
	return nil
}

type typeDef struct {
	compileError    error
	PkgName         string
	Import          []string
	Self            string
	Name            string `json:"-"`
	DivDurationType string
	MulDurationType string
	Units           []unitDef
	UnitGenerators  []unitGenerator
}

// compileRoutine should be started concurrently with go, and all typeDefs should be passed the same WaitGroup lookup.
func (t typeDef) compileRoutine(name string, p *pkgConfig, wgs map[string]*sync.WaitGroup) {
	defer wgs[name].Done()
	defer func() {
		p.Lock()
		p.Types[name] = t
		p.Unlock()
	}()

	var hasTime bool

	t.Name = name
	t.PkgName = p.PkgName

	// Append import time if needed.
	if len(t.DivDurationType) > 0 || len(t.MulDurationType) > 0 {
		hasTime = false
		for _, s := range t.Import {
			if s == "time" {
				hasTime = true
				break
			}
		}
		if !hasTime {
			t.Import = append(t.Import, "time")
		}
	}

	// Generate units -- wait for dependent types to complete.
	for i, g := range t.UnitGenerators {
		wg, ok := wgs[g.From]
		if !ok {
			t.compileError = fmt.Errorf("types.%s.unitGenerators[%d].from: '%s' not found in type lookup",
				name, i, g.From,
			)
			return
		}
		wg.Wait()

		p.RLock()
		units, err := g.units(name, p.Types)
		p.RUnlock()
		if err != nil {
			t.compileError = fmt.Errorf("types.%s.unitGenerators[%d]%s", name, i, err.Error())
			return
		}
		t.Units = append(t.Units, units...)
	}
	if len(t.UnitGenerators) > 0 {
		t.UnitGenerators = nil
	}
}

type unitDef struct {
	Name       string
	NamePlural string
	Value      string
}

type unitGenerator struct {
	From        string
	Pattern     string
	NameSuffix  string
	ValueSuffix string
}

func (g unitGenerator) units(name string, types map[string]typeDef) ([]unitDef, error) {
	var r *regexp.Regexp
	var err error

	t, ok := types[g.From]
	if !ok {
		return nil, fmt.Errorf(".from: '%s' not found in type lookup", g.From)
	}
	if g.Pattern != "" {
		r, err = regexp.Compile(g.Pattern)
		if err != nil {
			return nil, fmt.Errorf(".pattern: %s", err.Error())
		}
	}
	units := make([]unitDef, 0, len(t.Units))
	for _, from := range t.Units {
		if r != nil && !r.MatchString(from.Name) {
			continue
		}
		units = append(units, unitDef{
			Name:       from.Name + g.NameSuffix,
			NamePlural: from.NamePlural + g.NameSuffix,
			Value:      fmt.Sprintf("%s(%s)%s", name, from.Name, g.ValueSuffix),
		})
	}
	return units, nil
}

var tmplFunctions = template.FuncMap{
	"lower": strings.ToLower,
}

//go:embed template.go.tmpl
var tmpl string
