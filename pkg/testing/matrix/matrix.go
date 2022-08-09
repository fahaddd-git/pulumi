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

package matrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/engine"
	i "github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"github.com/stretchr/testify/assert"

	javagen "github.com/pulumi/pulumi-java/pkg/codegen/java"
	yamlgen "github.com/pulumi/pulumi-yaml/pkg/pulumiyaml/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	jsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pygen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
)

type LangTestOption struct {
	Language string
	Version  *semver.Version //Array of versions?
	Opts     *i.ProgramTestOptions
}

type PluginOptions struct {
	Name    string
	Kind    workspace.PluginKind
	Version semver.Version

	// Build is the set of commands to run to build the plugin and SDK
	// In the future we want to just generate the SDK using the plugin.
	Build []exec.Cmd

	// Bin is the path to the plugin to use.
	Bin string
}

type MatrixTestOptions struct {
	Languages  []LangTestOption
	Program    *i.ProgramTestOptions
	PulumiSDKs map[string]string
	Plugins    []PluginOptions
}

type projectGeneratorFunc func(directory string, project workspace.Project, p *pcl.Program, localProjects map[string]string) error

func Test(t *testing.T, opts MatrixTestOptions) {

	dir := opts.Program.Dir
	if !filepath.IsAbs(dir) {
		pwd, err := os.Getwd()
		assert.NoError(t, err)
		dir = filepath.Join(pwd, dir)
	}

	projfile := filepath.Join(dir, workspace.ProjectFile+".yaml")
	proj, err := workspace.LoadProject(projfile)
	assert.NoError(t, err)
	projinfo := &engine.Projinfo{Proj: proj, Root: dir}

	assert.NoError(t, err)

	//Assert runtime is YAML
	assert.Equal(t, projinfo.Proj.Runtime.Name(), "yaml")

	pwd, _, ctx, err := engine.ProjectInfoContext(projinfo, nil, cmdutil.Diag(), cmdutil.Diag(), false, nil)
	assert.NoError(t, err)

	var pclProgram *pcl.Program

	//check if *.pp file exists in dir
	files, err := ioutil.ReadDir(dir)
	assert.NoError(t, err)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".pp") {
			parser := syntax.NewParser()
			//get file content
			ppFilePath := filepath.Join(dir, file.Name())
			contents, err := ioutil.ReadFile(ppFilePath)
			assert.NoError(t, err)
			err = parser.ParseFile(bytes.NewReader(contents), ppFilePath)
			assert.NoError(t, err)
			program, diags, err := pcl.BindProgram(parser.Files, pcl.PluginHost(ctx.Host))
			assert.NoError(t, err)
			assert.Empty(t, diags)
			pclProgram = program

		}
	}
	if pclProgram == nil {
		loader := schema.NewPluginLoader(ctx.Host)
		_, pclProgram, err = yamlgen.Eject(pwd, loader)
		assert.NoError(t, err)
	}
	assert.NotNil(t, proj)
	assert.NotNil(t, pclProgram)

	localProjects := map[string]string{}

	//Here we are overriding the plugin links to use our own plugin links.
	proj.Plugins = &workspace.Plugins{}

	//This needs further testing
	for lang, path := range opts.PulumiSDKs {
		localProjects[lang] = path
	}

	//Execute build commands and add plugin links
	for _, plugin := range opts.Plugins {
		for _, cmd := range plugin.Build {
			err := cmd.Run()
			assert.NoError(t, err)
		}

		localProjects[plugin.Name] = fmt.Sprintf("%s/%s-sdk", dir, plugin.Name)

		assert.NotNil(t, plugin.Version)

		p := workspace.PluginOptions{
			Name:    plugin.Name,
			Path:    plugin.Bin,
			Version: plugin.Version.String(),
		}

		switch plugin.Kind {
		case workspace.AnalyzerPlugin:
			proj.Plugins.Analyzers = append(proj.Plugins.Analyzers, p)
		case workspace.LanguagePlugin:
			proj.Plugins.Languages = append(proj.Plugins.Languages, p)
		case workspace.ResourcePlugin:
			proj.Plugins.Providers = append(proj.Plugins.Providers, p)
		}
	}

	//Replace relative paths with absolute paths

	if proj.Plugins != nil {
		for i, provider := range proj.Plugins.Providers {
			if !filepath.IsAbs(provider.Path) {
				proj.Plugins.Providers[i].Path = filepath.Join(dir, provider.Path)
			}
		}
		for i, language := range proj.Plugins.Languages {
			if !filepath.IsAbs(language.Path) {
				proj.Plugins.Languages[i].Path = filepath.Join(dir, language.Path)
			}
		}

		for i, analyzer := range proj.Plugins.Analyzers {
			if !filepath.IsAbs(analyzer.Path) {
				proj.Plugins.Analyzers[i].Path = filepath.Join(dir, analyzer.Path)
			}
		}
	}

	//Generate SDKs
	pluginctx, err := plugin.NewContextWithRoot(cmdutil.Diag(), cmdutil.Diag(), nil, pwd, projinfo.Root,
		projinfo.Proj.Runtime.Options(), false, nil, proj.Plugins)
	assert.NoError(t, err)
	assert.NotNil(t, pluginctx)
	for _, plugin := range opts.Plugins {
		assert.NotNil(t, plugin.Version)

		if plugin.Kind != workspace.ResourcePlugin {
			continue
		}

		info, err := pluginctx.Host.ResolvePlugin(plugin.Kind, plugin.Name, &plugin.Version)
		assert.NoError(t, err)
		assert.NotNil(t, info)

		provider, err := pluginctx.Host.Provider(tokens.Package(info.Name), &plugin.Version)
		assert.NoError(t, err)
		assert.NotNil(t, provider)

		schemaBytes, err := provider.GetSchema(int(plugin.Version.Major))
		assert.NoError(t, err)
		assert.NotNil(t, schemaBytes)

		var spec schema.PackageSpec
		err = json.Unmarshal(schemaBytes, &spec)
		assert.NoError(t, err)

		pkg, diags, err := schema.BindSpec(spec, nil)
		assert.NoError(t, err)
		assert.Empty(t, diags)
		assert.NotNil(t, pkg)

		//Kludge
		/*for _, p := range opts.Plugins {
			if p.Name == pkg.Name {
				pkg.Version = &p.Version
			}
		}*/

		pkg.Test = true

		pkgName := pkg.Name

		for _, langOpt := range opts.Languages {
			lang := langOpt.Language

			var files map[string][]byte
			var err error
			switch lang {
			case "go":
				files, err = gogen.GeneratePackage(pkgName, pkg)
			case "python":
				PythonConfigurePkg(pkg)
				files, err = pygen.GeneratePackage(pkgName, pkg, files)
			case "nodejs":
				NodeConfigurePkg(pkg)
				files, err = jsgen.GeneratePackage(pkgName, pkg, files)
			case "dotnet":
				DotnetConfigurePkg(pkg)
				files, err = dotnetgen.GeneratePackage(pkgName, pkg, files)
			default:
				fmt.Printf("Unknown language: '%s'", lang)
				continue
			}
			assert.NoError(t, err)

			sdkDir := filepath.Join(dir, fmt.Sprintf("%s-sdk", pkgName), lang)
			for p, file := range files {
				// TODO: full conversion from path to filepath
				err = os.MkdirAll(filepath.Join(sdkDir, path.Dir(p)), 0700)
				assert.NoError(t, err)
				err = os.WriteFile(filepath.Join(sdkDir, p), file, 0600)
				assert.NoError(t, err)
			}
			if lang == "go" {
				//remove go.mod and go.sum files if they exist
				os.Remove(filepath.Join(sdkDir, "go.mod"))
				os.Remove(filepath.Join(sdkDir, "go.sum"))

				//run go mod init in the root of the sdk dir
				cmd := exec.Command("go", "mod", "init")
				cmd.Dir = sdkDir
				err = cmd.Run()
				//assert.NoError(t, err)

				//run go mod tidy in the root of the sdk dir
				cmd = exec.Command("go", "mod", "tidy", "-compat=1.18")
				cmd.Dir = sdkDir
				err = cmd.Run()
				//assert.NoError(t, err)
			}
		}
	}

	//Instantiate new directories and run pulumi convert on each language with new directories as output.
	for _, langOpt := range opts.Languages {
		var projectGenerator projectGeneratorFunc
		switch langOpt.Language {
		case "dotnet":
			projectGenerator = dotnetgen.GenerateProject
		case "go":
			projectGenerator = gogen.GenerateProject
		case "nodejs":
			projectGenerator = jsgen.GenerateProject
		case "python":
			projectGenerator = pygen.GenerateProject
		case "java":
			projectGenerator = func(directory string, project workspace.Project, p *pcl.Program, localProjects map[string]string) error {
				return javagen.GenerateProject(directory, project, p)
			}
		case "yaml": // nolint: goconst
			projectGenerator = func(directory string, project workspace.Project, p *pcl.Program, localProjects map[string]string) error {
				return yamlgen.GenerateProject(directory, project, p)
			}
		default:
			assert.FailNow(t, "Unsupported language: "+langOpt.Language)
		}

		subdir := filepath.Join(dir, langOpt.Language)
		//check if subdir exists
		if _, err := os.Stat(subdir); !os.IsNotExist(err) {
			assert.NoError(t, os.RemoveAll(subdir))
		}

		//create subdir
		err = os.MkdirAll(subdir, 0755)
		assert.NoError(t, err)

		//generate project
		err = projectGenerator(subdir, *proj, pclProgram, localProjects)
		assert.NoError(t, err)

		//Configure opts
		langPtOpts := i.ProgramTestOptions{}
		if opts.Program != nil {
			langPtOpts = *opts.Program
		}
		if langOpt.Opts != nil {
			langPtOpts = langPtOpts.With(*langOpt.Opts)
		}
		langPtOpts = langPtOpts.With(i.ProgramTestOptions{
			Dir: subdir,
		})
		t.Run(langOpt.Language, func(t *testing.T) {
			i.ProgramTest(t, &langPtOpts)
		})
		//clean up subdir

		//Including this seems to cause all tests to fail.
		/*t.Cleanup(func() {
			os.RemoveAll(subdir)
		})*/
	}
}

