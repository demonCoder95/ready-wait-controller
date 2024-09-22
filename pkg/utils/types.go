package utils

// DeploymentStatus represents the status of a Deployment
type DeploymentStatus struct {
	// The number of total replicas
	Replicas int32
	// The number of ready replicas
	ReadyReplicas int32
	// The number of updated replicas
	UpdatedReplicas int32
	// The number of available replicas
	AvailableReplicas int32
}

// StatefulSetStatus represents the status of a StatefulSet
type StatefulSetStatus struct {
	// The number of total replicas
	Replicas int32
	// The number of ready replicas
	ReadyReplicas int32
	// The number of updated replicas
	UpdatedReplicas int32
	// The number of available replicas
	AvailableReplicas int32
	// The number of current replicas
	CurrentReplicas int32
}

// DaemonSetStatus represents the status of a DaemonSet
type DaemonSetStatus struct {
	// The number of desired replicas
	DesiredReplicas int32
	// The number of current replicas
	CurrentReplicas int32
	// The number of ready replicas
	ReadyReplicas int32
	// The number of Up-to-date replicas
	UpToDateReplicas int32
	// The number of available replicas
	AvailableReplicas int32
}
