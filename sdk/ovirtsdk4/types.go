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
	"time"
)

type AffinityRules struct {
	AffinityRules []AffinityRule `xml:"affinity_rule"`
}

type AffinityRule struct {
	OvStruct
	Enabled   bool `xml:"enabled"`
	Enforcing bool `xml:"enforcing"`
	Positive  bool `xml:"positive"`
}

type AgentConfigurations struct {
	AgentConfigurations []AgentConfiguration `xml:"agent_configuration"`
}

type AgentConfiguration struct {
	OvStruct
	Address         string            `xml:"address"`
	BrokerType      MessageBrokerType `xml:"broker_type"`
	NetworkMappings string            `xml:"network_mappings"`
	Password        string            `xml:"password"`
	Port            int64             `xml:"port"`
	Username        string            `xml:"username"`
}

type Apis struct {
	Apis []Api `xml:"api"`
}

type Api struct {
	OvStruct
	ProductInfo    *ProductInfo    `xml:"product_info"`
	SpecialObjects *SpecialObjects `xml:"special_objects"`
	Summary        *ApiSummary     `xml:"summary"`
	Time           time.Time       `xml:"time"`
}

type ApiSummarys struct {
	ApiSummarys []ApiSummary `xml:"api_summary"`
}

type ApiSummary struct {
	OvStruct
	Hosts          *ApiSummaryItem `xml:"hosts"`
	StorageDomains *ApiSummaryItem `xml:"storage_domains"`
	Users          *ApiSummaryItem `xml:"users"`
	Vms            *ApiSummaryItem `xml:"vms"`
}

type ApiSummaryItems struct {
	ApiSummaryItems []ApiSummaryItem `xml:"api_summary_item"`
}

type ApiSummaryItem struct {
	OvStruct
	Active int64 `xml:"active"`
	Total  int64 `xml:"total"`
}

type Bioss struct {
	Bioss []Bios `xml:"bios"`
}

type Bios struct {
	OvStruct
	BootMenu *BootMenu `xml:"boot_menu"`
}

type BlockStatistics struct {
	BlockStatistics []BlockStatistic `xml:"block_statistic"`
}

type BlockStatistic struct {
	OvStruct
	Statistics []Statistic `xml:"statistics"`
}

type Bondings struct {
	Bondings []Bonding `xml:"bonding"`
}

type Bonding struct {
	OvStruct
	ActiveSlave  *HostNic  `xml:"active_slave"`
	AdPartnerMac *Mac      `xml:"ad_partner_mac"`
	Options      []Option  `xml:"options"`
	Slaves       []HostNic `xml:"slaves"`
}

type Boots struct {
	Boots []Boot `xml:"boot"`
}

type Boot struct {
	OvStruct
	Devices []BootDevice `xml:"devices"`
}

type BootMenus struct {
	BootMenus []BootMenu `xml:"boot_menu"`
}

type BootMenu struct {
	OvStruct
	Enabled bool `xml:"enabled"`
}

type CloudInits struct {
	CloudInits []CloudInit `xml:"cloud_init"`
}

type CloudInit struct {
	OvStruct
	AuthorizedKeys       []AuthorizedKey       `xml:"authorized_keys"`
	Files                []File                `xml:"files"`
	Host                 *Host                 `xml:"host"`
	NetworkConfiguration *NetworkConfiguration `xml:"network_configuration"`
	RegenerateSshKeys    bool                  `xml:"regenerate_ssh_keys"`
	Timezone             string                `xml:"timezone"`
	Users                []User                `xml:"users"`
}

type Configurations struct {
	Configurations []Configuration `xml:"configuration"`
}

type Configuration struct {
	OvStruct
	Data string            `xml:"data"`
	Type ConfigurationType `xml:"type"`
}

type Consoles struct {
	Consoles []Console `xml:"console"`
}

type Console struct {
	OvStruct
	Enabled bool `xml:"enabled"`
}

type Cores struct {
	Cores []Core `xml:"core"`
}

type Core struct {
	OvStruct
	Index  int64 `xml:"index"`
	Socket int64 `xml:"socket"`
}

type Cpus struct {
	Cpus []Cpu `xml:"cpu"`
}

type Cpu struct {
	OvStruct
	Architecture Architecture `xml:"architecture"`
	Cores        []Core       `xml:"cores"`
	CpuTune      *CpuTune     `xml:"cpu_tune"`
	Level        int64        `xml:"level"`
	Mode         CpuMode      `xml:"mode"`
	Name         string       `xml:"name"`
	Speed        float64      `xml:"speed"`
	Topology     *CpuTopology `xml:"topology"`
	Type         string       `xml:"type"`
}

type CpuTopologys struct {
	CpuTopologys []CpuTopology `xml:"cpu_topology"`
}

type CpuTopology struct {
	OvStruct
	Cores   int64 `xml:"cores"`
	Sockets int64 `xml:"sockets"`
	Threads int64 `xml:"threads"`
}

type CpuTunes struct {
	CpuTunes []CpuTune `xml:"cpu_tune"`
}

type CpuTune struct {
	OvStruct
	VcpuPins []VcpuPin `xml:"vcpu_pins"`
}

type CpuTypes struct {
	CpuTypes []CpuType `xml:"cpu_type"`
}

type CpuType struct {
	OvStruct
	Architecture Architecture `xml:"architecture"`
	Level        int64        `xml:"level"`
	Name         string       `xml:"name"`
}

type CustomPropertys struct {
	CustomPropertys []CustomProperty `xml:"custom_property"`
}

type CustomProperty struct {
	OvStruct
	Name   string `xml:"name"`
	Regexp string `xml:"regexp"`
	Value  string `xml:"value"`
}

type Displays struct {
	Displays []Display `xml:"display"`
}

type Display struct {
	OvStruct
	Address             string       `xml:"address"`
	AllowOverride       bool         `xml:"allow_override"`
	Certificate         *Certificate `xml:"certificate"`
	CopyPasteEnabled    bool         `xml:"copy_paste_enabled"`
	DisconnectAction    string       `xml:"disconnect_action"`
	FileTransferEnabled bool         `xml:"file_transfer_enabled"`
	KeyboardLayout      string       `xml:"keyboard_layout"`
	Monitors            int64        `xml:"monitors"`
	Port                int64        `xml:"port"`
	Proxy               string       `xml:"proxy"`
	SecurePort          int64        `xml:"secure_port"`
	SingleQxlPci        bool         `xml:"single_qxl_pci"`
	SmartcardEnabled    bool         `xml:"smartcard_enabled"`
	Type                DisplayType  `xml:"type"`
}

type Dnss struct {
	Dnss []Dns `xml:"dns"`
}

type Dns struct {
	OvStruct
	SearchDomains []Host `xml:"search_domains"`
	Servers       []Host `xml:"servers"`
}

type DnsResolverConfigurations struct {
	DnsResolverConfigurations []DnsResolverConfiguration `xml:"dns_resolver_configuration"`
}

type DnsResolverConfiguration struct {
	OvStruct
	NameServers []string `xml:"name_servers"`
}

type EntityProfileDetails struct {
	EntityProfileDetails []EntityProfileDetail `xml:"entity_profile_detail"`
}

type EntityProfileDetail struct {
	OvStruct
	ProfileDetails []ProfileDetail `xml:"profile_details"`
}

type ErrorHandlings struct {
	ErrorHandlings []ErrorHandling `xml:"error_handling"`
}

type ErrorHandling struct {
	OvStruct
	OnError MigrateOnError `xml:"on_error"`
}

type ExternalVmImports struct {
	ExternalVmImports []ExternalVmImport `xml:"external_vm_import"`
}

type ExternalVmImport struct {
	OvStruct
	Cluster       *Cluster               `xml:"cluster"`
	CpuProfile    *CpuProfile            `xml:"cpu_profile"`
	DriversIso    *File                  `xml:"drivers_iso"`
	Host          *Host                  `xml:"host"`
	Name          string                 `xml:"name"`
	Password      string                 `xml:"password"`
	Provider      ExternalVmProviderType `xml:"provider"`
	Quota         *Quota                 `xml:"quota"`
	Sparse        bool                   `xml:"sparse"`
	StorageDomain *StorageDomain         `xml:"storage_domain"`
	Url           string                 `xml:"url"`
	Username      string                 `xml:"username"`
	Vm            *Vm                    `xml:"vm"`
}

type Faults struct {
	Faults []Fault `xml:"fault"`
}

type Fault struct {
	OvStruct
	Detail string `xml:"detail"`
	Reason string `xml:"reason"`
}

type FencingPolicys struct {
	FencingPolicys []FencingPolicy `xml:"fencing_policy"`
}

type FencingPolicy struct {
	OvStruct
	Enabled                   bool                      `xml:"enabled"`
	SkipIfConnectivityBroken  *SkipIfConnectivityBroken `xml:"skip_if_connectivity_broken"`
	SkipIfGlusterBricksUp     bool                      `xml:"skip_if_gluster_bricks_up"`
	SkipIfGlusterQuorumNotMet bool                      `xml:"skip_if_gluster_quorum_not_met"`
	SkipIfSdActive            *SkipIfSdActive           `xml:"skip_if_sd_active"`
}

type FopStatistics struct {
	FopStatistics []FopStatistic `xml:"fop_statistic"`
}

type FopStatistic struct {
	OvStruct
	Name       string      `xml:"name"`
	Statistics []Statistic `xml:"statistics"`
}

type GlusterBrickMemoryInfos struct {
	GlusterBrickMemoryInfos []GlusterBrickMemoryInfo `xml:"gluster_brick_memory_info"`
}

type GlusterBrickMemoryInfo struct {
	OvStruct
	MemoryPools []GlusterMemoryPool `xml:"memory_pools"`
}

type GlusterClients struct {
	GlusterClients []GlusterClient `xml:"gluster_client"`
}

type GlusterClient struct {
	OvStruct
	BytesRead    int64  `xml:"bytes_read"`
	BytesWritten int64  `xml:"bytes_written"`
	ClientPort   int64  `xml:"client_port"`
	HostName     string `xml:"host_name"`
}

type GracePeriods struct {
	GracePeriods []GracePeriod `xml:"grace_period"`
}

type GracePeriod struct {
	OvStruct
	Expiry int64 `xml:"expiry"`
}

type GuestOperatingSystems struct {
	GuestOperatingSystems []GuestOperatingSystem `xml:"guest_operating_system"`
}

type GuestOperatingSystem struct {
	OvStruct
	Architecture string   `xml:"architecture"`
	Codename     string   `xml:"codename"`
	Distribution string   `xml:"distribution"`
	Family       string   `xml:"family"`
	Kernel       *Kernel  `xml:"kernel"`
	Version      *Version `xml:"version"`
}

type HardwareInformations struct {
	HardwareInformations []HardwareInformation `xml:"hardware_information"`
}

type HardwareInformation struct {
	OvStruct
	Family              string      `xml:"family"`
	Manufacturer        string      `xml:"manufacturer"`
	ProductName         string      `xml:"product_name"`
	SerialNumber        string      `xml:"serial_number"`
	SupportedRngSources []RngSource `xml:"supported_rng_sources"`
	Uuid                string      `xml:"uuid"`
	Version             string      `xml:"version"`
}

type HighAvailabilitys struct {
	HighAvailabilitys []HighAvailability `xml:"high_availability"`
}

type HighAvailability struct {
	OvStruct
	Enabled  bool  `xml:"enabled"`
	Priority int64 `xml:"priority"`
}

type HostDevicePassthroughs struct {
	HostDevicePassthroughs []HostDevicePassthrough `xml:"host_device_passthrough"`
}

type HostDevicePassthrough struct {
	OvStruct
	Enabled bool `xml:"enabled"`
}

type HostNicVirtualFunctionsConfigurations struct {
	HostNicVirtualFunctionsConfigurations []HostNicVirtualFunctionsConfiguration `xml:"host_nic_virtual_functions_configuration"`
}

type HostNicVirtualFunctionsConfiguration struct {
	OvStruct
	AllNetworksAllowed          bool  `xml:"all_networks_allowed"`
	MaxNumberOfVirtualFunctions int64 `xml:"max_number_of_virtual_functions"`
	NumberOfVirtualFunctions    int64 `xml:"number_of_virtual_functions"`
}

type HostedEngines struct {
	HostedEngines []HostedEngine `xml:"hosted_engine"`
}

type HostedEngine struct {
	OvStruct
	Active            bool  `xml:"active"`
	Configured        bool  `xml:"configured"`
	GlobalMaintenance bool  `xml:"global_maintenance"`
	LocalMaintenance  bool  `xml:"local_maintenance"`
	Score             int64 `xml:"score"`
}

type Identifieds struct {
	Identifieds []Identified `xml:"identified"`
}

type Identified struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
}

type Images struct {
	Images []Image `xml:"image"`
}

type Image struct {
	OvStruct
	Comment       string         `xml:"comment"`
	Description   string         `xml:"description"`
	Id            string         `xml:"id,attr"`
	Name          string         `xml:"name"`
	StorageDomain *StorageDomain `xml:"storage_domain"`
}

type ImageTransfers struct {
	ImageTransfers []ImageTransfer `xml:"image_transfer"`
}

type ImageTransfer struct {
	OvStruct
	Comment      string                 `xml:"comment"`
	Description  string                 `xml:"description"`
	Direction    ImageTransferDirection `xml:"direction"`
	Host         *Host                  `xml:"host"`
	Id           string                 `xml:"id,attr"`
	Image        *Image                 `xml:"image"`
	Name         string                 `xml:"name"`
	Phase        ImageTransferPhase     `xml:"phase"`
	ProxyUrl     string                 `xml:"proxy_url"`
	SignedTicket string                 `xml:"signed_ticket"`
}

type Initializations struct {
	Initializations []Initialization `xml:"initialization"`
}

type Initialization struct {
	OvStruct
	ActiveDirectoryOu string             `xml:"active_directory_ou"`
	AuthorizedSshKeys string             `xml:"authorized_ssh_keys"`
	CloudInit         *CloudInit         `xml:"cloud_init"`
	Configuration     *Configuration     `xml:"configuration"`
	CustomScript      string             `xml:"custom_script"`
	DnsSearch         string             `xml:"dns_search"`
	DnsServers        string             `xml:"dns_servers"`
	Domain            string             `xml:"domain"`
	HostName          string             `xml:"host_name"`
	InputLocale       string             `xml:"input_locale"`
	NicConfigurations []NicConfiguration `xml:"nic_configurations"`
	OrgName           string             `xml:"org_name"`
	RegenerateIds     bool               `xml:"regenerate_ids"`
	RegenerateSshKeys bool               `xml:"regenerate_ssh_keys"`
	RootPassword      string             `xml:"root_password"`
	SystemLocale      string             `xml:"system_locale"`
	Timezone          string             `xml:"timezone"`
	UiLanguage        string             `xml:"ui_language"`
	UserLocale        string             `xml:"user_locale"`
	UserName          string             `xml:"user_name"`
	WindowsLicenseKey string             `xml:"windows_license_key"`
}

type Ios struct {
	Ios []Io `xml:"io"`
}

type Io struct {
	OvStruct
	Threads int64 `xml:"threads"`
}

type Ips struct {
	Ips []Ip `xml:"ip"`
}

type Ip struct {
	OvStruct
	Address string    `xml:"address"`
	Gateway string    `xml:"gateway"`
	Netmask string    `xml:"netmask"`
	Version IpVersion `xml:"version"`
}

type IpAddressAssignments struct {
	IpAddressAssignments []IpAddressAssignment `xml:"ip_address_assignment"`
}

type IpAddressAssignment struct {
	OvStruct
	AssignmentMethod BootProtocol `xml:"assignment_method"`
	Ip               *Ip          `xml:"ip"`
}

