/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

func defaultIBMPowerVSMachineSpec(spec *IBMPowerVSMachineSpec) {

	if spec.Memory == "" {
		spec.Memory = "8"
	}
	if spec.Processors == "" {
		spec.Processors = "0.25"
	}
	if spec.SysType == "" {
		spec.SysType = "s922"
	}
	if spec.ProcType == "" {
		spec.ProcType = "shared"
	}
}

func validateIBMPowerVSSysType(spec IBMPowerVSMachineSpec) (bool, IBMPowerVSMachineSpec) {
	sysTypes := [...]string{"s922", "e980"}
	for _, st := range sysTypes {
		if spec.SysType == st {
			return true, IBMPowerVSMachineSpec{}
		}
	}
	return false, spec
}

func validateIBMPowerVSProcType(spec IBMPowerVSMachineSpec) (bool, IBMPowerVSMachineSpec) {
	procTypes := [...]string{"shared", "dedicated", "capped"}
	for _, pt := range procTypes {
		if spec.ProcType == pt {
			return true, IBMPowerVSMachineSpec{}
		}
	}
	return false, spec
}

func validateIBMPowerVSNetwork(network IBMPowerVSResourceReference) (bool, IBMPowerVSResourceReference) {
	if network.ID != nil && network.Name != nil {
		return false, network
	}
	return true, IBMPowerVSResourceReference{}
}