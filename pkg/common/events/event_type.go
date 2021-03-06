package events

// EventType is the type of event to record
type EventType string

// This file defines different event types. Please add new types here so that we can easily track them.
const (
	// ------ Cluster provisioning events

	// InstallSuccessful when the cluster has successfully provisioned.
	InstallSuccessful EventType = "InstallSuccessful"

	// FailedClusterProvision when the cluster has not been successfully provisioned.
	InstallFailed EventType = "InstallFailed"

	// HealthCheckSuccessful when the cluster health check has succeeded.
	HealthCheckSuccessful EventType = "HealthCheckSuccessful"

	// HealthCheckFailed when the cluster health check has failed.
	HealthCheckFailed EventType = "HealthCheckFailed"

	// UpgradeSuccessful when the upgrade was successful.
	UpgradeSuccessful EventType = "UpgradeSuccessful"

	// UpgradeFailed when the upgrade failed.
	UpgradeFailed EventType = "UpgradeFailed"

	// NoHiveLogs when no logs from Hive were collected after a cluster provisioning event.
	NoHiveLogs EventType = "NoHiveLogs"

	// InstallKubeconfigRetrievalSuccess when the Kubeconfig was retrieved successfully.
	InstallKubeconfigRetrievalSuccess EventType = "InstallKubeconfigRetrievalSuccess"

	// InstallKubeconfigRetrievalFailure when the Kubeconfig was not retrieved successfully.
	InstallKubeconfigRetrievalFailure EventType = "InstallKubeconfigRetrievalFailure"

	// ------ Addon installation events

	// InstallAddonsSuccessful when the addons installed successfully.
	InstallAddonsSuccessful EventType = "InstallAddonsSuccessful"

	// InstallAddonsFailed when the addons failed to install.
	InstallAddonsFailed EventType = "InstallAddonsFailed"
)