type IscsiBonds struct {
	IscsiBonds []IscsiBond `xml:"iscsi_bond"`
}

type IscsiBond struct {
	OvStruct
	Comment            string              `xml:"comment"`
	DataCenter         *DataCenter         `xml:"data_center"`
	Description        string              `xml:"description"`
	Id                 string              `xml:"id,attr"`
	Name               string              `xml:"name"`
	Networks           []Network           `xml:"networks"`
	StorageConnections []StorageConnection `xml:"storage_connections"`
}

type IscsiDetailss struct {
	IscsiDetailss []IscsiDetails `xml:"iscsi_details"`
}

type IscsiDetails struct {
	OvStruct
	Address         string `xml:"address"`
	DiskId          string `xml:"disk_id"`
	Initiator       string `xml:"initiator"`
	LunMapping      int64  `xml:"lun_mapping"`
	Password        string `xml:"password"`
	Paths           int64  `xml:"paths"`
	Port            int64  `xml:"port"`
	Portal          string `xml:"portal"`
	ProductId       string `xml:"product_id"`
	Serial          string `xml:"serial"`
	Size            int64  `xml:"size"`
	Status          string `xml:"status"`
	StorageDomainId string `xml:"storage_domain_id"`
	Target          string `xml:"target"`
	Username        string `xml:"username"`
	VendorId        string `xml:"vendor_id"`
	VolumeGroupId   string `xml:"volume_group_id"`
}

type Jobs struct {
	Jobs []Job `xml:"job"`
}

type Job struct {
	OvStruct
	AutoCleared bool      `xml:"auto_cleared"`
	Comment     string    `xml:"comment"`
	Description string    `xml:"description"`
	EndTime     time.Time `xml:"end_time"`
	External    bool      `xml:"external"`
	Id          string    `xml:"id,attr"`
	LastUpdated time.Time `xml:"last_updated"`
	Name        string    `xml:"name"`
	Owner       *User     `xml:"owner"`
	StartTime   time.Time `xml:"start_time"`
	Status      JobStatus `xml:"status"`
	Steps       []Step    `xml:"steps"`
}

type KatelloErratums struct {
	KatelloErratums []KatelloErratum `xml:"katello_erratum"`
}

type KatelloErratum struct {
	OvStruct
	Comment     string    `xml:"comment"`
	Description string    `xml:"description"`
	Host        *Host     `xml:"host"`
	Id          string    `xml:"id,attr"`
	Issued      time.Time `xml:"issued"`
	Name        string    `xml:"name"`
	Packages    []Package `xml:"packages"`
	Severity    string    `xml:"severity"`
	Solution    string    `xml:"solution"`
	Summary     string    `xml:"summary"`
	Title       string    `xml:"title"`
	Type        string    `xml:"type"`
	Vm          *Vm       `xml:"vm"`
}

type Kernels struct {
	Kernels []Kernel `xml:"kernel"`
}

type Kernel struct {
	OvStruct
	Version *Version `xml:"version"`
}

type Ksms struct {
	Ksms []Ksm `xml:"ksm"`
}

type Ksm struct {
	OvStruct
	Enabled          bool `xml:"enabled"`
	MergeAcrossNodes bool `xml:"merge_across_nodes"`
}

type LogicalUnits struct {
	LogicalUnits []LogicalUnit `xml:"logical_unit"`
}

type LogicalUnit struct {
	OvStruct
	Address           string    `xml:"address"`
	DiscardMaxSize    int64     `xml:"discard_max_size"`
	DiscardZeroesData bool      `xml:"discard_zeroes_data"`
	DiskId            string    `xml:"disk_id"`
	Id                string    `xml:"id,attr"`
	LunMapping        int64     `xml:"lun_mapping"`
	Password          string    `xml:"password"`
	Paths             int64     `xml:"paths"`
	Port              int64     `xml:"port"`
	Portal            string    `xml:"portal"`
	ProductId         string    `xml:"product_id"`
	Serial            string    `xml:"serial"`
	Size              int64     `xml:"size"`
	Status            LunStatus `xml:"status"`
	StorageDomainId   string    `xml:"storage_domain_id"`
	Target            string    `xml:"target"`
	Username          string    `xml:"username"`
	VendorId          string    `xml:"vendor_id"`
	VolumeGroupId     string    `xml:"volume_group_id"`
}

type Macs struct {
	Macs []Mac `xml:"mac"`
}

type Mac struct {
	OvStruct
	Address string `xml:"address"`
}

type MacPools struct {
	MacPools []MacPool `xml:"mac_pool"`
}

type MacPool struct {
	OvStruct
	AllowDuplicates bool    `xml:"allow_duplicates"`
	Comment         string  `xml:"comment"`
	DefaultPool     bool    `xml:"default_pool"`
	Description     string  `xml:"description"`
	Id              string  `xml:"id,attr"`
	Name            string  `xml:"name"`
	Ranges          []Range `xml:"ranges"`
}

type MemoryOverCommits struct {
	MemoryOverCommits []MemoryOverCommit `xml:"memory_over_commit"`
}

type MemoryOverCommit struct {
	OvStruct
	Percent int64 `xml:"percent"`
}

type MemoryPolicys struct {
	MemoryPolicys []MemoryPolicy `xml:"memory_policy"`
}

type MemoryPolicy struct {
	OvStruct
	Ballooning           bool                  `xml:"ballooning"`
	Guaranteed           int64                 `xml:"guaranteed"`
	Max                  int64                 `xml:"max"`
	OverCommit           *MemoryOverCommit     `xml:"over_commit"`
	TransparentHugePages *TransparentHugePages `xml:"transparent_huge_pages"`
}

type Methods struct {
	Methods []Method `xml:"method"`
}

type Method struct {
	OvStruct
	Id SsoMethod `xml:"id,attr"`
}

type MigrationBandwidths struct {
	MigrationBandwidths []MigrationBandwidth `xml:"migration_bandwidth"`
}

type MigrationBandwidth struct {
	OvStruct
	AssignmentMethod MigrationBandwidthAssignmentMethod `xml:"assignment_method"`
	CustomValue      int64                              `xml:"custom_value"`
}

type MigrationOptionss struct {
	MigrationOptionss []MigrationOptions `xml:"migration_options"`
}

type MigrationOptions struct {
	OvStruct
	AutoConverge InheritableBoolean  `xml:"auto_converge"`
	Bandwidth    *MigrationBandwidth `xml:"bandwidth"`
	Compressed   InheritableBoolean  `xml:"compressed"`
	Policy       *MigrationPolicy    `xml:"policy"`
}

type MigrationPolicys struct {
	MigrationPolicys []MigrationPolicy `xml:"migration_policy"`
}

type MigrationPolicy struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
}

type Networks struct {
	Networks []Network `xml:"network"`
}

type Network struct {
	OvStruct
	Cluster                  *Cluster                  `xml:"cluster"`
	Comment                  string                    `xml:"comment"`
	DataCenter               *DataCenter               `xml:"data_center"`
	Description              string                    `xml:"description"`
	Display                  bool                      `xml:"display"`
	DnsResolverConfiguration *DnsResolverConfiguration `xml:"dns_resolver_configuration"`
	Id                       string                    `xml:"id,attr"`
	Ip                       *Ip                       `xml:"ip"`
	Mtu                      int64                     `xml:"mtu"`
	Name                     string                    `xml:"name"`
	NetworkLabels            []NetworkLabel            `xml:"network_labels"`
	Permissions              []Permission              `xml:"permissions"`
	ProfileRequired          bool                      `xml:"profile_required"`
	Qos                      *Qos                      `xml:"qos"`
	Required                 bool                      `xml:"required"`
	Status                   NetworkStatus             `xml:"status"`
	Stp                      bool                      `xml:"stp"`
	Usages                   []NetworkUsage            `xml:"usages"`
	Vlan                     *Vlan                     `xml:"vlan"`
	VnicProfiles             []VnicProfile             `xml:"vnic_profiles"`
}

type NetworkAttachments struct {
	NetworkAttachments []NetworkAttachment `xml:"network_attachment"`
}

type NetworkAttachment struct {
	OvStruct
	Comment                  string                    `xml:"comment"`
	Description              string                    `xml:"description"`
	DnsResolverConfiguration *DnsResolverConfiguration `xml:"dns_resolver_configuration"`
	Host                     *Host                     `xml:"host"`
	HostNic                  *HostNic                  `xml:"host_nic"`
	Id                       string                    `xml:"id,attr"`
	InSync                   bool                      `xml:"in_sync"`
	IpAddressAssignments     []IpAddressAssignment     `xml:"ip_address_assignments"`
	Name                     string                    `xml:"name"`
	Network                  *Network                  `xml:"network"`
	Properties               []Property                `xml:"properties"`
	Qos                      *Qos                      `xml:"qos"`
	ReportedConfigurations   []ReportedConfiguration   `xml:"reported_configurations"`
}

type NetworkConfigurations struct {
	NetworkConfigurations []NetworkConfiguration `xml:"network_configuration"`
}

type NetworkConfiguration struct {
	OvStruct
	Dns  *Dns  `xml:"dns"`
	Nics []Nic `xml:"nics"`
}

type NetworkFilters struct {
	NetworkFilters []NetworkFilter `xml:"network_filter"`
}

type NetworkFilter struct {
	OvStruct
	Comment     string   `xml:"comment"`
	Description string   `xml:"description"`
	Id          string   `xml:"id,attr"`
	Name        string   `xml:"name"`
	Version     *Version `xml:"version"`
}

type NetworkFilterParameters struct {
	NetworkFilterParameters []NetworkFilterParameter `xml:"network_filter_parameter"`
}

type NetworkFilterParameter struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Value       string `xml:"value"`
}

type NetworkLabels struct {
	NetworkLabels []NetworkLabel `xml:"network_label"`
}

type NetworkLabel struct {
	OvStruct
	Comment     string   `xml:"comment"`
	Description string   `xml:"description"`
	HostNic     *HostNic `xml:"host_nic"`
	Id          string   `xml:"id,attr"`
	Name        string   `xml:"name"`
	Network     *Network `xml:"network"`
}

type NfsProfileDetails struct {
	NfsProfileDetails []NfsProfileDetail `xml:"nfs_profile_detail"`
}

type NfsProfileDetail struct {
	OvStruct
	NfsServerIp    string          `xml:"nfs_server_ip"`
	ProfileDetails []ProfileDetail `xml:"profile_details"`
}

type NicConfigurations struct {
	NicConfigurations []NicConfiguration `xml:"nic_configuration"`
}

type NicConfiguration struct {
	OvStruct
	BootProtocol     BootProtocol `xml:"boot_protocol"`
	Ip               *Ip          `xml:"ip"`
	Ipv6             *Ip          `xml:"ipv6"`
	Ipv6BootProtocol BootProtocol `xml:"ipv6_boot_protocol"`
	Name             string       `xml:"name"`
	OnBoot           bool         `xml:"on_boot"`
}

type NumaNodes struct {
	NumaNodes []NumaNode `xml:"numa_node"`
}

type NumaNode struct {
	OvStruct
	Comment      string      `xml:"comment"`
	Cpu          *Cpu        `xml:"cpu"`
	Description  string      `xml:"description"`
	Host         *Host       `xml:"host"`
	Id           string      `xml:"id,attr"`
	Index        int64       `xml:"index"`
	Memory       int64       `xml:"memory"`
	Name         string      `xml:"name"`
	NodeDistance string      `xml:"node_distance"`
	Statistics   []Statistic `xml:"statistics"`
}

type NumaNodePins struct {
	NumaNodePins []NumaNodePin `xml:"numa_node_pin"`
}

type NumaNodePin struct {
	OvStruct
	HostNumaNode *NumaNode `xml:"host_numa_node"`
	Index        int64     `xml:"index"`
	Pinned       bool      `xml:"pinned"`
}

type OpenStackImages struct {
	OpenStackImages []OpenStackImage `xml:"open_stack_image"`
}

type OpenStackImage struct {
	OvStruct
	Comment                string                  `xml:"comment"`
	Description            string                  `xml:"description"`
	Id                     string                  `xml:"id,attr"`
	Name                   string                  `xml:"name"`
	OpenstackImageProvider *OpenStackImageProvider `xml:"openstack_image_provider"`
}

type OpenStackNetworks struct {
	OpenStackNetworks []OpenStackNetwork `xml:"open_stack_network"`
}

type OpenStackNetwork struct {
	OvStruct
	Comment                  string                    `xml:"comment"`
	Description              string                    `xml:"description"`
	Id                       string                    `xml:"id,attr"`
	Name                     string                    `xml:"name"`
	OpenstackNetworkProvider *OpenStackNetworkProvider `xml:"openstack_network_provider"`
}

type OpenStackSubnets struct {
	OpenStackSubnets []OpenStackSubnet `xml:"open_stack_subnet"`
}

type OpenStackSubnet struct {
	OvStruct
	Cidr             string            `xml:"cidr"`
	Comment          string            `xml:"comment"`
	Description      string            `xml:"description"`
	DnsServers       []string          `xml:"dns_servers"`
	Gateway          string            `xml:"gateway"`
	Id               string            `xml:"id,attr"`
	IpVersion        string            `xml:"ip_version"`
	Name             string            `xml:"name"`
	OpenstackNetwork *OpenStackNetwork `xml:"openstack_network"`
}

type OpenStackVolumeTypes struct {
	OpenStackVolumeTypes []OpenStackVolumeType `xml:"open_stack_volume_type"`
}

type OpenStackVolumeType struct {
	OvStruct
	Comment                 string                   `xml:"comment"`
	Description             string                   `xml:"description"`
	Id                      string                   `xml:"id,attr"`
	Name                    string                   `xml:"name"`
	OpenstackVolumeProvider *OpenStackVolumeProvider `xml:"openstack_volume_provider"`
	Properties              []Property               `xml:"properties"`
}

type OpenstackVolumeAuthenticationKeys struct {
	OpenstackVolumeAuthenticationKeys []OpenstackVolumeAuthenticationKey `xml:"openstack_volume_authentication_key"`
}

type OpenstackVolumeAuthenticationKey struct {
	OvStruct
	Comment                 string                                    `xml:"comment"`
	CreationDate            time.Time                                 `xml:"creation_date"`
	Description             string                                    `xml:"description"`
	Id                      string                                    `xml:"id,attr"`
	Name                    string                                    `xml:"name"`
	OpenstackVolumeProvider *OpenStackVolumeProvider                  `xml:"openstack_volume_provider"`
	UsageType               OpenstackVolumeAuthenticationKeyUsageType `xml:"usage_type"`
	Uuid                    string                                    `xml:"uuid"`
	Value                   string                                    `xml:"value"`
}

type OperatingSystems struct {
	OperatingSystems []OperatingSystem `xml:"operating_system"`
}

type OperatingSystem struct {
	OvStruct
	Boot                  *Boot    `xml:"boot"`
	Cmdline               string   `xml:"cmdline"`
	CustomKernelCmdline   string   `xml:"custom_kernel_cmdline"`
	Initrd                string   `xml:"initrd"`
	Kernel                string   `xml:"kernel"`
	ReportedKernelCmdline string   `xml:"reported_kernel_cmdline"`
	Type                  string   `xml:"type"`
	Version               *Version `xml:"version"`
}

type OperatingSystemInfos struct {
	OperatingSystemInfos []OperatingSystemInfo `xml:"operating_system_info"`
}

type OperatingSystemInfo struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	LargeIcon   *Icon  `xml:"large_icon"`
	Name        string `xml:"name"`
	SmallIcon   *Icon  `xml:"small_icon"`
}

type Options struct {
	Options []Option `xml:"option"`
}

