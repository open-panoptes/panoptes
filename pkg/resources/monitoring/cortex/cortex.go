package cortex

import (
	"context"

	"log/slog"

	"github.com/cisco-open/operator-tools/pkg/reconciler"
	corev1beta1 "github.com/open-panoptes/opni/apis/core/v1beta1"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/resources"
	"github.com/open-panoptes/opni/pkg/util/k8sutil"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
	reconciler.ResourceReconciler
	ctx    context.Context
	client client.Client
	lg     *slog.Logger
	mc     *corev1beta1.MonitoringCluster
}

func NewReconciler(
	ctx context.Context,
	client client.Client,
	mc *corev1beta1.MonitoringCluster,
) *Reconciler {
	return &Reconciler{
		ResourceReconciler: reconciler.NewReconcilerWith(client,
			reconciler.WithEnableRecreateWorkload(),
			reconciler.WithRecreateErrorMessageCondition(reconciler.MatchImmutableErrorMessages),
			reconciler.WithLog(log.FromContext(ctx)),
			reconciler.WithScheme(client.Scheme()),
		),
		ctx:    ctx,
		client: client,
		mc:     mc,
		lg:     logger.New().WithGroup("controller").WithGroup("cortex"),
	}

}

func (r *Reconciler) Reconcile() (*reconcile.Result, error) {
	allResources := []resources.Resource{}

	updated, err := r.updateCortexVersionStatus()
	if err != nil {
		return k8sutil.RequeueErr(err).ResultPtr()
	}
	if updated {
		return k8sutil.Requeue().ResultPtr()
	}

	configs, configDigest, err := r.config()
	if err != nil {
		return k8sutil.RequeueErr(err).ResultPtr()
	}
	allResources = append(allResources, configs...)

	fallbackConfig := r.alertmanagerFallbackConfig()
	allResources = append(allResources, fallbackConfig)

	serviceAccount := r.serviceAccount()
	allResources = append(allResources, serviceAccount)

	workloads := r.cortexWorkloads(configDigest)
	allResources = append(allResources, workloads...)

	services := r.services()
	allResources = append(allResources, services...)

	if op := resources.ReconcileAll(r, allResources); op.ShouldRequeue() {
		return op.ResultPtr()
	}

	// watch cortex components until they are healthy
	if op := r.pollCortexHealth(workloads); op.ShouldRequeue() {
		return op.ResultPtr()
	}

	return nil, nil
}
