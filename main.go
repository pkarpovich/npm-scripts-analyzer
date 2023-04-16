package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/umputun/go-flags"
)

var opts struct {
	PackageJsonPath string `short:"p" long:"package-json-path" description:"Path to package.json file" default:"./package.json"`
}

func (g *Graph) FindDuplicates() map[string][]DuplicateInfo {
	visited := make(map[string]map[string]bool)
	duplicates := make(map[string][]DuplicateInfo)

	var dfs func(node *Node, origin string, parents []string)
	dfs = func(node *Node, origin string, parents []string) {
		if visited[origin] == nil {
			visited[origin] = make(map[string]bool)
		}

		if visited[origin][node.Name] {
			found := false
			for idx, duplicate := range duplicates[origin] {
				if duplicate.ScriptName == node.Name {
					duplicates[origin][idx].Count++
					duplicates[origin][idx].Parents = append(duplicates[origin][idx].Parents, parents...)
					found = true
					break
				}
			}
			if !found {
				duplicates[origin] = append(duplicates[origin], DuplicateInfo{
					ScriptName: node.Name,
					Count:      1,
					Parents:    parents,
				})
			}
			return
		}

		visited[origin][node.Name] = true
		parents = append(parents, node.Name)
		for _, child := range node.Children {
			dfs(child, origin, parents)
		}
		parents = parents[:len(parents)-1]
	}

	for _, node := range g.Nodes {
		dfs(node, node.Name, []string{})
	}

	return duplicates
}

func PrepareScriptsGraph(packageJsonPath string) (error, *Graph) {
	data, err := os.ReadFile(packageJsonPath)
	if err != nil {
		return fmt.Errorf("error reading package.json: %v", err), nil
	}

	var packageJSON map[string]interface{}
	if err := json.Unmarshal(data, &packageJSON); err != nil {
		return fmt.Errorf("error parsing package.json: %v", err), nil
	}

	scripts, ok := packageJSON["scripts"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("no scripts found in package.json"), nil
	}

	graph := NewGraph()
	npmRemoveRegex := regexp.MustCompile(`npm(?: run)?\s`)
	for scriptName, command := range scripts {
		commandStr, ok := command.(string)
		if !ok || commandStr == "" {
			fmt.Printf("Invalid script command for %s: %v\n", scriptName, command)
			continue
		}

		commandStr = npmRemoveRegex.ReplaceAllString(commandStr, "")
		commands := strings.Split(commandStr, "&&")

		for _, command := range commands {
			trimCommand := strings.TrimSpace(command)

			graph.AddEdge(scriptName, trimCommand)
		}
	}

	return nil, graph
}

func main() {
	p := flags.NewParser(&opts, flags.Default)
	p.SubcommandsOptional = true
	if _, err := p.Parse(); err != nil {
		if err.(*flags.Error).Type != flags.ErrHelp {
			log.Printf("[ERROR] cli error: %v", err)
		}
		os.Exit(2)
	}

	err, graph := PrepareScriptsGraph(opts.PackageJsonPath)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		os.Exit(1)
	}

	fmt.Println("NPM Scripts Graph:")
	for _, node := range graph.Nodes {
		if len(node.Children) > 0 {
			node.Print(0)
		}
	}

	duplicates := graph.FindDuplicates()
	if len(duplicates) > 0 {
		fmt.Println("\nDuplicates found:")
		for script, duplicateScripts := range duplicates {
			fmt.Printf("%s:\n", script)
			for _, duplicate := range duplicateScripts {
				fmt.Printf("  - %s: %d times (Parents: %v)\n\n", duplicate.ScriptName, duplicate.Count, duplicate.Parents)
			}
		}
	} else {
		fmt.Println("\nNo duplicates found.")
	}
}
