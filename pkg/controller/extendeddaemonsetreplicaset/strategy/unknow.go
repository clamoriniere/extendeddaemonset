// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package strategy

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	podutils "github.com/datadog/extendeddaemonset/pkg/controller/utils/pod"
)

// ManageUnknow use to manage ReplicaSet with unknow status
func ManageUnknow(client client.Client, params *Parameters) (*Result, error) {
	result := &Result{}
	// remove canary node if define
	for _, nodeName := range params.CanaryNodes {
		delete(params.PodByNodeName, nodeName)
	}
	now := time.Now()
	metaNow := metav1.NewTime(now)
	var desiredPods, currentPods, availablePods, readyPods int32

	for _, pod := range params.PodByNodeName {
		desiredPods++
		if pod == nil {
		} else if compareSpecTemplateMD5Hash(params.Replicaset.Spec.TemplateGeneration, pod) {
			currentPods++
			if podutils.IsPodAvailable(pod, 0, metaNow) {
				availablePods++
			}
			if podutils.IsPodReady(pod) {

				readyPods++
			}
		}
	}

	result.NewStatus = params.NewStatus.DeepCopy()
	result.NewStatus.Status = "Unknow"
	result.NewStatus.Desired = 0
	result.NewStatus.Ready = readyPods
	result.NewStatus.Current = currentPods
	result.NewStatus.Available = availablePods
	params.Logger.V(1).Info("Status:", "Desired", result.NewStatus.Desired, "Ready", readyPods, "Available", availablePods)

	if result.NewStatus.Desired != result.NewStatus.Ready {
		result.Result.Requeue = true
	}
	result.Result.RequeueAfter = time.Second
	return result, nil
}
