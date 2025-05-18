package helpers

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/joho/godotenv"
)


type PigenENV struct {
    ENV map[string]string
}


func ReplaceSecrets(in []byte) ([]byte, error) {
	// Load the .env.pigen file
	secretsMap, err := loadEnvFile(".env.pigen")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	pigenENV := PigenENV{
		ENV: secretsMap,
	}
	fmt.Println("Secrets map:", pigenENV)
	tmpl, err := template.New("SecretParser").Parse(string(in))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pigenENV)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	fmt.Printf("Replaced secrets: %s\n", buf.String())
	return buf.Bytes(), nil
}

func loadEnvFile(path string) (map[string]string, error) {
    return godotenv.Read(path)
}