type Option struct {
	OvStruct
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

type Packages struct {
	Packages []Package `xml:"package"`
}

type Package struct {
	OvStruct
	Name string `xml:"name"`
}

type Payloads struct {
	Payloads []Payload `xml:"payload"`
}

type Payload struct {
	OvStruct
	Files    []File       `xml:"files"`
	Type     VmDeviceType `xml:"type"`
	VolumeId string       `xml:"volume_id"`
}

type Permissions struct {
	Permissions []Permission `xml:"permission"`
}

type Permission struct {
	OvStruct
	Cluster       *Cluster       `xml:"cluster"`
	Comment       string         `xml:"comment"`
	DataCenter    *DataCenter    `xml:"data_center"`
	Description   string         `xml:"description"`
	Disk          *Disk          `xml:"disk"`
	Group         *Group         `xml:"group"`
	Host          *Host          `xml:"host"`
	Id            string         `xml:"id,attr"`
	Name          string         `xml:"name"`
	Role          *Role          `xml:"role"`
	StorageDomain *StorageDomain `xml:"storage_domain"`
	Template      *Template      `xml:"template"`
	User          *User          `xml:"user"`
	Vm            *Vm            `xml:"vm"`
	VmPool        *VmPool        `xml:"vm_pool"`
}

type Permits struct {
	Permits []Permit `xml:"permit"`
}

type Permit struct {
	OvStruct
	Administrative bool   `xml:"administrative"`
	Comment        string `xml:"comment"`
	Description    string `xml:"description"`
	Id             string `xml:"id,attr"`
	Name           string `xml:"name"`
	Role           *Role  `xml:"role"`
}

type PmProxys struct {
	PmProxys []PmProxy `xml:"pm_proxy"`
}

type PmProxy struct {
	OvStruct
	Type PmProxyType `xml:"type"`
}

type PortMirrorings struct {
	PortMirrorings []PortMirroring `xml:"port_mirroring"`
}

type PortMirroring struct {
	OvStruct
}

type PowerManagements struct {
	PowerManagements []PowerManagement `xml:"power_management"`
}

type PowerManagement struct {
	OvStruct
	Address            string                `xml:"address"`
	Agents             []Agent               `xml:"agents"`
	AutomaticPmEnabled bool                  `xml:"automatic_pm_enabled"`
	Enabled            bool                  `xml:"enabled"`
	KdumpDetection     bool                  `xml:"kdump_detection"`
	Options            []Option              `xml:"options"`
	Password           string                `xml:"password"`
	PmProxies          []PmProxy             `xml:"pm_proxies"`
	Status             PowerManagementStatus `xml:"status"`
	Type               string                `xml:"type"`
	Username           string                `xml:"username"`
}

type Products struct {
	Products []Product `xml:"product"`
}

type Product struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
}

type ProductInfos struct {
	ProductInfos []ProductInfo `xml:"product_info"`
}

type ProductInfo struct {
	OvStruct
	Name    string   `xml:"name"`
	Vendor  string   `xml:"vendor"`
	Version *Version `xml:"version"`
}

type ProfileDetails struct {
	ProfileDetails []ProfileDetail `xml:"profile_detail"`
}

type ProfileDetail struct {
	OvStruct
	BlockStatistics []BlockStatistic `xml:"block_statistics"`
	Duration        int64            `xml:"duration"`
	FopStatistics   []FopStatistic   `xml:"fop_statistics"`
	ProfileType     string           `xml:"profile_type"`
	Statistics      []Statistic      `xml:"statistics"`
}

type Propertys struct {
	Propertys []Property `xml:"property"`
}

