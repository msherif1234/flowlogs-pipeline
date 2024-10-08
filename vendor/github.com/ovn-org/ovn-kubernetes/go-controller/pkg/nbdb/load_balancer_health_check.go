// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import "github.com/ovn-org/libovsdb/model"

const LoadBalancerHealthCheckTable = "Load_Balancer_Health_Check"

// LoadBalancerHealthCheck defines an object in Load_Balancer_Health_Check table
type LoadBalancerHealthCheck struct {
	UUID        string            `ovsdb:"_uuid"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	Options     map[string]string `ovsdb:"options"`
	Vip         string            `ovsdb:"vip"`
}

func (a *LoadBalancerHealthCheck) GetUUID() string {
	return a.UUID
}

func (a *LoadBalancerHealthCheck) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyLoadBalancerHealthCheckExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLoadBalancerHealthCheckExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *LoadBalancerHealthCheck) GetOptions() map[string]string {
	return a.Options
}

func copyLoadBalancerHealthCheckOptions(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLoadBalancerHealthCheckOptions(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *LoadBalancerHealthCheck) GetVip() string {
	return a.Vip
}

func (a *LoadBalancerHealthCheck) DeepCopyInto(b *LoadBalancerHealthCheck) {
	*b = *a
	b.ExternalIDs = copyLoadBalancerHealthCheckExternalIDs(a.ExternalIDs)
	b.Options = copyLoadBalancerHealthCheckOptions(a.Options)
}

func (a *LoadBalancerHealthCheck) DeepCopy() *LoadBalancerHealthCheck {
	b := new(LoadBalancerHealthCheck)
	a.DeepCopyInto(b)
	return b
}

func (a *LoadBalancerHealthCheck) CloneModelInto(b model.Model) {
	c := b.(*LoadBalancerHealthCheck)
	a.DeepCopyInto(c)
}

func (a *LoadBalancerHealthCheck) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *LoadBalancerHealthCheck) Equals(b *LoadBalancerHealthCheck) bool {
	return a.UUID == b.UUID &&
		equalLoadBalancerHealthCheckExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		equalLoadBalancerHealthCheckOptions(a.Options, b.Options) &&
		a.Vip == b.Vip
}

func (a *LoadBalancerHealthCheck) EqualsModel(b model.Model) bool {
	c := b.(*LoadBalancerHealthCheck)
	return a.Equals(c)
}

var _ model.CloneableModel = &LoadBalancerHealthCheck{}
var _ model.ComparableModel = &LoadBalancerHealthCheck{}