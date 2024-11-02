package sprout

// RegistryGroup represents a collection of registries that can be registered
// together within a Handler.
//
// This struct simplifies the management of registries by allowing them to be grouped
// and added in bulk, enhancing modularity and organization within the Handler.
type RegistryGroup struct {
	Registries []Registry
}

// NewRegistryGroup initializes a RegistryGroup with the specified registries.
//
// Example:
//
//	group := NewRegistryGroup(registry1.NewRegistry(), registry2.NewRegistry())
func NewRegistryGroup(registries ...Registry) *RegistryGroup {
	if len(registries) == 0 {
		registries = make([]Registry, 0)
	}

	return &RegistryGroup{
		Registries: registries,
	}
}

// AddGroups registers one or multiple RegistryGroups within the DefaultHandler.
//
// This method allows grouped registries to be added collectively. Each registry
// within each group is registered individually, simplifying large-scale registry
// addition.
//
// Example:
//
//	handler.AddGroups(group1, group2)
func (dh *DefaultHandler) AddGroups(groups ...*RegistryGroup) error {
	for _, group := range groups {
		for _, registry := range group.Registries {
			if err := dh.AddRegistry(registry); err != nil {
				return err
			}
		}
	}
	return nil
}

// WithGroups provides a HandlerOption to configure a DefaultHandler with the
// specified RegistryGroups.
//
// This option allows for the addition of grouped registries during Handler
// creation, enhancing organization and simplifying the registration of multiple
// registries in one step.
//
// Example:
//
//	handler := New(WithGroups(group1, group2))
func WithGroups(groups ...*RegistryGroup) HandlerOption[*DefaultHandler] {
	return func(dh *DefaultHandler) error {
		return dh.AddGroups(groups...)
	}
}
