//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	apiv1 "github.com/Opster/opensearch-k8s-operator/opensearch-operator/api/v1"
	"github.com/rancher/opni/pkg/util/meta"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatorOTELConfigSpec) DeepCopyInto(out *AggregatorOTELConfigSpec) {
	*out = *in
	out.Processors = in.Processors
	out.Exporters = in.Exporters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatorOTELConfigSpec.
func (in *AggregatorOTELConfigSpec) DeepCopy() *AggregatorOTELConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AggregatorOTELConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatorOTELExporters) DeepCopyInto(out *AggregatorOTELExporters) {
	*out = *in
	out.OTLPHTTP = in.OTLPHTTP
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatorOTELExporters.
func (in *AggregatorOTELExporters) DeepCopy() *AggregatorOTELExporters {
	if in == nil {
		return nil
	}
	out := new(AggregatorOTELExporters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatorOTELProcessors) DeepCopyInto(out *AggregatorOTELProcessors) {
	*out = *in
	out.Batch = in.Batch
	out.MemoryLimiter = in.MemoryLimiter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatorOTELProcessors.
func (in *AggregatorOTELProcessors) DeepCopy() *AggregatorOTELProcessors {
	if in == nil {
		return nil
	}
	out := new(AggregatorOTELProcessors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManagerSpec) DeepCopyInto(out *AlertManagerSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(meta.ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	in.ApplicationSpec.DeepCopyInto(&out.ApplicationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManagerSpec.
func (in *AlertManagerSpec) DeepCopy() *AlertManagerSpec {
	if in == nil {
		return nil
	}
	out := new(AlertManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManagerStatus) DeepCopyInto(out *AlertManagerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManagerStatus.
func (in *AlertManagerStatus) DeepCopy() *AlertManagerStatus {
	if in == nil {
		return nil
	}
	out := new(AlertManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingApplicationSpec) DeepCopyInto(out *AlertingApplicationSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraVolumeSpec != nil {
		in, out := &in.ExtraVolumeSpec, &out.ExtraVolumeSpec
		*out = make([]VolumeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraEnvVars != nil {
		in, out := &in.ExtraEnvVars, &out.ExtraEnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceRequirements != nil {
		in, out := &in.ResourceRequirements, &out.ResourceRequirements
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.StatefulSetUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = new(apiv1.PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingApplicationSpec.
func (in *AlertingApplicationSpec) DeepCopy() *AlertingApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(AlertingApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingCluster) DeepCopyInto(out *AlertingCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingCluster.
func (in *AlertingCluster) DeepCopy() *AlertingCluster {
	if in == nil {
		return nil
	}
	out := new(AlertingCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertingCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingClusterList) DeepCopyInto(out *AlertingClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertingCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingClusterList.
func (in *AlertingClusterList) DeepCopy() *AlertingClusterList {
	if in == nil {
		return nil
	}
	out := new(AlertingClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertingClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingClusterSpec) DeepCopyInto(out *AlertingClusterSpec) {
	*out = *in
	out.Gateway = in.Gateway
	in.Alertmanager.DeepCopyInto(&out.Alertmanager)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingClusterSpec.
func (in *AlertingClusterSpec) DeepCopy() *AlertingClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AlertingClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingClusterStatus) DeepCopyInto(out *AlertingClusterStatus) {
	*out = *in
	in.Alertmanager.DeepCopyInto(&out.Alertmanager)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingClusterStatus.
func (in *AlertingClusterStatus) DeepCopy() *AlertingClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AlertingClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingSpec) DeepCopyInto(out *AlertingSpec) {
	*out = *in
	if in.GatewayVolumeMounts != nil {
		in, out := &in.GatewayVolumeMounts, &out.GatewayVolumeMounts
		*out = make([]meta.ExtraVolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingSpec.
func (in *AlertingSpec) DeepCopy() *AlertingSpec {
	if in == nil {
		return nil
	}
	out := new(AlertingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchProcessorConfig) DeepCopyInto(out *BatchProcessorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchProcessorConfig.
func (in *BatchProcessorConfig) DeepCopy() *BatchProcessorConfig {
	if in == nil {
		return nil
	}
	out := new(BatchProcessorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapToken) DeepCopyInto(out *BootstrapToken) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapToken.
func (in *BootstrapToken) DeepCopy() *BootstrapToken {
	if in == nil {
		return nil
	}
	out := new(BootstrapToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapTokenList) DeepCopyInto(out *BootstrapTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BootstrapToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapTokenList.
func (in *BootstrapTokenList) DeepCopy() *BootstrapTokenList {
	if in == nil {
		return nil
	}
	out := new(BootstrapTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collector) DeepCopyInto(out *Collector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collector.
func (in *Collector) DeepCopy() *Collector {
	if in == nil {
		return nil
	}
	out := new(Collector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Collector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorList) DeepCopyInto(out *CollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Collector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorList.
func (in *CollectorList) DeepCopy() *CollectorList {
	if in == nil {
		return nil
	}
	out := new(CollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSendingQueue) DeepCopyInto(out *CollectorSendingQueue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSendingQueue.
func (in *CollectorSendingQueue) DeepCopy() *CollectorSendingQueue {
	if in == nil {
		return nil
	}
	out := new(CollectorSendingQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSpec) DeepCopyInto(out *CollectorSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.TracesConfig != nil {
		in, out := &in.TracesConfig, &out.TracesConfig
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.MetricsConfig != nil {
		in, out := &in.MetricsConfig, &out.MetricsConfig
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.ConfigReloader != nil {
		in, out := &in.ConfigReloader, &out.ConfigReloader
		*out = new(ConfigReloaderSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AggregatorOTELConfigSpec != nil {
		in, out := &in.AggregatorOTELConfigSpec, &out.AggregatorOTELConfigSpec
		*out = new(AggregatorOTELConfigSpec)
		**out = **in
	}
	if in.NodeOTELConfigSpec != nil {
		in, out := &in.NodeOTELConfigSpec, &out.NodeOTELConfigSpec
		*out = new(NodeOTELConfigSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSpec.
func (in *CollectorSpec) DeepCopy() *CollectorSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorStatus) DeepCopyInto(out *CollectorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorStatus.
func (in *CollectorStatus) DeepCopy() *CollectorStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigReloaderSpec) DeepCopyInto(out *ConfigReloaderSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigReloaderSpec.
func (in *ConfigReloaderSpec) DeepCopy() *ConfigReloaderSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigReloaderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CortexSpec) DeepCopyInto(out *CortexSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.CortexWorkloads != nil {
		in, out := &in.CortexWorkloads, &out.CortexWorkloads
		*out = (*in).DeepCopy()
	}
	if in.CortexConfig != nil {
		in, out := &in.CortexConfig, &out.CortexConfig
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CortexSpec.
func (in *CortexSpec) DeepCopy() *CortexSpec {
	if in == nil {
		return nil
	}
	out := new(CortexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CortexStatus) DeepCopyInto(out *CortexStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkloadStatus != nil {
		in, out := &in.WorkloadStatus, &out.WorkloadStatus
		*out = make(map[string]WorkloadStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CortexStatus.
func (in *CortexStatus) DeepCopy() *CortexStatus {
	if in == nil {
		return nil
	}
	out := new(CortexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStorageSpec) DeepCopyInto(out *FileStorageSpec) {
	*out = *in
	in.JetStreamPersistenceSpec.DeepCopyInto(&out.JetStreamPersistenceSpec)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStorageSpec.
func (in *FileStorageSpec) DeepCopy() *FileStorageSpec {
	if in == nil {
		return nil
	}
	out := new(FileStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaSpec) DeepCopyInto(out *GrafanaSpec) {
	*out = *in
	if in.GrafanaConfig != nil {
		in, out := &in.GrafanaConfig, &out.GrafanaConfig
		*out = (*in).DeepCopy()
	}
	in.GrafanaSpec.DeepCopyInto(&out.GrafanaSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaSpec.
func (in *GrafanaSpec) DeepCopy() *GrafanaSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JetStreamPersistenceSpec) DeepCopyInto(out *JetStreamPersistenceSpec) {
	*out = *in
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(PVCSource)
		(*in).DeepCopyInto(*out)
	}
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(v1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JetStreamPersistenceSpec.
func (in *JetStreamPersistenceSpec) DeepCopy() *JetStreamPersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(JetStreamPersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JetStreamSpec) DeepCopyInto(out *JetStreamSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	out.MemoryStorageSize = in.MemoryStorageSize.DeepCopy()
	in.FileStorage.DeepCopyInto(&out.FileStorage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JetStreamSpec.
func (in *JetStreamSpec) DeepCopy() *JetStreamSpec {
	if in == nil {
		return nil
	}
	out := new(JetStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keyring) DeepCopyInto(out *Keyring) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keyring.
func (in *Keyring) DeepCopy() *Keyring {
	if in == nil {
		return nil
	}
	out := new(Keyring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Keyring) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyringList) DeepCopyInto(out *KeyringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Keyring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyringList.
func (in *KeyringList) DeepCopy() *KeyringList {
	if in == nil {
		return nil
	}
	out := new(KeyringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingCluster) DeepCopyInto(out *LoggingCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingCluster.
func (in *LoggingCluster) DeepCopy() *LoggingCluster {
	if in == nil {
		return nil
	}
	out := new(LoggingCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingClusterList) DeepCopyInto(out *LoggingClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoggingCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingClusterList.
func (in *LoggingClusterList) DeepCopy() *LoggingClusterList {
	if in == nil {
		return nil
	}
	out := new(LoggingClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingClusterSpec) DeepCopyInto(out *LoggingClusterSpec) {
	*out = *in
	if in.OpensearchClusterRef != nil {
		in, out := &in.OpensearchClusterRef, &out.OpensearchClusterRef
		*out = new(meta.OpensearchClusterRef)
		**out = **in
	}
	if in.IndexUserSecret != nil {
		in, out := &in.IndexUserSecret, &out.IndexUserSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	in.LastSync.DeepCopyInto(&out.LastSync)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingClusterSpec.
func (in *LoggingClusterSpec) DeepCopy() *LoggingClusterSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingClusterStatus) DeepCopyInto(out *LoggingClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingClusterStatus.
func (in *LoggingClusterStatus) DeepCopy() *LoggingClusterStatus {
	if in == nil {
		return nil
	}
	out := new(LoggingClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryLimiterProcessorConfig) DeepCopyInto(out *MemoryLimiterProcessorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryLimiterProcessorConfig.
func (in *MemoryLimiterProcessorConfig) DeepCopy() *MemoryLimiterProcessorConfig {
	if in == nil {
		return nil
	}
	out := new(MemoryLimiterProcessorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringCluster) DeepCopyInto(out *MonitoringCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringCluster.
func (in *MonitoringCluster) DeepCopy() *MonitoringCluster {
	if in == nil {
		return nil
	}
	out := new(MonitoringCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringClusterList) DeepCopyInto(out *MonitoringClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringClusterList.
func (in *MonitoringClusterList) DeepCopy() *MonitoringClusterList {
	if in == nil {
		return nil
	}
	out := new(MonitoringClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringClusterSpec) DeepCopyInto(out *MonitoringClusterSpec) {
	*out = *in
	out.Gateway = in.Gateway
	in.Cortex.DeepCopyInto(&out.Cortex)
	in.Grafana.DeepCopyInto(&out.Grafana)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringClusterSpec.
func (in *MonitoringClusterSpec) DeepCopy() *MonitoringClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringClusterStatus) DeepCopyInto(out *MonitoringClusterStatus) {
	*out = *in
	in.Cortex.DeepCopyInto(&out.Cortex)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringClusterStatus.
func (in *MonitoringClusterStatus) DeepCopy() *MonitoringClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsCluster) DeepCopyInto(out *NatsCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsCluster.
func (in *NatsCluster) DeepCopy() *NatsCluster {
	if in == nil {
		return nil
	}
	out := new(NatsCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NatsCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsClusterList) DeepCopyInto(out *NatsClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NatsCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsClusterList.
func (in *NatsClusterList) DeepCopy() *NatsClusterList {
	if in == nil {
		return nil
	}
	out := new(NatsClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NatsClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsSpec) DeepCopyInto(out *NatsSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.PasswordFrom != nil {
		in, out := &in.PasswordFrom, &out.PasswordFrom
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.JetStream.DeepCopyInto(&out.JetStream)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsSpec.
func (in *NatsSpec) DeepCopy() *NatsSpec {
	if in == nil {
		return nil
	}
	out := new(NatsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsStatus) DeepCopyInto(out *NatsStatus) {
	*out = *in
	if in.AuthSecretKeyRef != nil {
		in, out := &in.AuthSecretKeyRef, &out.AuthSecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsStatus.
func (in *NatsStatus) DeepCopy() *NatsStatus {
	if in == nil {
		return nil
	}
	out := new(NatsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOTELConfigSpec) DeepCopyInto(out *NodeOTELConfigSpec) {
	*out = *in
	out.Processors = in.Processors
	out.Exporters = in.Exporters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOTELConfigSpec.
func (in *NodeOTELConfigSpec) DeepCopy() *NodeOTELConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeOTELConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOTELExporters) DeepCopyInto(out *NodeOTELExporters) {
	*out = *in
	out.OTLP = in.OTLP
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOTELExporters.
func (in *NodeOTELExporters) DeepCopy() *NodeOTELExporters {
	if in == nil {
		return nil
	}
	out := new(NodeOTELExporters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOTELProcessors) DeepCopyInto(out *NodeOTELProcessors) {
	*out = *in
	out.MemoryLimiter = in.MemoryLimiter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOTELProcessors.
func (in *NodeOTELProcessors) DeepCopy() *NodeOTELProcessors {
	if in == nil {
		return nil
	}
	out := new(NodeOTELProcessors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OTLPExporterConfig) DeepCopyInto(out *OTLPExporterConfig) {
	*out = *in
	out.SendingQueue = in.SendingQueue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OTLPExporterConfig.
func (in *OTLPExporterConfig) DeepCopy() *OTLPExporterConfig {
	if in == nil {
		return nil
	}
	out := new(OTLPExporterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OTLPHTTPExporterConfig) DeepCopyInto(out *OTLPHTTPExporterConfig) {
	*out = *in
	out.SendingQueue = in.SendingQueue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OTLPHTTPExporterConfig.
func (in *OTLPHTTPExporterConfig) DeepCopy() *OTLPHTTPExporterConfig {
	if in == nil {
		return nil
	}
	out := new(OTLPHTTPExporterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCSource) DeepCopyInto(out *PVCSource) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCSource.
func (in *PVCSource) DeepCopy() *PVCSource {
	if in == nil {
		return nil
	}
	out := new(PVCSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
	in.VolumeMount.DeepCopyInto(&out.VolumeMount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}
