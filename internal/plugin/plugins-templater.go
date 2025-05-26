package plugin

import (
	"fmt"
	"regexp"

	"github.com/pigen-dev/pigen-cli/helpers"
	"github.com/pigen-dev/pigen-cli/internal/templater"
	shared "github.com/pigen-dev/shared"
	"gopkg.in/yaml.v3"
)

// I'm adding this middleware to avoid cycle dependency problem
func PigenFileParser(file []byte) ([]byte, error) {
	deps := ExtractTemplateDependencies(string(file))
	outputs := make(map[string]any)
	for _, dep := range deps {
		outputs[dep] = GetOutput(dep).Output
	}
	res, err := templater.PigenReplacer(file, outputs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Extract plugin labels used in template references like {{ .Plugins.LABEL.xyz }}
func ExtractTemplateDependencies(yamlStr string) []string {
    re := regexp.MustCompile(`{{\s*\.Plugins\.([^.}]+)\.[^}]+}}`)
    matches := re.FindAllStringSubmatch(yamlStr, -1)
		// Use a map to eliminate duplicate plugin names and ensure each dependency is listed only once.
		// To avoid hitting the pigen core api twice
    seen := make(map[string]struct{})
    var deps []string
    for _, match := range matches {
        if len(match) > 1 {
            plugin := match[1]
            if _, exists := seen[plugin]; !exists {
								// struct{}{} takes 0 memory
                seen[plugin] = struct{}{}
                deps = append(deps, plugin)
            }
        }
    }

    return deps
}

func PluginParser(plugin *shared.PluginStruct) error {
	bytes, err := yaml.Marshal(plugin)
	if err != nil {
		return fmt.Errorf("failed to marshal plugin data into yaml: %w", err)
	}
	parsedPlugin, err := PigenFileParser(bytes)
	if err != nil {
		return fmt.Errorf("failed to run plugin output parser: %w", err)
	}
	err = helpers.YamlToStruct(parsedPlugin, &plugin)
	if err != nil {
		return err
	}
	return nil
}