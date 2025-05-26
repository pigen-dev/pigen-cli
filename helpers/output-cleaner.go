package helpers

import "strings"

// CleanPluginOutput recursively removes wrapping quotes from string values in map[string]any
func CleanPluginOutput(data map[string]any) map[string]any {
    for k, v := range data {
        switch val := v.(type) {
        case string:
            data[k] = trimWrappedQuotes(val)
        case map[string]any:
            data[k] = CleanPluginOutput(val)
        case []any:
            data[k] = cleanSlice(val)
        }
    }
    return data
}

func cleanSlice(slice []any) []any {
    for i, v := range slice {
        switch val := v.(type) {
        case string:
            slice[i] = trimWrappedQuotes(val)
        case map[string]any:
            slice[i] = CleanPluginOutput(val)
        case []any:
            slice[i] = cleanSlice(val)
        }
    }
    return slice
}

func trimWrappedQuotes(s string) string {
    s = strings.TrimSpace(s)
    if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
        return strings.Trim(s, "\"")
    }
    return s
}