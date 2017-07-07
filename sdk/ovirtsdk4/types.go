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

type FopStatistics struct {
	XMLName       xml.Name       `xml:"fopstatistics"`
	FopStatistics []FopStatistic `xml:"fop_statistic,omitempty"`
}

type FopStatistic struct {
	OvStruct
	Name       *string     `xml:"name,omitempty"`
	Statistics []Statistic `xml:"statistics,omitempty"`
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

type GracePeriods struct {
	XMLName      xml.Name      `xml:"graceperiods"`
	GracePeriods []GracePeriod `xml:"grace_period,omitempty"`
}

type GracePeriod struct {
	OvStruct
	Expiry *int64 `xml:"expiry,omitempty"`
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

type HostDevicePassthroughs struct {
	XMLName                xml.Name                `xml:"hostdevicepassthroughs"`
	HostDevicePassthroughs []HostDevicePassthrough `xml:"host_device_passthrough,omitempty"`
}

type HostDevicePassthrough struct {
	OvStruct
	Enabled *bool `xml:"enabled,omitempty"`
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
	Build       *int64  `xml:"build,omitempty"`
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

type BrickProfileDetails struct {
	XMLName             xml.Name             `xml:"brickprofiledetails"`
	BrickProfileDetails []BrickProfileDetail `xml:"brick_profile_detail,omitempty"`
}

type BrickProfileDetail struct {
	OvStruct
	Brick          *GlusterBrick   `xml:"brick,omitempty"`
	ProfileDetails []ProfileDetail `xml:"profile_details,omitempty"`
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
