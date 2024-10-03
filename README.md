# Go API client for sdk2

MetalSoft REST API documentation

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 6.4.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://www.metalsoft.io/contact](https://www.metalsoft.io/contact)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./sdk2"
```

## Documentation for API Endpoints

All URIs are relative to **

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AIApi* | [**EliControllerGenerateEliResponse**](docs/AIApi.md#elicontrollergenerateeliresponse) | **Post** /api/v2/ai/generate | Request from AI a response for the given input
*AccountsApi* | [**AccountsControllerArchiveAccount**](docs/AccountsApi.md#accountscontrollerarchiveaccount) | **Post** /api/v2/accounts/{accountId}/actions/archive | Archive account
*AccountsApi* | [**AccountsControllerCreateAccount**](docs/AccountsApi.md#accountscontrollercreateaccount) | **Post** /api/v2/accounts | Create account
*AccountsApi* | [**AccountsControllerGetAccount**](docs/AccountsApi.md#accountscontrollergetaccount) | **Get** /api/v2/accounts/{accountId} | Get account by id
*AccountsApi* | [**AccountsControllerGetAccountLimits**](docs/AccountsApi.md#accountscontrollergetaccountlimits) | **Get** /api/v2/accounts/{accountId}/limits | Get account limits
*AccountsApi* | [**AccountsControllerGetAccounts**](docs/AccountsApi.md#accountscontrollergetaccounts) | **Get** /api/v2/accounts | Get all accounts
*AccountsApi* | [**AccountsControllerGetUsersForAccount**](docs/AccountsApi.md#accountscontrollergetusersforaccount) | **Get** /api/v2/accounts/{accountId}/users | Get users for account
*AccountsApi* | [**AccountsControllerUnarchiveAccount**](docs/AccountsApi.md#accountscontrollerunarchiveaccount) | **Post** /api/v2/accounts/{accountId}/actions/unarchive | Unarchive account
*AccountsApi* | [**AccountsControllerUpdateAccount**](docs/AccountsApi.md#accountscontrollerupdateaccount) | **Patch** /api/v2/accounts/{accountId} | Update account
*AccountsApi* | [**AccountsControllerUpdateAccountLimits**](docs/AccountsApi.md#accountscontrollerupdateaccountlimits) | **Patch** /api/v2/accounts/{accountId}/limits | Update account limits
*ConfigurationApi* | [**ConfigController1GetConfiguration**](docs/ConfigurationApi.md#configcontroller1getconfiguration) | **Get** /api/v2/config | Get configuration
*FileShareApi* | [**BlueprintControllerCreateFileShare**](docs/FileShareApi.md#blueprintcontrollercreatefileshare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares | Creates a File Share
*FileShareApi* | [**BlueprintControllerDeleteFileShare**](docs/FileShareApi.md#blueprintcontrollerdeletefileshare) | **Delete** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Deletes a File Share
*FileShareApi* | [**BlueprintControllerGetFileShare**](docs/FileShareApi.md#blueprintcontrollergetfileshare) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Get File Share information
*FileShareApi* | [**BlueprintControllerGetFileShareHosts**](docs/FileShareApi.md#blueprintcontrollergetfilesharehosts) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/hosts | Get the Hosts of File Share
*FileShareApi* | [**BlueprintControllerGetFileShares**](docs/FileShareApi.md#blueprintcontrollergetfileshares) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares | Get all File Shares
*FileShareApi* | [**BlueprintControllerStartFileShare**](docs/FileShareApi.md#blueprintcontrollerstartfileshare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/start | Starts a File Share
*FileShareApi* | [**BlueprintControllerStopFileShare**](docs/FileShareApi.md#blueprintcontrollerstopfileshare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/stop | Stops a File Share
*FileShareApi* | [**BlueprintControllerUpdateFileShare**](docs/FileShareApi.md#blueprintcontrollerupdatefileshare) | **Patch** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Updates File Share information
*FileShareApi* | [**BlueprintControllerUpdateFileShareInstanceArrayHostsBulk**](docs/FileShareApi.md#blueprintcontrollerupdatefileshareinstancearrayhostsbulk) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/modify-instance-array-hosts-bulk | Updates Instance Array Hosts on the File Share
*InfrastructureApi* | [**InventoryController1DeployInfrastructure**](docs/InfrastructureApi.md#inventorycontroller1deployinfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/deploy | Deploys the specified infrastructure
*InfrastructureApi* | [**InventoryController1GetInfrastructure**](docs/InfrastructureApi.md#inventorycontroller1getinfrastructure) | **Get** /api/v2/infrastructures/{infrastructureId} | Retrieves the specified infrastructure
*InfrastructureApi* | [**InventoryController1GetInfrastructures**](docs/InfrastructureApi.md#inventorycontroller1getinfrastructures) | **Get** /api/v2/infrastructures | Get all infrastructures
*InfrastructureApi* | [**InventoryController1RevertInfrastructure**](docs/InfrastructureApi.md#inventorycontroller1revertinfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/revert | Reverts the specified infrastructure
*NetworkApi* | [**BlueprintControllerCreateNetwork**](docs/NetworkApi.md#blueprintcontrollercreatenetwork) | **Post** /api/v2/infrastructures/{infrastructureId}/networks | Creates a new LAN network on the infrastructure
*NetworkApi* | [**BlueprintControllerDeleteNetwork**](docs/NetworkApi.md#blueprintcontrollerdeletenetwork) | **Delete** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Deletes a network from the infrastructure
*NetworkApi* | [**BlueprintControllerGetNetwork**](docs/NetworkApi.md#blueprintcontrollergetnetwork) | **Get** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Gets the specified network from the infrastructure
*NetworkApi* | [**BlueprintControllerGetNetworks**](docs/NetworkApi.md#blueprintcontrollergetnetworks) | **Get** /api/v2/infrastructures/{infrastructureId}/networks | Retrieves all networks on the infrastructure
*NetworkDevicesApi* | [**SwitchControllerChangeStatus**](docs/NetworkDevicesApi.md#switchcontrollerchangestatus) | **Patch** /api/v2/network-devices/{networkDeviceId}/actions/change-status | Change status of a network device
*NetworkDevicesApi* | [**SwitchControllerDiscover**](docs/NetworkDevicesApi.md#switchcontrollerdiscover) | **Post** /api/v2/network-devices/{networkDeviceId}/discover | Discover network device interfaces, hardware and software configuration
*NetworkDevicesApi* | [**SwitchControllerEnableSyslog**](docs/NetworkDevicesApi.md#switchcontrollerenablesyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-subscribe | Enables remote syslog for a network device
*NetworkDevicesApi* | [**SwitchControllerGetPorts**](docs/NetworkDevicesApi.md#switchcontrollergetports) | **Get** /api/v2/network-devices/{networkDeviceId}/ports | Get all ports for network device
*NetworkDevicesApi* | [**SwitchControllerResetNetworkDevice**](docs/NetworkDevicesApi.md#switchcontrollerresetnetworkdevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/reset | Resets a network device to default state
*NetworkDevicesApi* | [**SwitchControllerSetPortStatus**](docs/NetworkDevicesApi.md#switchcontrollersetportstatus) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-port-status | Set port status
*ResourcePoolsApi* | [**InventoryController1AddResourcePoolUser**](docs/ResourcePoolsApi.md#inventorycontroller1addresourcepooluser) | **Post** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Add a user to a Resource Pool
*ResourcePoolsApi* | [**InventoryController1AddServerToResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1addservertoresourcepool) | **Put** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Add a server to a Resource Pool
*ResourcePoolsApi* | [**InventoryController1AddSubnetPoolToResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1addsubnetpooltoresourcepool) | **Put** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Add a subnet pool to a resource pool
*ResourcePoolsApi* | [**InventoryController1CreateResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1createresourcepool) | **Post** /api/v2/resource-pools | Creates a Resource Pool
*ResourcePoolsApi* | [**InventoryController1DeleteResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1deleteresourcepool) | **Delete** /api/v2/resource-pools/{resourcePoolId} | Deletes a Resource Pool
*ResourcePoolsApi* | [**InventoryController1GetResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1getresourcepool) | **Get** /api/v2/resource-pools/{resourcePoolId} | Get Resource Pool information
*ResourcePoolsApi* | [**InventoryController1GetResourcePoolServers**](docs/ResourcePoolsApi.md#inventorycontroller1getresourcepoolservers) | **Get** /api/v2/resource-pools/{resourcePoolId}/servers | Get all servers that are part of a Resource Pool
*ResourcePoolsApi* | [**InventoryController1GetResourcePoolSubnetPools**](docs/ResourcePoolsApi.md#inventorycontroller1getresourcepoolsubnetpools) | **Get** /api/v2/resource-pools/{resourcePoolId}/subnet-pools | Get all subnet pools that are part of a resource pool
*ResourcePoolsApi* | [**InventoryController1GetResourcePoolUsers**](docs/ResourcePoolsApi.md#inventorycontroller1getresourcepoolusers) | **Get** /api/v2/resource-pools/{resourcePoolId}/users | Get all users that have access to a Resource Pool
*ResourcePoolsApi* | [**InventoryController1GetResourcePools**](docs/ResourcePoolsApi.md#inventorycontroller1getresourcepools) | **Get** /api/v2/resource-pools | Get all Resource Pools
*ResourcePoolsApi* | [**InventoryController1GetUserResourcePools**](docs/ResourcePoolsApi.md#inventorycontroller1getuserresourcepools) | **Get** /api/v2/resource-pools/user/{userId} | Get all Resource Pools that a user has access to
*ResourcePoolsApi* | [**InventoryController1RemoveResourcePoolUser**](docs/ResourcePoolsApi.md#inventorycontroller1removeresourcepooluser) | **Delete** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Remove a user from a Resource Pool
*ResourcePoolsApi* | [**InventoryController1RemoveServerFromResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1removeserverfromresourcepool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Remove a server from a Resource Pool
*ResourcePoolsApi* | [**InventoryController1RemoveSubnetPoolFromResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1removesubnetpoolfromresourcepool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Remove a subnet from a resource pool
*ResourcePoolsApi* | [**InventoryController1UpdateResourcePool**](docs/ResourcePoolsApi.md#inventorycontroller1updateresourcepool) | **Put** /api/v2/resource-pools/{resourcePoolId} | Updates Resource Pool information
*SecurityApi* | [**AuthenticationControllerListProviders**](docs/SecurityApi.md#authenticationcontrollerlistproviders) | **Get** /api/v2/authentication/providers | Get available authentication providers
*SecurityApi* | [**AuthenticationControllerUpdateProvider**](docs/SecurityApi.md#authenticationcontrollerupdateprovider) | **Patch** /api/v2/authentication/providers/{name} | Updates authentication provider
*ServerApi* | [**InventoryController1GetServerInfo**](docs/ServerApi.md#inventorycontroller1getserverinfo) | **Get** /api/v2/servers/{serverId} | Get Server information
*ServerApi* | [**InventoryController1RegisterServer**](docs/ServerApi.md#inventorycontroller1registerserver) | **Post** /api/v2/servers | Initialize server registration
*ServerApi* | [**ServersController1EnableSyslog**](docs/ServerApi.md#serverscontroller1enablesyslog) | **Post** /api/v2/servers/{serverId}/actions/syslog-subscribe | Enables remote syslog for a server
*ServerApi* | [**ServersController1GetPowerState**](docs/ServerApi.md#serverscontroller1getpowerstate) | **Post** /api/v2/servers/{serverId}/actions/get-power | Gets the power state of a server
*ServerApi* | [**ServersController1GetRemoteConsoleInfo**](docs/ServerApi.md#serverscontroller1getremoteconsoleinfo) | **Get** /api/v2/servers/{serverId}/remote-console-info | Get Remote Console information
*ServerApi* | [**ServersController1GetVNCInfo**](docs/ServerApi.md#serverscontroller1getvncinfo) | **Get** /api/v2/servers/{serverId}/vnc-info | Get VNC information
*ServerApi* | [**ServersController1ResetServerToFactoryDefaults**](docs/ServerApi.md#serverscontroller1resetservertofactorydefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
*ServerApi* | [**ServersController1SetPowerState**](docs/ServerApi.md#serverscontroller1setpowerstate) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server
*StorageApi* | [**InventoryController1CreateStorage**](docs/StorageApi.md#inventorycontroller1createstorage) | **Post** /api/v2/storages | Creates a Storage
*SystemApi* | [**SystemControllerGetVersion**](docs/SystemApi.md#systemcontrollergetversion) | **Get** /api/v2/version | Get MetalSoft system version
*UsersApi* | [**UsersControllerArchiveUser**](docs/UsersApi.md#userscontrollerarchiveuser) | **Post** /api/v2/users/{userId}/actions/archive | Archive user
*UsersApi* | [**UsersControllerChangeUserAccount**](docs/UsersApi.md#userscontrollerchangeuseraccount) | **Post** /api/v2/users/{userId}/actions/change-account | Change account for user
*UsersApi* | [**UsersControllerCreateUser**](docs/UsersApi.md#userscontrollercreateuser) | **Post** /api/v2/users | Creates a user
*UsersApi* | [**UsersControllerGetUser**](docs/UsersApi.md#userscontrollergetuser) | **Get** /api/v2/users/{userId} | Get user
*UsersApi* | [**UsersControllerGetUserLimits**](docs/UsersApi.md#userscontrollergetuserlimits) | **Get** /api/v2/users/{userId}/limits | Get user limits
*UsersApi* | [**UsersControllerGetUsers**](docs/UsersApi.md#userscontrollergetusers) | **Get** /api/v2/users | Get users
*UsersApi* | [**UsersControllerUnarchiveUser**](docs/UsersApi.md#userscontrollerunarchiveuser) | **Post** /api/v2/users/{userId}/actions/unarchive | Unarchive user
*UsersApi* | [**UsersControllerUpdateUser**](docs/UsersApi.md#userscontrollerupdateuser) | **Patch** /api/v2/users/{userId} | Update user
*UsersApi* | [**UsersControllerUpdateUserLimits**](docs/UsersApi.md#userscontrollerupdateuserlimits) | **Patch** /api/v2/users/{userId}/limits | Update user limits
*VMInstanceApi* | [**BlueprintControllerApplyVMTypeOnVMInstance**](docs/VMInstanceApi.md#blueprintcontrollerapplyvmtypeonvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance
*VMInstanceApi* | [**BlueprintControllerCreateVMInstance**](docs/VMInstanceApi.md#blueprintcontrollercreatevminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances | Creates a VM Instance
*VMInstanceApi* | [**BlueprintControllerDeleteVMInstance**](docs/VMInstanceApi.md#blueprintcontrollerdeletevminstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Deletes a VM Instance
*VMInstanceApi* | [**BlueprintControllerGetVMInstance**](docs/VMInstanceApi.md#blueprintcontrollergetvminstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Get VM Instance information
*VMInstanceApi* | [**BlueprintControllerPowerStatusVMInstance**](docs/VMInstanceApi.md#blueprintcontrollerpowerstatusvminstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/power-status | Retrieves the power status of the VM Instance
*VMInstanceApi* | [**BlueprintControllerRebootVMInstance**](docs/VMInstanceApi.md#blueprintcontrollerrebootvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/reboot | Reboots the VM Instance
*VMInstanceApi* | [**BlueprintControllerShutdownVMInstance**](docs/VMInstanceApi.md#blueprintcontrollershutdownvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/shutdown | Shuts down the VM Instance
*VMInstanceApi* | [**BlueprintControllerStartVMInstance**](docs/VMInstanceApi.md#blueprintcontrollerstartvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/start | Starts the VM Instance
*VMInstanceApi* | [**BlueprintControllerUpdateVMInstance**](docs/VMInstanceApi.md#blueprintcontrollerupdatevminstance) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Updates VM Instance information
*VMInstanceGroupApi* | [**BlueprintControllerApplyVMTypeOnVMInstanceGroup**](docs/VMInstanceGroupApi.md#blueprintcontrollerapplyvmtypeonvminstancegroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance Group
*VMInstanceGroupApi* | [**BlueprintControllerCreateVMInstanceGroup**](docs/VMInstanceGroupApi.md#blueprintcontrollercreatevminstancegroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Creates a VM Instance Group
*VMInstanceGroupApi* | [**BlueprintControllerCreateVMInterfaceOnVMInstanceGroup**](docs/VMInstanceGroupApi.md#blueprintcontrollercreatevminterfaceonvminstancegroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/interfaces | Creates a new Virtual Interface for the VM Instance Group
*VMInstanceGroupApi* | [**BlueprintControllerDeleteVMInstanceGroup**](docs/VMInstanceGroupApi.md#blueprintcontrollerdeletevminstancegroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Deletes a VM Instance Group
*VMInstanceGroupApi* | [**BlueprintControllerGetVMInstanceGroup**](docs/VMInstanceGroupApi.md#blueprintcontrollergetvminstancegroup) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Get VM Instance Group information
*VMInstanceGroupApi* | [**BlueprintControllerGetVMInstanceGroupVMInstances**](docs/VMInstanceGroupApi.md#blueprintcontrollergetvminstancegroupvminstances) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/vm-instances | Get the VM Instances of VM Instance Group
*VMInstanceGroupApi* | [**BlueprintControllerGetVMInstanceGroups**](docs/VMInstanceGroupApi.md#blueprintcontrollergetvminstancegroups) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Get all VM Instance Groups
*VMInstanceGroupApi* | [**BlueprintControllerPatchNetworkProfileOnVMInstanceGroupNetwork**](docs/VMInstanceGroupApi.md#blueprintcontrollerpatchnetworkprofileonvminstancegroupnetwork) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/network/{networkId} | Applies the given Network Profile to the specified VM Instance Group Network
*VMInstanceGroupApi* | [**BlueprintControllerUpdateVMInstanceGroup**](docs/VMInstanceGroupApi.md#blueprintcontrollerupdatevminstancegroup) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Updates VM Instance Group information
*VMPoolsApi* | [**InventoryController1CreateVMPool**](docs/VMPoolsApi.md#inventorycontroller1createvmpool) | **Post** /api/v2/vm-pools | Creates a VM Pool
*VMPoolsApi* | [**InventoryController1DeleteVMPool**](docs/VMPoolsApi.md#inventorycontroller1deletevmpool) | **Delete** /api/v2/vm-pools/{vmPoolId} | Deletes a VM Pool
*VMPoolsApi* | [**InventoryController1GetVMPool**](docs/VMPoolsApi.md#inventorycontroller1getvmpool) | **Get** /api/v2/vm-pools/{vmPoolId} | Get VM Pool information
*VMPoolsApi* | [**InventoryController1GetVMPoolClusterHost**](docs/VMPoolsApi.md#inventorycontroller1getvmpoolclusterhost) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId} | Retrieves a VM Cluster Host
*VMPoolsApi* | [**InventoryController1GetVMPoolClusterHostInterface**](docs/VMPoolsApi.md#inventorycontroller1getvmpoolclusterhostinterface) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Retrieves a VM Cluster Host Interface
*VMPoolsApi* | [**InventoryController1GetVMPoolClusterHostInterfaces**](docs/VMPoolsApi.md#inventorycontroller1getvmpoolclusterhostinterfaces) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces | Retrieves a list of VM Cluster Host Interfaces
*VMPoolsApi* | [**InventoryController1GetVMPoolClusterHostVMs**](docs/VMPoolsApi.md#inventorycontroller1getvmpoolclusterhostvms) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/vms | Retrieves a list of VM Cluster Host VMs
*VMPoolsApi* | [**InventoryController1GetVMPoolClusterHosts**](docs/VMPoolsApi.md#inventorycontroller1getvmpoolclusterhosts) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts | Get list of VM Cluster Hosts linked to the VM Pool
*VMPoolsApi* | [**InventoryController1GetVMPoolVMs**](docs/VMPoolsApi.md#inventorycontroller1getvmpoolvms) | **Get** /api/v2/vm-pools/{vmPoolId}/vms | Returns all VMs linked to the VM Pool
*VMPoolsApi* | [**InventoryController1GetVMPools**](docs/VMPoolsApi.md#inventorycontroller1getvmpools) | **Get** /api/v2/vm-pools | Get all VM Pools
*VMPoolsApi* | [**InventoryController1PatchVMPoolClusterHostInterface**](docs/VMPoolsApi.md#inventorycontroller1patchvmpoolclusterhostinterface) | **Patch** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Updates a VM Cluster Host Interface
*VMPoolsApi* | [**InventoryController1UpdateVMPool**](docs/VMPoolsApi.md#inventorycontroller1updatevmpool) | **Patch** /api/v2/vm-pools/{vmPoolId} | Updates VM Pool information
*VMTypesApi* | [**InventoryController1CreateVMType**](docs/VMTypesApi.md#inventorycontroller1createvmtype) | **Post** /api/v2/vm-types | Creates a VM Type
*VMTypesApi* | [**InventoryController1DeleteVMType**](docs/VMTypesApi.md#inventorycontroller1deletevmtype) | **Delete** /api/v2/vm-types/{vmTypeId} | Deletes a VM Type
*VMTypesApi* | [**InventoryController1GetVMType**](docs/VMTypesApi.md#inventorycontroller1getvmtype) | **Get** /api/v2/vm-types/{vmTypeId} | Get VM Type information
*VMTypesApi* | [**InventoryController1GetVMTypes**](docs/VMTypesApi.md#inventorycontroller1getvmtypes) | **Get** /api/v2/vm-types | Get all VM Types
*VMTypesApi* | [**InventoryController1GetVMsByVMType**](docs/VMTypesApi.md#inventorycontroller1getvmsbyvmtype) | **Get** /api/v2/vm-types/{vmTypeId}/vms | Returns all VMs linked to the VM Type
*VMTypesApi* | [**InventoryController1UpdateVMType**](docs/VMTypesApi.md#inventorycontroller1updatevmtype) | **Patch** /api/v2/vm-types/{vmTypeId} | Updates VM Type information
*VMsApi* | [**InventoryController1UpdateVM**](docs/VMsApi.md#inventorycontroller1updatevm) | **Patch** /api/v2/vms/{vmId} | Updates VM information
*VMsApi* | [**VMMicroserviceControllerGetRemoteConsoleInfo**](docs/VMsApi.md#vmmicroservicecontrollergetremoteconsoleinfo) | **Get** /api/v2/vms/{vmId}/remote-console-info | Get Remote Console information
*VMsApi* | [**VMMicroserviceControllerGetVM**](docs/VMsApi.md#vmmicroservicecontrollergetvm) | **Get** /api/v2/vms/{vmId} | Retrieves the VM information
*VMsApi* | [**VMMicroserviceControllerPowerStatusVMInstance**](docs/VMsApi.md#vmmicroservicecontrollerpowerstatusvminstance) | **Get** /api/v2/vms/{vmId}/power-status | Retrieves the power status of the VM
*VMsApi* | [**VMMicroserviceControllerRebootVM**](docs/VMsApi.md#vmmicroservicecontrollerrebootvm) | **Post** /api/v2/vms/{vmId}/reboot | Reboots the VM
*VMsApi* | [**VMMicroserviceControllerShutdownVM**](docs/VMsApi.md#vmmicroservicecontrollershutdownvm) | **Post** /api/v2/vms/{vmId}/shutdown | Shuts down the VM
*VMsApi* | [**VMMicroserviceControllerStartVM**](docs/VMsApi.md#vmmicroservicecontrollerstartvm) | **Post** /api/v2/vms/{vmId}/start | Starts the VM

## Documentation For Models

 - [AccountAddressDto](docs/AccountAddressDto.md)
 - [AccountDto](docs/AccountDto.md)
 - [AccountLimits](docs/AccountLimits.md)
 - [AccountLimitsDto](docs/AccountLimitsDto.md)
 - [AfcJobInfo](docs/AfcJobInfo.md)
 - [AiGenerateRequestDto](docs/AiGenerateRequestDto.md)
 - [AiGenerateResponseDto](docs/AiGenerateResponseDto.md)
 - [AllOfAccountDtoAddress](docs/AllOfAccountDtoAddress.md)
 - [AllOfAccountDtoLimits](docs/AllOfAccountDtoLimits.md)
 - [AllOfCreateAccountDtoAddress](docs/AllOfCreateAccountDtoAddress.md)
 - [AllOfFileShareGuiSettings](docs/AllOfFileShareGuiSettings.md)
 - [AllOfInfrastructureDeployOptionsShutdownOptions](docs/AllOfInfrastructureDeployOptionsShutdownOptions.md)
 - [AllOfResourcePoolWithStatsDtoStatistics](docs/AllOfResourcePoolWithStatsDtoStatistics.md)
 - [AllOfServerRegistrationResponseDtoJobInfo](docs/AllOfServerRegistrationResponseDtoJobInfo.md)
 - [AllOfUpdateAccountDtoAddress](docs/AllOfUpdateAccountDtoAddress.md)
 - [AllOfUpdateFileShareGuiSettings](docs/AllOfUpdateFileShareGuiSettings.md)
 - [AllOfUpdateUserDtoPassword](docs/AllOfUpdateUserDtoPassword.md)
 - [AllOfUpdateUserDtoPermissions](docs/AllOfUpdateUserDtoPermissions.md)
 - [AllOfUpdateUserDtoSuspension](docs/AllOfUpdateUserDtoSuspension.md)
 - [AllOfUpdateVmInstanceGroupDtoGuiSettings](docs/AllOfUpdateVmInstanceGroupDtoGuiSettings.md)
 - [AllOfUserDtoLimits](docs/AllOfUserDtoLimits.md)
 - [AllOfVmInstanceGroupDtoGuiSettings](docs/AllOfVmInstanceGroupDtoGuiSettings.md)
 - [AuthenticationProvider](docs/AuthenticationProvider.md)
 - [AuthenticationProviderUpdate](docs/AuthenticationProviderUpdate.md)
 - [ChangeUserAccountDto](docs/ChangeUserAccountDto.md)
 - [CreateAccountDto](docs/CreateAccountDto.md)
 - [CreateFileShare](docs/CreateFileShare.md)
 - [CreateNetworkDto](docs/CreateNetworkDto.md)
 - [CreateResourcePoolDto](docs/CreateResourcePoolDto.md)
 - [CreateStorage](docs/CreateStorage.md)
 - [CreateUserDto](docs/CreateUserDto.md)
 - [CreateVmInstanceDto](docs/CreateVmInstanceDto.md)
 - [CreateVmInstanceGroupDto](docs/CreateVmInstanceGroupDto.md)
 - [CreateVmInstanceGroupInterfaceDto](docs/CreateVmInstanceGroupInterfaceDto.md)
 - [CreateVmPoolDto](docs/CreateVmPoolDto.md)
 - [CreateVmTypeDto](docs/CreateVmTypeDto.md)
 - [DiscoveryQueryDto](docs/DiscoveryQueryDto.md)
 - [FileShare](docs/FileShare.md)
 - [FileShareHostBulkOperation](docs/FileShareHostBulkOperation.md)
 - [FileShareHosts](docs/FileShareHosts.md)
 - [FileShareHostsModifyBulk](docs/FileShareHostsModifyBulk.md)
 - [GenericGuiSettings](docs/GenericGuiSettings.md)
 - [InfrastructureDeployOptions](docs/InfrastructureDeployOptions.md)
 - [InfrastructureDeployShutdownOptions](docs/InfrastructureDeployShutdownOptions.md)
 - [InfrastructureDto](docs/InfrastructureDto.md)
 - [Link](docs/Link.md)
 - [NetworkDevicePortStatusDto](docs/NetworkDevicePortStatusDto.md)
 - [NetworkDto](docs/NetworkDto.md)
 - [RemoteConsoleInfoDto](docs/RemoteConsoleInfoDto.md)
 - [ResourcePoolDto](docs/ResourcePoolDto.md)
 - [ResourcePoolStatistics](docs/ResourcePoolStatistics.md)
 - [ResourcePoolWithStatsDto](docs/ResourcePoolWithStatsDto.md)
 - [Server](docs/Server.md)
 - [ServerPowerSetDto](docs/ServerPowerSetDto.md)
 - [ServerRegistrationDto](docs/ServerRegistrationDto.md)
 - [ServerRegistrationResponseDto](docs/ServerRegistrationResponseDto.md)
 - [ServerVncInfoDto](docs/ServerVncInfoDto.md)
 - [StorageRegistrationResponse](docs/StorageRegistrationResponse.md)
 - [UpdateAccountDto](docs/UpdateAccountDto.md)
 - [UpdateFileShare](docs/UpdateFileShare.md)
 - [UpdateResourcePoolDto](docs/UpdateResourcePoolDto.md)
 - [UpdateUserDto](docs/UpdateUserDto.md)
 - [UpdateVmDto](docs/UpdateVmDto.md)
 - [UpdateVmInstanceDto](docs/UpdateVmInstanceDto.md)
 - [UpdateVmInstanceGroupDto](docs/UpdateVmInstanceGroupDto.md)
 - [UpdateVmInstanceGroupInterfaceDto](docs/UpdateVmInstanceGroupInterfaceDto.md)
 - [UpdateVmInstanceGroupNetworkDto](docs/UpdateVmInstanceGroupNetworkDto.md)
 - [UpdateVmPoolClusterHostInterfaceDto](docs/UpdateVmPoolClusterHostInterfaceDto.md)
 - [UpdateVmPoolDto](docs/UpdateVmPoolDto.md)
 - [UpdateVmTypeDto](docs/UpdateVmTypeDto.md)
 - [UserDelegateDto](docs/UserDelegateDto.md)
 - [UserDto](docs/UserDto.md)
 - [UserExperimentalTagDto](docs/UserExperimentalTagDto.md)
 - [UserLimits](docs/UserLimits.md)
 - [UserLimitsDto](docs/UserLimitsDto.md)
 - [UserPromotionDto](docs/UserPromotionDto.md)
 - [UserResourcePermissionDto](docs/UserResourcePermissionDto.md)
 - [UserResourcePermissionsDto](docs/UserResourcePermissionsDto.md)
 - [UserSuspendDto](docs/UserSuspendDto.md)
 - [UserUpdatePasswordDto](docs/UserUpdatePasswordDto.md)
 - [Version](docs/Version.md)
 - [VmDiskDto](docs/VmDiskDto.md)
 - [VmDto](docs/VmDto.md)
 - [VmInstanceDto](docs/VmInstanceDto.md)
 - [VmInstanceGroupDto](docs/VmInstanceGroupDto.md)
 - [VmInstanceGroupGuiSettingsDto](docs/VmInstanceGroupGuiSettingsDto.md)
 - [VmInstanceGroupInterfaceDto](docs/VmInstanceGroupInterfaceDto.md)
 - [VmPoolDto](docs/VmPoolDto.md)
 - [VmPoolHostInterfacesDto](docs/VmPoolHostInterfacesDto.md)
 - [VmPoolHostsDto](docs/VmPoolHostsDto.md)
 - [VmTypeDto](docs/VmTypeDto.md)

## Documentation For Authorization

## JWT
## apiKey

## Author

support@metalsoft.io
