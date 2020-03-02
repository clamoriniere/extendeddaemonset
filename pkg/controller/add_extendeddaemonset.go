// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package controller

import (
	"github.com/datadog/extendeddaemonset/pkg/controller/extendeddaemonset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, extendeddaemonset.Add)
	AddToMetricsHandlerFuncs = append(AddToMetricsHandlerFuncs, extendeddaemonset.AddMetrics)
}