func PythonConfigurePkg(pkg *schema.Package) error {
	pkg.Version = &semver.Version{
		Major: pkg.Version.Major,
		Minor: pkg.Version.Minor,
		Patch: pkg.Version.Patch,
	} // Prune patch and build versions, since semver and python don't seem to agree on formatting.

	raw, notnil := pkg.Language["python"]
	message, ismessage := raw.(json.RawMessage)
	var pyPkg pygen.PackageInfo
	if notnil && ismessage {
		err := json.Unmarshal(message, &pyPkg)
		if err != nil {
			return err
		}
	} else if notnil && !ismessage {
		info, ok := raw.(pygen.PackageInfo)
		if !ok {
			return fmt.Errorf("invalid python package info")
		}
		pyPkg = info
	} else if !notnil {
		pyPkg = pygen.PackageInfo{}
	}
	pyPkg.RespectSchemaVersion = true
	pkg.Language["python"] = pyPkg
	return nil
}

func NodeConfigurePkg(pkg *schema.Package) error {
	raw, notnil := pkg.Language["nodejs"]
	message, ismessage := raw.(json.RawMessage)
	var nodePkg jsgen.NodePackageInfo
	if notnil && ismessage {
		err := json.Unmarshal(message, &nodePkg)
		if err != nil {
			return err
		}
	} else if notnil && !ismessage {
		info, ok := raw.(jsgen.NodePackageInfo)
		if !ok {
			return fmt.Errorf("invalid nodejs package info")
		}
		nodePkg = info
	} else if !notnil {
		nodePkg = jsgen.NodePackageInfo{}
	}
	nodePkg.RespectSchemaVersion = true
	pkg.Language["nodejs"] = nodePkg
	return nil
}

func DotnetConfigurePkg(pkg *schema.Package) error {
	raw, notnil := pkg.Language["csharp"]
	message, ismessage := raw.(json.RawMessage)
	var dotnetPkg dotnetgen.CSharpPackageInfo
	if notnil && ismessage {
		err := json.Unmarshal(message, &dotnetPkg)
		if err != nil {
			return err
		}
	} else if notnil && !ismessage {
		info, ok := raw.(dotnetgen.CSharpPackageInfo)
		if !ok {
			return fmt.Errorf("invalid dotnet package info")
		}
		dotnetPkg = info
	} else if !notnil {
		dotnetPkg = dotnetgen.CSharpPackageInfo{}
	}
	dotnetPkg.RespectSchemaVersion = true
	pkg.Language["csharp"] = dotnetPkg
	return nil
}