// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package report

import (
	"sync"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/version"
)

type Finish func(diags hcl.Diagnostics, err error)
type Guard func()

type Reporter interface {
	// Report a call to GenerateProgram.
	Start(title, language string, files []*syntax.File) (Finish, Guard)
	Summary() Summary
}

type Summary struct {
	Stats
	Name          string `json:"name"`
	Version       string `json:"version"`
	ReportVersion string `json:"reportVersion"`
	Languages     map[string]*Language
}

type Stats struct {
	NumConversions int
	Successes      int
}

type Language struct {
	Stats

	// A mapping from Error:(title:occurrences)
	Warnings map[string]map[string]int `json:"warning,omitempty"`
	Errors   map[string]map[string]int `json:"errors,omitempty"`

	// A mapping from title:files
	Files map[string][]string `json:"files,omitempty"`
}

func New(name, version string) Reporter {
	return &reporter{
		summary: Summary{
			Name:    name,
			Version: version}}
}

type reporter struct {
	summary Summary
	m       sync.Mutex
}

func (s *Stats) update(succeed bool) {
	s.NumConversions += 1
	if succeed {
		s.Successes += 1
	}
}

func (s *Summary) getLanguage(lang string) *Language {
	if s.Languages == nil {
		s.Languages = map[string]*Language{}
	}
	l, ok := s.Languages[lang]
	if !ok {
		l = new(Language)
		s.Languages[lang] = l
	}
	return l

}

func (r *reporter) Start(title, language string, files []*syntax.File) (Finish, Guard) {
	m := new(sync.Mutex)
	var reported bool

	finish := func(diags hcl.Diagnostics, err error) {
		m.Lock()
		defer m.Unlock()
		r.m.Lock()
		defer r.m.Unlock()
		failed := diags.HasErrors() || err != nil
		r.summary.Stats.update(!failed)
		lang := r.summary.getLanguage(language)
		lang.Stats.update(!failed)
		// TODO

		reported = true
	}

	guard := func() {
		m.Lock()
		defer m.Unlock()
		// The problem has already been reported. The panic happened later. Do nothing
		if reported {
			return
		}
		err := recover()
		if err == nil {
			return
		}
	}
	return finish, guard
}

func (r *reporter) Summary() Summary {
	if r == nil {
		return Summary{ReportVersion: version.Version}
	}
	r.summary.ReportVersion = version.Version
	return r.summary
}
