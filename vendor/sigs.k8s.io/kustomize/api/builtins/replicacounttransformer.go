// Code generated by pluginator on ReplicaCountTransformer; DO NOT EDIT.
// pluginator {unknown  1970-01-01T00:00:00Z  }

package builtins

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/transform"

	"sigs.k8s.io/kustomize/api/resid"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

// Find matching replicas declarations and replace the count.
// Eases the kustomization configuration of replica changes.
type ReplicaCountTransformerPlugin struct {
	Replica    types.Replica     `json:"replica,omitempty" yaml:"replica,omitempty"`
	FieldSpecs []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

func (p *ReplicaCountTransformerPlugin) Config(
	h *resmap.PluginHelpers, c []byte) (err error) {

	p.Replica = types.Replica{}
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *ReplicaCountTransformerPlugin) Transform(m resmap.ResMap) error {
	found := false
	for i, replicaSpec := range p.FieldSpecs {
		matcher := p.createMatcher(i)
		matchOriginal := m.GetMatchingResourcesByOriginalId(matcher)
		matchCurrent := m.GetMatchingResourcesByCurrentId(matcher)

		for _, res := range append(matchOriginal, matchCurrent...) {
			found = true
			err := transform.MutateField(
				res.Map(), replicaSpec.PathSlice(),
				replicaSpec.CreateIfNotPresent, p.addReplicas)
			if err != nil {
				return err
			}
		}
	}

	if !found {
		gvks := make([]string, len(p.FieldSpecs))
		for i, replicaSpec := range p.FieldSpecs {
			gvks[i] = replicaSpec.Gvk.String()
		}
		return fmt.Errorf("resource with name %s does not match a config with the following GVK %v",
			p.Replica.Name, gvks)
	}

	return nil
}

// Match Replica.Name and FieldSpec
func (p *ReplicaCountTransformerPlugin) createMatcher(i int) resmap.IdMatcher {
	return func(r resid.ResId) bool {
		return r.Name == p.Replica.Name &&
			r.Gvk.IsSelected(&p.FieldSpecs[i].Gvk)
	}
}

func (p *ReplicaCountTransformerPlugin) addReplicas(in interface{}) (interface{}, error) {
	switch m := in.(type) {
	case int64:
		// Was already in the field.
	case map[string]interface{}:
		if len(m) != 0 {
			// A map was already in the replicas field, don't want to
			// discard this data silently.
			return nil, fmt.Errorf("%#v is expected to be %T", in, m)
		}
		// Just got added, default type is map, but we can return anything.
	default:
		return nil, fmt.Errorf("%#v is expected to be %T", in, m)
	}
	return p.Replica.Count, nil
}

func NewReplicaCountTransformerPlugin() resmap.TransformerPlugin {
	return &ReplicaCountTransformerPlugin{}
}
