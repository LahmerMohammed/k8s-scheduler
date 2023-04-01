import (
    "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

type Plugin interface {
	Name() string
}

type QueueSortPlugin interface {
	Plugin

	Less(*QueuePodInfo, *QueuePodInfo) bool
}

type PreFilterPlugin interface {
	Plugin

	PreFilter(ctx context.Context, state *CycleState, p *v1.pod) *Status
}

type FilterPlugin {
	Plugin
	
	Filter(ctx context.Context, state *CycleState, pod *v1.Pod, nodeInfo *NodeInfo) *Status
}


type PostFilterPlugin {
	Plugin

	PostFilter(ctx context.Context, state *CycleState, pod *v1.Pod, filteredNodeStatusMap NodeToStatusMap) (*PostFilterResult, *Status)
}


type PreScorePlugin {
	Plugin

	PreScore(ctx context.Context, state *CycleState, pod *v1.Pod, nodes []*v1.Node) *Status
}

type ScorePlugin {
	Plugin

	Score(ctx context.Context, state *CycleState, p *v1.Pod, nodeName string) (int64, *Status)

	ScoreExtensions() ScoreExtensions
}
