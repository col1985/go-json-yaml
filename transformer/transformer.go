package transformer

import (
	"sigs.k8s.io/yaml"
)

func ToYaml(j []byte) ([]byte, error) {
	y, err := yaml.JSONToYAML(j)

	if err != nil {
		return nil, err
	}

	return y, nil

}

func ToJson(j []byte) ([]byte, error) {
	j, err := yaml.YAMLToJSON(j)
	if err != nil {
		return nil, err
	}

	return j, nil
}