type Property struct {
	OvStruct
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type ProxyTickets struct {
	ProxyTickets []ProxyTicket `xml:"proxy_ticket"`
}

type ProxyTicket struct {
	OvStruct
	Value string `xml:"value"`
}

type Qoss struct {
	Qoss []Qos `xml:"qos"`
}

type Qos struct {
	OvStruct
	Comment                   string      `xml:"comment"`
	CpuLimit                  int64       `xml:"cpu_limit"`
	DataCenter                *DataCenter `xml:"data_center"`
	Description               string      `xml:"description"`
	Id                        string      `xml:"id,attr"`
	InboundAverage            int64       `xml:"inbound_average"`
	InboundBurst              int64       `xml:"inbound_burst"`
	InboundPeak               int64       `xml:"inbound_peak"`
	MaxIops                   int64       `xml:"max_iops"`
	MaxReadIops               int64       `xml:"max_read_iops"`
	MaxReadThroughput         int64       `xml:"max_read_throughput"`
	MaxThroughput             int64       `xml:"max_throughput"`
	MaxWriteIops              int64       `xml:"max_write_iops"`
	MaxWriteThroughput        int64       `xml:"max_write_throughput"`
	Name                      string      `xml:"name"`
	OutboundAverage           int64       `xml:"outbound_average"`
	OutboundAverageLinkshare  int64       `xml:"outbound_average_linkshare"`
	OutboundAverageRealtime   int64       `xml:"outbound_average_realtime"`
	OutboundAverageUpperlimit int64       `xml:"outbound_average_upperlimit"`
	OutboundBurst             int64       `xml:"outbound_burst"`
	OutboundPeak              int64       `xml:"outbound_peak"`
	Type                      QosType     `xml:"type"`
}

type Quotas struct {
	Quotas []Quota `xml:"quota"`
}

type Quota struct {
	OvStruct
	ClusterHardLimitPct int64               `xml:"cluster_hard_limit_pct"`
	ClusterSoftLimitPct int64               `xml:"cluster_soft_limit_pct"`
	Comment             string              `xml:"comment"`
	DataCenter          *DataCenter         `xml:"data_center"`
	Description         string              `xml:"description"`
	Disks               []Disk              `xml:"disks"`
	Id                  string              `xml:"id,attr"`
	Name                string              `xml:"name"`
	Permissions         []Permission        `xml:"permissions"`
	QuotaClusterLimits  []QuotaClusterLimit `xml:"quota_cluster_limits"`
	QuotaStorageLimits  []QuotaStorageLimit `xml:"quota_storage_limits"`
	StorageHardLimitPct int64               `xml:"storage_hard_limit_pct"`
	StorageSoftLimitPct int64               `xml:"storage_soft_limit_pct"`
	Users               []User              `xml:"users"`
	Vms                 []Vm                `xml:"vms"`
}

type QuotaClusterLimits struct {
	QuotaClusterLimits []QuotaClusterLimit `xml:"quota_cluster_limit"`
}

type QuotaClusterLimit struct {
	OvStruct
	Cluster     *Cluster `xml:"cluster"`
	Comment     string   `xml:"comment"`
	Description string   `xml:"description"`
	Id          string   `xml:"id,attr"`
	MemoryLimit float64  `xml:"memory_limit"`
	MemoryUsage float64  `xml:"memory_usage"`
	Name        string   `xml:"name"`
	Quota       *Quota   `xml:"quota"`
	VcpuLimit   int64    `xml:"vcpu_limit"`
	VcpuUsage   int64    `xml:"vcpu_usage"`
}

type QuotaStorageLimits struct {
	QuotaStorageLimits []QuotaStorageLimit `xml:"quota_storage_limit"`
}

type QuotaStorageLimit struct {
	OvStruct
	Comment       string         `xml:"comment"`
	Description   string         `xml:"description"`
	Id            string         `xml:"id,attr"`
	Limit         int64          `xml:"limit"`
	Name          string         `xml:"name"`
	Quota         *Quota         `xml:"quota"`
	StorageDomain *StorageDomain `xml:"storage_domain"`
	Usage         float64        `xml:"usage"`
}

type Ranges struct {
	Ranges []Range `xml:"range"`
}

type Range struct {
	OvStruct
	From string `xml:"from"`
	To   string `xml:"to"`
}

type Rates struct {
	Rates []Rate `xml:"rate"`
}

type Rate struct {
	OvStruct
	Bytes  int64 `xml:"bytes"`
	Period int64 `xml:"period"`
}

type ReportedConfigurations struct {
	ReportedConfigurations []ReportedConfiguration `xml:"reported_configuration"`
}

type ReportedConfiguration struct {
	OvStruct
	ActualValue   string `xml:"actual_value"`
	ExpectedValue string `xml:"expected_value"`
	InSync        bool   `xml:"in_sync"`
	Name          string `xml:"name"`
}

type ReportedDevices struct {
	ReportedDevices []ReportedDevice `xml:"reported_device"`
}

type ReportedDevice struct {
	OvStruct
	Comment     string             `xml:"comment"`
	Description string             `xml:"description"`
	Id          string             `xml:"id,attr"`
	Ips         []Ip               `xml:"ips"`
	Mac         *Mac               `xml:"mac"`
	Name        string             `xml:"name"`
	Type        ReportedDeviceType `xml:"type"`
	Vm          *Vm                `xml:"vm"`
}

type RngDevices struct {
	RngDevices []RngDevice `xml:"rng_device"`
}

type RngDevice struct {
	OvStruct
	Rate   *Rate     `xml:"rate"`
	Source RngSource `xml:"source"`
}

type Roles struct {
	Roles []Role `xml:"role"`
}

type Role struct {
	OvStruct
	Administrative bool     `xml:"administrative"`
	Comment        string   `xml:"comment"`
	Description    string   `xml:"description"`
	Id             string   `xml:"id,attr"`
	Mutable        bool     `xml:"mutable"`
	Name           string   `xml:"name"`
	Permits        []Permit `xml:"permits"`
	User           *User    `xml:"user"`
}

type SchedulingPolicys struct {
	SchedulingPolicys []SchedulingPolicy `xml:"scheduling_policy"`
}

type SchedulingPolicy struct {
	OvStruct
	Balances      []Balance  `xml:"balances"`
	Comment       string     `xml:"comment"`
	DefaultPolicy bool       `xml:"default_policy"`
	Description   string     `xml:"description"`
	Filters       []Filter   `xml:"filters"`
	Id            string     `xml:"id,attr"`
	Locked        bool       `xml:"locked"`
	Name          string     `xml:"name"`
	Properties    []Property `xml:"properties"`
	Weight        []Weight   `xml:"weight"`
}

type SchedulingPolicyUnits struct {
	SchedulingPolicyUnits []SchedulingPolicyUnit `xml:"scheduling_policy_unit"`
}

type SchedulingPolicyUnit struct {
	OvStruct
	Comment     string         `xml:"comment"`
	Description string         `xml:"description"`
	Enabled     bool           `xml:"enabled"`
	Id          string         `xml:"id,attr"`
	Internal    bool           `xml:"internal"`
	Name        string         `xml:"name"`
	Properties  []Property     `xml:"properties"`
	Type        PolicyUnitType `xml:"type"`
}

type SeLinuxs struct {
	SeLinuxs []SeLinux `xml:"se_linux"`
}

type SeLinux struct {
	OvStruct
	Mode SeLinuxMode `xml:"mode"`
}

type SerialNumbers struct {
	SerialNumbers []SerialNumber `xml:"serial_number"`
}

type SerialNumber struct {
	OvStruct
	Policy SerialNumberPolicy `xml:"policy"`
	Value  string             `xml:"value"`
}

type Sessions struct {
	Sessions []Session `xml:"session"`
}

type Session struct {
	OvStruct
	Comment     string `xml:"comment"`
	ConsoleUser bool   `xml:"console_user"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Ip          *Ip    `xml:"ip"`
	Name        string `xml:"name"`
	Protocol    string `xml:"protocol"`
	User        *User  `xml:"user"`
	Vm          *Vm    `xml:"vm"`
}

type SkipIfConnectivityBrokens struct {
	SkipIfConnectivityBrokens []SkipIfConnectivityBroken `xml:"skip_if_connectivity_broken"`
}

type SkipIfConnectivityBroken struct {
	OvStruct
	Enabled   bool  `xml:"enabled"`
	Threshold int64 `xml:"threshold"`
}

type SkipIfSdActives struct {
	SkipIfSdActives []SkipIfSdActive `xml:"skip_if_sd_active"`
}

type SkipIfSdActive struct {
	OvStruct
	Enabled bool `xml:"enabled"`
}

type SpecialObjectss struct {
	SpecialObjectss []SpecialObjects `xml:"special_objects"`
}

type SpecialObjects struct {
	OvStruct
	BlankTemplate *Template `xml:"blank_template"`
	RootTag       *Tag      `xml:"root_tag"`
}

type Spms struct {
	Spms []Spm `xml:"spm"`
}

type Spm struct {
	OvStruct
	Priority int64     `xml:"priority"`
	Status   SpmStatus `xml:"status"`
}

type Sshs struct {
	Sshs []Ssh `xml:"ssh"`
}

type Ssh struct {
	OvStruct
	AuthenticationMethod SshAuthenticationMethod `xml:"authentication_method"`
	Comment              string                  `xml:"comment"`
	Description          string                  `xml:"description"`
	Fingerprint          string                  `xml:"fingerprint"`
	Id                   string                  `xml:"id,attr"`
	Name                 string                  `xml:"name"`
	Port                 int64                   `xml:"port"`
	User                 *User                   `xml:"user"`
}

type SshPublicKeys struct {
	SshPublicKeys []SshPublicKey `xml:"ssh_public_key"`
}

type SshPublicKey struct {
	OvStruct
	Comment     string `xml:"comment"`
	Content     string `xml:"content"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
	User        *User  `xml:"user"`
}

type Ssos struct {
	Ssos []Sso `xml:"sso"`
}

type Sso struct {
	OvStruct
	Methods []Method `xml:"methods"`
}

type Statistics struct {
	Statistics []Statistic `xml:"statistic"`
}

type Statistic struct {
	OvStruct
	Brick         *GlusterBrick  `xml:"brick"`
	Comment       string         `xml:"comment"`
	Description   string         `xml:"description"`
	Disk          *Disk          `xml:"disk"`
	GlusterVolume *GlusterVolume `xml:"gluster_volume"`
	Host          *Host          `xml:"host"`
	HostNic       *HostNic       `xml:"host_nic"`
	HostNumaNode  *NumaNode      `xml:"host_numa_node"`
	Id            string         `xml:"id,attr"`
	Kind          StatisticKind  `xml:"kind"`
	Name          string         `xml:"name"`
	Nic           *Nic           `xml:"nic"`
	Step          *Step          `xml:"step"`
	Type          ValueType      `xml:"type"`
	Unit          StatisticUnit  `xml:"unit"`
	Values        []Value        `xml:"values"`
	Vm            *Vm            `xml:"vm"`
}

type Steps struct {
	Steps []Step `xml:"step"`
}

type Step struct {
	OvStruct
	Comment       string             `xml:"comment"`
	Description   string             `xml:"description"`
	EndTime       time.Time          `xml:"end_time"`
	ExecutionHost *Host              `xml:"execution_host"`
	External      bool               `xml:"external"`
	ExternalType  ExternalSystemType `xml:"external_type"`
	Id            string             `xml:"id,attr"`
	Job           *Job               `xml:"job"`
	Name          string             `xml:"name"`
	Number        int64              `xml:"number"`
	ParentStep    *Step              `xml:"parent_step"`
	Progress      int64              `xml:"progress"`
	StartTime     time.Time          `xml:"start_time"`
	Statistics    []Statistic        `xml:"statistics"`
	Status        StepStatus         `xml:"status"`
	Type          StepEnum           `xml:"type"`
}

type StorageConnections struct {
	StorageConnections []StorageConnection `xml:"storage_connection"`
}

type StorageConnection struct {
	OvStruct
	Address      string      `xml:"address"`
	Comment      string      `xml:"comment"`
	Description  string      `xml:"description"`
	Host         *Host       `xml:"host"`
	Id           string      `xml:"id,attr"`
	MountOptions string      `xml:"mount_options"`
	Name         string      `xml:"name"`
	NfsRetrans   int64       `xml:"nfs_retrans"`
	NfsTimeo     int64       `xml:"nfs_timeo"`
	NfsVersion   NfsVersion  `xml:"nfs_version"`
	Password     string      `xml:"password"`
	Path         string      `xml:"path"`
	Port         int64       `xml:"port"`
	Portal       string      `xml:"portal"`
	Target       string      `xml:"target"`
	Type         StorageType `xml:"type"`
	Username     string      `xml:"username"`
	VfsType      string      `xml:"vfs_type"`
}

type StorageConnectionExtensions struct {
	StorageConnectionExtensions []StorageConnectionExtension `xml:"storage_connection_extension"`
}

type StorageConnectionExtension struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Host        *Host  `xml:"host"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Password    string `xml:"password"`
	Target      string `xml:"target"`
	Username    string `xml:"username"`
}

type StorageDomains struct {
	StorageDomains []StorageDomain `xml:"storage_domain"`
}

type StorageDomain struct {
	OvStruct
	Available                  int64               `xml:"available"`
	Comment                    string              `xml:"comment"`
	Committed                  int64               `xml:"committed"`
	CriticalSpaceActionBlocker int64               `xml:"critical_space_action_blocker"`
	DataCenter                 *DataCenter         `xml:"data_center"`
	DataCenters                []DataCenter        `xml:"data_centers"`
	Description                string              `xml:"description"`
	DiscardAfterDelete         bool                `xml:"discard_after_delete"`
	DiskProfiles               []DiskProfile       `xml:"disk_profiles"`
	DiskSnapshots              []DiskSnapshot      `xml:"disk_snapshots"`
	Disks                      []Disk              `xml:"disks"`
	ExternalStatus             ExternalStatus      `xml:"external_status"`
	Files                      []File              `xml:"files"`
	Host                       *Host               `xml:"host"`
	Id                         string              `xml:"id,attr"`
	Images                     []Image             `xml:"images"`
	Import                     bool                `xml:"import"`
	Master                     bool                `xml:"master"`
	Name                       string              `xml:"name"`
	Permissions                []Permission        `xml:"permissions"`
	Status                     StorageDomainStatus `xml:"status"`
	Storage                    *HostStorage        `xml:"storage"`
	StorageConnections         []StorageConnection `xml:"storage_connections"`
	StorageFormat              StorageFormat       `xml:"storage_format"`
	SupportsDiscard            bool                `xml:"supports_discard"`
	SupportsDiscardZeroesData  bool                `xml:"supports_discard_zeroes_data"`
	Templates                  []Template          `xml:"templates"`
	Type                       StorageDomainType   `xml:"type"`
	Used                       int64               `xml:"used"`
	Vms                        []Vm                `xml:"vms"`
	WarningLowSpaceIndicator   int64               `xml:"warning_low_space_indicator"`
	WipeAfterDelete            bool                `xml:"wipe_after_delete"`
}

type StorageDomainLeases struct {
	StorageDomainLeases []StorageDomainLease `xml:"storage_domain_lease"`
}

type StorageDomainLease struct {
	OvStruct
	StorageDomain *StorageDomain `xml:"storage_domain"`
}

type Tags struct {
	Tags []Tag `xml:"tag"`
}

type Tag struct {
	OvStruct
	Comment     string    `xml:"comment"`
	Description string    `xml:"description"`
	Group       *Group    `xml:"group"`
	Host        *Host     `xml:"host"`
	Id          string    `xml:"id,attr"`
	Name        string    `xml:"name"`
	Parent      *Tag      `xml:"parent"`
	Template    *Template `xml:"template"`
	User        *User     `xml:"user"`
	Vm          *Vm       `xml:"vm"`
}

type TemplateVersions struct {
	TemplateVersions []TemplateVersion `xml:"template_version"`
}

type TemplateVersion struct {
	OvStruct
	BaseTemplate  *Template `xml:"base_template"`
	VersionName   string    `xml:"version_name"`
	VersionNumber int64     `xml:"version_number"`
}

type Tickets struct {
	Tickets []Ticket `xml:"ticket"`
}

type Ticket struct {
	OvStruct
	Expiry int64  `xml:"expiry"`
	Value  string `xml:"value"`
}

type TimeZones struct {
	TimeZones []TimeZone `xml:"time_zone"`
}

type TimeZone struct {
	OvStruct
	Name      string `xml:"name"`
	UtcOffset string `xml:"utc_offset"`
}

type TransparentHugePagess struct {
	TransparentHugePagess []TransparentHugePages `xml:"transparent_huge_pages"`
}

type TransparentHugePages struct {
	OvStruct
	Enabled bool `xml:"enabled"`
}

type UnmanagedNetworks struct {
	UnmanagedNetworks []UnmanagedNetwork `xml:"unmanaged_network"`
}

type UnmanagedNetwork struct {
	OvStruct
	Comment     string   `xml:"comment"`
	Description string   `xml:"description"`
	Host        *Host    `xml:"host"`
	HostNic     *HostNic `xml:"host_nic"`
	Id          string   `xml:"id,attr"`
	Name        string   `xml:"name"`
}

type Usbs struct {
	Usbs []Usb `xml:"usb"`
}

type Usb struct {
	OvStruct
	Enabled bool    `xml:"enabled"`
	Type    UsbType `xml:"type"`
}

type Users struct {
	Users []User `xml:"user"`
}

type User struct {
	OvStruct
	Comment       string         `xml:"comment"`
	Department    string         `xml:"department"`
	Description   string         `xml:"description"`
	Domain        *Domain        `xml:"domain"`
	DomainEntryId string         `xml:"domain_entry_id"`
	Email         string         `xml:"email"`
	Groups        []Group        `xml:"groups"`
	Id            string         `xml:"id,attr"`
	LastName      string         `xml:"last_name"`
	LoggedIn      bool           `xml:"logged_in"`
	Name          string         `xml:"name"`
	Namespace     string         `xml:"namespace"`
	Password      string         `xml:"password"`
	Permissions   []Permission   `xml:"permissions"`
	Principal     string         `xml:"principal"`
	Roles         []Role         `xml:"roles"`
	SshPublicKeys []SshPublicKey `xml:"ssh_public_keys"`
	Tags          []Tag          `xml:"tags"`
	UserName      string         `xml:"user_name"`
}

type Values struct {
	Values []Value `xml:"value"`
}

type Value struct {
	OvStruct
	Datum  float64 `xml:"datum"`
	Detail string  `xml:"detail"`
}

type VcpuPins struct {
	VcpuPins []VcpuPin `xml:"vcpu_pin"`
}

type VcpuPin struct {
	OvStruct
	CpuSet string `xml:"cpu_set"`
	Vcpu   int64  `xml:"vcpu"`
}

type Vendors struct {
	Vendors []Vendor `xml:"vendor"`
}

type Vendor struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
}

type Versions struct {
	Versions []Version `xml:"version"`
}

type Version struct {
	OvStruct
	Build       int64  `xml:"build"`
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	FullVersion string `xml:"full_version"`
	Id          string `xml:"id,attr"`
	Major       int64  `xml:"major"`
	Minor       int64  `xml:"minor"`
	Name        string `xml:"name"`
	Revision    int64  `xml:"revision"`
}

type VirtioScsis struct {
	VirtioScsis []VirtioScsi `xml:"virtio_scsi"`
}

type VirtioScsi struct {
	OvStruct
	Enabled bool `xml:"enabled"`
}

type VirtualNumaNodes struct {
	VirtualNumaNodes []VirtualNumaNode `xml:"virtual_numa_node"`
}

type VirtualNumaNode struct {
	OvStruct
	Comment      string        `xml:"comment"`
	Cpu          *Cpu          `xml:"cpu"`
	Description  string        `xml:"description"`
	Host         *Host         `xml:"host"`
	Id           string        `xml:"id,attr"`
	Index        int64         `xml:"index"`
	Memory       int64         `xml:"memory"`
	Name         string        `xml:"name"`
	NodeDistance string        `xml:"node_distance"`
	NumaNodePins []NumaNodePin `xml:"numa_node_pins"`
	Statistics   []Statistic   `xml:"statistics"`
	Vm           *Vm           `xml:"vm"`
}

type Vlans struct {
	Vlans []Vlan `xml:"vlan"`
}

type Vlan struct {
	OvStruct
	Id int64 `xml:"id,attr"`
}

type VmBases struct {
	VmBases []VmBase `xml:"vm_base"`
}

type VmBase struct {
	OvStruct
	Bios                       *Bios               `xml:"bios"`
	Cluster                    *Cluster            `xml:"cluster"`
	Comment                    string              `xml:"comment"`
	Console                    *Console            `xml:"console"`
	Cpu                        *Cpu                `xml:"cpu"`
	CpuProfile                 *CpuProfile         `xml:"cpu_profile"`
	CpuShares                  int64               `xml:"cpu_shares"`
	CreationTime               time.Time           `xml:"creation_time"`
	CustomCompatibilityVersion *Version            `xml:"custom_compatibility_version"`
	CustomCpuModel             string              `xml:"custom_cpu_model"`
	CustomEmulatedMachine      string              `xml:"custom_emulated_machine"`
	CustomProperties           []CustomProperty    `xml:"custom_properties"`
	DeleteProtected            bool                `xml:"delete_protected"`
	Description                string              `xml:"description"`
	Display                    *Display            `xml:"display"`
	Domain                     *Domain             `xml:"domain"`
	HighAvailability           *HighAvailability   `xml:"high_availability"`
	Id                         string              `xml:"id,attr"`
	Initialization             *Initialization     `xml:"initialization"`
	Io                         *Io                 `xml:"io"`
	LargeIcon                  *Icon               `xml:"large_icon"`
	Lease                      *StorageDomainLease `xml:"lease"`
	Memory                     int64               `xml:"memory"`
	MemoryPolicy               *MemoryPolicy       `xml:"memory_policy"`
	Migration                  *MigrationOptions   `xml:"migration"`
	MigrationDowntime          int64               `xml:"migration_downtime"`
	Name                       string              `xml:"name"`
	Origin                     string              `xml:"origin"`
	Os                         *OperatingSystem    `xml:"os"`
	Quota                      *Quota              `xml:"quota"`
	RngDevice                  *RngDevice          `xml:"rng_device"`
	SerialNumber               *SerialNumber       `xml:"serial_number"`
	SmallIcon                  *Icon               `xml:"small_icon"`
	SoundcardEnabled           bool                `xml:"soundcard_enabled"`
	Sso                        *Sso                `xml:"sso"`
	StartPaused                bool                `xml:"start_paused"`
	Stateless                  bool                `xml:"stateless"`
	StorageDomain              *StorageDomain      `xml:"storage_domain"`
	TimeZone                   *TimeZone           `xml:"time_zone"`
	TunnelMigration            bool                `xml:"tunnel_migration"`
	Type                       VmType              `xml:"type"`
	Usb                        *Usb                `xml:"usb"`
	VirtioScsi                 *VirtioScsi         `xml:"virtio_scsi"`
}

type VmPlacementPolicys struct {
	VmPlacementPolicys []VmPlacementPolicy `xml:"vm_placement_policy"`
}

type VmPlacementPolicy struct {
	OvStruct
	Affinity VmAffinity `xml:"affinity"`
	Hosts    []Host     `xml:"hosts"`
}

type VmPools struct {
	VmPools []VmPool `xml:"vm_pool"`
}

type VmPool struct {
	OvStruct
	AutoStorageSelect        bool          `xml:"auto_storage_select"`
	Cluster                  *Cluster      `xml:"cluster"`
	Comment                  string        `xml:"comment"`
	Description              string        `xml:"description"`
	Display                  *Display      `xml:"display"`
	Id                       string        `xml:"id,attr"`
	InstanceType             *InstanceType `xml:"instance_type"`
	MaxUserVms               int64         `xml:"max_user_vms"`
	Name                     string        `xml:"name"`
	Permissions              []Permission  `xml:"permissions"`
	PrestartedVms            int64         `xml:"prestarted_vms"`
	RngDevice                *RngDevice    `xml:"rng_device"`
	Size                     int64         `xml:"size"`
	SoundcardEnabled         bool          `xml:"soundcard_enabled"`
	Stateful                 bool          `xml:"stateful"`
	Template                 *Template     `xml:"template"`
	Type                     VmPoolType    `xml:"type"`
	UseLatestTemplateVersion bool          `xml:"use_latest_template_version"`
	Vm                       *Vm           `xml:"vm"`
}

type VmSummarys struct {
	VmSummarys []VmSummary `xml:"vm_summary"`
}

type VmSummary struct {
	OvStruct
	Active    int64 `xml:"active"`
	Migrating int64 `xml:"migrating"`
	Total     int64 `xml:"total"`
}

type VnicPassThroughs struct {
	VnicPassThroughs []VnicPassThrough `xml:"vnic_pass_through"`
}

type VnicPassThrough struct {
	OvStruct
	Mode VnicPassThroughMode `xml:"mode"`
}

type VnicProfiles struct {
	VnicProfiles []VnicProfile `xml:"vnic_profile"`
}

type VnicProfile struct {
	OvStruct
	Comment          string           `xml:"comment"`
	CustomProperties []CustomProperty `xml:"custom_properties"`
	Description      string           `xml:"description"`
	Id               string           `xml:"id,attr"`
	Migratable       bool             `xml:"migratable"`
	Name             string           `xml:"name"`
	Network          *Network         `xml:"network"`
	NetworkFilter    *NetworkFilter   `xml:"network_filter"`
	PassThrough      *VnicPassThrough `xml:"pass_through"`
	Permissions      []Permission     `xml:"permissions"`
	PortMirroring    bool             `xml:"port_mirroring"`
	Qos              *Qos             `xml:"qos"`
}

type VnicProfileMappings struct {
	VnicProfileMappings []VnicProfileMapping `xml:"vnic_profile_mapping"`
}

type VnicProfileMapping struct {
	OvStruct
	SourceNetworkName        string       `xml:"source_network_name"`
	SourceNetworkProfileName string       `xml:"source_network_profile_name"`
	TargetVnicProfile        *VnicProfile `xml:"target_vnic_profile"`
}

type VolumeGroups struct {
	VolumeGroups []VolumeGroup `xml:"volume_group"`
}

type VolumeGroup struct {
	OvStruct
	Id           string        `xml:"id,attr"`
	LogicalUnits []LogicalUnit `xml:"logical_units"`
	Name         string        `xml:"name"`
}

type Weights struct {
	Weights []Weight `xml:"weight"`
}

type Weight struct {
	OvStruct
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	Factor               int64                 `xml:"factor"`
	Id                   string                `xml:"id,attr"`
	Name                 string                `xml:"name"`
	SchedulingPolicy     *SchedulingPolicy     `xml:"scheduling_policy"`
	SchedulingPolicyUnit *SchedulingPolicyUnit `xml:"scheduling_policy_unit"`
}

type Actions struct {
	Actions []Action `xml:"action"`
}

type Action struct {
	OvStruct
	AllowPartialImport             bool                                  `xml:"allow_partial_import"`
	Async                          bool                                  `xml:"async"`
	Bricks                         []GlusterBrick                        `xml:"bricks"`
	Certificates                   []Certificate                         `xml:"certificates"`
	CheckConnectivity              bool                                  `xml:"check_connectivity"`
	Clone                          bool                                  `xml:"clone"`
	Cluster                        *Cluster                              `xml:"cluster"`
	CollapseSnapshots              bool                                  `xml:"collapse_snapshots"`
	Comment                        string                                `xml:"comment"`
	ConnectivityTimeout            int64                                 `xml:"connectivity_timeout"`
	DataCenter                     *DataCenter                           `xml:"data_center"`
	DeployHostedEngine             bool                                  `xml:"deploy_hosted_engine"`
	Description                    string                                `xml:"description"`
	Details                        *GlusterVolumeProfileDetails          `xml:"details"`
	DiscardSnapshots               bool                                  `xml:"discard_snapshots"`
	Disk                           *Disk                                 `xml:"disk"`
	Disks                          []Disk                                `xml:"disks"`
	Exclusive                      bool                                  `xml:"exclusive"`
	Fault                          *Fault                                `xml:"fault"`
	FenceType                      string                                `xml:"fence_type"`
	Filter                         bool                                  `xml:"filter"`
	FixLayout                      bool                                  `xml:"fix_layout"`
	Force                          bool                                  `xml:"force"`
	GracePeriod                    *GracePeriod                          `xml:"grace_period"`
	Host                           *Host                                 `xml:"host"`
	Id                             string                                `xml:"id,attr"`
	Image                          string                                `xml:"image"`
	ImportAsTemplate               bool                                  `xml:"import_as_template"`
	IsAttached                     bool                                  `xml:"is_attached"`
	Iscsi                          *IscsiDetails                         `xml:"iscsi"`
	IscsiTargets                   []string                              `xml:"iscsi_targets"`
	Job                            *Job                                  `xml:"job"`
	LogicalUnits                   []LogicalUnit                         `xml:"logical_units"`
	MaintenanceEnabled             bool                                  `xml:"maintenance_enabled"`
	ModifiedBonds                  []HostNic                             `xml:"modified_bonds"`
	ModifiedLabels                 []NetworkLabel                        `xml:"modified_labels"`
	ModifiedNetworkAttachments     []NetworkAttachment                   `xml:"modified_network_attachments"`
	Name                           string                                `xml:"name"`
	Option                         *Option                               `xml:"option"`
	Pause                          bool                                  `xml:"pause"`
	PowerManagement                *PowerManagement                      `xml:"power_management"`
	ProxyTicket                    *ProxyTicket                          `xml:"proxy_ticket"`
	Reason                         string                                `xml:"reason"`
	ReassignBadMacs                bool                                  `xml:"reassign_bad_macs"`
	RemoteViewerConnectionFile     string                                `xml:"remote_viewer_connection_file"`
	RemovedBonds                   []HostNic                             `xml:"removed_bonds"`
	RemovedLabels                  []NetworkLabel                        `xml:"removed_labels"`
	RemovedNetworkAttachments      []NetworkAttachment                   `xml:"removed_network_attachments"`
	ResolutionType                 string                                `xml:"resolution_type"`
	RestoreMemory                  bool                                  `xml:"restore_memory"`
	RootPassword                   string                                `xml:"root_password"`
	Snapshot                       *Snapshot                             `xml:"snapshot"`
	Ssh                            *Ssh                                  `xml:"ssh"`
	Status                         string                                `xml:"status"`
	StopGlusterService             bool                                  `xml:"stop_gluster_service"`
	StorageDomain                  *StorageDomain                        `xml:"storage_domain"`
	StorageDomains                 []StorageDomain                       `xml:"storage_domains"`
	Succeeded                      bool                                  `xml:"succeeded"`
	SynchronizedNetworkAttachments []NetworkAttachment                   `xml:"synchronized_network_attachments"`
	Template                       *Template                             `xml:"template"`
	Ticket                         *Ticket                               `xml:"ticket"`
	UndeployHostedEngine           bool                                  `xml:"undeploy_hosted_engine"`
	UseCloudInit                   bool                                  `xml:"use_cloud_init"`
	UseSysprep                     bool                                  `xml:"use_sysprep"`
	VirtualFunctionsConfiguration  *HostNicVirtualFunctionsConfiguration `xml:"virtual_functions_configuration"`
	Vm                             *Vm                                   `xml:"vm"`
	VnicProfileMappings            []VnicProfileMapping                  `xml:"vnic_profile_mappings"`
}

type AffinityGroups struct {
	AffinityGroups []AffinityGroup `xml:"affinity_group"`
}

type AffinityGroup struct {
	OvStruct
	Cluster     *Cluster      `xml:"cluster"`
	Comment     string        `xml:"comment"`
	Description string        `xml:"description"`
	Enforcing   bool          `xml:"enforcing"`
	Hosts       []Host        `xml:"hosts"`
	HostsRule   *AffinityRule `xml:"hosts_rule"`
	Id          string        `xml:"id,attr"`
	Name        string        `xml:"name"`
	Positive    bool          `xml:"positive"`
	Vms         []Vm          `xml:"vms"`
	VmsRule     *AffinityRule `xml:"vms_rule"`
}

type AffinityLabels struct {
	AffinityLabels []AffinityLabel `xml:"affinity_label"`
}

type AffinityLabel struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Hosts       []Host `xml:"hosts"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
	ReadOnly    bool   `xml:"read_only"`
	Vms         []Vm   `xml:"vms"`
}

