package randomscore

import (
	"math/rand"
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)
type Plugin interface {
	Name() string
}

type ScoreExtensions interface {
	NormalizeScore(ctx context.Context, state *CycleState, p *v1.Pod, scores NodeScoreList) *Status
}

type ScorePlugin interface {
	Plugin

	Score(ctx context.Context, state *CycleState, p *v1.Pod, nodeName string) (int64, *Status)

	ScoreExtensions() ScoreExtensions
}

type RandomScore struct {
	handle framework.FrameworkHandle
}

const pluginName = "RandomScore"

var _ = framework.ScorePlugin(&RandomScore{})

func New(obj runtime.Object, h framework.FrameworkHandle) (framework.Plugin, error) {
	args, ok := obj.(*config.RandomScoreArgs)
	if !ok {
		return nil, fmt.Errorf("want args to be of type RandomScoreArgs, got %T", obj)
	}
	
	return &RandomScore{
		handle:     h,
	}, nil
}

func (r *RandomScore) Name() string {
	return pluginName
}


func (r *RandomScore) Score(ctx context.Context, state *CycleState, p *v1.Pod, nodeName string) (int64, *Status) {
	randomScore := rand.Int63n(101)

	klog.Infof("[RandomScore] node '%s' score: %s", nodeName, randomScore)
	return randomScore, nil
}

