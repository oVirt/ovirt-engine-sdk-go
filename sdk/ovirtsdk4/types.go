//
// Copyright (c) 2017 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package ovirtsdk4

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Actions struct {
	XMLName xml.Name `xml:"actions"`
	Actions []Action `xml:"action,omitempty"`
}

type Action struct {
	OvStruct
	AllowPartialImport             *bool                                 `xml:"allow_partial_import,omitempty"`
	Async                          *bool                                 `xml:"async,omitempty"`
	Bricks                         []GlusterBrick                        `xml:"bricks,omitempty"`
	Certificates                   []Certificate                         `xml:"certificates,omitempty"`
	CheckConnectivity              *bool                                 `xml:"check_connectivity,omitempty"`
	Clone                          *bool                                 `xml:"clone,omitempty"`
	Cluster                        *Cluster                              `xml:"cluster,omitempty"`
	CollapseSnapshots              *bool                                 `xml:"collapse_snapshots,omitempty"`
	Comment                        *string                               `xml:"comment,omitempty"`
	ConnectivityTimeout            *int64                                `xml:"connectivity_timeout,omitempty"`
	DataCenter                     *DataCenter                           `xml:"data_center,omitempty"`
	DeployHostedEngine             *bool                                 `xml:"deploy_hosted_engine,omitempty"`
	Description                    *string                               `xml:"description,omitempty"`
	Details                        *GlusterVolumeProfileDetails          `xml:"details,omitempty"`
	DiscardSnapshots               *bool                                 `xml:"discard_snapshots,omitempty"`
	Disk                           *Disk                                 `xml:"disk,omitempty"`
	Disks                          []Disk                                `xml:"disks,omitempty"`
	Exclusive                      *bool                                 `xml:"exclusive,omitempty"`
	Fault                          *Fault                                `xml:"fault,omitempty"`
	FenceType                      *string                               `xml:"fence_type,omitempty"`
	Filter                         *bool                                 `xml:"filter,omitempty"`
	FixLayout                      *bool                                 `xml:"fix_layout,omitempty"`
	Force                          *bool                                 `xml:"force,omitempty"`
	GracePeriod                    *GracePeriod                          `xml:"grace_period,omitempty"`
	Host                           *Host                                 `xml:"host,omitempty"`
	Id                             *string                               `xml:"id,attr,omitempty"`
	Image                          *string                               `xml:"image,omitempty"`
	ImportAsTemplate               *bool                                 `xml:"import_as_template,omitempty"`
	IsAttached                     *bool                                 `xml:"is_attached,omitempty"`
	Iscsi                          *IscsiDetails                         `xml:"iscsi,omitempty"`
	IscsiTargets                   []string                              `xml:"iscsi_targets,omitempty"`
	Job                            *Job                                  `xml:"job,omitempty"`
	LogicalUnits                   []LogicalUnit                         `xml:"logical_units,omitempty"`
	MaintenanceEnabled             *bool                                 `xml:"maintenance_enabled,omitempty"`
	ModifiedBonds                  []HostNic                             `xml:"modified_bonds,omitempty"`
	ModifiedLabels                 []NetworkLabel                        `xml:"modified_labels,omitempty"`
	ModifiedNetworkAttachments     []NetworkAttachment                   `xml:"modified_network_attachments,omitempty"`
	Name                           *string                               `xml:"name,omitempty"`
	Option                         *Option                               `xml:"option,omitempty"`
	Pause                          *bool                                 `xml:"pause,omitempty"`
	PowerManagement                *PowerManagement                      `xml:"power_management,omitempty"`
	ProxyTicket                    *ProxyTicket                          `xml:"proxy_ticket,omitempty"`
	Reason                         *string                               `xml:"reason,omitempty"`
	ReassignBadMacs                *bool                                 `xml:"reassign_bad_macs,omitempty"`
	RemoteViewerConnectionFile     *string                               `xml:"remote_viewer_connection_file,omitempty"`
	RemovedBonds                   []HostNic                             `xml:"removed_bonds,omitempty"`
	RemovedLabels                  []NetworkLabel                        `xml:"removed_labels,omitempty"`
	RemovedNetworkAttachments      []NetworkAttachment                   `xml:"removed_network_attachments,omitempty"`
	ResolutionType                 *string                               `xml:"resolution_type,omitempty"`
	RestoreMemory                  *bool                                 `xml:"restore_memory,omitempty"`
	RootPassword                   *string                               `xml:"root_password,omitempty"`
	Snapshot                       *Snapshot                             `xml:"snapshot,omitempty"`
	Ssh                            *Ssh                                  `xml:"ssh,omitempty"`
	Status                         *string                               `xml:"status,omitempty"`
	StopGlusterService             *bool                                 `xml:"stop_gluster_service,omitempty"`
	StorageDomain                  *StorageDomain                        `xml:"storage_domain,omitempty"`
	StorageDomains                 []StorageDomain                       `xml:"storage_domains,omitempty"`
	Succeeded                      *bool                                 `xml:"succeeded,omitempty"`
	SynchronizedNetworkAttachments []NetworkAttachment                   `xml:"synchronized_network_attachments,omitempty"`
	Template                       *Template                             `xml:"template,omitempty"`
	Ticket                         *Ticket                               `xml:"ticket,omitempty"`
	UndeployHostedEngine           *bool                                 `xml:"undeploy_hosted_engine,omitempty"`
	UseCloudInit                   *bool                                 `xml:"use_cloud_init,omitempty"`
	UseSysprep                     *bool                                 `xml:"use_sysprep,omitempty"`
	VirtualFunctionsConfiguration  *HostNicVirtualFunctionsConfiguration `xml:"virtual_functions_configuration,omitempty"`
	Vm                             *Vm                                   `xml:"vm,omitempty"`
	VnicProfileMappings            []VnicProfileMapping                  `xml:"vnic_profile_mappings,omitempty"`
}

type AffinityGroups struct {
	XMLName        xml.Name        `xml:"affinitygroups"`
	AffinityGroups []AffinityGroup `xml:"affinity_group,omitempty"`
}

type AffinityGroup struct {
	OvStruct
	Cluster     *Cluster      `xml:"cluster,omitempty"`
	Comment     *string       `xml:"comment,omitempty"`
	Description *string       `xml:"description,omitempty"`
	Enforcing   *bool         `xml:"enforcing,omitempty"`
	Hosts       []Host        `xml:"hosts,omitempty"`
	HostsRule   *AffinityRule `xml:"hosts_rule,omitempty"`
	Id          *string       `xml:"id,attr,omitempty"`
	Name        *string       `xml:"name,omitempty"`
	Positive    *bool         `xml:"positive,omitempty"`
	Vms         []Vm          `xml:"vms,omitempty"`
	VmsRule     *AffinityRule `xml:"vms_rule,omitempty"`
}

type AffinityLabels struct {
	XMLName        xml.Name        `xml:"affinitylabels"`
	AffinityLabels []AffinityLabel `xml:"affinity_label,omitempty"`
}

type AffinityLabel struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Hosts       []Host  `xml:"hosts,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	ReadOnly    *bool   `xml:"read_only,omitempty"`
	Vms         []Vm    `xml:"vms,omitempty"`
}

type AffinityRules struct {
	XMLName       xml.Name       `xml:"affinityrules"`
	AffinityRules []AffinityRule `xml:"affinity_rule,omitempty"`
}

type AffinityRule struct {
	OvStruct
	Enabled   *bool `xml:"enabled,omitempty"`
	Enforcing *bool `xml:"enforcing,omitempty"`
	Positive  *bool `xml:"positive,omitempty"`
}

type Agents struct {
	XMLName xml.Name `xml:"agents"`
	Agents  []Agent  `xml:"agent,omitempty"`
}

type Agent struct {
	OvStruct
	Address        *string  `xml:"address,omitempty"`
	Comment        *string  `xml:"comment,omitempty"`
	Concurrent     *bool    `xml:"concurrent,omitempty"`
	Description    *string  `xml:"description,omitempty"`
	EncryptOptions *bool    `xml:"encrypt_options,omitempty"`
	Host           *Host    `xml:"host,omitempty"`
	Id             *string  `xml:"id,attr,omitempty"`
	Name           *string  `xml:"name,omitempty"`
	Options        []Option `xml:"options,omitempty"`
	Order          *int64   `xml:"order,omitempty"`
	Password       *string  `xml:"password,omitempty"`
	Port           *int64   `xml:"port,omitempty"`
	Type           *string  `xml:"type,omitempty"`
	Username       *string  `xml:"username,omitempty"`
}

type AgentConfigurations struct {
	XMLName             xml.Name             `xml:"agentconfigurations"`
	AgentConfigurations []AgentConfiguration `xml:"agent_configuration,omitempty"`
}

type AgentConfiguration struct {
	OvStruct
	Address         *string           `xml:"address,omitempty"`
	BrokerType      MessageBrokerType `xml:"broker_type,omitempty"`
	NetworkMappings *string           `xml:"network_mappings,omitempty"`
	Password        *string           `xml:"password,omitempty"`
	Port            *int64            `xml:"port,omitempty"`
	Username        *string           `xml:"username,omitempty"`
}

type Apis struct {
	XMLName xml.Name `xml:"apis"`
	Apis    []Api    `xml:"api,omitempty"`
}

type Api struct {
	OvStruct
	ProductInfo    *ProductInfo    `xml:"product_info,omitempty"`
	SpecialObjects *SpecialObjects `xml:"special_objects,omitempty"`
	Summary        *ApiSummary     `xml:"summary,omitempty"`
	Time           time.Time       `xml:"time,omitempty"`
}

type ApiSummarys struct {
	XMLName     xml.Name     `xml:"apisummarys"`
	ApiSummarys []ApiSummary `xml:"api_summary,omitempty"`
}

type ApiSummary struct {
	OvStruct
	Hosts          *ApiSummaryItem `xml:"hosts,omitempty"`
	StorageDomains *ApiSummaryItem `xml:"storage_domains,omitempty"`
	Users          *ApiSummaryItem `xml:"users,omitempty"`
	Vms            *ApiSummaryItem `xml:"vms,omitempty"`
}

type ApiSummaryItems struct {
	XMLName         xml.Name         `xml:"apisummaryitems"`
	ApiSummaryItems []ApiSummaryItem `xml:"api_summary_item,omitempty"`
}

type ApiSummaryItem struct {
	OvStruct
	Active *int64 `xml:"active,omitempty"`
	Total  *int64 `xml:"total,omitempty"`
}

type Applications struct {
	XMLName      xml.Name      `xml:"applications"`
	Applications []Application `xml:"application,omitempty"`
}

type Application struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Vm          *Vm     `xml:"vm,omitempty"`
}

type AuthorizedKeys struct {
	XMLName        xml.Name        `xml:"authorizedkeys"`
	AuthorizedKeys []AuthorizedKey `xml:"authorized_key,omitempty"`
}

type AuthorizedKey struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Key         *string `xml:"key,omitempty"`
	Name        *string `xml:"name,omitempty"`
	User        *User   `xml:"user,omitempty"`
}

type Balances struct {
	XMLName  xml.Name  `xml:"balances"`
	Balances []Balance `xml:"balance,omitempty"`
}

type Balance struct {
	OvStruct
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
	SchedulingPolicy     *SchedulingPolicy     `xml:"scheduling_policy,omitempty"`
	SchedulingPolicyUnit *SchedulingPolicyUnit `xml:"scheduling_policy_unit,omitempty"`
}

type Bioss struct {
	XMLName xml.Name `xml:"bioss"`
	Bioss   []Bios   `xml:"bios,omitempty"`
}

type Bios struct {
	OvStruct
	BootMenu *BootMenu `xml:"boot_menu,omitempty"`
}

type BlockStatistics struct {
	XMLName         xml.Name         `xml:"blockstatistics"`
	BlockStatistics []BlockStatistic `xml:"block_statistic,omitempty"`
}

type BlockStatistic struct {
	OvStruct
	Statistics []Statistic `xml:"statistics,omitempty"`
}

type Bondings struct {
	XMLName  xml.Name  `xml:"bondings"`
	Bondings []Bonding `xml:"bonding,omitempty"`
}

type Bonding struct {
	OvStruct
	ActiveSlave  *HostNic  `xml:"active_slave,omitempty"`
	AdPartnerMac *Mac      `xml:"ad_partner_mac,omitempty"`
	Options      []Option  `xml:"options,omitempty"`
	Slaves       []HostNic `xml:"slaves,omitempty"`
}

type Bookmarks struct {
	XMLName   xml.Name   `xml:"bookmarks"`
	Bookmarks []Bookmark `xml:"bookmark,omitempty"`
}

type Bookmark struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Value       *string `xml:"value,omitempty"`
}

type Boots struct {
	XMLName xml.Name `xml:"boots"`
	Boots   []Boot   `xml:"boot,omitempty"`
}

type Boot struct {
	OvStruct
	Devices []BootDevice `xml:"devices,omitempty"`
}

type BootMenus struct {
	XMLName   xml.Name   `xml:"bootmenus"`
	BootMenus []BootMenu `xml:"boot_menu,omitempty"`
}

type BootMenu struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
}

type BrickProfileDetails struct {
	XMLName             xml.Name             `xml:"brickprofiledetails"`
	BrickProfileDetails []BrickProfileDetail `xml:"brick_profile_detail,omitempty"`
}

type BrickProfileDetail struct {
	OvStruct
	Brick          *GlusterBrick   `xml:"brick,omitempty"`
	ProfileDetails []ProfileDetail `xml:"profile_details,omitempty"`
}

type Cdroms struct {
	XMLName xml.Name `xml:"cdroms"`
	Cdroms  []Cdrom  `xml:"cdrom,omitempty"`
}

type Cdrom struct {
	OvStruct
	Comment      *string       `xml:"comment,omitempty"`
	Description  *string       `xml:"description,omitempty"`
	File         *File         `xml:"file,omitempty"`
	Id           *string       `xml:"id,attr,omitempty"`
	InstanceType *InstanceType `xml:"instance_type,omitempty"`
	Name         *string       `xml:"name,omitempty"`
	Template     *Template     `xml:"template,omitempty"`
	Vm           *Vm           `xml:"vm,omitempty"`
	Vms          []Vm          `xml:"vms,omitempty"`
}

type Certificates struct {
	XMLName      xml.Name      `xml:"certificates"`
	Certificates []Certificate `xml:"certificate,omitempty"`
}

type Certificate struct {
	OvStruct
	Comment      *string `xml:"comment,omitempty"`
	Content      *string `xml:"content,omitempty"`
	Description  *string `xml:"description,omitempty"`
	Id           *string `xml:"id,attr,omitempty"`
	Name         *string `xml:"name,omitempty"`
	Organization *string `xml:"organization,omitempty"`
	Subject      *string `xml:"subject,omitempty"`
}

type CloudInits struct {
	XMLName    xml.Name    `xml:"cloudinits"`
	CloudInits []CloudInit `xml:"cloud_init,omitempty"`
}

type CloudInit struct {
	OvStruct
	AuthorizedKeys       []AuthorizedKey       `xml:"authorized_keys,omitempty"`
	Files                []File                `xml:"files,omitempty"`
	Host                 *Host                 `xml:"host,omitempty"`
	NetworkConfiguration *NetworkConfiguration `xml:"network_configuration,omitempty"`
	RegenerateSshKeys    *bool                 `xml:"regenerate_ssh_keys,omitempty"`
	Timezone             *string               `xml:"timezone,omitempty"`
	Users                []User                `xml:"users,omitempty"`
}

type Clusters struct {
	XMLName  xml.Name  `xml:"clusters"`
	Clusters []Cluster `xml:"cluster,omitempty"`
}

type Cluster struct {
	OvStruct
	AffinityGroups                   []AffinityGroup   `xml:"affinity_groups,omitempty"`
	BallooningEnabled                *bool             `xml:"ballooning_enabled,omitempty"`
	Comment                          *string           `xml:"comment,omitempty"`
	Cpu                              *Cpu              `xml:"cpu,omitempty"`
	CpuProfiles                      []CpuProfile      `xml:"cpu_profiles,omitempty"`
	CustomSchedulingPolicyProperties []Property        `xml:"custom_scheduling_policy_properties,omitempty"`
	DataCenter                       *DataCenter       `xml:"data_center,omitempty"`
	Description                      *string           `xml:"description,omitempty"`
	Display                          *Display          `xml:"display,omitempty"`
	ErrorHandling                    *ErrorHandling    `xml:"error_handling,omitempty"`
	FencingPolicy                    *FencingPolicy    `xml:"fencing_policy,omitempty"`
	GlusterHooks                     []GlusterHook     `xml:"gluster_hooks,omitempty"`
	GlusterService                   *bool             `xml:"gluster_service,omitempty"`
	GlusterTunedProfile              *string           `xml:"gluster_tuned_profile,omitempty"`
	GlusterVolumes                   []GlusterVolume   `xml:"gluster_volumes,omitempty"`
	HaReservation                    *bool             `xml:"ha_reservation,omitempty"`
	Id                               *string           `xml:"id,attr,omitempty"`
	Ksm                              *Ksm              `xml:"ksm,omitempty"`
	MacPool                          *MacPool          `xml:"mac_pool,omitempty"`
	MaintenanceReasonRequired        *bool             `xml:"maintenance_reason_required,omitempty"`
	ManagementNetwork                *Network          `xml:"management_network,omitempty"`
	MemoryPolicy                     *MemoryPolicy     `xml:"memory_policy,omitempty"`
	Migration                        *MigrationOptions `xml:"migration,omitempty"`
	Name                             *string           `xml:"name,omitempty"`
	NetworkFilters                   []NetworkFilter   `xml:"network_filters,omitempty"`
	Networks                         []Network         `xml:"networks,omitempty"`
	OptionalReason                   *bool             `xml:"optional_reason,omitempty"`
	Permissions                      []Permission      `xml:"permissions,omitempty"`
	RequiredRngSources               []RngSource       `xml:"required_rng_sources,omitempty"`
	SchedulingPolicy                 *SchedulingPolicy `xml:"scheduling_policy,omitempty"`
	SerialNumber                     *SerialNumber     `xml:"serial_number,omitempty"`
	SupportedVersions                []Version         `xml:"supported_versions,omitempty"`
	SwitchType                       SwitchType        `xml:"switch_type,omitempty"`
	ThreadsAsCores                   *bool             `xml:"threads_as_cores,omitempty"`
	TrustedService                   *bool             `xml:"trusted_service,omitempty"`
	TunnelMigration                  *bool             `xml:"tunnel_migration,omitempty"`
	Version                          *Version          `xml:"version,omitempty"`
	VirtService                      *bool             `xml:"virt_service,omitempty"`
}

type ClusterLevels struct {
	XMLName       xml.Name       `xml:"clusterlevels"`
	ClusterLevels []ClusterLevel `xml:"cluster_level,omitempty"`
}

type ClusterLevel struct {
	OvStruct
	Comment     *string   `xml:"comment,omitempty"`
	CpuTypes    []CpuType `xml:"cpu_types,omitempty"`
	Description *string   `xml:"description,omitempty"`
	Id          *string   `xml:"id,attr,omitempty"`
	Name        *string   `xml:"name,omitempty"`
	Permits     []Permit  `xml:"permits,omitempty"`
}

type Configurations struct {
	XMLName        xml.Name        `xml:"configurations"`
	Configurations []Configuration `xml:"configuration,omitempty"`
}

type Configuration struct {
	OvStruct
	Data *string           `xml:"data,omitempty"`
	Type ConfigurationType `xml:"type,omitempty"`
}

type Consoles struct {
	XMLName  xml.Name  `xml:"consoles"`
	Consoles []Console `xml:"console,omitempty"`
}

type Console struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
}

type Cores struct {
	XMLName xml.Name `xml:"cores"`
	Cores   []Core   `xml:"core,omitempty"`
}

type Core struct {
	OvStruct
	Index  *int64 `xml:"index,omitempty"`
	Socket *int64 `xml:"socket,omitempty"`
}

type Cpus struct {
	XMLName xml.Name `xml:"cpus"`
	Cpus    []Cpu    `xml:"cpu,omitempty"`
}

type Cpu struct {
	OvStruct
	Architecture Architecture `xml:"architecture,omitempty"`
	Cores        []Core       `xml:"cores,omitempty"`
	CpuTune      *CpuTune     `xml:"cpu_tune,omitempty"`
	Level        *int64       `xml:"level,omitempty"`
	Mode         CpuMode      `xml:"mode,omitempty"`
	Name         *string      `xml:"name,omitempty"`
	Speed        *float64     `xml:"speed,omitempty"`
	Topology     *CpuTopology `xml:"topology,omitempty"`
	Type         *string      `xml:"type,omitempty"`
}

type CpuProfiles struct {
	XMLName     xml.Name     `xml:"cpuprofiles"`
	CpuProfiles []CpuProfile `xml:"cpu_profile,omitempty"`
}

type CpuProfile struct {
	OvStruct
	Cluster     *Cluster     `xml:"cluster,omitempty"`
	Comment     *string      `xml:"comment,omitempty"`
	Description *string      `xml:"description,omitempty"`
	Id          *string      `xml:"id,attr,omitempty"`
	Name        *string      `xml:"name,omitempty"`
	Permissions []Permission `xml:"permissions,omitempty"`
	Qos         *Qos         `xml:"qos,omitempty"`
}

type CpuTopologys struct {
	XMLName      xml.Name      `xml:"cputopologys"`
	CpuTopologys []CpuTopology `xml:"cpu_topology,omitempty"`
}

type CpuTopology struct {
	OvStruct
	Cores   *int64 `xml:"cores,omitempty"`
	Sockets *int64 `xml:"sockets,omitempty"`
	Threads *int64 `xml:"threads,omitempty"`
}

type CpuTunes struct {
	XMLName  xml.Name  `xml:"cputunes"`
	CpuTunes []CpuTune `xml:"cpu_tune,omitempty"`
}

type CpuTune struct {
	OvStruct
	VcpuPins []VcpuPin `xml:"vcpu_pins,omitempty"`
}

type CpuTypes struct {
	XMLName  xml.Name  `xml:"cputypes"`
	CpuTypes []CpuType `xml:"cpu_type,omitempty"`
}

type CpuType struct {
	OvStruct
	Architecture Architecture `xml:"architecture,omitempty"`
	Level        *int64       `xml:"level,omitempty"`
	Name         *string      `xml:"name,omitempty"`
}

type CustomPropertys struct {
	XMLName         xml.Name         `xml:"custompropertys"`
	CustomPropertys []CustomProperty `xml:"custom_property,omitempty"`
}

type CustomProperty struct {
	OvStruct
	Name   *string `xml:"name,omitempty"`
	Regexp *string `xml:"regexp,omitempty"`
	Value  *string `xml:"value,omitempty"`
}

type DataCenters struct {
	XMLName     xml.Name     `xml:"datacenters"`
	DataCenters []DataCenter `xml:"data_center,omitempty"`
}

type DataCenter struct {
	OvStruct
	Clusters          []Cluster        `xml:"clusters,omitempty"`
	Comment           *string          `xml:"comment,omitempty"`
	Description       *string          `xml:"description,omitempty"`
	Id                *string          `xml:"id,attr,omitempty"`
	IscsiBonds        []IscsiBond      `xml:"iscsi_bonds,omitempty"`
	Local             *bool            `xml:"local,omitempty"`
	MacPool           *MacPool         `xml:"mac_pool,omitempty"`
	Name              *string          `xml:"name,omitempty"`
	Networks          []Network        `xml:"networks,omitempty"`
	Permissions       []Permission     `xml:"permissions,omitempty"`
	Qoss              []Qos            `xml:"qoss,omitempty"`
	QuotaMode         QuotaModeType    `xml:"quota_mode,omitempty"`
	Quotas            []Quota          `xml:"quotas,omitempty"`
	Status            DataCenterStatus `xml:"status,omitempty"`
	StorageDomains    []StorageDomain  `xml:"storage_domains,omitempty"`
	StorageFormat     StorageFormat    `xml:"storage_format,omitempty"`
	SupportedVersions []Version        `xml:"supported_versions,omitempty"`
	Version           *Version         `xml:"version,omitempty"`
}

type Devices struct {
	XMLName xml.Name `xml:"devices"`
	Devices []Device `xml:"device,omitempty"`
}

type Device struct {
	OvStruct
	Comment      *string       `xml:"comment,omitempty"`
	Description  *string       `xml:"description,omitempty"`
	Id           *string       `xml:"id,attr,omitempty"`
	InstanceType *InstanceType `xml:"instance_type,omitempty"`
	Name         *string       `xml:"name,omitempty"`
	Template     *Template     `xml:"template,omitempty"`
	Vm           *Vm           `xml:"vm,omitempty"`
	Vms          []Vm          `xml:"vms,omitempty"`
}

type Disks struct {
	XMLName xml.Name `xml:"disks"`
	Disks   []Disk   `xml:"disk,omitempty"`
}

type Disk struct {
	OvStruct
	Active              *bool                `xml:"active,omitempty"`
	ActualSize          *int64               `xml:"actual_size,omitempty"`
	Alias               *string              `xml:"alias,omitempty"`
	Bootable            *bool                `xml:"bootable,omitempty"`
	Comment             *string              `xml:"comment,omitempty"`
	Description         *string              `xml:"description,omitempty"`
	DiskProfile         *DiskProfile         `xml:"disk_profile,omitempty"`
	Format              DiskFormat           `xml:"format,omitempty"`
	Id                  *string              `xml:"id,attr,omitempty"`
	ImageId             *string              `xml:"image_id,omitempty"`
	InitialSize         *int64               `xml:"initial_size,omitempty"`
	InstanceType        *InstanceType        `xml:"instance_type,omitempty"`
	Interface           DiskInterface        `xml:"interface,omitempty"`
	LogicalName         *string              `xml:"logical_name,omitempty"`
	LunStorage          *HostStorage         `xml:"lun_storage,omitempty"`
	Name                *string              `xml:"name,omitempty"`
	OpenstackVolumeType *OpenStackVolumeType `xml:"openstack_volume_type,omitempty"`
	Permissions         []Permission         `xml:"permissions,omitempty"`
	PropagateErrors     *bool                `xml:"propagate_errors,omitempty"`
	ProvisionedSize     *int64               `xml:"provisioned_size,omitempty"`
	QcowVersion         QcowVersion          `xml:"qcow_version,omitempty"`
	Quota               *Quota               `xml:"quota,omitempty"`
	ReadOnly            *bool                `xml:"read_only,omitempty"`
	Sgio                ScsiGenericIO        `xml:"sgio,omitempty"`
	Shareable           *bool                `xml:"shareable,omitempty"`
	Snapshot            *Snapshot            `xml:"snapshot,omitempty"`
	Sparse              *bool                `xml:"sparse,omitempty"`
	Statistics          []Statistic          `xml:"statistics,omitempty"`
	Status              DiskStatus           `xml:"status,omitempty"`
	StorageDomain       *StorageDomain       `xml:"storage_domain,omitempty"`
	StorageDomains      []StorageDomain      `xml:"storage_domains,omitempty"`
	StorageType         DiskStorageType      `xml:"storage_type,omitempty"`
	Template            *Template            `xml:"template,omitempty"`
	UsesScsiReservation *bool                `xml:"uses_scsi_reservation,omitempty"`
	Vm                  *Vm                  `xml:"vm,omitempty"`
	Vms                 []Vm                 `xml:"vms,omitempty"`
	WipeAfterDelete     *bool                `xml:"wipe_after_delete,omitempty"`
}

type DiskAttachments struct {
	XMLName         xml.Name         `xml:"diskattachments"`
	DiskAttachments []DiskAttachment `xml:"disk_attachment,omitempty"`
}

type DiskAttachment struct {
	OvStruct
	Active              *bool         `xml:"active,omitempty"`
	Bootable            *bool         `xml:"bootable,omitempty"`
	Comment             *string       `xml:"comment,omitempty"`
	Description         *string       `xml:"description,omitempty"`
	Disk                *Disk         `xml:"disk,omitempty"`
	Id                  *string       `xml:"id,attr,omitempty"`
	Interface           DiskInterface `xml:"interface,omitempty"`
	LogicalName         *string       `xml:"logical_name,omitempty"`
	Name                *string       `xml:"name,omitempty"`
	PassDiscard         *bool         `xml:"pass_discard,omitempty"`
	Template            *Template     `xml:"template,omitempty"`
	UsesScsiReservation *bool         `xml:"uses_scsi_reservation,omitempty"`
	Vm                  *Vm           `xml:"vm,omitempty"`
}

type DiskProfiles struct {
	XMLName      xml.Name      `xml:"diskprofiles"`
	DiskProfiles []DiskProfile `xml:"disk_profile,omitempty"`
}

type DiskProfile struct {
	OvStruct
	Comment       *string        `xml:"comment,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Permissions   []Permission   `xml:"permissions,omitempty"`
	Qos           *Qos           `xml:"qos,omitempty"`
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
}

type DiskSnapshots struct {
	XMLName       xml.Name       `xml:"disksnapshots"`
	DiskSnapshots []DiskSnapshot `xml:"disk_snapshot,omitempty"`
}

type DiskSnapshot struct {
	OvStruct
	Active              *bool                `xml:"active,omitempty"`
	ActualSize          *int64               `xml:"actual_size,omitempty"`
	Alias               *string              `xml:"alias,omitempty"`
	Bootable            *bool                `xml:"bootable,omitempty"`
	Comment             *string              `xml:"comment,omitempty"`
	Description         *string              `xml:"description,omitempty"`
	Disk                *Disk                `xml:"disk,omitempty"`
	DiskProfile         *DiskProfile         `xml:"disk_profile,omitempty"`
	Format              DiskFormat           `xml:"format,omitempty"`
	Id                  *string              `xml:"id,attr,omitempty"`
	ImageId             *string              `xml:"image_id,omitempty"`
	InitialSize         *int64               `xml:"initial_size,omitempty"`
	InstanceType        *InstanceType        `xml:"instance_type,omitempty"`
	Interface           DiskInterface        `xml:"interface,omitempty"`
	LogicalName         *string              `xml:"logical_name,omitempty"`
	LunStorage          *HostStorage         `xml:"lun_storage,omitempty"`
	Name                *string              `xml:"name,omitempty"`
	OpenstackVolumeType *OpenStackVolumeType `xml:"openstack_volume_type,omitempty"`
	Permissions         []Permission         `xml:"permissions,omitempty"`
	PropagateErrors     *bool                `xml:"propagate_errors,omitempty"`
	ProvisionedSize     *int64               `xml:"provisioned_size,omitempty"`
	QcowVersion         QcowVersion          `xml:"qcow_version,omitempty"`
	Quota               *Quota               `xml:"quota,omitempty"`
	ReadOnly            *bool                `xml:"read_only,omitempty"`
	Sgio                ScsiGenericIO        `xml:"sgio,omitempty"`
	Shareable           *bool                `xml:"shareable,omitempty"`
	Snapshot            *Snapshot            `xml:"snapshot,omitempty"`
	Sparse              *bool                `xml:"sparse,omitempty"`
	Statistics          []Statistic          `xml:"statistics,omitempty"`
	Status              DiskStatus           `xml:"status,omitempty"`
	StorageDomain       *StorageDomain       `xml:"storage_domain,omitempty"`
	StorageDomains      []StorageDomain      `xml:"storage_domains,omitempty"`
	StorageType         DiskStorageType      `xml:"storage_type,omitempty"`
	Template            *Template            `xml:"template,omitempty"`
	UsesScsiReservation *bool                `xml:"uses_scsi_reservation,omitempty"`
	Vm                  *Vm                  `xml:"vm,omitempty"`
	Vms                 []Vm                 `xml:"vms,omitempty"`
	WipeAfterDelete     *bool                `xml:"wipe_after_delete,omitempty"`
}

type Displays struct {
	XMLName  xml.Name  `xml:"displays"`
	Displays []Display `xml:"display,omitempty"`
}

type Display struct {
	OvStruct
	Address             *string      `xml:"address,omitempty"`
	AllowOverride       *bool        `xml:"allow_override,omitempty"`
	Certificate         *Certificate `xml:"certificate,omitempty"`
	CopyPasteEnabled    *bool        `xml:"copy_paste_enabled,omitempty"`
	DisconnectAction    *string      `xml:"disconnect_action,omitempty"`
	FileTransferEnabled *bool        `xml:"file_transfer_enabled,omitempty"`
	KeyboardLayout      *string      `xml:"keyboard_layout,omitempty"`
	Monitors            *int64       `xml:"monitors,omitempty"`
	Port                *int64       `xml:"port,omitempty"`
	Proxy               *string      `xml:"proxy,omitempty"`
	SecurePort          *int64       `xml:"secure_port,omitempty"`
	SingleQxlPci        *bool        `xml:"single_qxl_pci,omitempty"`
	SmartcardEnabled    *bool        `xml:"smartcard_enabled,omitempty"`
	Type                DisplayType  `xml:"type,omitempty"`
}

type Dnss struct {
	XMLName xml.Name `xml:"dnss"`
	Dnss    []Dns    `xml:"dns,omitempty"`
}

type Dns struct {
	OvStruct
	SearchDomains []Host `xml:"search_domains,omitempty"`
	Servers       []Host `xml:"servers,omitempty"`
}

type DnsResolverConfigurations struct {
	XMLName                   xml.Name                   `xml:"dnsresolverconfigurations"`
	DnsResolverConfigurations []DnsResolverConfiguration `xml:"dns_resolver_configuration,omitempty"`
}

type DnsResolverConfiguration struct {
	OvStruct
	NameServers []string `xml:"name_servers,omitempty"`
}

type Domains struct {
	XMLName xml.Name `xml:"domains"`
	Domains []Domain `xml:"domain,omitempty"`
}

type Domain struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Groups      []Group `xml:"groups,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	User        *User   `xml:"user,omitempty"`
	Users       []User  `xml:"users,omitempty"`
}

type EntityProfileDetails struct {
	XMLName              xml.Name              `xml:"entityprofiledetails"`
	EntityProfileDetails []EntityProfileDetail `xml:"entity_profile_detail,omitempty"`
}

type EntityProfileDetail struct {
	OvStruct
	ProfileDetails []ProfileDetail `xml:"profile_details,omitempty"`
}

type ErrorHandlings struct {
	XMLName        xml.Name        `xml:"errorhandlings"`
	ErrorHandlings []ErrorHandling `xml:"error_handling,omitempty"`
}

type ErrorHandling struct {
	OvStruct
	OnError MigrateOnError `xml:"on_error,omitempty"`
}

type Events struct {
	XMLName xml.Name `xml:"events"`
	Events  []Event  `xml:"event,omitempty"`
}

type Event struct {
	OvStruct
	Cluster       *Cluster       `xml:"cluster,omitempty"`
	Code          *int64         `xml:"code,omitempty"`
	Comment       *string        `xml:"comment,omitempty"`
	CorrelationId *string        `xml:"correlation_id,omitempty"`
	CustomData    *string        `xml:"custom_data,omitempty"`
	CustomId      *int64         `xml:"custom_id,omitempty"`
	DataCenter    *DataCenter    `xml:"data_center,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	FloodRate     *int64         `xml:"flood_rate,omitempty"`
	Host          *Host          `xml:"host,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Origin        *string        `xml:"origin,omitempty"`
	Severity      LogSeverity    `xml:"severity,omitempty"`
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
	Template      *Template      `xml:"template,omitempty"`
	Time          time.Time      `xml:"time,omitempty"`
	User          *User          `xml:"user,omitempty"`
	Vm            *Vm            `xml:"vm,omitempty"`
}

type ExternalComputeResources struct {
	XMLName                  xml.Name                  `xml:"externalcomputeresources"`
	ExternalComputeResources []ExternalComputeResource `xml:"external_compute_resource,omitempty"`
}

type ExternalComputeResource struct {
	OvStruct
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
	Provider             *string               `xml:"provider,omitempty"`
	Url                  *string               `xml:"url,omitempty"`
	User                 *string               `xml:"user,omitempty"`
}

type ExternalDiscoveredHosts struct {
	XMLName                 xml.Name                 `xml:"externaldiscoveredhosts"`
	ExternalDiscoveredHosts []ExternalDiscoveredHost `xml:"external_discovered_host,omitempty"`
}

type ExternalDiscoveredHost struct {
	OvStruct
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Ip                   *string               `xml:"ip,omitempty"`
	LastReport           *string               `xml:"last_report,omitempty"`
	Mac                  *string               `xml:"mac,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
	SubnetName           *string               `xml:"subnet_name,omitempty"`
}

type ExternalHosts struct {
	XMLName       xml.Name       `xml:"externalhosts"`
	ExternalHosts []ExternalHost `xml:"external_host,omitempty"`
}

type ExternalHost struct {
	OvStruct
	Address              *string               `xml:"address,omitempty"`
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
}

type ExternalHostGroups struct {
	XMLName            xml.Name            `xml:"externalhostgroups"`
	ExternalHostGroups []ExternalHostGroup `xml:"external_host_group,omitempty"`
}

type ExternalHostGroup struct {
	OvStruct
	ArchitectureName     *string               `xml:"architecture_name,omitempty"`
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	DomainName           *string               `xml:"domain_name,omitempty"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
	OperatingSystemName  *string               `xml:"operating_system_name,omitempty"`
	SubnetName           *string               `xml:"subnet_name,omitempty"`
}

type ExternalHostProviders struct {
	XMLName               xml.Name               `xml:"externalhostproviders"`
	ExternalHostProviders []ExternalHostProvider `xml:"external_host_provider,omitempty"`
}

type ExternalHostProvider struct {
	OvStruct
	AuthenticationUrl      *string                   `xml:"authentication_url,omitempty"`
	Certificates           []Certificate             `xml:"certificates,omitempty"`
	Comment                *string                   `xml:"comment,omitempty"`
	ComputeResources       []ExternalComputeResource `xml:"compute_resources,omitempty"`
	Description            *string                   `xml:"description,omitempty"`
	DiscoveredHosts        []ExternalDiscoveredHost  `xml:"discovered_hosts,omitempty"`
	HostGroups             []ExternalHostGroup       `xml:"host_groups,omitempty"`
	Hosts                  []Host                    `xml:"hosts,omitempty"`
	Id                     *string                   `xml:"id,attr,omitempty"`
	Name                   *string                   `xml:"name,omitempty"`
	Password               *string                   `xml:"password,omitempty"`
	Properties             []Property                `xml:"properties,omitempty"`
	RequiresAuthentication *bool                     `xml:"requires_authentication,omitempty"`
	Url                    *string                   `xml:"url,omitempty"`
	Username               *string                   `xml:"username,omitempty"`
}

type ExternalProviders struct {
	XMLName           xml.Name           `xml:"externalproviders"`
	ExternalProviders []ExternalProvider `xml:"external_provider,omitempty"`
}

type ExternalProvider struct {
	OvStruct
	AuthenticationUrl      *string    `xml:"authentication_url,omitempty"`
	Comment                *string    `xml:"comment,omitempty"`
	Description            *string    `xml:"description,omitempty"`
	Id                     *string    `xml:"id,attr,omitempty"`
	Name                   *string    `xml:"name,omitempty"`
	Password               *string    `xml:"password,omitempty"`
	Properties             []Property `xml:"properties,omitempty"`
	RequiresAuthentication *bool      `xml:"requires_authentication,omitempty"`
	Url                    *string    `xml:"url,omitempty"`
	Username               *string    `xml:"username,omitempty"`
}

type ExternalVmImports struct {
	XMLName           xml.Name           `xml:"externalvmimports"`
	ExternalVmImports []ExternalVmImport `xml:"external_vm_import,omitempty"`
}

type ExternalVmImport struct {
	OvStruct
	Cluster       *Cluster               `xml:"cluster,omitempty"`
	CpuProfile    *CpuProfile            `xml:"cpu_profile,omitempty"`
	DriversIso    *File                  `xml:"drivers_iso,omitempty"`
	Host          *Host                  `xml:"host,omitempty"`
	Name          *string                `xml:"name,omitempty"`
	Password      *string                `xml:"password,omitempty"`
	Provider      ExternalVmProviderType `xml:"provider,omitempty"`
	Quota         *Quota                 `xml:"quota,omitempty"`
	Sparse        *bool                  `xml:"sparse,omitempty"`
	StorageDomain *StorageDomain         `xml:"storage_domain,omitempty"`
	Url           *string                `xml:"url,omitempty"`
	Username      *string                `xml:"username,omitempty"`
	Vm            *Vm                    `xml:"vm,omitempty"`
}

type Faults struct {
	XMLName xml.Name `xml:"faults"`
	Faults  []Fault  `xml:"fault,omitempty"`
}

type Fault struct {
	OvStruct
	Detail *string `xml:"detail,omitempty"`
	Reason *string `xml:"reason,omitempty"`
}

type FencingPolicys struct {
	XMLName        xml.Name        `xml:"fencingpolicys"`
	FencingPolicys []FencingPolicy `xml:"fencing_policy,omitempty"`
}

type FencingPolicy struct {
	OvStruct
	Enabled                   *bool                     `xml:"enabled,omitempty"`
	SkipIfConnectivityBroken  *SkipIfConnectivityBroken `xml:"skip_if_connectivity_broken,omitempty"`
	SkipIfGlusterBricksUp     *bool                     `xml:"skip_if_gluster_bricks_up,omitempty"`
	SkipIfGlusterQuorumNotMet *bool                     `xml:"skip_if_gluster_quorum_not_met,omitempty"`
	SkipIfSdActive            *SkipIfSdActive           `xml:"skip_if_sd_active,omitempty"`
}

type Files struct {
	XMLName xml.Name `xml:"files"`
	Files   []File   `xml:"file,omitempty"`
}

type File struct {
	OvStruct
	Comment       *string        `xml:"comment,omitempty"`
	Content       *string        `xml:"content,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
	Type          *string        `xml:"type,omitempty"`
}

type Filters struct {
	XMLName xml.Name `xml:"filters"`
	Filters []Filter `xml:"filter,omitempty"`
}

type Filter struct {
	OvStruct
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
	Position             *int64                `xml:"position,omitempty"`
	SchedulingPolicyUnit *SchedulingPolicyUnit `xml:"scheduling_policy_unit,omitempty"`
}

type Floppys struct {
	XMLName xml.Name `xml:"floppys"`
	Floppys []Floppy `xml:"floppy,omitempty"`
}

type Floppy struct {
	OvStruct
	Comment      *string       `xml:"comment,omitempty"`
	Description  *string       `xml:"description,omitempty"`
	File         *File         `xml:"file,omitempty"`
	Id           *string       `xml:"id,attr,omitempty"`
	InstanceType *InstanceType `xml:"instance_type,omitempty"`
	Name         *string       `xml:"name,omitempty"`
	Template     *Template     `xml:"template,omitempty"`
	Vm           *Vm           `xml:"vm,omitempty"`
	Vms          []Vm          `xml:"vms,omitempty"`
}

type FopStatistics struct {
	XMLName       xml.Name       `xml:"fopstatistics"`
	FopStatistics []FopStatistic `xml:"fop_statistic,omitempty"`
}

type FopStatistic struct {
	OvStruct
	Name       *string     `xml:"name,omitempty"`
	Statistics []Statistic `xml:"statistics,omitempty"`
}

type GlusterBricks struct {
	XMLName       xml.Name       `xml:"glusterbricks"`
	GlusterBricks []GlusterBrick `xml:"gluster_brick,omitempty"`
}

type GlusterBrick struct {
	OvStruct
	BrickDir       *string             `xml:"brick_dir,omitempty"`
	Comment        *string             `xml:"comment,omitempty"`
	Description    *string             `xml:"description,omitempty"`
	Device         *string             `xml:"device,omitempty"`
	FsName         *string             `xml:"fs_name,omitempty"`
	GlusterClients []GlusterClient     `xml:"gluster_clients,omitempty"`
	GlusterVolume  *GlusterVolume      `xml:"gluster_volume,omitempty"`
	Id             *string             `xml:"id,attr,omitempty"`
	InstanceType   *InstanceType       `xml:"instance_type,omitempty"`
	MemoryPools    []GlusterMemoryPool `xml:"memory_pools,omitempty"`
	MntOptions     *string             `xml:"mnt_options,omitempty"`
	Name           *string             `xml:"name,omitempty"`
	Pid            *int64              `xml:"pid,omitempty"`
	Port           *int64              `xml:"port,omitempty"`
	ServerId       *string             `xml:"server_id,omitempty"`
	Statistics     []Statistic         `xml:"statistics,omitempty"`
	Status         GlusterBrickStatus  `xml:"status,omitempty"`
	Template       *Template           `xml:"template,omitempty"`
	Vm             *Vm                 `xml:"vm,omitempty"`
	Vms            []Vm                `xml:"vms,omitempty"`
}

type GlusterBrickAdvancedDetailss struct {
	XMLName                      xml.Name                      `xml:"glusterbrickadvanceddetailss"`
	GlusterBrickAdvancedDetailss []GlusterBrickAdvancedDetails `xml:"gluster_brick_advanced_details,omitempty"`
}

type GlusterBrickAdvancedDetails struct {
	OvStruct
	Comment        *string             `xml:"comment,omitempty"`
	Description    *string             `xml:"description,omitempty"`
	Device         *string             `xml:"device,omitempty"`
	FsName         *string             `xml:"fs_name,omitempty"`
	GlusterClients []GlusterClient     `xml:"gluster_clients,omitempty"`
	Id             *string             `xml:"id,attr,omitempty"`
	InstanceType   *InstanceType       `xml:"instance_type,omitempty"`
	MemoryPools    []GlusterMemoryPool `xml:"memory_pools,omitempty"`
	MntOptions     *string             `xml:"mnt_options,omitempty"`
	Name           *string             `xml:"name,omitempty"`
	Pid            *int64              `xml:"pid,omitempty"`
	Port           *int64              `xml:"port,omitempty"`
	Template       *Template           `xml:"template,omitempty"`
	Vm             *Vm                 `xml:"vm,omitempty"`
	Vms            []Vm                `xml:"vms,omitempty"`
}

type GlusterBrickMemoryInfos struct {
	XMLName                 xml.Name                 `xml:"glusterbrickmemoryinfos"`
	GlusterBrickMemoryInfos []GlusterBrickMemoryInfo `xml:"gluster_brick_memory_info,omitempty"`
}

type GlusterBrickMemoryInfo struct {
	OvStruct
	MemoryPools []GlusterMemoryPool `xml:"memory_pools,omitempty"`
}

type GlusterClients struct {
	XMLName        xml.Name        `xml:"glusterclients"`
	GlusterClients []GlusterClient `xml:"gluster_client,omitempty"`
}

type GlusterClient struct {
	OvStruct
	BytesRead    *int64  `xml:"bytes_read,omitempty"`
	BytesWritten *int64  `xml:"bytes_written,omitempty"`
	ClientPort   *int64  `xml:"client_port,omitempty"`
	HostName     *string `xml:"host_name,omitempty"`
}

type GlusterHooks struct {
	XMLName      xml.Name      `xml:"glusterhooks"`
	GlusterHooks []GlusterHook `xml:"gluster_hook,omitempty"`
}

type GlusterHook struct {
	OvStruct
	Checksum       *string             `xml:"checksum,omitempty"`
	Cluster        *Cluster            `xml:"cluster,omitempty"`
	Comment        *string             `xml:"comment,omitempty"`
	ConflictStatus *int64              `xml:"conflict_status,omitempty"`
	Conflicts      *string             `xml:"conflicts,omitempty"`
	Content        *string             `xml:"content,omitempty"`
	ContentType    HookContentType     `xml:"content_type,omitempty"`
	Description    *string             `xml:"description,omitempty"`
	GlusterCommand *string             `xml:"gluster_command,omitempty"`
	Id             *string             `xml:"id,attr,omitempty"`
	Name           *string             `xml:"name,omitempty"`
	ServerHooks    []GlusterServerHook `xml:"server_hooks,omitempty"`
	Stage          HookStage           `xml:"stage,omitempty"`
	Status         GlusterHookStatus   `xml:"status,omitempty"`
}

type GlusterMemoryPools struct {
	XMLName            xml.Name            `xml:"glustermemorypools"`
	GlusterMemoryPools []GlusterMemoryPool `xml:"gluster_memory_pool,omitempty"`
}

type GlusterMemoryPool struct {
	OvStruct
	AllocCount  *int64  `xml:"alloc_count,omitempty"`
	ColdCount   *int64  `xml:"cold_count,omitempty"`
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	HotCount    *int64  `xml:"hot_count,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	MaxAlloc    *int64  `xml:"max_alloc,omitempty"`
	MaxStdalloc *int64  `xml:"max_stdalloc,omitempty"`
	Name        *string `xml:"name,omitempty"`
	PaddedSize  *int64  `xml:"padded_size,omitempty"`
	PoolMisses  *int64  `xml:"pool_misses,omitempty"`
	Type        *string `xml:"type,omitempty"`
}

type GlusterServerHooks struct {
	XMLName            xml.Name            `xml:"glusterserverhooks"`
	GlusterServerHooks []GlusterServerHook `xml:"gluster_server_hook,omitempty"`
}

type GlusterServerHook struct {
	OvStruct
	Checksum    *string           `xml:"checksum,omitempty"`
	Comment     *string           `xml:"comment,omitempty"`
	ContentType HookContentType   `xml:"content_type,omitempty"`
	Description *string           `xml:"description,omitempty"`
	Host        *Host             `xml:"host,omitempty"`
	Id          *string           `xml:"id,attr,omitempty"`
	Name        *string           `xml:"name,omitempty"`
	Status      GlusterHookStatus `xml:"status,omitempty"`
}

type GlusterVolumes struct {
	XMLName        xml.Name        `xml:"glustervolumes"`
	GlusterVolumes []GlusterVolume `xml:"gluster_volume,omitempty"`
}

type GlusterVolume struct {
	OvStruct
	Bricks          []GlusterBrick      `xml:"bricks,omitempty"`
	Cluster         *Cluster            `xml:"cluster,omitempty"`
	Comment         *string             `xml:"comment,omitempty"`
	Description     *string             `xml:"description,omitempty"`
	DisperseCount   *int64              `xml:"disperse_count,omitempty"`
	Id              *string             `xml:"id,attr,omitempty"`
	Name            *string             `xml:"name,omitempty"`
	Options         []Option            `xml:"options,omitempty"`
	RedundancyCount *int64              `xml:"redundancy_count,omitempty"`
	ReplicaCount    *int64              `xml:"replica_count,omitempty"`
	Statistics      []Statistic         `xml:"statistics,omitempty"`
	Status          GlusterVolumeStatus `xml:"status,omitempty"`
	StripeCount     *int64              `xml:"stripe_count,omitempty"`
	TransportTypes  []TransportType     `xml:"transport_types,omitempty"`
	VolumeType      GlusterVolumeType   `xml:"volume_type,omitempty"`
}

type GlusterVolumeProfileDetailss struct {
	XMLName                      xml.Name                      `xml:"glustervolumeprofiledetailss"`
	GlusterVolumeProfileDetailss []GlusterVolumeProfileDetails `xml:"gluster_volume_profile_details,omitempty"`
}

type GlusterVolumeProfileDetails struct {
	OvStruct
	BrickProfileDetails []BrickProfileDetail `xml:"brick_profile_details,omitempty"`
	Comment             *string              `xml:"comment,omitempty"`
	Description         *string              `xml:"description,omitempty"`
	Id                  *string              `xml:"id,attr,omitempty"`
	Name                *string              `xml:"name,omitempty"`
	NfsProfileDetails   []NfsProfileDetail   `xml:"nfs_profile_details,omitempty"`
}

type GracePeriods struct {
	XMLName      xml.Name      `xml:"graceperiods"`
	GracePeriods []GracePeriod `xml:"grace_period,omitempty"`
}

type GracePeriod struct {
	OvStruct
	Expiry *int64 `xml:"expiry,omitempty"`
}

type GraphicsConsoles struct {
	XMLName          xml.Name          `xml:"graphicsconsoles"`
	GraphicsConsoles []GraphicsConsole `xml:"graphics_console,omitempty"`
}

type GraphicsConsole struct {
	OvStruct
	Address      *string       `xml:"address,omitempty"`
	Comment      *string       `xml:"comment,omitempty"`
	Description  *string       `xml:"description,omitempty"`
	Id           *string       `xml:"id,attr,omitempty"`
	InstanceType *InstanceType `xml:"instance_type,omitempty"`
	Name         *string       `xml:"name,omitempty"`
	Port         *int64        `xml:"port,omitempty"`
	Protocol     GraphicsType  `xml:"protocol,omitempty"`
	Template     *Template     `xml:"template,omitempty"`
	TlsPort      *int64        `xml:"tls_port,omitempty"`
	Vm           *Vm           `xml:"vm,omitempty"`
}

type Groups struct {
	XMLName xml.Name `xml:"groups"`
	Groups  []Group  `xml:"group,omitempty"`
}

type Group struct {
	OvStruct
	Comment       *string      `xml:"comment,omitempty"`
	Description   *string      `xml:"description,omitempty"`
	Domain        *Domain      `xml:"domain,omitempty"`
	DomainEntryId *string      `xml:"domain_entry_id,omitempty"`
	Id            *string      `xml:"id,attr,omitempty"`
	Name          *string      `xml:"name,omitempty"`
	Namespace     *string      `xml:"namespace,omitempty"`
	Permissions   []Permission `xml:"permissions,omitempty"`
	Roles         []Role       `xml:"roles,omitempty"`
	Tags          []Tag        `xml:"tags,omitempty"`
}

type GuestOperatingSystems struct {
	XMLName               xml.Name               `xml:"guestoperatingsystems"`
	GuestOperatingSystems []GuestOperatingSystem `xml:"guest_operating_system,omitempty"`
}

type GuestOperatingSystem struct {
	OvStruct
	Architecture *string  `xml:"architecture,omitempty"`
	Codename     *string  `xml:"codename,omitempty"`
	Distribution *string  `xml:"distribution,omitempty"`
	Family       *string  `xml:"family,omitempty"`
	Kernel       *Kernel  `xml:"kernel,omitempty"`
	Version      *Version `xml:"version,omitempty"`
}

type HardwareInformations struct {
	XMLName              xml.Name              `xml:"hardwareinformations"`
	HardwareInformations []HardwareInformation `xml:"hardware_information,omitempty"`
}

type HardwareInformation struct {
	OvStruct
	Family              *string     `xml:"family,omitempty"`
	Manufacturer        *string     `xml:"manufacturer,omitempty"`
	ProductName         *string     `xml:"product_name,omitempty"`
	SerialNumber        *string     `xml:"serial_number,omitempty"`
	SupportedRngSources []RngSource `xml:"supported_rng_sources,omitempty"`
	Uuid                *string     `xml:"uuid,omitempty"`
	Version             *string     `xml:"version,omitempty"`
}

type HighAvailabilitys struct {
	XMLName           xml.Name           `xml:"highavailabilitys"`
	HighAvailabilitys []HighAvailability `xml:"high_availability,omitempty"`
}

type HighAvailability struct {
	OvStruct
	Enabled  *bool  `xml:"enabled,omitempty"`
	Priority *int64 `xml:"priority,omitempty"`
}

type Hooks struct {
	XMLName xml.Name `xml:"hooks"`
	Hooks   []Hook   `xml:"hook,omitempty"`
}

type Hook struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	EventName   *string `xml:"event_name,omitempty"`
	Host        *Host   `xml:"host,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Md5         *string `xml:"md5,omitempty"`
	Name        *string `xml:"name,omitempty"`
}

type Hosts struct {
	XMLName xml.Name `xml:"hosts"`
	Hosts   []Host   `xml:"host,omitempty"`
}

type Host struct {
	OvStruct
	Address                     *string                      `xml:"address,omitempty"`
	AffinityLabels              []AffinityLabel              `xml:"affinity_labels,omitempty"`
	Agents                      []Agent                      `xml:"agents,omitempty"`
	AutoNumaStatus              AutoNumaStatus               `xml:"auto_numa_status,omitempty"`
	Certificate                 *Certificate                 `xml:"certificate,omitempty"`
	Cluster                     *Cluster                     `xml:"cluster,omitempty"`
	Comment                     *string                      `xml:"comment,omitempty"`
	Cpu                         *Cpu                         `xml:"cpu,omitempty"`
	Description                 *string                      `xml:"description,omitempty"`
	DevicePassthrough           *HostDevicePassthrough       `xml:"device_passthrough,omitempty"`
	Devices                     []Device                     `xml:"devices,omitempty"`
	Display                     *Display                     `xml:"display,omitempty"`
	ExternalHostProvider        *ExternalHostProvider        `xml:"external_host_provider,omitempty"`
	ExternalStatus              ExternalStatus               `xml:"external_status,omitempty"`
	HardwareInformation         *HardwareInformation         `xml:"hardware_information,omitempty"`
	Hooks                       []Hook                       `xml:"hooks,omitempty"`
	HostedEngine                *HostedEngine                `xml:"hosted_engine,omitempty"`
	Id                          *string                      `xml:"id,attr,omitempty"`
	Iscsi                       *IscsiDetails                `xml:"iscsi,omitempty"`
	KatelloErrata               []KatelloErratum             `xml:"katello_errata,omitempty"`
	KdumpStatus                 KdumpStatus                  `xml:"kdump_status,omitempty"`
	Ksm                         *Ksm                         `xml:"ksm,omitempty"`
	LibvirtVersion              *Version                     `xml:"libvirt_version,omitempty"`
	MaxSchedulingMemory         *int64                       `xml:"max_scheduling_memory,omitempty"`
	Memory                      *int64                       `xml:"memory,omitempty"`
	Name                        *string                      `xml:"name,omitempty"`
	NetworkAttachments          []NetworkAttachment          `xml:"network_attachments,omitempty"`
	Nics                        []Nic                        `xml:"nics,omitempty"`
	NumaNodes                   []NumaNode                   `xml:"numa_nodes,omitempty"`
	NumaSupported               *bool                        `xml:"numa_supported,omitempty"`
	Os                          *OperatingSystem             `xml:"os,omitempty"`
	OverrideIptables            *bool                        `xml:"override_iptables,omitempty"`
	Permissions                 []Permission                 `xml:"permissions,omitempty"`
	Port                        *int64                       `xml:"port,omitempty"`
	PowerManagement             *PowerManagement             `xml:"power_management,omitempty"`
	Protocol                    HostProtocol                 `xml:"protocol,omitempty"`
	RootPassword                *string                      `xml:"root_password,omitempty"`
	SeLinux                     *SeLinux                     `xml:"se_linux,omitempty"`
	Spm                         *Spm                         `xml:"spm,omitempty"`
	Ssh                         *Ssh                         `xml:"ssh,omitempty"`
	Statistics                  []Statistic                  `xml:"statistics,omitempty"`
	Status                      HostStatus                   `xml:"status,omitempty"`
	StatusDetail                *string                      `xml:"status_detail,omitempty"`
	StorageConnectionExtensions []StorageConnectionExtension `xml:"storage_connection_extensions,omitempty"`
	Storages                    []HostStorage                `xml:"storages,omitempty"`
	Summary                     *VmSummary                   `xml:"summary,omitempty"`
	Tags                        []Tag                        `xml:"tags,omitempty"`
	TransparentHugePages        *TransparentHugePages        `xml:"transparent_huge_pages,omitempty"`
	Type                        HostType                     `xml:"type,omitempty"`
	UnmanagedNetworks           []UnmanagedNetwork           `xml:"unmanaged_networks,omitempty"`
	UpdateAvailable             *bool                        `xml:"update_available,omitempty"`
	Version                     *Version                     `xml:"version,omitempty"`
}

type HostDevices struct {
	XMLName     xml.Name     `xml:"hostdevices"`
	HostDevices []HostDevice `xml:"host_device,omitempty"`
}

type HostDevice struct {
	OvStruct
	Capability       *string     `xml:"capability,omitempty"`
	Comment          *string     `xml:"comment,omitempty"`
	Description      *string     `xml:"description,omitempty"`
	Host             *Host       `xml:"host,omitempty"`
	Id               *string     `xml:"id,attr,omitempty"`
	IommuGroup       *int64      `xml:"iommu_group,omitempty"`
	Name             *string     `xml:"name,omitempty"`
	ParentDevice     *HostDevice `xml:"parent_device,omitempty"`
	PhysicalFunction *HostDevice `xml:"physical_function,omitempty"`
	Placeholder      *bool       `xml:"placeholder,omitempty"`
	Product          *Product    `xml:"product,omitempty"`
	Vendor           *Vendor     `xml:"vendor,omitempty"`
	VirtualFunctions *int64      `xml:"virtual_functions,omitempty"`
	Vm               *Vm         `xml:"vm,omitempty"`
}

type HostDevicePassthroughs struct {
	XMLName                xml.Name                `xml:"hostdevicepassthroughs"`
	HostDevicePassthroughs []HostDevicePassthrough `xml:"host_device_passthrough,omitempty"`
}

type HostDevicePassthrough struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
}

type HostNics struct {
	XMLName  xml.Name  `xml:"hostnics"`
	HostNics []HostNic `xml:"host_nic,omitempty"`
}

type HostNic struct {
	OvStruct
	AdAggregatorId                *int64                                `xml:"ad_aggregator_id,omitempty"`
	BaseInterface                 *string                               `xml:"base_interface,omitempty"`
	Bonding                       *Bonding                              `xml:"bonding,omitempty"`
	BootProtocol                  BootProtocol                          `xml:"boot_protocol,omitempty"`
	Bridged                       *bool                                 `xml:"bridged,omitempty"`
	CheckConnectivity             *bool                                 `xml:"check_connectivity,omitempty"`
	Comment                       *string                               `xml:"comment,omitempty"`
	CustomConfiguration           *bool                                 `xml:"custom_configuration,omitempty"`
	Description                   *string                               `xml:"description,omitempty"`
	Host                          *Host                                 `xml:"host,omitempty"`
	Id                            *string                               `xml:"id,attr,omitempty"`
	Ip                            *Ip                                   `xml:"ip,omitempty"`
	Ipv6                          *Ip                                   `xml:"ipv6,omitempty"`
	Ipv6BootProtocol              BootProtocol                          `xml:"ipv6_boot_protocol,omitempty"`
	Mac                           *Mac                                  `xml:"mac,omitempty"`
	Mtu                           *int64                                `xml:"mtu,omitempty"`
	Name                          *string                               `xml:"name,omitempty"`
	Network                       *Network                              `xml:"network,omitempty"`
	NetworkLabels                 []NetworkLabel                        `xml:"network_labels,omitempty"`
	OverrideConfiguration         *bool                                 `xml:"override_configuration,omitempty"`
	PhysicalFunction              *HostNic                              `xml:"physical_function,omitempty"`
	Properties                    []Property                            `xml:"properties,omitempty"`
	Qos                           *Qos                                  `xml:"qos,omitempty"`
	Speed                         *int64                                `xml:"speed,omitempty"`
	Statistics                    []Statistic                           `xml:"statistics,omitempty"`
	Status                        NicStatus                             `xml:"status,omitempty"`
	VirtualFunctionsConfiguration *HostNicVirtualFunctionsConfiguration `xml:"virtual_functions_configuration,omitempty"`
	Vlan                          *Vlan                                 `xml:"vlan,omitempty"`
}

type HostNicVirtualFunctionsConfigurations struct {
	XMLName                               xml.Name                               `xml:"hostnicvirtualfunctionsconfigurations"`
	HostNicVirtualFunctionsConfigurations []HostNicVirtualFunctionsConfiguration `xml:"host_nic_virtual_functions_configuration,omitempty"`
}

type HostNicVirtualFunctionsConfiguration struct {
	OvStruct
	AllNetworksAllowed          *bool  `xml:"all_networks_allowed,omitempty"`
	MaxNumberOfVirtualFunctions *int64 `xml:"max_number_of_virtual_functions,omitempty"`
	NumberOfVirtualFunctions    *int64 `xml:"number_of_virtual_functions,omitempty"`
}

type HostStorages struct {
	XMLName      xml.Name      `xml:"hoststorages"`
	HostStorages []HostStorage `xml:"host_storage,omitempty"`
}

type HostStorage struct {
	OvStruct
	Address      *string       `xml:"address,omitempty"`
	Comment      *string       `xml:"comment,omitempty"`
	Description  *string       `xml:"description,omitempty"`
	Host         *Host         `xml:"host,omitempty"`
	Id           *string       `xml:"id,attr,omitempty"`
	LogicalUnits []LogicalUnit `xml:"logical_units,omitempty"`
	MountOptions *string       `xml:"mount_options,omitempty"`
	Name         *string       `xml:"name,omitempty"`
	NfsRetrans   *int64        `xml:"nfs_retrans,omitempty"`
	NfsTimeo     *int64        `xml:"nfs_timeo,omitempty"`
	NfsVersion   NfsVersion    `xml:"nfs_version,omitempty"`
	OverrideLuns *bool         `xml:"override_luns,omitempty"`
	Password     *string       `xml:"password,omitempty"`
	Path         *string       `xml:"path,omitempty"`
	Port         *int64        `xml:"port,omitempty"`
	Portal       *string       `xml:"portal,omitempty"`
	Target       *string       `xml:"target,omitempty"`
	Type         StorageType   `xml:"type,omitempty"`
	Username     *string       `xml:"username,omitempty"`
	VfsType      *string       `xml:"vfs_type,omitempty"`
	VolumeGroup  *VolumeGroup  `xml:"volume_group,omitempty"`
}

type HostedEngines struct {
	XMLName       xml.Name       `xml:"hostedengines"`
	HostedEngines []HostedEngine `xml:"hosted_engine,omitempty"`
}

type HostedEngine struct {
	OvStruct
	Active            *bool  `xml:"active,omitempty"`
	Configured        *bool  `xml:"configured,omitempty"`
	GlobalMaintenance *bool  `xml:"global_maintenance,omitempty"`
	LocalMaintenance  *bool  `xml:"local_maintenance,omitempty"`
	Score             *int64 `xml:"score,omitempty"`
}

type Icons struct {
	XMLName xml.Name `xml:"icons"`
	Icons   []Icon   `xml:"icon,omitempty"`
}

type Icon struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Data        *string `xml:"data,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	MediaType   *string `xml:"media_type,omitempty"`
	Name        *string `xml:"name,omitempty"`
}

type Identifieds struct {
	XMLName     xml.Name     `xml:"identifieds"`
	Identifieds []Identified `xml:"identified,omitempty"`
}

type Identified struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
}

type Images struct {
	XMLName xml.Name `xml:"images"`
	Images  []Image  `xml:"image,omitempty"`
}

type Image struct {
	OvStruct
	Comment       *string        `xml:"comment,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
}

type ImageTransfers struct {
	XMLName        xml.Name        `xml:"imagetransfers"`
	ImageTransfers []ImageTransfer `xml:"image_transfer,omitempty"`
}

type ImageTransfer struct {
	OvStruct
	Comment      *string                `xml:"comment,omitempty"`
	Description  *string                `xml:"description,omitempty"`
	Direction    ImageTransferDirection `xml:"direction,omitempty"`
	Host         *Host                  `xml:"host,omitempty"`
	Id           *string                `xml:"id,attr,omitempty"`
	Image        *Image                 `xml:"image,omitempty"`
	Name         *string                `xml:"name,omitempty"`
	Phase        ImageTransferPhase     `xml:"phase,omitempty"`
	ProxyUrl     *string                `xml:"proxy_url,omitempty"`
	SignedTicket *string                `xml:"signed_ticket,omitempty"`
}

type Initializations struct {
	XMLName         xml.Name         `xml:"initializations"`
	Initializations []Initialization `xml:"initialization,omitempty"`
}

type Initialization struct {
	OvStruct
	ActiveDirectoryOu *string            `xml:"active_directory_ou,omitempty"`
	AuthorizedSshKeys *string            `xml:"authorized_ssh_keys,omitempty"`
	CloudInit         *CloudInit         `xml:"cloud_init,omitempty"`
	Configuration     *Configuration     `xml:"configuration,omitempty"`
	CustomScript      *string            `xml:"custom_script,omitempty"`
	DnsSearch         *string            `xml:"dns_search,omitempty"`
	DnsServers        *string            `xml:"dns_servers,omitempty"`
	Domain            *string            `xml:"domain,omitempty"`
	HostName          *string            `xml:"host_name,omitempty"`
	InputLocale       *string            `xml:"input_locale,omitempty"`
	NicConfigurations []NicConfiguration `xml:"nic_configurations,omitempty"`
	OrgName           *string            `xml:"org_name,omitempty"`
	RegenerateIds     *bool              `xml:"regenerate_ids,omitempty"`
	RegenerateSshKeys *bool              `xml:"regenerate_ssh_keys,omitempty"`
	RootPassword      *string            `xml:"root_password,omitempty"`
	SystemLocale      *string            `xml:"system_locale,omitempty"`
	Timezone          *string            `xml:"timezone,omitempty"`
	UiLanguage        *string            `xml:"ui_language,omitempty"`
	UserLocale        *string            `xml:"user_locale,omitempty"`
	UserName          *string            `xml:"user_name,omitempty"`
	WindowsLicenseKey *string            `xml:"windows_license_key,omitempty"`
}

type InstanceTypes struct {
	XMLName       xml.Name       `xml:"instancetypes"`
	InstanceTypes []InstanceType `xml:"instance_type,omitempty"`
}

type InstanceType struct {
	OvStruct
	Bios                       *Bios               `xml:"bios,omitempty"`
	Cdroms                     []Cdrom             `xml:"cdroms,omitempty"`
	Cluster                    *Cluster            `xml:"cluster,omitempty"`
	Comment                    *string             `xml:"comment,omitempty"`
	Console                    *Console            `xml:"console,omitempty"`
	Cpu                        *Cpu                `xml:"cpu,omitempty"`
	CpuProfile                 *CpuProfile         `xml:"cpu_profile,omitempty"`
	CpuShares                  *int64              `xml:"cpu_shares,omitempty"`
	CreationTime               time.Time           `xml:"creation_time,omitempty"`
	CustomCompatibilityVersion *Version            `xml:"custom_compatibility_version,omitempty"`
	CustomCpuModel             *string             `xml:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      *string             `xml:"custom_emulated_machine,omitempty"`
	CustomProperties           []CustomProperty    `xml:"custom_properties,omitempty"`
	DeleteProtected            *bool               `xml:"delete_protected,omitempty"`
	Description                *string             `xml:"description,omitempty"`
	DiskAttachments            []DiskAttachment    `xml:"disk_attachments,omitempty"`
	Display                    *Display            `xml:"display,omitempty"`
	Domain                     *Domain             `xml:"domain,omitempty"`
	GraphicsConsoles           []GraphicsConsole   `xml:"graphics_consoles,omitempty"`
	HighAvailability           *HighAvailability   `xml:"high_availability,omitempty"`
	Id                         *string             `xml:"id,attr,omitempty"`
	Initialization             *Initialization     `xml:"initialization,omitempty"`
	Io                         *Io                 `xml:"io,omitempty"`
	LargeIcon                  *Icon               `xml:"large_icon,omitempty"`
	Lease                      *StorageDomainLease `xml:"lease,omitempty"`
	Memory                     *int64              `xml:"memory,omitempty"`
	MemoryPolicy               *MemoryPolicy       `xml:"memory_policy,omitempty"`
	Migration                  *MigrationOptions   `xml:"migration,omitempty"`
	MigrationDowntime          *int64              `xml:"migration_downtime,omitempty"`
	Name                       *string             `xml:"name,omitempty"`
	Nics                       []Nic               `xml:"nics,omitempty"`
	Origin                     *string             `xml:"origin,omitempty"`
	Os                         *OperatingSystem    `xml:"os,omitempty"`
	Permissions                []Permission        `xml:"permissions,omitempty"`
	Quota                      *Quota              `xml:"quota,omitempty"`
	RngDevice                  *RngDevice          `xml:"rng_device,omitempty"`
	SerialNumber               *SerialNumber       `xml:"serial_number,omitempty"`
	SmallIcon                  *Icon               `xml:"small_icon,omitempty"`
	SoundcardEnabled           *bool               `xml:"soundcard_enabled,omitempty"`
	Sso                        *Sso                `xml:"sso,omitempty"`
	StartPaused                *bool               `xml:"start_paused,omitempty"`
	Stateless                  *bool               `xml:"stateless,omitempty"`
	Status                     TemplateStatus      `xml:"status,omitempty"`
	StorageDomain              *StorageDomain      `xml:"storage_domain,omitempty"`
	Tags                       []Tag               `xml:"tags,omitempty"`
	TimeZone                   *TimeZone           `xml:"time_zone,omitempty"`
	TunnelMigration            *bool               `xml:"tunnel_migration,omitempty"`
	Type                       VmType              `xml:"type,omitempty"`
	Usb                        *Usb                `xml:"usb,omitempty"`
	Version                    *TemplateVersion    `xml:"version,omitempty"`
	VirtioScsi                 *VirtioScsi         `xml:"virtio_scsi,omitempty"`
	Vm                         *Vm                 `xml:"vm,omitempty"`
	Watchdogs                  []Watchdog          `xml:"watchdogs,omitempty"`
}

type Ios struct {
	XMLName xml.Name `xml:"ios"`
	Ios     []Io     `xml:"io,omitempty"`
}

type Io struct {
	OvStruct
	Threads *int64 `xml:"threads,omitempty"`
}

type Ips struct {
	XMLName xml.Name `xml:"ips"`
	Ips     []Ip     `xml:"ip,omitempty"`
}

type Ip struct {
	OvStruct
	Address *string   `xml:"address,omitempty"`
	Gateway *string   `xml:"gateway,omitempty"`
	Netmask *string   `xml:"netmask,omitempty"`
	Version IpVersion `xml:"version,omitempty"`
}

type IpAddressAssignments struct {
	XMLName              xml.Name              `xml:"ipaddressassignments"`
	IpAddressAssignments []IpAddressAssignment `xml:"ip_address_assignment,omitempty"`
}

type IpAddressAssignment struct {
	OvStruct
	AssignmentMethod BootProtocol `xml:"assignment_method,omitempty"`
	Ip               *Ip          `xml:"ip,omitempty"`
}

type IscsiBonds struct {
	XMLName    xml.Name    `xml:"iscsibonds"`
	IscsiBonds []IscsiBond `xml:"iscsi_bond,omitempty"`
}

type IscsiBond struct {
	OvStruct
	Comment            *string             `xml:"comment,omitempty"`
	DataCenter         *DataCenter         `xml:"data_center,omitempty"`
	Description        *string             `xml:"description,omitempty"`
	Id                 *string             `xml:"id,attr,omitempty"`
	Name               *string             `xml:"name,omitempty"`
	Networks           []Network           `xml:"networks,omitempty"`
	StorageConnections []StorageConnection `xml:"storage_connections,omitempty"`
}

type IscsiDetailss struct {
	XMLName       xml.Name       `xml:"iscsidetailss"`
	IscsiDetailss []IscsiDetails `xml:"iscsi_details,omitempty"`
}

type IscsiDetails struct {
	OvStruct
	Address         *string `xml:"address,omitempty"`
	DiskId          *string `xml:"disk_id,omitempty"`
	Initiator       *string `xml:"initiator,omitempty"`
	LunMapping      *int64  `xml:"lun_mapping,omitempty"`
	Password        *string `xml:"password,omitempty"`
	Paths           *int64  `xml:"paths,omitempty"`
	Port            *int64  `xml:"port,omitempty"`
	Portal          *string `xml:"portal,omitempty"`
	ProductId       *string `xml:"product_id,omitempty"`
	Serial          *string `xml:"serial,omitempty"`
	Size            *int64  `xml:"size,omitempty"`
	Status          *string `xml:"status,omitempty"`
	StorageDomainId *string `xml:"storage_domain_id,omitempty"`
	Target          *string `xml:"target,omitempty"`
	Username        *string `xml:"username,omitempty"`
	VendorId        *string `xml:"vendor_id,omitempty"`
	VolumeGroupId   *string `xml:"volume_group_id,omitempty"`
}

type Jobs struct {
	XMLName xml.Name `xml:"jobs"`
	Jobs    []Job    `xml:"job,omitempty"`
}

type Job struct {
	OvStruct
	AutoCleared *bool     `xml:"auto_cleared,omitempty"`
	Comment     *string   `xml:"comment,omitempty"`
	Description *string   `xml:"description,omitempty"`
	EndTime     time.Time `xml:"end_time,omitempty"`
	External    *bool     `xml:"external,omitempty"`
	Id          *string   `xml:"id,attr,omitempty"`
	LastUpdated time.Time `xml:"last_updated,omitempty"`
	Name        *string   `xml:"name,omitempty"`
	Owner       *User     `xml:"owner,omitempty"`
	StartTime   time.Time `xml:"start_time,omitempty"`
	Status      JobStatus `xml:"status,omitempty"`
	Steps       []Step    `xml:"steps,omitempty"`
}

type KatelloErratums struct {
	XMLName         xml.Name         `xml:"katelloerratums"`
	KatelloErratums []KatelloErratum `xml:"katello_erratum,omitempty"`
}

type KatelloErratum struct {
	OvStruct
	Comment     *string   `xml:"comment,omitempty"`
	Description *string   `xml:"description,omitempty"`
	Host        *Host     `xml:"host,omitempty"`
	Id          *string   `xml:"id,attr,omitempty"`
	Issued      time.Time `xml:"issued,omitempty"`
	Name        *string   `xml:"name,omitempty"`
	Packages    []Package `xml:"packages,omitempty"`
	Severity    *string   `xml:"severity,omitempty"`
	Solution    *string   `xml:"solution,omitempty"`
	Summary     *string   `xml:"summary,omitempty"`
	Title       *string   `xml:"title,omitempty"`
	Type        *string   `xml:"type,omitempty"`
	Vm          *Vm       `xml:"vm,omitempty"`
}

type Kernels struct {
	XMLName xml.Name `xml:"kernels"`
	Kernels []Kernel `xml:"kernel,omitempty"`
}

type Kernel struct {
	OvStruct
	Version *Version `xml:"version,omitempty"`
}

type Ksms struct {
	XMLName xml.Name `xml:"ksms"`
	Ksms    []Ksm    `xml:"ksm,omitempty"`
}

type Ksm struct {
	OvStruct
	Enabled          *bool `xml:"enabled,omitempty"`
	MergeAcrossNodes *bool `xml:"merge_across_nodes,omitempty"`
}

type LogicalUnits struct {
	XMLName      xml.Name      `xml:"logicalunits"`
	LogicalUnits []LogicalUnit `xml:"logical_unit,omitempty"`
}

type LogicalUnit struct {
	OvStruct
	Address           *string   `xml:"address,omitempty"`
	DiscardMaxSize    *int64    `xml:"discard_max_size,omitempty"`
	DiscardZeroesData *bool     `xml:"discard_zeroes_data,omitempty"`
	DiskId            *string   `xml:"disk_id,omitempty"`
	Id                *string   `xml:"id,attr,omitempty"`
	LunMapping        *int64    `xml:"lun_mapping,omitempty"`
	Password          *string   `xml:"password,omitempty"`
	Paths             *int64    `xml:"paths,omitempty"`
	Port              *int64    `xml:"port,omitempty"`
	Portal            *string   `xml:"portal,omitempty"`
	ProductId         *string   `xml:"product_id,omitempty"`
	Serial            *string   `xml:"serial,omitempty"`
	Size              *int64    `xml:"size,omitempty"`
	Status            LunStatus `xml:"status,omitempty"`
	StorageDomainId   *string   `xml:"storage_domain_id,omitempty"`
	Target            *string   `xml:"target,omitempty"`
	Username          *string   `xml:"username,omitempty"`
	VendorId          *string   `xml:"vendor_id,omitempty"`
	VolumeGroupId     *string   `xml:"volume_group_id,omitempty"`
}

type Macs struct {
	XMLName xml.Name `xml:"macs"`
	Macs    []Mac    `xml:"mac,omitempty"`
}

type Mac struct {
	OvStruct
	Address *string `xml:"address,omitempty"`
}

type MacPools struct {
	XMLName  xml.Name  `xml:"macpools"`
	MacPools []MacPool `xml:"mac_pool,omitempty"`
}

type MacPool struct {
	OvStruct
	AllowDuplicates *bool   `xml:"allow_duplicates,omitempty"`
	Comment         *string `xml:"comment,omitempty"`
	DefaultPool     *bool   `xml:"default_pool,omitempty"`
	Description     *string `xml:"description,omitempty"`
	Id              *string `xml:"id,attr,omitempty"`
	Name            *string `xml:"name,omitempty"`
	Ranges          []Range `xml:"ranges,omitempty"`
}

type MemoryOverCommits struct {
	XMLName           xml.Name           `xml:"memoryovercommits"`
	MemoryOverCommits []MemoryOverCommit `xml:"memory_over_commit,omitempty"`
}

type MemoryOverCommit struct {
	OvStruct
	Percent *int64 `xml:"percent,omitempty"`
}

type MemoryPolicys struct {
	XMLName       xml.Name       `xml:"memorypolicys"`
	MemoryPolicys []MemoryPolicy `xml:"memory_policy,omitempty"`
}

type MemoryPolicy struct {
	OvStruct
	Ballooning           *bool                 `xml:"ballooning,omitempty"`
	Guaranteed           *int64                `xml:"guaranteed,omitempty"`
	Max                  *int64                `xml:"max,omitempty"`
	OverCommit           *MemoryOverCommit     `xml:"over_commit,omitempty"`
	TransparentHugePages *TransparentHugePages `xml:"transparent_huge_pages,omitempty"`
}

type Methods struct {
	XMLName xml.Name `xml:"methods"`
	Methods []Method `xml:"method,omitempty"`
}

type Method struct {
	OvStruct
	Id SsoMethod `xml:"id,attr,omitempty"`
}

type MigrationBandwidths struct {
	XMLName             xml.Name             `xml:"migrationbandwidths"`
	MigrationBandwidths []MigrationBandwidth `xml:"migration_bandwidth,omitempty"`
}

type MigrationBandwidth struct {
	OvStruct
	AssignmentMethod MigrationBandwidthAssignmentMethod `xml:"assignment_method,omitempty"`
	CustomValue      *int64                             `xml:"custom_value,omitempty"`
}

type MigrationOptionss struct {
	XMLName           xml.Name           `xml:"migrationoptionss"`
	MigrationOptionss []MigrationOptions `xml:"migration_options,omitempty"`
}

type MigrationOptions struct {
	OvStruct
	AutoConverge InheritableBoolean  `xml:"auto_converge,omitempty"`
	Bandwidth    *MigrationBandwidth `xml:"bandwidth,omitempty"`
	Compressed   InheritableBoolean  `xml:"compressed,omitempty"`
	Policy       *MigrationPolicy    `xml:"policy,omitempty"`
}

type MigrationPolicys struct {
	XMLName          xml.Name          `xml:"migrationpolicys"`
	MigrationPolicys []MigrationPolicy `xml:"migration_policy,omitempty"`
}

type MigrationPolicy struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
}

type Networks struct {
	XMLName  xml.Name  `xml:"networks"`
	Networks []Network `xml:"network,omitempty"`
}

type Network struct {
	OvStruct
	Cluster                  *Cluster                  `xml:"cluster,omitempty"`
	Comment                  *string                   `xml:"comment,omitempty"`
	DataCenter               *DataCenter               `xml:"data_center,omitempty"`
	Description              *string                   `xml:"description,omitempty"`
	Display                  *bool                     `xml:"display,omitempty"`
	DnsResolverConfiguration *DnsResolverConfiguration `xml:"dns_resolver_configuration,omitempty"`
	Id                       *string                   `xml:"id,attr,omitempty"`
	Ip                       *Ip                       `xml:"ip,omitempty"`
	Mtu                      *int64                    `xml:"mtu,omitempty"`
	Name                     *string                   `xml:"name,omitempty"`
	NetworkLabels            []NetworkLabel            `xml:"network_labels,omitempty"`
	Permissions              []Permission              `xml:"permissions,omitempty"`
	ProfileRequired          *bool                     `xml:"profile_required,omitempty"`
	Qos                      *Qos                      `xml:"qos,omitempty"`
	Required                 *bool                     `xml:"required,omitempty"`
	Status                   NetworkStatus             `xml:"status,omitempty"`
	Stp                      *bool                     `xml:"stp,omitempty"`
	Usages                   []NetworkUsage            `xml:"usages,omitempty"`
	Vlan                     *Vlan                     `xml:"vlan,omitempty"`
	VnicProfiles             []VnicProfile             `xml:"vnic_profiles,omitempty"`
}

type NetworkAttachments struct {
	XMLName            xml.Name            `xml:"networkattachments"`
	NetworkAttachments []NetworkAttachment `xml:"network_attachment,omitempty"`
}

type NetworkAttachment struct {
	OvStruct
	Comment                  *string                   `xml:"comment,omitempty"`
	Description              *string                   `xml:"description,omitempty"`
	DnsResolverConfiguration *DnsResolverConfiguration `xml:"dns_resolver_configuration,omitempty"`
	Host                     *Host                     `xml:"host,omitempty"`
	HostNic                  *HostNic                  `xml:"host_nic,omitempty"`
	Id                       *string                   `xml:"id,attr,omitempty"`
	InSync                   *bool                     `xml:"in_sync,omitempty"`
	IpAddressAssignments     []IpAddressAssignment     `xml:"ip_address_assignments,omitempty"`
	Name                     *string                   `xml:"name,omitempty"`
	Network                  *Network                  `xml:"network,omitempty"`
	Properties               []Property                `xml:"properties,omitempty"`
	Qos                      *Qos                      `xml:"qos,omitempty"`
	ReportedConfigurations   []ReportedConfiguration   `xml:"reported_configurations,omitempty"`
}

type NetworkConfigurations struct {
	XMLName               xml.Name               `xml:"networkconfigurations"`
	NetworkConfigurations []NetworkConfiguration `xml:"network_configuration,omitempty"`
}

type NetworkConfiguration struct {
	OvStruct
	Dns  *Dns  `xml:"dns,omitempty"`
	Nics []Nic `xml:"nics,omitempty"`
}

type NetworkFilters struct {
	XMLName        xml.Name        `xml:"networkfilters"`
	NetworkFilters []NetworkFilter `xml:"network_filter,omitempty"`
}

type NetworkFilter struct {
	OvStruct
	Comment     *string  `xml:"comment,omitempty"`
	Description *string  `xml:"description,omitempty"`
	Id          *string  `xml:"id,attr,omitempty"`
	Name        *string  `xml:"name,omitempty"`
	Version     *Version `xml:"version,omitempty"`
}

type NetworkFilterParameters struct {
	XMLName                 xml.Name                 `xml:"networkfilterparameters"`
	NetworkFilterParameters []NetworkFilterParameter `xml:"network_filter_parameter,omitempty"`
}

type NetworkFilterParameter struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Value       *string `xml:"value,omitempty"`
}

type NetworkLabels struct {
	XMLName       xml.Name       `xml:"networklabels"`
	NetworkLabels []NetworkLabel `xml:"network_label,omitempty"`
}

type NetworkLabel struct {
	OvStruct
	Comment     *string  `xml:"comment,omitempty"`
	Description *string  `xml:"description,omitempty"`
	HostNic     *HostNic `xml:"host_nic,omitempty"`
	Id          *string  `xml:"id,attr,omitempty"`
	Name        *string  `xml:"name,omitempty"`
	Network     *Network `xml:"network,omitempty"`
}

type NfsProfileDetails struct {
	XMLName           xml.Name           `xml:"nfsprofiledetails"`
	NfsProfileDetails []NfsProfileDetail `xml:"nfs_profile_detail,omitempty"`
}

type NfsProfileDetail struct {
	OvStruct
	NfsServerIp    *string         `xml:"nfs_server_ip,omitempty"`
	ProfileDetails []ProfileDetail `xml:"profile_details,omitempty"`
}

type Nics struct {
	XMLName xml.Name `xml:"nics"`
	Nics    []Nic    `xml:"nic,omitempty"`
}

type Nic struct {
	OvStruct
	BootProtocol                   BootProtocol             `xml:"boot_protocol,omitempty"`
	Comment                        *string                  `xml:"comment,omitempty"`
	Description                    *string                  `xml:"description,omitempty"`
	Id                             *string                  `xml:"id,attr,omitempty"`
	InstanceType                   *InstanceType            `xml:"instance_type,omitempty"`
	Interface                      NicInterface             `xml:"interface,omitempty"`
	Linked                         *bool                    `xml:"linked,omitempty"`
	Mac                            *Mac                     `xml:"mac,omitempty"`
	Name                           *string                  `xml:"name,omitempty"`
	Network                        *Network                 `xml:"network,omitempty"`
	NetworkAttachments             []NetworkAttachment      `xml:"network_attachments,omitempty"`
	NetworkFilterParameters        []NetworkFilterParameter `xml:"network_filter_parameters,omitempty"`
	NetworkLabels                  []NetworkLabel           `xml:"network_labels,omitempty"`
	OnBoot                         *bool                    `xml:"on_boot,omitempty"`
	Plugged                        *bool                    `xml:"plugged,omitempty"`
	ReportedDevices                []ReportedDevice         `xml:"reported_devices,omitempty"`
	Statistics                     []Statistic              `xml:"statistics,omitempty"`
	Template                       *Template                `xml:"template,omitempty"`
	VirtualFunctionAllowedLabels   []NetworkLabel           `xml:"virtual_function_allowed_labels,omitempty"`
	VirtualFunctionAllowedNetworks []Network                `xml:"virtual_function_allowed_networks,omitempty"`
	Vm                             *Vm                      `xml:"vm,omitempty"`
	Vms                            []Vm                     `xml:"vms,omitempty"`
	VnicProfile                    *VnicProfile             `xml:"vnic_profile,omitempty"`
}

type NicConfigurations struct {
	XMLName           xml.Name           `xml:"nicconfigurations"`
	NicConfigurations []NicConfiguration `xml:"nic_configuration,omitempty"`
}

type NicConfiguration struct {
	OvStruct
	BootProtocol     BootProtocol `xml:"boot_protocol,omitempty"`
	Ip               *Ip          `xml:"ip,omitempty"`
	Ipv6             *Ip          `xml:"ipv6,omitempty"`
	Ipv6BootProtocol BootProtocol `xml:"ipv6_boot_protocol,omitempty"`
	Name             *string      `xml:"name,omitempty"`
	OnBoot           *bool        `xml:"on_boot,omitempty"`
}

type NumaNodes struct {
	XMLName   xml.Name   `xml:"numanodes"`
	NumaNodes []NumaNode `xml:"numa_node,omitempty"`
}

type NumaNode struct {
	OvStruct
	Comment      *string     `xml:"comment,omitempty"`
	Cpu          *Cpu        `xml:"cpu,omitempty"`
	Description  *string     `xml:"description,omitempty"`
	Host         *Host       `xml:"host,omitempty"`
	Id           *string     `xml:"id,attr,omitempty"`
	Index        *int64      `xml:"index,omitempty"`
	Memory       *int64      `xml:"memory,omitempty"`
	Name         *string     `xml:"name,omitempty"`
	NodeDistance *string     `xml:"node_distance,omitempty"`
	Statistics   []Statistic `xml:"statistics,omitempty"`
}

type NumaNodePins struct {
	XMLName      xml.Name      `xml:"numanodepins"`
	NumaNodePins []NumaNodePin `xml:"numa_node_pin,omitempty"`
}

type NumaNodePin struct {
	OvStruct
	HostNumaNode *NumaNode `xml:"host_numa_node,omitempty"`
	Index        *int64    `xml:"index,omitempty"`
	Pinned       *bool     `xml:"pinned,omitempty"`
}

type OpenStackImages struct {
	XMLName         xml.Name         `xml:"openstackimages"`
	OpenStackImages []OpenStackImage `xml:"open_stack_image,omitempty"`
}

type OpenStackImage struct {
	OvStruct
	Comment                *string                 `xml:"comment,omitempty"`
	Description            *string                 `xml:"description,omitempty"`
	Id                     *string                 `xml:"id,attr,omitempty"`
	Name                   *string                 `xml:"name,omitempty"`
	OpenstackImageProvider *OpenStackImageProvider `xml:"openstack_image_provider,omitempty"`
}

type OpenStackImageProviders struct {
	XMLName                 xml.Name                 `xml:"openstackimageproviders"`
	OpenStackImageProviders []OpenStackImageProvider `xml:"open_stack_image_provider,omitempty"`
}

type OpenStackImageProvider struct {
	OvStruct
	AuthenticationUrl      *string          `xml:"authentication_url,omitempty"`
	Certificates           []Certificate    `xml:"certificates,omitempty"`
	Comment                *string          `xml:"comment,omitempty"`
	Description            *string          `xml:"description,omitempty"`
	Id                     *string          `xml:"id,attr,omitempty"`
	Images                 []OpenStackImage `xml:"images,omitempty"`
	Name                   *string          `xml:"name,omitempty"`
	Password               *string          `xml:"password,omitempty"`
	Properties             []Property       `xml:"properties,omitempty"`
	RequiresAuthentication *bool            `xml:"requires_authentication,omitempty"`
	TenantName             *string          `xml:"tenant_name,omitempty"`
	Url                    *string          `xml:"url,omitempty"`
	Username               *string          `xml:"username,omitempty"`
}

type OpenStackNetworks struct {
	XMLName           xml.Name           `xml:"openstacknetworks"`
	OpenStackNetworks []OpenStackNetwork `xml:"open_stack_network,omitempty"`
}

type OpenStackNetwork struct {
	OvStruct
	Comment                  *string                   `xml:"comment,omitempty"`
	Description              *string                   `xml:"description,omitempty"`
	Id                       *string                   `xml:"id,attr,omitempty"`
	Name                     *string                   `xml:"name,omitempty"`
	OpenstackNetworkProvider *OpenStackNetworkProvider `xml:"openstack_network_provider,omitempty"`
}

type OpenStackNetworkProviders struct {
	XMLName                   xml.Name                   `xml:"openstacknetworkproviders"`
	OpenStackNetworkProviders []OpenStackNetworkProvider `xml:"open_stack_network_provider,omitempty"`
}

type OpenStackNetworkProvider struct {
	OvStruct
	AgentConfiguration     *AgentConfiguration          `xml:"agent_configuration,omitempty"`
	AuthenticationUrl      *string                      `xml:"authentication_url,omitempty"`
	Certificates           []Certificate                `xml:"certificates,omitempty"`
	Comment                *string                      `xml:"comment,omitempty"`
	Description            *string                      `xml:"description,omitempty"`
	Id                     *string                      `xml:"id,attr,omitempty"`
	Name                   *string                      `xml:"name,omitempty"`
	Networks               []OpenStackNetwork           `xml:"networks,omitempty"`
	Password               *string                      `xml:"password,omitempty"`
	PluginType             NetworkPluginType            `xml:"plugin_type,omitempty"`
	Properties             []Property                   `xml:"properties,omitempty"`
	ReadOnly               *bool                        `xml:"read_only,omitempty"`
	RequiresAuthentication *bool                        `xml:"requires_authentication,omitempty"`
	Subnets                []OpenStackSubnet            `xml:"subnets,omitempty"`
	TenantName             *string                      `xml:"tenant_name,omitempty"`
	Type                   OpenStackNetworkProviderType `xml:"type,omitempty"`
	Url                    *string                      `xml:"url,omitempty"`
	Username               *string                      `xml:"username,omitempty"`
}

type OpenStackProviders struct {
	XMLName            xml.Name            `xml:"openstackproviders"`
	OpenStackProviders []OpenStackProvider `xml:"open_stack_provider,omitempty"`
}

type OpenStackProvider struct {
	OvStruct
	AuthenticationUrl      *string    `xml:"authentication_url,omitempty"`
	Comment                *string    `xml:"comment,omitempty"`
	Description            *string    `xml:"description,omitempty"`
	Id                     *string    `xml:"id,attr,omitempty"`
	Name                   *string    `xml:"name,omitempty"`
	Password               *string    `xml:"password,omitempty"`
	Properties             []Property `xml:"properties,omitempty"`
	RequiresAuthentication *bool      `xml:"requires_authentication,omitempty"`
	TenantName             *string    `xml:"tenant_name,omitempty"`
	Url                    *string    `xml:"url,omitempty"`
	Username               *string    `xml:"username,omitempty"`
}

type OpenStackSubnets struct {
	XMLName          xml.Name          `xml:"openstacksubnets"`
	OpenStackSubnets []OpenStackSubnet `xml:"open_stack_subnet,omitempty"`
}

type OpenStackSubnet struct {
	OvStruct
	Cidr             *string           `xml:"cidr,omitempty"`
	Comment          *string           `xml:"comment,omitempty"`
	Description      *string           `xml:"description,omitempty"`
	DnsServers       []string          `xml:"dns_servers,omitempty"`
	Gateway          *string           `xml:"gateway,omitempty"`
	Id               *string           `xml:"id,attr,omitempty"`
	IpVersion        *string           `xml:"ip_version,omitempty"`
	Name             *string           `xml:"name,omitempty"`
	OpenstackNetwork *OpenStackNetwork `xml:"openstack_network,omitempty"`
}

type OpenStackVolumeProviders struct {
	XMLName                  xml.Name                  `xml:"openstackvolumeproviders"`
	OpenStackVolumeProviders []OpenStackVolumeProvider `xml:"open_stack_volume_provider,omitempty"`
}

type OpenStackVolumeProvider struct {
	OvStruct
	AuthenticationKeys     []OpenstackVolumeAuthenticationKey `xml:"authentication_keys,omitempty"`
	AuthenticationUrl      *string                            `xml:"authentication_url,omitempty"`
	Certificates           []Certificate                      `xml:"certificates,omitempty"`
	Comment                *string                            `xml:"comment,omitempty"`
	DataCenter             *DataCenter                        `xml:"data_center,omitempty"`
	Description            *string                            `xml:"description,omitempty"`
	Id                     *string                            `xml:"id,attr,omitempty"`
	Name                   *string                            `xml:"name,omitempty"`
	Password               *string                            `xml:"password,omitempty"`
	Properties             []Property                         `xml:"properties,omitempty"`
	RequiresAuthentication *bool                              `xml:"requires_authentication,omitempty"`
	TenantName             *string                            `xml:"tenant_name,omitempty"`
	Url                    *string                            `xml:"url,omitempty"`
	Username               *string                            `xml:"username,omitempty"`
	VolumeTypes            []OpenStackVolumeType              `xml:"volume_types,omitempty"`
}

type OpenStackVolumeTypes struct {
	XMLName              xml.Name              `xml:"openstackvolumetypes"`
	OpenStackVolumeTypes []OpenStackVolumeType `xml:"open_stack_volume_type,omitempty"`
}

type OpenStackVolumeType struct {
	OvStruct
	Comment                 *string                  `xml:"comment,omitempty"`
	Description             *string                  `xml:"description,omitempty"`
	Id                      *string                  `xml:"id,attr,omitempty"`
	Name                    *string                  `xml:"name,omitempty"`
	OpenstackVolumeProvider *OpenStackVolumeProvider `xml:"openstack_volume_provider,omitempty"`
	Properties              []Property               `xml:"properties,omitempty"`
}

type OpenstackVolumeAuthenticationKeys struct {
	XMLName                           xml.Name                           `xml:"openstackvolumeauthenticationkeys"`
	OpenstackVolumeAuthenticationKeys []OpenstackVolumeAuthenticationKey `xml:"openstack_volume_authentication_key,omitempty"`
}

type OpenstackVolumeAuthenticationKey struct {
	OvStruct
	Comment                 *string                                   `xml:"comment,omitempty"`
	CreationDate            time.Time                                 `xml:"creation_date,omitempty"`
	Description             *string                                   `xml:"description,omitempty"`
	Id                      *string                                   `xml:"id,attr,omitempty"`
	Name                    *string                                   `xml:"name,omitempty"`
	OpenstackVolumeProvider *OpenStackVolumeProvider                  `xml:"openstack_volume_provider,omitempty"`
	UsageType               OpenstackVolumeAuthenticationKeyUsageType `xml:"usage_type,omitempty"`
	Uuid                    *string                                   `xml:"uuid,omitempty"`
	Value                   *string                                   `xml:"value,omitempty"`
}

type OperatingSystems struct {
	XMLName          xml.Name          `xml:"operatingsystems"`
	OperatingSystems []OperatingSystem `xml:"operating_system,omitempty"`
}

type OperatingSystem struct {
	OvStruct
	Boot                  *Boot    `xml:"boot,omitempty"`
	Cmdline               *string  `xml:"cmdline,omitempty"`
	CustomKernelCmdline   *string  `xml:"custom_kernel_cmdline,omitempty"`
	Initrd                *string  `xml:"initrd,omitempty"`
	Kernel                *string  `xml:"kernel,omitempty"`
	ReportedKernelCmdline *string  `xml:"reported_kernel_cmdline,omitempty"`
	Type                  *string  `xml:"type,omitempty"`
	Version               *Version `xml:"version,omitempty"`
}

type OperatingSystemInfos struct {
	XMLName              xml.Name              `xml:"operatingsysteminfos"`
	OperatingSystemInfos []OperatingSystemInfo `xml:"operating_system_info,omitempty"`
}

type OperatingSystemInfo struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	LargeIcon   *Icon   `xml:"large_icon,omitempty"`
	Name        *string `xml:"name,omitempty"`
	SmallIcon   *Icon   `xml:"small_icon,omitempty"`
}

type Options struct {
	XMLName xml.Name `xml:"options"`
	Options []Option `xml:"option,omitempty"`
}

type Option struct {
	OvStruct
	Name  *string `xml:"name,omitempty"`
	Type  *string `xml:"type,omitempty"`
	Value *string `xml:"value,omitempty"`
}

type Packages struct {
	XMLName  xml.Name  `xml:"packages"`
	Packages []Package `xml:"package,omitempty"`
}

type Package struct {
	OvStruct
	Name *string `xml:"name,omitempty"`
}

type Payloads struct {
	XMLName  xml.Name  `xml:"payloads"`
	Payloads []Payload `xml:"payload,omitempty"`
}

type Payload struct {
	OvStruct
	Files    []File       `xml:"files,omitempty"`
	Type     VmDeviceType `xml:"type,omitempty"`
	VolumeId *string      `xml:"volume_id,omitempty"`
}

type Permissions struct {
	XMLName     xml.Name     `xml:"permissions"`
	Permissions []Permission `xml:"permission,omitempty"`
}

type Permission struct {
	OvStruct
	Cluster       *Cluster       `xml:"cluster,omitempty"`
	Comment       *string        `xml:"comment,omitempty"`
	DataCenter    *DataCenter    `xml:"data_center,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Disk          *Disk          `xml:"disk,omitempty"`
	Group         *Group         `xml:"group,omitempty"`
	Host          *Host          `xml:"host,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Role          *Role          `xml:"role,omitempty"`
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
	Template      *Template      `xml:"template,omitempty"`
	User          *User          `xml:"user,omitempty"`
	Vm            *Vm            `xml:"vm,omitempty"`
	VmPool        *VmPool        `xml:"vm_pool,omitempty"`
}

type Permits struct {
	XMLName xml.Name `xml:"permits"`
	Permits []Permit `xml:"permit,omitempty"`
}

type Permit struct {
	OvStruct
	Administrative *bool   `xml:"administrative,omitempty"`
	Comment        *string `xml:"comment,omitempty"`
	Description    *string `xml:"description,omitempty"`
	Id             *string `xml:"id,attr,omitempty"`
	Name           *string `xml:"name,omitempty"`
	Role           *Role   `xml:"role,omitempty"`
}

type PmProxys struct {
	XMLName  xml.Name  `xml:"pmproxys"`
	PmProxys []PmProxy `xml:"pm_proxy,omitempty"`
}

type PmProxy struct {
	OvStruct
	Type PmProxyType `xml:"type,omitempty"`
}

type PortMirrorings struct {
	XMLName        xml.Name        `xml:"portmirrorings"`
	PortMirrorings []PortMirroring `xml:"port_mirroring,omitempty"`
}

type PortMirroring struct {
	OvStruct
}

type PowerManagements struct {
	XMLName          xml.Name          `xml:"powermanagements"`
	PowerManagements []PowerManagement `xml:"power_management,omitempty"`
}

type PowerManagement struct {
	OvStruct
	Address            *string               `xml:"address,omitempty"`
	Agents             []Agent               `xml:"agents,omitempty"`
	AutomaticPmEnabled *bool                 `xml:"automatic_pm_enabled,omitempty"`
	Enabled            *bool                 `xml:"enabled,omitempty"`
	KdumpDetection     *bool                 `xml:"kdump_detection,omitempty"`
	Options            []Option              `xml:"options,omitempty"`
	Password           *string               `xml:"password,omitempty"`
	PmProxies          []PmProxy             `xml:"pm_proxies,omitempty"`
	Status             PowerManagementStatus `xml:"status,omitempty"`
	Type               *string               `xml:"type,omitempty"`
	Username           *string               `xml:"username,omitempty"`
}

type Products struct {
	XMLName  xml.Name  `xml:"products"`
	Products []Product `xml:"product,omitempty"`
}

type Product struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
}

type ProductInfos struct {
	XMLName      xml.Name      `xml:"productinfos"`
	ProductInfos []ProductInfo `xml:"product_info,omitempty"`
}

type ProductInfo struct {
	OvStruct
	Name    *string  `xml:"name,omitempty"`
	Vendor  *string  `xml:"vendor,omitempty"`
	Version *Version `xml:"version,omitempty"`
}

type ProfileDetails struct {
	XMLName        xml.Name        `xml:"profiledetails"`
	ProfileDetails []ProfileDetail `xml:"profile_detail,omitempty"`
}

type ProfileDetail struct {
	OvStruct
	BlockStatistics []BlockStatistic `xml:"block_statistics,omitempty"`
	Duration        *int64           `xml:"duration,omitempty"`
	FopStatistics   []FopStatistic   `xml:"fop_statistics,omitempty"`
	ProfileType     *string          `xml:"profile_type,omitempty"`
	Statistics      []Statistic      `xml:"statistics,omitempty"`
}

type Propertys struct {
	XMLName   xml.Name   `xml:"propertys"`
	Propertys []Property `xml:"property,omitempty"`
}

type Property struct {
	OvStruct
	Name  *string `xml:"name,omitempty"`
	Value *string `xml:"value,omitempty"`
}

type ProxyTickets struct {
	XMLName      xml.Name      `xml:"proxytickets"`
	ProxyTickets []ProxyTicket `xml:"proxy_ticket,omitempty"`
}

type ProxyTicket struct {
	OvStruct
	Value *string `xml:"value,omitempty"`
}

type Qoss struct {
	XMLName xml.Name `xml:"qoss"`
	Qoss    []Qos    `xml:"qos,omitempty"`
}

type Qos struct {
	OvStruct
	Comment                   *string     `xml:"comment,omitempty"`
	CpuLimit                  *int64      `xml:"cpu_limit,omitempty"`
	DataCenter                *DataCenter `xml:"data_center,omitempty"`
	Description               *string     `xml:"description,omitempty"`
	Id                        *string     `xml:"id,attr,omitempty"`
	InboundAverage            *int64      `xml:"inbound_average,omitempty"`
	InboundBurst              *int64      `xml:"inbound_burst,omitempty"`
	InboundPeak               *int64      `xml:"inbound_peak,omitempty"`
	MaxIops                   *int64      `xml:"max_iops,omitempty"`
	MaxReadIops               *int64      `xml:"max_read_iops,omitempty"`
	MaxReadThroughput         *int64      `xml:"max_read_throughput,omitempty"`
	MaxThroughput             *int64      `xml:"max_throughput,omitempty"`
	MaxWriteIops              *int64      `xml:"max_write_iops,omitempty"`
	MaxWriteThroughput        *int64      `xml:"max_write_throughput,omitempty"`
	Name                      *string     `xml:"name,omitempty"`
	OutboundAverage           *int64      `xml:"outbound_average,omitempty"`
	OutboundAverageLinkshare  *int64      `xml:"outbound_average_linkshare,omitempty"`
	OutboundAverageRealtime   *int64      `xml:"outbound_average_realtime,omitempty"`
	OutboundAverageUpperlimit *int64      `xml:"outbound_average_upperlimit,omitempty"`
	OutboundBurst             *int64      `xml:"outbound_burst,omitempty"`
	OutboundPeak              *int64      `xml:"outbound_peak,omitempty"`
	Type                      QosType     `xml:"type,omitempty"`
}

type Quotas struct {
	XMLName xml.Name `xml:"quotas"`
	Quotas  []Quota  `xml:"quota,omitempty"`
}

type Quota struct {
	OvStruct
	ClusterHardLimitPct *int64              `xml:"cluster_hard_limit_pct,omitempty"`
	ClusterSoftLimitPct *int64              `xml:"cluster_soft_limit_pct,omitempty"`
	Comment             *string             `xml:"comment,omitempty"`
	DataCenter          *DataCenter         `xml:"data_center,omitempty"`
	Description         *string             `xml:"description,omitempty"`
	Disks               []Disk              `xml:"disks,omitempty"`
	Id                  *string             `xml:"id,attr,omitempty"`
	Name                *string             `xml:"name,omitempty"`
	Permissions         []Permission        `xml:"permissions,omitempty"`
	QuotaClusterLimits  []QuotaClusterLimit `xml:"quota_cluster_limits,omitempty"`
	QuotaStorageLimits  []QuotaStorageLimit `xml:"quota_storage_limits,omitempty"`
	StorageHardLimitPct *int64              `xml:"storage_hard_limit_pct,omitempty"`
	StorageSoftLimitPct *int64              `xml:"storage_soft_limit_pct,omitempty"`
	Users               []User              `xml:"users,omitempty"`
	Vms                 []Vm                `xml:"vms,omitempty"`
}

type QuotaClusterLimits struct {
	XMLName            xml.Name            `xml:"quotaclusterlimits"`
	QuotaClusterLimits []QuotaClusterLimit `xml:"quota_cluster_limit,omitempty"`
}

type QuotaClusterLimit struct {
	OvStruct
	Cluster     *Cluster `xml:"cluster,omitempty"`
	Comment     *string  `xml:"comment,omitempty"`
	Description *string  `xml:"description,omitempty"`
	Id          *string  `xml:"id,attr,omitempty"`
	MemoryLimit *float64 `xml:"memory_limit,omitempty"`
	MemoryUsage *float64 `xml:"memory_usage,omitempty"`
	Name        *string  `xml:"name,omitempty"`
	Quota       *Quota   `xml:"quota,omitempty"`
	VcpuLimit   *int64   `xml:"vcpu_limit,omitempty"`
	VcpuUsage   *int64   `xml:"vcpu_usage,omitempty"`
}

type QuotaStorageLimits struct {
	XMLName            xml.Name            `xml:"quotastoragelimits"`
	QuotaStorageLimits []QuotaStorageLimit `xml:"quota_storage_limit,omitempty"`
}

type QuotaStorageLimit struct {
	OvStruct
	Comment       *string        `xml:"comment,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Limit         *int64         `xml:"limit,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Quota         *Quota         `xml:"quota,omitempty"`
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
	Usage         *float64       `xml:"usage,omitempty"`
}

type Ranges struct {
	XMLName xml.Name `xml:"ranges"`
	Ranges  []Range  `xml:"range,omitempty"`
}

type Range struct {
	OvStruct
	From *string `xml:"from,omitempty"`
	To   *string `xml:"to,omitempty"`
}

type Rates struct {
	XMLName xml.Name `xml:"rates"`
	Rates   []Rate   `xml:"rate,omitempty"`
}

type Rate struct {
	OvStruct
	Bytes  *int64 `xml:"bytes,omitempty"`
	Period *int64 `xml:"period,omitempty"`
}

type ReportedConfigurations struct {
	XMLName                xml.Name                `xml:"reportedconfigurations"`
	ReportedConfigurations []ReportedConfiguration `xml:"reported_configuration,omitempty"`
}

type ReportedConfiguration struct {
	OvStruct
	ActualValue   *string `xml:"actual_value,omitempty"`
	ExpectedValue *string `xml:"expected_value,omitempty"`
	InSync        *bool   `xml:"in_sync,omitempty"`
	Name          *string `xml:"name,omitempty"`
}

type ReportedDevices struct {
	XMLName         xml.Name         `xml:"reporteddevices"`
	ReportedDevices []ReportedDevice `xml:"reported_device,omitempty"`
}

type ReportedDevice struct {
	OvStruct
	Comment     *string            `xml:"comment,omitempty"`
	Description *string            `xml:"description,omitempty"`
	Id          *string            `xml:"id,attr,omitempty"`
	Ips         []Ip               `xml:"ips,omitempty"`
	Mac         *Mac               `xml:"mac,omitempty"`
	Name        *string            `xml:"name,omitempty"`
	Type        ReportedDeviceType `xml:"type,omitempty"`
	Vm          *Vm                `xml:"vm,omitempty"`
}

type RngDevices struct {
	XMLName    xml.Name    `xml:"rngdevices"`
	RngDevices []RngDevice `xml:"rng_device,omitempty"`
}

type RngDevice struct {
	OvStruct
	Rate   *Rate     `xml:"rate,omitempty"`
	Source RngSource `xml:"source,omitempty"`
}

type Roles struct {
	XMLName xml.Name `xml:"roles"`
	Roles   []Role   `xml:"role,omitempty"`
}

type Role struct {
	OvStruct
	Administrative *bool    `xml:"administrative,omitempty"`
	Comment        *string  `xml:"comment,omitempty"`
	Description    *string  `xml:"description,omitempty"`
	Id             *string  `xml:"id,attr,omitempty"`
	Mutable        *bool    `xml:"mutable,omitempty"`
	Name           *string  `xml:"name,omitempty"`
	Permits        []Permit `xml:"permits,omitempty"`
	User           *User    `xml:"user,omitempty"`
}

type SchedulingPolicys struct {
	XMLName           xml.Name           `xml:"schedulingpolicys"`
	SchedulingPolicys []SchedulingPolicy `xml:"scheduling_policy,omitempty"`
}

type SchedulingPolicy struct {
	OvStruct
	Balances      []Balance  `xml:"balances,omitempty"`
	Comment       *string    `xml:"comment,omitempty"`
	DefaultPolicy *bool      `xml:"default_policy,omitempty"`
	Description   *string    `xml:"description,omitempty"`
	Filters       []Filter   `xml:"filters,omitempty"`
	Id            *string    `xml:"id,attr,omitempty"`
	Locked        *bool      `xml:"locked,omitempty"`
	Name          *string    `xml:"name,omitempty"`
	Properties    []Property `xml:"properties,omitempty"`
	Weight        []Weight   `xml:"weight,omitempty"`
}

type SchedulingPolicyUnits struct {
	XMLName               xml.Name               `xml:"schedulingpolicyunits"`
	SchedulingPolicyUnits []SchedulingPolicyUnit `xml:"scheduling_policy_unit,omitempty"`
}

type SchedulingPolicyUnit struct {
	OvStruct
	Comment     *string        `xml:"comment,omitempty"`
	Description *string        `xml:"description,omitempty"`
	Enabled     *bool          `xml:"enabled,omitempty"`
	Id          *string        `xml:"id,attr,omitempty"`
	Internal    *bool          `xml:"internal,omitempty"`
	Name        *string        `xml:"name,omitempty"`
	Properties  []Property     `xml:"properties,omitempty"`
	Type        PolicyUnitType `xml:"type,omitempty"`
}

type SeLinuxs struct {
	XMLName  xml.Name  `xml:"selinuxs"`
	SeLinuxs []SeLinux `xml:"se_linux,omitempty"`
}

type SeLinux struct {
	OvStruct
	Mode SeLinuxMode `xml:"mode,omitempty"`
}

type SerialNumbers struct {
	XMLName       xml.Name       `xml:"serialnumbers"`
	SerialNumbers []SerialNumber `xml:"serial_number,omitempty"`
}

type SerialNumber struct {
	OvStruct
	Policy SerialNumberPolicy `xml:"policy,omitempty"`
	Value  *string            `xml:"value,omitempty"`
}

type Sessions struct {
	XMLName  xml.Name  `xml:"sessions"`
	Sessions []Session `xml:"session,omitempty"`
}

type Session struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	ConsoleUser *bool   `xml:"console_user,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Ip          *Ip     `xml:"ip,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Protocol    *string `xml:"protocol,omitempty"`
	User        *User   `xml:"user,omitempty"`
	Vm          *Vm     `xml:"vm,omitempty"`
}

type SkipIfConnectivityBrokens struct {
	XMLName                   xml.Name                   `xml:"skipifconnectivitybrokens"`
	SkipIfConnectivityBrokens []SkipIfConnectivityBroken `xml:"skip_if_connectivity_broken,omitempty"`
}

type SkipIfConnectivityBroken struct {
	OvStruct
	Enabled   *bool  `xml:"enabled,omitempty"`
	Threshold *int64 `xml:"threshold,omitempty"`
}

type SkipIfSdActives struct {
	XMLName         xml.Name         `xml:"skipifsdactives"`
	SkipIfSdActives []SkipIfSdActive `xml:"skip_if_sd_active,omitempty"`
}

type SkipIfSdActive struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
}

type Snapshots struct {
	XMLName   xml.Name   `xml:"snapshots"`
	Snapshots []Snapshot `xml:"snapshot,omitempty"`
}

type Snapshot struct {
	OvStruct
	AffinityLabels             []AffinityLabel       `xml:"affinity_labels,omitempty"`
	Applications               []Application         `xml:"applications,omitempty"`
	Bios                       *Bios                 `xml:"bios,omitempty"`
	Cdroms                     []Cdrom               `xml:"cdroms,omitempty"`
	Cluster                    *Cluster              `xml:"cluster,omitempty"`
	Comment                    *string               `xml:"comment,omitempty"`
	Console                    *Console              `xml:"console,omitempty"`
	Cpu                        *Cpu                  `xml:"cpu,omitempty"`
	CpuProfile                 *CpuProfile           `xml:"cpu_profile,omitempty"`
	CpuShares                  *int64                `xml:"cpu_shares,omitempty"`
	CreationTime               time.Time             `xml:"creation_time,omitempty"`
	CustomCompatibilityVersion *Version              `xml:"custom_compatibility_version,omitempty"`
	CustomCpuModel             *string               `xml:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      *string               `xml:"custom_emulated_machine,omitempty"`
	CustomProperties           []CustomProperty      `xml:"custom_properties,omitempty"`
	Date                       time.Time             `xml:"date,omitempty"`
	DeleteProtected            *bool                 `xml:"delete_protected,omitempty"`
	Description                *string               `xml:"description,omitempty"`
	DiskAttachments            []DiskAttachment      `xml:"disk_attachments,omitempty"`
	Display                    *Display              `xml:"display,omitempty"`
	Domain                     *Domain               `xml:"domain,omitempty"`
	ExternalHostProvider       *ExternalHostProvider `xml:"external_host_provider,omitempty"`
	Floppies                   []Floppy              `xml:"floppies,omitempty"`
	Fqdn                       *string               `xml:"fqdn,omitempty"`
	GraphicsConsoles           []GraphicsConsole     `xml:"graphics_consoles,omitempty"`
	GuestOperatingSystem       *GuestOperatingSystem `xml:"guest_operating_system,omitempty"`
	GuestTimeZone              *TimeZone             `xml:"guest_time_zone,omitempty"`
	HighAvailability           *HighAvailability     `xml:"high_availability,omitempty"`
	Host                       *Host                 `xml:"host,omitempty"`
	HostDevices                []HostDevice          `xml:"host_devices,omitempty"`
	Id                         *string               `xml:"id,attr,omitempty"`
	Initialization             *Initialization       `xml:"initialization,omitempty"`
	InstanceType               *InstanceType         `xml:"instance_type,omitempty"`
	Io                         *Io                   `xml:"io,omitempty"`
	KatelloErrata              []KatelloErratum      `xml:"katello_errata,omitempty"`
	LargeIcon                  *Icon                 `xml:"large_icon,omitempty"`
	Lease                      *StorageDomainLease   `xml:"lease,omitempty"`
	Memory                     *int64                `xml:"memory,omitempty"`
	MemoryPolicy               *MemoryPolicy         `xml:"memory_policy,omitempty"`
	Migration                  *MigrationOptions     `xml:"migration,omitempty"`
	MigrationDowntime          *int64                `xml:"migration_downtime,omitempty"`
	Name                       *string               `xml:"name,omitempty"`
	NextRunConfigurationExists *bool                 `xml:"next_run_configuration_exists,omitempty"`
	Nics                       []Nic                 `xml:"nics,omitempty"`
	NumaNodes                  []NumaNode            `xml:"numa_nodes,omitempty"`
	NumaTuneMode               NumaTuneMode          `xml:"numa_tune_mode,omitempty"`
	Origin                     *string               `xml:"origin,omitempty"`
	OriginalTemplate           *Template             `xml:"original_template,omitempty"`
	Os                         *OperatingSystem      `xml:"os,omitempty"`
	Payloads                   []Payload             `xml:"payloads,omitempty"`
	Permissions                []Permission          `xml:"permissions,omitempty"`
	PersistMemorystate         *bool                 `xml:"persist_memorystate,omitempty"`
	PlacementPolicy            *VmPlacementPolicy    `xml:"placement_policy,omitempty"`
	Quota                      *Quota                `xml:"quota,omitempty"`
	ReportedDevices            []ReportedDevice      `xml:"reported_devices,omitempty"`
	RngDevice                  *RngDevice            `xml:"rng_device,omitempty"`
	RunOnce                    *bool                 `xml:"run_once,omitempty"`
	SerialNumber               *SerialNumber         `xml:"serial_number,omitempty"`
	Sessions                   []Session             `xml:"sessions,omitempty"`
	SmallIcon                  *Icon                 `xml:"small_icon,omitempty"`
	SnapshotStatus             SnapshotStatus        `xml:"snapshot_status,omitempty"`
	SnapshotType               SnapshotType          `xml:"snapshot_type,omitempty"`
	Snapshots                  []Snapshot            `xml:"snapshots,omitempty"`
	SoundcardEnabled           *bool                 `xml:"soundcard_enabled,omitempty"`
	Sso                        *Sso                  `xml:"sso,omitempty"`
	StartPaused                *bool                 `xml:"start_paused,omitempty"`
	StartTime                  time.Time             `xml:"start_time,omitempty"`
	Stateless                  *bool                 `xml:"stateless,omitempty"`
	Statistics                 []Statistic           `xml:"statistics,omitempty"`
	Status                     VmStatus              `xml:"status,omitempty"`
	StatusDetail               *string               `xml:"status_detail,omitempty"`
	StopReason                 *string               `xml:"stop_reason,omitempty"`
	StopTime                   time.Time             `xml:"stop_time,omitempty"`
	StorageDomain              *StorageDomain        `xml:"storage_domain,omitempty"`
	Tags                       []Tag                 `xml:"tags,omitempty"`
	Template                   *Template             `xml:"template,omitempty"`
	TimeZone                   *TimeZone             `xml:"time_zone,omitempty"`
	TunnelMigration            *bool                 `xml:"tunnel_migration,omitempty"`
	Type                       VmType                `xml:"type,omitempty"`
	Usb                        *Usb                  `xml:"usb,omitempty"`
	UseLatestTemplateVersion   *bool                 `xml:"use_latest_template_version,omitempty"`
	VirtioScsi                 *VirtioScsi           `xml:"virtio_scsi,omitempty"`
	Vm                         *Vm                   `xml:"vm,omitempty"`
	VmPool                     *VmPool               `xml:"vm_pool,omitempty"`
	Watchdogs                  []Watchdog            `xml:"watchdogs,omitempty"`
}

type SpecialObjectss struct {
	XMLName         xml.Name         `xml:"specialobjectss"`
	SpecialObjectss []SpecialObjects `xml:"special_objects,omitempty"`
}

type SpecialObjects struct {
	OvStruct
	BlankTemplate *Template `xml:"blank_template,omitempty"`
	RootTag       *Tag      `xml:"root_tag,omitempty"`
}

type Spms struct {
	XMLName xml.Name `xml:"spms"`
	Spms    []Spm    `xml:"spm,omitempty"`
}

type Spm struct {
	OvStruct
	Priority *int64    `xml:"priority,omitempty"`
	Status   SpmStatus `xml:"status,omitempty"`
}

type Sshs struct {
	XMLName xml.Name `xml:"sshs"`
	Sshs    []Ssh    `xml:"ssh,omitempty"`
}

type Ssh struct {
	OvStruct
	AuthenticationMethod SshAuthenticationMethod `xml:"authentication_method,omitempty"`
	Comment              *string                 `xml:"comment,omitempty"`
	Description          *string                 `xml:"description,omitempty"`
	Fingerprint          *string                 `xml:"fingerprint,omitempty"`
	Id                   *string                 `xml:"id,attr,omitempty"`
	Name                 *string                 `xml:"name,omitempty"`
	Port                 *int64                  `xml:"port,omitempty"`
	User                 *User                   `xml:"user,omitempty"`
}

type SshPublicKeys struct {
	XMLName       xml.Name       `xml:"sshpublickeys"`
	SshPublicKeys []SshPublicKey `xml:"ssh_public_key,omitempty"`
}

type SshPublicKey struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Content     *string `xml:"content,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	User        *User   `xml:"user,omitempty"`
}

type Ssos struct {
	XMLName xml.Name `xml:"ssos"`
	Ssos    []Sso    `xml:"sso,omitempty"`
}

type Sso struct {
	OvStruct
	Methods []Method `xml:"methods,omitempty"`
}

type Statistics struct {
	XMLName    xml.Name    `xml:"statistics"`
	Statistics []Statistic `xml:"statistic,omitempty"`
}

type Statistic struct {
	OvStruct
	Brick         *GlusterBrick  `xml:"brick,omitempty"`
	Comment       *string        `xml:"comment,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Disk          *Disk          `xml:"disk,omitempty"`
	GlusterVolume *GlusterVolume `xml:"gluster_volume,omitempty"`
	Host          *Host          `xml:"host,omitempty"`
	HostNic       *HostNic       `xml:"host_nic,omitempty"`
	HostNumaNode  *NumaNode      `xml:"host_numa_node,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	Kind          StatisticKind  `xml:"kind,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Nic           *Nic           `xml:"nic,omitempty"`
	Step          *Step          `xml:"step,omitempty"`
	Type          ValueType      `xml:"type,omitempty"`
	Unit          StatisticUnit  `xml:"unit,omitempty"`
	Values        []Value        `xml:"values,omitempty"`
	Vm            *Vm            `xml:"vm,omitempty"`
}

type Steps struct {
	XMLName xml.Name `xml:"steps"`
	Steps   []Step   `xml:"step,omitempty"`
}

type Step struct {
	OvStruct
	Comment       *string            `xml:"comment,omitempty"`
	Description   *string            `xml:"description,omitempty"`
	EndTime       time.Time          `xml:"end_time,omitempty"`
	ExecutionHost *Host              `xml:"execution_host,omitempty"`
	External      *bool              `xml:"external,omitempty"`
	ExternalType  ExternalSystemType `xml:"external_type,omitempty"`
	Id            *string            `xml:"id,attr,omitempty"`
	Job           *Job               `xml:"job,omitempty"`
	Name          *string            `xml:"name,omitempty"`
	Number        *int64             `xml:"number,omitempty"`
	ParentStep    *Step              `xml:"parent_step,omitempty"`
	Progress      *int64             `xml:"progress,omitempty"`
	StartTime     time.Time          `xml:"start_time,omitempty"`
	Statistics    []Statistic        `xml:"statistics,omitempty"`
	Status        StepStatus         `xml:"status,omitempty"`
	Type          StepEnum           `xml:"type,omitempty"`
}

type StorageConnections struct {
	XMLName            xml.Name            `xml:"storageconnections"`
	StorageConnections []StorageConnection `xml:"storage_connection,omitempty"`
}

type StorageConnection struct {
	OvStruct
	Address      *string     `xml:"address,omitempty"`
	Comment      *string     `xml:"comment,omitempty"`
	Description  *string     `xml:"description,omitempty"`
	Host         *Host       `xml:"host,omitempty"`
	Id           *string     `xml:"id,attr,omitempty"`
	MountOptions *string     `xml:"mount_options,omitempty"`
	Name         *string     `xml:"name,omitempty"`
	NfsRetrans   *int64      `xml:"nfs_retrans,omitempty"`
	NfsTimeo     *int64      `xml:"nfs_timeo,omitempty"`
	NfsVersion   NfsVersion  `xml:"nfs_version,omitempty"`
	Password     *string     `xml:"password,omitempty"`
	Path         *string     `xml:"path,omitempty"`
	Port         *int64      `xml:"port,omitempty"`
	Portal       *string     `xml:"portal,omitempty"`
	Target       *string     `xml:"target,omitempty"`
	Type         StorageType `xml:"type,omitempty"`
	Username     *string     `xml:"username,omitempty"`
	VfsType      *string     `xml:"vfs_type,omitempty"`
}

type StorageConnectionExtensions struct {
	XMLName                     xml.Name                     `xml:"storageconnectionextensions"`
	StorageConnectionExtensions []StorageConnectionExtension `xml:"storage_connection_extension,omitempty"`
}

type StorageConnectionExtension struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Host        *Host   `xml:"host,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Password    *string `xml:"password,omitempty"`
	Target      *string `xml:"target,omitempty"`
	Username    *string `xml:"username,omitempty"`
}

type StorageDomains struct {
	XMLName        xml.Name        `xml:"storagedomains"`
	StorageDomains []StorageDomain `xml:"storage_domain,omitempty"`
}

type StorageDomain struct {
	OvStruct
	Available                  *int64              `xml:"available,omitempty"`
	Comment                    *string             `xml:"comment,omitempty"`
	Committed                  *int64              `xml:"committed,omitempty"`
	CriticalSpaceActionBlocker *int64              `xml:"critical_space_action_blocker,omitempty"`
	DataCenter                 *DataCenter         `xml:"data_center,omitempty"`
	DataCenters                []DataCenter        `xml:"data_centers,omitempty"`
	Description                *string             `xml:"description,omitempty"`
	DiscardAfterDelete         *bool               `xml:"discard_after_delete,omitempty"`
	DiskProfiles               []DiskProfile       `xml:"disk_profiles,omitempty"`
	DiskSnapshots              []DiskSnapshot      `xml:"disk_snapshots,omitempty"`
	Disks                      []Disk              `xml:"disks,omitempty"`
	ExternalStatus             ExternalStatus      `xml:"external_status,omitempty"`
	Files                      []File              `xml:"files,omitempty"`
	Host                       *Host               `xml:"host,omitempty"`
	Id                         *string             `xml:"id,attr,omitempty"`
	Images                     []Image             `xml:"images,omitempty"`
	Import                     *bool               `xml:"import,omitempty"`
	Master                     *bool               `xml:"master,omitempty"`
	Name                       *string             `xml:"name,omitempty"`
	Permissions                []Permission        `xml:"permissions,omitempty"`
	Status                     StorageDomainStatus `xml:"status,omitempty"`
	Storage                    *HostStorage        `xml:"storage,omitempty"`
	StorageConnections         []StorageConnection `xml:"storage_connections,omitempty"`
	StorageFormat              StorageFormat       `xml:"storage_format,omitempty"`
	SupportsDiscard            *bool               `xml:"supports_discard,omitempty"`
	SupportsDiscardZeroesData  *bool               `xml:"supports_discard_zeroes_data,omitempty"`
	Templates                  []Template          `xml:"templates,omitempty"`
	Type                       StorageDomainType   `xml:"type,omitempty"`
	Used                       *int64              `xml:"used,omitempty"`
	Vms                        []Vm                `xml:"vms,omitempty"`
	WarningLowSpaceIndicator   *int64              `xml:"warning_low_space_indicator,omitempty"`
	WipeAfterDelete            *bool               `xml:"wipe_after_delete,omitempty"`
}

type StorageDomainLeases struct {
	XMLName             xml.Name             `xml:"storagedomainleases"`
	StorageDomainLeases []StorageDomainLease `xml:"storage_domain_lease,omitempty"`
}

type StorageDomainLease struct {
	OvStruct
	StorageDomain *StorageDomain `xml:"storage_domain,omitempty"`
}

type Tags struct {
	XMLName xml.Name `xml:"tags"`
	Tags    []Tag    `xml:"tag,omitempty"`
}

type Tag struct {
	OvStruct
	Comment     *string   `xml:"comment,omitempty"`
	Description *string   `xml:"description,omitempty"`
	Group       *Group    `xml:"group,omitempty"`
	Host        *Host     `xml:"host,omitempty"`
	Id          *string   `xml:"id,attr,omitempty"`
	Name        *string   `xml:"name,omitempty"`
	Parent      *Tag      `xml:"parent,omitempty"`
	Template    *Template `xml:"template,omitempty"`
	User        *User     `xml:"user,omitempty"`
	Vm          *Vm       `xml:"vm,omitempty"`
}

type Templates struct {
	XMLName   xml.Name   `xml:"templates"`
	Templates []Template `xml:"template,omitempty"`
}

type Template struct {
	OvStruct
	Bios                       *Bios               `xml:"bios,omitempty"`
	Cdroms                     []Cdrom             `xml:"cdroms,omitempty"`
	Cluster                    *Cluster            `xml:"cluster,omitempty"`
	Comment                    *string             `xml:"comment,omitempty"`
	Console                    *Console            `xml:"console,omitempty"`
	Cpu                        *Cpu                `xml:"cpu,omitempty"`
	CpuProfile                 *CpuProfile         `xml:"cpu_profile,omitempty"`
	CpuShares                  *int64              `xml:"cpu_shares,omitempty"`
	CreationTime               time.Time           `xml:"creation_time,omitempty"`
	CustomCompatibilityVersion *Version            `xml:"custom_compatibility_version,omitempty"`
	CustomCpuModel             *string             `xml:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      *string             `xml:"custom_emulated_machine,omitempty"`
	CustomProperties           []CustomProperty    `xml:"custom_properties,omitempty"`
	DeleteProtected            *bool               `xml:"delete_protected,omitempty"`
	Description                *string             `xml:"description,omitempty"`
	DiskAttachments            []DiskAttachment    `xml:"disk_attachments,omitempty"`
	Display                    *Display            `xml:"display,omitempty"`
	Domain                     *Domain             `xml:"domain,omitempty"`
	GraphicsConsoles           []GraphicsConsole   `xml:"graphics_consoles,omitempty"`
	HighAvailability           *HighAvailability   `xml:"high_availability,omitempty"`
	Id                         *string             `xml:"id,attr,omitempty"`
	Initialization             *Initialization     `xml:"initialization,omitempty"`
	Io                         *Io                 `xml:"io,omitempty"`
	LargeIcon                  *Icon               `xml:"large_icon,omitempty"`
	Lease                      *StorageDomainLease `xml:"lease,omitempty"`
	Memory                     *int64              `xml:"memory,omitempty"`
	MemoryPolicy               *MemoryPolicy       `xml:"memory_policy,omitempty"`
	Migration                  *MigrationOptions   `xml:"migration,omitempty"`
	MigrationDowntime          *int64              `xml:"migration_downtime,omitempty"`
	Name                       *string             `xml:"name,omitempty"`
	Nics                       []Nic               `xml:"nics,omitempty"`
	Origin                     *string             `xml:"origin,omitempty"`
	Os                         *OperatingSystem    `xml:"os,omitempty"`
	Permissions                []Permission        `xml:"permissions,omitempty"`
	Quota                      *Quota              `xml:"quota,omitempty"`
	RngDevice                  *RngDevice          `xml:"rng_device,omitempty"`
	SerialNumber               *SerialNumber       `xml:"serial_number,omitempty"`
	SmallIcon                  *Icon               `xml:"small_icon,omitempty"`
	SoundcardEnabled           *bool               `xml:"soundcard_enabled,omitempty"`
	Sso                        *Sso                `xml:"sso,omitempty"`
	StartPaused                *bool               `xml:"start_paused,omitempty"`
	Stateless                  *bool               `xml:"stateless,omitempty"`
	Status                     TemplateStatus      `xml:"status,omitempty"`
	StorageDomain              *StorageDomain      `xml:"storage_domain,omitempty"`
	Tags                       []Tag               `xml:"tags,omitempty"`
	TimeZone                   *TimeZone           `xml:"time_zone,omitempty"`
	TunnelMigration            *bool               `xml:"tunnel_migration,omitempty"`
	Type                       VmType              `xml:"type,omitempty"`
	Usb                        *Usb                `xml:"usb,omitempty"`
	Version                    *TemplateVersion    `xml:"version,omitempty"`
	VirtioScsi                 *VirtioScsi         `xml:"virtio_scsi,omitempty"`
	Vm                         *Vm                 `xml:"vm,omitempty"`
	Watchdogs                  []Watchdog          `xml:"watchdogs,omitempty"`
}

type TemplateVersions struct {
	XMLName          xml.Name          `xml:"templateversions"`
	TemplateVersions []TemplateVersion `xml:"template_version,omitempty"`
}

type TemplateVersion struct {
	OvStruct
	BaseTemplate  *Template `xml:"base_template,omitempty"`
	VersionName   *string   `xml:"version_name,omitempty"`
	VersionNumber *int64    `xml:"version_number,omitempty"`
}

type Tickets struct {
	XMLName xml.Name `xml:"tickets"`
	Tickets []Ticket `xml:"ticket,omitempty"`
}

type Ticket struct {
	OvStruct
	Expiry *int64  `xml:"expiry,omitempty"`
	Value  *string `xml:"value,omitempty"`
}

type TimeZones struct {
	XMLName   xml.Name   `xml:"timezones"`
	TimeZones []TimeZone `xml:"time_zone,omitempty"`
}

type TimeZone struct {
	OvStruct
	Name      *string `xml:"name,omitempty"`
	UtcOffset *string `xml:"utc_offset,omitempty"`
}

type TransparentHugePagess struct {
	XMLName               xml.Name               `xml:"transparenthugepagess"`
	TransparentHugePagess []TransparentHugePages `xml:"transparent_huge_pages,omitempty"`
}

type TransparentHugePages struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
}

type UnmanagedNetworks struct {
	XMLName           xml.Name           `xml:"unmanagednetworks"`
	UnmanagedNetworks []UnmanagedNetwork `xml:"unmanaged_network,omitempty"`
}

type UnmanagedNetwork struct {
	OvStruct
	Comment     *string  `xml:"comment,omitempty"`
	Description *string  `xml:"description,omitempty"`
	Host        *Host    `xml:"host,omitempty"`
	HostNic     *HostNic `xml:"host_nic,omitempty"`
	Id          *string  `xml:"id,attr,omitempty"`
	Name        *string  `xml:"name,omitempty"`
}

type Usbs struct {
	XMLName xml.Name `xml:"usbs"`
	Usbs    []Usb    `xml:"usb,omitempty"`
}

type Usb struct {
	OvStruct
	Enabled *bool   `xml:"enabled,omitempty"`
	Type    UsbType `xml:"type,omitempty"`
}

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user,omitempty"`
}

type User struct {
	OvStruct
	Comment       *string        `xml:"comment,omitempty"`
	Department    *string        `xml:"department,omitempty"`
	Description   *string        `xml:"description,omitempty"`
	Domain        *Domain        `xml:"domain,omitempty"`
	DomainEntryId *string        `xml:"domain_entry_id,omitempty"`
	Email         *string        `xml:"email,omitempty"`
	Groups        []Group        `xml:"groups,omitempty"`
	Id            *string        `xml:"id,attr,omitempty"`
	LastName      *string        `xml:"last_name,omitempty"`
	LoggedIn      *bool          `xml:"logged_in,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Namespace     *string        `xml:"namespace,omitempty"`
	Password      *string        `xml:"password,omitempty"`
	Permissions   []Permission   `xml:"permissions,omitempty"`
	Principal     *string        `xml:"principal,omitempty"`
	Roles         []Role         `xml:"roles,omitempty"`
	SshPublicKeys []SshPublicKey `xml:"ssh_public_keys,omitempty"`
	Tags          []Tag          `xml:"tags,omitempty"`
	UserName      *string        `xml:"user_name,omitempty"`
}

type Values struct {
	XMLName xml.Name `xml:"values"`
	Values  []Value  `xml:"value,omitempty"`
}

type Value struct {
	OvStruct
	Datum  *float64 `xml:"datum,omitempty"`
	Detail *string  `xml:"detail,omitempty"`
}

type VcpuPins struct {
	XMLName  xml.Name  `xml:"vcpupins"`
	VcpuPins []VcpuPin `xml:"vcpu_pin,omitempty"`
}

type VcpuPin struct {
	OvStruct
	CpuSet *string `xml:"cpu_set,omitempty"`
	Vcpu   *int64  `xml:"vcpu,omitempty"`
}

type Vendors struct {
	XMLName xml.Name `xml:"vendors"`
	Vendors []Vendor `xml:"vendor,omitempty"`
}

type Vendor struct {
	OvStruct
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Name        *string `xml:"name,omitempty"`
}

type Versions struct {
	XMLName  xml.Name  `xml:"versions"`
	Versions []Version `xml:"version,omitempty"`
}

type Version struct {
	OvStruct
	Build_      *int64  `xml:"build,omitempty"`
	Comment     *string `xml:"comment,omitempty"`
	Description *string `xml:"description,omitempty"`
	FullVersion *string `xml:"full_version,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Major       *int64  `xml:"major,omitempty"`
	Minor       *int64  `xml:"minor,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Revision    *int64  `xml:"revision,omitempty"`
}

type VirtioScsis struct {
	XMLName     xml.Name     `xml:"virtioscsis"`
	VirtioScsis []VirtioScsi `xml:"virtio_scsi,omitempty"`
}

type VirtioScsi struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
}

type VirtualNumaNodes struct {
	XMLName          xml.Name          `xml:"virtualnumanodes"`
	VirtualNumaNodes []VirtualNumaNode `xml:"virtual_numa_node,omitempty"`
}

type VirtualNumaNode struct {
	OvStruct
	Comment      *string       `xml:"comment,omitempty"`
	Cpu          *Cpu          `xml:"cpu,omitempty"`
	Description  *string       `xml:"description,omitempty"`
	Host         *Host         `xml:"host,omitempty"`
	Id           *string       `xml:"id,attr,omitempty"`
	Index        *int64        `xml:"index,omitempty"`
	Memory       *int64        `xml:"memory,omitempty"`
	Name         *string       `xml:"name,omitempty"`
	NodeDistance *string       `xml:"node_distance,omitempty"`
	NumaNodePins []NumaNodePin `xml:"numa_node_pins,omitempty"`
	Statistics   []Statistic   `xml:"statistics,omitempty"`
	Vm           *Vm           `xml:"vm,omitempty"`
}

type Vlans struct {
	XMLName xml.Name `xml:"vlans"`
	Vlans   []Vlan   `xml:"vlan,omitempty"`
}

type Vlan struct {
	OvStruct
	Id *int64 `xml:"id,attr,omitempty"`
}

type Vms struct {
	XMLName xml.Name `xml:"vms"`
	Vms     []Vm     `xml:"vm,omitempty"`
}

type Vm struct {
	OvStruct
	AffinityLabels             []AffinityLabel       `xml:"affinity_labels,omitempty"`
	Applications               []Application         `xml:"applications,omitempty"`
	Bios                       *Bios                 `xml:"bios,omitempty"`
	Cdroms                     []Cdrom               `xml:"cdroms,omitempty"`
	Cluster                    *Cluster              `xml:"cluster,omitempty"`
	Comment                    *string               `xml:"comment,omitempty"`
	Console                    *Console              `xml:"console,omitempty"`
	Cpu                        *Cpu                  `xml:"cpu,omitempty"`
	CpuProfile                 *CpuProfile           `xml:"cpu_profile,omitempty"`
	CpuShares                  *int64                `xml:"cpu_shares,omitempty"`
	CreationTime               time.Time             `xml:"creation_time,omitempty"`
	CustomCompatibilityVersion *Version              `xml:"custom_compatibility_version,omitempty"`
	CustomCpuModel             *string               `xml:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      *string               `xml:"custom_emulated_machine,omitempty"`
	CustomProperties           []CustomProperty      `xml:"custom_properties,omitempty"`
	DeleteProtected            *bool                 `xml:"delete_protected,omitempty"`
	Description                *string               `xml:"description,omitempty"`
	DiskAttachments            []DiskAttachment      `xml:"disk_attachments,omitempty"`
	Display                    *Display              `xml:"display,omitempty"`
	Domain                     *Domain               `xml:"domain,omitempty"`
	ExternalHostProvider       *ExternalHostProvider `xml:"external_host_provider,omitempty"`
	Floppies                   []Floppy              `xml:"floppies,omitempty"`
	Fqdn                       *string               `xml:"fqdn,omitempty"`
	GraphicsConsoles           []GraphicsConsole     `xml:"graphics_consoles,omitempty"`
	GuestOperatingSystem       *GuestOperatingSystem `xml:"guest_operating_system,omitempty"`
	GuestTimeZone              *TimeZone             `xml:"guest_time_zone,omitempty"`
	HighAvailability           *HighAvailability     `xml:"high_availability,omitempty"`
	Host                       *Host                 `xml:"host,omitempty"`
	HostDevices                []HostDevice          `xml:"host_devices,omitempty"`
	Id                         *string               `xml:"id,attr,omitempty"`
	Initialization             *Initialization       `xml:"initialization,omitempty"`
	InstanceType               *InstanceType         `xml:"instance_type,omitempty"`
	Io                         *Io                   `xml:"io,omitempty"`
	KatelloErrata              []KatelloErratum      `xml:"katello_errata,omitempty"`
	LargeIcon                  *Icon                 `xml:"large_icon,omitempty"`
	Lease                      *StorageDomainLease   `xml:"lease,omitempty"`
	Memory                     *int64                `xml:"memory,omitempty"`
	MemoryPolicy               *MemoryPolicy         `xml:"memory_policy,omitempty"`
	Migration                  *MigrationOptions     `xml:"migration,omitempty"`
	MigrationDowntime          *int64                `xml:"migration_downtime,omitempty"`
	Name                       *string               `xml:"name,omitempty"`
	NextRunConfigurationExists *bool                 `xml:"next_run_configuration_exists,omitempty"`
	Nics                       []Nic                 `xml:"nics,omitempty"`
	NumaNodes                  []NumaNode            `xml:"numa_nodes,omitempty"`
	NumaTuneMode               NumaTuneMode          `xml:"numa_tune_mode,omitempty"`
	Origin                     *string               `xml:"origin,omitempty"`
	OriginalTemplate           *Template             `xml:"original_template,omitempty"`
	Os                         *OperatingSystem      `xml:"os,omitempty"`
	Payloads                   []Payload             `xml:"payloads,omitempty"`
	Permissions                []Permission          `xml:"permissions,omitempty"`
	PlacementPolicy            *VmPlacementPolicy    `xml:"placement_policy,omitempty"`
	Quota                      *Quota                `xml:"quota,omitempty"`
	ReportedDevices            []ReportedDevice      `xml:"reported_devices,omitempty"`
	RngDevice                  *RngDevice            `xml:"rng_device,omitempty"`
	RunOnce                    *bool                 `xml:"run_once,omitempty"`
	SerialNumber               *SerialNumber         `xml:"serial_number,omitempty"`
	Sessions                   []Session             `xml:"sessions,omitempty"`
	SmallIcon                  *Icon                 `xml:"small_icon,omitempty"`
	Snapshots                  []Snapshot            `xml:"snapshots,omitempty"`
	SoundcardEnabled           *bool                 `xml:"soundcard_enabled,omitempty"`
	Sso                        *Sso                  `xml:"sso,omitempty"`
	StartPaused                *bool                 `xml:"start_paused,omitempty"`
	StartTime                  time.Time             `xml:"start_time,omitempty"`
	Stateless                  *bool                 `xml:"stateless,omitempty"`
	Statistics                 []Statistic           `xml:"statistics,omitempty"`
	Status                     VmStatus              `xml:"status,omitempty"`
	StatusDetail               *string               `xml:"status_detail,omitempty"`
	StopReason                 *string               `xml:"stop_reason,omitempty"`
	StopTime                   time.Time             `xml:"stop_time,omitempty"`
	StorageDomain              *StorageDomain        `xml:"storage_domain,omitempty"`
	Tags                       []Tag                 `xml:"tags,omitempty"`
	Template                   *Template             `xml:"template,omitempty"`
	TimeZone                   *TimeZone             `xml:"time_zone,omitempty"`
	TunnelMigration            *bool                 `xml:"tunnel_migration,omitempty"`
	Type                       VmType                `xml:"type,omitempty"`
	Usb                        *Usb                  `xml:"usb,omitempty"`
	UseLatestTemplateVersion   *bool                 `xml:"use_latest_template_version,omitempty"`
	VirtioScsi                 *VirtioScsi           `xml:"virtio_scsi,omitempty"`
	VmPool                     *VmPool               `xml:"vm_pool,omitempty"`
	Watchdogs                  []Watchdog            `xml:"watchdogs,omitempty"`
}

type VmBases struct {
	XMLName xml.Name `xml:"vmbases"`
	VmBases []VmBase `xml:"vm_base,omitempty"`
}

type VmBase struct {
	OvStruct
	Bios                       *Bios               `xml:"bios,omitempty"`
	Cluster                    *Cluster            `xml:"cluster,omitempty"`
	Comment                    *string             `xml:"comment,omitempty"`
	Console                    *Console            `xml:"console,omitempty"`
	Cpu                        *Cpu                `xml:"cpu,omitempty"`
	CpuProfile                 *CpuProfile         `xml:"cpu_profile,omitempty"`
	CpuShares                  *int64              `xml:"cpu_shares,omitempty"`
	CreationTime               time.Time           `xml:"creation_time,omitempty"`
	CustomCompatibilityVersion *Version            `xml:"custom_compatibility_version,omitempty"`
	CustomCpuModel             *string             `xml:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      *string             `xml:"custom_emulated_machine,omitempty"`
	CustomProperties           []CustomProperty    `xml:"custom_properties,omitempty"`
	DeleteProtected            *bool               `xml:"delete_protected,omitempty"`
	Description                *string             `xml:"description,omitempty"`
	Display                    *Display            `xml:"display,omitempty"`
	Domain                     *Domain             `xml:"domain,omitempty"`
	HighAvailability           *HighAvailability   `xml:"high_availability,omitempty"`
	Id                         *string             `xml:"id,attr,omitempty"`
	Initialization             *Initialization     `xml:"initialization,omitempty"`
	Io                         *Io                 `xml:"io,omitempty"`
	LargeIcon                  *Icon               `xml:"large_icon,omitempty"`
	Lease                      *StorageDomainLease `xml:"lease,omitempty"`
	Memory                     *int64              `xml:"memory,omitempty"`
	MemoryPolicy               *MemoryPolicy       `xml:"memory_policy,omitempty"`
	Migration                  *MigrationOptions   `xml:"migration,omitempty"`
	MigrationDowntime          *int64              `xml:"migration_downtime,omitempty"`
	Name                       *string             `xml:"name,omitempty"`
	Origin                     *string             `xml:"origin,omitempty"`
	Os                         *OperatingSystem    `xml:"os,omitempty"`
	Quota                      *Quota              `xml:"quota,omitempty"`
	RngDevice                  *RngDevice          `xml:"rng_device,omitempty"`
	SerialNumber               *SerialNumber       `xml:"serial_number,omitempty"`
	SmallIcon                  *Icon               `xml:"small_icon,omitempty"`
	SoundcardEnabled           *bool               `xml:"soundcard_enabled,omitempty"`
	Sso                        *Sso                `xml:"sso,omitempty"`
	StartPaused                *bool               `xml:"start_paused,omitempty"`
	Stateless                  *bool               `xml:"stateless,omitempty"`
	StorageDomain              *StorageDomain      `xml:"storage_domain,omitempty"`
	TimeZone                   *TimeZone           `xml:"time_zone,omitempty"`
	TunnelMigration            *bool               `xml:"tunnel_migration,omitempty"`
	Type                       VmType              `xml:"type,omitempty"`
	Usb                        *Usb                `xml:"usb,omitempty"`
	VirtioScsi                 *VirtioScsi         `xml:"virtio_scsi,omitempty"`
}

type VmPlacementPolicys struct {
	XMLName            xml.Name            `xml:"vmplacementpolicys"`
	VmPlacementPolicys []VmPlacementPolicy `xml:"vm_placement_policy,omitempty"`
}

type VmPlacementPolicy struct {
	OvStruct
	Affinity VmAffinity `xml:"affinity,omitempty"`
	Hosts    []Host     `xml:"hosts,omitempty"`
}

type VmPools struct {
	XMLName xml.Name `xml:"vmpools"`
	VmPools []VmPool `xml:"vm_pool,omitempty"`
}

type VmPool struct {
	OvStruct
	AutoStorageSelect        *bool         `xml:"auto_storage_select,omitempty"`
	Cluster                  *Cluster      `xml:"cluster,omitempty"`
	Comment                  *string       `xml:"comment,omitempty"`
	Description              *string       `xml:"description,omitempty"`
	Display                  *Display      `xml:"display,omitempty"`
	Id                       *string       `xml:"id,attr,omitempty"`
	InstanceType             *InstanceType `xml:"instance_type,omitempty"`
	MaxUserVms               *int64        `xml:"max_user_vms,omitempty"`
	Name                     *string       `xml:"name,omitempty"`
	Permissions              []Permission  `xml:"permissions,omitempty"`
	PrestartedVms            *int64        `xml:"prestarted_vms,omitempty"`
	RngDevice                *RngDevice    `xml:"rng_device,omitempty"`
	Size                     *int64        `xml:"size,omitempty"`
	SoundcardEnabled         *bool         `xml:"soundcard_enabled,omitempty"`
	Stateful                 *bool         `xml:"stateful,omitempty"`
	Template                 *Template     `xml:"template,omitempty"`
	Type                     VmPoolType    `xml:"type,omitempty"`
	UseLatestTemplateVersion *bool         `xml:"use_latest_template_version,omitempty"`
	Vm                       *Vm           `xml:"vm,omitempty"`
}

type VmSummarys struct {
	XMLName    xml.Name    `xml:"vmsummarys"`
	VmSummarys []VmSummary `xml:"vm_summary,omitempty"`
}

type VmSummary struct {
	OvStruct
	Active    *int64 `xml:"active,omitempty"`
	Migrating *int64 `xml:"migrating,omitempty"`
	Total     *int64 `xml:"total,omitempty"`
}

type VnicPassThroughs struct {
	XMLName          xml.Name          `xml:"vnicpassthroughs"`
	VnicPassThroughs []VnicPassThrough `xml:"vnic_pass_through,omitempty"`
}

type VnicPassThrough struct {
	OvStruct
	Mode VnicPassThroughMode `xml:"mode,omitempty"`
}

type VnicProfiles struct {
	XMLName      xml.Name      `xml:"vnicprofiles"`
	VnicProfiles []VnicProfile `xml:"vnic_profile,omitempty"`
}

type VnicProfile struct {
	OvStruct
	Comment          *string          `xml:"comment,omitempty"`
	CustomProperties []CustomProperty `xml:"custom_properties,omitempty"`
	Description      *string          `xml:"description,omitempty"`
	Id               *string          `xml:"id,attr,omitempty"`
	Migratable       *bool            `xml:"migratable,omitempty"`
	Name             *string          `xml:"name,omitempty"`
	Network          *Network         `xml:"network,omitempty"`
	NetworkFilter    *NetworkFilter   `xml:"network_filter,omitempty"`
	PassThrough      *VnicPassThrough `xml:"pass_through,omitempty"`
	Permissions      []Permission     `xml:"permissions,omitempty"`
	PortMirroring    *bool            `xml:"port_mirroring,omitempty"`
	Qos              *Qos             `xml:"qos,omitempty"`
}

type VnicProfileMappings struct {
	XMLName             xml.Name             `xml:"vnicprofilemappings"`
	VnicProfileMappings []VnicProfileMapping `xml:"vnic_profile_mapping,omitempty"`
}

type VnicProfileMapping struct {
	OvStruct
	SourceNetworkName        *string      `xml:"source_network_name,omitempty"`
	SourceNetworkProfileName *string      `xml:"source_network_profile_name,omitempty"`
	TargetVnicProfile        *VnicProfile `xml:"target_vnic_profile,omitempty"`
}

type VolumeGroups struct {
	XMLName      xml.Name      `xml:"volumegroups"`
	VolumeGroups []VolumeGroup `xml:"volume_group,omitempty"`
}

type VolumeGroup struct {
	OvStruct
	Id           *string       `xml:"id,attr,omitempty"`
	LogicalUnits []LogicalUnit `xml:"logical_units,omitempty"`
	Name         *string       `xml:"name,omitempty"`
}

type Watchdogs struct {
	XMLName   xml.Name   `xml:"watchdogs"`
	Watchdogs []Watchdog `xml:"watchdog,omitempty"`
}

type Watchdog struct {
	OvStruct
	Action       WatchdogAction `xml:"action,omitempty"`
	Comment      *string        `xml:"comment,omitempty"`
	Description  *string        `xml:"description,omitempty"`
	Id           *string        `xml:"id,attr,omitempty"`
	InstanceType *InstanceType  `xml:"instance_type,omitempty"`
	Model        WatchdogModel  `xml:"model,omitempty"`
	Name         *string        `xml:"name,omitempty"`
	Template     *Template      `xml:"template,omitempty"`
	Vm           *Vm            `xml:"vm,omitempty"`
	Vms          []Vm           `xml:"vms,omitempty"`
}

type Weights struct {
	XMLName xml.Name `xml:"weights"`
	Weights []Weight `xml:"weight,omitempty"`
}

type Weight struct {
	OvStruct
	Comment              *string               `xml:"comment,omitempty"`
	Description          *string               `xml:"description,omitempty"`
	Factor               *int64                `xml:"factor,omitempty"`
	Id                   *string               `xml:"id,attr,omitempty"`
	Name                 *string               `xml:"name,omitempty"`
	SchedulingPolicy     *SchedulingPolicy     `xml:"scheduling_policy,omitempty"`
	SchedulingPolicyUnit *SchedulingPolicyUnit `xml:"scheduling_policy_unit,omitempty"`
}

type actionBuilder struct {
	action *Action
	err    error
}

func NewActionBuilder() *actionBuilder {
	return &actionBuilder{action: &Action{}, err: nil}
}

func (builder *actionBuilder) AllowPartialImport(allowPartialImport bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.AllowPartialImport = &allowPartialImport
	return builder
}

func (builder *actionBuilder) Async(async bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Async = &async
	return builder
}

func (builder *actionBuilder) Bricks(bricks []GlusterBrick) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Bricks = bricks
	return builder
}

func (builder *actionBuilder) Certificates(certificates []Certificate) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Certificates = certificates
	return builder
}

func (builder *actionBuilder) CheckConnectivity(checkConnectivity bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.CheckConnectivity = &checkConnectivity
	return builder
}

func (builder *actionBuilder) Clone(clone bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Clone = &clone
	return builder
}

func (builder *actionBuilder) Cluster(cluster *Cluster) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Cluster = cluster
	return builder
}

func (builder *actionBuilder) CollapseSnapshots(collapseSnapshots bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.CollapseSnapshots = &collapseSnapshots
	return builder
}

func (builder *actionBuilder) Comment(comment string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Comment = &comment
	return builder
}

func (builder *actionBuilder) ConnectivityTimeout(connectivityTimeout int64) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ConnectivityTimeout = &connectivityTimeout
	return builder
}

func (builder *actionBuilder) DataCenter(dataCenter *DataCenter) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.DataCenter = dataCenter
	return builder
}

func (builder *actionBuilder) DeployHostedEngine(deployHostedEngine bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.DeployHostedEngine = &deployHostedEngine
	return builder
}

func (builder *actionBuilder) Description(description string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Description = &description
	return builder
}

func (builder *actionBuilder) Details(details *GlusterVolumeProfileDetails) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Details = details
	return builder
}

func (builder *actionBuilder) DiscardSnapshots(discardSnapshots bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.DiscardSnapshots = &discardSnapshots
	return builder
}

func (builder *actionBuilder) Disk(disk *Disk) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Disk = disk
	return builder
}

func (builder *actionBuilder) Disks(disks []Disk) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Disks = disks
	return builder
}

func (builder *actionBuilder) Exclusive(exclusive bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Exclusive = &exclusive
	return builder
}

func (builder *actionBuilder) Fault(fault *Fault) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Fault = fault
	return builder
}

func (builder *actionBuilder) FenceType(fenceType string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.FenceType = &fenceType
	return builder
}

func (builder *actionBuilder) Filter(filter bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Filter = &filter
	return builder
}

func (builder *actionBuilder) FixLayout(fixLayout bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.FixLayout = &fixLayout
	return builder
}

func (builder *actionBuilder) Force(force bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Force = &force
	return builder
}

func (builder *actionBuilder) GracePeriod(gracePeriod *GracePeriod) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.GracePeriod = gracePeriod
	return builder
}

func (builder *actionBuilder) Host(host *Host) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Host = host
	return builder
}

func (builder *actionBuilder) Id(id string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Id = &id
	return builder
}

func (builder *actionBuilder) Image(image string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Image = &image
	return builder
}

func (builder *actionBuilder) ImportAsTemplate(importAsTemplate bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ImportAsTemplate = &importAsTemplate
	return builder
}

func (builder *actionBuilder) IsAttached(isAttached bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.IsAttached = &isAttached
	return builder
}

func (builder *actionBuilder) Iscsi(iscsi *IscsiDetails) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Iscsi = iscsi
	return builder
}

func (builder *actionBuilder) IscsiTargets(iscsiTargets []string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.IscsiTargets = iscsiTargets
	return builder
}

func (builder *actionBuilder) Job(job *Job) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Job = job
	return builder
}

func (builder *actionBuilder) LogicalUnits(logicalUnits []LogicalUnit) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.LogicalUnits = logicalUnits
	return builder
}

func (builder *actionBuilder) MaintenanceEnabled(maintenanceEnabled bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.MaintenanceEnabled = &maintenanceEnabled
	return builder
}

func (builder *actionBuilder) ModifiedBonds(modifiedBonds []HostNic) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ModifiedBonds = modifiedBonds
	return builder
}

func (builder *actionBuilder) ModifiedLabels(modifiedLabels []NetworkLabel) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ModifiedLabels = modifiedLabels
	return builder
}

func (builder *actionBuilder) ModifiedNetworkAttachments(modifiedNetworkAttachments []NetworkAttachment) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ModifiedNetworkAttachments = modifiedNetworkAttachments
	return builder
}

func (builder *actionBuilder) Name(name string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Name = &name
	return builder
}

func (builder *actionBuilder) Option(option *Option) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Option = option
	return builder
}

func (builder *actionBuilder) Pause(pause bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Pause = &pause
	return builder
}

func (builder *actionBuilder) PowerManagement(powerManagement *PowerManagement) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.PowerManagement = powerManagement
	return builder
}

func (builder *actionBuilder) ProxyTicket(proxyTicket *ProxyTicket) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ProxyTicket = proxyTicket
	return builder
}

func (builder *actionBuilder) Reason(reason string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Reason = &reason
	return builder
}

func (builder *actionBuilder) ReassignBadMacs(reassignBadMacs bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ReassignBadMacs = &reassignBadMacs
	return builder
}

func (builder *actionBuilder) RemoteViewerConnectionFile(remoteViewerConnectionFile string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.RemoteViewerConnectionFile = &remoteViewerConnectionFile
	return builder
}

func (builder *actionBuilder) RemovedBonds(removedBonds []HostNic) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.RemovedBonds = removedBonds
	return builder
}

func (builder *actionBuilder) RemovedLabels(removedLabels []NetworkLabel) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.RemovedLabels = removedLabels
	return builder
}

func (builder *actionBuilder) RemovedNetworkAttachments(removedNetworkAttachments []NetworkAttachment) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.RemovedNetworkAttachments = removedNetworkAttachments
	return builder
}

func (builder *actionBuilder) ResolutionType(resolutionType string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.ResolutionType = &resolutionType
	return builder
}

func (builder *actionBuilder) RestoreMemory(restoreMemory bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.RestoreMemory = &restoreMemory
	return builder
}

func (builder *actionBuilder) RootPassword(rootPassword string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.RootPassword = &rootPassword
	return builder
}

func (builder *actionBuilder) Snapshot(snapshot *Snapshot) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Snapshot = snapshot
	return builder
}

func (builder *actionBuilder) Ssh(ssh *Ssh) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Ssh = ssh
	return builder
}

func (builder *actionBuilder) Status(status string) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Status = &status
	return builder
}

func (builder *actionBuilder) StopGlusterService(stopGlusterService bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.StopGlusterService = &stopGlusterService
	return builder
}

func (builder *actionBuilder) StorageDomain(storageDomain *StorageDomain) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.StorageDomain = storageDomain
	return builder
}

func (builder *actionBuilder) StorageDomains(storageDomains []StorageDomain) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.StorageDomains = storageDomains
	return builder
}

func (builder *actionBuilder) Succeeded(succeeded bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Succeeded = &succeeded
	return builder
}

func (builder *actionBuilder) SynchronizedNetworkAttachments(synchronizedNetworkAttachments []NetworkAttachment) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.SynchronizedNetworkAttachments = synchronizedNetworkAttachments
	return builder
}

func (builder *actionBuilder) Template(template *Template) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Template = template
	return builder
}

func (builder *actionBuilder) Ticket(ticket *Ticket) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Ticket = ticket
	return builder
}

func (builder *actionBuilder) UndeployHostedEngine(undeployHostedEngine bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.UndeployHostedEngine = &undeployHostedEngine
	return builder
}

func (builder *actionBuilder) UseCloudInit(useCloudInit bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.UseCloudInit = &useCloudInit
	return builder
}

func (builder *actionBuilder) UseSysprep(useSysprep bool) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.UseSysprep = &useSysprep
	return builder
}

func (builder *actionBuilder) VirtualFunctionsConfiguration(virtualFunctionsConfiguration *HostNicVirtualFunctionsConfiguration) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.VirtualFunctionsConfiguration = virtualFunctionsConfiguration
	return builder
}

func (builder *actionBuilder) Vm(vm *Vm) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.Vm = vm
	return builder
}

func (builder *actionBuilder) VnicProfileMappings(vnicProfileMappings []VnicProfileMapping) *actionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.action.VnicProfileMappings = vnicProfileMappings
	return builder
}

func (builder *actionBuilder) Build() (*Action, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.action, nil
}

type affinityGroupBuilder struct {
	affinityGroup *AffinityGroup
	err           error
}

func NewAffinityGroupBuilder() *affinityGroupBuilder {
	return &affinityGroupBuilder{affinityGroup: &AffinityGroup{}, err: nil}
}

func (builder *affinityGroupBuilder) Cluster(cluster *Cluster) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Cluster = cluster
	return builder
}

func (builder *affinityGroupBuilder) Comment(comment string) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Comment = &comment
	return builder
}

func (builder *affinityGroupBuilder) Description(description string) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Description = &description
	return builder
}

func (builder *affinityGroupBuilder) Enforcing(enforcing bool) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Enforcing = &enforcing
	return builder
}

func (builder *affinityGroupBuilder) Hosts(hosts []Host) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Hosts = hosts
	return builder
}

func (builder *affinityGroupBuilder) HostsRule(hostsRule *AffinityRule) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.HostsRule = hostsRule
	return builder
}

func (builder *affinityGroupBuilder) Id(id string) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Id = &id
	return builder
}

func (builder *affinityGroupBuilder) Name(name string) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Name = &name
	return builder
}

func (builder *affinityGroupBuilder) Positive(positive bool) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Positive = &positive
	return builder
}

func (builder *affinityGroupBuilder) Vms(vms []Vm) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.Vms = vms
	return builder
}

func (builder *affinityGroupBuilder) VmsRule(vmsRule *AffinityRule) *affinityGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityGroup.VmsRule = vmsRule
	return builder
}

func (builder *affinityGroupBuilder) Build() (*AffinityGroup, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.affinityGroup, nil
}

type affinityLabelBuilder struct {
	affinityLabel *AffinityLabel
	err           error
}

func NewAffinityLabelBuilder() *affinityLabelBuilder {
	return &affinityLabelBuilder{affinityLabel: &AffinityLabel{}, err: nil}
}

func (builder *affinityLabelBuilder) Comment(comment string) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.Comment = &comment
	return builder
}

func (builder *affinityLabelBuilder) Description(description string) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.Description = &description
	return builder
}

func (builder *affinityLabelBuilder) Hosts(hosts []Host) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.Hosts = hosts
	return builder
}

func (builder *affinityLabelBuilder) Id(id string) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.Id = &id
	return builder
}

func (builder *affinityLabelBuilder) Name(name string) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.Name = &name
	return builder
}

func (builder *affinityLabelBuilder) ReadOnly(readOnly bool) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.ReadOnly = &readOnly
	return builder
}

func (builder *affinityLabelBuilder) Vms(vms []Vm) *affinityLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityLabel.Vms = vms
	return builder
}

func (builder *affinityLabelBuilder) Build() (*AffinityLabel, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.affinityLabel, nil
}

type affinityRuleBuilder struct {
	affinityRule *AffinityRule
	err          error
}

func NewAffinityRuleBuilder() *affinityRuleBuilder {
	return &affinityRuleBuilder{affinityRule: &AffinityRule{}, err: nil}
}

func (builder *affinityRuleBuilder) Enabled(enabled bool) *affinityRuleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityRule.Enabled = &enabled
	return builder
}

func (builder *affinityRuleBuilder) Enforcing(enforcing bool) *affinityRuleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityRule.Enforcing = &enforcing
	return builder
}

func (builder *affinityRuleBuilder) Positive(positive bool) *affinityRuleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.affinityRule.Positive = &positive
	return builder
}

func (builder *affinityRuleBuilder) Build() (*AffinityRule, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.affinityRule, nil
}

type agentBuilder struct {
	agent *Agent
	err   error
}

func NewAgentBuilder() *agentBuilder {
	return &agentBuilder{agent: &Agent{}, err: nil}
}

func (builder *agentBuilder) Address(address string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Address = &address
	return builder
}

func (builder *agentBuilder) Comment(comment string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Comment = &comment
	return builder
}

func (builder *agentBuilder) Concurrent(concurrent bool) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Concurrent = &concurrent
	return builder
}

func (builder *agentBuilder) Description(description string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Description = &description
	return builder
}

func (builder *agentBuilder) EncryptOptions(encryptOptions bool) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.EncryptOptions = &encryptOptions
	return builder
}

func (builder *agentBuilder) Host(host *Host) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Host = host
	return builder
}

func (builder *agentBuilder) Id(id string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Id = &id
	return builder
}

func (builder *agentBuilder) Name(name string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Name = &name
	return builder
}

func (builder *agentBuilder) Options(options []Option) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Options = options
	return builder
}

func (builder *agentBuilder) Order(order int64) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Order = &order
	return builder
}

func (builder *agentBuilder) Password(password string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Password = &password
	return builder
}

func (builder *agentBuilder) Port(port int64) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Port = &port
	return builder
}

func (builder *agentBuilder) Type(type_ string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Type = &type_
	return builder
}

func (builder *agentBuilder) Username(username string) *agentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agent.Username = &username
	return builder
}

func (builder *agentBuilder) Build() (*Agent, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.agent, nil
}

type agentConfigurationBuilder struct {
	agentConfiguration *AgentConfiguration
	err                error
}

func NewAgentConfigurationBuilder() *agentConfigurationBuilder {
	return &agentConfigurationBuilder{agentConfiguration: &AgentConfiguration{}, err: nil}
}

func (builder *agentConfigurationBuilder) Address(address string) *agentConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agentConfiguration.Address = &address
	return builder
}

func (builder *agentConfigurationBuilder) BrokerType(brokerType MessageBrokerType) *agentConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agentConfiguration.BrokerType = brokerType
	return builder
}

func (builder *agentConfigurationBuilder) NetworkMappings(networkMappings string) *agentConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agentConfiguration.NetworkMappings = &networkMappings
	return builder
}

func (builder *agentConfigurationBuilder) Password(password string) *agentConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agentConfiguration.Password = &password
	return builder
}

func (builder *agentConfigurationBuilder) Port(port int64) *agentConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agentConfiguration.Port = &port
	return builder
}

func (builder *agentConfigurationBuilder) Username(username string) *agentConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.agentConfiguration.Username = &username
	return builder
}

func (builder *agentConfigurationBuilder) Build() (*AgentConfiguration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.agentConfiguration, nil
}

type apiBuilder struct {
	api *Api
	err error
}

func NewApiBuilder() *apiBuilder {
	return &apiBuilder{api: &Api{}, err: nil}
}

func (builder *apiBuilder) ProductInfo(productInfo *ProductInfo) *apiBuilder {
	if builder.err != nil {
		return builder
	}

	builder.api.ProductInfo = productInfo
	return builder
}

func (builder *apiBuilder) SpecialObjects(specialObjects *SpecialObjects) *apiBuilder {
	if builder.err != nil {
		return builder
	}

	builder.api.SpecialObjects = specialObjects
	return builder
}

func (builder *apiBuilder) Summary(summary *ApiSummary) *apiBuilder {
	if builder.err != nil {
		return builder
	}

	builder.api.Summary = summary
	return builder
}

func (builder *apiBuilder) Time(time time.Time) *apiBuilder {
	if builder.err != nil {
		return builder
	}

	builder.api.Time = time
	return builder
}

func (builder *apiBuilder) Build() (*Api, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.api, nil
}

type apiSummaryBuilder struct {
	apiSummary *ApiSummary
	err        error
}

func NewApiSummaryBuilder() *apiSummaryBuilder {
	return &apiSummaryBuilder{apiSummary: &ApiSummary{}, err: nil}
}

func (builder *apiSummaryBuilder) Hosts(hosts *ApiSummaryItem) *apiSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.apiSummary.Hosts = hosts
	return builder
}

func (builder *apiSummaryBuilder) StorageDomains(storageDomains *ApiSummaryItem) *apiSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.apiSummary.StorageDomains = storageDomains
	return builder
}

func (builder *apiSummaryBuilder) Users(users *ApiSummaryItem) *apiSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.apiSummary.Users = users
	return builder
}

func (builder *apiSummaryBuilder) Vms(vms *ApiSummaryItem) *apiSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.apiSummary.Vms = vms
	return builder
}

func (builder *apiSummaryBuilder) Build() (*ApiSummary, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.apiSummary, nil
}

type apiSummaryItemBuilder struct {
	apiSummaryItem *ApiSummaryItem
	err            error
}

func NewApiSummaryItemBuilder() *apiSummaryItemBuilder {
	return &apiSummaryItemBuilder{apiSummaryItem: &ApiSummaryItem{}, err: nil}
}

func (builder *apiSummaryItemBuilder) Active(active int64) *apiSummaryItemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.apiSummaryItem.Active = &active
	return builder
}

func (builder *apiSummaryItemBuilder) Total(total int64) *apiSummaryItemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.apiSummaryItem.Total = &total
	return builder
}

func (builder *apiSummaryItemBuilder) Build() (*ApiSummaryItem, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.apiSummaryItem, nil
}

type applicationBuilder struct {
	application *Application
	err         error
}

func NewApplicationBuilder() *applicationBuilder {
	return &applicationBuilder{application: &Application{}, err: nil}
}

func (builder *applicationBuilder) Comment(comment string) *applicationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.application.Comment = &comment
	return builder
}

func (builder *applicationBuilder) Description(description string) *applicationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.application.Description = &description
	return builder
}

func (builder *applicationBuilder) Id(id string) *applicationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.application.Id = &id
	return builder
}

func (builder *applicationBuilder) Name(name string) *applicationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.application.Name = &name
	return builder
}

func (builder *applicationBuilder) Vm(vm *Vm) *applicationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.application.Vm = vm
	return builder
}

func (builder *applicationBuilder) Build() (*Application, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.application, nil
}

type authorizedKeyBuilder struct {
	authorizedKey *AuthorizedKey
	err           error
}

func NewAuthorizedKeyBuilder() *authorizedKeyBuilder {
	return &authorizedKeyBuilder{authorizedKey: &AuthorizedKey{}, err: nil}
}

func (builder *authorizedKeyBuilder) Comment(comment string) *authorizedKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.authorizedKey.Comment = &comment
	return builder
}

func (builder *authorizedKeyBuilder) Description(description string) *authorizedKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.authorizedKey.Description = &description
	return builder
}

func (builder *authorizedKeyBuilder) Id(id string) *authorizedKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.authorizedKey.Id = &id
	return builder
}

func (builder *authorizedKeyBuilder) Key(key string) *authorizedKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.authorizedKey.Key = &key
	return builder
}

func (builder *authorizedKeyBuilder) Name(name string) *authorizedKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.authorizedKey.Name = &name
	return builder
}

func (builder *authorizedKeyBuilder) User(user *User) *authorizedKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.authorizedKey.User = user
	return builder
}

func (builder *authorizedKeyBuilder) Build() (*AuthorizedKey, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.authorizedKey, nil
}

type balanceBuilder struct {
	balance *Balance
	err     error
}

func NewBalanceBuilder() *balanceBuilder {
	return &balanceBuilder{balance: &Balance{}, err: nil}
}

func (builder *balanceBuilder) Comment(comment string) *balanceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.balance.Comment = &comment
	return builder
}

func (builder *balanceBuilder) Description(description string) *balanceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.balance.Description = &description
	return builder
}

func (builder *balanceBuilder) Id(id string) *balanceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.balance.Id = &id
	return builder
}

func (builder *balanceBuilder) Name(name string) *balanceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.balance.Name = &name
	return builder
}

func (builder *balanceBuilder) SchedulingPolicy(schedulingPolicy *SchedulingPolicy) *balanceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.balance.SchedulingPolicy = schedulingPolicy
	return builder
}

func (builder *balanceBuilder) SchedulingPolicyUnit(schedulingPolicyUnit *SchedulingPolicyUnit) *balanceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.balance.SchedulingPolicyUnit = schedulingPolicyUnit
	return builder
}

func (builder *balanceBuilder) Build() (*Balance, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.balance, nil
}

type biosBuilder struct {
	bios *Bios
	err  error
}

func NewBiosBuilder() *biosBuilder {
	return &biosBuilder{bios: &Bios{}, err: nil}
}

func (builder *biosBuilder) BootMenu(bootMenu *BootMenu) *biosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bios.BootMenu = bootMenu
	return builder
}

func (builder *biosBuilder) Build() (*Bios, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.bios, nil
}

type blockStatisticBuilder struct {
	blockStatistic *BlockStatistic
	err            error
}

func NewBlockStatisticBuilder() *blockStatisticBuilder {
	return &blockStatisticBuilder{blockStatistic: &BlockStatistic{}, err: nil}
}

func (builder *blockStatisticBuilder) Statistics(statistics []Statistic) *blockStatisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.blockStatistic.Statistics = statistics
	return builder
}

func (builder *blockStatisticBuilder) Build() (*BlockStatistic, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.blockStatistic, nil
}

type bondingBuilder struct {
	bonding *Bonding
	err     error
}

func NewBondingBuilder() *bondingBuilder {
	return &bondingBuilder{bonding: &Bonding{}, err: nil}
}

func (builder *bondingBuilder) ActiveSlave(activeSlave *HostNic) *bondingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bonding.ActiveSlave = activeSlave
	return builder
}

func (builder *bondingBuilder) AdPartnerMac(adPartnerMac *Mac) *bondingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bonding.AdPartnerMac = adPartnerMac
	return builder
}

func (builder *bondingBuilder) Options(options []Option) *bondingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bonding.Options = options
	return builder
}

func (builder *bondingBuilder) Slaves(slaves []HostNic) *bondingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bonding.Slaves = slaves
	return builder
}

func (builder *bondingBuilder) Build() (*Bonding, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.bonding, nil
}

type bookmarkBuilder struct {
	bookmark *Bookmark
	err      error
}

func NewBookmarkBuilder() *bookmarkBuilder {
	return &bookmarkBuilder{bookmark: &Bookmark{}, err: nil}
}

func (builder *bookmarkBuilder) Comment(comment string) *bookmarkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bookmark.Comment = &comment
	return builder
}

func (builder *bookmarkBuilder) Description(description string) *bookmarkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bookmark.Description = &description
	return builder
}

func (builder *bookmarkBuilder) Id(id string) *bookmarkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bookmark.Id = &id
	return builder
}

func (builder *bookmarkBuilder) Name(name string) *bookmarkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bookmark.Name = &name
	return builder
}

func (builder *bookmarkBuilder) Value(value string) *bookmarkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bookmark.Value = &value
	return builder
}

func (builder *bookmarkBuilder) Build() (*Bookmark, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.bookmark, nil
}

type bootBuilder struct {
	boot *Boot
	err  error
}

func NewBootBuilder() *bootBuilder {
	return &bootBuilder{boot: &Boot{}, err: nil}
}

func (builder *bootBuilder) Devices(devices []BootDevice) *bootBuilder {
	if builder.err != nil {
		return builder
	}

	builder.boot.Devices = devices
	return builder
}

func (builder *bootBuilder) Build() (*Boot, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.boot, nil
}

type bootMenuBuilder struct {
	bootMenu *BootMenu
	err      error
}

func NewBootMenuBuilder() *bootMenuBuilder {
	return &bootMenuBuilder{bootMenu: &BootMenu{}, err: nil}
}

func (builder *bootMenuBuilder) Enabled(enabled bool) *bootMenuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.bootMenu.Enabled = &enabled
	return builder
}

func (builder *bootMenuBuilder) Build() (*BootMenu, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.bootMenu, nil
}

type brickProfileDetailBuilder struct {
	brickProfileDetail *BrickProfileDetail
	err                error
}

func NewBrickProfileDetailBuilder() *brickProfileDetailBuilder {
	return &brickProfileDetailBuilder{brickProfileDetail: &BrickProfileDetail{}, err: nil}
}

func (builder *brickProfileDetailBuilder) Brick(brick *GlusterBrick) *brickProfileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.brickProfileDetail.Brick = brick
	return builder
}

func (builder *brickProfileDetailBuilder) ProfileDetails(profileDetails []ProfileDetail) *brickProfileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.brickProfileDetail.ProfileDetails = profileDetails
	return builder
}

func (builder *brickProfileDetailBuilder) Build() (*BrickProfileDetail, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.brickProfileDetail, nil
}

type cdromBuilder struct {
	cdrom *Cdrom
	err   error
}

func NewCdromBuilder() *cdromBuilder {
	return &cdromBuilder{cdrom: &Cdrom{}, err: nil}
}

func (builder *cdromBuilder) Comment(comment string) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Comment = &comment
	return builder
}

func (builder *cdromBuilder) Description(description string) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Description = &description
	return builder
}

func (builder *cdromBuilder) File(file *File) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.File = file
	return builder
}

func (builder *cdromBuilder) Id(id string) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Id = &id
	return builder
}

func (builder *cdromBuilder) InstanceType(instanceType *InstanceType) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.InstanceType = instanceType
	return builder
}

func (builder *cdromBuilder) Name(name string) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Name = &name
	return builder
}

func (builder *cdromBuilder) Template(template *Template) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Template = template
	return builder
}

func (builder *cdromBuilder) Vm(vm *Vm) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Vm = vm
	return builder
}

func (builder *cdromBuilder) Vms(vms []Vm) *cdromBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cdrom.Vms = vms
	return builder
}

func (builder *cdromBuilder) Build() (*Cdrom, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cdrom, nil
}

type certificateBuilder struct {
	certificate *Certificate
	err         error
}

func NewCertificateBuilder() *certificateBuilder {
	return &certificateBuilder{certificate: &Certificate{}, err: nil}
}

func (builder *certificateBuilder) Comment(comment string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Comment = &comment
	return builder
}

func (builder *certificateBuilder) Content(content string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Content = &content
	return builder
}

func (builder *certificateBuilder) Description(description string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Description = &description
	return builder
}

func (builder *certificateBuilder) Id(id string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Id = &id
	return builder
}

func (builder *certificateBuilder) Name(name string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Name = &name
	return builder
}

func (builder *certificateBuilder) Organization(organization string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Organization = &organization
	return builder
}

func (builder *certificateBuilder) Subject(subject string) *certificateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.certificate.Subject = &subject
	return builder
}

func (builder *certificateBuilder) Build() (*Certificate, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.certificate, nil
}

type cloudInitBuilder struct {
	cloudInit *CloudInit
	err       error
}

func NewCloudInitBuilder() *cloudInitBuilder {
	return &cloudInitBuilder{cloudInit: &CloudInit{}, err: nil}
}

func (builder *cloudInitBuilder) AuthorizedKeys(authorizedKeys []AuthorizedKey) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.AuthorizedKeys = authorizedKeys
	return builder
}

func (builder *cloudInitBuilder) Files(files []File) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.Files = files
	return builder
}

func (builder *cloudInitBuilder) Host(host *Host) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.Host = host
	return builder
}

func (builder *cloudInitBuilder) NetworkConfiguration(networkConfiguration *NetworkConfiguration) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.NetworkConfiguration = networkConfiguration
	return builder
}

func (builder *cloudInitBuilder) RegenerateSshKeys(regenerateSshKeys bool) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.RegenerateSshKeys = &regenerateSshKeys
	return builder
}

func (builder *cloudInitBuilder) Timezone(timezone string) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.Timezone = &timezone
	return builder
}

func (builder *cloudInitBuilder) Users(users []User) *cloudInitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cloudInit.Users = users
	return builder
}

func (builder *cloudInitBuilder) Build() (*CloudInit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cloudInit, nil
}

type clusterBuilder struct {
	cluster *Cluster
	err     error
}

func NewClusterBuilder() *clusterBuilder {
	return &clusterBuilder{cluster: &Cluster{}, err: nil}
}

func (builder *clusterBuilder) AffinityGroups(affinityGroups []AffinityGroup) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.AffinityGroups = affinityGroups
	return builder
}

func (builder *clusterBuilder) BallooningEnabled(ballooningEnabled bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.BallooningEnabled = &ballooningEnabled
	return builder
}

func (builder *clusterBuilder) Comment(comment string) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Comment = &comment
	return builder
}

func (builder *clusterBuilder) Cpu(cpu *Cpu) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Cpu = cpu
	return builder
}

func (builder *clusterBuilder) CpuProfiles(cpuProfiles []CpuProfile) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.CpuProfiles = cpuProfiles
	return builder
}

func (builder *clusterBuilder) CustomSchedulingPolicyProperties(customSchedulingPolicyProperties []Property) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.CustomSchedulingPolicyProperties = customSchedulingPolicyProperties
	return builder
}

func (builder *clusterBuilder) DataCenter(dataCenter *DataCenter) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.DataCenter = dataCenter
	return builder
}

func (builder *clusterBuilder) Description(description string) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Description = &description
	return builder
}

func (builder *clusterBuilder) Display(display *Display) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Display = display
	return builder
}

func (builder *clusterBuilder) ErrorHandling(errorHandling *ErrorHandling) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.ErrorHandling = errorHandling
	return builder
}

func (builder *clusterBuilder) FencingPolicy(fencingPolicy *FencingPolicy) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.FencingPolicy = fencingPolicy
	return builder
}

func (builder *clusterBuilder) GlusterHooks(glusterHooks []GlusterHook) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.GlusterHooks = glusterHooks
	return builder
}

func (builder *clusterBuilder) GlusterService(glusterService bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.GlusterService = &glusterService
	return builder
}

func (builder *clusterBuilder) GlusterTunedProfile(glusterTunedProfile string) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.GlusterTunedProfile = &glusterTunedProfile
	return builder
}

func (builder *clusterBuilder) GlusterVolumes(glusterVolumes []GlusterVolume) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.GlusterVolumes = glusterVolumes
	return builder
}

func (builder *clusterBuilder) HaReservation(haReservation bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.HaReservation = &haReservation
	return builder
}

func (builder *clusterBuilder) Id(id string) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Id = &id
	return builder
}

func (builder *clusterBuilder) Ksm(ksm *Ksm) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Ksm = ksm
	return builder
}

func (builder *clusterBuilder) MacPool(macPool *MacPool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.MacPool = macPool
	return builder
}

func (builder *clusterBuilder) MaintenanceReasonRequired(maintenanceReasonRequired bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.MaintenanceReasonRequired = &maintenanceReasonRequired
	return builder
}

func (builder *clusterBuilder) ManagementNetwork(managementNetwork *Network) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.ManagementNetwork = managementNetwork
	return builder
}

func (builder *clusterBuilder) MemoryPolicy(memoryPolicy *MemoryPolicy) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.MemoryPolicy = memoryPolicy
	return builder
}

func (builder *clusterBuilder) Migration(migration *MigrationOptions) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Migration = migration
	return builder
}

func (builder *clusterBuilder) Name(name string) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Name = &name
	return builder
}

func (builder *clusterBuilder) NetworkFilters(networkFilters []NetworkFilter) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.NetworkFilters = networkFilters
	return builder
}

func (builder *clusterBuilder) Networks(networks []Network) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Networks = networks
	return builder
}

func (builder *clusterBuilder) OptionalReason(optionalReason bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.OptionalReason = &optionalReason
	return builder
}

func (builder *clusterBuilder) Permissions(permissions []Permission) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Permissions = permissions
	return builder
}

func (builder *clusterBuilder) RequiredRngSources(requiredRngSources []RngSource) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.RequiredRngSources = requiredRngSources
	return builder
}

func (builder *clusterBuilder) SchedulingPolicy(schedulingPolicy *SchedulingPolicy) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.SchedulingPolicy = schedulingPolicy
	return builder
}

func (builder *clusterBuilder) SerialNumber(serialNumber *SerialNumber) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.SerialNumber = serialNumber
	return builder
}

func (builder *clusterBuilder) SupportedVersions(supportedVersions []Version) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.SupportedVersions = supportedVersions
	return builder
}

func (builder *clusterBuilder) SwitchType(switchType SwitchType) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.SwitchType = switchType
	return builder
}

func (builder *clusterBuilder) ThreadsAsCores(threadsAsCores bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.ThreadsAsCores = &threadsAsCores
	return builder
}

func (builder *clusterBuilder) TrustedService(trustedService bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.TrustedService = &trustedService
	return builder
}

func (builder *clusterBuilder) TunnelMigration(tunnelMigration bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.TunnelMigration = &tunnelMigration
	return builder
}

func (builder *clusterBuilder) Version(version *Version) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.Version = version
	return builder
}

func (builder *clusterBuilder) VirtService(virtService bool) *clusterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cluster.VirtService = &virtService
	return builder
}

func (builder *clusterBuilder) Build() (*Cluster, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cluster, nil
}

type clusterLevelBuilder struct {
	clusterLevel *ClusterLevel
	err          error
}

func NewClusterLevelBuilder() *clusterLevelBuilder {
	return &clusterLevelBuilder{clusterLevel: &ClusterLevel{}, err: nil}
}

func (builder *clusterLevelBuilder) Comment(comment string) *clusterLevelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.clusterLevel.Comment = &comment
	return builder
}

func (builder *clusterLevelBuilder) CpuTypes(cpuTypes []CpuType) *clusterLevelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.clusterLevel.CpuTypes = cpuTypes
	return builder
}

func (builder *clusterLevelBuilder) Description(description string) *clusterLevelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.clusterLevel.Description = &description
	return builder
}

func (builder *clusterLevelBuilder) Id(id string) *clusterLevelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.clusterLevel.Id = &id
	return builder
}

func (builder *clusterLevelBuilder) Name(name string) *clusterLevelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.clusterLevel.Name = &name
	return builder
}

func (builder *clusterLevelBuilder) Permits(permits []Permit) *clusterLevelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.clusterLevel.Permits = permits
	return builder
}

func (builder *clusterLevelBuilder) Build() (*ClusterLevel, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.clusterLevel, nil
}

type configurationBuilder struct {
	configuration *Configuration
	err           error
}

func NewConfigurationBuilder() *configurationBuilder {
	return &configurationBuilder{configuration: &Configuration{}, err: nil}
}

func (builder *configurationBuilder) Data(data string) *configurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.configuration.Data = &data
	return builder
}

func (builder *configurationBuilder) Type(type_ ConfigurationType) *configurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.configuration.Type = type_
	return builder
}

func (builder *configurationBuilder) Build() (*Configuration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.configuration, nil
}

type consoleBuilder struct {
	console *Console
	err     error
}

func NewConsoleBuilder() *consoleBuilder {
	return &consoleBuilder{console: &Console{}, err: nil}
}

func (builder *consoleBuilder) Enabled(enabled bool) *consoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.console.Enabled = &enabled
	return builder
}

func (builder *consoleBuilder) Build() (*Console, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.console, nil
}

type coreBuilder struct {
	core *Core
	err  error
}

func NewCoreBuilder() *coreBuilder {
	return &coreBuilder{core: &Core{}, err: nil}
}

func (builder *coreBuilder) Index(index int64) *coreBuilder {
	if builder.err != nil {
		return builder
	}

	builder.core.Index = &index
	return builder
}

func (builder *coreBuilder) Socket(socket int64) *coreBuilder {
	if builder.err != nil {
		return builder
	}

	builder.core.Socket = &socket
	return builder
}

func (builder *coreBuilder) Build() (*Core, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.core, nil
}

type cpuBuilder struct {
	cpu *Cpu
	err error
}

func NewCpuBuilder() *cpuBuilder {
	return &cpuBuilder{cpu: &Cpu{}, err: nil}
}

func (builder *cpuBuilder) Architecture(architecture Architecture) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Architecture = architecture
	return builder
}

func (builder *cpuBuilder) Cores(cores []Core) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Cores = cores
	return builder
}

func (builder *cpuBuilder) CpuTune(cpuTune *CpuTune) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.CpuTune = cpuTune
	return builder
}

func (builder *cpuBuilder) Level(level int64) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Level = &level
	return builder
}

func (builder *cpuBuilder) Mode(mode CpuMode) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Mode = mode
	return builder
}

func (builder *cpuBuilder) Name(name string) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Name = &name
	return builder
}

func (builder *cpuBuilder) Speed(speed float64) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Speed = &speed
	return builder
}

func (builder *cpuBuilder) Topology(topology *CpuTopology) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Topology = topology
	return builder
}

func (builder *cpuBuilder) Type(type_ string) *cpuBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpu.Type = &type_
	return builder
}

func (builder *cpuBuilder) Build() (*Cpu, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cpu, nil
}

type cpuProfileBuilder struct {
	cpuProfile *CpuProfile
	err        error
}

func NewCpuProfileBuilder() *cpuProfileBuilder {
	return &cpuProfileBuilder{cpuProfile: &CpuProfile{}, err: nil}
}

func (builder *cpuProfileBuilder) Cluster(cluster *Cluster) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Cluster = cluster
	return builder
}

func (builder *cpuProfileBuilder) Comment(comment string) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Comment = &comment
	return builder
}

func (builder *cpuProfileBuilder) Description(description string) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Description = &description
	return builder
}

func (builder *cpuProfileBuilder) Id(id string) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Id = &id
	return builder
}

func (builder *cpuProfileBuilder) Name(name string) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Name = &name
	return builder
}

func (builder *cpuProfileBuilder) Permissions(permissions []Permission) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Permissions = permissions
	return builder
}

func (builder *cpuProfileBuilder) Qos(qos *Qos) *cpuProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuProfile.Qos = qos
	return builder
}

func (builder *cpuProfileBuilder) Build() (*CpuProfile, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cpuProfile, nil
}

type cpuTopologyBuilder struct {
	cpuTopology *CpuTopology
	err         error
}

func NewCpuTopologyBuilder() *cpuTopologyBuilder {
	return &cpuTopologyBuilder{cpuTopology: &CpuTopology{}, err: nil}
}

func (builder *cpuTopologyBuilder) Cores(cores int64) *cpuTopologyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuTopology.Cores = &cores
	return builder
}

func (builder *cpuTopologyBuilder) Sockets(sockets int64) *cpuTopologyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuTopology.Sockets = &sockets
	return builder
}

func (builder *cpuTopologyBuilder) Threads(threads int64) *cpuTopologyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuTopology.Threads = &threads
	return builder
}

func (builder *cpuTopologyBuilder) Build() (*CpuTopology, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cpuTopology, nil
}

type cpuTuneBuilder struct {
	cpuTune *CpuTune
	err     error
}

func NewCpuTuneBuilder() *cpuTuneBuilder {
	return &cpuTuneBuilder{cpuTune: &CpuTune{}, err: nil}
}

func (builder *cpuTuneBuilder) VcpuPins(vcpuPins []VcpuPin) *cpuTuneBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuTune.VcpuPins = vcpuPins
	return builder
}

func (builder *cpuTuneBuilder) Build() (*CpuTune, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cpuTune, nil
}

type cpuTypeBuilder struct {
	cpuType *CpuType
	err     error
}

func NewCpuTypeBuilder() *cpuTypeBuilder {
	return &cpuTypeBuilder{cpuType: &CpuType{}, err: nil}
}

func (builder *cpuTypeBuilder) Architecture(architecture Architecture) *cpuTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuType.Architecture = architecture
	return builder
}

func (builder *cpuTypeBuilder) Level(level int64) *cpuTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuType.Level = &level
	return builder
}

func (builder *cpuTypeBuilder) Name(name string) *cpuTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.cpuType.Name = &name
	return builder
}

func (builder *cpuTypeBuilder) Build() (*CpuType, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.cpuType, nil
}

type customPropertyBuilder struct {
	customProperty *CustomProperty
	err            error
}

func NewCustomPropertyBuilder() *customPropertyBuilder {
	return &customPropertyBuilder{customProperty: &CustomProperty{}, err: nil}
}

func (builder *customPropertyBuilder) Name(name string) *customPropertyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.customProperty.Name = &name
	return builder
}

func (builder *customPropertyBuilder) Regexp(regexp string) *customPropertyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.customProperty.Regexp = &regexp
	return builder
}

func (builder *customPropertyBuilder) Value(value string) *customPropertyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.customProperty.Value = &value
	return builder
}

func (builder *customPropertyBuilder) Build() (*CustomProperty, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.customProperty, nil
}

type dataCenterBuilder struct {
	dataCenter *DataCenter
	err        error
}

func NewDataCenterBuilder() *dataCenterBuilder {
	return &dataCenterBuilder{dataCenter: &DataCenter{}, err: nil}
}

func (builder *dataCenterBuilder) Clusters(clusters []Cluster) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Clusters = clusters
	return builder
}

func (builder *dataCenterBuilder) Comment(comment string) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Comment = &comment
	return builder
}

func (builder *dataCenterBuilder) Description(description string) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Description = &description
	return builder
}

func (builder *dataCenterBuilder) Id(id string) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Id = &id
	return builder
}

func (builder *dataCenterBuilder) IscsiBonds(iscsiBonds []IscsiBond) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.IscsiBonds = iscsiBonds
	return builder
}

func (builder *dataCenterBuilder) Local(local bool) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Local = &local
	return builder
}

func (builder *dataCenterBuilder) MacPool(macPool *MacPool) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.MacPool = macPool
	return builder
}

func (builder *dataCenterBuilder) Name(name string) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Name = &name
	return builder
}

func (builder *dataCenterBuilder) Networks(networks []Network) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Networks = networks
	return builder
}

func (builder *dataCenterBuilder) Permissions(permissions []Permission) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Permissions = permissions
	return builder
}

func (builder *dataCenterBuilder) Qoss(qoss []Qos) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Qoss = qoss
	return builder
}

func (builder *dataCenterBuilder) QuotaMode(quotaMode QuotaModeType) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.QuotaMode = quotaMode
	return builder
}

func (builder *dataCenterBuilder) Quotas(quotas []Quota) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Quotas = quotas
	return builder
}

func (builder *dataCenterBuilder) Status(status DataCenterStatus) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Status = status
	return builder
}

func (builder *dataCenterBuilder) StorageDomains(storageDomains []StorageDomain) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.StorageDomains = storageDomains
	return builder
}

func (builder *dataCenterBuilder) StorageFormat(storageFormat StorageFormat) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.StorageFormat = storageFormat
	return builder
}

func (builder *dataCenterBuilder) SupportedVersions(supportedVersions []Version) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.SupportedVersions = supportedVersions
	return builder
}

func (builder *dataCenterBuilder) Version(version *Version) *dataCenterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dataCenter.Version = version
	return builder
}

func (builder *dataCenterBuilder) Build() (*DataCenter, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.dataCenter, nil
}

type deviceBuilder struct {
	device *Device
	err    error
}

func NewDeviceBuilder() *deviceBuilder {
	return &deviceBuilder{device: &Device{}, err: nil}
}

func (builder *deviceBuilder) Comment(comment string) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Comment = &comment
	return builder
}

func (builder *deviceBuilder) Description(description string) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Description = &description
	return builder
}

func (builder *deviceBuilder) Id(id string) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Id = &id
	return builder
}

func (builder *deviceBuilder) InstanceType(instanceType *InstanceType) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.InstanceType = instanceType
	return builder
}

func (builder *deviceBuilder) Name(name string) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Name = &name
	return builder
}

func (builder *deviceBuilder) Template(template *Template) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Template = template
	return builder
}

func (builder *deviceBuilder) Vm(vm *Vm) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Vm = vm
	return builder
}

func (builder *deviceBuilder) Vms(vms []Vm) *deviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.device.Vms = vms
	return builder
}

func (builder *deviceBuilder) Build() (*Device, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.device, nil
}

type diskBuilder struct {
	disk *Disk
	err  error
}

func NewDiskBuilder() *diskBuilder {
	return &diskBuilder{disk: &Disk{}, err: nil}
}

func (builder *diskBuilder) Active(active bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Active = &active
	return builder
}

func (builder *diskBuilder) ActualSize(actualSize int64) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.ActualSize = &actualSize
	return builder
}

func (builder *diskBuilder) Alias(alias string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Alias = &alias
	return builder
}

func (builder *diskBuilder) Bootable(bootable bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Bootable = &bootable
	return builder
}

func (builder *diskBuilder) Comment(comment string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Comment = &comment
	return builder
}

func (builder *diskBuilder) Description(description string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Description = &description
	return builder
}

func (builder *diskBuilder) DiskProfile(diskProfile *DiskProfile) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.DiskProfile = diskProfile
	return builder
}

func (builder *diskBuilder) Format(format DiskFormat) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Format = format
	return builder
}

func (builder *diskBuilder) Id(id string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Id = &id
	return builder
}

func (builder *diskBuilder) ImageId(imageId string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.ImageId = &imageId
	return builder
}

func (builder *diskBuilder) InitialSize(initialSize int64) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.InitialSize = &initialSize
	return builder
}

func (builder *diskBuilder) InstanceType(instanceType *InstanceType) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.InstanceType = instanceType
	return builder
}

func (builder *diskBuilder) Interface(interface_ DiskInterface) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Interface = interface_
	return builder
}

func (builder *diskBuilder) LogicalName(logicalName string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.LogicalName = &logicalName
	return builder
}

func (builder *diskBuilder) LunStorage(lunStorage *HostStorage) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.LunStorage = lunStorage
	return builder
}

func (builder *diskBuilder) Name(name string) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Name = &name
	return builder
}

func (builder *diskBuilder) OpenstackVolumeType(openstackVolumeType *OpenStackVolumeType) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.OpenstackVolumeType = openstackVolumeType
	return builder
}

func (builder *diskBuilder) Permissions(permissions []Permission) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Permissions = permissions
	return builder
}

func (builder *diskBuilder) PropagateErrors(propagateErrors bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.PropagateErrors = &propagateErrors
	return builder
}

func (builder *diskBuilder) ProvisionedSize(provisionedSize int64) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.ProvisionedSize = &provisionedSize
	return builder
}

func (builder *diskBuilder) QcowVersion(qcowVersion QcowVersion) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.QcowVersion = qcowVersion
	return builder
}

func (builder *diskBuilder) Quota(quota *Quota) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Quota = quota
	return builder
}

func (builder *diskBuilder) ReadOnly(readOnly bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.ReadOnly = &readOnly
	return builder
}

func (builder *diskBuilder) Sgio(sgio ScsiGenericIO) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Sgio = sgio
	return builder
}

func (builder *diskBuilder) Shareable(shareable bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Shareable = &shareable
	return builder
}

func (builder *diskBuilder) Snapshot(snapshot *Snapshot) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Snapshot = snapshot
	return builder
}

func (builder *diskBuilder) Sparse(sparse bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Sparse = &sparse
	return builder
}

func (builder *diskBuilder) Statistics(statistics []Statistic) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Statistics = statistics
	return builder
}

func (builder *diskBuilder) Status(status DiskStatus) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Status = status
	return builder
}

func (builder *diskBuilder) StorageDomain(storageDomain *StorageDomain) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.StorageDomain = storageDomain
	return builder
}

func (builder *diskBuilder) StorageDomains(storageDomains []StorageDomain) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.StorageDomains = storageDomains
	return builder
}

func (builder *diskBuilder) StorageType(storageType DiskStorageType) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.StorageType = storageType
	return builder
}

func (builder *diskBuilder) Template(template *Template) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Template = template
	return builder
}

func (builder *diskBuilder) UsesScsiReservation(usesScsiReservation bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.UsesScsiReservation = &usesScsiReservation
	return builder
}

func (builder *diskBuilder) Vm(vm *Vm) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Vm = vm
	return builder
}

func (builder *diskBuilder) Vms(vms []Vm) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.Vms = vms
	return builder
}

func (builder *diskBuilder) WipeAfterDelete(wipeAfterDelete bool) *diskBuilder {
	if builder.err != nil {
		return builder
	}

	builder.disk.WipeAfterDelete = &wipeAfterDelete
	return builder
}

func (builder *diskBuilder) Build() (*Disk, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.disk, nil
}

type diskAttachmentBuilder struct {
	diskAttachment *DiskAttachment
	err            error
}

func NewDiskAttachmentBuilder() *diskAttachmentBuilder {
	return &diskAttachmentBuilder{diskAttachment: &DiskAttachment{}, err: nil}
}

func (builder *diskAttachmentBuilder) Active(active bool) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Active = &active
	return builder
}

func (builder *diskAttachmentBuilder) Bootable(bootable bool) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Bootable = &bootable
	return builder
}

func (builder *diskAttachmentBuilder) Comment(comment string) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Comment = &comment
	return builder
}

func (builder *diskAttachmentBuilder) Description(description string) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Description = &description
	return builder
}

func (builder *diskAttachmentBuilder) Disk(disk *Disk) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Disk = disk
	return builder
}

func (builder *diskAttachmentBuilder) Id(id string) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Id = &id
	return builder
}

func (builder *diskAttachmentBuilder) Interface(interface_ DiskInterface) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Interface = interface_
	return builder
}

func (builder *diskAttachmentBuilder) LogicalName(logicalName string) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.LogicalName = &logicalName
	return builder
}

func (builder *diskAttachmentBuilder) Name(name string) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Name = &name
	return builder
}

func (builder *diskAttachmentBuilder) PassDiscard(passDiscard bool) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.PassDiscard = &passDiscard
	return builder
}

func (builder *diskAttachmentBuilder) Template(template *Template) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Template = template
	return builder
}

func (builder *diskAttachmentBuilder) UsesScsiReservation(usesScsiReservation bool) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.UsesScsiReservation = &usesScsiReservation
	return builder
}

func (builder *diskAttachmentBuilder) Vm(vm *Vm) *diskAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskAttachment.Vm = vm
	return builder
}

func (builder *diskAttachmentBuilder) Build() (*DiskAttachment, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.diskAttachment, nil
}

type diskProfileBuilder struct {
	diskProfile *DiskProfile
	err         error
}

func NewDiskProfileBuilder() *diskProfileBuilder {
	return &diskProfileBuilder{diskProfile: &DiskProfile{}, err: nil}
}

func (builder *diskProfileBuilder) Comment(comment string) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.Comment = &comment
	return builder
}

func (builder *diskProfileBuilder) Description(description string) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.Description = &description
	return builder
}

func (builder *diskProfileBuilder) Id(id string) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.Id = &id
	return builder
}

func (builder *diskProfileBuilder) Name(name string) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.Name = &name
	return builder
}

func (builder *diskProfileBuilder) Permissions(permissions []Permission) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.Permissions = permissions
	return builder
}

func (builder *diskProfileBuilder) Qos(qos *Qos) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.Qos = qos
	return builder
}

func (builder *diskProfileBuilder) StorageDomain(storageDomain *StorageDomain) *diskProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskProfile.StorageDomain = storageDomain
	return builder
}

func (builder *diskProfileBuilder) Build() (*DiskProfile, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.diskProfile, nil
}

type diskSnapshotBuilder struct {
	diskSnapshot *DiskSnapshot
	err          error
}

func NewDiskSnapshotBuilder() *diskSnapshotBuilder {
	return &diskSnapshotBuilder{diskSnapshot: &DiskSnapshot{}, err: nil}
}

func (builder *diskSnapshotBuilder) Active(active bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Active = &active
	return builder
}

func (builder *diskSnapshotBuilder) ActualSize(actualSize int64) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.ActualSize = &actualSize
	return builder
}

func (builder *diskSnapshotBuilder) Alias(alias string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Alias = &alias
	return builder
}

func (builder *diskSnapshotBuilder) Bootable(bootable bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Bootable = &bootable
	return builder
}

func (builder *diskSnapshotBuilder) Comment(comment string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Comment = &comment
	return builder
}

func (builder *diskSnapshotBuilder) Description(description string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Description = &description
	return builder
}

func (builder *diskSnapshotBuilder) Disk(disk *Disk) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Disk = disk
	return builder
}

func (builder *diskSnapshotBuilder) DiskProfile(diskProfile *DiskProfile) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.DiskProfile = diskProfile
	return builder
}

func (builder *diskSnapshotBuilder) Format(format DiskFormat) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Format = format
	return builder
}

func (builder *diskSnapshotBuilder) Id(id string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Id = &id
	return builder
}

func (builder *diskSnapshotBuilder) ImageId(imageId string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.ImageId = &imageId
	return builder
}

func (builder *diskSnapshotBuilder) InitialSize(initialSize int64) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.InitialSize = &initialSize
	return builder
}

func (builder *diskSnapshotBuilder) InstanceType(instanceType *InstanceType) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.InstanceType = instanceType
	return builder
}

func (builder *diskSnapshotBuilder) Interface(interface_ DiskInterface) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Interface = interface_
	return builder
}

func (builder *diskSnapshotBuilder) LogicalName(logicalName string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.LogicalName = &logicalName
	return builder
}

func (builder *diskSnapshotBuilder) LunStorage(lunStorage *HostStorage) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.LunStorage = lunStorage
	return builder
}

func (builder *diskSnapshotBuilder) Name(name string) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Name = &name
	return builder
}

func (builder *diskSnapshotBuilder) OpenstackVolumeType(openstackVolumeType *OpenStackVolumeType) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.OpenstackVolumeType = openstackVolumeType
	return builder
}

func (builder *diskSnapshotBuilder) Permissions(permissions []Permission) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Permissions = permissions
	return builder
}

func (builder *diskSnapshotBuilder) PropagateErrors(propagateErrors bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.PropagateErrors = &propagateErrors
	return builder
}

func (builder *diskSnapshotBuilder) ProvisionedSize(provisionedSize int64) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.ProvisionedSize = &provisionedSize
	return builder
}

func (builder *diskSnapshotBuilder) QcowVersion(qcowVersion QcowVersion) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.QcowVersion = qcowVersion
	return builder
}

func (builder *diskSnapshotBuilder) Quota(quota *Quota) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Quota = quota
	return builder
}

func (builder *diskSnapshotBuilder) ReadOnly(readOnly bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.ReadOnly = &readOnly
	return builder
}

func (builder *diskSnapshotBuilder) Sgio(sgio ScsiGenericIO) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Sgio = sgio
	return builder
}

func (builder *diskSnapshotBuilder) Shareable(shareable bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Shareable = &shareable
	return builder
}

func (builder *diskSnapshotBuilder) Snapshot(snapshot *Snapshot) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Snapshot = snapshot
	return builder
}

func (builder *diskSnapshotBuilder) Sparse(sparse bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Sparse = &sparse
	return builder
}

func (builder *diskSnapshotBuilder) Statistics(statistics []Statistic) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Statistics = statistics
	return builder
}

func (builder *diskSnapshotBuilder) Status(status DiskStatus) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Status = status
	return builder
}

func (builder *diskSnapshotBuilder) StorageDomain(storageDomain *StorageDomain) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.StorageDomain = storageDomain
	return builder
}

func (builder *diskSnapshotBuilder) StorageDomains(storageDomains []StorageDomain) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.StorageDomains = storageDomains
	return builder
}

func (builder *diskSnapshotBuilder) StorageType(storageType DiskStorageType) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.StorageType = storageType
	return builder
}

func (builder *diskSnapshotBuilder) Template(template *Template) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Template = template
	return builder
}

func (builder *diskSnapshotBuilder) UsesScsiReservation(usesScsiReservation bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.UsesScsiReservation = &usesScsiReservation
	return builder
}

func (builder *diskSnapshotBuilder) Vm(vm *Vm) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Vm = vm
	return builder
}

func (builder *diskSnapshotBuilder) Vms(vms []Vm) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.Vms = vms
	return builder
}

func (builder *diskSnapshotBuilder) WipeAfterDelete(wipeAfterDelete bool) *diskSnapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.diskSnapshot.WipeAfterDelete = &wipeAfterDelete
	return builder
}

func (builder *diskSnapshotBuilder) Build() (*DiskSnapshot, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.diskSnapshot, nil
}

type displayBuilder struct {
	display *Display
	err     error
}

func NewDisplayBuilder() *displayBuilder {
	return &displayBuilder{display: &Display{}, err: nil}
}

func (builder *displayBuilder) Address(address string) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.Address = &address
	return builder
}

func (builder *displayBuilder) AllowOverride(allowOverride bool) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.AllowOverride = &allowOverride
	return builder
}

func (builder *displayBuilder) Certificate(certificate *Certificate) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.Certificate = certificate
	return builder
}

func (builder *displayBuilder) CopyPasteEnabled(copyPasteEnabled bool) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.CopyPasteEnabled = &copyPasteEnabled
	return builder
}

func (builder *displayBuilder) DisconnectAction(disconnectAction string) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.DisconnectAction = &disconnectAction
	return builder
}

func (builder *displayBuilder) FileTransferEnabled(fileTransferEnabled bool) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.FileTransferEnabled = &fileTransferEnabled
	return builder
}

func (builder *displayBuilder) KeyboardLayout(keyboardLayout string) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.KeyboardLayout = &keyboardLayout
	return builder
}

func (builder *displayBuilder) Monitors(monitors int64) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.Monitors = &monitors
	return builder
}

func (builder *displayBuilder) Port(port int64) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.Port = &port
	return builder
}

func (builder *displayBuilder) Proxy(proxy string) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.Proxy = &proxy
	return builder
}

func (builder *displayBuilder) SecurePort(securePort int64) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.SecurePort = &securePort
	return builder
}

func (builder *displayBuilder) SingleQxlPci(singleQxlPci bool) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.SingleQxlPci = &singleQxlPci
	return builder
}

func (builder *displayBuilder) SmartcardEnabled(smartcardEnabled bool) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.SmartcardEnabled = &smartcardEnabled
	return builder
}

func (builder *displayBuilder) Type(type_ DisplayType) *displayBuilder {
	if builder.err != nil {
		return builder
	}

	builder.display.Type = type_
	return builder
}

func (builder *displayBuilder) Build() (*Display, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.display, nil
}

type dnsBuilder struct {
	dns *Dns
	err error
}

func NewDnsBuilder() *dnsBuilder {
	return &dnsBuilder{dns: &Dns{}, err: nil}
}

func (builder *dnsBuilder) SearchDomains(searchDomains []Host) *dnsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dns.SearchDomains = searchDomains
	return builder
}

func (builder *dnsBuilder) Servers(servers []Host) *dnsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dns.Servers = servers
	return builder
}

func (builder *dnsBuilder) Build() (*Dns, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.dns, nil
}

type dnsResolverConfigurationBuilder struct {
	dnsResolverConfiguration *DnsResolverConfiguration
	err                      error
}

func NewDnsResolverConfigurationBuilder() *dnsResolverConfigurationBuilder {
	return &dnsResolverConfigurationBuilder{dnsResolverConfiguration: &DnsResolverConfiguration{}, err: nil}
}

func (builder *dnsResolverConfigurationBuilder) NameServers(nameServers []string) *dnsResolverConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.dnsResolverConfiguration.NameServers = nameServers
	return builder
}

func (builder *dnsResolverConfigurationBuilder) Build() (*DnsResolverConfiguration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.dnsResolverConfiguration, nil
}

type domainBuilder struct {
	domain *Domain
	err    error
}

func NewDomainBuilder() *domainBuilder {
	return &domainBuilder{domain: &Domain{}, err: nil}
}

func (builder *domainBuilder) Comment(comment string) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.Comment = &comment
	return builder
}

func (builder *domainBuilder) Description(description string) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.Description = &description
	return builder
}

func (builder *domainBuilder) Groups(groups []Group) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.Groups = groups
	return builder
}

func (builder *domainBuilder) Id(id string) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.Id = &id
	return builder
}

func (builder *domainBuilder) Name(name string) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.Name = &name
	return builder
}

func (builder *domainBuilder) User(user *User) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.User = user
	return builder
}

func (builder *domainBuilder) Users(users []User) *domainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.domain.Users = users
	return builder
}

func (builder *domainBuilder) Build() (*Domain, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.domain, nil
}

type entityProfileDetailBuilder struct {
	entityProfileDetail *EntityProfileDetail
	err                 error
}

func NewEntityProfileDetailBuilder() *entityProfileDetailBuilder {
	return &entityProfileDetailBuilder{entityProfileDetail: &EntityProfileDetail{}, err: nil}
}

func (builder *entityProfileDetailBuilder) ProfileDetails(profileDetails []ProfileDetail) *entityProfileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.entityProfileDetail.ProfileDetails = profileDetails
	return builder
}

func (builder *entityProfileDetailBuilder) Build() (*EntityProfileDetail, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.entityProfileDetail, nil
}

type errorHandlingBuilder struct {
	errorHandling *ErrorHandling
	err           error
}

func NewErrorHandlingBuilder() *errorHandlingBuilder {
	return &errorHandlingBuilder{errorHandling: &ErrorHandling{}, err: nil}
}

func (builder *errorHandlingBuilder) OnError(onError MigrateOnError) *errorHandlingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.errorHandling.OnError = onError
	return builder
}

func (builder *errorHandlingBuilder) Build() (*ErrorHandling, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.errorHandling, nil
}

type eventBuilder struct {
	event *Event
	err   error
}

func NewEventBuilder() *eventBuilder {
	return &eventBuilder{event: &Event{}, err: nil}
}

func (builder *eventBuilder) Cluster(cluster *Cluster) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Cluster = cluster
	return builder
}

func (builder *eventBuilder) Code(code int64) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Code = &code
	return builder
}

func (builder *eventBuilder) Comment(comment string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Comment = &comment
	return builder
}

func (builder *eventBuilder) CorrelationId(correlationId string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.CorrelationId = &correlationId
	return builder
}

func (builder *eventBuilder) CustomData(customData string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.CustomData = &customData
	return builder
}

func (builder *eventBuilder) CustomId(customId int64) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.CustomId = &customId
	return builder
}

func (builder *eventBuilder) DataCenter(dataCenter *DataCenter) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.DataCenter = dataCenter
	return builder
}

func (builder *eventBuilder) Description(description string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Description = &description
	return builder
}

func (builder *eventBuilder) FloodRate(floodRate int64) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.FloodRate = &floodRate
	return builder
}

func (builder *eventBuilder) Host(host *Host) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Host = host
	return builder
}

func (builder *eventBuilder) Id(id string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Id = &id
	return builder
}

func (builder *eventBuilder) Name(name string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Name = &name
	return builder
}

func (builder *eventBuilder) Origin(origin string) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Origin = &origin
	return builder
}

func (builder *eventBuilder) Severity(severity LogSeverity) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Severity = severity
	return builder
}

func (builder *eventBuilder) StorageDomain(storageDomain *StorageDomain) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.StorageDomain = storageDomain
	return builder
}

func (builder *eventBuilder) Template(template *Template) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Template = template
	return builder
}

func (builder *eventBuilder) Time(time time.Time) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Time = time
	return builder
}

func (builder *eventBuilder) User(user *User) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.User = user
	return builder
}

func (builder *eventBuilder) Vm(vm *Vm) *eventBuilder {
	if builder.err != nil {
		return builder
	}

	builder.event.Vm = vm
	return builder
}

func (builder *eventBuilder) Build() (*Event, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.event, nil
}

type externalComputeResourceBuilder struct {
	externalComputeResource *ExternalComputeResource
	err                     error
}

func NewExternalComputeResourceBuilder() *externalComputeResourceBuilder {
	return &externalComputeResourceBuilder{externalComputeResource: &ExternalComputeResource{}, err: nil}
}

func (builder *externalComputeResourceBuilder) Comment(comment string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.Comment = &comment
	return builder
}

func (builder *externalComputeResourceBuilder) Description(description string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.Description = &description
	return builder
}

func (builder *externalComputeResourceBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *externalComputeResourceBuilder) Id(id string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.Id = &id
	return builder
}

func (builder *externalComputeResourceBuilder) Name(name string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.Name = &name
	return builder
}

func (builder *externalComputeResourceBuilder) Provider(provider string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.Provider = &provider
	return builder
}

func (builder *externalComputeResourceBuilder) Url(url string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.Url = &url
	return builder
}

func (builder *externalComputeResourceBuilder) User(user string) *externalComputeResourceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalComputeResource.User = &user
	return builder
}

func (builder *externalComputeResourceBuilder) Build() (*ExternalComputeResource, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalComputeResource, nil
}

type externalDiscoveredHostBuilder struct {
	externalDiscoveredHost *ExternalDiscoveredHost
	err                    error
}

func NewExternalDiscoveredHostBuilder() *externalDiscoveredHostBuilder {
	return &externalDiscoveredHostBuilder{externalDiscoveredHost: &ExternalDiscoveredHost{}, err: nil}
}

func (builder *externalDiscoveredHostBuilder) Comment(comment string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.Comment = &comment
	return builder
}

func (builder *externalDiscoveredHostBuilder) Description(description string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.Description = &description
	return builder
}

func (builder *externalDiscoveredHostBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *externalDiscoveredHostBuilder) Id(id string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.Id = &id
	return builder
}

func (builder *externalDiscoveredHostBuilder) Ip(ip string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.Ip = &ip
	return builder
}

func (builder *externalDiscoveredHostBuilder) LastReport(lastReport string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.LastReport = &lastReport
	return builder
}

func (builder *externalDiscoveredHostBuilder) Mac(mac string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.Mac = &mac
	return builder
}

func (builder *externalDiscoveredHostBuilder) Name(name string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.Name = &name
	return builder
}

func (builder *externalDiscoveredHostBuilder) SubnetName(subnetName string) *externalDiscoveredHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalDiscoveredHost.SubnetName = &subnetName
	return builder
}

func (builder *externalDiscoveredHostBuilder) Build() (*ExternalDiscoveredHost, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalDiscoveredHost, nil
}

type externalHostBuilder struct {
	externalHost *ExternalHost
	err          error
}

func NewExternalHostBuilder() *externalHostBuilder {
	return &externalHostBuilder{externalHost: &ExternalHost{}, err: nil}
}

func (builder *externalHostBuilder) Address(address string) *externalHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHost.Address = &address
	return builder
}

func (builder *externalHostBuilder) Comment(comment string) *externalHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHost.Comment = &comment
	return builder
}

func (builder *externalHostBuilder) Description(description string) *externalHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHost.Description = &description
	return builder
}

func (builder *externalHostBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *externalHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHost.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *externalHostBuilder) Id(id string) *externalHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHost.Id = &id
	return builder
}

func (builder *externalHostBuilder) Name(name string) *externalHostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHost.Name = &name
	return builder
}

func (builder *externalHostBuilder) Build() (*ExternalHost, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalHost, nil
}

type externalHostGroupBuilder struct {
	externalHostGroup *ExternalHostGroup
	err               error
}

func NewExternalHostGroupBuilder() *externalHostGroupBuilder {
	return &externalHostGroupBuilder{externalHostGroup: &ExternalHostGroup{}, err: nil}
}

func (builder *externalHostGroupBuilder) ArchitectureName(architectureName string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.ArchitectureName = &architectureName
	return builder
}

func (builder *externalHostGroupBuilder) Comment(comment string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.Comment = &comment
	return builder
}

func (builder *externalHostGroupBuilder) Description(description string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.Description = &description
	return builder
}

func (builder *externalHostGroupBuilder) DomainName(domainName string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.DomainName = &domainName
	return builder
}

func (builder *externalHostGroupBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *externalHostGroupBuilder) Id(id string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.Id = &id
	return builder
}

func (builder *externalHostGroupBuilder) Name(name string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.Name = &name
	return builder
}

func (builder *externalHostGroupBuilder) OperatingSystemName(operatingSystemName string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.OperatingSystemName = &operatingSystemName
	return builder
}

func (builder *externalHostGroupBuilder) SubnetName(subnetName string) *externalHostGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostGroup.SubnetName = &subnetName
	return builder
}

func (builder *externalHostGroupBuilder) Build() (*ExternalHostGroup, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalHostGroup, nil
}

type externalHostProviderBuilder struct {
	externalHostProvider *ExternalHostProvider
	err                  error
}

func NewExternalHostProviderBuilder() *externalHostProviderBuilder {
	return &externalHostProviderBuilder{externalHostProvider: &ExternalHostProvider{}, err: nil}
}

func (builder *externalHostProviderBuilder) AuthenticationUrl(authenticationUrl string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.AuthenticationUrl = &authenticationUrl
	return builder
}

func (builder *externalHostProviderBuilder) Certificates(certificates []Certificate) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Certificates = certificates
	return builder
}

func (builder *externalHostProviderBuilder) Comment(comment string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Comment = &comment
	return builder
}

func (builder *externalHostProviderBuilder) ComputeResources(computeResources []ExternalComputeResource) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.ComputeResources = computeResources
	return builder
}

func (builder *externalHostProviderBuilder) Description(description string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Description = &description
	return builder
}

func (builder *externalHostProviderBuilder) DiscoveredHosts(discoveredHosts []ExternalDiscoveredHost) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.DiscoveredHosts = discoveredHosts
	return builder
}

func (builder *externalHostProviderBuilder) HostGroups(hostGroups []ExternalHostGroup) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.HostGroups = hostGroups
	return builder
}

func (builder *externalHostProviderBuilder) Hosts(hosts []Host) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Hosts = hosts
	return builder
}

func (builder *externalHostProviderBuilder) Id(id string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Id = &id
	return builder
}

func (builder *externalHostProviderBuilder) Name(name string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Name = &name
	return builder
}

func (builder *externalHostProviderBuilder) Password(password string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Password = &password
	return builder
}

func (builder *externalHostProviderBuilder) Properties(properties []Property) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Properties = properties
	return builder
}

func (builder *externalHostProviderBuilder) RequiresAuthentication(requiresAuthentication bool) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.RequiresAuthentication = &requiresAuthentication
	return builder
}

func (builder *externalHostProviderBuilder) Url(url string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Url = &url
	return builder
}

func (builder *externalHostProviderBuilder) Username(username string) *externalHostProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalHostProvider.Username = &username
	return builder
}

func (builder *externalHostProviderBuilder) Build() (*ExternalHostProvider, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalHostProvider, nil
}

type externalProviderBuilder struct {
	externalProvider *ExternalProvider
	err              error
}

func NewExternalProviderBuilder() *externalProviderBuilder {
	return &externalProviderBuilder{externalProvider: &ExternalProvider{}, err: nil}
}

func (builder *externalProviderBuilder) AuthenticationUrl(authenticationUrl string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.AuthenticationUrl = &authenticationUrl
	return builder
}

func (builder *externalProviderBuilder) Comment(comment string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Comment = &comment
	return builder
}

func (builder *externalProviderBuilder) Description(description string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Description = &description
	return builder
}

func (builder *externalProviderBuilder) Id(id string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Id = &id
	return builder
}

func (builder *externalProviderBuilder) Name(name string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Name = &name
	return builder
}

func (builder *externalProviderBuilder) Password(password string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Password = &password
	return builder
}

func (builder *externalProviderBuilder) Properties(properties []Property) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Properties = properties
	return builder
}

func (builder *externalProviderBuilder) RequiresAuthentication(requiresAuthentication bool) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.RequiresAuthentication = &requiresAuthentication
	return builder
}

func (builder *externalProviderBuilder) Url(url string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Url = &url
	return builder
}

func (builder *externalProviderBuilder) Username(username string) *externalProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalProvider.Username = &username
	return builder
}

func (builder *externalProviderBuilder) Build() (*ExternalProvider, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalProvider, nil
}

type externalVmImportBuilder struct {
	externalVmImport *ExternalVmImport
	err              error
}

func NewExternalVmImportBuilder() *externalVmImportBuilder {
	return &externalVmImportBuilder{externalVmImport: &ExternalVmImport{}, err: nil}
}

func (builder *externalVmImportBuilder) Cluster(cluster *Cluster) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Cluster = cluster
	return builder
}

func (builder *externalVmImportBuilder) CpuProfile(cpuProfile *CpuProfile) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.CpuProfile = cpuProfile
	return builder
}

func (builder *externalVmImportBuilder) DriversIso(driversIso *File) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.DriversIso = driversIso
	return builder
}

func (builder *externalVmImportBuilder) Host(host *Host) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Host = host
	return builder
}

func (builder *externalVmImportBuilder) Name(name string) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Name = &name
	return builder
}

func (builder *externalVmImportBuilder) Password(password string) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Password = &password
	return builder
}

func (builder *externalVmImportBuilder) Provider(provider ExternalVmProviderType) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Provider = provider
	return builder
}

func (builder *externalVmImportBuilder) Quota(quota *Quota) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Quota = quota
	return builder
}

func (builder *externalVmImportBuilder) Sparse(sparse bool) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Sparse = &sparse
	return builder
}

func (builder *externalVmImportBuilder) StorageDomain(storageDomain *StorageDomain) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.StorageDomain = storageDomain
	return builder
}

func (builder *externalVmImportBuilder) Url(url string) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Url = &url
	return builder
}

func (builder *externalVmImportBuilder) Username(username string) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Username = &username
	return builder
}

func (builder *externalVmImportBuilder) Vm(vm *Vm) *externalVmImportBuilder {
	if builder.err != nil {
		return builder
	}

	builder.externalVmImport.Vm = vm
	return builder
}

func (builder *externalVmImportBuilder) Build() (*ExternalVmImport, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.externalVmImport, nil
}

type faultBuilder struct {
	fault *Fault
	err   error
}

func NewFaultBuilder() *faultBuilder {
	return &faultBuilder{fault: &Fault{}, err: nil}
}

func (builder *faultBuilder) Detail(detail string) *faultBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fault.Detail = &detail
	return builder
}

func (builder *faultBuilder) Reason(reason string) *faultBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fault.Reason = &reason
	return builder
}

func (builder *faultBuilder) Build() (*Fault, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.fault, nil
}

type fencingPolicyBuilder struct {
	fencingPolicy *FencingPolicy
	err           error
}

func NewFencingPolicyBuilder() *fencingPolicyBuilder {
	return &fencingPolicyBuilder{fencingPolicy: &FencingPolicy{}, err: nil}
}

func (builder *fencingPolicyBuilder) Enabled(enabled bool) *fencingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fencingPolicy.Enabled = &enabled
	return builder
}

func (builder *fencingPolicyBuilder) SkipIfConnectivityBroken(skipIfConnectivityBroken *SkipIfConnectivityBroken) *fencingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fencingPolicy.SkipIfConnectivityBroken = skipIfConnectivityBroken
	return builder
}

func (builder *fencingPolicyBuilder) SkipIfGlusterBricksUp(skipIfGlusterBricksUp bool) *fencingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fencingPolicy.SkipIfGlusterBricksUp = &skipIfGlusterBricksUp
	return builder
}

func (builder *fencingPolicyBuilder) SkipIfGlusterQuorumNotMet(skipIfGlusterQuorumNotMet bool) *fencingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fencingPolicy.SkipIfGlusterQuorumNotMet = &skipIfGlusterQuorumNotMet
	return builder
}

func (builder *fencingPolicyBuilder) SkipIfSdActive(skipIfSdActive *SkipIfSdActive) *fencingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fencingPolicy.SkipIfSdActive = skipIfSdActive
	return builder
}

func (builder *fencingPolicyBuilder) Build() (*FencingPolicy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.fencingPolicy, nil
}

type fileBuilder struct {
	file *File
	err  error
}

func NewFileBuilder() *fileBuilder {
	return &fileBuilder{file: &File{}, err: nil}
}

func (builder *fileBuilder) Comment(comment string) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.Comment = &comment
	return builder
}

func (builder *fileBuilder) Content(content string) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.Content = &content
	return builder
}

func (builder *fileBuilder) Description(description string) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.Description = &description
	return builder
}

func (builder *fileBuilder) Id(id string) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.Id = &id
	return builder
}

func (builder *fileBuilder) Name(name string) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.Name = &name
	return builder
}

func (builder *fileBuilder) StorageDomain(storageDomain *StorageDomain) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.StorageDomain = storageDomain
	return builder
}

func (builder *fileBuilder) Type(type_ string) *fileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.file.Type = &type_
	return builder
}

func (builder *fileBuilder) Build() (*File, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.file, nil
}

type filterBuilder struct {
	filter *Filter
	err    error
}

func NewFilterBuilder() *filterBuilder {
	return &filterBuilder{filter: &Filter{}, err: nil}
}

func (builder *filterBuilder) Comment(comment string) *filterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.filter.Comment = &comment
	return builder
}

func (builder *filterBuilder) Description(description string) *filterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.filter.Description = &description
	return builder
}

func (builder *filterBuilder) Id(id string) *filterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.filter.Id = &id
	return builder
}

func (builder *filterBuilder) Name(name string) *filterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.filter.Name = &name
	return builder
}

func (builder *filterBuilder) Position(position int64) *filterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.filter.Position = &position
	return builder
}

func (builder *filterBuilder) SchedulingPolicyUnit(schedulingPolicyUnit *SchedulingPolicyUnit) *filterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.filter.SchedulingPolicyUnit = schedulingPolicyUnit
	return builder
}

func (builder *filterBuilder) Build() (*Filter, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.filter, nil
}

type floppyBuilder struct {
	floppy *Floppy
	err    error
}

func NewFloppyBuilder() *floppyBuilder {
	return &floppyBuilder{floppy: &Floppy{}, err: nil}
}

func (builder *floppyBuilder) Comment(comment string) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Comment = &comment
	return builder
}

func (builder *floppyBuilder) Description(description string) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Description = &description
	return builder
}

func (builder *floppyBuilder) File(file *File) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.File = file
	return builder
}

func (builder *floppyBuilder) Id(id string) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Id = &id
	return builder
}

func (builder *floppyBuilder) InstanceType(instanceType *InstanceType) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.InstanceType = instanceType
	return builder
}

func (builder *floppyBuilder) Name(name string) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Name = &name
	return builder
}

func (builder *floppyBuilder) Template(template *Template) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Template = template
	return builder
}

func (builder *floppyBuilder) Vm(vm *Vm) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Vm = vm
	return builder
}

func (builder *floppyBuilder) Vms(vms []Vm) *floppyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.floppy.Vms = vms
	return builder
}

func (builder *floppyBuilder) Build() (*Floppy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.floppy, nil
}

type fopStatisticBuilder struct {
	fopStatistic *FopStatistic
	err          error
}

func NewFopStatisticBuilder() *fopStatisticBuilder {
	return &fopStatisticBuilder{fopStatistic: &FopStatistic{}, err: nil}
}

func (builder *fopStatisticBuilder) Name(name string) *fopStatisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fopStatistic.Name = &name
	return builder
}

func (builder *fopStatisticBuilder) Statistics(statistics []Statistic) *fopStatisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.fopStatistic.Statistics = statistics
	return builder
}

func (builder *fopStatisticBuilder) Build() (*FopStatistic, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.fopStatistic, nil
}

type glusterBrickBuilder struct {
	glusterBrick *GlusterBrick
	err          error
}

func NewGlusterBrickBuilder() *glusterBrickBuilder {
	return &glusterBrickBuilder{glusterBrick: &GlusterBrick{}, err: nil}
}

func (builder *glusterBrickBuilder) BrickDir(brickDir string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.BrickDir = &brickDir
	return builder
}

func (builder *glusterBrickBuilder) Comment(comment string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Comment = &comment
	return builder
}

func (builder *glusterBrickBuilder) Description(description string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Description = &description
	return builder
}

func (builder *glusterBrickBuilder) Device(device string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Device = &device
	return builder
}

func (builder *glusterBrickBuilder) FsName(fsName string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.FsName = &fsName
	return builder
}

func (builder *glusterBrickBuilder) GlusterClients(glusterClients []GlusterClient) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.GlusterClients = glusterClients
	return builder
}

func (builder *glusterBrickBuilder) GlusterVolume(glusterVolume *GlusterVolume) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.GlusterVolume = glusterVolume
	return builder
}

func (builder *glusterBrickBuilder) Id(id string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Id = &id
	return builder
}

func (builder *glusterBrickBuilder) InstanceType(instanceType *InstanceType) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.InstanceType = instanceType
	return builder
}

func (builder *glusterBrickBuilder) MemoryPools(memoryPools []GlusterMemoryPool) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.MemoryPools = memoryPools
	return builder
}

func (builder *glusterBrickBuilder) MntOptions(mntOptions string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.MntOptions = &mntOptions
	return builder
}

func (builder *glusterBrickBuilder) Name(name string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Name = &name
	return builder
}

func (builder *glusterBrickBuilder) Pid(pid int64) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Pid = &pid
	return builder
}

func (builder *glusterBrickBuilder) Port(port int64) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Port = &port
	return builder
}

func (builder *glusterBrickBuilder) ServerId(serverId string) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.ServerId = &serverId
	return builder
}

func (builder *glusterBrickBuilder) Statistics(statistics []Statistic) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Statistics = statistics
	return builder
}

func (builder *glusterBrickBuilder) Status(status GlusterBrickStatus) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Status = status
	return builder
}

func (builder *glusterBrickBuilder) Template(template *Template) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Template = template
	return builder
}

func (builder *glusterBrickBuilder) Vm(vm *Vm) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Vm = vm
	return builder
}

func (builder *glusterBrickBuilder) Vms(vms []Vm) *glusterBrickBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrick.Vms = vms
	return builder
}

func (builder *glusterBrickBuilder) Build() (*GlusterBrick, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterBrick, nil
}

type glusterBrickAdvancedDetailsBuilder struct {
	glusterBrickAdvancedDetails *GlusterBrickAdvancedDetails
	err                         error
}

func NewGlusterBrickAdvancedDetailsBuilder() *glusterBrickAdvancedDetailsBuilder {
	return &glusterBrickAdvancedDetailsBuilder{glusterBrickAdvancedDetails: &GlusterBrickAdvancedDetails{}, err: nil}
}

func (builder *glusterBrickAdvancedDetailsBuilder) Comment(comment string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Comment = &comment
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Description(description string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Description = &description
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Device(device string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Device = &device
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) FsName(fsName string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.FsName = &fsName
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) GlusterClients(glusterClients []GlusterClient) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.GlusterClients = glusterClients
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Id(id string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Id = &id
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) InstanceType(instanceType *InstanceType) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.InstanceType = instanceType
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) MemoryPools(memoryPools []GlusterMemoryPool) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.MemoryPools = memoryPools
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) MntOptions(mntOptions string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.MntOptions = &mntOptions
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Name(name string) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Name = &name
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Pid(pid int64) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Pid = &pid
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Port(port int64) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Port = &port
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Template(template *Template) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Template = template
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Vm(vm *Vm) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Vm = vm
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Vms(vms []Vm) *glusterBrickAdvancedDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickAdvancedDetails.Vms = vms
	return builder
}

func (builder *glusterBrickAdvancedDetailsBuilder) Build() (*GlusterBrickAdvancedDetails, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterBrickAdvancedDetails, nil
}

type glusterBrickMemoryInfoBuilder struct {
	glusterBrickMemoryInfo *GlusterBrickMemoryInfo
	err                    error
}

func NewGlusterBrickMemoryInfoBuilder() *glusterBrickMemoryInfoBuilder {
	return &glusterBrickMemoryInfoBuilder{glusterBrickMemoryInfo: &GlusterBrickMemoryInfo{}, err: nil}
}

func (builder *glusterBrickMemoryInfoBuilder) MemoryPools(memoryPools []GlusterMemoryPool) *glusterBrickMemoryInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterBrickMemoryInfo.MemoryPools = memoryPools
	return builder
}

func (builder *glusterBrickMemoryInfoBuilder) Build() (*GlusterBrickMemoryInfo, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterBrickMemoryInfo, nil
}

type glusterClientBuilder struct {
	glusterClient *GlusterClient
	err           error
}

func NewGlusterClientBuilder() *glusterClientBuilder {
	return &glusterClientBuilder{glusterClient: &GlusterClient{}, err: nil}
}

func (builder *glusterClientBuilder) BytesRead(bytesRead int64) *glusterClientBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterClient.BytesRead = &bytesRead
	return builder
}

func (builder *glusterClientBuilder) BytesWritten(bytesWritten int64) *glusterClientBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterClient.BytesWritten = &bytesWritten
	return builder
}

func (builder *glusterClientBuilder) ClientPort(clientPort int64) *glusterClientBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterClient.ClientPort = &clientPort
	return builder
}

func (builder *glusterClientBuilder) HostName(hostName string) *glusterClientBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterClient.HostName = &hostName
	return builder
}

func (builder *glusterClientBuilder) Build() (*GlusterClient, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterClient, nil
}

type glusterHookBuilder struct {
	glusterHook *GlusterHook
	err         error
}

func NewGlusterHookBuilder() *glusterHookBuilder {
	return &glusterHookBuilder{glusterHook: &GlusterHook{}, err: nil}
}

func (builder *glusterHookBuilder) Checksum(checksum string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Checksum = &checksum
	return builder
}

func (builder *glusterHookBuilder) Cluster(cluster *Cluster) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Cluster = cluster
	return builder
}

func (builder *glusterHookBuilder) Comment(comment string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Comment = &comment
	return builder
}

func (builder *glusterHookBuilder) ConflictStatus(conflictStatus int64) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.ConflictStatus = &conflictStatus
	return builder
}

func (builder *glusterHookBuilder) Conflicts(conflicts string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Conflicts = &conflicts
	return builder
}

func (builder *glusterHookBuilder) Content(content string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Content = &content
	return builder
}

func (builder *glusterHookBuilder) ContentType(contentType HookContentType) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.ContentType = contentType
	return builder
}

func (builder *glusterHookBuilder) Description(description string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Description = &description
	return builder
}

func (builder *glusterHookBuilder) GlusterCommand(glusterCommand string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.GlusterCommand = &glusterCommand
	return builder
}

func (builder *glusterHookBuilder) Id(id string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Id = &id
	return builder
}

func (builder *glusterHookBuilder) Name(name string) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Name = &name
	return builder
}

func (builder *glusterHookBuilder) ServerHooks(serverHooks []GlusterServerHook) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.ServerHooks = serverHooks
	return builder
}

func (builder *glusterHookBuilder) Stage(stage HookStage) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Stage = stage
	return builder
}

func (builder *glusterHookBuilder) Status(status GlusterHookStatus) *glusterHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterHook.Status = status
	return builder
}

func (builder *glusterHookBuilder) Build() (*GlusterHook, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterHook, nil
}

type glusterMemoryPoolBuilder struct {
	glusterMemoryPool *GlusterMemoryPool
	err               error
}

func NewGlusterMemoryPoolBuilder() *glusterMemoryPoolBuilder {
	return &glusterMemoryPoolBuilder{glusterMemoryPool: &GlusterMemoryPool{}, err: nil}
}

func (builder *glusterMemoryPoolBuilder) AllocCount(allocCount int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.AllocCount = &allocCount
	return builder
}

func (builder *glusterMemoryPoolBuilder) ColdCount(coldCount int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.ColdCount = &coldCount
	return builder
}

func (builder *glusterMemoryPoolBuilder) Comment(comment string) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.Comment = &comment
	return builder
}

func (builder *glusterMemoryPoolBuilder) Description(description string) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.Description = &description
	return builder
}

func (builder *glusterMemoryPoolBuilder) HotCount(hotCount int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.HotCount = &hotCount
	return builder
}

func (builder *glusterMemoryPoolBuilder) Id(id string) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.Id = &id
	return builder
}

func (builder *glusterMemoryPoolBuilder) MaxAlloc(maxAlloc int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.MaxAlloc = &maxAlloc
	return builder
}

func (builder *glusterMemoryPoolBuilder) MaxStdalloc(maxStdalloc int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.MaxStdalloc = &maxStdalloc
	return builder
}

func (builder *glusterMemoryPoolBuilder) Name(name string) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.Name = &name
	return builder
}

func (builder *glusterMemoryPoolBuilder) PaddedSize(paddedSize int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.PaddedSize = &paddedSize
	return builder
}

func (builder *glusterMemoryPoolBuilder) PoolMisses(poolMisses int64) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.PoolMisses = &poolMisses
	return builder
}

func (builder *glusterMemoryPoolBuilder) Type(type_ string) *glusterMemoryPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterMemoryPool.Type = &type_
	return builder
}

func (builder *glusterMemoryPoolBuilder) Build() (*GlusterMemoryPool, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterMemoryPool, nil
}

type glusterServerHookBuilder struct {
	glusterServerHook *GlusterServerHook
	err               error
}

func NewGlusterServerHookBuilder() *glusterServerHookBuilder {
	return &glusterServerHookBuilder{glusterServerHook: &GlusterServerHook{}, err: nil}
}

func (builder *glusterServerHookBuilder) Checksum(checksum string) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Checksum = &checksum
	return builder
}

func (builder *glusterServerHookBuilder) Comment(comment string) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Comment = &comment
	return builder
}

func (builder *glusterServerHookBuilder) ContentType(contentType HookContentType) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.ContentType = contentType
	return builder
}

func (builder *glusterServerHookBuilder) Description(description string) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Description = &description
	return builder
}

func (builder *glusterServerHookBuilder) Host(host *Host) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Host = host
	return builder
}

func (builder *glusterServerHookBuilder) Id(id string) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Id = &id
	return builder
}

func (builder *glusterServerHookBuilder) Name(name string) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Name = &name
	return builder
}

func (builder *glusterServerHookBuilder) Status(status GlusterHookStatus) *glusterServerHookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterServerHook.Status = status
	return builder
}

func (builder *glusterServerHookBuilder) Build() (*GlusterServerHook, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterServerHook, nil
}

type glusterVolumeBuilder struct {
	glusterVolume *GlusterVolume
	err           error
}

func NewGlusterVolumeBuilder() *glusterVolumeBuilder {
	return &glusterVolumeBuilder{glusterVolume: &GlusterVolume{}, err: nil}
}

func (builder *glusterVolumeBuilder) Bricks(bricks []GlusterBrick) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Bricks = bricks
	return builder
}

func (builder *glusterVolumeBuilder) Cluster(cluster *Cluster) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Cluster = cluster
	return builder
}

func (builder *glusterVolumeBuilder) Comment(comment string) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Comment = &comment
	return builder
}

func (builder *glusterVolumeBuilder) Description(description string) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Description = &description
	return builder
}

func (builder *glusterVolumeBuilder) DisperseCount(disperseCount int64) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.DisperseCount = &disperseCount
	return builder
}

func (builder *glusterVolumeBuilder) Id(id string) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Id = &id
	return builder
}

func (builder *glusterVolumeBuilder) Name(name string) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Name = &name
	return builder
}

func (builder *glusterVolumeBuilder) Options(options []Option) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Options = options
	return builder
}

func (builder *glusterVolumeBuilder) RedundancyCount(redundancyCount int64) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.RedundancyCount = &redundancyCount
	return builder
}

func (builder *glusterVolumeBuilder) ReplicaCount(replicaCount int64) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.ReplicaCount = &replicaCount
	return builder
}

func (builder *glusterVolumeBuilder) Statistics(statistics []Statistic) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Statistics = statistics
	return builder
}

func (builder *glusterVolumeBuilder) Status(status GlusterVolumeStatus) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.Status = status
	return builder
}

func (builder *glusterVolumeBuilder) StripeCount(stripeCount int64) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.StripeCount = &stripeCount
	return builder
}

func (builder *glusterVolumeBuilder) TransportTypes(transportTypes []TransportType) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.TransportTypes = transportTypes
	return builder
}

func (builder *glusterVolumeBuilder) VolumeType(volumeType GlusterVolumeType) *glusterVolumeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolume.VolumeType = volumeType
	return builder
}

func (builder *glusterVolumeBuilder) Build() (*GlusterVolume, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterVolume, nil
}

type glusterVolumeProfileDetailsBuilder struct {
	glusterVolumeProfileDetails *GlusterVolumeProfileDetails
	err                         error
}

func NewGlusterVolumeProfileDetailsBuilder() *glusterVolumeProfileDetailsBuilder {
	return &glusterVolumeProfileDetailsBuilder{glusterVolumeProfileDetails: &GlusterVolumeProfileDetails{}, err: nil}
}

func (builder *glusterVolumeProfileDetailsBuilder) BrickProfileDetails(brickProfileDetails []BrickProfileDetail) *glusterVolumeProfileDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolumeProfileDetails.BrickProfileDetails = brickProfileDetails
	return builder
}

func (builder *glusterVolumeProfileDetailsBuilder) Comment(comment string) *glusterVolumeProfileDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolumeProfileDetails.Comment = &comment
	return builder
}

func (builder *glusterVolumeProfileDetailsBuilder) Description(description string) *glusterVolumeProfileDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolumeProfileDetails.Description = &description
	return builder
}

func (builder *glusterVolumeProfileDetailsBuilder) Id(id string) *glusterVolumeProfileDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolumeProfileDetails.Id = &id
	return builder
}

func (builder *glusterVolumeProfileDetailsBuilder) Name(name string) *glusterVolumeProfileDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolumeProfileDetails.Name = &name
	return builder
}

func (builder *glusterVolumeProfileDetailsBuilder) NfsProfileDetails(nfsProfileDetails []NfsProfileDetail) *glusterVolumeProfileDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.glusterVolumeProfileDetails.NfsProfileDetails = nfsProfileDetails
	return builder
}

func (builder *glusterVolumeProfileDetailsBuilder) Build() (*GlusterVolumeProfileDetails, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.glusterVolumeProfileDetails, nil
}

type gracePeriodBuilder struct {
	gracePeriod *GracePeriod
	err         error
}

func NewGracePeriodBuilder() *gracePeriodBuilder {
	return &gracePeriodBuilder{gracePeriod: &GracePeriod{}, err: nil}
}

func (builder *gracePeriodBuilder) Expiry(expiry int64) *gracePeriodBuilder {
	if builder.err != nil {
		return builder
	}

	builder.gracePeriod.Expiry = &expiry
	return builder
}

func (builder *gracePeriodBuilder) Build() (*GracePeriod, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.gracePeriod, nil
}

type graphicsConsoleBuilder struct {
	graphicsConsole *GraphicsConsole
	err             error
}

func NewGraphicsConsoleBuilder() *graphicsConsoleBuilder {
	return &graphicsConsoleBuilder{graphicsConsole: &GraphicsConsole{}, err: nil}
}

func (builder *graphicsConsoleBuilder) Address(address string) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Address = &address
	return builder
}

func (builder *graphicsConsoleBuilder) Comment(comment string) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Comment = &comment
	return builder
}

func (builder *graphicsConsoleBuilder) Description(description string) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Description = &description
	return builder
}

func (builder *graphicsConsoleBuilder) Id(id string) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Id = &id
	return builder
}

func (builder *graphicsConsoleBuilder) InstanceType(instanceType *InstanceType) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.InstanceType = instanceType
	return builder
}

func (builder *graphicsConsoleBuilder) Name(name string) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Name = &name
	return builder
}

func (builder *graphicsConsoleBuilder) Port(port int64) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Port = &port
	return builder
}

func (builder *graphicsConsoleBuilder) Protocol(protocol GraphicsType) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Protocol = protocol
	return builder
}

func (builder *graphicsConsoleBuilder) Template(template *Template) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Template = template
	return builder
}

func (builder *graphicsConsoleBuilder) TlsPort(tlsPort int64) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.TlsPort = &tlsPort
	return builder
}

func (builder *graphicsConsoleBuilder) Vm(vm *Vm) *graphicsConsoleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.graphicsConsole.Vm = vm
	return builder
}

func (builder *graphicsConsoleBuilder) Build() (*GraphicsConsole, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.graphicsConsole, nil
}

type groupBuilder struct {
	group *Group
	err   error
}

func NewGroupBuilder() *groupBuilder {
	return &groupBuilder{group: &Group{}, err: nil}
}

func (builder *groupBuilder) Comment(comment string) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Comment = &comment
	return builder
}

func (builder *groupBuilder) Description(description string) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Description = &description
	return builder
}

func (builder *groupBuilder) Domain(domain *Domain) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Domain = domain
	return builder
}

func (builder *groupBuilder) DomainEntryId(domainEntryId string) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.DomainEntryId = &domainEntryId
	return builder
}

func (builder *groupBuilder) Id(id string) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Id = &id
	return builder
}

func (builder *groupBuilder) Name(name string) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Name = &name
	return builder
}

func (builder *groupBuilder) Namespace(namespace string) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Namespace = &namespace
	return builder
}

func (builder *groupBuilder) Permissions(permissions []Permission) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Permissions = permissions
	return builder
}

func (builder *groupBuilder) Roles(roles []Role) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Roles = roles
	return builder
}

func (builder *groupBuilder) Tags(tags []Tag) *groupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.group.Tags = tags
	return builder
}

func (builder *groupBuilder) Build() (*Group, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.group, nil
}

type guestOperatingSystemBuilder struct {
	guestOperatingSystem *GuestOperatingSystem
	err                  error
}

func NewGuestOperatingSystemBuilder() *guestOperatingSystemBuilder {
	return &guestOperatingSystemBuilder{guestOperatingSystem: &GuestOperatingSystem{}, err: nil}
}

func (builder *guestOperatingSystemBuilder) Architecture(architecture string) *guestOperatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.guestOperatingSystem.Architecture = &architecture
	return builder
}

func (builder *guestOperatingSystemBuilder) Codename(codename string) *guestOperatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.guestOperatingSystem.Codename = &codename
	return builder
}

func (builder *guestOperatingSystemBuilder) Distribution(distribution string) *guestOperatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.guestOperatingSystem.Distribution = &distribution
	return builder
}

func (builder *guestOperatingSystemBuilder) Family(family string) *guestOperatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.guestOperatingSystem.Family = &family
	return builder
}

func (builder *guestOperatingSystemBuilder) Kernel(kernel *Kernel) *guestOperatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.guestOperatingSystem.Kernel = kernel
	return builder
}

func (builder *guestOperatingSystemBuilder) Version(version *Version) *guestOperatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.guestOperatingSystem.Version = version
	return builder
}

func (builder *guestOperatingSystemBuilder) Build() (*GuestOperatingSystem, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.guestOperatingSystem, nil
}

type hardwareInformationBuilder struct {
	hardwareInformation *HardwareInformation
	err                 error
}

func NewHardwareInformationBuilder() *hardwareInformationBuilder {
	return &hardwareInformationBuilder{hardwareInformation: &HardwareInformation{}, err: nil}
}

func (builder *hardwareInformationBuilder) Family(family string) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.Family = &family
	return builder
}

func (builder *hardwareInformationBuilder) Manufacturer(manufacturer string) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.Manufacturer = &manufacturer
	return builder
}

func (builder *hardwareInformationBuilder) ProductName(productName string) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.ProductName = &productName
	return builder
}

func (builder *hardwareInformationBuilder) SerialNumber(serialNumber string) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.SerialNumber = &serialNumber
	return builder
}

func (builder *hardwareInformationBuilder) SupportedRngSources(supportedRngSources []RngSource) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.SupportedRngSources = supportedRngSources
	return builder
}

func (builder *hardwareInformationBuilder) Uuid(uuid string) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.Uuid = &uuid
	return builder
}

func (builder *hardwareInformationBuilder) Version(version string) *hardwareInformationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hardwareInformation.Version = &version
	return builder
}

func (builder *hardwareInformationBuilder) Build() (*HardwareInformation, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hardwareInformation, nil
}

type highAvailabilityBuilder struct {
	highAvailability *HighAvailability
	err              error
}

func NewHighAvailabilityBuilder() *highAvailabilityBuilder {
	return &highAvailabilityBuilder{highAvailability: &HighAvailability{}, err: nil}
}

func (builder *highAvailabilityBuilder) Enabled(enabled bool) *highAvailabilityBuilder {
	if builder.err != nil {
		return builder
	}

	builder.highAvailability.Enabled = &enabled
	return builder
}

func (builder *highAvailabilityBuilder) Priority(priority int64) *highAvailabilityBuilder {
	if builder.err != nil {
		return builder
	}

	builder.highAvailability.Priority = &priority
	return builder
}

func (builder *highAvailabilityBuilder) Build() (*HighAvailability, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.highAvailability, nil
}

type hookBuilder struct {
	hook *Hook
	err  error
}

func NewHookBuilder() *hookBuilder {
	return &hookBuilder{hook: &Hook{}, err: nil}
}

func (builder *hookBuilder) Comment(comment string) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.Comment = &comment
	return builder
}

func (builder *hookBuilder) Description(description string) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.Description = &description
	return builder
}

func (builder *hookBuilder) EventName(eventName string) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.EventName = &eventName
	return builder
}

func (builder *hookBuilder) Host(host *Host) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.Host = host
	return builder
}

func (builder *hookBuilder) Id(id string) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.Id = &id
	return builder
}

func (builder *hookBuilder) Md5(md5 string) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.Md5 = &md5
	return builder
}

func (builder *hookBuilder) Name(name string) *hookBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hook.Name = &name
	return builder
}

func (builder *hookBuilder) Build() (*Hook, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hook, nil
}

type hostBuilder struct {
	host *Host
	err  error
}

func NewHostBuilder() *hostBuilder {
	return &hostBuilder{host: &Host{}, err: nil}
}

func (builder *hostBuilder) Address(address string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Address = &address
	return builder
}

func (builder *hostBuilder) AffinityLabels(affinityLabels []AffinityLabel) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.AffinityLabels = affinityLabels
	return builder
}

func (builder *hostBuilder) Agents(agents []Agent) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Agents = agents
	return builder
}

func (builder *hostBuilder) AutoNumaStatus(autoNumaStatus AutoNumaStatus) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.AutoNumaStatus = autoNumaStatus
	return builder
}

func (builder *hostBuilder) Certificate(certificate *Certificate) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Certificate = certificate
	return builder
}

func (builder *hostBuilder) Cluster(cluster *Cluster) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Cluster = cluster
	return builder
}

func (builder *hostBuilder) Comment(comment string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Comment = &comment
	return builder
}

func (builder *hostBuilder) Cpu(cpu *Cpu) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Cpu = cpu
	return builder
}

func (builder *hostBuilder) Description(description string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Description = &description
	return builder
}

func (builder *hostBuilder) DevicePassthrough(devicePassthrough *HostDevicePassthrough) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.DevicePassthrough = devicePassthrough
	return builder
}

func (builder *hostBuilder) Devices(devices []Device) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Devices = devices
	return builder
}

func (builder *hostBuilder) Display(display *Display) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Display = display
	return builder
}

func (builder *hostBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *hostBuilder) ExternalStatus(externalStatus ExternalStatus) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.ExternalStatus = externalStatus
	return builder
}

func (builder *hostBuilder) HardwareInformation(hardwareInformation *HardwareInformation) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.HardwareInformation = hardwareInformation
	return builder
}

func (builder *hostBuilder) Hooks(hooks []Hook) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Hooks = hooks
	return builder
}

func (builder *hostBuilder) HostedEngine(hostedEngine *HostedEngine) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.HostedEngine = hostedEngine
	return builder
}

func (builder *hostBuilder) Id(id string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Id = &id
	return builder
}

func (builder *hostBuilder) Iscsi(iscsi *IscsiDetails) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Iscsi = iscsi
	return builder
}

func (builder *hostBuilder) KatelloErrata(katelloErrata []KatelloErratum) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.KatelloErrata = katelloErrata
	return builder
}

func (builder *hostBuilder) KdumpStatus(kdumpStatus KdumpStatus) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.KdumpStatus = kdumpStatus
	return builder
}

func (builder *hostBuilder) Ksm(ksm *Ksm) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Ksm = ksm
	return builder
}

func (builder *hostBuilder) LibvirtVersion(libvirtVersion *Version) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.LibvirtVersion = libvirtVersion
	return builder
}

func (builder *hostBuilder) MaxSchedulingMemory(maxSchedulingMemory int64) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.MaxSchedulingMemory = &maxSchedulingMemory
	return builder
}

func (builder *hostBuilder) Memory(memory int64) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Memory = &memory
	return builder
}

func (builder *hostBuilder) Name(name string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Name = &name
	return builder
}

func (builder *hostBuilder) NetworkAttachments(networkAttachments []NetworkAttachment) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.NetworkAttachments = networkAttachments
	return builder
}

func (builder *hostBuilder) Nics(nics []Nic) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Nics = nics
	return builder
}

func (builder *hostBuilder) NumaNodes(numaNodes []NumaNode) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.NumaNodes = numaNodes
	return builder
}

func (builder *hostBuilder) NumaSupported(numaSupported bool) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.NumaSupported = &numaSupported
	return builder
}

func (builder *hostBuilder) Os(os *OperatingSystem) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Os = os
	return builder
}

func (builder *hostBuilder) OverrideIptables(overrideIptables bool) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.OverrideIptables = &overrideIptables
	return builder
}

func (builder *hostBuilder) Permissions(permissions []Permission) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Permissions = permissions
	return builder
}

func (builder *hostBuilder) Port(port int64) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Port = &port
	return builder
}

func (builder *hostBuilder) PowerManagement(powerManagement *PowerManagement) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.PowerManagement = powerManagement
	return builder
}

func (builder *hostBuilder) Protocol(protocol HostProtocol) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Protocol = protocol
	return builder
}

func (builder *hostBuilder) RootPassword(rootPassword string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.RootPassword = &rootPassword
	return builder
}

func (builder *hostBuilder) SeLinux(seLinux *SeLinux) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.SeLinux = seLinux
	return builder
}

func (builder *hostBuilder) Spm(spm *Spm) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Spm = spm
	return builder
}

func (builder *hostBuilder) Ssh(ssh *Ssh) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Ssh = ssh
	return builder
}

func (builder *hostBuilder) Statistics(statistics []Statistic) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Statistics = statistics
	return builder
}

func (builder *hostBuilder) Status(status HostStatus) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Status = status
	return builder
}

func (builder *hostBuilder) StatusDetail(statusDetail string) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.StatusDetail = &statusDetail
	return builder
}

func (builder *hostBuilder) StorageConnectionExtensions(storageConnectionExtensions []StorageConnectionExtension) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.StorageConnectionExtensions = storageConnectionExtensions
	return builder
}

func (builder *hostBuilder) Storages(storages []HostStorage) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Storages = storages
	return builder
}

func (builder *hostBuilder) Summary(summary *VmSummary) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Summary = summary
	return builder
}

func (builder *hostBuilder) Tags(tags []Tag) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Tags = tags
	return builder
}

func (builder *hostBuilder) TransparentHugePages(transparentHugePages *TransparentHugePages) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.TransparentHugePages = transparentHugePages
	return builder
}

func (builder *hostBuilder) Type(type_ HostType) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Type = type_
	return builder
}

func (builder *hostBuilder) UnmanagedNetworks(unmanagedNetworks []UnmanagedNetwork) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.UnmanagedNetworks = unmanagedNetworks
	return builder
}

func (builder *hostBuilder) UpdateAvailable(updateAvailable bool) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.UpdateAvailable = &updateAvailable
	return builder
}

func (builder *hostBuilder) Version(version *Version) *hostBuilder {
	if builder.err != nil {
		return builder
	}

	builder.host.Version = version
	return builder
}

func (builder *hostBuilder) Build() (*Host, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.host, nil
}

type hostDeviceBuilder struct {
	hostDevice *HostDevice
	err        error
}

func NewHostDeviceBuilder() *hostDeviceBuilder {
	return &hostDeviceBuilder{hostDevice: &HostDevice{}, err: nil}
}

func (builder *hostDeviceBuilder) Capability(capability string) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Capability = &capability
	return builder
}

func (builder *hostDeviceBuilder) Comment(comment string) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Comment = &comment
	return builder
}

func (builder *hostDeviceBuilder) Description(description string) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Description = &description
	return builder
}

func (builder *hostDeviceBuilder) Host(host *Host) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Host = host
	return builder
}

func (builder *hostDeviceBuilder) Id(id string) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Id = &id
	return builder
}

func (builder *hostDeviceBuilder) IommuGroup(iommuGroup int64) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.IommuGroup = &iommuGroup
	return builder
}

func (builder *hostDeviceBuilder) Name(name string) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Name = &name
	return builder
}

func (builder *hostDeviceBuilder) ParentDevice(parentDevice *HostDevice) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.ParentDevice = parentDevice
	return builder
}

func (builder *hostDeviceBuilder) PhysicalFunction(physicalFunction *HostDevice) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.PhysicalFunction = physicalFunction
	return builder
}

func (builder *hostDeviceBuilder) Placeholder(placeholder bool) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Placeholder = &placeholder
	return builder
}

func (builder *hostDeviceBuilder) Product(product *Product) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Product = product
	return builder
}

func (builder *hostDeviceBuilder) Vendor(vendor *Vendor) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Vendor = vendor
	return builder
}

func (builder *hostDeviceBuilder) VirtualFunctions(virtualFunctions int64) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.VirtualFunctions = &virtualFunctions
	return builder
}

func (builder *hostDeviceBuilder) Vm(vm *Vm) *hostDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevice.Vm = vm
	return builder
}

func (builder *hostDeviceBuilder) Build() (*HostDevice, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hostDevice, nil
}

type hostDevicePassthroughBuilder struct {
	hostDevicePassthrough *HostDevicePassthrough
	err                   error
}

func NewHostDevicePassthroughBuilder() *hostDevicePassthroughBuilder {
	return &hostDevicePassthroughBuilder{hostDevicePassthrough: &HostDevicePassthrough{}, err: nil}
}

func (builder *hostDevicePassthroughBuilder) Enabled(enabled bool) *hostDevicePassthroughBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostDevicePassthrough.Enabled = &enabled
	return builder
}

func (builder *hostDevicePassthroughBuilder) Build() (*HostDevicePassthrough, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hostDevicePassthrough, nil
}

type hostNicBuilder struct {
	hostNic *HostNic
	err     error
}

func NewHostNicBuilder() *hostNicBuilder {
	return &hostNicBuilder{hostNic: &HostNic{}, err: nil}
}

func (builder *hostNicBuilder) AdAggregatorId(adAggregatorId int64) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.AdAggregatorId = &adAggregatorId
	return builder
}

func (builder *hostNicBuilder) BaseInterface(baseInterface string) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.BaseInterface = &baseInterface
	return builder
}

func (builder *hostNicBuilder) Bonding(bonding *Bonding) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Bonding = bonding
	return builder
}

func (builder *hostNicBuilder) BootProtocol(bootProtocol BootProtocol) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.BootProtocol = bootProtocol
	return builder
}

func (builder *hostNicBuilder) Bridged(bridged bool) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Bridged = &bridged
	return builder
}

func (builder *hostNicBuilder) CheckConnectivity(checkConnectivity bool) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.CheckConnectivity = &checkConnectivity
	return builder
}

func (builder *hostNicBuilder) Comment(comment string) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Comment = &comment
	return builder
}

func (builder *hostNicBuilder) CustomConfiguration(customConfiguration bool) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.CustomConfiguration = &customConfiguration
	return builder
}

func (builder *hostNicBuilder) Description(description string) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Description = &description
	return builder
}

func (builder *hostNicBuilder) Host(host *Host) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Host = host
	return builder
}

func (builder *hostNicBuilder) Id(id string) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Id = &id
	return builder
}

func (builder *hostNicBuilder) Ip(ip *Ip) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Ip = ip
	return builder
}

func (builder *hostNicBuilder) Ipv6(ipv6 *Ip) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Ipv6 = ipv6
	return builder
}

func (builder *hostNicBuilder) Ipv6BootProtocol(ipv6BootProtocol BootProtocol) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Ipv6BootProtocol = ipv6BootProtocol
	return builder
}

func (builder *hostNicBuilder) Mac(mac *Mac) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Mac = mac
	return builder
}

func (builder *hostNicBuilder) Mtu(mtu int64) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Mtu = &mtu
	return builder
}

func (builder *hostNicBuilder) Name(name string) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Name = &name
	return builder
}

func (builder *hostNicBuilder) Network(network *Network) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Network = network
	return builder
}

func (builder *hostNicBuilder) NetworkLabels(networkLabels []NetworkLabel) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.NetworkLabels = networkLabels
	return builder
}

func (builder *hostNicBuilder) OverrideConfiguration(overrideConfiguration bool) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.OverrideConfiguration = &overrideConfiguration
	return builder
}

func (builder *hostNicBuilder) PhysicalFunction(physicalFunction *HostNic) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.PhysicalFunction = physicalFunction
	return builder
}

func (builder *hostNicBuilder) Properties(properties []Property) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Properties = properties
	return builder
}

func (builder *hostNicBuilder) Qos(qos *Qos) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Qos = qos
	return builder
}

func (builder *hostNicBuilder) Speed(speed int64) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Speed = &speed
	return builder
}

func (builder *hostNicBuilder) Statistics(statistics []Statistic) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Statistics = statistics
	return builder
}

func (builder *hostNicBuilder) Status(status NicStatus) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Status = status
	return builder
}

func (builder *hostNicBuilder) VirtualFunctionsConfiguration(virtualFunctionsConfiguration *HostNicVirtualFunctionsConfiguration) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.VirtualFunctionsConfiguration = virtualFunctionsConfiguration
	return builder
}

func (builder *hostNicBuilder) Vlan(vlan *Vlan) *hostNicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNic.Vlan = vlan
	return builder
}

func (builder *hostNicBuilder) Build() (*HostNic, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hostNic, nil
}

type hostNicVirtualFunctionsConfigurationBuilder struct {
	hostNicVirtualFunctionsConfiguration *HostNicVirtualFunctionsConfiguration
	err                                  error
}

func NewHostNicVirtualFunctionsConfigurationBuilder() *hostNicVirtualFunctionsConfigurationBuilder {
	return &hostNicVirtualFunctionsConfigurationBuilder{hostNicVirtualFunctionsConfiguration: &HostNicVirtualFunctionsConfiguration{}, err: nil}
}

func (builder *hostNicVirtualFunctionsConfigurationBuilder) AllNetworksAllowed(allNetworksAllowed bool) *hostNicVirtualFunctionsConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNicVirtualFunctionsConfiguration.AllNetworksAllowed = &allNetworksAllowed
	return builder
}

func (builder *hostNicVirtualFunctionsConfigurationBuilder) MaxNumberOfVirtualFunctions(maxNumberOfVirtualFunctions int64) *hostNicVirtualFunctionsConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNicVirtualFunctionsConfiguration.MaxNumberOfVirtualFunctions = &maxNumberOfVirtualFunctions
	return builder
}

func (builder *hostNicVirtualFunctionsConfigurationBuilder) NumberOfVirtualFunctions(numberOfVirtualFunctions int64) *hostNicVirtualFunctionsConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostNicVirtualFunctionsConfiguration.NumberOfVirtualFunctions = &numberOfVirtualFunctions
	return builder
}

func (builder *hostNicVirtualFunctionsConfigurationBuilder) Build() (*HostNicVirtualFunctionsConfiguration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hostNicVirtualFunctionsConfiguration, nil
}

type hostStorageBuilder struct {
	hostStorage *HostStorage
	err         error
}

func NewHostStorageBuilder() *hostStorageBuilder {
	return &hostStorageBuilder{hostStorage: &HostStorage{}, err: nil}
}

func (builder *hostStorageBuilder) Address(address string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Address = &address
	return builder
}

func (builder *hostStorageBuilder) Comment(comment string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Comment = &comment
	return builder
}

func (builder *hostStorageBuilder) Description(description string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Description = &description
	return builder
}

func (builder *hostStorageBuilder) Host(host *Host) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Host = host
	return builder
}

func (builder *hostStorageBuilder) Id(id string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Id = &id
	return builder
}

func (builder *hostStorageBuilder) LogicalUnits(logicalUnits []LogicalUnit) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.LogicalUnits = logicalUnits
	return builder
}

func (builder *hostStorageBuilder) MountOptions(mountOptions string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.MountOptions = &mountOptions
	return builder
}

func (builder *hostStorageBuilder) Name(name string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Name = &name
	return builder
}

func (builder *hostStorageBuilder) NfsRetrans(nfsRetrans int64) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.NfsRetrans = &nfsRetrans
	return builder
}

func (builder *hostStorageBuilder) NfsTimeo(nfsTimeo int64) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.NfsTimeo = &nfsTimeo
	return builder
}

func (builder *hostStorageBuilder) NfsVersion(nfsVersion NfsVersion) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.NfsVersion = nfsVersion
	return builder
}

func (builder *hostStorageBuilder) OverrideLuns(overrideLuns bool) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.OverrideLuns = &overrideLuns
	return builder
}

func (builder *hostStorageBuilder) Password(password string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Password = &password
	return builder
}

func (builder *hostStorageBuilder) Path(path string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Path = &path
	return builder
}

func (builder *hostStorageBuilder) Port(port int64) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Port = &port
	return builder
}

func (builder *hostStorageBuilder) Portal(portal string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Portal = &portal
	return builder
}

func (builder *hostStorageBuilder) Target(target string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Target = &target
	return builder
}

func (builder *hostStorageBuilder) Type(type_ StorageType) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Type = type_
	return builder
}

func (builder *hostStorageBuilder) Username(username string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.Username = &username
	return builder
}

func (builder *hostStorageBuilder) VfsType(vfsType string) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.VfsType = &vfsType
	return builder
}

func (builder *hostStorageBuilder) VolumeGroup(volumeGroup *VolumeGroup) *hostStorageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostStorage.VolumeGroup = volumeGroup
	return builder
}

func (builder *hostStorageBuilder) Build() (*HostStorage, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hostStorage, nil
}

type hostedEngineBuilder struct {
	hostedEngine *HostedEngine
	err          error
}

func NewHostedEngineBuilder() *hostedEngineBuilder {
	return &hostedEngineBuilder{hostedEngine: &HostedEngine{}, err: nil}
}

func (builder *hostedEngineBuilder) Active(active bool) *hostedEngineBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostedEngine.Active = &active
	return builder
}

func (builder *hostedEngineBuilder) Configured(configured bool) *hostedEngineBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostedEngine.Configured = &configured
	return builder
}

func (builder *hostedEngineBuilder) GlobalMaintenance(globalMaintenance bool) *hostedEngineBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostedEngine.GlobalMaintenance = &globalMaintenance
	return builder
}

func (builder *hostedEngineBuilder) LocalMaintenance(localMaintenance bool) *hostedEngineBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostedEngine.LocalMaintenance = &localMaintenance
	return builder
}

func (builder *hostedEngineBuilder) Score(score int64) *hostedEngineBuilder {
	if builder.err != nil {
		return builder
	}

	builder.hostedEngine.Score = &score
	return builder
}

func (builder *hostedEngineBuilder) Build() (*HostedEngine, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.hostedEngine, nil
}

type iconBuilder struct {
	icon *Icon
	err  error
}

func NewIconBuilder() *iconBuilder {
	return &iconBuilder{icon: &Icon{}, err: nil}
}

func (builder *iconBuilder) Comment(comment string) *iconBuilder {
	if builder.err != nil {
		return builder
	}

	builder.icon.Comment = &comment
	return builder
}

func (builder *iconBuilder) Data(data string) *iconBuilder {
	if builder.err != nil {
		return builder
	}

	builder.icon.Data = &data
	return builder
}

func (builder *iconBuilder) Description(description string) *iconBuilder {
	if builder.err != nil {
		return builder
	}

	builder.icon.Description = &description
	return builder
}

func (builder *iconBuilder) Id(id string) *iconBuilder {
	if builder.err != nil {
		return builder
	}

	builder.icon.Id = &id
	return builder
}

func (builder *iconBuilder) MediaType(mediaType string) *iconBuilder {
	if builder.err != nil {
		return builder
	}

	builder.icon.MediaType = &mediaType
	return builder
}

func (builder *iconBuilder) Name(name string) *iconBuilder {
	if builder.err != nil {
		return builder
	}

	builder.icon.Name = &name
	return builder
}

func (builder *iconBuilder) Build() (*Icon, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.icon, nil
}

type identifiedBuilder struct {
	identified *Identified
	err        error
}

func NewIdentifiedBuilder() *identifiedBuilder {
	return &identifiedBuilder{identified: &Identified{}, err: nil}
}

func (builder *identifiedBuilder) Comment(comment string) *identifiedBuilder {
	if builder.err != nil {
		return builder
	}

	builder.identified.Comment = &comment
	return builder
}

func (builder *identifiedBuilder) Description(description string) *identifiedBuilder {
	if builder.err != nil {
		return builder
	}

	builder.identified.Description = &description
	return builder
}

func (builder *identifiedBuilder) Id(id string) *identifiedBuilder {
	if builder.err != nil {
		return builder
	}

	builder.identified.Id = &id
	return builder
}

func (builder *identifiedBuilder) Name(name string) *identifiedBuilder {
	if builder.err != nil {
		return builder
	}

	builder.identified.Name = &name
	return builder
}

func (builder *identifiedBuilder) Build() (*Identified, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.identified, nil
}

type imageBuilder struct {
	image *Image
	err   error
}

func NewImageBuilder() *imageBuilder {
	return &imageBuilder{image: &Image{}, err: nil}
}

func (builder *imageBuilder) Comment(comment string) *imageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.image.Comment = &comment
	return builder
}

func (builder *imageBuilder) Description(description string) *imageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.image.Description = &description
	return builder
}

func (builder *imageBuilder) Id(id string) *imageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.image.Id = &id
	return builder
}

func (builder *imageBuilder) Name(name string) *imageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.image.Name = &name
	return builder
}

func (builder *imageBuilder) StorageDomain(storageDomain *StorageDomain) *imageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.image.StorageDomain = storageDomain
	return builder
}

func (builder *imageBuilder) Build() (*Image, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.image, nil
}

type imageTransferBuilder struct {
	imageTransfer *ImageTransfer
	err           error
}

func NewImageTransferBuilder() *imageTransferBuilder {
	return &imageTransferBuilder{imageTransfer: &ImageTransfer{}, err: nil}
}

func (builder *imageTransferBuilder) Comment(comment string) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Comment = &comment
	return builder
}

func (builder *imageTransferBuilder) Description(description string) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Description = &description
	return builder
}

func (builder *imageTransferBuilder) Direction(direction ImageTransferDirection) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Direction = direction
	return builder
}

func (builder *imageTransferBuilder) Host(host *Host) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Host = host
	return builder
}

func (builder *imageTransferBuilder) Id(id string) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Id = &id
	return builder
}

func (builder *imageTransferBuilder) Image(image *Image) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Image = image
	return builder
}

func (builder *imageTransferBuilder) Name(name string) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Name = &name
	return builder
}

func (builder *imageTransferBuilder) Phase(phase ImageTransferPhase) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.Phase = phase
	return builder
}

func (builder *imageTransferBuilder) ProxyUrl(proxyUrl string) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.ProxyUrl = &proxyUrl
	return builder
}

func (builder *imageTransferBuilder) SignedTicket(signedTicket string) *imageTransferBuilder {
	if builder.err != nil {
		return builder
	}

	builder.imageTransfer.SignedTicket = &signedTicket
	return builder
}

func (builder *imageTransferBuilder) Build() (*ImageTransfer, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.imageTransfer, nil
}

type initializationBuilder struct {
	initialization *Initialization
	err            error
}

func NewInitializationBuilder() *initializationBuilder {
	return &initializationBuilder{initialization: &Initialization{}, err: nil}
}

func (builder *initializationBuilder) ActiveDirectoryOu(activeDirectoryOu string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.ActiveDirectoryOu = &activeDirectoryOu
	return builder
}

func (builder *initializationBuilder) AuthorizedSshKeys(authorizedSshKeys string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.AuthorizedSshKeys = &authorizedSshKeys
	return builder
}

func (builder *initializationBuilder) CloudInit(cloudInit *CloudInit) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.CloudInit = cloudInit
	return builder
}

func (builder *initializationBuilder) Configuration(configuration *Configuration) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.Configuration = configuration
	return builder
}

func (builder *initializationBuilder) CustomScript(customScript string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.CustomScript = &customScript
	return builder
}

func (builder *initializationBuilder) DnsSearch(dnsSearch string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.DnsSearch = &dnsSearch
	return builder
}

func (builder *initializationBuilder) DnsServers(dnsServers string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.DnsServers = &dnsServers
	return builder
}

func (builder *initializationBuilder) Domain(domain string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.Domain = &domain
	return builder
}

func (builder *initializationBuilder) HostName(hostName string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.HostName = &hostName
	return builder
}

func (builder *initializationBuilder) InputLocale(inputLocale string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.InputLocale = &inputLocale
	return builder
}

func (builder *initializationBuilder) NicConfigurations(nicConfigurations []NicConfiguration) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.NicConfigurations = nicConfigurations
	return builder
}

func (builder *initializationBuilder) OrgName(orgName string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.OrgName = &orgName
	return builder
}

func (builder *initializationBuilder) RegenerateIds(regenerateIds bool) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.RegenerateIds = &regenerateIds
	return builder
}

func (builder *initializationBuilder) RegenerateSshKeys(regenerateSshKeys bool) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.RegenerateSshKeys = &regenerateSshKeys
	return builder
}

func (builder *initializationBuilder) RootPassword(rootPassword string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.RootPassword = &rootPassword
	return builder
}

func (builder *initializationBuilder) SystemLocale(systemLocale string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.SystemLocale = &systemLocale
	return builder
}

func (builder *initializationBuilder) Timezone(timezone string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.Timezone = &timezone
	return builder
}

func (builder *initializationBuilder) UiLanguage(uiLanguage string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.UiLanguage = &uiLanguage
	return builder
}

func (builder *initializationBuilder) UserLocale(userLocale string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.UserLocale = &userLocale
	return builder
}

func (builder *initializationBuilder) UserName(userName string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.UserName = &userName
	return builder
}

func (builder *initializationBuilder) WindowsLicenseKey(windowsLicenseKey string) *initializationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.initialization.WindowsLicenseKey = &windowsLicenseKey
	return builder
}

func (builder *initializationBuilder) Build() (*Initialization, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.initialization, nil
}

type instanceTypeBuilder struct {
	instanceType *InstanceType
	err          error
}

func NewInstanceTypeBuilder() *instanceTypeBuilder {
	return &instanceTypeBuilder{instanceType: &InstanceType{}, err: nil}
}

func (builder *instanceTypeBuilder) Bios(bios *Bios) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Bios = bios
	return builder
}

func (builder *instanceTypeBuilder) Cdroms(cdroms []Cdrom) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Cdroms = cdroms
	return builder
}

func (builder *instanceTypeBuilder) Cluster(cluster *Cluster) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Cluster = cluster
	return builder
}

func (builder *instanceTypeBuilder) Comment(comment string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Comment = &comment
	return builder
}

func (builder *instanceTypeBuilder) Console(console *Console) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Console = console
	return builder
}

func (builder *instanceTypeBuilder) Cpu(cpu *Cpu) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Cpu = cpu
	return builder
}

func (builder *instanceTypeBuilder) CpuProfile(cpuProfile *CpuProfile) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CpuProfile = cpuProfile
	return builder
}

func (builder *instanceTypeBuilder) CpuShares(cpuShares int64) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CpuShares = &cpuShares
	return builder
}

func (builder *instanceTypeBuilder) CreationTime(creationTime time.Time) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CreationTime = creationTime
	return builder
}

func (builder *instanceTypeBuilder) CustomCompatibilityVersion(customCompatibilityVersion *Version) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CustomCompatibilityVersion = customCompatibilityVersion
	return builder
}

func (builder *instanceTypeBuilder) CustomCpuModel(customCpuModel string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CustomCpuModel = &customCpuModel
	return builder
}

func (builder *instanceTypeBuilder) CustomEmulatedMachine(customEmulatedMachine string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CustomEmulatedMachine = &customEmulatedMachine
	return builder
}

func (builder *instanceTypeBuilder) CustomProperties(customProperties []CustomProperty) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.CustomProperties = customProperties
	return builder
}

func (builder *instanceTypeBuilder) DeleteProtected(deleteProtected bool) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.DeleteProtected = &deleteProtected
	return builder
}

func (builder *instanceTypeBuilder) Description(description string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Description = &description
	return builder
}

func (builder *instanceTypeBuilder) DiskAttachments(diskAttachments []DiskAttachment) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.DiskAttachments = diskAttachments
	return builder
}

func (builder *instanceTypeBuilder) Display(display *Display) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Display = display
	return builder
}

func (builder *instanceTypeBuilder) Domain(domain *Domain) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Domain = domain
	return builder
}

func (builder *instanceTypeBuilder) GraphicsConsoles(graphicsConsoles []GraphicsConsole) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.GraphicsConsoles = graphicsConsoles
	return builder
}

func (builder *instanceTypeBuilder) HighAvailability(highAvailability *HighAvailability) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.HighAvailability = highAvailability
	return builder
}

func (builder *instanceTypeBuilder) Id(id string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Id = &id
	return builder
}

func (builder *instanceTypeBuilder) Initialization(initialization *Initialization) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Initialization = initialization
	return builder
}

func (builder *instanceTypeBuilder) Io(io *Io) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Io = io
	return builder
}

func (builder *instanceTypeBuilder) LargeIcon(largeIcon *Icon) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.LargeIcon = largeIcon
	return builder
}

func (builder *instanceTypeBuilder) Lease(lease *StorageDomainLease) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Lease = lease
	return builder
}

func (builder *instanceTypeBuilder) Memory(memory int64) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Memory = &memory
	return builder
}

func (builder *instanceTypeBuilder) MemoryPolicy(memoryPolicy *MemoryPolicy) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.MemoryPolicy = memoryPolicy
	return builder
}

func (builder *instanceTypeBuilder) Migration(migration *MigrationOptions) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Migration = migration
	return builder
}

func (builder *instanceTypeBuilder) MigrationDowntime(migrationDowntime int64) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.MigrationDowntime = &migrationDowntime
	return builder
}

func (builder *instanceTypeBuilder) Name(name string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Name = &name
	return builder
}

func (builder *instanceTypeBuilder) Nics(nics []Nic) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Nics = nics
	return builder
}

func (builder *instanceTypeBuilder) Origin(origin string) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Origin = &origin
	return builder
}

func (builder *instanceTypeBuilder) Os(os *OperatingSystem) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Os = os
	return builder
}

func (builder *instanceTypeBuilder) Permissions(permissions []Permission) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Permissions = permissions
	return builder
}

func (builder *instanceTypeBuilder) Quota(quota *Quota) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Quota = quota
	return builder
}

func (builder *instanceTypeBuilder) RngDevice(rngDevice *RngDevice) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.RngDevice = rngDevice
	return builder
}

func (builder *instanceTypeBuilder) SerialNumber(serialNumber *SerialNumber) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.SerialNumber = serialNumber
	return builder
}

func (builder *instanceTypeBuilder) SmallIcon(smallIcon *Icon) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.SmallIcon = smallIcon
	return builder
}

func (builder *instanceTypeBuilder) SoundcardEnabled(soundcardEnabled bool) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.SoundcardEnabled = &soundcardEnabled
	return builder
}

func (builder *instanceTypeBuilder) Sso(sso *Sso) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Sso = sso
	return builder
}

func (builder *instanceTypeBuilder) StartPaused(startPaused bool) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.StartPaused = &startPaused
	return builder
}

func (builder *instanceTypeBuilder) Stateless(stateless bool) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Stateless = &stateless
	return builder
}

func (builder *instanceTypeBuilder) Status(status TemplateStatus) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Status = status
	return builder
}

func (builder *instanceTypeBuilder) StorageDomain(storageDomain *StorageDomain) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.StorageDomain = storageDomain
	return builder
}

func (builder *instanceTypeBuilder) Tags(tags []Tag) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Tags = tags
	return builder
}

func (builder *instanceTypeBuilder) TimeZone(timeZone *TimeZone) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.TimeZone = timeZone
	return builder
}

func (builder *instanceTypeBuilder) TunnelMigration(tunnelMigration bool) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.TunnelMigration = &tunnelMigration
	return builder
}

func (builder *instanceTypeBuilder) Type(type_ VmType) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Type = type_
	return builder
}

func (builder *instanceTypeBuilder) Usb(usb *Usb) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Usb = usb
	return builder
}

func (builder *instanceTypeBuilder) Version(version *TemplateVersion) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Version = version
	return builder
}

func (builder *instanceTypeBuilder) VirtioScsi(virtioScsi *VirtioScsi) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.VirtioScsi = virtioScsi
	return builder
}

func (builder *instanceTypeBuilder) Vm(vm *Vm) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Vm = vm
	return builder
}

func (builder *instanceTypeBuilder) Watchdogs(watchdogs []Watchdog) *instanceTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.instanceType.Watchdogs = watchdogs
	return builder
}

func (builder *instanceTypeBuilder) Build() (*InstanceType, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.instanceType, nil
}

type ioBuilder struct {
	io  *Io
	err error
}

func NewIoBuilder() *ioBuilder {
	return &ioBuilder{io: &Io{}, err: nil}
}

func (builder *ioBuilder) Threads(threads int64) *ioBuilder {
	if builder.err != nil {
		return builder
	}

	builder.io.Threads = &threads
	return builder
}

func (builder *ioBuilder) Build() (*Io, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.io, nil
}

type ipBuilder struct {
	ip  *Ip
	err error
}

func NewIpBuilder() *ipBuilder {
	return &ipBuilder{ip: &Ip{}, err: nil}
}

func (builder *ipBuilder) Address(address string) *ipBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ip.Address = &address
	return builder
}

func (builder *ipBuilder) Gateway(gateway string) *ipBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ip.Gateway = &gateway
	return builder
}

func (builder *ipBuilder) Netmask(netmask string) *ipBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ip.Netmask = &netmask
	return builder
}

func (builder *ipBuilder) Version(version IpVersion) *ipBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ip.Version = version
	return builder
}

func (builder *ipBuilder) Build() (*Ip, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.ip, nil
}

type ipAddressAssignmentBuilder struct {
	ipAddressAssignment *IpAddressAssignment
	err                 error
}

func NewIpAddressAssignmentBuilder() *ipAddressAssignmentBuilder {
	return &ipAddressAssignmentBuilder{ipAddressAssignment: &IpAddressAssignment{}, err: nil}
}

func (builder *ipAddressAssignmentBuilder) AssignmentMethod(assignmentMethod BootProtocol) *ipAddressAssignmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ipAddressAssignment.AssignmentMethod = assignmentMethod
	return builder
}

func (builder *ipAddressAssignmentBuilder) Ip(ip *Ip) *ipAddressAssignmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ipAddressAssignment.Ip = ip
	return builder
}

func (builder *ipAddressAssignmentBuilder) Build() (*IpAddressAssignment, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.ipAddressAssignment, nil
}

type iscsiBondBuilder struct {
	iscsiBond *IscsiBond
	err       error
}

func NewIscsiBondBuilder() *iscsiBondBuilder {
	return &iscsiBondBuilder{iscsiBond: &IscsiBond{}, err: nil}
}

func (builder *iscsiBondBuilder) Comment(comment string) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.Comment = &comment
	return builder
}

func (builder *iscsiBondBuilder) DataCenter(dataCenter *DataCenter) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.DataCenter = dataCenter
	return builder
}

func (builder *iscsiBondBuilder) Description(description string) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.Description = &description
	return builder
}

func (builder *iscsiBondBuilder) Id(id string) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.Id = &id
	return builder
}

func (builder *iscsiBondBuilder) Name(name string) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.Name = &name
	return builder
}

func (builder *iscsiBondBuilder) Networks(networks []Network) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.Networks = networks
	return builder
}

func (builder *iscsiBondBuilder) StorageConnections(storageConnections []StorageConnection) *iscsiBondBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiBond.StorageConnections = storageConnections
	return builder
}

func (builder *iscsiBondBuilder) Build() (*IscsiBond, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.iscsiBond, nil
}

type iscsiDetailsBuilder struct {
	iscsiDetails *IscsiDetails
	err          error
}

func NewIscsiDetailsBuilder() *iscsiDetailsBuilder {
	return &iscsiDetailsBuilder{iscsiDetails: &IscsiDetails{}, err: nil}
}

func (builder *iscsiDetailsBuilder) Address(address string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Address = &address
	return builder
}

func (builder *iscsiDetailsBuilder) DiskId(diskId string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.DiskId = &diskId
	return builder
}

func (builder *iscsiDetailsBuilder) Initiator(initiator string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Initiator = &initiator
	return builder
}

func (builder *iscsiDetailsBuilder) LunMapping(lunMapping int64) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.LunMapping = &lunMapping
	return builder
}

func (builder *iscsiDetailsBuilder) Password(password string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Password = &password
	return builder
}

func (builder *iscsiDetailsBuilder) Paths(paths int64) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Paths = &paths
	return builder
}

func (builder *iscsiDetailsBuilder) Port(port int64) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Port = &port
	return builder
}

func (builder *iscsiDetailsBuilder) Portal(portal string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Portal = &portal
	return builder
}

func (builder *iscsiDetailsBuilder) ProductId(productId string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.ProductId = &productId
	return builder
}

func (builder *iscsiDetailsBuilder) Serial(serial string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Serial = &serial
	return builder
}

func (builder *iscsiDetailsBuilder) Size(size int64) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Size = &size
	return builder
}

func (builder *iscsiDetailsBuilder) Status(status string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Status = &status
	return builder
}

func (builder *iscsiDetailsBuilder) StorageDomainId(storageDomainId string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.StorageDomainId = &storageDomainId
	return builder
}

func (builder *iscsiDetailsBuilder) Target(target string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Target = &target
	return builder
}

func (builder *iscsiDetailsBuilder) Username(username string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.Username = &username
	return builder
}

func (builder *iscsiDetailsBuilder) VendorId(vendorId string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.VendorId = &vendorId
	return builder
}

func (builder *iscsiDetailsBuilder) VolumeGroupId(volumeGroupId string) *iscsiDetailsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.iscsiDetails.VolumeGroupId = &volumeGroupId
	return builder
}

func (builder *iscsiDetailsBuilder) Build() (*IscsiDetails, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.iscsiDetails, nil
}

type jobBuilder struct {
	job *Job
	err error
}

func NewJobBuilder() *jobBuilder {
	return &jobBuilder{job: &Job{}, err: nil}
}

func (builder *jobBuilder) AutoCleared(autoCleared bool) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.AutoCleared = &autoCleared
	return builder
}

func (builder *jobBuilder) Comment(comment string) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Comment = &comment
	return builder
}

func (builder *jobBuilder) Description(description string) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Description = &description
	return builder
}

func (builder *jobBuilder) EndTime(endTime time.Time) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.EndTime = endTime
	return builder
}

func (builder *jobBuilder) External(external bool) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.External = &external
	return builder
}

func (builder *jobBuilder) Id(id string) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Id = &id
	return builder
}

func (builder *jobBuilder) LastUpdated(lastUpdated time.Time) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.LastUpdated = lastUpdated
	return builder
}

func (builder *jobBuilder) Name(name string) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Name = &name
	return builder
}

func (builder *jobBuilder) Owner(owner *User) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Owner = owner
	return builder
}

func (builder *jobBuilder) StartTime(startTime time.Time) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.StartTime = startTime
	return builder
}

func (builder *jobBuilder) Status(status JobStatus) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Status = status
	return builder
}

func (builder *jobBuilder) Steps(steps []Step) *jobBuilder {
	if builder.err != nil {
		return builder
	}

	builder.job.Steps = steps
	return builder
}

func (builder *jobBuilder) Build() (*Job, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.job, nil
}

type katelloErratumBuilder struct {
	katelloErratum *KatelloErratum
	err            error
}

func NewKatelloErratumBuilder() *katelloErratumBuilder {
	return &katelloErratumBuilder{katelloErratum: &KatelloErratum{}, err: nil}
}

func (builder *katelloErratumBuilder) Comment(comment string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Comment = &comment
	return builder
}

func (builder *katelloErratumBuilder) Description(description string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Description = &description
	return builder
}

func (builder *katelloErratumBuilder) Host(host *Host) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Host = host
	return builder
}

func (builder *katelloErratumBuilder) Id(id string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Id = &id
	return builder
}

func (builder *katelloErratumBuilder) Issued(issued time.Time) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Issued = issued
	return builder
}

func (builder *katelloErratumBuilder) Name(name string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Name = &name
	return builder
}

func (builder *katelloErratumBuilder) Packages(packages []Package) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Packages = packages
	return builder
}

func (builder *katelloErratumBuilder) Severity(severity string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Severity = &severity
	return builder
}

func (builder *katelloErratumBuilder) Solution(solution string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Solution = &solution
	return builder
}

func (builder *katelloErratumBuilder) Summary(summary string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Summary = &summary
	return builder
}

func (builder *katelloErratumBuilder) Title(title string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Title = &title
	return builder
}

func (builder *katelloErratumBuilder) Type(type_ string) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Type = &type_
	return builder
}

func (builder *katelloErratumBuilder) Vm(vm *Vm) *katelloErratumBuilder {
	if builder.err != nil {
		return builder
	}

	builder.katelloErratum.Vm = vm
	return builder
}

func (builder *katelloErratumBuilder) Build() (*KatelloErratum, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.katelloErratum, nil
}

type kernelBuilder struct {
	kernel *Kernel
	err    error
}

func NewKernelBuilder() *kernelBuilder {
	return &kernelBuilder{kernel: &Kernel{}, err: nil}
}

func (builder *kernelBuilder) Version(version *Version) *kernelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.kernel.Version = version
	return builder
}

func (builder *kernelBuilder) Build() (*Kernel, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.kernel, nil
}

type ksmBuilder struct {
	ksm *Ksm
	err error
}

func NewKsmBuilder() *ksmBuilder {
	return &ksmBuilder{ksm: &Ksm{}, err: nil}
}

func (builder *ksmBuilder) Enabled(enabled bool) *ksmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ksm.Enabled = &enabled
	return builder
}

func (builder *ksmBuilder) MergeAcrossNodes(mergeAcrossNodes bool) *ksmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ksm.MergeAcrossNodes = &mergeAcrossNodes
	return builder
}

func (builder *ksmBuilder) Build() (*Ksm, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.ksm, nil
}

type logicalUnitBuilder struct {
	logicalUnit *LogicalUnit
	err         error
}

func NewLogicalUnitBuilder() *logicalUnitBuilder {
	return &logicalUnitBuilder{logicalUnit: &LogicalUnit{}, err: nil}
}

func (builder *logicalUnitBuilder) Address(address string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Address = &address
	return builder
}

func (builder *logicalUnitBuilder) DiscardMaxSize(discardMaxSize int64) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.DiscardMaxSize = &discardMaxSize
	return builder
}

func (builder *logicalUnitBuilder) DiscardZeroesData(discardZeroesData bool) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.DiscardZeroesData = &discardZeroesData
	return builder
}

func (builder *logicalUnitBuilder) DiskId(diskId string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.DiskId = &diskId
	return builder
}

func (builder *logicalUnitBuilder) Id(id string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Id = &id
	return builder
}

func (builder *logicalUnitBuilder) LunMapping(lunMapping int64) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.LunMapping = &lunMapping
	return builder
}

func (builder *logicalUnitBuilder) Password(password string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Password = &password
	return builder
}

func (builder *logicalUnitBuilder) Paths(paths int64) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Paths = &paths
	return builder
}

func (builder *logicalUnitBuilder) Port(port int64) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Port = &port
	return builder
}

func (builder *logicalUnitBuilder) Portal(portal string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Portal = &portal
	return builder
}

func (builder *logicalUnitBuilder) ProductId(productId string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.ProductId = &productId
	return builder
}

func (builder *logicalUnitBuilder) Serial(serial string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Serial = &serial
	return builder
}

func (builder *logicalUnitBuilder) Size(size int64) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Size = &size
	return builder
}

func (builder *logicalUnitBuilder) Status(status LunStatus) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Status = status
	return builder
}

func (builder *logicalUnitBuilder) StorageDomainId(storageDomainId string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.StorageDomainId = &storageDomainId
	return builder
}

func (builder *logicalUnitBuilder) Target(target string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Target = &target
	return builder
}

func (builder *logicalUnitBuilder) Username(username string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.Username = &username
	return builder
}

func (builder *logicalUnitBuilder) VendorId(vendorId string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.VendorId = &vendorId
	return builder
}

func (builder *logicalUnitBuilder) VolumeGroupId(volumeGroupId string) *logicalUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.logicalUnit.VolumeGroupId = &volumeGroupId
	return builder
}

func (builder *logicalUnitBuilder) Build() (*LogicalUnit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.logicalUnit, nil
}

type macBuilder struct {
	mac *Mac
	err error
}

func NewMacBuilder() *macBuilder {
	return &macBuilder{mac: &Mac{}, err: nil}
}

func (builder *macBuilder) Address(address string) *macBuilder {
	if builder.err != nil {
		return builder
	}

	builder.mac.Address = &address
	return builder
}

func (builder *macBuilder) Build() (*Mac, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.mac, nil
}

type macPoolBuilder struct {
	macPool *MacPool
	err     error
}

func NewMacPoolBuilder() *macPoolBuilder {
	return &macPoolBuilder{macPool: &MacPool{}, err: nil}
}

func (builder *macPoolBuilder) AllowDuplicates(allowDuplicates bool) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.AllowDuplicates = &allowDuplicates
	return builder
}

func (builder *macPoolBuilder) Comment(comment string) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.Comment = &comment
	return builder
}

func (builder *macPoolBuilder) DefaultPool(defaultPool bool) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.DefaultPool = &defaultPool
	return builder
}

func (builder *macPoolBuilder) Description(description string) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.Description = &description
	return builder
}

func (builder *macPoolBuilder) Id(id string) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.Id = &id
	return builder
}

func (builder *macPoolBuilder) Name(name string) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.Name = &name
	return builder
}

func (builder *macPoolBuilder) Ranges(ranges []Range) *macPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.macPool.Ranges = ranges
	return builder
}

func (builder *macPoolBuilder) Build() (*MacPool, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.macPool, nil
}

type memoryOverCommitBuilder struct {
	memoryOverCommit *MemoryOverCommit
	err              error
}

func NewMemoryOverCommitBuilder() *memoryOverCommitBuilder {
	return &memoryOverCommitBuilder{memoryOverCommit: &MemoryOverCommit{}, err: nil}
}

func (builder *memoryOverCommitBuilder) Percent(percent int64) *memoryOverCommitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.memoryOverCommit.Percent = &percent
	return builder
}

func (builder *memoryOverCommitBuilder) Build() (*MemoryOverCommit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.memoryOverCommit, nil
}

type memoryPolicyBuilder struct {
	memoryPolicy *MemoryPolicy
	err          error
}

func NewMemoryPolicyBuilder() *memoryPolicyBuilder {
	return &memoryPolicyBuilder{memoryPolicy: &MemoryPolicy{}, err: nil}
}

func (builder *memoryPolicyBuilder) Ballooning(ballooning bool) *memoryPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.memoryPolicy.Ballooning = &ballooning
	return builder
}

func (builder *memoryPolicyBuilder) Guaranteed(guaranteed int64) *memoryPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.memoryPolicy.Guaranteed = &guaranteed
	return builder
}

func (builder *memoryPolicyBuilder) Max(max int64) *memoryPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.memoryPolicy.Max = &max
	return builder
}

func (builder *memoryPolicyBuilder) OverCommit(overCommit *MemoryOverCommit) *memoryPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.memoryPolicy.OverCommit = overCommit
	return builder
}

func (builder *memoryPolicyBuilder) TransparentHugePages(transparentHugePages *TransparentHugePages) *memoryPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.memoryPolicy.TransparentHugePages = transparentHugePages
	return builder
}

func (builder *memoryPolicyBuilder) Build() (*MemoryPolicy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.memoryPolicy, nil
}

type methodBuilder struct {
	method *Method
	err    error
}

func NewMethodBuilder() *methodBuilder {
	return &methodBuilder{method: &Method{}, err: nil}
}

func (builder *methodBuilder) Id(id SsoMethod) *methodBuilder {
	if builder.err != nil {
		return builder
	}

	builder.method.Id = id
	return builder
}

func (builder *methodBuilder) Build() (*Method, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.method, nil
}

type migrationBandwidthBuilder struct {
	migrationBandwidth *MigrationBandwidth
	err                error
}

func NewMigrationBandwidthBuilder() *migrationBandwidthBuilder {
	return &migrationBandwidthBuilder{migrationBandwidth: &MigrationBandwidth{}, err: nil}
}

func (builder *migrationBandwidthBuilder) AssignmentMethod(assignmentMethod MigrationBandwidthAssignmentMethod) *migrationBandwidthBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationBandwidth.AssignmentMethod = assignmentMethod
	return builder
}

func (builder *migrationBandwidthBuilder) CustomValue(customValue int64) *migrationBandwidthBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationBandwidth.CustomValue = &customValue
	return builder
}

func (builder *migrationBandwidthBuilder) Build() (*MigrationBandwidth, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.migrationBandwidth, nil
}

type migrationOptionsBuilder struct {
	migrationOptions *MigrationOptions
	err              error
}

func NewMigrationOptionsBuilder() *migrationOptionsBuilder {
	return &migrationOptionsBuilder{migrationOptions: &MigrationOptions{}, err: nil}
}

func (builder *migrationOptionsBuilder) AutoConverge(autoConverge InheritableBoolean) *migrationOptionsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationOptions.AutoConverge = autoConverge
	return builder
}

func (builder *migrationOptionsBuilder) Bandwidth(bandwidth *MigrationBandwidth) *migrationOptionsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationOptions.Bandwidth = bandwidth
	return builder
}

func (builder *migrationOptionsBuilder) Compressed(compressed InheritableBoolean) *migrationOptionsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationOptions.Compressed = compressed
	return builder
}

func (builder *migrationOptionsBuilder) Policy(policy *MigrationPolicy) *migrationOptionsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationOptions.Policy = policy
	return builder
}

func (builder *migrationOptionsBuilder) Build() (*MigrationOptions, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.migrationOptions, nil
}

type migrationPolicyBuilder struct {
	migrationPolicy *MigrationPolicy
	err             error
}

func NewMigrationPolicyBuilder() *migrationPolicyBuilder {
	return &migrationPolicyBuilder{migrationPolicy: &MigrationPolicy{}, err: nil}
}

func (builder *migrationPolicyBuilder) Comment(comment string) *migrationPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationPolicy.Comment = &comment
	return builder
}

func (builder *migrationPolicyBuilder) Description(description string) *migrationPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationPolicy.Description = &description
	return builder
}

func (builder *migrationPolicyBuilder) Id(id string) *migrationPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationPolicy.Id = &id
	return builder
}

func (builder *migrationPolicyBuilder) Name(name string) *migrationPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.migrationPolicy.Name = &name
	return builder
}

func (builder *migrationPolicyBuilder) Build() (*MigrationPolicy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.migrationPolicy, nil
}

type networkBuilder struct {
	network *Network
	err     error
}

func NewNetworkBuilder() *networkBuilder {
	return &networkBuilder{network: &Network{}, err: nil}
}

func (builder *networkBuilder) Cluster(cluster *Cluster) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Cluster = cluster
	return builder
}

func (builder *networkBuilder) Comment(comment string) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Comment = &comment
	return builder
}

func (builder *networkBuilder) DataCenter(dataCenter *DataCenter) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.DataCenter = dataCenter
	return builder
}

func (builder *networkBuilder) Description(description string) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Description = &description
	return builder
}

func (builder *networkBuilder) Display(display bool) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Display = &display
	return builder
}

func (builder *networkBuilder) DnsResolverConfiguration(dnsResolverConfiguration *DnsResolverConfiguration) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.DnsResolverConfiguration = dnsResolverConfiguration
	return builder
}

func (builder *networkBuilder) Id(id string) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Id = &id
	return builder
}

func (builder *networkBuilder) Ip(ip *Ip) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Ip = ip
	return builder
}

func (builder *networkBuilder) Mtu(mtu int64) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Mtu = &mtu
	return builder
}

func (builder *networkBuilder) Name(name string) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Name = &name
	return builder
}

func (builder *networkBuilder) NetworkLabels(networkLabels []NetworkLabel) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.NetworkLabels = networkLabels
	return builder
}

func (builder *networkBuilder) Permissions(permissions []Permission) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Permissions = permissions
	return builder
}

func (builder *networkBuilder) ProfileRequired(profileRequired bool) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.ProfileRequired = &profileRequired
	return builder
}

func (builder *networkBuilder) Qos(qos *Qos) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Qos = qos
	return builder
}

func (builder *networkBuilder) Required(required bool) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Required = &required
	return builder
}

func (builder *networkBuilder) Status(status NetworkStatus) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Status = status
	return builder
}

func (builder *networkBuilder) Stp(stp bool) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Stp = &stp
	return builder
}

func (builder *networkBuilder) Usages(usages []NetworkUsage) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Usages = usages
	return builder
}

func (builder *networkBuilder) Vlan(vlan *Vlan) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.Vlan = vlan
	return builder
}

func (builder *networkBuilder) VnicProfiles(vnicProfiles []VnicProfile) *networkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.network.VnicProfiles = vnicProfiles
	return builder
}

func (builder *networkBuilder) Build() (*Network, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.network, nil
}

type networkAttachmentBuilder struct {
	networkAttachment *NetworkAttachment
	err               error
}

func NewNetworkAttachmentBuilder() *networkAttachmentBuilder {
	return &networkAttachmentBuilder{networkAttachment: &NetworkAttachment{}, err: nil}
}

func (builder *networkAttachmentBuilder) Comment(comment string) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Comment = &comment
	return builder
}

func (builder *networkAttachmentBuilder) Description(description string) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Description = &description
	return builder
}

func (builder *networkAttachmentBuilder) DnsResolverConfiguration(dnsResolverConfiguration *DnsResolverConfiguration) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.DnsResolverConfiguration = dnsResolverConfiguration
	return builder
}

func (builder *networkAttachmentBuilder) Host(host *Host) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Host = host
	return builder
}

func (builder *networkAttachmentBuilder) HostNic(hostNic *HostNic) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.HostNic = hostNic
	return builder
}

func (builder *networkAttachmentBuilder) Id(id string) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Id = &id
	return builder
}

func (builder *networkAttachmentBuilder) InSync(inSync bool) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.InSync = &inSync
	return builder
}

func (builder *networkAttachmentBuilder) IpAddressAssignments(ipAddressAssignments []IpAddressAssignment) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.IpAddressAssignments = ipAddressAssignments
	return builder
}

func (builder *networkAttachmentBuilder) Name(name string) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Name = &name
	return builder
}

func (builder *networkAttachmentBuilder) Network(network *Network) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Network = network
	return builder
}

func (builder *networkAttachmentBuilder) Properties(properties []Property) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Properties = properties
	return builder
}

func (builder *networkAttachmentBuilder) Qos(qos *Qos) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.Qos = qos
	return builder
}

func (builder *networkAttachmentBuilder) ReportedConfigurations(reportedConfigurations []ReportedConfiguration) *networkAttachmentBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkAttachment.ReportedConfigurations = reportedConfigurations
	return builder
}

func (builder *networkAttachmentBuilder) Build() (*NetworkAttachment, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.networkAttachment, nil
}

type networkConfigurationBuilder struct {
	networkConfiguration *NetworkConfiguration
	err                  error
}

func NewNetworkConfigurationBuilder() *networkConfigurationBuilder {
	return &networkConfigurationBuilder{networkConfiguration: &NetworkConfiguration{}, err: nil}
}

func (builder *networkConfigurationBuilder) Dns(dns *Dns) *networkConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkConfiguration.Dns = dns
	return builder
}

func (builder *networkConfigurationBuilder) Nics(nics []Nic) *networkConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkConfiguration.Nics = nics
	return builder
}

func (builder *networkConfigurationBuilder) Build() (*NetworkConfiguration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.networkConfiguration, nil
}

type networkFilterBuilder struct {
	networkFilter *NetworkFilter
	err           error
}

func NewNetworkFilterBuilder() *networkFilterBuilder {
	return &networkFilterBuilder{networkFilter: &NetworkFilter{}, err: nil}
}

func (builder *networkFilterBuilder) Comment(comment string) *networkFilterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilter.Comment = &comment
	return builder
}

func (builder *networkFilterBuilder) Description(description string) *networkFilterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilter.Description = &description
	return builder
}

func (builder *networkFilterBuilder) Id(id string) *networkFilterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilter.Id = &id
	return builder
}

func (builder *networkFilterBuilder) Name(name string) *networkFilterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilter.Name = &name
	return builder
}

func (builder *networkFilterBuilder) Version(version *Version) *networkFilterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilter.Version = version
	return builder
}

func (builder *networkFilterBuilder) Build() (*NetworkFilter, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.networkFilter, nil
}

type networkFilterParameterBuilder struct {
	networkFilterParameter *NetworkFilterParameter
	err                    error
}

func NewNetworkFilterParameterBuilder() *networkFilterParameterBuilder {
	return &networkFilterParameterBuilder{networkFilterParameter: &NetworkFilterParameter{}, err: nil}
}

func (builder *networkFilterParameterBuilder) Comment(comment string) *networkFilterParameterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilterParameter.Comment = &comment
	return builder
}

func (builder *networkFilterParameterBuilder) Description(description string) *networkFilterParameterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilterParameter.Description = &description
	return builder
}

func (builder *networkFilterParameterBuilder) Id(id string) *networkFilterParameterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilterParameter.Id = &id
	return builder
}

func (builder *networkFilterParameterBuilder) Name(name string) *networkFilterParameterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilterParameter.Name = &name
	return builder
}

func (builder *networkFilterParameterBuilder) Value(value string) *networkFilterParameterBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkFilterParameter.Value = &value
	return builder
}

func (builder *networkFilterParameterBuilder) Build() (*NetworkFilterParameter, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.networkFilterParameter, nil
}

type networkLabelBuilder struct {
	networkLabel *NetworkLabel
	err          error
}

func NewNetworkLabelBuilder() *networkLabelBuilder {
	return &networkLabelBuilder{networkLabel: &NetworkLabel{}, err: nil}
}

func (builder *networkLabelBuilder) Comment(comment string) *networkLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkLabel.Comment = &comment
	return builder
}

func (builder *networkLabelBuilder) Description(description string) *networkLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkLabel.Description = &description
	return builder
}

func (builder *networkLabelBuilder) HostNic(hostNic *HostNic) *networkLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkLabel.HostNic = hostNic
	return builder
}

func (builder *networkLabelBuilder) Id(id string) *networkLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkLabel.Id = &id
	return builder
}

func (builder *networkLabelBuilder) Name(name string) *networkLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkLabel.Name = &name
	return builder
}

func (builder *networkLabelBuilder) Network(network *Network) *networkLabelBuilder {
	if builder.err != nil {
		return builder
	}

	builder.networkLabel.Network = network
	return builder
}

func (builder *networkLabelBuilder) Build() (*NetworkLabel, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.networkLabel, nil
}

type nfsProfileDetailBuilder struct {
	nfsProfileDetail *NfsProfileDetail
	err              error
}

func NewNfsProfileDetailBuilder() *nfsProfileDetailBuilder {
	return &nfsProfileDetailBuilder{nfsProfileDetail: &NfsProfileDetail{}, err: nil}
}

func (builder *nfsProfileDetailBuilder) NfsServerIp(nfsServerIp string) *nfsProfileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nfsProfileDetail.NfsServerIp = &nfsServerIp
	return builder
}

func (builder *nfsProfileDetailBuilder) ProfileDetails(profileDetails []ProfileDetail) *nfsProfileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nfsProfileDetail.ProfileDetails = profileDetails
	return builder
}

func (builder *nfsProfileDetailBuilder) Build() (*NfsProfileDetail, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.nfsProfileDetail, nil
}

type nicBuilder struct {
	nic *Nic
	err error
}

func NewNicBuilder() *nicBuilder {
	return &nicBuilder{nic: &Nic{}, err: nil}
}

func (builder *nicBuilder) BootProtocol(bootProtocol BootProtocol) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.BootProtocol = bootProtocol
	return builder
}

func (builder *nicBuilder) Comment(comment string) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Comment = &comment
	return builder
}

func (builder *nicBuilder) Description(description string) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Description = &description
	return builder
}

func (builder *nicBuilder) Id(id string) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Id = &id
	return builder
}

func (builder *nicBuilder) InstanceType(instanceType *InstanceType) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.InstanceType = instanceType
	return builder
}

func (builder *nicBuilder) Interface(interface_ NicInterface) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Interface = interface_
	return builder
}

func (builder *nicBuilder) Linked(linked bool) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Linked = &linked
	return builder
}

func (builder *nicBuilder) Mac(mac *Mac) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Mac = mac
	return builder
}

func (builder *nicBuilder) Name(name string) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Name = &name
	return builder
}

func (builder *nicBuilder) Network(network *Network) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Network = network
	return builder
}

func (builder *nicBuilder) NetworkAttachments(networkAttachments []NetworkAttachment) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.NetworkAttachments = networkAttachments
	return builder
}

func (builder *nicBuilder) NetworkFilterParameters(networkFilterParameters []NetworkFilterParameter) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.NetworkFilterParameters = networkFilterParameters
	return builder
}

func (builder *nicBuilder) NetworkLabels(networkLabels []NetworkLabel) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.NetworkLabels = networkLabels
	return builder
}

func (builder *nicBuilder) OnBoot(onBoot bool) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.OnBoot = &onBoot
	return builder
}

func (builder *nicBuilder) Plugged(plugged bool) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Plugged = &plugged
	return builder
}

func (builder *nicBuilder) ReportedDevices(reportedDevices []ReportedDevice) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.ReportedDevices = reportedDevices
	return builder
}

func (builder *nicBuilder) Statistics(statistics []Statistic) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Statistics = statistics
	return builder
}

func (builder *nicBuilder) Template(template *Template) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Template = template
	return builder
}

func (builder *nicBuilder) VirtualFunctionAllowedLabels(virtualFunctionAllowedLabels []NetworkLabel) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.VirtualFunctionAllowedLabels = virtualFunctionAllowedLabels
	return builder
}

func (builder *nicBuilder) VirtualFunctionAllowedNetworks(virtualFunctionAllowedNetworks []Network) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.VirtualFunctionAllowedNetworks = virtualFunctionAllowedNetworks
	return builder
}

func (builder *nicBuilder) Vm(vm *Vm) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Vm = vm
	return builder
}

func (builder *nicBuilder) Vms(vms []Vm) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.Vms = vms
	return builder
}

func (builder *nicBuilder) VnicProfile(vnicProfile *VnicProfile) *nicBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nic.VnicProfile = vnicProfile
	return builder
}

func (builder *nicBuilder) Build() (*Nic, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.nic, nil
}

type nicConfigurationBuilder struct {
	nicConfiguration *NicConfiguration
	err              error
}

func NewNicConfigurationBuilder() *nicConfigurationBuilder {
	return &nicConfigurationBuilder{nicConfiguration: &NicConfiguration{}, err: nil}
}

func (builder *nicConfigurationBuilder) BootProtocol(bootProtocol BootProtocol) *nicConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nicConfiguration.BootProtocol = bootProtocol
	return builder
}

func (builder *nicConfigurationBuilder) Ip(ip *Ip) *nicConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nicConfiguration.Ip = ip
	return builder
}

func (builder *nicConfigurationBuilder) Ipv6(ipv6 *Ip) *nicConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nicConfiguration.Ipv6 = ipv6
	return builder
}

func (builder *nicConfigurationBuilder) Ipv6BootProtocol(ipv6BootProtocol BootProtocol) *nicConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nicConfiguration.Ipv6BootProtocol = ipv6BootProtocol
	return builder
}

func (builder *nicConfigurationBuilder) Name(name string) *nicConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nicConfiguration.Name = &name
	return builder
}

func (builder *nicConfigurationBuilder) OnBoot(onBoot bool) *nicConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.nicConfiguration.OnBoot = &onBoot
	return builder
}

func (builder *nicConfigurationBuilder) Build() (*NicConfiguration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.nicConfiguration, nil
}

type numaNodeBuilder struct {
	numaNode *NumaNode
	err      error
}

func NewNumaNodeBuilder() *numaNodeBuilder {
	return &numaNodeBuilder{numaNode: &NumaNode{}, err: nil}
}

func (builder *numaNodeBuilder) Comment(comment string) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Comment = &comment
	return builder
}

func (builder *numaNodeBuilder) Cpu(cpu *Cpu) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Cpu = cpu
	return builder
}

func (builder *numaNodeBuilder) Description(description string) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Description = &description
	return builder
}

func (builder *numaNodeBuilder) Host(host *Host) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Host = host
	return builder
}

func (builder *numaNodeBuilder) Id(id string) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Id = &id
	return builder
}

func (builder *numaNodeBuilder) Index(index int64) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Index = &index
	return builder
}

func (builder *numaNodeBuilder) Memory(memory int64) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Memory = &memory
	return builder
}

func (builder *numaNodeBuilder) Name(name string) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Name = &name
	return builder
}

func (builder *numaNodeBuilder) NodeDistance(nodeDistance string) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.NodeDistance = &nodeDistance
	return builder
}

func (builder *numaNodeBuilder) Statistics(statistics []Statistic) *numaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNode.Statistics = statistics
	return builder
}

func (builder *numaNodeBuilder) Build() (*NumaNode, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.numaNode, nil
}

type numaNodePinBuilder struct {
	numaNodePin *NumaNodePin
	err         error
}

func NewNumaNodePinBuilder() *numaNodePinBuilder {
	return &numaNodePinBuilder{numaNodePin: &NumaNodePin{}, err: nil}
}

func (builder *numaNodePinBuilder) HostNumaNode(hostNumaNode *NumaNode) *numaNodePinBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNodePin.HostNumaNode = hostNumaNode
	return builder
}

func (builder *numaNodePinBuilder) Index(index int64) *numaNodePinBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNodePin.Index = &index
	return builder
}

func (builder *numaNodePinBuilder) Pinned(pinned bool) *numaNodePinBuilder {
	if builder.err != nil {
		return builder
	}

	builder.numaNodePin.Pinned = &pinned
	return builder
}

func (builder *numaNodePinBuilder) Build() (*NumaNodePin, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.numaNodePin, nil
}

type openStackImageBuilder struct {
	openStackImage *OpenStackImage
	err            error
}

func NewOpenStackImageBuilder() *openStackImageBuilder {
	return &openStackImageBuilder{openStackImage: &OpenStackImage{}, err: nil}
}

func (builder *openStackImageBuilder) Comment(comment string) *openStackImageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImage.Comment = &comment
	return builder
}

func (builder *openStackImageBuilder) Description(description string) *openStackImageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImage.Description = &description
	return builder
}

func (builder *openStackImageBuilder) Id(id string) *openStackImageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImage.Id = &id
	return builder
}

func (builder *openStackImageBuilder) Name(name string) *openStackImageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImage.Name = &name
	return builder
}

func (builder *openStackImageBuilder) OpenstackImageProvider(openstackImageProvider *OpenStackImageProvider) *openStackImageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImage.OpenstackImageProvider = openstackImageProvider
	return builder
}

func (builder *openStackImageBuilder) Build() (*OpenStackImage, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackImage, nil
}

type openStackImageProviderBuilder struct {
	openStackImageProvider *OpenStackImageProvider
	err                    error
}

func NewOpenStackImageProviderBuilder() *openStackImageProviderBuilder {
	return &openStackImageProviderBuilder{openStackImageProvider: &OpenStackImageProvider{}, err: nil}
}

func (builder *openStackImageProviderBuilder) AuthenticationUrl(authenticationUrl string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.AuthenticationUrl = &authenticationUrl
	return builder
}

func (builder *openStackImageProviderBuilder) Certificates(certificates []Certificate) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Certificates = certificates
	return builder
}

func (builder *openStackImageProviderBuilder) Comment(comment string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Comment = &comment
	return builder
}

func (builder *openStackImageProviderBuilder) Description(description string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Description = &description
	return builder
}

func (builder *openStackImageProviderBuilder) Id(id string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Id = &id
	return builder
}

func (builder *openStackImageProviderBuilder) Images(images []OpenStackImage) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Images = images
	return builder
}

func (builder *openStackImageProviderBuilder) Name(name string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Name = &name
	return builder
}

func (builder *openStackImageProviderBuilder) Password(password string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Password = &password
	return builder
}

func (builder *openStackImageProviderBuilder) Properties(properties []Property) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Properties = properties
	return builder
}

func (builder *openStackImageProviderBuilder) RequiresAuthentication(requiresAuthentication bool) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.RequiresAuthentication = &requiresAuthentication
	return builder
}

func (builder *openStackImageProviderBuilder) TenantName(tenantName string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.TenantName = &tenantName
	return builder
}

func (builder *openStackImageProviderBuilder) Url(url string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Url = &url
	return builder
}

func (builder *openStackImageProviderBuilder) Username(username string) *openStackImageProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackImageProvider.Username = &username
	return builder
}

func (builder *openStackImageProviderBuilder) Build() (*OpenStackImageProvider, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackImageProvider, nil
}

type openStackNetworkBuilder struct {
	openStackNetwork *OpenStackNetwork
	err              error
}

func NewOpenStackNetworkBuilder() *openStackNetworkBuilder {
	return &openStackNetworkBuilder{openStackNetwork: &OpenStackNetwork{}, err: nil}
}

func (builder *openStackNetworkBuilder) Comment(comment string) *openStackNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetwork.Comment = &comment
	return builder
}

func (builder *openStackNetworkBuilder) Description(description string) *openStackNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetwork.Description = &description
	return builder
}

func (builder *openStackNetworkBuilder) Id(id string) *openStackNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetwork.Id = &id
	return builder
}

func (builder *openStackNetworkBuilder) Name(name string) *openStackNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetwork.Name = &name
	return builder
}

func (builder *openStackNetworkBuilder) OpenstackNetworkProvider(openstackNetworkProvider *OpenStackNetworkProvider) *openStackNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetwork.OpenstackNetworkProvider = openstackNetworkProvider
	return builder
}

func (builder *openStackNetworkBuilder) Build() (*OpenStackNetwork, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackNetwork, nil
}

type openStackNetworkProviderBuilder struct {
	openStackNetworkProvider *OpenStackNetworkProvider
	err                      error
}

func NewOpenStackNetworkProviderBuilder() *openStackNetworkProviderBuilder {
	return &openStackNetworkProviderBuilder{openStackNetworkProvider: &OpenStackNetworkProvider{}, err: nil}
}

func (builder *openStackNetworkProviderBuilder) AgentConfiguration(agentConfiguration *AgentConfiguration) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.AgentConfiguration = agentConfiguration
	return builder
}

func (builder *openStackNetworkProviderBuilder) AuthenticationUrl(authenticationUrl string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.AuthenticationUrl = &authenticationUrl
	return builder
}

func (builder *openStackNetworkProviderBuilder) Certificates(certificates []Certificate) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Certificates = certificates
	return builder
}

func (builder *openStackNetworkProviderBuilder) Comment(comment string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Comment = &comment
	return builder
}

func (builder *openStackNetworkProviderBuilder) Description(description string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Description = &description
	return builder
}

func (builder *openStackNetworkProviderBuilder) Id(id string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Id = &id
	return builder
}

func (builder *openStackNetworkProviderBuilder) Name(name string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Name = &name
	return builder
}

func (builder *openStackNetworkProviderBuilder) Networks(networks []OpenStackNetwork) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Networks = networks
	return builder
}

func (builder *openStackNetworkProviderBuilder) Password(password string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Password = &password
	return builder
}

func (builder *openStackNetworkProviderBuilder) PluginType(pluginType NetworkPluginType) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.PluginType = pluginType
	return builder
}

func (builder *openStackNetworkProviderBuilder) Properties(properties []Property) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Properties = properties
	return builder
}

func (builder *openStackNetworkProviderBuilder) ReadOnly(readOnly bool) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.ReadOnly = &readOnly
	return builder
}

func (builder *openStackNetworkProviderBuilder) RequiresAuthentication(requiresAuthentication bool) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.RequiresAuthentication = &requiresAuthentication
	return builder
}

func (builder *openStackNetworkProviderBuilder) Subnets(subnets []OpenStackSubnet) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Subnets = subnets
	return builder
}

func (builder *openStackNetworkProviderBuilder) TenantName(tenantName string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.TenantName = &tenantName
	return builder
}

func (builder *openStackNetworkProviderBuilder) Type(type_ OpenStackNetworkProviderType) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Type = type_
	return builder
}

func (builder *openStackNetworkProviderBuilder) Url(url string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Url = &url
	return builder
}

func (builder *openStackNetworkProviderBuilder) Username(username string) *openStackNetworkProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackNetworkProvider.Username = &username
	return builder
}

func (builder *openStackNetworkProviderBuilder) Build() (*OpenStackNetworkProvider, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackNetworkProvider, nil
}

type openStackProviderBuilder struct {
	openStackProvider *OpenStackProvider
	err               error
}

func NewOpenStackProviderBuilder() *openStackProviderBuilder {
	return &openStackProviderBuilder{openStackProvider: &OpenStackProvider{}, err: nil}
}

func (builder *openStackProviderBuilder) AuthenticationUrl(authenticationUrl string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.AuthenticationUrl = &authenticationUrl
	return builder
}

func (builder *openStackProviderBuilder) Comment(comment string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Comment = &comment
	return builder
}

func (builder *openStackProviderBuilder) Description(description string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Description = &description
	return builder
}

func (builder *openStackProviderBuilder) Id(id string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Id = &id
	return builder
}

func (builder *openStackProviderBuilder) Name(name string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Name = &name
	return builder
}

func (builder *openStackProviderBuilder) Password(password string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Password = &password
	return builder
}

func (builder *openStackProviderBuilder) Properties(properties []Property) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Properties = properties
	return builder
}

func (builder *openStackProviderBuilder) RequiresAuthentication(requiresAuthentication bool) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.RequiresAuthentication = &requiresAuthentication
	return builder
}

func (builder *openStackProviderBuilder) TenantName(tenantName string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.TenantName = &tenantName
	return builder
}

func (builder *openStackProviderBuilder) Url(url string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Url = &url
	return builder
}

func (builder *openStackProviderBuilder) Username(username string) *openStackProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackProvider.Username = &username
	return builder
}

func (builder *openStackProviderBuilder) Build() (*OpenStackProvider, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackProvider, nil
}

type openStackSubnetBuilder struct {
	openStackSubnet *OpenStackSubnet
	err             error
}

func NewOpenStackSubnetBuilder() *openStackSubnetBuilder {
	return &openStackSubnetBuilder{openStackSubnet: &OpenStackSubnet{}, err: nil}
}

func (builder *openStackSubnetBuilder) Cidr(cidr string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.Cidr = &cidr
	return builder
}

func (builder *openStackSubnetBuilder) Comment(comment string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.Comment = &comment
	return builder
}

func (builder *openStackSubnetBuilder) Description(description string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.Description = &description
	return builder
}

func (builder *openStackSubnetBuilder) DnsServers(dnsServers []string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.DnsServers = dnsServers
	return builder
}

func (builder *openStackSubnetBuilder) Gateway(gateway string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.Gateway = &gateway
	return builder
}

func (builder *openStackSubnetBuilder) Id(id string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.Id = &id
	return builder
}

func (builder *openStackSubnetBuilder) IpVersion(ipVersion string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.IpVersion = &ipVersion
	return builder
}

func (builder *openStackSubnetBuilder) Name(name string) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.Name = &name
	return builder
}

func (builder *openStackSubnetBuilder) OpenstackNetwork(openstackNetwork *OpenStackNetwork) *openStackSubnetBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackSubnet.OpenstackNetwork = openstackNetwork
	return builder
}

func (builder *openStackSubnetBuilder) Build() (*OpenStackSubnet, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackSubnet, nil
}

type openStackVolumeProviderBuilder struct {
	openStackVolumeProvider *OpenStackVolumeProvider
	err                     error
}

func NewOpenStackVolumeProviderBuilder() *openStackVolumeProviderBuilder {
	return &openStackVolumeProviderBuilder{openStackVolumeProvider: &OpenStackVolumeProvider{}, err: nil}
}

func (builder *openStackVolumeProviderBuilder) AuthenticationKeys(authenticationKeys []OpenstackVolumeAuthenticationKey) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.AuthenticationKeys = authenticationKeys
	return builder
}

func (builder *openStackVolumeProviderBuilder) AuthenticationUrl(authenticationUrl string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.AuthenticationUrl = &authenticationUrl
	return builder
}

func (builder *openStackVolumeProviderBuilder) Certificates(certificates []Certificate) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Certificates = certificates
	return builder
}

func (builder *openStackVolumeProviderBuilder) Comment(comment string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Comment = &comment
	return builder
}

func (builder *openStackVolumeProviderBuilder) DataCenter(dataCenter *DataCenter) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.DataCenter = dataCenter
	return builder
}

func (builder *openStackVolumeProviderBuilder) Description(description string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Description = &description
	return builder
}

func (builder *openStackVolumeProviderBuilder) Id(id string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Id = &id
	return builder
}

func (builder *openStackVolumeProviderBuilder) Name(name string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Name = &name
	return builder
}

func (builder *openStackVolumeProviderBuilder) Password(password string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Password = &password
	return builder
}

func (builder *openStackVolumeProviderBuilder) Properties(properties []Property) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Properties = properties
	return builder
}

func (builder *openStackVolumeProviderBuilder) RequiresAuthentication(requiresAuthentication bool) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.RequiresAuthentication = &requiresAuthentication
	return builder
}

func (builder *openStackVolumeProviderBuilder) TenantName(tenantName string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.TenantName = &tenantName
	return builder
}

func (builder *openStackVolumeProviderBuilder) Url(url string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Url = &url
	return builder
}

func (builder *openStackVolumeProviderBuilder) Username(username string) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.Username = &username
	return builder
}

func (builder *openStackVolumeProviderBuilder) VolumeTypes(volumeTypes []OpenStackVolumeType) *openStackVolumeProviderBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeProvider.VolumeTypes = volumeTypes
	return builder
}

func (builder *openStackVolumeProviderBuilder) Build() (*OpenStackVolumeProvider, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackVolumeProvider, nil
}

type openStackVolumeTypeBuilder struct {
	openStackVolumeType *OpenStackVolumeType
	err                 error
}

func NewOpenStackVolumeTypeBuilder() *openStackVolumeTypeBuilder {
	return &openStackVolumeTypeBuilder{openStackVolumeType: &OpenStackVolumeType{}, err: nil}
}

func (builder *openStackVolumeTypeBuilder) Comment(comment string) *openStackVolumeTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeType.Comment = &comment
	return builder
}

func (builder *openStackVolumeTypeBuilder) Description(description string) *openStackVolumeTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeType.Description = &description
	return builder
}

func (builder *openStackVolumeTypeBuilder) Id(id string) *openStackVolumeTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeType.Id = &id
	return builder
}

func (builder *openStackVolumeTypeBuilder) Name(name string) *openStackVolumeTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeType.Name = &name
	return builder
}

func (builder *openStackVolumeTypeBuilder) OpenstackVolumeProvider(openstackVolumeProvider *OpenStackVolumeProvider) *openStackVolumeTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeType.OpenstackVolumeProvider = openstackVolumeProvider
	return builder
}

func (builder *openStackVolumeTypeBuilder) Properties(properties []Property) *openStackVolumeTypeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openStackVolumeType.Properties = properties
	return builder
}

func (builder *openStackVolumeTypeBuilder) Build() (*OpenStackVolumeType, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openStackVolumeType, nil
}

type openstackVolumeAuthenticationKeyBuilder struct {
	openstackVolumeAuthenticationKey *OpenstackVolumeAuthenticationKey
	err                              error
}

func NewOpenstackVolumeAuthenticationKeyBuilder() *openstackVolumeAuthenticationKeyBuilder {
	return &openstackVolumeAuthenticationKeyBuilder{openstackVolumeAuthenticationKey: &OpenstackVolumeAuthenticationKey{}, err: nil}
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Comment(comment string) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.Comment = &comment
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) CreationDate(creationDate time.Time) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.CreationDate = creationDate
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Description(description string) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.Description = &description
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Id(id string) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.Id = &id
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Name(name string) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.Name = &name
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) OpenstackVolumeProvider(openstackVolumeProvider *OpenStackVolumeProvider) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.OpenstackVolumeProvider = openstackVolumeProvider
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) UsageType(usageType OpenstackVolumeAuthenticationKeyUsageType) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.UsageType = usageType
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Uuid(uuid string) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.Uuid = &uuid
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Value(value string) *openstackVolumeAuthenticationKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.openstackVolumeAuthenticationKey.Value = &value
	return builder
}

func (builder *openstackVolumeAuthenticationKeyBuilder) Build() (*OpenstackVolumeAuthenticationKey, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.openstackVolumeAuthenticationKey, nil
}

type operatingSystemBuilder struct {
	operatingSystem *OperatingSystem
	err             error
}

func NewOperatingSystemBuilder() *operatingSystemBuilder {
	return &operatingSystemBuilder{operatingSystem: &OperatingSystem{}, err: nil}
}

func (builder *operatingSystemBuilder) Boot(boot *Boot) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.Boot = boot
	return builder
}

func (builder *operatingSystemBuilder) Cmdline(cmdline string) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.Cmdline = &cmdline
	return builder
}

func (builder *operatingSystemBuilder) CustomKernelCmdline(customKernelCmdline string) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.CustomKernelCmdline = &customKernelCmdline
	return builder
}

func (builder *operatingSystemBuilder) Initrd(initrd string) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.Initrd = &initrd
	return builder
}

func (builder *operatingSystemBuilder) Kernel(kernel string) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.Kernel = &kernel
	return builder
}

func (builder *operatingSystemBuilder) ReportedKernelCmdline(reportedKernelCmdline string) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.ReportedKernelCmdline = &reportedKernelCmdline
	return builder
}

func (builder *operatingSystemBuilder) Type(type_ string) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.Type = &type_
	return builder
}

func (builder *operatingSystemBuilder) Version(version *Version) *operatingSystemBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystem.Version = version
	return builder
}

func (builder *operatingSystemBuilder) Build() (*OperatingSystem, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.operatingSystem, nil
}

type operatingSystemInfoBuilder struct {
	operatingSystemInfo *OperatingSystemInfo
	err                 error
}

func NewOperatingSystemInfoBuilder() *operatingSystemInfoBuilder {
	return &operatingSystemInfoBuilder{operatingSystemInfo: &OperatingSystemInfo{}, err: nil}
}

func (builder *operatingSystemInfoBuilder) Comment(comment string) *operatingSystemInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystemInfo.Comment = &comment
	return builder
}

func (builder *operatingSystemInfoBuilder) Description(description string) *operatingSystemInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystemInfo.Description = &description
	return builder
}

func (builder *operatingSystemInfoBuilder) Id(id string) *operatingSystemInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystemInfo.Id = &id
	return builder
}

func (builder *operatingSystemInfoBuilder) LargeIcon(largeIcon *Icon) *operatingSystemInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystemInfo.LargeIcon = largeIcon
	return builder
}

func (builder *operatingSystemInfoBuilder) Name(name string) *operatingSystemInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystemInfo.Name = &name
	return builder
}

func (builder *operatingSystemInfoBuilder) SmallIcon(smallIcon *Icon) *operatingSystemInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.operatingSystemInfo.SmallIcon = smallIcon
	return builder
}

func (builder *operatingSystemInfoBuilder) Build() (*OperatingSystemInfo, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.operatingSystemInfo, nil
}

type optionBuilder struct {
	option *Option
	err    error
}

func NewOptionBuilder() *optionBuilder {
	return &optionBuilder{option: &Option{}, err: nil}
}

func (builder *optionBuilder) Name(name string) *optionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.option.Name = &name
	return builder
}

func (builder *optionBuilder) Type(type_ string) *optionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.option.Type = &type_
	return builder
}

func (builder *optionBuilder) Value(value string) *optionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.option.Value = &value
	return builder
}

func (builder *optionBuilder) Build() (*Option, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.option, nil
}

type packageBuilder struct {
	package_ *Package
	err      error
}

func NewPackageBuilder() *packageBuilder {
	return &packageBuilder{package_: &Package{}, err: nil}
}

func (builder *packageBuilder) Name(name string) *packageBuilder {
	if builder.err != nil {
		return builder
	}

	builder.package_.Name = &name
	return builder
}

func (builder *packageBuilder) Build() (*Package, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.package_, nil
}

type payloadBuilder struct {
	payload *Payload
	err     error
}

func NewPayloadBuilder() *payloadBuilder {
	return &payloadBuilder{payload: &Payload{}, err: nil}
}

func (builder *payloadBuilder) Files(files []File) *payloadBuilder {
	if builder.err != nil {
		return builder
	}

	builder.payload.Files = files
	return builder
}

func (builder *payloadBuilder) Type(type_ VmDeviceType) *payloadBuilder {
	if builder.err != nil {
		return builder
	}

	builder.payload.Type = type_
	return builder
}

func (builder *payloadBuilder) VolumeId(volumeId string) *payloadBuilder {
	if builder.err != nil {
		return builder
	}

	builder.payload.VolumeId = &volumeId
	return builder
}

func (builder *payloadBuilder) Build() (*Payload, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.payload, nil
}

type permissionBuilder struct {
	permission *Permission
	err        error
}

func NewPermissionBuilder() *permissionBuilder {
	return &permissionBuilder{permission: &Permission{}, err: nil}
}

func (builder *permissionBuilder) Cluster(cluster *Cluster) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Cluster = cluster
	return builder
}

func (builder *permissionBuilder) Comment(comment string) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Comment = &comment
	return builder
}

func (builder *permissionBuilder) DataCenter(dataCenter *DataCenter) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.DataCenter = dataCenter
	return builder
}

func (builder *permissionBuilder) Description(description string) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Description = &description
	return builder
}

func (builder *permissionBuilder) Disk(disk *Disk) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Disk = disk
	return builder
}

func (builder *permissionBuilder) Group(group *Group) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Group = group
	return builder
}

func (builder *permissionBuilder) Host(host *Host) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Host = host
	return builder
}

func (builder *permissionBuilder) Id(id string) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Id = &id
	return builder
}

func (builder *permissionBuilder) Name(name string) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Name = &name
	return builder
}

func (builder *permissionBuilder) Role(role *Role) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Role = role
	return builder
}

func (builder *permissionBuilder) StorageDomain(storageDomain *StorageDomain) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.StorageDomain = storageDomain
	return builder
}

func (builder *permissionBuilder) Template(template *Template) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Template = template
	return builder
}

func (builder *permissionBuilder) User(user *User) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.User = user
	return builder
}

func (builder *permissionBuilder) Vm(vm *Vm) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.Vm = vm
	return builder
}

func (builder *permissionBuilder) VmPool(vmPool *VmPool) *permissionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permission.VmPool = vmPool
	return builder
}

func (builder *permissionBuilder) Build() (*Permission, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.permission, nil
}

type permitBuilder struct {
	permit *Permit
	err    error
}

func NewPermitBuilder() *permitBuilder {
	return &permitBuilder{permit: &Permit{}, err: nil}
}

func (builder *permitBuilder) Administrative(administrative bool) *permitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permit.Administrative = &administrative
	return builder
}

func (builder *permitBuilder) Comment(comment string) *permitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permit.Comment = &comment
	return builder
}

func (builder *permitBuilder) Description(description string) *permitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permit.Description = &description
	return builder
}

func (builder *permitBuilder) Id(id string) *permitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permit.Id = &id
	return builder
}

func (builder *permitBuilder) Name(name string) *permitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permit.Name = &name
	return builder
}

func (builder *permitBuilder) Role(role *Role) *permitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.permit.Role = role
	return builder
}

func (builder *permitBuilder) Build() (*Permit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.permit, nil
}

type pmProxyBuilder struct {
	pmProxy *PmProxy
	err     error
}

func NewPmProxyBuilder() *pmProxyBuilder {
	return &pmProxyBuilder{pmProxy: &PmProxy{}, err: nil}
}

func (builder *pmProxyBuilder) Type(type_ PmProxyType) *pmProxyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.pmProxy.Type = type_
	return builder
}

func (builder *pmProxyBuilder) Build() (*PmProxy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.pmProxy, nil
}

type portMirroringBuilder struct {
	portMirroring *PortMirroring
	err           error
}

func NewPortMirroringBuilder() *portMirroringBuilder {
	return &portMirroringBuilder{portMirroring: &PortMirroring{}, err: nil}
}

func (builder *portMirroringBuilder) Build() (*PortMirroring, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.portMirroring, nil
}

type powerManagementBuilder struct {
	powerManagement *PowerManagement
	err             error
}

func NewPowerManagementBuilder() *powerManagementBuilder {
	return &powerManagementBuilder{powerManagement: &PowerManagement{}, err: nil}
}

func (builder *powerManagementBuilder) Address(address string) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Address = &address
	return builder
}

func (builder *powerManagementBuilder) Agents(agents []Agent) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Agents = agents
	return builder
}

func (builder *powerManagementBuilder) AutomaticPmEnabled(automaticPmEnabled bool) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.AutomaticPmEnabled = &automaticPmEnabled
	return builder
}

func (builder *powerManagementBuilder) Enabled(enabled bool) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Enabled = &enabled
	return builder
}

func (builder *powerManagementBuilder) KdumpDetection(kdumpDetection bool) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.KdumpDetection = &kdumpDetection
	return builder
}

func (builder *powerManagementBuilder) Options(options []Option) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Options = options
	return builder
}

func (builder *powerManagementBuilder) Password(password string) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Password = &password
	return builder
}

func (builder *powerManagementBuilder) PmProxies(pmProxies []PmProxy) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.PmProxies = pmProxies
	return builder
}

func (builder *powerManagementBuilder) Status(status PowerManagementStatus) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Status = status
	return builder
}

func (builder *powerManagementBuilder) Type(type_ string) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Type = &type_
	return builder
}

func (builder *powerManagementBuilder) Username(username string) *powerManagementBuilder {
	if builder.err != nil {
		return builder
	}

	builder.powerManagement.Username = &username
	return builder
}

func (builder *powerManagementBuilder) Build() (*PowerManagement, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.powerManagement, nil
}

type productBuilder struct {
	product *Product
	err     error
}

func NewProductBuilder() *productBuilder {
	return &productBuilder{product: &Product{}, err: nil}
}

func (builder *productBuilder) Comment(comment string) *productBuilder {
	if builder.err != nil {
		return builder
	}

	builder.product.Comment = &comment
	return builder
}

func (builder *productBuilder) Description(description string) *productBuilder {
	if builder.err != nil {
		return builder
	}

	builder.product.Description = &description
	return builder
}

func (builder *productBuilder) Id(id string) *productBuilder {
	if builder.err != nil {
		return builder
	}

	builder.product.Id = &id
	return builder
}

func (builder *productBuilder) Name(name string) *productBuilder {
	if builder.err != nil {
		return builder
	}

	builder.product.Name = &name
	return builder
}

func (builder *productBuilder) Build() (*Product, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.product, nil
}

type productInfoBuilder struct {
	productInfo *ProductInfo
	err         error
}

func NewProductInfoBuilder() *productInfoBuilder {
	return &productInfoBuilder{productInfo: &ProductInfo{}, err: nil}
}

func (builder *productInfoBuilder) Name(name string) *productInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.productInfo.Name = &name
	return builder
}

func (builder *productInfoBuilder) Vendor(vendor string) *productInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.productInfo.Vendor = &vendor
	return builder
}

func (builder *productInfoBuilder) Version(version *Version) *productInfoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.productInfo.Version = version
	return builder
}

func (builder *productInfoBuilder) Build() (*ProductInfo, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.productInfo, nil
}

type profileDetailBuilder struct {
	profileDetail *ProfileDetail
	err           error
}

func NewProfileDetailBuilder() *profileDetailBuilder {
	return &profileDetailBuilder{profileDetail: &ProfileDetail{}, err: nil}
}

func (builder *profileDetailBuilder) BlockStatistics(blockStatistics []BlockStatistic) *profileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.profileDetail.BlockStatistics = blockStatistics
	return builder
}

func (builder *profileDetailBuilder) Duration(duration int64) *profileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.profileDetail.Duration = &duration
	return builder
}

func (builder *profileDetailBuilder) FopStatistics(fopStatistics []FopStatistic) *profileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.profileDetail.FopStatistics = fopStatistics
	return builder
}

func (builder *profileDetailBuilder) ProfileType(profileType string) *profileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.profileDetail.ProfileType = &profileType
	return builder
}

func (builder *profileDetailBuilder) Statistics(statistics []Statistic) *profileDetailBuilder {
	if builder.err != nil {
		return builder
	}

	builder.profileDetail.Statistics = statistics
	return builder
}

func (builder *profileDetailBuilder) Build() (*ProfileDetail, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.profileDetail, nil
}

type propertyBuilder struct {
	property *Property
	err      error
}

func NewPropertyBuilder() *propertyBuilder {
	return &propertyBuilder{property: &Property{}, err: nil}
}

func (builder *propertyBuilder) Name(name string) *propertyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.property.Name = &name
	return builder
}

func (builder *propertyBuilder) Value(value string) *propertyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.property.Value = &value
	return builder
}

func (builder *propertyBuilder) Build() (*Property, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.property, nil
}

type proxyTicketBuilder struct {
	proxyTicket *ProxyTicket
	err         error
}

func NewProxyTicketBuilder() *proxyTicketBuilder {
	return &proxyTicketBuilder{proxyTicket: &ProxyTicket{}, err: nil}
}

func (builder *proxyTicketBuilder) Value(value string) *proxyTicketBuilder {
	if builder.err != nil {
		return builder
	}

	builder.proxyTicket.Value = &value
	return builder
}

func (builder *proxyTicketBuilder) Build() (*ProxyTicket, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.proxyTicket, nil
}

type qosBuilder struct {
	qos *Qos
	err error
}

func NewQosBuilder() *qosBuilder {
	return &qosBuilder{qos: &Qos{}, err: nil}
}

func (builder *qosBuilder) Comment(comment string) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.Comment = &comment
	return builder
}

func (builder *qosBuilder) CpuLimit(cpuLimit int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.CpuLimit = &cpuLimit
	return builder
}

func (builder *qosBuilder) DataCenter(dataCenter *DataCenter) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.DataCenter = dataCenter
	return builder
}

func (builder *qosBuilder) Description(description string) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.Description = &description
	return builder
}

func (builder *qosBuilder) Id(id string) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.Id = &id
	return builder
}

func (builder *qosBuilder) InboundAverage(inboundAverage int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.InboundAverage = &inboundAverage
	return builder
}

func (builder *qosBuilder) InboundBurst(inboundBurst int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.InboundBurst = &inboundBurst
	return builder
}

func (builder *qosBuilder) InboundPeak(inboundPeak int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.InboundPeak = &inboundPeak
	return builder
}

func (builder *qosBuilder) MaxIops(maxIops int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.MaxIops = &maxIops
	return builder
}

func (builder *qosBuilder) MaxReadIops(maxReadIops int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.MaxReadIops = &maxReadIops
	return builder
}

func (builder *qosBuilder) MaxReadThroughput(maxReadThroughput int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.MaxReadThroughput = &maxReadThroughput
	return builder
}

func (builder *qosBuilder) MaxThroughput(maxThroughput int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.MaxThroughput = &maxThroughput
	return builder
}

func (builder *qosBuilder) MaxWriteIops(maxWriteIops int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.MaxWriteIops = &maxWriteIops
	return builder
}

func (builder *qosBuilder) MaxWriteThroughput(maxWriteThroughput int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.MaxWriteThroughput = &maxWriteThroughput
	return builder
}

func (builder *qosBuilder) Name(name string) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.Name = &name
	return builder
}

func (builder *qosBuilder) OutboundAverage(outboundAverage int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.OutboundAverage = &outboundAverage
	return builder
}

func (builder *qosBuilder) OutboundAverageLinkshare(outboundAverageLinkshare int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.OutboundAverageLinkshare = &outboundAverageLinkshare
	return builder
}

func (builder *qosBuilder) OutboundAverageRealtime(outboundAverageRealtime int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.OutboundAverageRealtime = &outboundAverageRealtime
	return builder
}

func (builder *qosBuilder) OutboundAverageUpperlimit(outboundAverageUpperlimit int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.OutboundAverageUpperlimit = &outboundAverageUpperlimit
	return builder
}

func (builder *qosBuilder) OutboundBurst(outboundBurst int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.OutboundBurst = &outboundBurst
	return builder
}

func (builder *qosBuilder) OutboundPeak(outboundPeak int64) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.OutboundPeak = &outboundPeak
	return builder
}

func (builder *qosBuilder) Type(type_ QosType) *qosBuilder {
	if builder.err != nil {
		return builder
	}

	builder.qos.Type = type_
	return builder
}

func (builder *qosBuilder) Build() (*Qos, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.qos, nil
}

type quotaBuilder struct {
	quota *Quota
	err   error
}

func NewQuotaBuilder() *quotaBuilder {
	return &quotaBuilder{quota: &Quota{}, err: nil}
}

func (builder *quotaBuilder) ClusterHardLimitPct(clusterHardLimitPct int64) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.ClusterHardLimitPct = &clusterHardLimitPct
	return builder
}

func (builder *quotaBuilder) ClusterSoftLimitPct(clusterSoftLimitPct int64) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.ClusterSoftLimitPct = &clusterSoftLimitPct
	return builder
}

func (builder *quotaBuilder) Comment(comment string) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Comment = &comment
	return builder
}

func (builder *quotaBuilder) DataCenter(dataCenter *DataCenter) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.DataCenter = dataCenter
	return builder
}

func (builder *quotaBuilder) Description(description string) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Description = &description
	return builder
}

func (builder *quotaBuilder) Disks(disks []Disk) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Disks = disks
	return builder
}

func (builder *quotaBuilder) Id(id string) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Id = &id
	return builder
}

func (builder *quotaBuilder) Name(name string) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Name = &name
	return builder
}

func (builder *quotaBuilder) Permissions(permissions []Permission) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Permissions = permissions
	return builder
}

func (builder *quotaBuilder) QuotaClusterLimits(quotaClusterLimits []QuotaClusterLimit) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.QuotaClusterLimits = quotaClusterLimits
	return builder
}

func (builder *quotaBuilder) QuotaStorageLimits(quotaStorageLimits []QuotaStorageLimit) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.QuotaStorageLimits = quotaStorageLimits
	return builder
}

func (builder *quotaBuilder) StorageHardLimitPct(storageHardLimitPct int64) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.StorageHardLimitPct = &storageHardLimitPct
	return builder
}

func (builder *quotaBuilder) StorageSoftLimitPct(storageSoftLimitPct int64) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.StorageSoftLimitPct = &storageSoftLimitPct
	return builder
}

func (builder *quotaBuilder) Users(users []User) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Users = users
	return builder
}

func (builder *quotaBuilder) Vms(vms []Vm) *quotaBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quota.Vms = vms
	return builder
}

func (builder *quotaBuilder) Build() (*Quota, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.quota, nil
}

type quotaClusterLimitBuilder struct {
	quotaClusterLimit *QuotaClusterLimit
	err               error
}

func NewQuotaClusterLimitBuilder() *quotaClusterLimitBuilder {
	return &quotaClusterLimitBuilder{quotaClusterLimit: &QuotaClusterLimit{}, err: nil}
}

func (builder *quotaClusterLimitBuilder) Cluster(cluster *Cluster) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.Cluster = cluster
	return builder
}

func (builder *quotaClusterLimitBuilder) Comment(comment string) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.Comment = &comment
	return builder
}

func (builder *quotaClusterLimitBuilder) Description(description string) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.Description = &description
	return builder
}

func (builder *quotaClusterLimitBuilder) Id(id string) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.Id = &id
	return builder
}

func (builder *quotaClusterLimitBuilder) MemoryLimit(memoryLimit float64) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.MemoryLimit = &memoryLimit
	return builder
}

func (builder *quotaClusterLimitBuilder) MemoryUsage(memoryUsage float64) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.MemoryUsage = &memoryUsage
	return builder
}

func (builder *quotaClusterLimitBuilder) Name(name string) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.Name = &name
	return builder
}

func (builder *quotaClusterLimitBuilder) Quota(quota *Quota) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.Quota = quota
	return builder
}

func (builder *quotaClusterLimitBuilder) VcpuLimit(vcpuLimit int64) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.VcpuLimit = &vcpuLimit
	return builder
}

func (builder *quotaClusterLimitBuilder) VcpuUsage(vcpuUsage int64) *quotaClusterLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaClusterLimit.VcpuUsage = &vcpuUsage
	return builder
}

func (builder *quotaClusterLimitBuilder) Build() (*QuotaClusterLimit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.quotaClusterLimit, nil
}

type quotaStorageLimitBuilder struct {
	quotaStorageLimit *QuotaStorageLimit
	err               error
}

func NewQuotaStorageLimitBuilder() *quotaStorageLimitBuilder {
	return &quotaStorageLimitBuilder{quotaStorageLimit: &QuotaStorageLimit{}, err: nil}
}

func (builder *quotaStorageLimitBuilder) Comment(comment string) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Comment = &comment
	return builder
}

func (builder *quotaStorageLimitBuilder) Description(description string) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Description = &description
	return builder
}

func (builder *quotaStorageLimitBuilder) Id(id string) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Id = &id
	return builder
}

func (builder *quotaStorageLimitBuilder) Limit(limit int64) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Limit = &limit
	return builder
}

func (builder *quotaStorageLimitBuilder) Name(name string) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Name = &name
	return builder
}

func (builder *quotaStorageLimitBuilder) Quota(quota *Quota) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Quota = quota
	return builder
}

func (builder *quotaStorageLimitBuilder) StorageDomain(storageDomain *StorageDomain) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.StorageDomain = storageDomain
	return builder
}

func (builder *quotaStorageLimitBuilder) Usage(usage float64) *quotaStorageLimitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.quotaStorageLimit.Usage = &usage
	return builder
}

func (builder *quotaStorageLimitBuilder) Build() (*QuotaStorageLimit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.quotaStorageLimit, nil
}

type rangeBuilder struct {
	range_ *Range
	err    error
}

func NewRangeBuilder() *rangeBuilder {
	return &rangeBuilder{range_: &Range{}, err: nil}
}

func (builder *rangeBuilder) From(from string) *rangeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.range_.From = &from
	return builder
}

func (builder *rangeBuilder) To(to string) *rangeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.range_.To = &to
	return builder
}

func (builder *rangeBuilder) Build() (*Range, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.range_, nil
}

type rateBuilder struct {
	rate *Rate
	err  error
}

func NewRateBuilder() *rateBuilder {
	return &rateBuilder{rate: &Rate{}, err: nil}
}

func (builder *rateBuilder) Bytes(bytes int64) *rateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.rate.Bytes = &bytes
	return builder
}

func (builder *rateBuilder) Period(period int64) *rateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.rate.Period = &period
	return builder
}

func (builder *rateBuilder) Build() (*Rate, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.rate, nil
}

type reportedConfigurationBuilder struct {
	reportedConfiguration *ReportedConfiguration
	err                   error
}

func NewReportedConfigurationBuilder() *reportedConfigurationBuilder {
	return &reportedConfigurationBuilder{reportedConfiguration: &ReportedConfiguration{}, err: nil}
}

func (builder *reportedConfigurationBuilder) ActualValue(actualValue string) *reportedConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedConfiguration.ActualValue = &actualValue
	return builder
}

func (builder *reportedConfigurationBuilder) ExpectedValue(expectedValue string) *reportedConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedConfiguration.ExpectedValue = &expectedValue
	return builder
}

func (builder *reportedConfigurationBuilder) InSync(inSync bool) *reportedConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedConfiguration.InSync = &inSync
	return builder
}

func (builder *reportedConfigurationBuilder) Name(name string) *reportedConfigurationBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedConfiguration.Name = &name
	return builder
}

func (builder *reportedConfigurationBuilder) Build() (*ReportedConfiguration, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.reportedConfiguration, nil
}

type reportedDeviceBuilder struct {
	reportedDevice *ReportedDevice
	err            error
}

func NewReportedDeviceBuilder() *reportedDeviceBuilder {
	return &reportedDeviceBuilder{reportedDevice: &ReportedDevice{}, err: nil}
}

func (builder *reportedDeviceBuilder) Comment(comment string) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Comment = &comment
	return builder
}

func (builder *reportedDeviceBuilder) Description(description string) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Description = &description
	return builder
}

func (builder *reportedDeviceBuilder) Id(id string) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Id = &id
	return builder
}

func (builder *reportedDeviceBuilder) Ips(ips []Ip) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Ips = ips
	return builder
}

func (builder *reportedDeviceBuilder) Mac(mac *Mac) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Mac = mac
	return builder
}

func (builder *reportedDeviceBuilder) Name(name string) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Name = &name
	return builder
}

func (builder *reportedDeviceBuilder) Type(type_ ReportedDeviceType) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Type = type_
	return builder
}

func (builder *reportedDeviceBuilder) Vm(vm *Vm) *reportedDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.reportedDevice.Vm = vm
	return builder
}

func (builder *reportedDeviceBuilder) Build() (*ReportedDevice, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.reportedDevice, nil
}

type rngDeviceBuilder struct {
	rngDevice *RngDevice
	err       error
}

func NewRngDeviceBuilder() *rngDeviceBuilder {
	return &rngDeviceBuilder{rngDevice: &RngDevice{}, err: nil}
}

func (builder *rngDeviceBuilder) Rate(rate *Rate) *rngDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.rngDevice.Rate = rate
	return builder
}

func (builder *rngDeviceBuilder) Source(source RngSource) *rngDeviceBuilder {
	if builder.err != nil {
		return builder
	}

	builder.rngDevice.Source = source
	return builder
}

func (builder *rngDeviceBuilder) Build() (*RngDevice, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.rngDevice, nil
}

type roleBuilder struct {
	role *Role
	err  error
}

func NewRoleBuilder() *roleBuilder {
	return &roleBuilder{role: &Role{}, err: nil}
}

func (builder *roleBuilder) Administrative(administrative bool) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Administrative = &administrative
	return builder
}

func (builder *roleBuilder) Comment(comment string) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Comment = &comment
	return builder
}

func (builder *roleBuilder) Description(description string) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Description = &description
	return builder
}

func (builder *roleBuilder) Id(id string) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Id = &id
	return builder
}

func (builder *roleBuilder) Mutable(mutable bool) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Mutable = &mutable
	return builder
}

func (builder *roleBuilder) Name(name string) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Name = &name
	return builder
}

func (builder *roleBuilder) Permits(permits []Permit) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.Permits = permits
	return builder
}

func (builder *roleBuilder) User(user *User) *roleBuilder {
	if builder.err != nil {
		return builder
	}

	builder.role.User = user
	return builder
}

func (builder *roleBuilder) Build() (*Role, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.role, nil
}

type schedulingPolicyBuilder struct {
	schedulingPolicy *SchedulingPolicy
	err              error
}

func NewSchedulingPolicyBuilder() *schedulingPolicyBuilder {
	return &schedulingPolicyBuilder{schedulingPolicy: &SchedulingPolicy{}, err: nil}
}

func (builder *schedulingPolicyBuilder) Balances(balances []Balance) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Balances = balances
	return builder
}

func (builder *schedulingPolicyBuilder) Comment(comment string) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Comment = &comment
	return builder
}

func (builder *schedulingPolicyBuilder) DefaultPolicy(defaultPolicy bool) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.DefaultPolicy = &defaultPolicy
	return builder
}

func (builder *schedulingPolicyBuilder) Description(description string) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Description = &description
	return builder
}

func (builder *schedulingPolicyBuilder) Filters(filters []Filter) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Filters = filters
	return builder
}

func (builder *schedulingPolicyBuilder) Id(id string) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Id = &id
	return builder
}

func (builder *schedulingPolicyBuilder) Locked(locked bool) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Locked = &locked
	return builder
}

func (builder *schedulingPolicyBuilder) Name(name string) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Name = &name
	return builder
}

func (builder *schedulingPolicyBuilder) Properties(properties []Property) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Properties = properties
	return builder
}

func (builder *schedulingPolicyBuilder) Weight(weight []Weight) *schedulingPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicy.Weight = weight
	return builder
}

func (builder *schedulingPolicyBuilder) Build() (*SchedulingPolicy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.schedulingPolicy, nil
}

type schedulingPolicyUnitBuilder struct {
	schedulingPolicyUnit *SchedulingPolicyUnit
	err                  error
}

func NewSchedulingPolicyUnitBuilder() *schedulingPolicyUnitBuilder {
	return &schedulingPolicyUnitBuilder{schedulingPolicyUnit: &SchedulingPolicyUnit{}, err: nil}
}

func (builder *schedulingPolicyUnitBuilder) Comment(comment string) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Comment = &comment
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Description(description string) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Description = &description
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Enabled(enabled bool) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Enabled = &enabled
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Id(id string) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Id = &id
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Internal(internal bool) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Internal = &internal
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Name(name string) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Name = &name
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Properties(properties []Property) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Properties = properties
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Type(type_ PolicyUnitType) *schedulingPolicyUnitBuilder {
	if builder.err != nil {
		return builder
	}

	builder.schedulingPolicyUnit.Type = type_
	return builder
}

func (builder *schedulingPolicyUnitBuilder) Build() (*SchedulingPolicyUnit, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.schedulingPolicyUnit, nil
}

type seLinuxBuilder struct {
	seLinux *SeLinux
	err     error
}

func NewSeLinuxBuilder() *seLinuxBuilder {
	return &seLinuxBuilder{seLinux: &SeLinux{}, err: nil}
}

func (builder *seLinuxBuilder) Mode(mode SeLinuxMode) *seLinuxBuilder {
	if builder.err != nil {
		return builder
	}

	builder.seLinux.Mode = mode
	return builder
}

func (builder *seLinuxBuilder) Build() (*SeLinux, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.seLinux, nil
}

type serialNumberBuilder struct {
	serialNumber *SerialNumber
	err          error
}

func NewSerialNumberBuilder() *serialNumberBuilder {
	return &serialNumberBuilder{serialNumber: &SerialNumber{}, err: nil}
}

func (builder *serialNumberBuilder) Policy(policy SerialNumberPolicy) *serialNumberBuilder {
	if builder.err != nil {
		return builder
	}

	builder.serialNumber.Policy = policy
	return builder
}

func (builder *serialNumberBuilder) Value(value string) *serialNumberBuilder {
	if builder.err != nil {
		return builder
	}

	builder.serialNumber.Value = &value
	return builder
}

func (builder *serialNumberBuilder) Build() (*SerialNumber, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.serialNumber, nil
}

type sessionBuilder struct {
	session *Session
	err     error
}

func NewSessionBuilder() *sessionBuilder {
	return &sessionBuilder{session: &Session{}, err: nil}
}

func (builder *sessionBuilder) Comment(comment string) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Comment = &comment
	return builder
}

func (builder *sessionBuilder) ConsoleUser(consoleUser bool) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.ConsoleUser = &consoleUser
	return builder
}

func (builder *sessionBuilder) Description(description string) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Description = &description
	return builder
}

func (builder *sessionBuilder) Id(id string) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Id = &id
	return builder
}

func (builder *sessionBuilder) Ip(ip *Ip) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Ip = ip
	return builder
}

func (builder *sessionBuilder) Name(name string) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Name = &name
	return builder
}

func (builder *sessionBuilder) Protocol(protocol string) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Protocol = &protocol
	return builder
}

func (builder *sessionBuilder) User(user *User) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.User = user
	return builder
}

func (builder *sessionBuilder) Vm(vm *Vm) *sessionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.session.Vm = vm
	return builder
}

func (builder *sessionBuilder) Build() (*Session, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.session, nil
}

type skipIfConnectivityBrokenBuilder struct {
	skipIfConnectivityBroken *SkipIfConnectivityBroken
	err                      error
}

func NewSkipIfConnectivityBrokenBuilder() *skipIfConnectivityBrokenBuilder {
	return &skipIfConnectivityBrokenBuilder{skipIfConnectivityBroken: &SkipIfConnectivityBroken{}, err: nil}
}

func (builder *skipIfConnectivityBrokenBuilder) Enabled(enabled bool) *skipIfConnectivityBrokenBuilder {
	if builder.err != nil {
		return builder
	}

	builder.skipIfConnectivityBroken.Enabled = &enabled
	return builder
}

func (builder *skipIfConnectivityBrokenBuilder) Threshold(threshold int64) *skipIfConnectivityBrokenBuilder {
	if builder.err != nil {
		return builder
	}

	builder.skipIfConnectivityBroken.Threshold = &threshold
	return builder
}

func (builder *skipIfConnectivityBrokenBuilder) Build() (*SkipIfConnectivityBroken, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.skipIfConnectivityBroken, nil
}

type skipIfSdActiveBuilder struct {
	skipIfSdActive *SkipIfSdActive
	err            error
}

func NewSkipIfSdActiveBuilder() *skipIfSdActiveBuilder {
	return &skipIfSdActiveBuilder{skipIfSdActive: &SkipIfSdActive{}, err: nil}
}

func (builder *skipIfSdActiveBuilder) Enabled(enabled bool) *skipIfSdActiveBuilder {
	if builder.err != nil {
		return builder
	}

	builder.skipIfSdActive.Enabled = &enabled
	return builder
}

func (builder *skipIfSdActiveBuilder) Build() (*SkipIfSdActive, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.skipIfSdActive, nil
}

type snapshotBuilder struct {
	snapshot *Snapshot
	err      error
}

func NewSnapshotBuilder() *snapshotBuilder {
	return &snapshotBuilder{snapshot: &Snapshot{}, err: nil}
}

func (builder *snapshotBuilder) AffinityLabels(affinityLabels []AffinityLabel) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.AffinityLabels = affinityLabels
	return builder
}

func (builder *snapshotBuilder) Applications(applications []Application) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Applications = applications
	return builder
}

func (builder *snapshotBuilder) Bios(bios *Bios) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Bios = bios
	return builder
}

func (builder *snapshotBuilder) Cdroms(cdroms []Cdrom) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Cdroms = cdroms
	return builder
}

func (builder *snapshotBuilder) Cluster(cluster *Cluster) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Cluster = cluster
	return builder
}

func (builder *snapshotBuilder) Comment(comment string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Comment = &comment
	return builder
}

func (builder *snapshotBuilder) Console(console *Console) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Console = console
	return builder
}

func (builder *snapshotBuilder) Cpu(cpu *Cpu) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Cpu = cpu
	return builder
}

func (builder *snapshotBuilder) CpuProfile(cpuProfile *CpuProfile) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CpuProfile = cpuProfile
	return builder
}

func (builder *snapshotBuilder) CpuShares(cpuShares int64) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CpuShares = &cpuShares
	return builder
}

func (builder *snapshotBuilder) CreationTime(creationTime time.Time) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CreationTime = creationTime
	return builder
}

func (builder *snapshotBuilder) CustomCompatibilityVersion(customCompatibilityVersion *Version) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CustomCompatibilityVersion = customCompatibilityVersion
	return builder
}

func (builder *snapshotBuilder) CustomCpuModel(customCpuModel string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CustomCpuModel = &customCpuModel
	return builder
}

func (builder *snapshotBuilder) CustomEmulatedMachine(customEmulatedMachine string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CustomEmulatedMachine = &customEmulatedMachine
	return builder
}

func (builder *snapshotBuilder) CustomProperties(customProperties []CustomProperty) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.CustomProperties = customProperties
	return builder
}

func (builder *snapshotBuilder) Date(date time.Time) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Date = date
	return builder
}

func (builder *snapshotBuilder) DeleteProtected(deleteProtected bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.DeleteProtected = &deleteProtected
	return builder
}

func (builder *snapshotBuilder) Description(description string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Description = &description
	return builder
}

func (builder *snapshotBuilder) DiskAttachments(diskAttachments []DiskAttachment) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.DiskAttachments = diskAttachments
	return builder
}

func (builder *snapshotBuilder) Display(display *Display) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Display = display
	return builder
}

func (builder *snapshotBuilder) Domain(domain *Domain) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Domain = domain
	return builder
}

func (builder *snapshotBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *snapshotBuilder) Floppies(floppies []Floppy) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Floppies = floppies
	return builder
}

func (builder *snapshotBuilder) Fqdn(fqdn string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Fqdn = &fqdn
	return builder
}

func (builder *snapshotBuilder) GraphicsConsoles(graphicsConsoles []GraphicsConsole) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.GraphicsConsoles = graphicsConsoles
	return builder
}

func (builder *snapshotBuilder) GuestOperatingSystem(guestOperatingSystem *GuestOperatingSystem) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.GuestOperatingSystem = guestOperatingSystem
	return builder
}

func (builder *snapshotBuilder) GuestTimeZone(guestTimeZone *TimeZone) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.GuestTimeZone = guestTimeZone
	return builder
}

func (builder *snapshotBuilder) HighAvailability(highAvailability *HighAvailability) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.HighAvailability = highAvailability
	return builder
}

func (builder *snapshotBuilder) Host(host *Host) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Host = host
	return builder
}

func (builder *snapshotBuilder) HostDevices(hostDevices []HostDevice) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.HostDevices = hostDevices
	return builder
}

func (builder *snapshotBuilder) Id(id string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Id = &id
	return builder
}

func (builder *snapshotBuilder) Initialization(initialization *Initialization) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Initialization = initialization
	return builder
}

func (builder *snapshotBuilder) InstanceType(instanceType *InstanceType) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.InstanceType = instanceType
	return builder
}

func (builder *snapshotBuilder) Io(io *Io) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Io = io
	return builder
}

func (builder *snapshotBuilder) KatelloErrata(katelloErrata []KatelloErratum) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.KatelloErrata = katelloErrata
	return builder
}

func (builder *snapshotBuilder) LargeIcon(largeIcon *Icon) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.LargeIcon = largeIcon
	return builder
}

func (builder *snapshotBuilder) Lease(lease *StorageDomainLease) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Lease = lease
	return builder
}

func (builder *snapshotBuilder) Memory(memory int64) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Memory = &memory
	return builder
}

func (builder *snapshotBuilder) MemoryPolicy(memoryPolicy *MemoryPolicy) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.MemoryPolicy = memoryPolicy
	return builder
}

func (builder *snapshotBuilder) Migration(migration *MigrationOptions) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Migration = migration
	return builder
}

func (builder *snapshotBuilder) MigrationDowntime(migrationDowntime int64) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.MigrationDowntime = &migrationDowntime
	return builder
}

func (builder *snapshotBuilder) Name(name string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Name = &name
	return builder
}

func (builder *snapshotBuilder) NextRunConfigurationExists(nextRunConfigurationExists bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.NextRunConfigurationExists = &nextRunConfigurationExists
	return builder
}

func (builder *snapshotBuilder) Nics(nics []Nic) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Nics = nics
	return builder
}

func (builder *snapshotBuilder) NumaNodes(numaNodes []NumaNode) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.NumaNodes = numaNodes
	return builder
}

func (builder *snapshotBuilder) NumaTuneMode(numaTuneMode NumaTuneMode) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.NumaTuneMode = numaTuneMode
	return builder
}

func (builder *snapshotBuilder) Origin(origin string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Origin = &origin
	return builder
}

func (builder *snapshotBuilder) OriginalTemplate(originalTemplate *Template) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.OriginalTemplate = originalTemplate
	return builder
}

func (builder *snapshotBuilder) Os(os *OperatingSystem) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Os = os
	return builder
}

func (builder *snapshotBuilder) Payloads(payloads []Payload) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Payloads = payloads
	return builder
}

func (builder *snapshotBuilder) Permissions(permissions []Permission) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Permissions = permissions
	return builder
}

func (builder *snapshotBuilder) PersistMemorystate(persistMemorystate bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.PersistMemorystate = &persistMemorystate
	return builder
}

func (builder *snapshotBuilder) PlacementPolicy(placementPolicy *VmPlacementPolicy) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.PlacementPolicy = placementPolicy
	return builder
}

func (builder *snapshotBuilder) Quota(quota *Quota) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Quota = quota
	return builder
}

func (builder *snapshotBuilder) ReportedDevices(reportedDevices []ReportedDevice) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.ReportedDevices = reportedDevices
	return builder
}

func (builder *snapshotBuilder) RngDevice(rngDevice *RngDevice) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.RngDevice = rngDevice
	return builder
}

func (builder *snapshotBuilder) RunOnce(runOnce bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.RunOnce = &runOnce
	return builder
}

func (builder *snapshotBuilder) SerialNumber(serialNumber *SerialNumber) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.SerialNumber = serialNumber
	return builder
}

func (builder *snapshotBuilder) Sessions(sessions []Session) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Sessions = sessions
	return builder
}

func (builder *snapshotBuilder) SmallIcon(smallIcon *Icon) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.SmallIcon = smallIcon
	return builder
}

func (builder *snapshotBuilder) SnapshotStatus(snapshotStatus SnapshotStatus) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.SnapshotStatus = snapshotStatus
	return builder
}

func (builder *snapshotBuilder) SnapshotType(snapshotType SnapshotType) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.SnapshotType = snapshotType
	return builder
}

func (builder *snapshotBuilder) Snapshots(snapshots []Snapshot) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Snapshots = snapshots
	return builder
}

func (builder *snapshotBuilder) SoundcardEnabled(soundcardEnabled bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.SoundcardEnabled = &soundcardEnabled
	return builder
}

func (builder *snapshotBuilder) Sso(sso *Sso) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Sso = sso
	return builder
}

func (builder *snapshotBuilder) StartPaused(startPaused bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.StartPaused = &startPaused
	return builder
}

func (builder *snapshotBuilder) StartTime(startTime time.Time) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.StartTime = startTime
	return builder
}

func (builder *snapshotBuilder) Stateless(stateless bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Stateless = &stateless
	return builder
}

func (builder *snapshotBuilder) Statistics(statistics []Statistic) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Statistics = statistics
	return builder
}

func (builder *snapshotBuilder) Status(status VmStatus) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Status = status
	return builder
}

func (builder *snapshotBuilder) StatusDetail(statusDetail string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.StatusDetail = &statusDetail
	return builder
}

func (builder *snapshotBuilder) StopReason(stopReason string) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.StopReason = &stopReason
	return builder
}

func (builder *snapshotBuilder) StopTime(stopTime time.Time) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.StopTime = stopTime
	return builder
}

func (builder *snapshotBuilder) StorageDomain(storageDomain *StorageDomain) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.StorageDomain = storageDomain
	return builder
}

func (builder *snapshotBuilder) Tags(tags []Tag) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Tags = tags
	return builder
}

func (builder *snapshotBuilder) Template(template *Template) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Template = template
	return builder
}

func (builder *snapshotBuilder) TimeZone(timeZone *TimeZone) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.TimeZone = timeZone
	return builder
}

func (builder *snapshotBuilder) TunnelMigration(tunnelMigration bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.TunnelMigration = &tunnelMigration
	return builder
}

func (builder *snapshotBuilder) Type(type_ VmType) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Type = type_
	return builder
}

func (builder *snapshotBuilder) Usb(usb *Usb) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Usb = usb
	return builder
}

func (builder *snapshotBuilder) UseLatestTemplateVersion(useLatestTemplateVersion bool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.UseLatestTemplateVersion = &useLatestTemplateVersion
	return builder
}

func (builder *snapshotBuilder) VirtioScsi(virtioScsi *VirtioScsi) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.VirtioScsi = virtioScsi
	return builder
}

func (builder *snapshotBuilder) Vm(vm *Vm) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Vm = vm
	return builder
}

func (builder *snapshotBuilder) VmPool(vmPool *VmPool) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.VmPool = vmPool
	return builder
}

func (builder *snapshotBuilder) Watchdogs(watchdogs []Watchdog) *snapshotBuilder {
	if builder.err != nil {
		return builder
	}

	builder.snapshot.Watchdogs = watchdogs
	return builder
}

func (builder *snapshotBuilder) Build() (*Snapshot, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.snapshot, nil
}

type specialObjectsBuilder struct {
	specialObjects *SpecialObjects
	err            error
}

func NewSpecialObjectsBuilder() *specialObjectsBuilder {
	return &specialObjectsBuilder{specialObjects: &SpecialObjects{}, err: nil}
}

func (builder *specialObjectsBuilder) BlankTemplate(blankTemplate *Template) *specialObjectsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.specialObjects.BlankTemplate = blankTemplate
	return builder
}

func (builder *specialObjectsBuilder) RootTag(rootTag *Tag) *specialObjectsBuilder {
	if builder.err != nil {
		return builder
	}

	builder.specialObjects.RootTag = rootTag
	return builder
}

func (builder *specialObjectsBuilder) Build() (*SpecialObjects, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.specialObjects, nil
}

type spmBuilder struct {
	spm *Spm
	err error
}

func NewSpmBuilder() *spmBuilder {
	return &spmBuilder{spm: &Spm{}, err: nil}
}

func (builder *spmBuilder) Priority(priority int64) *spmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.spm.Priority = &priority
	return builder
}

func (builder *spmBuilder) Status(status SpmStatus) *spmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.spm.Status = status
	return builder
}

func (builder *spmBuilder) Build() (*Spm, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.spm, nil
}

type sshBuilder struct {
	ssh *Ssh
	err error
}

func NewSshBuilder() *sshBuilder {
	return &sshBuilder{ssh: &Ssh{}, err: nil}
}

func (builder *sshBuilder) AuthenticationMethod(authenticationMethod SshAuthenticationMethod) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.AuthenticationMethod = authenticationMethod
	return builder
}

func (builder *sshBuilder) Comment(comment string) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.Comment = &comment
	return builder
}

func (builder *sshBuilder) Description(description string) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.Description = &description
	return builder
}

func (builder *sshBuilder) Fingerprint(fingerprint string) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.Fingerprint = &fingerprint
	return builder
}

func (builder *sshBuilder) Id(id string) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.Id = &id
	return builder
}

func (builder *sshBuilder) Name(name string) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.Name = &name
	return builder
}

func (builder *sshBuilder) Port(port int64) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.Port = &port
	return builder
}

func (builder *sshBuilder) User(user *User) *sshBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ssh.User = user
	return builder
}

func (builder *sshBuilder) Build() (*Ssh, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.ssh, nil
}

type sshPublicKeyBuilder struct {
	sshPublicKey *SshPublicKey
	err          error
}

func NewSshPublicKeyBuilder() *sshPublicKeyBuilder {
	return &sshPublicKeyBuilder{sshPublicKey: &SshPublicKey{}, err: nil}
}

func (builder *sshPublicKeyBuilder) Comment(comment string) *sshPublicKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sshPublicKey.Comment = &comment
	return builder
}

func (builder *sshPublicKeyBuilder) Content(content string) *sshPublicKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sshPublicKey.Content = &content
	return builder
}

func (builder *sshPublicKeyBuilder) Description(description string) *sshPublicKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sshPublicKey.Description = &description
	return builder
}

func (builder *sshPublicKeyBuilder) Id(id string) *sshPublicKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sshPublicKey.Id = &id
	return builder
}

func (builder *sshPublicKeyBuilder) Name(name string) *sshPublicKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sshPublicKey.Name = &name
	return builder
}

func (builder *sshPublicKeyBuilder) User(user *User) *sshPublicKeyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sshPublicKey.User = user
	return builder
}

func (builder *sshPublicKeyBuilder) Build() (*SshPublicKey, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.sshPublicKey, nil
}

type ssoBuilder struct {
	sso *Sso
	err error
}

func NewSsoBuilder() *ssoBuilder {
	return &ssoBuilder{sso: &Sso{}, err: nil}
}

func (builder *ssoBuilder) Methods(methods []Method) *ssoBuilder {
	if builder.err != nil {
		return builder
	}

	builder.sso.Methods = methods
	return builder
}

func (builder *ssoBuilder) Build() (*Sso, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.sso, nil
}

type statisticBuilder struct {
	statistic *Statistic
	err       error
}

func NewStatisticBuilder() *statisticBuilder {
	return &statisticBuilder{statistic: &Statistic{}, err: nil}
}

func (builder *statisticBuilder) Brick(brick *GlusterBrick) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Brick = brick
	return builder
}

func (builder *statisticBuilder) Comment(comment string) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Comment = &comment
	return builder
}

func (builder *statisticBuilder) Description(description string) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Description = &description
	return builder
}

func (builder *statisticBuilder) Disk(disk *Disk) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Disk = disk
	return builder
}

func (builder *statisticBuilder) GlusterVolume(glusterVolume *GlusterVolume) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.GlusterVolume = glusterVolume
	return builder
}

func (builder *statisticBuilder) Host(host *Host) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Host = host
	return builder
}

func (builder *statisticBuilder) HostNic(hostNic *HostNic) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.HostNic = hostNic
	return builder
}

func (builder *statisticBuilder) HostNumaNode(hostNumaNode *NumaNode) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.HostNumaNode = hostNumaNode
	return builder
}

func (builder *statisticBuilder) Id(id string) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Id = &id
	return builder
}

func (builder *statisticBuilder) Kind(kind StatisticKind) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Kind = kind
	return builder
}

func (builder *statisticBuilder) Name(name string) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Name = &name
	return builder
}

func (builder *statisticBuilder) Nic(nic *Nic) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Nic = nic
	return builder
}

func (builder *statisticBuilder) Step(step *Step) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Step = step
	return builder
}

func (builder *statisticBuilder) Type(type_ ValueType) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Type = type_
	return builder
}

func (builder *statisticBuilder) Unit(unit StatisticUnit) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Unit = unit
	return builder
}

func (builder *statisticBuilder) Values(values []Value) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Values = values
	return builder
}

func (builder *statisticBuilder) Vm(vm *Vm) *statisticBuilder {
	if builder.err != nil {
		return builder
	}

	builder.statistic.Vm = vm
	return builder
}

func (builder *statisticBuilder) Build() (*Statistic, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.statistic, nil
}

type stepBuilder struct {
	step *Step
	err  error
}

func NewStepBuilder() *stepBuilder {
	return &stepBuilder{step: &Step{}, err: nil}
}

func (builder *stepBuilder) Comment(comment string) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Comment = &comment
	return builder
}

func (builder *stepBuilder) Description(description string) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Description = &description
	return builder
}

func (builder *stepBuilder) EndTime(endTime time.Time) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.EndTime = endTime
	return builder
}

func (builder *stepBuilder) ExecutionHost(executionHost *Host) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.ExecutionHost = executionHost
	return builder
}

func (builder *stepBuilder) External(external bool) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.External = &external
	return builder
}

func (builder *stepBuilder) ExternalType(externalType ExternalSystemType) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.ExternalType = externalType
	return builder
}

func (builder *stepBuilder) Id(id string) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Id = &id
	return builder
}

func (builder *stepBuilder) Job(job *Job) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Job = job
	return builder
}

func (builder *stepBuilder) Name(name string) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Name = &name
	return builder
}

func (builder *stepBuilder) Number(number int64) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Number = &number
	return builder
}

func (builder *stepBuilder) ParentStep(parentStep *Step) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.ParentStep = parentStep
	return builder
}

func (builder *stepBuilder) Progress(progress int64) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Progress = &progress
	return builder
}

func (builder *stepBuilder) StartTime(startTime time.Time) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.StartTime = startTime
	return builder
}

func (builder *stepBuilder) Statistics(statistics []Statistic) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Statistics = statistics
	return builder
}

func (builder *stepBuilder) Status(status StepStatus) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Status = status
	return builder
}

func (builder *stepBuilder) Type(type_ StepEnum) *stepBuilder {
	if builder.err != nil {
		return builder
	}

	builder.step.Type = type_
	return builder
}

func (builder *stepBuilder) Build() (*Step, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.step, nil
}

type storageConnectionBuilder struct {
	storageConnection *StorageConnection
	err               error
}

func NewStorageConnectionBuilder() *storageConnectionBuilder {
	return &storageConnectionBuilder{storageConnection: &StorageConnection{}, err: nil}
}

func (builder *storageConnectionBuilder) Address(address string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Address = &address
	return builder
}

func (builder *storageConnectionBuilder) Comment(comment string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Comment = &comment
	return builder
}

func (builder *storageConnectionBuilder) Description(description string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Description = &description
	return builder
}

func (builder *storageConnectionBuilder) Host(host *Host) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Host = host
	return builder
}

func (builder *storageConnectionBuilder) Id(id string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Id = &id
	return builder
}

func (builder *storageConnectionBuilder) MountOptions(mountOptions string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.MountOptions = &mountOptions
	return builder
}

func (builder *storageConnectionBuilder) Name(name string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Name = &name
	return builder
}

func (builder *storageConnectionBuilder) NfsRetrans(nfsRetrans int64) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.NfsRetrans = &nfsRetrans
	return builder
}

func (builder *storageConnectionBuilder) NfsTimeo(nfsTimeo int64) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.NfsTimeo = &nfsTimeo
	return builder
}

func (builder *storageConnectionBuilder) NfsVersion(nfsVersion NfsVersion) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.NfsVersion = nfsVersion
	return builder
}

func (builder *storageConnectionBuilder) Password(password string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Password = &password
	return builder
}

func (builder *storageConnectionBuilder) Path(path string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Path = &path
	return builder
}

func (builder *storageConnectionBuilder) Port(port int64) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Port = &port
	return builder
}

func (builder *storageConnectionBuilder) Portal(portal string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Portal = &portal
	return builder
}

func (builder *storageConnectionBuilder) Target(target string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Target = &target
	return builder
}

func (builder *storageConnectionBuilder) Type(type_ StorageType) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Type = type_
	return builder
}

func (builder *storageConnectionBuilder) Username(username string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.Username = &username
	return builder
}

func (builder *storageConnectionBuilder) VfsType(vfsType string) *storageConnectionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnection.VfsType = &vfsType
	return builder
}

func (builder *storageConnectionBuilder) Build() (*StorageConnection, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.storageConnection, nil
}

type storageConnectionExtensionBuilder struct {
	storageConnectionExtension *StorageConnectionExtension
	err                        error
}

func NewStorageConnectionExtensionBuilder() *storageConnectionExtensionBuilder {
	return &storageConnectionExtensionBuilder{storageConnectionExtension: &StorageConnectionExtension{}, err: nil}
}

func (builder *storageConnectionExtensionBuilder) Comment(comment string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Comment = &comment
	return builder
}

func (builder *storageConnectionExtensionBuilder) Description(description string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Description = &description
	return builder
}

func (builder *storageConnectionExtensionBuilder) Host(host *Host) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Host = host
	return builder
}

func (builder *storageConnectionExtensionBuilder) Id(id string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Id = &id
	return builder
}

func (builder *storageConnectionExtensionBuilder) Name(name string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Name = &name
	return builder
}

func (builder *storageConnectionExtensionBuilder) Password(password string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Password = &password
	return builder
}

func (builder *storageConnectionExtensionBuilder) Target(target string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Target = &target
	return builder
}

func (builder *storageConnectionExtensionBuilder) Username(username string) *storageConnectionExtensionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageConnectionExtension.Username = &username
	return builder
}

func (builder *storageConnectionExtensionBuilder) Build() (*StorageConnectionExtension, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.storageConnectionExtension, nil
}

type storageDomainBuilder struct {
	storageDomain *StorageDomain
	err           error
}

func NewStorageDomainBuilder() *storageDomainBuilder {
	return &storageDomainBuilder{storageDomain: &StorageDomain{}, err: nil}
}

func (builder *storageDomainBuilder) Available(available int64) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Available = &available
	return builder
}

func (builder *storageDomainBuilder) Comment(comment string) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Comment = &comment
	return builder
}

func (builder *storageDomainBuilder) Committed(committed int64) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Committed = &committed
	return builder
}

func (builder *storageDomainBuilder) CriticalSpaceActionBlocker(criticalSpaceActionBlocker int64) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.CriticalSpaceActionBlocker = &criticalSpaceActionBlocker
	return builder
}

func (builder *storageDomainBuilder) DataCenter(dataCenter *DataCenter) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.DataCenter = dataCenter
	return builder
}

func (builder *storageDomainBuilder) DataCenters(dataCenters []DataCenter) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.DataCenters = dataCenters
	return builder
}

func (builder *storageDomainBuilder) Description(description string) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Description = &description
	return builder
}

func (builder *storageDomainBuilder) DiscardAfterDelete(discardAfterDelete bool) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.DiscardAfterDelete = &discardAfterDelete
	return builder
}

func (builder *storageDomainBuilder) DiskProfiles(diskProfiles []DiskProfile) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.DiskProfiles = diskProfiles
	return builder
}

func (builder *storageDomainBuilder) DiskSnapshots(diskSnapshots []DiskSnapshot) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.DiskSnapshots = diskSnapshots
	return builder
}

func (builder *storageDomainBuilder) Disks(disks []Disk) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Disks = disks
	return builder
}

func (builder *storageDomainBuilder) ExternalStatus(externalStatus ExternalStatus) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.ExternalStatus = externalStatus
	return builder
}

func (builder *storageDomainBuilder) Files(files []File) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Files = files
	return builder
}

func (builder *storageDomainBuilder) Host(host *Host) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Host = host
	return builder
}

func (builder *storageDomainBuilder) Id(id string) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Id = &id
	return builder
}

func (builder *storageDomainBuilder) Images(images []Image) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Images = images
	return builder
}

func (builder *storageDomainBuilder) Import(import_ bool) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Import = &import_
	return builder
}

func (builder *storageDomainBuilder) Master(master bool) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Master = &master
	return builder
}

func (builder *storageDomainBuilder) Name(name string) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Name = &name
	return builder
}

func (builder *storageDomainBuilder) Permissions(permissions []Permission) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Permissions = permissions
	return builder
}

func (builder *storageDomainBuilder) Status(status StorageDomainStatus) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Status = status
	return builder
}

func (builder *storageDomainBuilder) Storage(storage *HostStorage) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Storage = storage
	return builder
}

func (builder *storageDomainBuilder) StorageConnections(storageConnections []StorageConnection) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.StorageConnections = storageConnections
	return builder
}

func (builder *storageDomainBuilder) StorageFormat(storageFormat StorageFormat) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.StorageFormat = storageFormat
	return builder
}

func (builder *storageDomainBuilder) SupportsDiscard(supportsDiscard bool) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.SupportsDiscard = &supportsDiscard
	return builder
}

func (builder *storageDomainBuilder) SupportsDiscardZeroesData(supportsDiscardZeroesData bool) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.SupportsDiscardZeroesData = &supportsDiscardZeroesData
	return builder
}

func (builder *storageDomainBuilder) Templates(templates []Template) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Templates = templates
	return builder
}

func (builder *storageDomainBuilder) Type(type_ StorageDomainType) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Type = type_
	return builder
}

func (builder *storageDomainBuilder) Used(used int64) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Used = &used
	return builder
}

func (builder *storageDomainBuilder) Vms(vms []Vm) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.Vms = vms
	return builder
}

func (builder *storageDomainBuilder) WarningLowSpaceIndicator(warningLowSpaceIndicator int64) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.WarningLowSpaceIndicator = &warningLowSpaceIndicator
	return builder
}

func (builder *storageDomainBuilder) WipeAfterDelete(wipeAfterDelete bool) *storageDomainBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomain.WipeAfterDelete = &wipeAfterDelete
	return builder
}

func (builder *storageDomainBuilder) Build() (*StorageDomain, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.storageDomain, nil
}

type storageDomainLeaseBuilder struct {
	storageDomainLease *StorageDomainLease
	err                error
}

func NewStorageDomainLeaseBuilder() *storageDomainLeaseBuilder {
	return &storageDomainLeaseBuilder{storageDomainLease: &StorageDomainLease{}, err: nil}
}

func (builder *storageDomainLeaseBuilder) StorageDomain(storageDomain *StorageDomain) *storageDomainLeaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.storageDomainLease.StorageDomain = storageDomain
	return builder
}

func (builder *storageDomainLeaseBuilder) Build() (*StorageDomainLease, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.storageDomainLease, nil
}

type tagBuilder struct {
	tag *Tag
	err error
}

func NewTagBuilder() *tagBuilder {
	return &tagBuilder{tag: &Tag{}, err: nil}
}

func (builder *tagBuilder) Comment(comment string) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Comment = &comment
	return builder
}

func (builder *tagBuilder) Description(description string) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Description = &description
	return builder
}

func (builder *tagBuilder) Group(group *Group) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Group = group
	return builder
}

func (builder *tagBuilder) Host(host *Host) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Host = host
	return builder
}

func (builder *tagBuilder) Id(id string) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Id = &id
	return builder
}

func (builder *tagBuilder) Name(name string) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Name = &name
	return builder
}

func (builder *tagBuilder) Parent(parent *Tag) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Parent = parent
	return builder
}

func (builder *tagBuilder) Template(template *Template) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Template = template
	return builder
}

func (builder *tagBuilder) User(user *User) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.User = user
	return builder
}

func (builder *tagBuilder) Vm(vm *Vm) *tagBuilder {
	if builder.err != nil {
		return builder
	}

	builder.tag.Vm = vm
	return builder
}

func (builder *tagBuilder) Build() (*Tag, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.tag, nil
}

type templateBuilder struct {
	template *Template
	err      error
}

func NewTemplateBuilder() *templateBuilder {
	return &templateBuilder{template: &Template{}, err: nil}
}

func (builder *templateBuilder) Bios(bios *Bios) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Bios = bios
	return builder
}

func (builder *templateBuilder) Cdroms(cdroms []Cdrom) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Cdroms = cdroms
	return builder
}

func (builder *templateBuilder) Cluster(cluster *Cluster) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Cluster = cluster
	return builder
}

func (builder *templateBuilder) Comment(comment string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Comment = &comment
	return builder
}

func (builder *templateBuilder) Console(console *Console) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Console = console
	return builder
}

func (builder *templateBuilder) Cpu(cpu *Cpu) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Cpu = cpu
	return builder
}

func (builder *templateBuilder) CpuProfile(cpuProfile *CpuProfile) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CpuProfile = cpuProfile
	return builder
}

func (builder *templateBuilder) CpuShares(cpuShares int64) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CpuShares = &cpuShares
	return builder
}

func (builder *templateBuilder) CreationTime(creationTime time.Time) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CreationTime = creationTime
	return builder
}

func (builder *templateBuilder) CustomCompatibilityVersion(customCompatibilityVersion *Version) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CustomCompatibilityVersion = customCompatibilityVersion
	return builder
}

func (builder *templateBuilder) CustomCpuModel(customCpuModel string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CustomCpuModel = &customCpuModel
	return builder
}

func (builder *templateBuilder) CustomEmulatedMachine(customEmulatedMachine string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CustomEmulatedMachine = &customEmulatedMachine
	return builder
}

func (builder *templateBuilder) CustomProperties(customProperties []CustomProperty) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.CustomProperties = customProperties
	return builder
}

func (builder *templateBuilder) DeleteProtected(deleteProtected bool) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.DeleteProtected = &deleteProtected
	return builder
}

func (builder *templateBuilder) Description(description string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Description = &description
	return builder
}

func (builder *templateBuilder) DiskAttachments(diskAttachments []DiskAttachment) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.DiskAttachments = diskAttachments
	return builder
}

func (builder *templateBuilder) Display(display *Display) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Display = display
	return builder
}

func (builder *templateBuilder) Domain(domain *Domain) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Domain = domain
	return builder
}

func (builder *templateBuilder) GraphicsConsoles(graphicsConsoles []GraphicsConsole) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.GraphicsConsoles = graphicsConsoles
	return builder
}

func (builder *templateBuilder) HighAvailability(highAvailability *HighAvailability) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.HighAvailability = highAvailability
	return builder
}

func (builder *templateBuilder) Id(id string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Id = &id
	return builder
}

func (builder *templateBuilder) Initialization(initialization *Initialization) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Initialization = initialization
	return builder
}

func (builder *templateBuilder) Io(io *Io) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Io = io
	return builder
}

func (builder *templateBuilder) LargeIcon(largeIcon *Icon) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.LargeIcon = largeIcon
	return builder
}

func (builder *templateBuilder) Lease(lease *StorageDomainLease) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Lease = lease
	return builder
}

func (builder *templateBuilder) Memory(memory int64) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Memory = &memory
	return builder
}

func (builder *templateBuilder) MemoryPolicy(memoryPolicy *MemoryPolicy) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.MemoryPolicy = memoryPolicy
	return builder
}

func (builder *templateBuilder) Migration(migration *MigrationOptions) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Migration = migration
	return builder
}

func (builder *templateBuilder) MigrationDowntime(migrationDowntime int64) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.MigrationDowntime = &migrationDowntime
	return builder
}

func (builder *templateBuilder) Name(name string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Name = &name
	return builder
}

func (builder *templateBuilder) Nics(nics []Nic) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Nics = nics
	return builder
}

func (builder *templateBuilder) Origin(origin string) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Origin = &origin
	return builder
}

func (builder *templateBuilder) Os(os *OperatingSystem) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Os = os
	return builder
}

func (builder *templateBuilder) Permissions(permissions []Permission) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Permissions = permissions
	return builder
}

func (builder *templateBuilder) Quota(quota *Quota) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Quota = quota
	return builder
}

func (builder *templateBuilder) RngDevice(rngDevice *RngDevice) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.RngDevice = rngDevice
	return builder
}

func (builder *templateBuilder) SerialNumber(serialNumber *SerialNumber) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.SerialNumber = serialNumber
	return builder
}

func (builder *templateBuilder) SmallIcon(smallIcon *Icon) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.SmallIcon = smallIcon
	return builder
}

func (builder *templateBuilder) SoundcardEnabled(soundcardEnabled bool) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.SoundcardEnabled = &soundcardEnabled
	return builder
}

func (builder *templateBuilder) Sso(sso *Sso) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Sso = sso
	return builder
}

func (builder *templateBuilder) StartPaused(startPaused bool) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.StartPaused = &startPaused
	return builder
}

func (builder *templateBuilder) Stateless(stateless bool) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Stateless = &stateless
	return builder
}

func (builder *templateBuilder) Status(status TemplateStatus) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Status = status
	return builder
}

func (builder *templateBuilder) StorageDomain(storageDomain *StorageDomain) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.StorageDomain = storageDomain
	return builder
}

func (builder *templateBuilder) Tags(tags []Tag) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Tags = tags
	return builder
}

func (builder *templateBuilder) TimeZone(timeZone *TimeZone) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.TimeZone = timeZone
	return builder
}

func (builder *templateBuilder) TunnelMigration(tunnelMigration bool) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.TunnelMigration = &tunnelMigration
	return builder
}

func (builder *templateBuilder) Type(type_ VmType) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Type = type_
	return builder
}

func (builder *templateBuilder) Usb(usb *Usb) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Usb = usb
	return builder
}

func (builder *templateBuilder) Version(version *TemplateVersion) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Version = version
	return builder
}

func (builder *templateBuilder) VirtioScsi(virtioScsi *VirtioScsi) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.VirtioScsi = virtioScsi
	return builder
}

func (builder *templateBuilder) Vm(vm *Vm) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Vm = vm
	return builder
}

func (builder *templateBuilder) Watchdogs(watchdogs []Watchdog) *templateBuilder {
	if builder.err != nil {
		return builder
	}

	builder.template.Watchdogs = watchdogs
	return builder
}

func (builder *templateBuilder) Build() (*Template, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.template, nil
}

type templateVersionBuilder struct {
	templateVersion *TemplateVersion
	err             error
}

func NewTemplateVersionBuilder() *templateVersionBuilder {
	return &templateVersionBuilder{templateVersion: &TemplateVersion{}, err: nil}
}

func (builder *templateVersionBuilder) BaseTemplate(baseTemplate *Template) *templateVersionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.templateVersion.BaseTemplate = baseTemplate
	return builder
}

func (builder *templateVersionBuilder) VersionName(versionName string) *templateVersionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.templateVersion.VersionName = &versionName
	return builder
}

func (builder *templateVersionBuilder) VersionNumber(versionNumber int64) *templateVersionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.templateVersion.VersionNumber = &versionNumber
	return builder
}

func (builder *templateVersionBuilder) Build() (*TemplateVersion, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.templateVersion, nil
}

type ticketBuilder struct {
	ticket *Ticket
	err    error
}

func NewTicketBuilder() *ticketBuilder {
	return &ticketBuilder{ticket: &Ticket{}, err: nil}
}

func (builder *ticketBuilder) Expiry(expiry int64) *ticketBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ticket.Expiry = &expiry
	return builder
}

func (builder *ticketBuilder) Value(value string) *ticketBuilder {
	if builder.err != nil {
		return builder
	}

	builder.ticket.Value = &value
	return builder
}

func (builder *ticketBuilder) Build() (*Ticket, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.ticket, nil
}

type timeZoneBuilder struct {
	timeZone *TimeZone
	err      error
}

func NewTimeZoneBuilder() *timeZoneBuilder {
	return &timeZoneBuilder{timeZone: &TimeZone{}, err: nil}
}

func (builder *timeZoneBuilder) Name(name string) *timeZoneBuilder {
	if builder.err != nil {
		return builder
	}

	builder.timeZone.Name = &name
	return builder
}

func (builder *timeZoneBuilder) UtcOffset(utcOffset string) *timeZoneBuilder {
	if builder.err != nil {
		return builder
	}

	builder.timeZone.UtcOffset = &utcOffset
	return builder
}

func (builder *timeZoneBuilder) Build() (*TimeZone, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.timeZone, nil
}

type transparentHugePagesBuilder struct {
	transparentHugePages *TransparentHugePages
	err                  error
}

func NewTransparentHugePagesBuilder() *transparentHugePagesBuilder {
	return &transparentHugePagesBuilder{transparentHugePages: &TransparentHugePages{}, err: nil}
}

func (builder *transparentHugePagesBuilder) Enabled(enabled bool) *transparentHugePagesBuilder {
	if builder.err != nil {
		return builder
	}

	builder.transparentHugePages.Enabled = &enabled
	return builder
}

func (builder *transparentHugePagesBuilder) Build() (*TransparentHugePages, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.transparentHugePages, nil
}

type unmanagedNetworkBuilder struct {
	unmanagedNetwork *UnmanagedNetwork
	err              error
}

func NewUnmanagedNetworkBuilder() *unmanagedNetworkBuilder {
	return &unmanagedNetworkBuilder{unmanagedNetwork: &UnmanagedNetwork{}, err: nil}
}

func (builder *unmanagedNetworkBuilder) Comment(comment string) *unmanagedNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.unmanagedNetwork.Comment = &comment
	return builder
}

func (builder *unmanagedNetworkBuilder) Description(description string) *unmanagedNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.unmanagedNetwork.Description = &description
	return builder
}

func (builder *unmanagedNetworkBuilder) Host(host *Host) *unmanagedNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.unmanagedNetwork.Host = host
	return builder
}

func (builder *unmanagedNetworkBuilder) HostNic(hostNic *HostNic) *unmanagedNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.unmanagedNetwork.HostNic = hostNic
	return builder
}

func (builder *unmanagedNetworkBuilder) Id(id string) *unmanagedNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.unmanagedNetwork.Id = &id
	return builder
}

func (builder *unmanagedNetworkBuilder) Name(name string) *unmanagedNetworkBuilder {
	if builder.err != nil {
		return builder
	}

	builder.unmanagedNetwork.Name = &name
	return builder
}

func (builder *unmanagedNetworkBuilder) Build() (*UnmanagedNetwork, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.unmanagedNetwork, nil
}

type usbBuilder struct {
	usb *Usb
	err error
}

func NewUsbBuilder() *usbBuilder {
	return &usbBuilder{usb: &Usb{}, err: nil}
}

func (builder *usbBuilder) Enabled(enabled bool) *usbBuilder {
	if builder.err != nil {
		return builder
	}

	builder.usb.Enabled = &enabled
	return builder
}

func (builder *usbBuilder) Type(type_ UsbType) *usbBuilder {
	if builder.err != nil {
		return builder
	}

	builder.usb.Type = type_
	return builder
}

func (builder *usbBuilder) Build() (*Usb, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.usb, nil
}

type userBuilder struct {
	user *User
	err  error
}

func NewUserBuilder() *userBuilder {
	return &userBuilder{user: &User{}, err: nil}
}

func (builder *userBuilder) Comment(comment string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Comment = &comment
	return builder
}

func (builder *userBuilder) Department(department string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Department = &department
	return builder
}

func (builder *userBuilder) Description(description string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Description = &description
	return builder
}

func (builder *userBuilder) Domain(domain *Domain) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Domain = domain
	return builder
}

func (builder *userBuilder) DomainEntryId(domainEntryId string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.DomainEntryId = &domainEntryId
	return builder
}

func (builder *userBuilder) Email(email string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Email = &email
	return builder
}

func (builder *userBuilder) Groups(groups []Group) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Groups = groups
	return builder
}

func (builder *userBuilder) Id(id string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Id = &id
	return builder
}

func (builder *userBuilder) LastName(lastName string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.LastName = &lastName
	return builder
}

func (builder *userBuilder) LoggedIn(loggedIn bool) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.LoggedIn = &loggedIn
	return builder
}

func (builder *userBuilder) Name(name string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Name = &name
	return builder
}

func (builder *userBuilder) Namespace(namespace string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Namespace = &namespace
	return builder
}

func (builder *userBuilder) Password(password string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Password = &password
	return builder
}

func (builder *userBuilder) Permissions(permissions []Permission) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Permissions = permissions
	return builder
}

func (builder *userBuilder) Principal(principal string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Principal = &principal
	return builder
}

func (builder *userBuilder) Roles(roles []Role) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Roles = roles
	return builder
}

func (builder *userBuilder) SshPublicKeys(sshPublicKeys []SshPublicKey) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.SshPublicKeys = sshPublicKeys
	return builder
}

func (builder *userBuilder) Tags(tags []Tag) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.Tags = tags
	return builder
}

func (builder *userBuilder) UserName(userName string) *userBuilder {
	if builder.err != nil {
		return builder
	}

	builder.user.UserName = &userName
	return builder
}

func (builder *userBuilder) Build() (*User, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.user, nil
}

type valueBuilder struct {
	value *Value
	err   error
}

func NewValueBuilder() *valueBuilder {
	return &valueBuilder{value: &Value{}, err: nil}
}

func (builder *valueBuilder) Datum(datum float64) *valueBuilder {
	if builder.err != nil {
		return builder
	}

	builder.value.Datum = &datum
	return builder
}

func (builder *valueBuilder) Detail(detail string) *valueBuilder {
	if builder.err != nil {
		return builder
	}

	builder.value.Detail = &detail
	return builder
}

func (builder *valueBuilder) Build() (*Value, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.value, nil
}

type vcpuPinBuilder struct {
	vcpuPin *VcpuPin
	err     error
}

func NewVcpuPinBuilder() *vcpuPinBuilder {
	return &vcpuPinBuilder{vcpuPin: &VcpuPin{}, err: nil}
}

func (builder *vcpuPinBuilder) CpuSet(cpuSet string) *vcpuPinBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vcpuPin.CpuSet = &cpuSet
	return builder
}

func (builder *vcpuPinBuilder) Vcpu(vcpu int64) *vcpuPinBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vcpuPin.Vcpu = &vcpu
	return builder
}

func (builder *vcpuPinBuilder) Build() (*VcpuPin, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vcpuPin, nil
}

type vendorBuilder struct {
	vendor *Vendor
	err    error
}

func NewVendorBuilder() *vendorBuilder {
	return &vendorBuilder{vendor: &Vendor{}, err: nil}
}

func (builder *vendorBuilder) Comment(comment string) *vendorBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vendor.Comment = &comment
	return builder
}

func (builder *vendorBuilder) Description(description string) *vendorBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vendor.Description = &description
	return builder
}

func (builder *vendorBuilder) Id(id string) *vendorBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vendor.Id = &id
	return builder
}

func (builder *vendorBuilder) Name(name string) *vendorBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vendor.Name = &name
	return builder
}

func (builder *vendorBuilder) Build() (*Vendor, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vendor, nil
}

type versionBuilder struct {
	version *Version
	err     error
}

func NewVersionBuilder() *versionBuilder {
	return &versionBuilder{version: &Version{}, err: nil}
}

func (builder *versionBuilder) Build_(build_ int64) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Build_ = &build_
	return builder
}

func (builder *versionBuilder) Comment(comment string) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Comment = &comment
	return builder
}

func (builder *versionBuilder) Description(description string) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Description = &description
	return builder
}

func (builder *versionBuilder) FullVersion(fullVersion string) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.FullVersion = &fullVersion
	return builder
}

func (builder *versionBuilder) Id(id string) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Id = &id
	return builder
}

func (builder *versionBuilder) Major(major int64) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Major = &major
	return builder
}

func (builder *versionBuilder) Minor(minor int64) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Minor = &minor
	return builder
}

func (builder *versionBuilder) Name(name string) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Name = &name
	return builder
}

func (builder *versionBuilder) Revision(revision int64) *versionBuilder {
	if builder.err != nil {
		return builder
	}

	builder.version.Revision = &revision
	return builder
}

func (builder *versionBuilder) Build() (*Version, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.version, nil
}

type virtioScsiBuilder struct {
	virtioScsi *VirtioScsi
	err        error
}

func NewVirtioScsiBuilder() *virtioScsiBuilder {
	return &virtioScsiBuilder{virtioScsi: &VirtioScsi{}, err: nil}
}

func (builder *virtioScsiBuilder) Enabled(enabled bool) *virtioScsiBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtioScsi.Enabled = &enabled
	return builder
}

func (builder *virtioScsiBuilder) Build() (*VirtioScsi, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.virtioScsi, nil
}

type virtualNumaNodeBuilder struct {
	virtualNumaNode *VirtualNumaNode
	err             error
}

func NewVirtualNumaNodeBuilder() *virtualNumaNodeBuilder {
	return &virtualNumaNodeBuilder{virtualNumaNode: &VirtualNumaNode{}, err: nil}
}

func (builder *virtualNumaNodeBuilder) Comment(comment string) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Comment = &comment
	return builder
}

func (builder *virtualNumaNodeBuilder) Cpu(cpu *Cpu) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Cpu = cpu
	return builder
}

func (builder *virtualNumaNodeBuilder) Description(description string) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Description = &description
	return builder
}

func (builder *virtualNumaNodeBuilder) Host(host *Host) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Host = host
	return builder
}

func (builder *virtualNumaNodeBuilder) Id(id string) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Id = &id
	return builder
}

func (builder *virtualNumaNodeBuilder) Index(index int64) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Index = &index
	return builder
}

func (builder *virtualNumaNodeBuilder) Memory(memory int64) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Memory = &memory
	return builder
}

func (builder *virtualNumaNodeBuilder) Name(name string) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Name = &name
	return builder
}

func (builder *virtualNumaNodeBuilder) NodeDistance(nodeDistance string) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.NodeDistance = &nodeDistance
	return builder
}

func (builder *virtualNumaNodeBuilder) NumaNodePins(numaNodePins []NumaNodePin) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.NumaNodePins = numaNodePins
	return builder
}

func (builder *virtualNumaNodeBuilder) Statistics(statistics []Statistic) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Statistics = statistics
	return builder
}

func (builder *virtualNumaNodeBuilder) Vm(vm *Vm) *virtualNumaNodeBuilder {
	if builder.err != nil {
		return builder
	}

	builder.virtualNumaNode.Vm = vm
	return builder
}

func (builder *virtualNumaNodeBuilder) Build() (*VirtualNumaNode, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.virtualNumaNode, nil
}

type vlanBuilder struct {
	vlan *Vlan
	err  error
}

func NewVlanBuilder() *vlanBuilder {
	return &vlanBuilder{vlan: &Vlan{}, err: nil}
}

func (builder *vlanBuilder) Id(id int64) *vlanBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vlan.Id = &id
	return builder
}

func (builder *vlanBuilder) Build() (*Vlan, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vlan, nil
}

type vmBuilder struct {
	vm  *Vm
	err error
}

func NewVmBuilder() *vmBuilder {
	return &vmBuilder{vm: &Vm{}, err: nil}
}

func (builder *vmBuilder) AffinityLabels(affinityLabels []AffinityLabel) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.AffinityLabels = affinityLabels
	return builder
}

func (builder *vmBuilder) Applications(applications []Application) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Applications = applications
	return builder
}

func (builder *vmBuilder) Bios(bios *Bios) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Bios = bios
	return builder
}

func (builder *vmBuilder) Cdroms(cdroms []Cdrom) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Cdroms = cdroms
	return builder
}

func (builder *vmBuilder) Cluster(cluster *Cluster) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Cluster = cluster
	return builder
}

func (builder *vmBuilder) Comment(comment string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Comment = &comment
	return builder
}

func (builder *vmBuilder) Console(console *Console) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Console = console
	return builder
}

func (builder *vmBuilder) Cpu(cpu *Cpu) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Cpu = cpu
	return builder
}

func (builder *vmBuilder) CpuProfile(cpuProfile *CpuProfile) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CpuProfile = cpuProfile
	return builder
}

func (builder *vmBuilder) CpuShares(cpuShares int64) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CpuShares = &cpuShares
	return builder
}

func (builder *vmBuilder) CreationTime(creationTime time.Time) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CreationTime = creationTime
	return builder
}

func (builder *vmBuilder) CustomCompatibilityVersion(customCompatibilityVersion *Version) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CustomCompatibilityVersion = customCompatibilityVersion
	return builder
}

func (builder *vmBuilder) CustomCpuModel(customCpuModel string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CustomCpuModel = &customCpuModel
	return builder
}

func (builder *vmBuilder) CustomEmulatedMachine(customEmulatedMachine string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CustomEmulatedMachine = &customEmulatedMachine
	return builder
}

func (builder *vmBuilder) CustomProperties(customProperties []CustomProperty) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.CustomProperties = customProperties
	return builder
}

func (builder *vmBuilder) DeleteProtected(deleteProtected bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.DeleteProtected = &deleteProtected
	return builder
}

func (builder *vmBuilder) Description(description string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Description = &description
	return builder
}

func (builder *vmBuilder) DiskAttachments(diskAttachments []DiskAttachment) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.DiskAttachments = diskAttachments
	return builder
}

func (builder *vmBuilder) Display(display *Display) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Display = display
	return builder
}

func (builder *vmBuilder) Domain(domain *Domain) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Domain = domain
	return builder
}

func (builder *vmBuilder) ExternalHostProvider(externalHostProvider *ExternalHostProvider) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.ExternalHostProvider = externalHostProvider
	return builder
}

func (builder *vmBuilder) Floppies(floppies []Floppy) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Floppies = floppies
	return builder
}

func (builder *vmBuilder) Fqdn(fqdn string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Fqdn = &fqdn
	return builder
}

func (builder *vmBuilder) GraphicsConsoles(graphicsConsoles []GraphicsConsole) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.GraphicsConsoles = graphicsConsoles
	return builder
}

func (builder *vmBuilder) GuestOperatingSystem(guestOperatingSystem *GuestOperatingSystem) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.GuestOperatingSystem = guestOperatingSystem
	return builder
}

func (builder *vmBuilder) GuestTimeZone(guestTimeZone *TimeZone) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.GuestTimeZone = guestTimeZone
	return builder
}

func (builder *vmBuilder) HighAvailability(highAvailability *HighAvailability) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.HighAvailability = highAvailability
	return builder
}

func (builder *vmBuilder) Host(host *Host) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Host = host
	return builder
}

func (builder *vmBuilder) HostDevices(hostDevices []HostDevice) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.HostDevices = hostDevices
	return builder
}

func (builder *vmBuilder) Id(id string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Id = &id
	return builder
}

func (builder *vmBuilder) Initialization(initialization *Initialization) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Initialization = initialization
	return builder
}

func (builder *vmBuilder) InstanceType(instanceType *InstanceType) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.InstanceType = instanceType
	return builder
}

func (builder *vmBuilder) Io(io *Io) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Io = io
	return builder
}

func (builder *vmBuilder) KatelloErrata(katelloErrata []KatelloErratum) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.KatelloErrata = katelloErrata
	return builder
}

func (builder *vmBuilder) LargeIcon(largeIcon *Icon) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.LargeIcon = largeIcon
	return builder
}

func (builder *vmBuilder) Lease(lease *StorageDomainLease) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Lease = lease
	return builder
}

func (builder *vmBuilder) Memory(memory int64) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Memory = &memory
	return builder
}

func (builder *vmBuilder) MemoryPolicy(memoryPolicy *MemoryPolicy) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.MemoryPolicy = memoryPolicy
	return builder
}

func (builder *vmBuilder) Migration(migration *MigrationOptions) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Migration = migration
	return builder
}

func (builder *vmBuilder) MigrationDowntime(migrationDowntime int64) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.MigrationDowntime = &migrationDowntime
	return builder
}

func (builder *vmBuilder) Name(name string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Name = &name
	return builder
}

func (builder *vmBuilder) NextRunConfigurationExists(nextRunConfigurationExists bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.NextRunConfigurationExists = &nextRunConfigurationExists
	return builder
}

func (builder *vmBuilder) Nics(nics []Nic) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Nics = nics
	return builder
}

func (builder *vmBuilder) NumaNodes(numaNodes []NumaNode) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.NumaNodes = numaNodes
	return builder
}

func (builder *vmBuilder) NumaTuneMode(numaTuneMode NumaTuneMode) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.NumaTuneMode = numaTuneMode
	return builder
}

func (builder *vmBuilder) Origin(origin string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Origin = &origin
	return builder
}

func (builder *vmBuilder) OriginalTemplate(originalTemplate *Template) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.OriginalTemplate = originalTemplate
	return builder
}

func (builder *vmBuilder) Os(os *OperatingSystem) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Os = os
	return builder
}

func (builder *vmBuilder) Payloads(payloads []Payload) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Payloads = payloads
	return builder
}

func (builder *vmBuilder) Permissions(permissions []Permission) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Permissions = permissions
	return builder
}

func (builder *vmBuilder) PlacementPolicy(placementPolicy *VmPlacementPolicy) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.PlacementPolicy = placementPolicy
	return builder
}

func (builder *vmBuilder) Quota(quota *Quota) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Quota = quota
	return builder
}

func (builder *vmBuilder) ReportedDevices(reportedDevices []ReportedDevice) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.ReportedDevices = reportedDevices
	return builder
}

func (builder *vmBuilder) RngDevice(rngDevice *RngDevice) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.RngDevice = rngDevice
	return builder
}

func (builder *vmBuilder) RunOnce(runOnce bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.RunOnce = &runOnce
	return builder
}

func (builder *vmBuilder) SerialNumber(serialNumber *SerialNumber) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.SerialNumber = serialNumber
	return builder
}

func (builder *vmBuilder) Sessions(sessions []Session) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Sessions = sessions
	return builder
}

func (builder *vmBuilder) SmallIcon(smallIcon *Icon) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.SmallIcon = smallIcon
	return builder
}

func (builder *vmBuilder) Snapshots(snapshots []Snapshot) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Snapshots = snapshots
	return builder
}

func (builder *vmBuilder) SoundcardEnabled(soundcardEnabled bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.SoundcardEnabled = &soundcardEnabled
	return builder
}

func (builder *vmBuilder) Sso(sso *Sso) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Sso = sso
	return builder
}

func (builder *vmBuilder) StartPaused(startPaused bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.StartPaused = &startPaused
	return builder
}

func (builder *vmBuilder) StartTime(startTime time.Time) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.StartTime = startTime
	return builder
}

func (builder *vmBuilder) Stateless(stateless bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Stateless = &stateless
	return builder
}

func (builder *vmBuilder) Statistics(statistics []Statistic) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Statistics = statistics
	return builder
}

func (builder *vmBuilder) Status(status VmStatus) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Status = status
	return builder
}

func (builder *vmBuilder) StatusDetail(statusDetail string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.StatusDetail = &statusDetail
	return builder
}

func (builder *vmBuilder) StopReason(stopReason string) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.StopReason = &stopReason
	return builder
}

func (builder *vmBuilder) StopTime(stopTime time.Time) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.StopTime = stopTime
	return builder
}

func (builder *vmBuilder) StorageDomain(storageDomain *StorageDomain) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.StorageDomain = storageDomain
	return builder
}

func (builder *vmBuilder) Tags(tags []Tag) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Tags = tags
	return builder
}

func (builder *vmBuilder) Template(template *Template) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Template = template
	return builder
}

func (builder *vmBuilder) TimeZone(timeZone *TimeZone) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.TimeZone = timeZone
	return builder
}

func (builder *vmBuilder) TunnelMigration(tunnelMigration bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.TunnelMigration = &tunnelMigration
	return builder
}

func (builder *vmBuilder) Type(type_ VmType) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Type = type_
	return builder
}

func (builder *vmBuilder) Usb(usb *Usb) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Usb = usb
	return builder
}

func (builder *vmBuilder) UseLatestTemplateVersion(useLatestTemplateVersion bool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.UseLatestTemplateVersion = &useLatestTemplateVersion
	return builder
}

func (builder *vmBuilder) VirtioScsi(virtioScsi *VirtioScsi) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.VirtioScsi = virtioScsi
	return builder
}

func (builder *vmBuilder) VmPool(vmPool *VmPool) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.VmPool = vmPool
	return builder
}

func (builder *vmBuilder) Watchdogs(watchdogs []Watchdog) *vmBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vm.Watchdogs = watchdogs
	return builder
}

func (builder *vmBuilder) Build() (*Vm, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vm, nil
}

type vmBaseBuilder struct {
	vmBase *VmBase
	err    error
}

func NewVmBaseBuilder() *vmBaseBuilder {
	return &vmBaseBuilder{vmBase: &VmBase{}, err: nil}
}

func (builder *vmBaseBuilder) Bios(bios *Bios) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Bios = bios
	return builder
}

func (builder *vmBaseBuilder) Cluster(cluster *Cluster) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Cluster = cluster
	return builder
}

func (builder *vmBaseBuilder) Comment(comment string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Comment = &comment
	return builder
}

func (builder *vmBaseBuilder) Console(console *Console) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Console = console
	return builder
}

func (builder *vmBaseBuilder) Cpu(cpu *Cpu) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Cpu = cpu
	return builder
}

func (builder *vmBaseBuilder) CpuProfile(cpuProfile *CpuProfile) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CpuProfile = cpuProfile
	return builder
}

func (builder *vmBaseBuilder) CpuShares(cpuShares int64) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CpuShares = &cpuShares
	return builder
}

func (builder *vmBaseBuilder) CreationTime(creationTime time.Time) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CreationTime = creationTime
	return builder
}

func (builder *vmBaseBuilder) CustomCompatibilityVersion(customCompatibilityVersion *Version) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CustomCompatibilityVersion = customCompatibilityVersion
	return builder
}

func (builder *vmBaseBuilder) CustomCpuModel(customCpuModel string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CustomCpuModel = &customCpuModel
	return builder
}

func (builder *vmBaseBuilder) CustomEmulatedMachine(customEmulatedMachine string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CustomEmulatedMachine = &customEmulatedMachine
	return builder
}

func (builder *vmBaseBuilder) CustomProperties(customProperties []CustomProperty) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.CustomProperties = customProperties
	return builder
}

func (builder *vmBaseBuilder) DeleteProtected(deleteProtected bool) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.DeleteProtected = &deleteProtected
	return builder
}

func (builder *vmBaseBuilder) Description(description string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Description = &description
	return builder
}

func (builder *vmBaseBuilder) Display(display *Display) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Display = display
	return builder
}

func (builder *vmBaseBuilder) Domain(domain *Domain) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Domain = domain
	return builder
}

func (builder *vmBaseBuilder) HighAvailability(highAvailability *HighAvailability) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.HighAvailability = highAvailability
	return builder
}

func (builder *vmBaseBuilder) Id(id string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Id = &id
	return builder
}

func (builder *vmBaseBuilder) Initialization(initialization *Initialization) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Initialization = initialization
	return builder
}

func (builder *vmBaseBuilder) Io(io *Io) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Io = io
	return builder
}

func (builder *vmBaseBuilder) LargeIcon(largeIcon *Icon) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.LargeIcon = largeIcon
	return builder
}

func (builder *vmBaseBuilder) Lease(lease *StorageDomainLease) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Lease = lease
	return builder
}

func (builder *vmBaseBuilder) Memory(memory int64) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Memory = &memory
	return builder
}

func (builder *vmBaseBuilder) MemoryPolicy(memoryPolicy *MemoryPolicy) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.MemoryPolicy = memoryPolicy
	return builder
}

func (builder *vmBaseBuilder) Migration(migration *MigrationOptions) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Migration = migration
	return builder
}

func (builder *vmBaseBuilder) MigrationDowntime(migrationDowntime int64) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.MigrationDowntime = &migrationDowntime
	return builder
}

func (builder *vmBaseBuilder) Name(name string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Name = &name
	return builder
}

func (builder *vmBaseBuilder) Origin(origin string) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Origin = &origin
	return builder
}

func (builder *vmBaseBuilder) Os(os *OperatingSystem) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Os = os
	return builder
}

func (builder *vmBaseBuilder) Quota(quota *Quota) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Quota = quota
	return builder
}

func (builder *vmBaseBuilder) RngDevice(rngDevice *RngDevice) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.RngDevice = rngDevice
	return builder
}

func (builder *vmBaseBuilder) SerialNumber(serialNumber *SerialNumber) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.SerialNumber = serialNumber
	return builder
}

func (builder *vmBaseBuilder) SmallIcon(smallIcon *Icon) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.SmallIcon = smallIcon
	return builder
}

func (builder *vmBaseBuilder) SoundcardEnabled(soundcardEnabled bool) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.SoundcardEnabled = &soundcardEnabled
	return builder
}

func (builder *vmBaseBuilder) Sso(sso *Sso) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Sso = sso
	return builder
}

func (builder *vmBaseBuilder) StartPaused(startPaused bool) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.StartPaused = &startPaused
	return builder
}

func (builder *vmBaseBuilder) Stateless(stateless bool) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Stateless = &stateless
	return builder
}

func (builder *vmBaseBuilder) StorageDomain(storageDomain *StorageDomain) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.StorageDomain = storageDomain
	return builder
}

func (builder *vmBaseBuilder) TimeZone(timeZone *TimeZone) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.TimeZone = timeZone
	return builder
}

func (builder *vmBaseBuilder) TunnelMigration(tunnelMigration bool) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.TunnelMigration = &tunnelMigration
	return builder
}

func (builder *vmBaseBuilder) Type(type_ VmType) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Type = type_
	return builder
}

func (builder *vmBaseBuilder) Usb(usb *Usb) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.Usb = usb
	return builder
}

func (builder *vmBaseBuilder) VirtioScsi(virtioScsi *VirtioScsi) *vmBaseBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmBase.VirtioScsi = virtioScsi
	return builder
}

func (builder *vmBaseBuilder) Build() (*VmBase, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vmBase, nil
}

type vmPlacementPolicyBuilder struct {
	vmPlacementPolicy *VmPlacementPolicy
	err               error
}

func NewVmPlacementPolicyBuilder() *vmPlacementPolicyBuilder {
	return &vmPlacementPolicyBuilder{vmPlacementPolicy: &VmPlacementPolicy{}, err: nil}
}

func (builder *vmPlacementPolicyBuilder) Affinity(affinity VmAffinity) *vmPlacementPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPlacementPolicy.Affinity = affinity
	return builder
}

func (builder *vmPlacementPolicyBuilder) Hosts(hosts []Host) *vmPlacementPolicyBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPlacementPolicy.Hosts = hosts
	return builder
}

func (builder *vmPlacementPolicyBuilder) Build() (*VmPlacementPolicy, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vmPlacementPolicy, nil
}

type vmPoolBuilder struct {
	vmPool *VmPool
	err    error
}

func NewVmPoolBuilder() *vmPoolBuilder {
	return &vmPoolBuilder{vmPool: &VmPool{}, err: nil}
}

func (builder *vmPoolBuilder) AutoStorageSelect(autoStorageSelect bool) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.AutoStorageSelect = &autoStorageSelect
	return builder
}

func (builder *vmPoolBuilder) Cluster(cluster *Cluster) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Cluster = cluster
	return builder
}

func (builder *vmPoolBuilder) Comment(comment string) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Comment = &comment
	return builder
}

func (builder *vmPoolBuilder) Description(description string) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Description = &description
	return builder
}

func (builder *vmPoolBuilder) Display(display *Display) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Display = display
	return builder
}

func (builder *vmPoolBuilder) Id(id string) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Id = &id
	return builder
}

func (builder *vmPoolBuilder) InstanceType(instanceType *InstanceType) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.InstanceType = instanceType
	return builder
}

func (builder *vmPoolBuilder) MaxUserVms(maxUserVms int64) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.MaxUserVms = &maxUserVms
	return builder
}

func (builder *vmPoolBuilder) Name(name string) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Name = &name
	return builder
}

func (builder *vmPoolBuilder) Permissions(permissions []Permission) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Permissions = permissions
	return builder
}

func (builder *vmPoolBuilder) PrestartedVms(prestartedVms int64) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.PrestartedVms = &prestartedVms
	return builder
}

func (builder *vmPoolBuilder) RngDevice(rngDevice *RngDevice) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.RngDevice = rngDevice
	return builder
}

func (builder *vmPoolBuilder) Size(size int64) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Size = &size
	return builder
}

func (builder *vmPoolBuilder) SoundcardEnabled(soundcardEnabled bool) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.SoundcardEnabled = &soundcardEnabled
	return builder
}

func (builder *vmPoolBuilder) Stateful(stateful bool) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Stateful = &stateful
	return builder
}

func (builder *vmPoolBuilder) Template(template *Template) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Template = template
	return builder
}

func (builder *vmPoolBuilder) Type(type_ VmPoolType) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Type = type_
	return builder
}

func (builder *vmPoolBuilder) UseLatestTemplateVersion(useLatestTemplateVersion bool) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.UseLatestTemplateVersion = &useLatestTemplateVersion
	return builder
}

func (builder *vmPoolBuilder) Vm(vm *Vm) *vmPoolBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmPool.Vm = vm
	return builder
}

func (builder *vmPoolBuilder) Build() (*VmPool, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vmPool, nil
}

type vmSummaryBuilder struct {
	vmSummary *VmSummary
	err       error
}

func NewVmSummaryBuilder() *vmSummaryBuilder {
	return &vmSummaryBuilder{vmSummary: &VmSummary{}, err: nil}
}

func (builder *vmSummaryBuilder) Active(active int64) *vmSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmSummary.Active = &active
	return builder
}

func (builder *vmSummaryBuilder) Migrating(migrating int64) *vmSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmSummary.Migrating = &migrating
	return builder
}

func (builder *vmSummaryBuilder) Total(total int64) *vmSummaryBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vmSummary.Total = &total
	return builder
}

func (builder *vmSummaryBuilder) Build() (*VmSummary, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vmSummary, nil
}

type vnicPassThroughBuilder struct {
	vnicPassThrough *VnicPassThrough
	err             error
}

func NewVnicPassThroughBuilder() *vnicPassThroughBuilder {
	return &vnicPassThroughBuilder{vnicPassThrough: &VnicPassThrough{}, err: nil}
}

func (builder *vnicPassThroughBuilder) Mode(mode VnicPassThroughMode) *vnicPassThroughBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicPassThrough.Mode = mode
	return builder
}

func (builder *vnicPassThroughBuilder) Build() (*VnicPassThrough, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vnicPassThrough, nil
}

type vnicProfileBuilder struct {
	vnicProfile *VnicProfile
	err         error
}

func NewVnicProfileBuilder() *vnicProfileBuilder {
	return &vnicProfileBuilder{vnicProfile: &VnicProfile{}, err: nil}
}

func (builder *vnicProfileBuilder) Comment(comment string) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Comment = &comment
	return builder
}

func (builder *vnicProfileBuilder) CustomProperties(customProperties []CustomProperty) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.CustomProperties = customProperties
	return builder
}

func (builder *vnicProfileBuilder) Description(description string) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Description = &description
	return builder
}

func (builder *vnicProfileBuilder) Id(id string) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Id = &id
	return builder
}

func (builder *vnicProfileBuilder) Migratable(migratable bool) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Migratable = &migratable
	return builder
}

func (builder *vnicProfileBuilder) Name(name string) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Name = &name
	return builder
}

func (builder *vnicProfileBuilder) Network(network *Network) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Network = network
	return builder
}

func (builder *vnicProfileBuilder) NetworkFilter(networkFilter *NetworkFilter) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.NetworkFilter = networkFilter
	return builder
}

func (builder *vnicProfileBuilder) PassThrough(passThrough *VnicPassThrough) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.PassThrough = passThrough
	return builder
}

func (builder *vnicProfileBuilder) Permissions(permissions []Permission) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Permissions = permissions
	return builder
}

func (builder *vnicProfileBuilder) PortMirroring(portMirroring bool) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.PortMirroring = &portMirroring
	return builder
}

func (builder *vnicProfileBuilder) Qos(qos *Qos) *vnicProfileBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfile.Qos = qos
	return builder
}

func (builder *vnicProfileBuilder) Build() (*VnicProfile, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vnicProfile, nil
}

type vnicProfileMappingBuilder struct {
	vnicProfileMapping *VnicProfileMapping
	err                error
}

func NewVnicProfileMappingBuilder() *vnicProfileMappingBuilder {
	return &vnicProfileMappingBuilder{vnicProfileMapping: &VnicProfileMapping{}, err: nil}
}

func (builder *vnicProfileMappingBuilder) SourceNetworkName(sourceNetworkName string) *vnicProfileMappingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfileMapping.SourceNetworkName = &sourceNetworkName
	return builder
}

func (builder *vnicProfileMappingBuilder) SourceNetworkProfileName(sourceNetworkProfileName string) *vnicProfileMappingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfileMapping.SourceNetworkProfileName = &sourceNetworkProfileName
	return builder
}

func (builder *vnicProfileMappingBuilder) TargetVnicProfile(targetVnicProfile *VnicProfile) *vnicProfileMappingBuilder {
	if builder.err != nil {
		return builder
	}

	builder.vnicProfileMapping.TargetVnicProfile = targetVnicProfile
	return builder
}

func (builder *vnicProfileMappingBuilder) Build() (*VnicProfileMapping, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.vnicProfileMapping, nil
}

type volumeGroupBuilder struct {
	volumeGroup *VolumeGroup
	err         error
}

func NewVolumeGroupBuilder() *volumeGroupBuilder {
	return &volumeGroupBuilder{volumeGroup: &VolumeGroup{}, err: nil}
}

func (builder *volumeGroupBuilder) Id(id string) *volumeGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.volumeGroup.Id = &id
	return builder
}

func (builder *volumeGroupBuilder) LogicalUnits(logicalUnits []LogicalUnit) *volumeGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.volumeGroup.LogicalUnits = logicalUnits
	return builder
}

func (builder *volumeGroupBuilder) Name(name string) *volumeGroupBuilder {
	if builder.err != nil {
		return builder
	}

	builder.volumeGroup.Name = &name
	return builder
}

func (builder *volumeGroupBuilder) Build() (*VolumeGroup, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.volumeGroup, nil
}

type watchdogBuilder struct {
	watchdog *Watchdog
	err      error
}

func NewWatchdogBuilder() *watchdogBuilder {
	return &watchdogBuilder{watchdog: &Watchdog{}, err: nil}
}

func (builder *watchdogBuilder) Action(action WatchdogAction) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Action = action
	return builder
}

func (builder *watchdogBuilder) Comment(comment string) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Comment = &comment
	return builder
}

func (builder *watchdogBuilder) Description(description string) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Description = &description
	return builder
}

func (builder *watchdogBuilder) Id(id string) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Id = &id
	return builder
}

func (builder *watchdogBuilder) InstanceType(instanceType *InstanceType) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.InstanceType = instanceType
	return builder
}

func (builder *watchdogBuilder) Model(model WatchdogModel) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Model = model
	return builder
}

func (builder *watchdogBuilder) Name(name string) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Name = &name
	return builder
}

func (builder *watchdogBuilder) Template(template *Template) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Template = template
	return builder
}

func (builder *watchdogBuilder) Vm(vm *Vm) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Vm = vm
	return builder
}

func (builder *watchdogBuilder) Vms(vms []Vm) *watchdogBuilder {
	if builder.err != nil {
		return builder
	}

	builder.watchdog.Vms = vms
	return builder
}

func (builder *watchdogBuilder) Build() (*Watchdog, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.watchdog, nil
}

type weightBuilder struct {
	weight *Weight
	err    error
}

func NewWeightBuilder() *weightBuilder {
	return &weightBuilder{weight: &Weight{}, err: nil}
}

func (builder *weightBuilder) Comment(comment string) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.Comment = &comment
	return builder
}

func (builder *weightBuilder) Description(description string) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.Description = &description
	return builder
}

func (builder *weightBuilder) Factor(factor int64) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.Factor = &factor
	return builder
}

func (builder *weightBuilder) Id(id string) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.Id = &id
	return builder
}

func (builder *weightBuilder) Name(name string) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.Name = &name
	return builder
}

func (builder *weightBuilder) SchedulingPolicy(schedulingPolicy *SchedulingPolicy) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.SchedulingPolicy = schedulingPolicy
	return builder
}

func (builder *weightBuilder) SchedulingPolicyUnit(schedulingPolicyUnit *SchedulingPolicyUnit) *weightBuilder {
	if builder.err != nil {
		return builder
	}

	builder.weight.SchedulingPolicyUnit = schedulingPolicyUnit
	return builder
}

func (builder *weightBuilder) Build() (*Weight, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.weight, nil
}

type AccessProtocol string

const (
	ACCESSPROTOCOL_CIFS    AccessProtocol = "cifs"
	ACCESSPROTOCOL_GLUSTER AccessProtocol = "gluster"
	ACCESSPROTOCOL_NFS     AccessProtocol = "nfs"
)

type Architecture string

const (
	ARCHITECTURE_PPC64     Architecture = "ppc64"
	ARCHITECTURE_UNDEFINED Architecture = "undefined"
	ARCHITECTURE_X86_64    Architecture = "x86_64"
)

type AutoNumaStatus string

const (
	AUTONUMASTATUS_DISABLE AutoNumaStatus = "disable"
	AUTONUMASTATUS_ENABLE  AutoNumaStatus = "enable"
	AUTONUMASTATUS_UNKNOWN AutoNumaStatus = "unknown"
)

type BootDevice string

const (
	BOOTDEVICE_CDROM   BootDevice = "cdrom"
	BOOTDEVICE_HD      BootDevice = "hd"
	BOOTDEVICE_NETWORK BootDevice = "network"
)

type BootProtocol string

const (
	BOOTPROTOCOL_AUTOCONF BootProtocol = "autoconf"
	BOOTPROTOCOL_DHCP     BootProtocol = "dhcp"
	BOOTPROTOCOL_NONE     BootProtocol = "none"
	BOOTPROTOCOL_STATIC   BootProtocol = "static"
)

type ConfigurationType string

const (
	CONFIGURATIONTYPE_OVF ConfigurationType = "ovf"
)

type CpuMode string

const (
	CPUMODE_CUSTOM           CpuMode = "custom"
	CPUMODE_HOST_MODEL       CpuMode = "host_model"
	CPUMODE_HOST_PASSTHROUGH CpuMode = "host_passthrough"
)

type CreationStatus string

const (
	CREATIONSTATUS_COMPLETE    CreationStatus = "complete"
	CREATIONSTATUS_FAILED      CreationStatus = "failed"
	CREATIONSTATUS_IN_PROGRESS CreationStatus = "in_progress"
	CREATIONSTATUS_PENDING     CreationStatus = "pending"
)

type DataCenterStatus string

const (
	DATACENTERSTATUS_CONTEND         DataCenterStatus = "contend"
	DATACENTERSTATUS_MAINTENANCE     DataCenterStatus = "maintenance"
	DATACENTERSTATUS_NOT_OPERATIONAL DataCenterStatus = "not_operational"
	DATACENTERSTATUS_PROBLEMATIC     DataCenterStatus = "problematic"
	DATACENTERSTATUS_UNINITIALIZED   DataCenterStatus = "uninitialized"
	DATACENTERSTATUS_UP              DataCenterStatus = "up"
)

type DiskFormat string

const (
	DISKFORMAT_COW DiskFormat = "cow"
	DISKFORMAT_RAW DiskFormat = "raw"
)

type DiskInterface string

const (
	DISKINTERFACE_IDE         DiskInterface = "ide"
	DISKINTERFACE_SPAPR_VSCSI DiskInterface = "spapr_vscsi"
	DISKINTERFACE_VIRTIO      DiskInterface = "virtio"
	DISKINTERFACE_VIRTIO_SCSI DiskInterface = "virtio_scsi"
)

type DiskStatus string

const (
	DISKSTATUS_ILLEGAL DiskStatus = "illegal"
	DISKSTATUS_LOCKED  DiskStatus = "locked"
	DISKSTATUS_OK      DiskStatus = "ok"
)

type DiskStorageType string

const (
	DISKSTORAGETYPE_CINDER DiskStorageType = "cinder"
	DISKSTORAGETYPE_IMAGE  DiskStorageType = "image"
	DISKSTORAGETYPE_LUN    DiskStorageType = "lun"
)

type DiskType string

const (
	DISKTYPE_DATA   DiskType = "data"
	DISKTYPE_SYSTEM DiskType = "system"
)

type DisplayType string

const (
	DISPLAYTYPE_SPICE DisplayType = "spice"
	DISPLAYTYPE_VNC   DisplayType = "vnc"
)

type EntityExternalStatus string

const (
	ENTITYEXTERNALSTATUS_ERROR   EntityExternalStatus = "error"
	ENTITYEXTERNALSTATUS_FAILURE EntityExternalStatus = "failure"
	ENTITYEXTERNALSTATUS_INFO    EntityExternalStatus = "info"
	ENTITYEXTERNALSTATUS_OK      EntityExternalStatus = "ok"
	ENTITYEXTERNALSTATUS_WARNING EntityExternalStatus = "warning"
)

type ExternalStatus string

const (
	EXTERNALSTATUS_ERROR   ExternalStatus = "error"
	EXTERNALSTATUS_FAILURE ExternalStatus = "failure"
	EXTERNALSTATUS_INFO    ExternalStatus = "info"
	EXTERNALSTATUS_OK      ExternalStatus = "ok"
	EXTERNALSTATUS_WARNING ExternalStatus = "warning"
)

type ExternalSystemType string

const (
	EXTERNALSYSTEMTYPE_GLUSTER ExternalSystemType = "gluster"
	EXTERNALSYSTEMTYPE_VDSM    ExternalSystemType = "vdsm"
)

type ExternalVmProviderType string

const (
	EXTERNALVMPROVIDERTYPE_KVM    ExternalVmProviderType = "kvm"
	EXTERNALVMPROVIDERTYPE_VMWARE ExternalVmProviderType = "vmware"
	EXTERNALVMPROVIDERTYPE_XEN    ExternalVmProviderType = "xen"
)

type FenceType string

const (
	FENCETYPE_MANUAL  FenceType = "manual"
	FENCETYPE_RESTART FenceType = "restart"
	FENCETYPE_START   FenceType = "start"
	FENCETYPE_STATUS  FenceType = "status"
	FENCETYPE_STOP    FenceType = "stop"
)

type GlusterBrickStatus string

const (
	GLUSTERBRICKSTATUS_DOWN    GlusterBrickStatus = "down"
	GLUSTERBRICKSTATUS_UNKNOWN GlusterBrickStatus = "unknown"
	GLUSTERBRICKSTATUS_UP      GlusterBrickStatus = "up"
)

type GlusterHookStatus string

const (
	GLUSTERHOOKSTATUS_DISABLED GlusterHookStatus = "disabled"
	GLUSTERHOOKSTATUS_ENABLED  GlusterHookStatus = "enabled"
	GLUSTERHOOKSTATUS_MISSING  GlusterHookStatus = "missing"
)

type GlusterState string

const (
	GLUSTERSTATE_DOWN    GlusterState = "down"
	GLUSTERSTATE_UNKNOWN GlusterState = "unknown"
	GLUSTERSTATE_UP      GlusterState = "up"
)

type GlusterVolumeStatus string

const (
	GLUSTERVOLUMESTATUS_DOWN    GlusterVolumeStatus = "down"
	GLUSTERVOLUMESTATUS_UNKNOWN GlusterVolumeStatus = "unknown"
	GLUSTERVOLUMESTATUS_UP      GlusterVolumeStatus = "up"
)

type GlusterVolumeType string

const (
	GLUSTERVOLUMETYPE_DISPERSE                      GlusterVolumeType = "disperse"
	GLUSTERVOLUMETYPE_DISTRIBUTE                    GlusterVolumeType = "distribute"
	GLUSTERVOLUMETYPE_DISTRIBUTED_DISPERSE          GlusterVolumeType = "distributed_disperse"
	GLUSTERVOLUMETYPE_DISTRIBUTED_REPLICATE         GlusterVolumeType = "distributed_replicate"
	GLUSTERVOLUMETYPE_DISTRIBUTED_STRIPE            GlusterVolumeType = "distributed_stripe"
	GLUSTERVOLUMETYPE_DISTRIBUTED_STRIPED_REPLICATE GlusterVolumeType = "distributed_striped_replicate"
	GLUSTERVOLUMETYPE_REPLICATE                     GlusterVolumeType = "replicate"
	GLUSTERVOLUMETYPE_STRIPE                        GlusterVolumeType = "stripe"
	GLUSTERVOLUMETYPE_STRIPED_REPLICATE             GlusterVolumeType = "striped_replicate"
)

type GraphicsType string

const (
	GRAPHICSTYPE_SPICE GraphicsType = "spice"
	GRAPHICSTYPE_VNC   GraphicsType = "vnc"
)

type HookContentType string

const (
	HOOKCONTENTTYPE_BINARY HookContentType = "binary"
	HOOKCONTENTTYPE_TEXT   HookContentType = "text"
)

type HookStage string

const (
	HOOKSTAGE_POST HookStage = "post"
	HOOKSTAGE_PRE  HookStage = "pre"
)

type HookStatus string

const (
	HOOKSTATUS_DISABLED HookStatus = "disabled"
	HOOKSTATUS_ENABLED  HookStatus = "enabled"
	HOOKSTATUS_MISSING  HookStatus = "missing"
)

type HostProtocol string

const (
	HOSTPROTOCOL_STOMP HostProtocol = "stomp"
	HOSTPROTOCOL_XML   HostProtocol = "xml"
)

type HostStatus string

const (
	HOSTSTATUS_CONNECTING                HostStatus = "connecting"
	HOSTSTATUS_DOWN                      HostStatus = "down"
	HOSTSTATUS_ERROR                     HostStatus = "error"
	HOSTSTATUS_INITIALIZING              HostStatus = "initializing"
	HOSTSTATUS_INSTALL_FAILED            HostStatus = "install_failed"
	HOSTSTATUS_INSTALLING                HostStatus = "installing"
	HOSTSTATUS_INSTALLING_OS             HostStatus = "installing_os"
	HOSTSTATUS_KDUMPING                  HostStatus = "kdumping"
	HOSTSTATUS_MAINTENANCE               HostStatus = "maintenance"
	HOSTSTATUS_NON_OPERATIONAL           HostStatus = "non_operational"
	HOSTSTATUS_NON_RESPONSIVE            HostStatus = "non_responsive"
	HOSTSTATUS_PENDING_APPROVAL          HostStatus = "pending_approval"
	HOSTSTATUS_PREPARING_FOR_MAINTENANCE HostStatus = "preparing_for_maintenance"
	HOSTSTATUS_REBOOT                    HostStatus = "reboot"
	HOSTSTATUS_UNASSIGNED                HostStatus = "unassigned"
	HOSTSTATUS_UP                        HostStatus = "up"
)

type HostType string

const (
	HOSTTYPE_OVIRT_NODE HostType = "ovirt_node"
	HOSTTYPE_RHEL       HostType = "rhel"
	HOSTTYPE_RHEV_H     HostType = "rhev_h"
)

type ImageTransferDirection string

const (
	IMAGETRANSFERDIRECTION_DOWNLOAD ImageTransferDirection = "download"
	IMAGETRANSFERDIRECTION_UPLOAD   ImageTransferDirection = "upload"
)

type ImageTransferPhase string

const (
	IMAGETRANSFERPHASE_CANCELLED          ImageTransferPhase = "cancelled"
	IMAGETRANSFERPHASE_FINALIZING_FAILURE ImageTransferPhase = "finalizing_failure"
	IMAGETRANSFERPHASE_FINALIZING_SUCCESS ImageTransferPhase = "finalizing_success"
	IMAGETRANSFERPHASE_FINISHED_FAILURE   ImageTransferPhase = "finished_failure"
	IMAGETRANSFERPHASE_FINISHED_SUCCESS   ImageTransferPhase = "finished_success"
	IMAGETRANSFERPHASE_INITIALIZING       ImageTransferPhase = "initializing"
	IMAGETRANSFERPHASE_PAUSED_SYSTEM      ImageTransferPhase = "paused_system"
	IMAGETRANSFERPHASE_PAUSED_USER        ImageTransferPhase = "paused_user"
	IMAGETRANSFERPHASE_RESUMING           ImageTransferPhase = "resuming"
	IMAGETRANSFERPHASE_TRANSFERRING       ImageTransferPhase = "transferring"
	IMAGETRANSFERPHASE_UNKNOWN            ImageTransferPhase = "unknown"
)

type InheritableBoolean string

const (
	INHERITABLEBOOLEAN_FALSE   InheritableBoolean = "false"
	INHERITABLEBOOLEAN_INHERIT InheritableBoolean = "inherit"
	INHERITABLEBOOLEAN_TRUE    InheritableBoolean = "true"
)

type IpVersion string

const (
	IPVERSION_V4 IpVersion = "v4"
	IPVERSION_V6 IpVersion = "v6"
)

type JobStatus string

const (
	JOBSTATUS_ABORTED  JobStatus = "aborted"
	JOBSTATUS_FAILED   JobStatus = "failed"
	JOBSTATUS_FINISHED JobStatus = "finished"
	JOBSTATUS_STARTED  JobStatus = "started"
	JOBSTATUS_UNKNOWN  JobStatus = "unknown"
)

type KdumpStatus string

const (
	KDUMPSTATUS_DISABLED KdumpStatus = "disabled"
	KDUMPSTATUS_ENABLED  KdumpStatus = "enabled"
	KDUMPSTATUS_UNKNOWN  KdumpStatus = "unknown"
)

type LogSeverity string

const (
	LOGSEVERITY_ALERT   LogSeverity = "alert"
	LOGSEVERITY_ERROR   LogSeverity = "error"
	LOGSEVERITY_NORMAL  LogSeverity = "normal"
	LOGSEVERITY_WARNING LogSeverity = "warning"
)

type LunStatus string

const (
	LUNSTATUS_FREE     LunStatus = "free"
	LUNSTATUS_UNUSABLE LunStatus = "unusable"
	LUNSTATUS_USED     LunStatus = "used"
)

type MessageBrokerType string

const (
	MESSAGEBROKERTYPE_QPID      MessageBrokerType = "qpid"
	MESSAGEBROKERTYPE_RABBIT_MQ MessageBrokerType = "rabbit_mq"
)

type MigrateOnError string

const (
	MIGRATEONERROR_DO_NOT_MIGRATE           MigrateOnError = "do_not_migrate"
	MIGRATEONERROR_MIGRATE                  MigrateOnError = "migrate"
	MIGRATEONERROR_MIGRATE_HIGHLY_AVAILABLE MigrateOnError = "migrate_highly_available"
)

type MigrationBandwidthAssignmentMethod string

const (
	MIGRATIONBANDWIDTHASSIGNMENTMETHOD_AUTO               MigrationBandwidthAssignmentMethod = "auto"
	MIGRATIONBANDWIDTHASSIGNMENTMETHOD_CUSTOM             MigrationBandwidthAssignmentMethod = "custom"
	MIGRATIONBANDWIDTHASSIGNMENTMETHOD_HYPERVISOR_DEFAULT MigrationBandwidthAssignmentMethod = "hypervisor_default"
)

type NetworkPluginType string

const (
	NETWORKPLUGINTYPE_OPEN_VSWITCH NetworkPluginType = "open_vswitch"
)

type NetworkStatus string

const (
	NETWORKSTATUS_NON_OPERATIONAL NetworkStatus = "non_operational"
	NETWORKSTATUS_OPERATIONAL     NetworkStatus = "operational"
)

type NetworkUsage string

const (
	NETWORKUSAGE_DEFAULT_ROUTE NetworkUsage = "default_route"
	NETWORKUSAGE_DISPLAY       NetworkUsage = "display"
	NETWORKUSAGE_GLUSTER       NetworkUsage = "gluster"
	NETWORKUSAGE_MANAGEMENT    NetworkUsage = "management"
	NETWORKUSAGE_MIGRATION     NetworkUsage = "migration"
	NETWORKUSAGE_VM            NetworkUsage = "vm"
)

type NfsVersion string

const (
	NFSVERSION_AUTO NfsVersion = "auto"
	NFSVERSION_V3   NfsVersion = "v3"
	NFSVERSION_V4   NfsVersion = "v4"
	NFSVERSION_V4_1 NfsVersion = "v4_1"
	NFSVERSION_V4_2 NfsVersion = "v4_2"
)

type NicInterface string

const (
	NICINTERFACE_E1000           NicInterface = "e1000"
	NICINTERFACE_PCI_PASSTHROUGH NicInterface = "pci_passthrough"
	NICINTERFACE_RTL8139         NicInterface = "rtl8139"
	NICINTERFACE_RTL8139_VIRTIO  NicInterface = "rtl8139_virtio"
	NICINTERFACE_SPAPR_VLAN      NicInterface = "spapr_vlan"
	NICINTERFACE_VIRTIO          NicInterface = "virtio"
)

type NicStatus string

const (
	NICSTATUS_DOWN NicStatus = "down"
	NICSTATUS_UP   NicStatus = "up"
)

type NumaTuneMode string

const (
	NUMATUNEMODE_INTERLEAVE NumaTuneMode = "interleave"
	NUMATUNEMODE_PREFERRED  NumaTuneMode = "preferred"
	NUMATUNEMODE_STRICT     NumaTuneMode = "strict"
)

type OpenStackNetworkProviderType string

const (
	OPENSTACKNETWORKPROVIDERTYPE_EXTERNAL OpenStackNetworkProviderType = "external"
	OPENSTACKNETWORKPROVIDERTYPE_NEUTRON  OpenStackNetworkProviderType = "neutron"
)

type OpenstackVolumeAuthenticationKeyUsageType string

const (
	OPENSTACKVOLUMEAUTHENTICATIONKEYUSAGETYPE_CEPH OpenstackVolumeAuthenticationKeyUsageType = "ceph"
)

type OsType string

const (
	OSTYPE_OTHER             OsType = "other"
	OSTYPE_OTHER_LINUX       OsType = "other_linux"
	OSTYPE_RHEL_3            OsType = "rhel_3"
	OSTYPE_RHEL_3X64         OsType = "rhel_3x64"
	OSTYPE_RHEL_4            OsType = "rhel_4"
	OSTYPE_RHEL_4X64         OsType = "rhel_4x64"
	OSTYPE_RHEL_5            OsType = "rhel_5"
	OSTYPE_RHEL_5X64         OsType = "rhel_5x64"
	OSTYPE_RHEL_6            OsType = "rhel_6"
	OSTYPE_RHEL_6X64         OsType = "rhel_6x64"
	OSTYPE_UNASSIGNED        OsType = "unassigned"
	OSTYPE_WINDOWS_2003      OsType = "windows_2003"
	OSTYPE_WINDOWS_2003X64   OsType = "windows_2003x64"
	OSTYPE_WINDOWS_2008      OsType = "windows_2008"
	OSTYPE_WINDOWS_2008R2X64 OsType = "windows_2008r2x64"
	OSTYPE_WINDOWS_2008X64   OsType = "windows_2008x64"
	OSTYPE_WINDOWS_2012X64   OsType = "windows_2012x64"
	OSTYPE_WINDOWS_7         OsType = "windows_7"
	OSTYPE_WINDOWS_7X64      OsType = "windows_7x64"
	OSTYPE_WINDOWS_8         OsType = "windows_8"
	OSTYPE_WINDOWS_8X64      OsType = "windows_8x64"
	OSTYPE_WINDOWS_XP        OsType = "windows_xp"
)

type PayloadEncoding string

const (
	PAYLOADENCODING_BASE64    PayloadEncoding = "base64"
	PAYLOADENCODING_PLAINTEXT PayloadEncoding = "plaintext"
)

type PmProxyType string

const (
	PMPROXYTYPE_CLUSTER  PmProxyType = "cluster"
	PMPROXYTYPE_DC       PmProxyType = "dc"
	PMPROXYTYPE_OTHER_DC PmProxyType = "other_dc"
)

type PolicyUnitType string

const (
	POLICYUNITTYPE_FILTER         PolicyUnitType = "filter"
	POLICYUNITTYPE_LOAD_BALANCING PolicyUnitType = "load_balancing"
	POLICYUNITTYPE_WEIGHT         PolicyUnitType = "weight"
)

type PowerManagementStatus string

const (
	POWERMANAGEMENTSTATUS_OFF     PowerManagementStatus = "off"
	POWERMANAGEMENTSTATUS_ON      PowerManagementStatus = "on"
	POWERMANAGEMENTSTATUS_UNKNOWN PowerManagementStatus = "unknown"
)

type QcowVersion string

const (
	QCOWVERSION_QCOW2_V2 QcowVersion = "qcow2_v2"
	QCOWVERSION_QCOW2_V3 QcowVersion = "qcow2_v3"
)

type QosType string

const (
	QOSTYPE_CPU         QosType = "cpu"
	QOSTYPE_HOSTNETWORK QosType = "hostnetwork"
	QOSTYPE_NETWORK     QosType = "network"
	QOSTYPE_STORAGE     QosType = "storage"
)

type QuotaModeType string

const (
	QUOTAMODETYPE_AUDIT    QuotaModeType = "audit"
	QUOTAMODETYPE_DISABLED QuotaModeType = "disabled"
	QUOTAMODETYPE_ENABLED  QuotaModeType = "enabled"
)

type ReportedDeviceType string

const (
	REPORTEDDEVICETYPE_NETWORK ReportedDeviceType = "network"
)

type ResolutionType string

const (
	RESOLUTIONTYPE_ADD  ResolutionType = "add"
	RESOLUTIONTYPE_COPY ResolutionType = "copy"
)

type RngSource string

const (
	RNGSOURCE_HWRNG   RngSource = "hwrng"
	RNGSOURCE_RANDOM  RngSource = "random"
	RNGSOURCE_URANDOM RngSource = "urandom"
)

type RoleType string

const (
	ROLETYPE_ADMIN RoleType = "admin"
	ROLETYPE_USER  RoleType = "user"
)

type ScsiGenericIO string

const (
	SCSIGENERICIO_FILTERED   ScsiGenericIO = "filtered"
	SCSIGENERICIO_UNFILTERED ScsiGenericIO = "unfiltered"
)

type SeLinuxMode string

const (
	SELINUXMODE_DISABLED   SeLinuxMode = "disabled"
	SELINUXMODE_ENFORCING  SeLinuxMode = "enforcing"
	SELINUXMODE_PERMISSIVE SeLinuxMode = "permissive"
)

type SerialNumberPolicy string

const (
	SERIALNUMBERPOLICY_CUSTOM SerialNumberPolicy = "custom"
	SERIALNUMBERPOLICY_HOST   SerialNumberPolicy = "host"
	SERIALNUMBERPOLICY_VM     SerialNumberPolicy = "vm"
)

type SnapshotStatus string

const (
	SNAPSHOTSTATUS_IN_PREVIEW SnapshotStatus = "in_preview"
	SNAPSHOTSTATUS_LOCKED     SnapshotStatus = "locked"
	SNAPSHOTSTATUS_OK         SnapshotStatus = "ok"
)

type SnapshotType string

const (
	SNAPSHOTTYPE_ACTIVE    SnapshotType = "active"
	SNAPSHOTTYPE_PREVIEW   SnapshotType = "preview"
	SNAPSHOTTYPE_REGULAR   SnapshotType = "regular"
	SNAPSHOTTYPE_STATELESS SnapshotType = "stateless"
)

type SpmStatus string

const (
	SPMSTATUS_CONTENDING SpmStatus = "contending"
	SPMSTATUS_NONE       SpmStatus = "none"
	SPMSTATUS_SPM        SpmStatus = "spm"
)

type SshAuthenticationMethod string

const (
	SSHAUTHENTICATIONMETHOD_PASSWORD  SshAuthenticationMethod = "password"
	SSHAUTHENTICATIONMETHOD_PUBLICKEY SshAuthenticationMethod = "publickey"
)

type SsoMethod string

const (
	SSOMETHOD_GUEST_AGENT SsoMethod = "guest_agent"
)

type StatisticKind string

const (
	STATISTICKIND_COUNTER StatisticKind = "counter"
	STATISTICKIND_GAUGE   StatisticKind = "gauge"
)

type StatisticUnit string

const (
	STATISTICUNIT_BITS_PER_SECOND  StatisticUnit = "bits_per_second"
	STATISTICUNIT_BYTES            StatisticUnit = "bytes"
	STATISTICUNIT_BYTES_PER_SECOND StatisticUnit = "bytes_per_second"
	STATISTICUNIT_COUNT_PER_SECOND StatisticUnit = "count_per_second"
	STATISTICUNIT_NONE             StatisticUnit = "none"
	STATISTICUNIT_PERCENT          StatisticUnit = "percent"
	STATISTICUNIT_SECONDS          StatisticUnit = "seconds"
)

type StepEnum string

const (
	STEPENUM_EXECUTING          StepEnum = "executing"
	STEPENUM_FINALIZING         StepEnum = "finalizing"
	STEPENUM_REBALANCING_VOLUME StepEnum = "rebalancing_volume"
	STEPENUM_REMOVING_BRICKS    StepEnum = "removing_bricks"
	STEPENUM_UNKNOWN            StepEnum = "unknown"
	STEPENUM_VALIDATING         StepEnum = "validating"
)

type StepStatus string

const (
	STEPSTATUS_ABORTED  StepStatus = "aborted"
	STEPSTATUS_FAILED   StepStatus = "failed"
	STEPSTATUS_FINISHED StepStatus = "finished"
	STEPSTATUS_STARTED  StepStatus = "started"
	STEPSTATUS_UNKNOWN  StepStatus = "unknown"
)

type StorageDomainStatus string

const (
	STORAGEDOMAINSTATUS_ACTIVATING                StorageDomainStatus = "activating"
	STORAGEDOMAINSTATUS_ACTIVE                    StorageDomainStatus = "active"
	STORAGEDOMAINSTATUS_DETACHING                 StorageDomainStatus = "detaching"
	STORAGEDOMAINSTATUS_INACTIVE                  StorageDomainStatus = "inactive"
	STORAGEDOMAINSTATUS_LOCKED                    StorageDomainStatus = "locked"
	STORAGEDOMAINSTATUS_MAINTENANCE               StorageDomainStatus = "maintenance"
	STORAGEDOMAINSTATUS_MIXED                     StorageDomainStatus = "mixed"
	STORAGEDOMAINSTATUS_PREPARING_FOR_MAINTENANCE StorageDomainStatus = "preparing_for_maintenance"
	STORAGEDOMAINSTATUS_UNATTACHED                StorageDomainStatus = "unattached"
	STORAGEDOMAINSTATUS_UNKNOWN                   StorageDomainStatus = "unknown"
)

type StorageDomainType string

const (
	STORAGEDOMAINTYPE_DATA   StorageDomainType = "data"
	STORAGEDOMAINTYPE_EXPORT StorageDomainType = "export"
	STORAGEDOMAINTYPE_IMAGE  StorageDomainType = "image"
	STORAGEDOMAINTYPE_ISO    StorageDomainType = "iso"
	STORAGEDOMAINTYPE_VOLUME StorageDomainType = "volume"
)

type StorageFormat string

const (
	STORAGEFORMAT_V1 StorageFormat = "v1"
	STORAGEFORMAT_V2 StorageFormat = "v2"
	STORAGEFORMAT_V3 StorageFormat = "v3"
	STORAGEFORMAT_V4 StorageFormat = "v4"
)

type StorageType string

const (
	STORAGETYPE_CINDER    StorageType = "cinder"
	STORAGETYPE_FCP       StorageType = "fcp"
	STORAGETYPE_GLANCE    StorageType = "glance"
	STORAGETYPE_GLUSTERFS StorageType = "glusterfs"
	STORAGETYPE_ISCSI     StorageType = "iscsi"
	STORAGETYPE_LOCALFS   StorageType = "localfs"
	STORAGETYPE_NFS       StorageType = "nfs"
	STORAGETYPE_POSIXFS   StorageType = "posixfs"
)

type SwitchType string

const (
	SWITCHTYPE_LEGACY SwitchType = "legacy"
	SWITCHTYPE_OVS    SwitchType = "ovs"
)

type TemplateStatus string

const (
	TEMPLATESTATUS_ILLEGAL TemplateStatus = "illegal"
	TEMPLATESTATUS_LOCKED  TemplateStatus = "locked"
	TEMPLATESTATUS_OK      TemplateStatus = "ok"
)

type TransportType string

const (
	TRANSPORTTYPE_RDMA TransportType = "rdma"
	TRANSPORTTYPE_TCP  TransportType = "tcp"
)

type UsbType string

const (
	USBTYPE_LEGACY UsbType = "legacy"
	USBTYPE_NATIVE UsbType = "native"
)

type ValueType string

const (
	VALUETYPE_DECIMAL ValueType = "decimal"
	VALUETYPE_INTEGER ValueType = "integer"
	VALUETYPE_STRING  ValueType = "string"
)

type VmAffinity string

const (
	VMAFFINITY_MIGRATABLE      VmAffinity = "migratable"
	VMAFFINITY_PINNED          VmAffinity = "pinned"
	VMAFFINITY_USER_MIGRATABLE VmAffinity = "user_migratable"
)

type VmDeviceType string

const (
	VMDEVICETYPE_CDROM  VmDeviceType = "cdrom"
	VMDEVICETYPE_FLOPPY VmDeviceType = "floppy"
)

type VmPoolType string

const (
	VMPOOLTYPE_AUTOMATIC VmPoolType = "automatic"
	VMPOOLTYPE_MANUAL    VmPoolType = "manual"
)

type VmStatus string

const (
	VMSTATUS_DOWN               VmStatus = "down"
	VMSTATUS_IMAGE_LOCKED       VmStatus = "image_locked"
	VMSTATUS_MIGRATING          VmStatus = "migrating"
	VMSTATUS_NOT_RESPONDING     VmStatus = "not_responding"
	VMSTATUS_PAUSED             VmStatus = "paused"
	VMSTATUS_POWERING_DOWN      VmStatus = "powering_down"
	VMSTATUS_POWERING_UP        VmStatus = "powering_up"
	VMSTATUS_REBOOT_IN_PROGRESS VmStatus = "reboot_in_progress"
	VMSTATUS_RESTORING_STATE    VmStatus = "restoring_state"
	VMSTATUS_SAVING_STATE       VmStatus = "saving_state"
	VMSTATUS_SUSPENDED          VmStatus = "suspended"
	VMSTATUS_UNASSIGNED         VmStatus = "unassigned"
	VMSTATUS_UNKNOWN            VmStatus = "unknown"
	VMSTATUS_UP                 VmStatus = "up"
	VMSTATUS_WAIT_FOR_LAUNCH    VmStatus = "wait_for_launch"
)

type VmType string

const (
	VMTYPE_DESKTOP VmType = "desktop"
	VMTYPE_SERVER  VmType = "server"
)

type VnicPassThroughMode string

const (
	VNICPASSTHROUGHMODE_DISABLED VnicPassThroughMode = "disabled"
	VNICPASSTHROUGHMODE_ENABLED  VnicPassThroughMode = "enabled"
)

type WatchdogAction string

const (
	WATCHDOGACTION_DUMP     WatchdogAction = "dump"
	WATCHDOGACTION_NONE     WatchdogAction = "none"
	WATCHDOGACTION_PAUSE    WatchdogAction = "pause"
	WATCHDOGACTION_POWEROFF WatchdogAction = "poweroff"
	WATCHDOGACTION_RESET    WatchdogAction = "reset"
)

type WatchdogModel string

const (
	WATCHDOGMODEL_I6300ESB WatchdogModel = "i6300esb"
)

func (fault *Fault) Error() string {
	return fmt.Sprintf("Error details is %s, reason is %s", fault.Detail, fault.Reason)
}
