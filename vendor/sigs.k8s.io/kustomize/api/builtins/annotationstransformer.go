// Code generated by pluginator on AnnotationsTransformer; DO NOT EDIT.
// pluginator {unknown  1970-01-01T00:00:00Z  }

package builtins

import (
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/transform"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

// Add the given annotations to the given field specifications.
type AnnotationsTransformerPlugin struct {
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	FieldSpecs  []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

func (p *AnnotationsTransformerPlugin) Config(
	h *resmap.PluginHelpers, c []byte) (err error) {
	p.Annotations = nil
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *AnnotationsTransformerPlugin) Transform(m resmap.ResMap) error {
	t, err := transform.NewMapTransformer(
		p.FieldSpecs,
		p.Annotations,
	)
	if err != nil {
		return err
	}
	return t.Transform(m)
}

func NewAnnotationsTransformerPlugin() resmap.TransformerPlugin {
	return &AnnotationsTransformerPlugin{}
}
