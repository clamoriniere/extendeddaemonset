// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package extendeddaemonset

import (
	"testing"
	"time"

	datadoghqv1alpha1 "github.com/datadog/extendeddaemonset/pkg/apis/datadoghq/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsCanaryDeploymentPaused(t *testing.T) {
	type args struct {
		dsAnnotations map[string]string
	}

	tests := []struct {
		name       string
		args       args
		want       bool
		wantReason datadoghqv1alpha1.ExtendedDaemonSetStatusReason
	}{
		{
			name: "pause annotation is false, expect false",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-paused": "false",
				},
			},
			want:       false,
			wantReason: "",
		},
		{
			name: "pause annotation doesn't exist, expect false",
			args: args{
				dsAnnotations: map[string]string{},
			},
			want:       false,
			wantReason: "",
		},
		{
			name: "pause annotation is `truee`, expect false",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-paused": "truee",
				},
			},
			want:       false,
			wantReason: "",
		},
		{
			name: "pause annotation is `true` and has a reason, expect true and the reason",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-paused":        "true",
					"extendeddaemonset.datadoghq.com/canary-paused-reason": string(datadoghqv1alpha1.ExtendedDaemonSetStatusReasonOOM),
				},
			},
			want:       true,
			wantReason: datadoghqv1alpha1.ExtendedDaemonSetStatusReasonOOM,
		},
		{
			name: "pause annotation is `true` and has no reason, expect true and `unknown` reason",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-paused": "true",
				},
			},
			want:       true,
			wantReason: datadoghqv1alpha1.ExtendedDaemonSetStatusReasonUnknown,
		},
		{
			name: "pause annotation is `true` and has an unsupported reason, expect true and `unknown` reason",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-paused":        "true",
					"extendeddaemonset.datadoghq.com/canary-paused-reason": "SomeUnsupportedReason",
				},
			},
			want:       true,
			wantReason: datadoghqv1alpha1.ExtendedDaemonSetStatusReasonUnknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotReason := IsCanaryDeploymentPaused(tt.args.dsAnnotations)
			if got != tt.want {
				t.Errorf("IsCanaryDeploymentPaused() = %v, want %v", got, tt.want)
			}
			if gotReason != tt.wantReason {
				t.Errorf("IsCanaryDeploymentePaused() = %v, wantReason %v", gotReason, tt.wantReason)
			}
		})
	}
}

func TestIsCanaryDeploymentEnded(t *testing.T) {
	now := time.Now()
	type args struct {
		specCanary *datadoghqv1alpha1.ExtendedDaemonSetSpecStrategyCanary
		rs         *datadoghqv1alpha1.ExtendedDaemonSetReplicaSet
		now        time.Time
	}
	tests := []struct {
		name         string
		args         args
		want         bool
		wantDuration time.Duration
	}{
		{
			name: "not spec == nil",
			args: args{
				specCanary: nil,
				rs:         &datadoghqv1alpha1.ExtendedDaemonSetReplicaSet{},
				now:        now,
			},
			want: true,
		},
		{
			name: "not canary not done",
			args: args{
				specCanary: &datadoghqv1alpha1.ExtendedDaemonSetSpecStrategyCanary{
					Duration: &metav1.Duration{Duration: time.Hour},
				},
				rs: &datadoghqv1alpha1.ExtendedDaemonSetReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						CreationTimestamp: metav1.NewTime(now.Add(-time.Minute)),
					},
				},
				now: now,
			},
			want:         false,
			wantDuration: 59 * time.Minute,
		},
		{
			name: "not canary duration not set",
			args: args{
				specCanary: &datadoghqv1alpha1.ExtendedDaemonSetSpecStrategyCanary{},
				rs: &datadoghqv1alpha1.ExtendedDaemonSetReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						CreationTimestamp: metav1.NewTime(now.Add(-time.Minute)),
					},
				},
				now: now,
			},
			want: false,
		},
		{
			name: "not canary done",
			args: args{
				specCanary: &datadoghqv1alpha1.ExtendedDaemonSetSpecStrategyCanary{
					Duration: &metav1.Duration{Duration: time.Hour},
				},
				rs: &datadoghqv1alpha1.ExtendedDaemonSetReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						CreationTimestamp: metav1.NewTime(now.Add(-2 * time.Hour)),
					},
				},
				now: now,
			},
			want:         true,
			wantDuration: -time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotDuration := IsCanaryDeploymentEnded(tt.args.specCanary, tt.args.rs, tt.args.now)
			if got != tt.want {
				t.Errorf("IsCanaryDeploymentEnded() = %v, want %v", got, tt.want)
			}
			if gotDuration != tt.wantDuration {
				t.Errorf("IsCanaryDeploymenteEnded() = %v, wantDuration %v", gotDuration, tt.wantDuration)
			}
		})
	}
}

func TestIsCanaryDeploymentValid(t *testing.T) {
	type args struct {
		dsAnnotations map[string]string
		rsName        string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "annotation found - correct rs name",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-valid": "rsName",
				},
				rsName: "rsName",
			},
			want: true,
		},
		{
			name: "annotation found - incorrect rs name",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-valid": "rsName",
				},
				rsName: "anotherRsName",
			},
			want: false,
		},
		{
			name: "annotation not found",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/another-annotation": "rsName",
				},
				rsName: "rsName",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCanaryDeploymentValid(tt.args.dsAnnotations, tt.args.rsName); got != tt.want {
				t.Errorf("IsCanaryDeploymentValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCanaryDeploymentFailed(t *testing.T) {
	type args struct {
		dsAnnotations map[string]string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "annotation found - correct rs name",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-failed": "true",
				},
			},
			want: true,
		},
		{
			name: "annotation found - incorrect rs name",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-failed": "false",
				},
			},
			want: false,
		},
		{
			name: "annotation not found",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-failed": "random",
				},
			},
			want: false,
		},
		{
			name: "annotation not found",
			args: args{
				dsAnnotations: map[string]string{
					"extendeddaemonset.datadoghq.com/canary-something-else": "true",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCanaryDeploymentFailed(tt.args.dsAnnotations); got != tt.want {
				t.Errorf("IsCanaryDeploymentFailed() = %v, want %v", got, tt.want)
			}
		})
	}
}
