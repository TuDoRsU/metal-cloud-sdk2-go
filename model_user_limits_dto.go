/*
 * MetalSoft REST API
 *
 * MetalSoft REST API documentation
 *
 * API version: 2.0
 * Contact: support@metalsoft.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk2

type UserLimitsDto struct {
	ContainerArrayContainersMaxCount float64 `json:"containerArrayContainersMaxCount,omitempty"`
	ContainerArrayContainersMinCount float64 `json:"containerArrayContainersMinCount,omitempty"`
	ContainerArrayDriveArraysMaxCount float64 `json:"containerArrayDriveArraysMaxCount,omitempty"`
	ContainerArrayDriveArraysMinCount float64 `json:"containerArrayDriveArraysMinCount,omitempty"`
	ContainerArraySecretsMaxCount float64 `json:"containerArraySecretsMaxCount,omitempty"`
	ContainerPlatformContainerArrayMaxCount float64 `json:"containerPlatformContainerArrayMaxCount,omitempty"`
	DriveArrayDrivesMaxCount float64 `json:"driveArrayDrivesMaxCount,omitempty"`
	DriveArrayDrivesMinCount float64 `json:"driveArrayDrivesMinCount,omitempty"`
	DriveMaxSizeMbytes float64 `json:"driveMaxSizeMbytes,omitempty"`
	DriveMinSizeMbytes float64 `json:"driveMinSizeMbytes,omitempty"`
	InfrastructureActiveMaxCount float64 `json:"infrastructureActiveMaxCount,omitempty"`
	InfrastructureClusterMaxCount float64 `json:"infrastructureClusterMaxCount,omitempty"`
	InfrastructureClusterMysqlAllowed bool `json:"infrastructureClusterMysqlAllowed,omitempty"`
	InfrastructureContainerClusterKafkaAllowed bool `json:"infrastructureContainerClusterKafkaAllowed,omitempty"`
	InfrastructureContainerClusterMaxCount float64 `json:"infrastructureContainerClusterMaxCount,omitempty"`
	InfrastructureContainerClusterPostgresqlAllowed bool `json:"infrastructureContainerClusterPostgresqlAllowed,omitempty"`
	InfrastructureContainerClusterSparkArrayAllowed bool `json:"infrastructureContainerClusterSparkArrayAllowed,omitempty"`
	InfrastructureContainerClusterSparksqlAllowed bool `json:"infrastructureContainerClusterSparksqlAllowed,omitempty"`
	InfrastructureContainerClusterStreamsetsAllowed bool `json:"infrastructureContainerClusterStreamsetsAllowed,omitempty"`
	InfrastructureContainerClusterZookeeperAllowed bool `json:"infrastructureContainerClusterZookeeperAllowed,omitempty"`
	InfrastructureContainerClusterZoomdataAllowed bool `json:"infrastructureContainerClusterZoomdataAllowed,omitempty"`
	InfrastructureContainerPlatformMaxCount float64 `json:"infrastructureContainerPlatformMaxCount,omitempty"`
	InfrastructureDataLakeEnabled bool `json:"infrastructureDataLakeEnabled,omitempty"`
	InfrastructureDataLakeMaxCount float64 `json:"infrastructureDataLakeMaxCount,omitempty"`
	InfrastructureDeletedMaxCount float64 `json:"infrastructureDeletedMaxCount,omitempty"`
	InfrastructureDriveArrayMaxCount float64 `json:"infrastructureDriveArrayMaxCount,omitempty"`
	InfrastructureInactiveMaxCount float64 `json:"infrastructureInactiveMaxCount,omitempty"`
	InfrastructureInstanceArrayMaxCount float64 `json:"infrastructureInstanceArrayMaxCount,omitempty"`
	InfrastructureLanMaxCount float64 `json:"infrastructureLanMaxCount,omitempty"`
	InfrastructureSanMaxCount float64 `json:"infrastructureSanMaxCount,omitempty"`
	InfrastructureSharedDriveMaxCount float64 `json:"infrastructureSharedDriveMaxCount,omitempty"`
	InfrastructureVolumeTemplateExperimentalAllowed bool `json:"infrastructureVolumeTemplateExperimentalAllowed,omitempty"`
	InfrastructureWanMaxCount float64 `json:"infrastructureWanMaxCount,omitempty"`
	InstanceArrayInstancesMaxCount float64 `json:"instanceArrayInstancesMaxCount,omitempty"`
	InstanceArrayInstancesMinCount float64 `json:"instanceArrayInstancesMinCount,omitempty"`
	InfrastructureVmInstanceGroupMaxCount float64 `json:"infrastructureVmInstanceGroupMaxCount,omitempty"`
	VmInstanceGroupVmInstancesMaxCount float64 `json:"vmInstanceGroupVmInstancesMaxCount,omitempty"`
	VmInstanceMaxDiskSizeMbytes float64 `json:"vmInstanceMaxDiskSizeMbytes,omitempty"`
	OwnerIsBillable bool `json:"ownerIsBillable,omitempty"`
	ServerTypeReservationMaxCount float64 `json:"serverTypeReservationMaxCount,omitempty"`
	ServerTypeReservationMaxQuantity float64 `json:"serverTypeReservationMaxQuantity,omitempty"`
	SharedDriveMaxSizeMbytes float64 `json:"sharedDriveMaxSizeMbytes,omitempty"`
	SharedDriveMinSizeMbytes float64 `json:"sharedDriveMinSizeMbytes,omitempty"`
	AllowVlanOverrides bool `json:"allowVlanOverrides,omitempty"`
	AllowNetworkProfiles bool `json:"allowNetworkProfiles,omitempty"`
	ShowOperatingSystemImagesTab bool `json:"showOperatingSystemImagesTab,omitempty"`
	ShowTemplateAssetsView bool `json:"showTemplateAssetsView,omitempty"`
	StorageTypes []string `json:"storageTypes,omitempty"`
	ThresholdMaxCount float64 `json:"thresholdMaxCount,omitempty"`
	UserResourceIscsiStorageSpaceMaxGbytes float64 `json:"userResourceIscsiStorageSpaceMaxGbytes,omitempty"`
	UserResourceServersMaxCount float64 `json:"userResourceServersMaxCount,omitempty"`
	UserResourceServerTypeNameToMaxCount *interface{} `json:"userResourceServerTypeNameToMaxCount,omitempty"`
	UserSshKeysCountMax float64 `json:"userSshKeysCountMax,omitempty"`
	WanSubnetIpv4MaxCount float64 `json:"wanSubnetIpv4MaxCount,omitempty"`
	WanSubnetIpv6MaxCount float64 `json:"wanSubnetIpv6MaxCount,omitempty"`
	WanSubnetPrefixSizeToMaxCount *interface{} `json:"wanSubnetPrefixSizeToMaxCount,omitempty"`
	ShowLegacyPages bool `json:"showLegacyPages,omitempty"`
	ShowExperimentalPages bool `json:"showExperimentalPages,omitempty"`
	ShowDiagramAppsGlobal bool `json:"showDiagramAppsGlobal,omitempty"`
	ShowDiagramAppClusterTypeCloudera bool `json:"showDiagramAppClusterTypeCloudera,omitempty"`
	ShowDiagramAppClusterTypeCouchbase bool `json:"showDiagramAppClusterTypeCouchbase,omitempty"`
	ShowDiagramAppClusterTypeDatameer bool `json:"showDiagramAppClusterTypeDatameer,omitempty"`
	ShowDiagramAppClusterTypeDatastax bool `json:"showDiagramAppClusterTypeDatastax,omitempty"`
	ShowDiagramAppClusterTypeElasticsearch bool `json:"showDiagramAppClusterTypeElasticsearch,omitempty"`
	ShowDiagramAppClusterTypeExasol bool `json:"showDiagramAppClusterTypeExasol,omitempty"`
	ShowDiagramAppClusterTypeHortonworks bool `json:"showDiagramAppClusterTypeHortonworks,omitempty"`
	ShowDiagramAppClusterTypeKubernetes bool `json:"showDiagramAppClusterTypeKubernetes,omitempty"`
	ShowDiagramAppClusterTypeMapr bool `json:"showDiagramAppClusterTypeMapr,omitempty"`
	ShowDiagramAppClusterTypeMesos bool `json:"showDiagramAppClusterTypeMesos,omitempty"`
	ShowDiagramAppClusterTypeMysqlPercona bool `json:"showDiagramAppClusterTypeMysqlPercona,omitempty"`
	ShowDiagramAppClusterTypeSplunk bool `json:"showDiagramAppClusterTypeSplunk,omitempty"`
	ShowDiagramAppClusterTypeVmwareVsphere bool `json:"showDiagramAppClusterTypeVmwareVsphere,omitempty"`
	ShowDiagramAppClusterTypeVmwareVcf bool `json:"showDiagramAppClusterTypeVmwareVcf,omitempty"`
	ShowDiagramAppClusterTypeKubernetesEksa bool `json:"showDiagramAppClusterTypeKubernetesEksa,omitempty"`
	ShowDiagramVmInstanceGroups bool `json:"showDiagramVmInstanceGroups,omitempty"`
	ShowEliChatBot bool `json:"showEliChatBot,omitempty"`
	EnableCustomRaidConfiguration bool `json:"enableCustomRaidConfiguration,omitempty"`
	EnableCustomSubnets bool `json:"enableCustomSubnets,omitempty"`
	ShowStackTrace bool `json:"showStackTrace,omitempty"`
	AllowedNetworkProfiles *interface{} `json:"allowedNetworkProfiles,omitempty"`
}
