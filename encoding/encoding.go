package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	JSONData, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	var data models.DockerCompose
	err = json.Unmarshal(JSONData, &data)
	if err != nil {
		return err
	}

	YAMLData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(j.FileOutput, YAMLData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	YAMLData, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	var data models.DockerCompose
	err = yaml.Unmarshal(YAMLData, &data)
	if err != nil {
		return err
	}

	JSONData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(y.FileOutput, JSONData, 0644)
	if err != nil {
		return err
	}

	return nil
}
