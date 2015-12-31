package api

import (
	"fmt"
)

// fixedStore provides fixed Machine, Spec, and config resources for testing
// purposes.
type fixedStore struct {
	Machines     map[string]*Machine
	Specs        map[string]*Spec
	CloudConfigs map[string]*CloudConfig
}

func (s *fixedStore) Machine(id string) (*Machine, error) {
	if machine, present := s.Machines[id]; present {
		return machine, nil
	}
	return nil, fmt.Errorf("no machine config %s", id)
}

func (s *fixedStore) Spec(id string) (*Spec, error) {
	if spec, present := s.Specs[id]; present {
		return spec, nil
	}
	return nil, fmt.Errorf("no spec %s", id)
}

func (s *fixedStore) CloudConfig(id string) (*CloudConfig, error) {
	if config, present := s.CloudConfigs[id]; present {
		return config, nil
	}
	return nil, fmt.Errorf("no cloud config %s", id)
}

type emptyStore struct{}

func (s *emptyStore) Machine(id string) (*Machine, error) {
	return nil, fmt.Errorf("no machine config %s", id)
}

func (s *emptyStore) Spec(id string) (*Spec, error) {
	return nil, fmt.Errorf("no group config %s", id)
}

func (s emptyStore) CloudConfig(id string) (*CloudConfig, error) {
	return nil, fmt.Errorf("no cloud config %s", id)
}