type Agents struct {
	Agents []Agent `xml:"agent"`
}

type Agent struct {
	OvStruct
	Address        string   `xml:"address"`
	Comment        string   `xml:"comment"`
	Concurrent     bool     `xml:"concurrent"`
	Description    string   `xml:"description"`
	EncryptOptions bool     `xml:"encrypt_options"`
	Host           *Host    `xml:"host"`
	Id             string   `xml:"id,attr"`
	Name           string   `xml:"name"`
	Options        []Option `xml:"options"`
	Order          int64    `xml:"order"`
	Password       string   `xml:"password"`
	Port           int64    `xml:"port"`
	Type           string   `xml:"type"`
	Username       string   `xml:"username"`
}

type Applications struct {
	Applications []Application `xml:"application"`
}

type Application struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Vm          *Vm    `xml:"vm"`
}

type AuthorizedKeys struct {
	AuthorizedKeys []AuthorizedKey `xml:"authorized_key"`
}

type AuthorizedKey struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Key         string `xml:"key"`
	Name        string `xml:"name"`
	User        *User  `xml:"user"`
}

type Balances struct {
	Balances []Balance `xml:"balance"`
}

type Balance struct {
	OvStruct
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	Id                   string                `xml:"id,attr"`
	Name                 string                `xml:"name"`
	SchedulingPolicy     *SchedulingPolicy     `xml:"scheduling_policy"`
	SchedulingPolicyUnit *SchedulingPolicyUnit `xml:"scheduling_policy_unit"`
}

type Bookmarks struct {
	Bookmarks []Bookmark `xml:"bookmark"`
}

type Bookmark struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Value       string `xml:"value"`
}

type BrickProfileDetails struct {
	BrickProfileDetails []BrickProfileDetail `xml:"brick_profile_detail"`
}

type BrickProfileDetail struct {
	OvStruct
	Brick          *GlusterBrick   `xml:"brick"`
	ProfileDetails []ProfileDetail `xml:"profile_details"`
}

type Certificates struct {
	Certificates []Certificate `xml:"certificate"`
}

type Certificate struct {
	OvStruct
	Comment      string `xml:"comment"`
	Content      string `xml:"content"`
	Description  string `xml:"description"`
	Id           string `xml:"id,attr"`
	Name         string `xml:"name"`
	Organization string `xml:"organization"`
	Subject      string `xml:"subject"`
}

type Clusters struct {
	Clusters []Cluster `xml:"cluster"`
}

type Cluster struct {
	OvStruct
	AffinityGroups                   []AffinityGroup   `xml:"affinity_groups"`
	BallooningEnabled                bool              `xml:"ballooning_enabled"`
	Comment                          string            `xml:"comment"`
	Cpu                              *Cpu              `xml:"cpu"`
	CpuProfiles                      []CpuProfile      `xml:"cpu_profiles"`
	CustomSchedulingPolicyProperties []Property        `xml:"custom_scheduling_policy_properties"`
	DataCenter                       *DataCenter       `xml:"data_center"`
	Description                      string            `xml:"description"`
	Display                          *Display          `xml:"display"`
	ErrorHandling                    *ErrorHandling    `xml:"error_handling"`
	FencingPolicy                    *FencingPolicy    `xml:"fencing_policy"`
	GlusterHooks                     []GlusterHook     `xml:"gluster_hooks"`
	GlusterService                   bool              `xml:"gluster_service"`
	GlusterTunedProfile              string            `xml:"gluster_tuned_profile"`
	GlusterVolumes                   []GlusterVolume   `xml:"gluster_volumes"`
	HaReservation                    bool              `xml:"ha_reservation"`
	Id                               string            `xml:"id,attr"`
	Ksm                              *Ksm              `xml:"ksm"`
	MacPool                          *MacPool          `xml:"mac_pool"`
	MaintenanceReasonRequired        bool              `xml:"maintenance_reason_required"`
	ManagementNetwork                *Network          `xml:"management_network"`
	MemoryPolicy                     *MemoryPolicy     `xml:"memory_policy"`
	Migration                        *MigrationOptions `xml:"migration"`
	Name                             string            `xml:"name"`
	NetworkFilters                   []NetworkFilter   `xml:"network_filters"`
	Networks                         []Network         `xml:"networks"`
	OptionalReason                   bool              `xml:"optional_reason"`
	Permissions                      []Permission      `xml:"permissions"`
	RequiredRngSources               []RngSource       `xml:"required_rng_sources"`
	SchedulingPolicy                 *SchedulingPolicy `xml:"scheduling_policy"`
	SerialNumber                     *SerialNumber     `xml:"serial_number"`
	SupportedVersions                []Version         `xml:"supported_versions"`
	SwitchType                       SwitchType        `xml:"switch_type"`
	ThreadsAsCores                   bool              `xml:"threads_as_cores"`
	TrustedService                   bool              `xml:"trusted_service"`
	TunnelMigration                  bool              `xml:"tunnel_migration"`
	Version                          *Version          `xml:"version"`
	VirtService                      bool              `xml:"virt_service"`
}

type ClusterLevels struct {
	ClusterLevels []ClusterLevel `xml:"cluster_level"`
}

type ClusterLevel struct {
	OvStruct
	Comment     string    `xml:"comment"`
	CpuTypes    []CpuType `xml:"cpu_types"`
	Description string    `xml:"description"`
	Id          string    `xml:"id,attr"`
	Name        string    `xml:"name"`
	Permits     []Permit  `xml:"permits"`
}

type CpuProfiles struct {
	CpuProfiles []CpuProfile `xml:"cpu_profile"`
}

type CpuProfile struct {
	OvStruct
	Cluster     *Cluster     `xml:"cluster"`
	Comment     string       `xml:"comment"`
	Description string       `xml:"description"`
	Id          string       `xml:"id,attr"`
	Name        string       `xml:"name"`
	Permissions []Permission `xml:"permissions"`
	Qos         *Qos         `xml:"qos"`
}

type DataCenters struct {
	DataCenters []DataCenter `xml:"data_center"`
}

type DataCenter struct {
	OvStruct
	Clusters          []Cluster        `xml:"clusters"`
	Comment           string           `xml:"comment"`
	Description       string           `xml:"description"`
	Id                string           `xml:"id,attr"`
	IscsiBonds        []IscsiBond      `xml:"iscsi_bonds"`
	Local             bool             `xml:"local"`
	MacPool           *MacPool         `xml:"mac_pool"`
	Name              string           `xml:"name"`
	Networks          []Network        `xml:"networks"`
	Permissions       []Permission     `xml:"permissions"`
	Qoss              []Qos            `xml:"qoss"`
	QuotaMode         QuotaModeType    `xml:"quota_mode"`
	Quotas            []Quota          `xml:"quotas"`
	Status            DataCenterStatus `xml:"status"`
	StorageDomains    []StorageDomain  `xml:"storage_domains"`
	StorageFormat     StorageFormat    `xml:"storage_format"`
	SupportedVersions []Version        `xml:"supported_versions"`
	Version           *Version         `xml:"version"`
}

type Devices struct {
	Devices []Device `xml:"device"`
}

type Device struct {
	OvStruct
	Comment      string        `xml:"comment"`
	Description  string        `xml:"description"`
	Id           string        `xml:"id,attr"`
	InstanceType *InstanceType `xml:"instance_type"`
	Name         string        `xml:"name"`
	Template     *Template     `xml:"template"`
	Vm           *Vm           `xml:"vm"`
	Vms          []Vm          `xml:"vms"`
}

type Disks struct {
	Disks []Disk `xml:"disk"`
}

type Disk struct {
	OvStruct
	Active              bool                 `xml:"active"`
	ActualSize          int64                `xml:"actual_size"`
	Alias               string               `xml:"alias"`
	Bootable            bool                 `xml:"bootable"`
	Comment             string               `xml:"comment"`
	Description         string               `xml:"description"`
	DiskProfile         *DiskProfile         `xml:"disk_profile"`
	Format              DiskFormat           `xml:"format"`
	Id                  string               `xml:"id,attr"`
	ImageId             string               `xml:"image_id"`
	InitialSize         int64                `xml:"initial_size"`
	InstanceType        *InstanceType        `xml:"instance_type"`
	Interface           DiskInterface        `xml:"interface"`
	LogicalName         string               `xml:"logical_name"`
	LunStorage          *HostStorage         `xml:"lun_storage"`
	Name                string               `xml:"name"`
	OpenstackVolumeType *OpenStackVolumeType `xml:"openstack_volume_type"`
	Permissions         []Permission         `xml:"permissions"`
	PropagateErrors     bool                 `xml:"propagate_errors"`
	ProvisionedSize     int64                `xml:"provisioned_size"`
	QcowVersion         QcowVersion          `xml:"qcow_version"`
	Quota               *Quota               `xml:"quota"`
	ReadOnly            bool                 `xml:"read_only"`
	Sgio                ScsiGenericIO        `xml:"sgio"`
	Shareable           bool                 `xml:"shareable"`
	Snapshot            *Snapshot            `xml:"snapshot"`
	Sparse              bool                 `xml:"sparse"`
	Statistics          []Statistic          `xml:"statistics"`
	Status              DiskStatus           `xml:"status"`
	StorageDomain       *StorageDomain       `xml:"storage_domain"`
	StorageDomains      []StorageDomain      `xml:"storage_domains"`
	StorageType         DiskStorageType      `xml:"storage_type"`
	Template            *Template            `xml:"template"`
	UsesScsiReservation bool                 `xml:"uses_scsi_reservation"`
	Vm                  *Vm                  `xml:"vm"`
	Vms                 []Vm                 `xml:"vms"`
	WipeAfterDelete     bool                 `xml:"wipe_after_delete"`
}

type DiskAttachments struct {
	DiskAttachments []DiskAttachment `xml:"disk_attachment"`
}

type DiskAttachment struct {
	OvStruct
	Active              bool          `xml:"active"`
	Bootable            bool          `xml:"bootable"`
	Comment             string        `xml:"comment"`
	Description         string        `xml:"description"`
	Disk                *Disk         `xml:"disk"`
	Id                  string        `xml:"id,attr"`
	Interface           DiskInterface `xml:"interface"`
	LogicalName         string        `xml:"logical_name"`
	Name                string        `xml:"name"`
	PassDiscard         bool          `xml:"pass_discard"`
	Template            *Template     `xml:"template"`
	UsesScsiReservation bool          `xml:"uses_scsi_reservation"`
	Vm                  *Vm           `xml:"vm"`
}

type DiskProfiles struct {
	DiskProfiles []DiskProfile `xml:"disk_profile"`
}

type DiskProfile struct {
	OvStruct
	Comment       string         `xml:"comment"`
	Description   string         `xml:"description"`
	Id            string         `xml:"id,attr"`
	Name          string         `xml:"name"`
	Permissions   []Permission   `xml:"permissions"`
	Qos           *Qos           `xml:"qos"`
	StorageDomain *StorageDomain `xml:"storage_domain"`
}

type DiskSnapshots struct {
	DiskSnapshots []DiskSnapshot `xml:"disk_snapshot"`
}

