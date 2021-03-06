package controller

import (
	"ovn4nfv-k8s-plugin/pkg/controller/providernetwork"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, providernetwork.Add)
}