type DiskSnapshot struct {
	OvStruct
	Active              bool                 `xml:"active"`
	ActualSize          int64                `xml:"actual_size"`
	Alias               string               `xml:"alias"`
	Bootable            bool                 `xml:"bootable"`
	Comment             string               `xml:"comment"`
	Description         string               `xml:"description"`
	Disk                *Disk                `xml:"disk"`
	DiskProfile         *DiskProfile         `xml:"disk_profile"`
	Format              DiskFormat           `xml:"format"`
	Id                  string               `xml:"id,attr"`
	ImageId             string               `xml:"image_id"`
	InitialSize         int64                `xml:"initial_size"`
	InstanceType        *InstanceType        `xml:"instance_type"`
	Interface           DiskInterface        `xml:"interface"`
	LogicalName         string               `xml:"logical_name"`
	LunStorage          *HostStorage         `xml:"lun_storage"`
	Name                string               `xml:"name"`
	OpenstackVolumeType *OpenStackVolumeType `xml:"openstack_volume_type"`
	Permissions         []Permission         `xml:"permissions"`
	PropagateErrors     bool                 `xml:"propagate_errors"`
	ProvisionedSize     int64                `xml:"provisioned_size"`
	QcowVersion         QcowVersion          `xml:"qcow_version"`
	Quota               *Quota               `xml:"quota"`
	ReadOnly            bool                 `xml:"read_only"`
	Sgio                ScsiGenericIO        `xml:"sgio"`
	Shareable           bool                 `xml:"shareable"`
	Snapshot            *Snapshot            `xml:"snapshot"`
	Sparse              bool                 `xml:"sparse"`
	Statistics          []Statistic          `xml:"statistics"`
	Status              DiskStatus           `xml:"status"`
	StorageDomain       *StorageDomain       `xml:"storage_domain"`
	StorageDomains      []StorageDomain      `xml:"storage_domains"`
	StorageType         DiskStorageType      `xml:"storage_type"`
	Template            *Template            `xml:"template"`
	UsesScsiReservation bool                 `xml:"uses_scsi_reservation"`
	Vm                  *Vm                  `xml:"vm"`
	Vms                 []Vm                 `xml:"vms"`
	WipeAfterDelete     bool                 `xml:"wipe_after_delete"`
}

type Domains struct {
	Domains []Domain `xml:"domain"`
}

type Domain struct {
	OvStruct
	Comment     string  `xml:"comment"`
	Description string  `xml:"description"`
	Groups      []Group `xml:"groups"`
	Id          string  `xml:"id,attr"`
	Name        string  `xml:"name"`
	User        *User   `xml:"user"`
	Users       []User  `xml:"users"`
}

type Events struct {
	Events []Event `xml:"event"`
}

type Event struct {
	OvStruct
	Cluster       *Cluster       `xml:"cluster"`
	Code          int64          `xml:"code"`
	Comment       string         `xml:"comment"`
	CorrelationId string         `xml:"correlation_id"`
	CustomData    string         `xml:"custom_data"`
	CustomId      int64          `xml:"custom_id"`
	DataCenter    *DataCenter    `xml:"data_center"`
	Description   string         `xml:"description"`
	FloodRate     int64          `xml:"flood_rate"`
	Host          *Host          `xml:"host"`
	Id            string         `xml:"id,attr"`
	Name          string         `xml:"name"`
	Origin        string         `xml:"origin"`
	Severity      LogSeverity    `xml:"severity"`
	StorageDomain *StorageDomain `xml:"storage_domain"`
	Template      *Template      `xml:"template"`
	Time          time.Time      `xml:"time"`
	User          *User          `xml:"user"`
	Vm            *Vm            `xml:"vm"`
}

type ExternalComputeResources struct {
	ExternalComputeResources []ExternalComputeResource `xml:"external_compute_resource"`
}

type ExternalComputeResource struct {
	OvStruct
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider"`
	Id                   string                `xml:"id,attr"`
	Name                 string                `xml:"name"`
	Provider             string                `xml:"provider"`
	Url                  string                `xml:"url"`
	User                 string                `xml:"user"`
}

type ExternalDiscoveredHosts struct {
	ExternalDiscoveredHosts []ExternalDiscoveredHost `xml:"external_discovered_host"`
}

type ExternalDiscoveredHost struct {
	OvStruct
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider"`
	Id                   string                `xml:"id,attr"`
	Ip                   string                `xml:"ip"`
	LastReport           string                `xml:"last_report"`
	Mac                  string                `xml:"mac"`
	Name                 string                `xml:"name"`
	SubnetName           string                `xml:"subnet_name"`
}

type ExternalHosts struct {
	ExternalHosts []ExternalHost `xml:"external_host"`
}

type ExternalHost struct {
	OvStruct
	Address              string                `xml:"address"`
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider"`
	Id                   string                `xml:"id,attr"`
	Name                 string                `xml:"name"`
}

type ExternalHostGroups struct {
	ExternalHostGroups []ExternalHostGroup `xml:"external_host_group"`
}

type ExternalHostGroup struct {
	OvStruct
	ArchitectureName     string                `xml:"architecture_name"`
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	DomainName           string                `xml:"domain_name"`
	ExternalHostProvider *ExternalHostProvider `xml:"external_host_provider"`
	Id                   string                `xml:"id,attr"`
	Name                 string                `xml:"name"`
	OperatingSystemName  string                `xml:"operating_system_name"`
	SubnetName           string                `xml:"subnet_name"`
}

type ExternalProviders struct {
	ExternalProviders []ExternalProvider `xml:"external_provider"`
}

type ExternalProvider struct {
	OvStruct
	AuthenticationUrl      string     `xml:"authentication_url"`
	Comment                string     `xml:"comment"`
	Description            string     `xml:"description"`
	Id                     string     `xml:"id,attr"`
	Name                   string     `xml:"name"`
	Password               string     `xml:"password"`
	Properties             []Property `xml:"properties"`
	RequiresAuthentication bool       `xml:"requires_authentication"`
	Url                    string     `xml:"url"`
	Username               string     `xml:"username"`
}

type Files struct {
	Files []File `xml:"file"`
}

type File struct {
	OvStruct
	Comment       string         `xml:"comment"`
	Content       string         `xml:"content"`
	Description   string         `xml:"description"`
	Id            string         `xml:"id,attr"`
	Name          string         `xml:"name"`
	StorageDomain *StorageDomain `xml:"storage_domain"`
	Type          string         `xml:"type"`
}

type Filters struct {
	Filters []Filter `xml:"filter"`
}

type Filter struct {
	OvStruct
	Comment              string                `xml:"comment"`
	Description          string                `xml:"description"`
	Id                   string                `xml:"id,attr"`
	Name                 string                `xml:"name"`
	Position             int64                 `xml:"position"`
	SchedulingPolicyUnit *SchedulingPolicyUnit `xml:"scheduling_policy_unit"`
}

type Floppys struct {
	Floppys []Floppy `xml:"floppy"`
}

type Floppy struct {
	OvStruct
	Comment      string        `xml:"comment"`
	Description  string        `xml:"description"`
	File         *File         `xml:"file"`
	Id           string        `xml:"id,attr"`
	InstanceType *InstanceType `xml:"instance_type"`
	Name         string        `xml:"name"`
	Template     *Template     `xml:"template"`
	Vm           *Vm           `xml:"vm"`
	Vms          []Vm          `xml:"vms"`
}

type GlusterBrickAdvancedDetailss struct {
	GlusterBrickAdvancedDetailss []GlusterBrickAdvancedDetails `xml:"gluster_brick_advanced_details"`
}

type GlusterBrickAdvancedDetails struct {
	OvStruct
	Comment        string              `xml:"comment"`
	Description    string              `xml:"description"`
	Device         string              `xml:"device"`
	FsName         string              `xml:"fs_name"`
	GlusterClients []GlusterClient     `xml:"gluster_clients"`
	Id             string              `xml:"id,attr"`
	InstanceType   *InstanceType       `xml:"instance_type"`
	MemoryPools    []GlusterMemoryPool `xml:"memory_pools"`
	MntOptions     string              `xml:"mnt_options"`
	Name           string              `xml:"name"`
	Pid            int64               `xml:"pid"`
	Port           int64               `xml:"port"`
	Template       *Template           `xml:"template"`
	Vm             *Vm                 `xml:"vm"`
	Vms            []Vm                `xml:"vms"`
}

type GlusterHooks struct {
	GlusterHooks []GlusterHook `xml:"gluster_hook"`
}

type GlusterHook struct {
	OvStruct
	Checksum       string              `xml:"checksum"`
	Cluster        *Cluster            `xml:"cluster"`
	Comment        string              `xml:"comment"`
	ConflictStatus int64               `xml:"conflict_status"`
	Conflicts      string              `xml:"conflicts"`
	Content        string              `xml:"content"`
	ContentType    HookContentType     `xml:"content_type"`
	Description    string              `xml:"description"`
	GlusterCommand string              `xml:"gluster_command"`
	Id             string              `xml:"id,attr"`
	Name           string              `xml:"name"`
	ServerHooks    []GlusterServerHook `xml:"server_hooks"`
	Stage          HookStage           `xml:"stage"`
	Status         GlusterHookStatus   `xml:"status"`
}

type GlusterMemoryPools struct {
	GlusterMemoryPools []GlusterMemoryPool `xml:"gluster_memory_pool"`
}

type GlusterMemoryPool struct {
	OvStruct
	AllocCount  int64  `xml:"alloc_count"`
	ColdCount   int64  `xml:"cold_count"`
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	HotCount    int64  `xml:"hot_count"`
	Id          string `xml:"id,attr"`
	MaxAlloc    int64  `xml:"max_alloc"`
	MaxStdalloc int64  `xml:"max_stdalloc"`
	Name        string `xml:"name"`
	PaddedSize  int64  `xml:"padded_size"`
	PoolMisses  int64  `xml:"pool_misses"`
	Type        string `xml:"type"`
}

type GlusterServerHooks struct {
	GlusterServerHooks []GlusterServerHook `xml:"gluster_server_hook"`
}

type GlusterServerHook struct {
	OvStruct
	Checksum    string            `xml:"checksum"`
	Comment     string            `xml:"comment"`
	ContentType HookContentType   `xml:"content_type"`
	Description string            `xml:"description"`
	Host        *Host             `xml:"host"`
	Id          string            `xml:"id,attr"`
	Name        string            `xml:"name"`
	Status      GlusterHookStatus `xml:"status"`
}

type GlusterVolumes struct {
	GlusterVolumes []GlusterVolume `xml:"gluster_volume"`
}

type GlusterVolume struct {
	OvStruct
	Bricks          []GlusterBrick      `xml:"bricks"`
	Cluster         *Cluster            `xml:"cluster"`
	Comment         string              `xml:"comment"`
	Description     string              `xml:"description"`
	DisperseCount   int64               `xml:"disperse_count"`
	Id              string              `xml:"id,attr"`
	Name            string              `xml:"name"`
	Options         []Option            `xml:"options"`
	RedundancyCount int64               `xml:"redundancy_count"`
	ReplicaCount    int64               `xml:"replica_count"`
	Statistics      []Statistic         `xml:"statistics"`
	Status          GlusterVolumeStatus `xml:"status"`
	StripeCount     int64               `xml:"stripe_count"`
	TransportTypes  []TransportType     `xml:"transport_types"`
	VolumeType      GlusterVolumeType   `xml:"volume_type"`
}

type GlusterVolumeProfileDetailss struct {
	GlusterVolumeProfileDetailss []GlusterVolumeProfileDetails `xml:"gluster_volume_profile_details"`
}

type GlusterVolumeProfileDetails struct {
	OvStruct
	BrickProfileDetails []BrickProfileDetail `xml:"brick_profile_details"`
	Comment             string               `xml:"comment"`
	Description         string               `xml:"description"`
	Id                  string               `xml:"id,attr"`
	Name                string               `xml:"name"`
	NfsProfileDetails   []NfsProfileDetail   `xml:"nfs_profile_details"`
}

type GraphicsConsoles struct {
	GraphicsConsoles []GraphicsConsole `xml:"graphics_console"`
}

type GraphicsConsole struct {
	OvStruct
	Address      string        `xml:"address"`
	Comment      string        `xml:"comment"`
	Description  string        `xml:"description"`
	Id           string        `xml:"id,attr"`
	InstanceType *InstanceType `xml:"instance_type"`
	Name         string        `xml:"name"`
	Port         int64         `xml:"port"`
	Protocol     GraphicsType  `xml:"protocol"`
	Template     *Template     `xml:"template"`
	TlsPort      int64         `xml:"tls_port"`
	Vm           *Vm           `xml:"vm"`
}

type Groups struct {
	Groups []Group `xml:"group"`
}

type Group struct {
	OvStruct
	Comment       string       `xml:"comment"`
	Description   string       `xml:"description"`
	Domain        *Domain      `xml:"domain"`
	DomainEntryId string       `xml:"domain_entry_id"`
	Id            string       `xml:"id,attr"`
	Name          string       `xml:"name"`
	Namespace     string       `xml:"namespace"`
	Permissions   []Permission `xml:"permissions"`
	Roles         []Role       `xml:"roles"`
	Tags          []Tag        `xml:"tags"`
}

type Hooks struct {
	Hooks []Hook `xml:"hook"`
}

type Hook struct {
	OvStruct
	Comment     string `xml:"comment"`
	Description string `xml:"description"`
	EventName   string `xml:"event_name"`
	Host        *Host  `xml:"host"`
	Id          string `xml:"id,attr"`
	Md5         string `xml:"md5"`
	Name        string `xml:"name"`
}

type Hosts struct {
	Hosts []Host `xml:"host"`
}

type Host struct {
	OvStruct
	Address                     string                       `xml:"address"`
	AffinityLabels              []AffinityLabel              `xml:"affinity_labels"`
	Agents                      []Agent                      `xml:"agents"`
	AutoNumaStatus              AutoNumaStatus               `xml:"auto_numa_status"`
	Certificate                 *Certificate                 `xml:"certificate"`
	Cluster                     *Cluster                     `xml:"cluster"`
	Comment                     string                       `xml:"comment"`
	Cpu                         *Cpu                         `xml:"cpu"`
	Description                 string                       `xml:"description"`
	DevicePassthrough           *HostDevicePassthrough       `xml:"device_passthrough"`
	Devices                     []Device                     `xml:"devices"`
	Display                     *Display                     `xml:"display"`
	ExternalHostProvider        *ExternalHostProvider        `xml:"external_host_provider"`
	ExternalStatus              ExternalStatus               `xml:"external_status"`
	HardwareInformation         *HardwareInformation         `xml:"hardware_information"`
	Hooks                       []Hook                       `xml:"hooks"`
	HostedEngine                *HostedEngine                `xml:"hosted_engine"`
	Id                          string                       `xml:"id,attr"`
	Iscsi                       *IscsiDetails                `xml:"iscsi"`
	KatelloErrata               []KatelloErratum             `xml:"katello_errata"`
	KdumpStatus                 KdumpStatus                  `xml:"kdump_status"`
	Ksm                         *Ksm                         `xml:"ksm"`
	LibvirtVersion              *Version                     `xml:"libvirt_version"`
	MaxSchedulingMemory         int64                        `xml:"max_scheduling_memory"`
	Memory                      int64                        `xml:"memory"`
	Name                        string                       `xml:"name"`
	NetworkAttachments          []NetworkAttachment          `xml:"network_attachments"`
	Nics                        []Nic                        `xml:"nics"`
	NumaNodes                   []NumaNode                   `xml:"numa_nodes"`
	NumaSupported               bool                         `xml:"numa_supported"`
	Os                          *OperatingSystem             `xml:"os"`
	OverrideIptables            bool                         `xml:"override_iptables"`
	Permissions                 []Permission                 `xml:"permissions"`
	Port                        int64                        `xml:"port"`
	PowerManagement             *PowerManagement             `xml:"power_management"`
	Protocol                    HostProtocol                 `xml:"protocol"`
	RootPassword                string                       `xml:"root_password"`
	SeLinux                     *SeLinux                     `xml:"se_linux"`
	Spm                         *Spm                         `xml:"spm"`
	Ssh                         *Ssh                         `xml:"ssh"`
	Statistics                  []Statistic                  `xml:"statistics"`
	Status                      HostStatus                   `xml:"status"`
	StatusDetail                string                       `xml:"status_detail"`
	StorageConnectionExtensions []StorageConnectionExtension `xml:"storage_connection_extensions"`
	Storages                    []HostStorage                `xml:"storages"`
	Summary                     *VmSummary                   `xml:"summary"`
	Tags                        []Tag                        `xml:"tags"`
	TransparentHugePages        *TransparentHugePages        `xml:"transparent_huge_pages"`
	Type                        HostType                     `xml:"type"`
	UnmanagedNetworks           []UnmanagedNetwork           `xml:"unmanaged_networks"`
	UpdateAvailable             bool                         `xml:"update_available"`
	Version                     *Version                     `xml:"version"`
}

type HostDevices struct {
	HostDevices []HostDevice `xml:"host_device"`
}

type HostDevice struct {
	OvStruct
	Capability       string      `xml:"capability"`
	Comment          string      `xml:"comment"`
	Description      string      `xml:"description"`
	Host             *Host       `xml:"host"`
	Id               string      `xml:"id,attr"`
	IommuGroup       int64       `xml:"iommu_group"`
	Name             string      `xml:"name"`
	ParentDevice     *HostDevice `xml:"parent_device"`
	PhysicalFunction *HostDevice `xml:"physical_function"`
	Placeholder      bool        `xml:"placeholder"`
	Product          *Product    `xml:"product"`
	Vendor           *Vendor     `xml:"vendor"`
	VirtualFunctions int64       `xml:"virtual_functions"`
	Vm               *Vm         `xml:"vm"`
}

type HostNics struct {
	HostNics []HostNic `xml:"host_nic"`
}

type HostNic struct {
	OvStruct
	AdAggregatorId                int64                                 `xml:"ad_aggregator_id"`
	BaseInterface                 string                                `xml:"base_interface"`
	Bonding                       *Bonding                              `xml:"bonding"`
	BootProtocol                  BootProtocol                          `xml:"boot_protocol"`
	Bridged                       bool                                  `xml:"bridged"`
	CheckConnectivity             bool                                  `xml:"check_connectivity"`
	Comment                       string                                `xml:"comment"`
	CustomConfiguration           bool                                  `xml:"custom_configuration"`
	Description                   string                                `xml:"description"`
	Host                          *Host                                 `xml:"host"`
	Id                            string                                `xml:"id,attr"`
	Ip                            *Ip                                   `xml:"ip"`
	Ipv6                          *Ip                                   `xml:"ipv6"`
	Ipv6BootProtocol              BootProtocol                          `xml:"ipv6_boot_protocol"`
	Mac                           *Mac                                  `xml:"mac"`
	Mtu                           int64                                 `xml:"mtu"`
	Name                          string                                `xml:"name"`
	Network                       *Network                              `xml:"network"`
	NetworkLabels                 []NetworkLabel                        `xml:"network_labels"`
	OverrideConfiguration         bool                                  `xml:"override_configuration"`
	PhysicalFunction              *HostNic                              `xml:"physical_function"`
	Properties                    []Property                            `xml:"properties"`
	Qos                           *Qos                                  `xml:"qos"`
	Speed                         int64                                 `xml:"speed"`
	Statistics                    []Statistic                           `xml:"statistics"`
	Status                        NicStatus                             `xml:"status"`
	VirtualFunctionsConfiguration *HostNicVirtualFunctionsConfiguration `xml:"virtual_functions_configuration"`
	Vlan                          *Vlan                                 `xml:"vlan"`
}

type HostStorages struct {
	HostStorages []HostStorage `xml:"host_storage"`
}

type HostStorage struct {
	OvStruct
	Address      string        `xml:"address"`
	Comment      string        `xml:"comment"`
	Description  string        `xml:"description"`
	Host         *Host         `xml:"host"`
	Id           string        `xml:"id,attr"`
	LogicalUnits []LogicalUnit `xml:"logical_units"`
	MountOptions string        `xml:"mount_options"`
	Name         string        `xml:"name"`
	NfsRetrans   int64         `xml:"nfs_retrans"`
	NfsTimeo     int64         `xml:"nfs_timeo"`
	NfsVersion   NfsVersion    `xml:"nfs_version"`
	OverrideLuns bool          `xml:"override_luns"`
	Password     string        `xml:"password"`
	Path         string        `xml:"path"`
	Port         int64         `xml:"port"`
	Portal       string        `xml:"portal"`
	Target       string        `xml:"target"`
	Type         StorageType   `xml:"type"`
	Username     string        `xml:"username"`
	VfsType      string        `xml:"vfs_type"`
	VolumeGroup  *VolumeGroup  `xml:"volume_group"`
}

type Icons struct {
	Icons []Icon `xml:"icon"`
}

type Icon struct {
	OvStruct
	Comment     string `xml:"comment"`
	Data        string `xml:"data"`
	Description string `xml:"description"`
	Id          string `xml:"id,attr"`
	MediaType   string `xml:"media_type"`
	Name        string `xml:"name"`
}

type Nics struct {
	Nics []Nic `xml:"nic"`
}

type Nic struct {
	OvStruct
	BootProtocol                   BootProtocol             `xml:"boot_protocol"`
	Comment                        string                   `xml:"comment"`
	Description                    string                   `xml:"description"`
	Id                             string                   `xml:"id,attr"`
	InstanceType                   *InstanceType            `xml:"instance_type"`
	Interface                      NicInterface             `xml:"interface"`
	Linked                         bool                     `xml:"linked"`
	Mac                            *Mac                     `xml:"mac"`
	Name                           string                   `xml:"name"`
	Network                        *Network                 `xml:"network"`
	NetworkAttachments             []NetworkAttachment      `xml:"network_attachments"`
	NetworkFilterParameters        []NetworkFilterParameter `xml:"network_filter_parameters"`
	NetworkLabels                  []NetworkLabel           `xml:"network_labels"`
	OnBoot                         bool                     `xml:"on_boot"`
	Plugged                        bool                     `xml:"plugged"`
	ReportedDevices                []ReportedDevice         `xml:"reported_devices"`
	Statistics                     []Statistic              `xml:"statistics"`
	Template                       *Template                `xml:"template"`
	VirtualFunctionAllowedLabels   []NetworkLabel           `xml:"virtual_function_allowed_labels"`
	VirtualFunctionAllowedNetworks []Network                `xml:"virtual_function_allowed_networks"`
	Vm                             *Vm                      `xml:"vm"`
	Vms                            []Vm                     `xml:"vms"`
	VnicProfile                    *VnicProfile             `xml:"vnic_profile"`
}

type OpenStackProviders struct {
	OpenStackProviders []OpenStackProvider `xml:"open_stack_provider"`
}

type OpenStackProvider struct {
	OvStruct
	AuthenticationUrl      string     `xml:"authentication_url"`
	Comment                string     `xml:"comment"`
	Description            string     `xml:"description"`
	Id                     string     `xml:"id,attr"`
	Name                   string     `xml:"name"`
	Password               string     `xml:"password"`
	Properties             []Property `xml:"properties"`
	RequiresAuthentication bool       `xml:"requires_authentication"`
	TenantName             string     `xml:"tenant_name"`
	Url                    string     `xml:"url"`
	Username               string     `xml:"username"`
}

type OpenStackVolumeProviders struct {
	OpenStackVolumeProviders []OpenStackVolumeProvider `xml:"open_stack_volume_provider"`
}

type OpenStackVolumeProvider struct {
	OvStruct
	AuthenticationKeys     []OpenstackVolumeAuthenticationKey `xml:"authentication_keys"`
	AuthenticationUrl      string                             `xml:"authentication_url"`
	Certificates           []Certificate                      `xml:"certificates"`
	Comment                string                             `xml:"comment"`
	DataCenter             *DataCenter                        `xml:"data_center"`
	Description            string                             `xml:"description"`
	Id                     string                             `xml:"id,attr"`
	Name                   string                             `xml:"name"`
	Password               string                             `xml:"password"`
	Properties             []Property                         `xml:"properties"`
	RequiresAuthentication bool                               `xml:"requires_authentication"`
	TenantName             string                             `xml:"tenant_name"`
	Url                    string                             `xml:"url"`
	Username               string                             `xml:"username"`
	VolumeTypes            []OpenStackVolumeType              `xml:"volume_types"`
}

type Templates struct {
	Templates []Template `xml:"template"`
}

type Template struct {
	OvStruct
	Bios                       *Bios               `xml:"bios"`
	Cdroms                     []Cdrom             `xml:"cdroms"`
	Cluster                    *Cluster            `xml:"cluster"`
	Comment                    string              `xml:"comment"`
	Console                    *Console            `xml:"console"`
	Cpu                        *Cpu                `xml:"cpu"`
	CpuProfile                 *CpuProfile         `xml:"cpu_profile"`
	CpuShares                  int64               `xml:"cpu_shares"`
	CreationTime               time.Time           `xml:"creation_time"`
	CustomCompatibilityVersion *Version            `xml:"custom_compatibility_version"`
	CustomCpuModel             string              `xml:"custom_cpu_model"`
	CustomEmulatedMachine      string              `xml:"custom_emulated_machine"`
	CustomProperties           []CustomProperty    `xml:"custom_properties"`
	DeleteProtected            bool                `xml:"delete_protected"`
	Description                string              `xml:"description"`
	DiskAttachments            []DiskAttachment    `xml:"disk_attachments"`
	Display                    *Display            `xml:"display"`
	Domain                     *Domain             `xml:"domain"`
	GraphicsConsoles           []GraphicsConsole   `xml:"graphics_consoles"`
	HighAvailability           *HighAvailability   `xml:"high_availability"`
	Id                         string              `xml:"id,attr"`
	Initialization             *Initialization     `xml:"initialization"`
	Io                         *Io                 `xml:"io"`
	LargeIcon                  *Icon               `xml:"large_icon"`
	Lease                      *StorageDomainLease `xml:"lease"`
	Memory                     int64               `xml:"memory"`
	MemoryPolicy               *MemoryPolicy       `xml:"memory_policy"`
	Migration                  *MigrationOptions   `xml:"migration"`
	MigrationDowntime          int64               `xml:"migration_downtime"`
	Name                       string              `xml:"name"`
	Nics                       []Nic               `xml:"nics"`
	Origin                     string              `xml:"origin"`
	Os                         *OperatingSystem    `xml:"os"`
	Permissions                []Permission        `xml:"permissions"`
	Quota                      *Quota              `xml:"quota"`
	RngDevice                  *RngDevice          `xml:"rng_device"`
	SerialNumber               *SerialNumber       `xml:"serial_number"`
	SmallIcon                  *Icon               `xml:"small_icon"`
	SoundcardEnabled           bool                `xml:"soundcard_enabled"`
	Sso                        *Sso                `xml:"sso"`
	StartPaused                bool                `xml:"start_paused"`
	Stateless                  bool                `xml:"stateless"`
	Status                     TemplateStatus      `xml:"status"`
	StorageDomain              *StorageDomain      `xml:"storage_domain"`
	Tags                       []Tag               `xml:"tags"`
	TimeZone                   *TimeZone           `xml:"time_zone"`
	TunnelMigration            bool                `xml:"tunnel_migration"`
	Type                       VmType              `xml:"type"`
	Usb                        *Usb                `xml:"usb"`
	Version                    *TemplateVersion    `xml:"version"`
	VirtioScsi                 *VirtioScsi         `xml:"virtio_scsi"`
	Vm                         *Vm                 `xml:"vm"`
	Watchdogs                  []Watchdog          `xml:"watchdogs"`
}

type Vms struct {
	Vms []Vm `xml:"vm"`
}

type Vm struct {
	OvStruct
	AffinityLabels             []AffinityLabel       `xml:"affinity_labels"`
	Applications               []Application         `xml:"applications"`
	Bios                       *Bios                 `xml:"bios"`
	Cdroms                     []Cdrom               `xml:"cdroms"`
	Cluster                    *Cluster              `xml:"cluster"`
	Comment                    string                `xml:"comment"`
	Console                    *Console              `xml:"console"`
	Cpu                        *Cpu                  `xml:"cpu"`
	CpuProfile                 *CpuProfile           `xml:"cpu_profile"`
	CpuShares                  int64                 `xml:"cpu_shares"`
	CreationTime               time.Time             `xml:"creation_time"`
	CustomCompatibilityVersion *Version              `xml:"custom_compatibility_version"`
	CustomCpuModel             string                `xml:"custom_cpu_model"`
	CustomEmulatedMachine      string                `xml:"custom_emulated_machine"`
	CustomProperties           []CustomProperty      `xml:"custom_properties"`
	DeleteProtected            bool                  `xml:"delete_protected"`
	Description                string                `xml:"description"`
	DiskAttachments            []DiskAttachment      `xml:"disk_attachments"`
	Display                    *Display              `xml:"display"`
	Domain                     *Domain               `xml:"domain"`
	ExternalHostProvider       *ExternalHostProvider `xml:"external_host_provider"`
	Floppies                   []Floppy              `xml:"floppies"`
	Fqdn                       string                `xml:"fqdn"`
	GraphicsConsoles           []GraphicsConsole     `xml:"graphics_consoles"`
	GuestOperatingSystem       *GuestOperatingSystem `xml:"guest_operating_system"`
	GuestTimeZone              *TimeZone             `xml:"guest_time_zone"`
	HighAvailability           *HighAvailability     `xml:"high_availability"`
	Host                       *Host                 `xml:"host"`
	HostDevices                []HostDevice          `xml:"host_devices"`
	Id                         string                `xml:"id,attr"`
	Initialization             *Initialization       `xml:"initialization"`
	InstanceType               *InstanceType         `xml:"instance_type"`
	Io                         *Io                   `xml:"io"`
	KatelloErrata              []KatelloErratum      `xml:"katello_errata"`
	LargeIcon                  *Icon                 `xml:"large_icon"`
	Lease                      *StorageDomainLease   `xml:"lease"`
	Memory                     int64                 `xml:"memory"`
	MemoryPolicy               *MemoryPolicy         `xml:"memory_policy"`
	Migration                  *MigrationOptions     `xml:"migration"`
	MigrationDowntime          int64                 `xml:"migration_downtime"`
	Name                       string                `xml:"name"`
	NextRunConfigurationExists bool                  `xml:"next_run_configuration_exists"`
	Nics                       []Nic                 `xml:"nics"`
	NumaNodes                  []NumaNode            `xml:"numa_nodes"`
	NumaTuneMode               NumaTuneMode          `xml:"numa_tune_mode"`
	Origin                     string                `xml:"origin"`
	OriginalTemplate           *Template             `xml:"original_template"`
	Os                         *OperatingSystem      `xml:"os"`
	Payloads                   []Payload             `xml:"payloads"`
	Permissions                []Permission          `xml:"permissions"`
	PlacementPolicy            *VmPlacementPolicy    `xml:"placement_policy"`
	Quota                      *Quota                `xml:"quota"`
	ReportedDevices            []ReportedDevice      `xml:"reported_devices"`
	RngDevice                  *RngDevice            `xml:"rng_device"`
	RunOnce                    bool                  `xml:"run_once"`
	SerialNumber               *SerialNumber         `xml:"serial_number"`
	Sessions                   []Session             `xml:"sessions"`
	SmallIcon                  *Icon                 `xml:"small_icon"`
	Snapshots                  []Snapshot            `xml:"snapshots"`
	SoundcardEnabled           bool                  `xml:"soundcard_enabled"`
	Sso                        *Sso                  `xml:"sso"`
	StartPaused                bool                  `xml:"start_paused"`
	StartTime                  time.Time             `xml:"start_time"`
	Stateless                  bool                  `xml:"stateless"`
	Statistics                 []Statistic           `xml:"statistics"`
	Status                     VmStatus              `xml:"status"`
	StatusDetail               string                `xml:"status_detail"`
	StopReason                 string                `xml:"stop_reason"`
	StopTime                   time.Time             `xml:"stop_time"`
	StorageDomain              *StorageDomain        `xml:"storage_domain"`
	Tags                       []Tag                 `xml:"tags"`
	Template                   *Template             `xml:"template"`
	TimeZone                   *TimeZone             `xml:"time_zone"`
	TunnelMigration            bool                  `xml:"tunnel_migration"`
	Type                       VmType                `xml:"type"`
	Usb                        *Usb                  `xml:"usb"`
	UseLatestTemplateVersion   bool                  `xml:"use_latest_template_version"`
	VirtioScsi                 *VirtioScsi           `xml:"virtio_scsi"`
	VmPool                     *VmPool               `xml:"vm_pool"`
	Watchdogs                  []Watchdog            `xml:"watchdogs"`
}

type Watchdogs struct {
	Watchdogs []Watchdog `xml:"watchdog"`
}

type Watchdog struct {
	OvStruct
	Action       WatchdogAction `xml:"action"`
	Comment      string         `xml:"comment"`
	Description  string         `xml:"description"`
	Id           string         `xml:"id,attr"`
	InstanceType *InstanceType  `xml:"instance_type"`
	Model        WatchdogModel  `xml:"model"`
	Name         string         `xml:"name"`
	Template     *Template      `xml:"template"`
	Vm           *Vm            `xml:"vm"`
	Vms          []Vm           `xml:"vms"`
}

type Cdroms struct {
	Cdroms []Cdrom `xml:"cdrom"`
}

type Cdrom struct {
	OvStruct
	Comment      string        `xml:"comment"`
	Description  string        `xml:"description"`
	File         *File         `xml:"file"`
	Id           string        `xml:"id,attr"`
	InstanceType *InstanceType `xml:"instance_type"`
	Name         string        `xml:"name"`
	Template     *Template     `xml:"template"`
	Vm           *Vm           `xml:"vm"`
	Vms          []Vm          `xml:"vms"`
}

type ExternalHostProviders struct {
	ExternalHostProviders []ExternalHostProvider `xml:"external_host_provider"`
}

type ExternalHostProvider struct {
	OvStruct
	AuthenticationUrl      string                    `xml:"authentication_url"`
	Certificates           []Certificate             `xml:"certificates"`
	Comment                string                    `xml:"comment"`
	ComputeResources       []ExternalComputeResource `xml:"compute_resources"`
	Description            string                    `xml:"description"`
	DiscoveredHosts        []ExternalDiscoveredHost  `xml:"discovered_hosts"`
	HostGroups             []ExternalHostGroup       `xml:"host_groups"`
	Hosts                  []Host                    `xml:"hosts"`
	Id                     string                    `xml:"id,attr"`
	Name                   string                    `xml:"name"`
	Password               string                    `xml:"password"`
	Properties             []Property                `xml:"properties"`
	RequiresAuthentication bool                      `xml:"requires_authentication"`
	Url                    string                    `xml:"url"`
	Username               string                    `xml:"username"`
}

type GlusterBricks struct {
	GlusterBricks []GlusterBrick `xml:"gluster_brick"`
}

type GlusterBrick struct {
	OvStruct
	BrickDir       string              `xml:"brick_dir"`
	Comment        string              `xml:"comment"`
	Description    string              `xml:"description"`
	Device         string              `xml:"device"`
	FsName         string              `xml:"fs_name"`
	GlusterClients []GlusterClient     `xml:"gluster_clients"`
	GlusterVolume  *GlusterVolume      `xml:"gluster_volume"`
	Id             string              `xml:"id,attr"`
	InstanceType   *InstanceType       `xml:"instance_type"`
	MemoryPools    []GlusterMemoryPool `xml:"memory_pools"`
	MntOptions     string              `xml:"mnt_options"`
	Name           string              `xml:"name"`
	Pid            int64               `xml:"pid"`
	Port           int64               `xml:"port"`
	ServerId       string              `xml:"server_id"`
	Statistics     []Statistic         `xml:"statistics"`
	Status         GlusterBrickStatus  `xml:"status"`
	Template       *Template           `xml:"template"`
	Vm             *Vm                 `xml:"vm"`
	Vms            []Vm                `xml:"vms"`
}

type InstanceTypes struct {
	InstanceTypes []InstanceType `xml:"instance_type"`
}

type InstanceType struct {
	OvStruct
	Bios                       *Bios               `xml:"bios"`
	Cdroms                     []Cdrom             `xml:"cdroms"`
	Cluster                    *Cluster            `xml:"cluster"`
	Comment                    string              `xml:"comment"`
	Console                    *Console            `xml:"console"`
	Cpu                        *Cpu                `xml:"cpu"`
	CpuProfile                 *CpuProfile         `xml:"cpu_profile"`
	CpuShares                  int64               `xml:"cpu_shares"`
	CreationTime               time.Time           `xml:"creation_time"`
	CustomCompatibilityVersion *Version            `xml:"custom_compatibility_version"`
	CustomCpuModel             string              `xml:"custom_cpu_model"`
	CustomEmulatedMachine      string              `xml:"custom_emulated_machine"`
	CustomProperties           []CustomProperty    `xml:"custom_properties"`
	DeleteProtected            bool                `xml:"delete_protected"`
	Description                string              `xml:"description"`
	DiskAttachments            []DiskAttachment    `xml:"disk_attachments"`
	Display                    *Display            `xml:"display"`
	Domain                     *Domain             `xml:"domain"`
	GraphicsConsoles           []GraphicsConsole   `xml:"graphics_consoles"`
	HighAvailability           *HighAvailability   `xml:"high_availability"`
	Id                         string              `xml:"id,attr"`
	Initialization             *Initialization     `xml:"initialization"`
	Io                         *Io                 `xml:"io"`
	LargeIcon                  *Icon               `xml:"large_icon"`
	Lease                      *StorageDomainLease `xml:"lease"`
	Memory                     int64               `xml:"memory"`
	MemoryPolicy               *MemoryPolicy       `xml:"memory_policy"`
	Migration                  *MigrationOptions   `xml:"migration"`
	MigrationDowntime          int64               `xml:"migration_downtime"`
	Name                       string              `xml:"name"`
	Nics                       []Nic               `xml:"nics"`
	Origin                     string              `xml:"origin"`
	Os                         *OperatingSystem    `xml:"os"`
	Permissions                []Permission        `xml:"permissions"`
	Quota                      *Quota              `xml:"quota"`
	RngDevice                  *RngDevice          `xml:"rng_device"`
	SerialNumber               *SerialNumber       `xml:"serial_number"`
	SmallIcon                  *Icon               `xml:"small_icon"`
	SoundcardEnabled           bool                `xml:"soundcard_enabled"`
	Sso                        *Sso                `xml:"sso"`
	StartPaused                bool                `xml:"start_paused"`
	Stateless                  bool                `xml:"stateless"`
	Status                     TemplateStatus      `xml:"status"`
	StorageDomain              *StorageDomain      `xml:"storage_domain"`
	Tags                       []Tag               `xml:"tags"`
	TimeZone                   *TimeZone           `xml:"time_zone"`
	TunnelMigration            bool                `xml:"tunnel_migration"`
	Type                       VmType              `xml:"type"`
	Usb                        *Usb                `xml:"usb"`
	Version                    *TemplateVersion    `xml:"version"`
	VirtioScsi                 *VirtioScsi         `xml:"virtio_scsi"`
	Vm                         *Vm                 `xml:"vm"`
	Watchdogs                  []Watchdog          `xml:"watchdogs"`
}

type OpenStackImageProviders struct {
	OpenStackImageProviders []OpenStackImageProvider `xml:"open_stack_image_provider"`
}

type OpenStackImageProvider struct {
	OvStruct
	AuthenticationUrl      string           `xml:"authentication_url"`
	Certificates           []Certificate    `xml:"certificates"`
	Comment                string           `xml:"comment"`
	Description            string           `xml:"description"`
	Id                     string           `xml:"id,attr"`
	Images                 []OpenStackImage `xml:"images"`
	Name                   string           `xml:"name"`
	Password               string           `xml:"password"`
	Properties             []Property       `xml:"properties"`
	RequiresAuthentication bool             `xml:"requires_authentication"`
	TenantName             string           `xml:"tenant_name"`
	Url                    string           `xml:"url"`
	Username               string           `xml:"username"`
}

type OpenStackNetworkProviders struct {
	OpenStackNetworkProviders []OpenStackNetworkProvider `xml:"open_stack_network_provider"`
}

type OpenStackNetworkProvider struct {
	OvStruct
	AgentConfiguration     *AgentConfiguration          `xml:"agent_configuration"`
	AuthenticationUrl      string                       `xml:"authentication_url"`
	Certificates           []Certificate                `xml:"certificates"`
	Comment                string                       `xml:"comment"`
	Description            string                       `xml:"description"`
	Id                     string                       `xml:"id,attr"`
	Name                   string                       `xml:"name"`
	Networks               []OpenStackNetwork           `xml:"networks"`
	Password               string                       `xml:"password"`
	PluginType             NetworkPluginType            `xml:"plugin_type"`
	Properties             []Property                   `xml:"properties"`
	ReadOnly               bool                         `xml:"read_only"`
	RequiresAuthentication bool                         `xml:"requires_authentication"`
	Subnets                []OpenStackSubnet            `xml:"subnets"`
	TenantName             string                       `xml:"tenant_name"`
	Type                   OpenStackNetworkProviderType `xml:"type"`
	Url                    string                       `xml:"url"`
	Username               string                       `xml:"username"`
}

type Snapshots struct {
	Snapshots []Snapshot `xml:"snapshot"`
}

type Snapshot struct {
	OvStruct
	AffinityLabels             []AffinityLabel       `xml:"affinity_labels"`
	Applications               []Application         `xml:"applications"`
	Bios                       *Bios                 `xml:"bios"`
	Cdroms                     []Cdrom               `xml:"cdroms"`
	Cluster                    *Cluster              `xml:"cluster"`
	Comment                    string                `xml:"comment"`
	Console                    *Console              `xml:"console"`
	Cpu                        *Cpu                  `xml:"cpu"`
	CpuProfile                 *CpuProfile           `xml:"cpu_profile"`
	CpuShares                  int64                 `xml:"cpu_shares"`
	CreationTime               time.Time             `xml:"creation_time"`
	CustomCompatibilityVersion *Version              `xml:"custom_compatibility_version"`
	CustomCpuModel             string                `xml:"custom_cpu_model"`
	CustomEmulatedMachine      string                `xml:"custom_emulated_machine"`
	CustomProperties           []CustomProperty      `xml:"custom_properties"`
	Date                       time.Time             `xml:"date"`
	DeleteProtected            bool                  `xml:"delete_protected"`
	Description                string                `xml:"description"`
	DiskAttachments            []DiskAttachment      `xml:"disk_attachments"`
	Display                    *Display              `xml:"display"`
	Domain                     *Domain               `xml:"domain"`
	ExternalHostProvider       *ExternalHostProvider `xml:"external_host_provider"`
	Floppies                   []Floppy              `xml:"floppies"`
	Fqdn                       string                `xml:"fqdn"`
	GraphicsConsoles           []GraphicsConsole     `xml:"graphics_consoles"`
	GuestOperatingSystem       *GuestOperatingSystem `xml:"guest_operating_system"`
	GuestTimeZone              *TimeZone             `xml:"guest_time_zone"`
	HighAvailability           *HighAvailability     `xml:"high_availability"`
	Host                       *Host                 `xml:"host"`
	HostDevices                []HostDevice          `xml:"host_devices"`
	Id                         string                `xml:"id,attr"`
	Initialization             *Initialization       `xml:"initialization"`
	InstanceType               *InstanceType         `xml:"instance_type"`
	Io                         *Io                   `xml:"io"`
	KatelloErrata              []KatelloErratum      `xml:"katello_errata"`
	LargeIcon                  *Icon                 `xml:"large_icon"`
	Lease                      *StorageDomainLease   `xml:"lease"`
	Memory                     int64                 `xml:"memory"`
	MemoryPolicy               *MemoryPolicy         `xml:"memory_policy"`
	Migration                  *MigrationOptions     `xml:"migration"`
	MigrationDowntime          int64                 `xml:"migration_downtime"`
	Name                       string                `xml:"name"`
	NextRunConfigurationExists bool                  `xml:"next_run_configuration_exists"`
	Nics                       []Nic                 `xml:"nics"`
	NumaNodes                  []NumaNode            `xml:"numa_nodes"`
	NumaTuneMode               NumaTuneMode          `xml:"numa_tune_mode"`
	Origin                     string                `xml:"origin"`
	OriginalTemplate           *Template             `xml:"original_template"`
	Os                         *OperatingSystem      `xml:"os"`
	Payloads                   []Payload             `xml:"payloads"`
	Permissions                []Permission          `xml:"permissions"`
	PersistMemorystate         bool                  `xml:"persist_memorystate"`
	PlacementPolicy            *VmPlacementPolicy    `xml:"placement_policy"`
	Quota                      *Quota                `xml:"quota"`
	ReportedDevices            []ReportedDevice      `xml:"reported_devices"`
	RngDevice                  *RngDevice            `xml:"rng_device"`
	RunOnce                    bool                  `xml:"run_once"`
	SerialNumber               *SerialNumber         `xml:"serial_number"`
	Sessions                   []Session             `xml:"sessions"`
	SmallIcon                  *Icon                 `xml:"small_icon"`
	SnapshotStatus             SnapshotStatus        `xml:"snapshot_status"`
	SnapshotType               SnapshotType          `xml:"snapshot_type"`
	Snapshots                  []Snapshot            `xml:"snapshots"`
	SoundcardEnabled           bool                  `xml:"soundcard_enabled"`
	Sso                        *Sso                  `xml:"sso"`
	StartPaused                bool                  `xml:"start_paused"`
	StartTime                  time.Time             `xml:"start_time"`
	Stateless                  bool                  `xml:"stateless"`
	Statistics                 []Statistic           `xml:"statistics"`
	Status                     VmStatus              `xml:"status"`
	StatusDetail               string                `xml:"status_detail"`
	StopReason                 string                `xml:"stop_reason"`
	StopTime                   time.Time             `xml:"stop_time"`
	StorageDomain              *StorageDomain        `xml:"storage_domain"`
	Tags                       []Tag                 `xml:"tags"`
	Template                   *Template             `xml:"template"`
	TimeZone                   *TimeZone             `xml:"time_zone"`
	TunnelMigration            bool                  `xml:"tunnel_migration"`
	Type                       VmType                `xml:"type"`
	Usb                        *Usb                  `xml:"usb"`
	UseLatestTemplateVersion   bool                  `xml:"use_latest_template_version"`
	VirtioScsi                 *VirtioScsi           `xml:"virtio_scsi"`
	Vm                         *Vm                   `xml:"vm"`
	VmPool                     *VmPool               `xml:"vm_pool"`
	Watchdogs                  []Watchdog            `xml:"watchdogs"`
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
