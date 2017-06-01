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

type AffinityRule struct {
    OvType
    Enabled    bool
    Enforcing    bool
    Positive    bool

}

type AgentConfiguration struct {
    OvType
    Address    string
    BrokerType    MessageBrokerType
    NetworkMappings    string
    Password    string
    Port    int64
    Username    string

}

type Api struct {
    OvType
    ProductInfo    ProductInfo
    SpecialObjects    SpecialObjects
    Summary    ApiSummary
    Time    time.Time

}

type ApiSummary struct {
    OvType
    Hosts    ApiSummaryItem
    StorageDomains    ApiSummaryItem
    Users    ApiSummaryItem
    Vms    ApiSummaryItem

}

type ApiSummaryItem struct {
    OvType
    Active    int64
    Total    int64

}

type Bios struct {
    OvType
    BootMenu    BootMenu

}

type BlockStatistic struct {
    OvType
    Statistics    []*Statistic

}

type Bonding struct {
    OvType
    ActiveSlave    HostNic
    AdPartnerMac    Mac
    Options    []*Option
    Slaves    []*HostNic

}

type Boot struct {
    OvType
    Devices    []*BootDevice

}

type BootMenu struct {
    OvType
    Enabled    bool

}

type CloudInit struct {
    OvType
    AuthorizedKeys    []*AuthorizedKey
    Files    []*File
    Host    Host
    NetworkConfiguration    NetworkConfiguration
    RegenerateSshKeys    bool
    Timezone    string
    Users    []*User

}

type Configuration struct {
    OvType
    Data    string
    Type    ConfigurationType

}

type Console struct {
    OvType
    Enabled    bool

}

type Core struct {
    OvType
    Index    int64
    Socket    int64

}

type Cpu struct {
    OvType
    Architecture    Architecture
    Cores    []*Core
    CpuTune    CpuTune
    Level    int64
    Mode    CpuMode
    Name    string
    Speed    float64
    Topology    CpuTopology
    Type    string

}

type CpuTopology struct {
    OvType
    Cores    int64
    Sockets    int64
    Threads    int64

}

type CpuTune struct {
    OvType
    VcpuPins    []*VcpuPin

}

type CpuType struct {
    OvType
    Architecture    Architecture
    Level    int64
    Name    string

}

type CustomProperty struct {
    OvType
    Name    string
    Regexp    string
    Value    string

}

type Display struct {
    OvType
    Address    string
    AllowOverride    bool
    Certificate    Certificate
    CopyPasteEnabled    bool
    DisconnectAction    string
    FileTransferEnabled    bool
    KeyboardLayout    string
    Monitors    int64
    Port    int64
    Proxy    string
    SecurePort    int64
    SingleQxlPci    bool
    SmartcardEnabled    bool
    Type    DisplayType

}

type Dns struct {
    OvType
    SearchDomains    []*Host
    Servers    []*Host

}

type DnsResolverConfiguration struct {
    OvType
    NameServers    []*String

}

type EntityProfileDetail struct {
    OvType
    ProfileDetails    []*ProfileDetail

}

type ErrorHandling struct {
    OvType
    OnError    MigrateOnError

}

type ExternalVmImport struct {
    OvType
    Cluster    Cluster
    CpuProfile    CpuProfile
    DriversIso    File
    Host    Host
    Name    string
    Password    string
    Provider    ExternalVmProviderType
    Quota    Quota
    Sparse    bool
    StorageDomain    StorageDomain
    Url    string
    Username    string
    Vm    Vm

}

type Fault struct {
    OvType
    Detail    string
    Reason    string

}

type FencingPolicy struct {
    OvType
    Enabled    bool
    SkipIfConnectivityBroken    SkipIfConnectivityBroken
    SkipIfGlusterBricksUp    bool
    SkipIfGlusterQuorumNotMet    bool
    SkipIfSdActive    SkipIfSdActive

}

type FopStatistic struct {
    OvType
    Name    string
    Statistics    []*Statistic

}

type GlusterBrickMemoryInfo struct {
    OvType
    MemoryPools    []*GlusterMemoryPool

}

type GlusterClient struct {
    OvType
    BytesRead    int64
    BytesWritten    int64
    ClientPort    int64
    HostName    string

}

type GracePeriod struct {
    OvType
    Expiry    int64

}

type GuestOperatingSystem struct {
    OvType
    Architecture    string
    Codename    string
    Distribution    string
    Family    string
    Kernel    Kernel
    Version    Version

}

type HardwareInformation struct {
    OvType
    Family    string
    Manufacturer    string
    ProductName    string
    SerialNumber    string
    SupportedRngSources    []*RngSource
    Uuid    string
    Version    string

}

type HighAvailability struct {
    OvType
    Enabled    bool
    Priority    int64

}

type HostDevicePassthrough struct {
    OvType
    Enabled    bool

}

type HostNicVirtualFunctionsConfiguration struct {
    OvType
    AllNetworksAllowed    bool
    MaxNumberOfVirtualFunctions    int64
    NumberOfVirtualFunctions    int64

}

type HostedEngine struct {
    OvType
    Active    bool
    Configured    bool
    GlobalMaintenance    bool
    LocalMaintenance    bool
    Score    int64

}

type Identified struct {
    OvType
    Comment    string
    Description    string
    Id    string
    Name    string

}

type Image struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    StorageDomain    StorageDomain

}

type ImageTransfer struct {
    Identified
    Comment    string
    Description    string
    Direction    ImageTransferDirection
    Host    Host
    Id    string
    Image    Image
    Name    string
    Phase    ImageTransferPhase
    ProxyUrl    string
    SignedTicket    string

}

type Initialization struct {
    OvType
    ActiveDirectoryOu    string
    AuthorizedSshKeys    string
    CloudInit    CloudInit
    Configuration    Configuration
    CustomScript    string
    DnsSearch    string
    DnsServers    string
    Domain    string
    HostName    string
    InputLocale    string
    NicConfigurations    []*NicConfiguration
    OrgName    string
    RegenerateIds    bool
    RegenerateSshKeys    bool
    RootPassword    string
    SystemLocale    string
    Timezone    string
    UiLanguage    string
    UserLocale    string
    UserName    string
    WindowsLicenseKey    string

}

type Io struct {
    OvType
    Threads    int64

}

type Ip struct {
    OvType
    Address    string
    Gateway    string
    Netmask    string
    Version    IpVersion

}

type IpAddressAssignment struct {
    OvType
    AssignmentMethod    BootProtocol
    Ip    Ip

}

type IscsiBond struct {
    Identified
    Comment    string
    DataCenter    DataCenter
    Description    string
    Id    string
    Name    string
    Networks    []*Network
    StorageConnections    []*StorageConnection

}

type IscsiDetails struct {
    OvType
    Address    string
    DiskId    string
    Initiator    string
    LunMapping    int64
    Password    string
    Paths    int64
    Port    int64
    Portal    string
    ProductId    string
    Serial    string
    Size    int64
    Status    string
    StorageDomainId    string
    Target    string
    Username    string
    VendorId    string
    VolumeGroupId    string

}

type Job struct {
    Identified
    AutoCleared    bool
    Comment    string
    Description    string
    EndTime    time.Time
    External    bool
    Id    string
    LastUpdated    time.Time
    Name    string
    Owner    User
    StartTime    time.Time
    Status    JobStatus
    Steps    []*Step

}

type KatelloErratum struct {
    Identified
    Comment    string
    Description    string
    Host    Host
    Id    string
    Issued    time.Time
    Name    string
    Packages    []*Package
    Severity    string
    Solution    string
    Summary    string
    Title    string
    Type    string
    Vm    Vm

}

type Kernel struct {
    OvType
    Version    Version

}

type Ksm struct {
    OvType
    Enabled    bool
    MergeAcrossNodes    bool

}

type LogicalUnit struct {
    OvType
    Address    string
    DiscardMaxSize    int64
    DiscardZeroesData    bool
    DiskId    string
    Id    string
    LunMapping    int64
    Password    string
    Paths    int64
    Port    int64
    Portal    string
    ProductId    string
    Serial    string
    Size    int64
    Status    LunStatus
    StorageDomainId    string
    Target    string
    Username    string
    VendorId    string
    VolumeGroupId    string

}

type Mac struct {
    OvType
    Address    string

}

type MacPool struct {
    Identified
    AllowDuplicates    bool
    Comment    string
    DefaultPool    bool
    Description    string
    Id    string
    Name    string
    Ranges    []*Range

}

type MemoryOverCommit struct {
    OvType
    Percent    int64

}

type MemoryPolicy struct {
    OvType
    Ballooning    bool
    Guaranteed    int64
    Max    int64
    OverCommit    MemoryOverCommit
    TransparentHugePages    TransparentHugePages

}

type Method struct {
    OvType
    Id    SsoMethod

}

type MigrationBandwidth struct {
    OvType
    AssignmentMethod    MigrationBandwidthAssignmentMethod
    CustomValue    int64

}

type MigrationOptions struct {
    OvType
    AutoConverge    InheritableBoolean
    Bandwidth    MigrationBandwidth
    Compressed    InheritableBoolean
    Policy    MigrationPolicy

}

type MigrationPolicy struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string

}

type Network struct {
    Identified
    Cluster    Cluster
    Comment    string
    DataCenter    DataCenter
    Description    string
    Display    bool
    DnsResolverConfiguration    DnsResolverConfiguration
    Id    string
    Ip    Ip
    Mtu    int64
    Name    string
    NetworkLabels    []*NetworkLabel
    Permissions    []*Permission
    ProfileRequired    bool
    Qos    Qos
    Required    bool
    Status    NetworkStatus
    Stp    bool
    Usages    []*NetworkUsage
    Vlan    Vlan
    VnicProfiles    []*VnicProfile

}

type NetworkAttachment struct {
    Identified
    Comment    string
    Description    string
    DnsResolverConfiguration    DnsResolverConfiguration
    Host    Host
    HostNic    HostNic
    Id    string
    InSync    bool
    IpAddressAssignments    []*IpAddressAssignment
    Name    string
    Network    Network
    Properties    []*Property
    Qos    Qos
    ReportedConfigurations    []*ReportedConfiguration

}

type NetworkConfiguration struct {
    OvType
    Dns    Dns
    Nics    []*Nic

}

type NetworkFilter struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    Version    Version

}

type NetworkFilterParameter struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    Value    string

}

type NetworkLabel struct {
    Identified
    Comment    string
    Description    string
    HostNic    HostNic
    Id    string
    Name    string
    Network    Network

}

type NfsProfileDetail struct {
    EntityProfileDetail
    NfsServerIp    string
    ProfileDetails    []*ProfileDetail

}

type NicConfiguration struct {
    OvType
    BootProtocol    BootProtocol
    Ip    Ip
    Ipv6    Ip
    Ipv6BootProtocol    BootProtocol
    Name    string
    OnBoot    bool

}

type NumaNode struct {
    Identified
    Comment    string
    Cpu    Cpu
    Description    string
    Host    Host
    Id    string
    Index    int64
    Memory    int64
    Name    string
    NodeDistance    string
    Statistics    []*Statistic

}

type NumaNodePin struct {
    OvType
    HostNumaNode    NumaNode
    Index    int64
    Pinned    bool

}

type OpenStackImage struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    OpenstackImageProvider    OpenStackImageProvider

}

type OpenStackNetwork struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    OpenstackNetworkProvider    OpenStackNetworkProvider

}

type OpenStackSubnet struct {
    Identified
    Cidr    string
    Comment    string
    Description    string
    DnsServers    []*String
    Gateway    string
    Id    string
    IpVersion    string
    Name    string
    OpenstackNetwork    OpenStackNetwork

}

type OpenStackVolumeType struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    OpenstackVolumeProvider    OpenStackVolumeProvider
    Properties    []*Property

}

type OpenstackVolumeAuthenticationKey struct {
    Identified
    Comment    string
    CreationDate    time.Time
    Description    string
    Id    string
    Name    string
    OpenstackVolumeProvider    OpenStackVolumeProvider
    UsageType    OpenstackVolumeAuthenticationKeyUsageType
    Uuid    string
    Value    string

}

type OperatingSystem struct {
    OvType
    Boot    Boot
    Cmdline    string
    CustomKernelCmdline    string
    Initrd    string
    Kernel    string
    ReportedKernelCmdline    string
    Type    string
    Version    Version

}

type OperatingSystemInfo struct {
    Identified
    Comment    string
    Description    string
    Id    string
    LargeIcon    Icon
    Name    string
    SmallIcon    Icon

}

type Option struct {
    OvType
    Name    string
    Type    string
    Value    string

}

type Package struct {
    OvType
    Name    string

}

type Payload struct {
    OvType
    Files    []*File
    Type    VmDeviceType
    VolumeId    string

}

type Permission struct {
    Identified
    Cluster    Cluster
    Comment    string
    DataCenter    DataCenter
    Description    string
    Disk    Disk
    Group    Group
    Host    Host
    Id    string
    Name    string
    Role    Role
    StorageDomain    StorageDomain
    Template    Template
    User    User
    Vm    Vm
    VmPool    VmPool

}

type Permit struct {
    Identified
    Administrative    bool
    Comment    string
    Description    string
    Id    string
    Name    string
    Role    Role

}

type PmProxy struct {
    OvType
    Type    PmProxyType

}

type PortMirroring struct {
    OvType

}

type PowerManagement struct {
    OvType
    Address    string
    Agents    []*Agent
    AutomaticPmEnabled    bool
    Enabled    bool
    KdumpDetection    bool
    Options    []*Option
    Password    string
    PmProxies    []*PmProxy
    Status    PowerManagementStatus
    Type    string
    Username    string

}

type Product struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string

}

type ProductInfo struct {
    OvType
    Name    string
    Vendor    string
    Version    Version

}

type ProfileDetail struct {
    OvType
    BlockStatistics    []*BlockStatistic
    Duration    int64
    FopStatistics    []*FopStatistic
    ProfileType    string
    Statistics    []*Statistic

}

type Property struct {
    OvType
    Name    string
    Value    string

}

type ProxyTicket struct {
    OvType
    Value    string

}

type Qos struct {
    Identified
    Comment    string
    CpuLimit    int64
    DataCenter    DataCenter
    Description    string
    Id    string
    InboundAverage    int64
    InboundBurst    int64
    InboundPeak    int64
    MaxIops    int64
    MaxReadIops    int64
    MaxReadThroughput    int64
    MaxThroughput    int64
    MaxWriteIops    int64
    MaxWriteThroughput    int64
    Name    string
    OutboundAverage    int64
    OutboundAverageLinkshare    int64
    OutboundAverageRealtime    int64
    OutboundAverageUpperlimit    int64
    OutboundBurst    int64
    OutboundPeak    int64
    Type    QosType

}

type Quota struct {
    Identified
    ClusterHardLimitPct    int64
    ClusterSoftLimitPct    int64
    Comment    string
    DataCenter    DataCenter
    Description    string
    Disks    []*Disk
    Id    string
    Name    string
    Permissions    []*Permission
    QuotaClusterLimits    []*QuotaClusterLimit
    QuotaStorageLimits    []*QuotaStorageLimit
    StorageHardLimitPct    int64
    StorageSoftLimitPct    int64
    Users    []*User
    Vms    []*Vm

}

type QuotaClusterLimit struct {
    Identified
    Cluster    Cluster
    Comment    string
    Description    string
    Id    string
    MemoryLimit    float64
    MemoryUsage    float64
    Name    string
    Quota    Quota
    VcpuLimit    int64
    VcpuUsage    int64

}

type QuotaStorageLimit struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Limit    int64
    Name    string
    Quota    Quota
    StorageDomain    StorageDomain
    Usage    float64

}

type Range struct {
    OvType
    From    string
    To    string

}

type Rate struct {
    OvType
    Bytes    int64
    Period    int64

}

type ReportedConfiguration struct {
    OvType
    ActualValue    string
    ExpectedValue    string
    InSync    bool
    Name    string

}

type ReportedDevice struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Ips    []*Ip
    Mac    Mac
    Name    string
    Type    ReportedDeviceType
    Vm    Vm

}

type RngDevice struct {
    OvType
    Rate    Rate
    Source    RngSource

}

type Role struct {
    Identified
    Administrative    bool
    Comment    string
    Description    string
    Id    string
    Mutable    bool
    Name    string
    Permits    []*Permit
    User    User

}

type SchedulingPolicy struct {
    Identified
    Balances    []*Balance
    Comment    string
    DefaultPolicy    bool
    Description    string
    Filters    []*Filter
    Id    string
    Locked    bool
    Name    string
    Properties    []*Property
    Weight    []*Weight

}

type SchedulingPolicyUnit struct {
    Identified
    Comment    string
    Description    string
    Enabled    bool
    Id    string
    Internal    bool
    Name    string
    Properties    []*Property
    Type    PolicyUnitType

}

type SeLinux struct {
    OvType
    Mode    SeLinuxMode

}

type SerialNumber struct {
    OvType
    Policy    SerialNumberPolicy
    Value    string

}

type Session struct {
    Identified
    Comment    string
    ConsoleUser    bool
    Description    string
    Id    string
    Ip    Ip
    Name    string
    Protocol    string
    User    User
    Vm    Vm

}

type SkipIfConnectivityBroken struct {
    OvType
    Enabled    bool
    Threshold    int64

}

type SkipIfSdActive struct {
    OvType
    Enabled    bool

}

type SpecialObjects struct {
    OvType
    BlankTemplate    Template
    RootTag    Tag

}

type Spm struct {
    OvType
    Priority    int64
    Status    SpmStatus

}

type Ssh struct {
    Identified
    AuthenticationMethod    SshAuthenticationMethod
    Comment    string
    Description    string
    Fingerprint    string
    Id    string
    Name    string
    Port    int64
    User    User

}

type SshPublicKey struct {
    Identified
    Comment    string
    Content    string
    Description    string
    Id    string
    Name    string
    User    User

}

type Sso struct {
    OvType
    Methods    []*Method

}

type Statistic struct {
    Identified
    Brick    GlusterBrick
    Comment    string
    Description    string
    Disk    Disk
    GlusterVolume    GlusterVolume
    Host    Host
    HostNic    HostNic
    HostNumaNode    NumaNode
    Id    string
    Kind    StatisticKind
    Name    string
    Nic    Nic
    Step    Step
    Type    ValueType
    Unit    StatisticUnit
    Values    []*Value
    Vm    Vm

}

type Step struct {
    Identified
    Comment    string
    Description    string
    EndTime    time.Time
    ExecutionHost    Host
    External    bool
    ExternalType    ExternalSystemType
    Id    string
    Job    Job
    Name    string
    Number    int64
    ParentStep    Step
    Progress    int64
    StartTime    time.Time
    Statistics    []*Statistic
    Status    StepStatus
    Type    StepEnum

}

type StorageConnection struct {
    Identified
    Address    string
    Comment    string
    Description    string
    Host    Host
    Id    string
    MountOptions    string
    Name    string
    NfsRetrans    int64
    NfsTimeo    int64
    NfsVersion    NfsVersion
    Password    string
    Path    string
    Port    int64
    Portal    string
    Target    string
    Type    StorageType
    Username    string
    VfsType    string

}

type StorageConnectionExtension struct {
    Identified
    Comment    string
    Description    string
    Host    Host
    Id    string
    Name    string
    Password    string
    Target    string
    Username    string

}

type StorageDomain struct {
    Identified
    Available    int64
    Comment    string
    Committed    int64
    CriticalSpaceActionBlocker    int64
    DataCenter    DataCenter
    DataCenters    []*DataCenter
    Description    string
    DiscardAfterDelete    bool
    DiskProfiles    []*DiskProfile
    DiskSnapshots    []*DiskSnapshot
    Disks    []*Disk
    ExternalStatus    ExternalStatus
    Files    []*File
    Host    Host
    Id    string
    Images    []*Image
    Import    bool
    Master    bool
    Name    string
    Permissions    []*Permission
    Status    StorageDomainStatus
    Storage    HostStorage
    StorageConnections    []*StorageConnection
    StorageFormat    StorageFormat
    SupportsDiscard    bool
    SupportsDiscardZeroesData    bool
    Templates    []*Template
    Type    StorageDomainType
    Used    int64
    Vms    []*Vm
    WarningLowSpaceIndicator    int64
    WipeAfterDelete    bool

}

type StorageDomainLease struct {
    OvType
    StorageDomain    StorageDomain

}

type Tag struct {
    Identified
    Comment    string
    Description    string
    Group    Group
    Host    Host
    Id    string
    Name    string
    Parent    Tag
    Template    Template
    User    User
    Vm    Vm

}

type TemplateVersion struct {
    OvType
    BaseTemplate    Template
    VersionName    string
    VersionNumber    int64

}

type Ticket struct {
    OvType
    Expiry    int64
    Value    string

}

type TimeZone struct {
    OvType
    Name    string
    UtcOffset    string

}

type TransparentHugePages struct {
    OvType
    Enabled    bool

}

type UnmanagedNetwork struct {
    Identified
    Comment    string
    Description    string
    Host    Host
    HostNic    HostNic
    Id    string
    Name    string

}

type Usb struct {
    OvType
    Enabled    bool
    Type    UsbType

}

type User struct {
    Identified
    Comment    string
    Department    string
    Description    string
    Domain    Domain
    DomainEntryId    string
    Email    string
    Groups    []*Group
    Id    string
    LastName    string
    LoggedIn    bool
    Name    string
    Namespace    string
    Password    string
    Permissions    []*Permission
    Principal    string
    Roles    []*Role
    SshPublicKeys    []*SshPublicKey
    Tags    []*Tag
    UserName    string

}

type Value struct {
    OvType
    Datum    float64
    Detail    string

}

type VcpuPin struct {
    OvType
    CpuSet    string
    Vcpu    int64

}

type Vendor struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string

}

type Version struct {
    Identified
    Build    int64
    Comment    string
    Description    string
    FullVersion    string
    Id    string
    Major    int64
    Minor    int64
    Name    string
    Revision    int64

}

type VirtioScsi struct {
    OvType
    Enabled    bool

}

type VirtualNumaNode struct {
    NumaNode
    Comment    string
    Cpu    Cpu
    Description    string
    Host    Host
    Id    string
    Index    int64
    Memory    int64
    Name    string
    NodeDistance    string
    NumaNodePins    []*NumaNodePin
    Statistics    []*Statistic
    Vm    Vm

}

type Vlan struct {
    OvType
    Id    int64

}

type VmBase struct {
    Identified
    Bios    Bios
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int64
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []*CustomProperty
    DeleteProtected    bool
    Description    string
    Display    Display
    Domain    Domain
    HighAvailability    HighAvailability
    Id    string
    Initialization    Initialization
    Io    Io
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int64
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int64
    Name    string
    Origin    string
    Os    OperatingSystem
    Quota    Quota
    RngDevice    RngDevice
    SerialNumber    SerialNumber
    SmallIcon    Icon
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    Stateless    bool
    StorageDomain    StorageDomain
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    VirtioScsi    VirtioScsi

}

type VmPlacementPolicy struct {
    OvType
    Affinity    VmAffinity
    Hosts    []*Host

}

type VmPool struct {
    Identified
    AutoStorageSelect    bool
    Cluster    Cluster
    Comment    string
    Description    string
    Display    Display
    Id    string
    InstanceType    InstanceType
    MaxUserVms    int64
    Name    string
    Permissions    []*Permission
    PrestartedVms    int64
    RngDevice    RngDevice
    Size    int64
    SoundcardEnabled    bool
    Stateful    bool
    Template    Template
    Type    VmPoolType
    UseLatestTemplateVersion    bool
    Vm    Vm

}

type VmSummary struct {
    OvType
    Active    int64
    Migrating    int64
    Total    int64

}

type VnicPassThrough struct {
    OvType
    Mode    VnicPassThroughMode

}

type VnicProfile struct {
    Identified
    Comment    string
    CustomProperties    []*CustomProperty
    Description    string
    Id    string
    Migratable    bool
    Name    string
    Network    Network
    NetworkFilter    NetworkFilter
    PassThrough    VnicPassThrough
    Permissions    []*Permission
    PortMirroring    bool
    Qos    Qos

}

type VnicProfileMapping struct {
    OvType
    SourceNetworkName    string
    SourceNetworkProfileName    string
    TargetVnicProfile    VnicProfile

}

type VolumeGroup struct {
    OvType
    Id    string
    LogicalUnits    []*LogicalUnit
    Name    string

}

type Weight struct {
    Identified
    Comment    string
    Description    string
    Factor    int64
    Id    string
    Name    string
    SchedulingPolicy    SchedulingPolicy
    SchedulingPolicyUnit    SchedulingPolicyUnit

}

type Action struct {
    Identified
    AllowPartialImport    bool
    Async    bool
    Bricks    []*GlusterBrick
    Certificates    []*Certificate
    CheckConnectivity    bool
    Clone    bool
    Cluster    Cluster
    CollapseSnapshots    bool
    Comment    string
    ConnectivityTimeout    int64
    DataCenter    DataCenter
    DeployHostedEngine    bool
    Description    string
    Details    GlusterVolumeProfileDetails
    DiscardSnapshots    bool
    Disk    Disk
    Disks    []*Disk
    Exclusive    bool
    Fault    Fault
    FenceType    string
    Filter    bool
    FixLayout    bool
    Force    bool
    GracePeriod    GracePeriod
    Host    Host
    Id    string
    Image    string
    ImportAsTemplate    bool
    IsAttached    bool
    Iscsi    IscsiDetails
    IscsiTargets    []*String
    Job    Job
    LogicalUnits    []*LogicalUnit
    MaintenanceEnabled    bool
    ModifiedBonds    []*HostNic
    ModifiedLabels    []*NetworkLabel
    ModifiedNetworkAttachments    []*NetworkAttachment
    Name    string
    Option    Option
    Pause    bool
    PowerManagement    PowerManagement
    ProxyTicket    ProxyTicket
    Reason    string
    ReassignBadMacs    bool
    RemoteViewerConnectionFile    string
    RemovedBonds    []*HostNic
    RemovedLabels    []*NetworkLabel
    RemovedNetworkAttachments    []*NetworkAttachment
    ResolutionType    string
    RestoreMemory    bool
    RootPassword    string
    Snapshot    Snapshot
    Ssh    Ssh
    Status    string
    StopGlusterService    bool
    StorageDomain    StorageDomain
    StorageDomains    []*StorageDomain
    Succeeded    bool
    SynchronizedNetworkAttachments    []*NetworkAttachment
    Template    Template
    Ticket    Ticket
    UndeployHostedEngine    bool
    UseCloudInit    bool
    UseSysprep    bool
    VirtualFunctionsConfiguration    HostNicVirtualFunctionsConfiguration
    Vm    Vm
    VnicProfileMappings    []*VnicProfileMapping

}

type AffinityGroup struct {
    Identified
    Cluster    Cluster
    Comment    string
    Description    string
    Enforcing    bool
    Hosts    []*Host
    HostsRule    AffinityRule
    Id    string
    Name    string
    Positive    bool
    Vms    []*Vm
    VmsRule    AffinityRule

}

type AffinityLabel struct {
    Identified
    Comment    string
    Description    string
    Hosts    []*Host
    Id    string
    Name    string
    ReadOnly    bool
    Vms    []*Vm

}

type Agent struct {
    Identified
    Address    string
    Comment    string
    Concurrent    bool
    Description    string
    EncryptOptions    bool
    Host    Host
    Id    string
    Name    string
    Options    []*Option
    Order    int64
    Password    string
    Port    int64
    Type    string
    Username    string

}

type Application struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    Vm    Vm

}

type AuthorizedKey struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Key    string
    Name    string
    User    User

}

type Balance struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    SchedulingPolicy    SchedulingPolicy
    SchedulingPolicyUnit    SchedulingPolicyUnit

}

type Bookmark struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    Value    string

}

type BrickProfileDetail struct {
    EntityProfileDetail
    Brick    GlusterBrick
    ProfileDetails    []*ProfileDetail

}

type Certificate struct {
    Identified
    Comment    string
    Content    string
    Description    string
    Id    string
    Name    string
    Organization    string
    Subject    string

}

type Cluster struct {
    Identified
    AffinityGroups    []*AffinityGroup
    BallooningEnabled    bool
    Comment    string
    Cpu    Cpu
    CpuProfiles    []*CpuProfile
    CustomSchedulingPolicyProperties    []*Property
    DataCenter    DataCenter
    Description    string
    Display    Display
    ErrorHandling    ErrorHandling
    FencingPolicy    FencingPolicy
    GlusterHooks    []*GlusterHook
    GlusterService    bool
    GlusterTunedProfile    string
    GlusterVolumes    []*GlusterVolume
    HaReservation    bool
    Id    string
    Ksm    Ksm
    MacPool    MacPool
    MaintenanceReasonRequired    bool
    ManagementNetwork    Network
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    Name    string
    NetworkFilters    []*NetworkFilter
    Networks    []*Network
    OptionalReason    bool
    Permissions    []*Permission
    RequiredRngSources    []*RngSource
    SchedulingPolicy    SchedulingPolicy
    SerialNumber    SerialNumber
    SupportedVersions    []*Version
    SwitchType    SwitchType
    ThreadsAsCores    bool
    TrustedService    bool
    TunnelMigration    bool
    Version    Version
    VirtService    bool

}

type ClusterLevel struct {
    Identified
    Comment    string
    CpuTypes    []*CpuType
    Description    string
    Id    string
    Name    string
    Permits    []*Permit

}

type CpuProfile struct {
    Identified
    Cluster    Cluster
    Comment    string
    Description    string
    Id    string
    Name    string
    Permissions    []*Permission
    Qos    Qos

}

type DataCenter struct {
    Identified
    Clusters    []*Cluster
    Comment    string
    Description    string
    Id    string
    IscsiBonds    []*IscsiBond
    Local    bool
    MacPool    MacPool
    Name    string
    Networks    []*Network
    Permissions    []*Permission
    Qoss    []*Qos
    QuotaMode    QuotaModeType
    Quotas    []*Quota
    Status    DataCenterStatus
    StorageDomains    []*StorageDomain
    StorageFormat    StorageFormat
    SupportedVersions    []*Version
    Version    Version

}

type Device struct {
    Identified
    Comment    string
    Description    string
    Id    string
    InstanceType    InstanceType
    Name    string
    Template    Template
    Vm    Vm
    Vms    []*Vm

}

type Disk struct {
    Device
    Active    bool
    ActualSize    int64
    Alias    string
    Bootable    bool
    Comment    string
    Description    string
    DiskProfile    DiskProfile
    Format    DiskFormat
    Id    string
    ImageId    string
    InitialSize    int64
    InstanceType    InstanceType
    Interface    DiskInterface
    LogicalName    string
    LunStorage    HostStorage
    Name    string
    OpenstackVolumeType    OpenStackVolumeType
    Permissions    []*Permission
    PropagateErrors    bool
    ProvisionedSize    int64
    QcowVersion    QcowVersion
    Quota    Quota
    ReadOnly    bool
    Sgio    ScsiGenericIO
    Shareable    bool
    Snapshot    Snapshot
    Sparse    bool
    Statistics    []*Statistic
    Status    DiskStatus
    StorageDomain    StorageDomain
    StorageDomains    []*StorageDomain
    StorageType    DiskStorageType
    Template    Template
    UsesScsiReservation    bool
    Vm    Vm
    Vms    []*Vm
    WipeAfterDelete    bool

}

type DiskAttachment struct {
    Identified
    Active    bool
    Bootable    bool
    Comment    string
    Description    string
    Disk    Disk
    Id    string
    Interface    DiskInterface
    LogicalName    string
    Name    string
    PassDiscard    bool
    Template    Template
    UsesScsiReservation    bool
    Vm    Vm

}

type DiskProfile struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    Permissions    []*Permission
    Qos    Qos
    StorageDomain    StorageDomain

}

type DiskSnapshot struct {
    Disk
    Active    bool
    ActualSize    int64
    Alias    string
    Bootable    bool
    Comment    string
    Description    string
    Disk    Disk
    DiskProfile    DiskProfile
    Format    DiskFormat
    Id    string
    ImageId    string
    InitialSize    int64
    InstanceType    InstanceType
    Interface    DiskInterface
    LogicalName    string
    LunStorage    HostStorage
    Name    string
    OpenstackVolumeType    OpenStackVolumeType
    Permissions    []*Permission
    PropagateErrors    bool
    ProvisionedSize    int64
    QcowVersion    QcowVersion
    Quota    Quota
    ReadOnly    bool
    Sgio    ScsiGenericIO
    Shareable    bool
    Snapshot    Snapshot
    Sparse    bool
    Statistics    []*Statistic
    Status    DiskStatus
    StorageDomain    StorageDomain
    StorageDomains    []*StorageDomain
    StorageType    DiskStorageType
    Template    Template
    UsesScsiReservation    bool
    Vm    Vm
    Vms    []*Vm
    WipeAfterDelete    bool

}

type Domain struct {
    Identified
    Comment    string
    Description    string
    Groups    []*Group
    Id    string
    Name    string
    User    User
    Users    []*User

}

type Event struct {
    Identified
    Cluster    Cluster
    Code    int64
    Comment    string
    CorrelationId    string
    CustomData    string
    CustomId    int64
    DataCenter    DataCenter
    Description    string
    FloodRate    int64
    Host    Host
    Id    string
    Name    string
    Origin    string
    Severity    LogSeverity
    StorageDomain    StorageDomain
    Template    Template
    Time    time.Time
    User    User
    Vm    Vm

}

type ExternalComputeResource struct {
    Identified
    Comment    string
    Description    string
    ExternalHostProvider    ExternalHostProvider
    Id    string
    Name    string
    Provider    string
    Url    string
    User    string

}

type ExternalDiscoveredHost struct {
    Identified
    Comment    string
    Description    string
    ExternalHostProvider    ExternalHostProvider
    Id    string
    Ip    string
    LastReport    string
    Mac    string
    Name    string
    SubnetName    string

}

type ExternalHost struct {
    Identified
    Address    string
    Comment    string
    Description    string
    ExternalHostProvider    ExternalHostProvider
    Id    string
    Name    string

}

type ExternalHostGroup struct {
    Identified
    ArchitectureName    string
    Comment    string
    Description    string
    DomainName    string
    ExternalHostProvider    ExternalHostProvider
    Id    string
    Name    string
    OperatingSystemName    string
    SubnetName    string

}

type ExternalProvider struct {
    Identified
    AuthenticationUrl    string
    Comment    string
    Description    string
    Id    string
    Name    string
    Password    string
    Properties    []*Property
    RequiresAuthentication    bool
    Url    string
    Username    string

}

type File struct {
    Identified
    Comment    string
    Content    string
    Description    string
    Id    string
    Name    string
    StorageDomain    StorageDomain
    Type    string

}

type Filter struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Name    string
    Position    int64
    SchedulingPolicyUnit    SchedulingPolicyUnit

}

type Floppy struct {
    Device
    Comment    string
    Description    string
    File    File
    Id    string
    InstanceType    InstanceType
    Name    string
    Template    Template
    Vm    Vm
    Vms    []*Vm

}

type GlusterBrickAdvancedDetails struct {
    Device
    Comment    string
    Description    string
    Device    string
    FsName    string
    GlusterClients    []*GlusterClient
    Id    string
    InstanceType    InstanceType
    MemoryPools    []*GlusterMemoryPool
    MntOptions    string
    Name    string
    Pid    int64
    Port    int64
    Template    Template
    Vm    Vm
    Vms    []*Vm

}

type GlusterHook struct {
    Identified
    Checksum    string
    Cluster    Cluster
    Comment    string
    ConflictStatus    int64
    Conflicts    string
    Content    string
    ContentType    HookContentType
    Description    string
    GlusterCommand    string
    Id    string
    Name    string
    ServerHooks    []*GlusterServerHook
    Stage    HookStage
    Status    GlusterHookStatus

}

type GlusterMemoryPool struct {
    Identified
    AllocCount    int64
    ColdCount    int64
    Comment    string
    Description    string
    HotCount    int64
    Id    string
    MaxAlloc    int64
    MaxStdalloc    int64
    Name    string
    PaddedSize    int64
    PoolMisses    int64
    Type    string

}

type GlusterServerHook struct {
    Identified
    Checksum    string
    Comment    string
    ContentType    HookContentType
    Description    string
    Host    Host
    Id    string
    Name    string
    Status    GlusterHookStatus

}

type GlusterVolume struct {
    Identified
    Bricks    []*GlusterBrick
    Cluster    Cluster
    Comment    string
    Description    string
    DisperseCount    int64
    Id    string
    Name    string
    Options    []*Option
    RedundancyCount    int64
    ReplicaCount    int64
    Statistics    []*Statistic
    Status    GlusterVolumeStatus
    StripeCount    int64
    TransportTypes    []*TransportType
    VolumeType    GlusterVolumeType

}

type GlusterVolumeProfileDetails struct {
    Identified
    BrickProfileDetails    []*BrickProfileDetail
    Comment    string
    Description    string
    Id    string
    Name    string
    NfsProfileDetails    []*NfsProfileDetail

}

type GraphicsConsole struct {
    Identified
    Address    string
    Comment    string
    Description    string
    Id    string
    InstanceType    InstanceType
    Name    string
    Port    int64
    Protocol    GraphicsType
    Template    Template
    TlsPort    int64
    Vm    Vm

}

type Group struct {
    Identified
    Comment    string
    Description    string
    Domain    Domain
    DomainEntryId    string
    Id    string
    Name    string
    Namespace    string
    Permissions    []*Permission
    Roles    []*Role
    Tags    []*Tag

}

type Hook struct {
    Identified
    Comment    string
    Description    string
    EventName    string
    Host    Host
    Id    string
    Md5    string
    Name    string

}

type Host struct {
    Identified
    Address    string
    AffinityLabels    []*AffinityLabel
    Agents    []*Agent
    AutoNumaStatus    AutoNumaStatus
    Certificate    Certificate
    Cluster    Cluster
    Comment    string
    Cpu    Cpu
    Description    string
    DevicePassthrough    HostDevicePassthrough
    Devices    []*Device
    Display    Display
    ExternalHostProvider    ExternalHostProvider
    ExternalStatus    ExternalStatus
    HardwareInformation    HardwareInformation
    Hooks    []*Hook
    HostedEngine    HostedEngine
    Id    string
    Iscsi    IscsiDetails
    KatelloErrata    []*KatelloErratum
    KdumpStatus    KdumpStatus
    Ksm    Ksm
    LibvirtVersion    Version
    MaxSchedulingMemory    int64
    Memory    int64
    Name    string
    NetworkAttachments    []*NetworkAttachment
    Nics    []*Nic
    NumaNodes    []*NumaNode
    NumaSupported    bool
    Os    OperatingSystem
    OverrideIptables    bool
    Permissions    []*Permission
    Port    int64
    PowerManagement    PowerManagement
    Protocol    HostProtocol
    RootPassword    string
    SeLinux    SeLinux
    Spm    Spm
    Ssh    Ssh
    Statistics    []*Statistic
    Status    HostStatus
    StatusDetail    string
    StorageConnectionExtensions    []*StorageConnectionExtension
    Storages    []*HostStorage
    Summary    VmSummary
    Tags    []*Tag
    TransparentHugePages    TransparentHugePages
    Type    HostType
    UnmanagedNetworks    []*UnmanagedNetwork
    UpdateAvailable    bool
    Version    Version

}

type HostDevice struct {
    Identified
    Capability    string
    Comment    string
    Description    string
    Host    Host
    Id    string
    IommuGroup    int64
    Name    string
    ParentDevice    HostDevice
    PhysicalFunction    HostDevice
    Placeholder    bool
    Product    Product
    Vendor    Vendor
    VirtualFunctions    int64
    Vm    Vm

}

type HostNic struct {
    Identified
    AdAggregatorId    int64
    BaseInterface    string
    Bonding    Bonding
    BootProtocol    BootProtocol
    Bridged    bool
    CheckConnectivity    bool
    Comment    string
    CustomConfiguration    bool
    Description    string
    Host    Host
    Id    string
    Ip    Ip
    Ipv6    Ip
    Ipv6BootProtocol    BootProtocol
    Mac    Mac
    Mtu    int64
    Name    string
    Network    Network
    NetworkLabels    []*NetworkLabel
    OverrideConfiguration    bool
    PhysicalFunction    HostNic
    Properties    []*Property
    Qos    Qos
    Speed    int64
    Statistics    []*Statistic
    Status    NicStatus
    VirtualFunctionsConfiguration    HostNicVirtualFunctionsConfiguration
    Vlan    Vlan

}

type HostStorage struct {
    Identified
    Address    string
    Comment    string
    Description    string
    Host    Host
    Id    string
    LogicalUnits    []*LogicalUnit
    MountOptions    string
    Name    string
    NfsRetrans    int64
    NfsTimeo    int64
    NfsVersion    NfsVersion
    OverrideLuns    bool
    Password    string
    Path    string
    Port    int64
    Portal    string
    Target    string
    Type    StorageType
    Username    string
    VfsType    string
    VolumeGroup    VolumeGroup

}

type Icon struct {
    Identified
    Comment    string
    Data    string
    Description    string
    Id    string
    MediaType    string
    Name    string

}

type Nic struct {
    Device
    BootProtocol    BootProtocol
    Comment    string
    Description    string
    Id    string
    InstanceType    InstanceType
    Interface    NicInterface
    Linked    bool
    Mac    Mac
    Name    string
    Network    Network
    NetworkAttachments    []*NetworkAttachment
    NetworkFilterParameters    []*NetworkFilterParameter
    NetworkLabels    []*NetworkLabel
    OnBoot    bool
    Plugged    bool
    ReportedDevices    []*ReportedDevice
    Statistics    []*Statistic
    Template    Template
    VirtualFunctionAllowedLabels    []*NetworkLabel
    VirtualFunctionAllowedNetworks    []*Network
    Vm    Vm
    Vms    []*Vm
    VnicProfile    VnicProfile

}

type OpenStackProvider struct {
    ExternalProvider
    AuthenticationUrl    string
    Comment    string
    Description    string
    Id    string
    Name    string
    Password    string
    Properties    []*Property
    RequiresAuthentication    bool
    TenantName    string
    Url    string
    Username    string

}

type OpenStackVolumeProvider struct {
    OpenStackProvider
    AuthenticationKeys    []*OpenstackVolumeAuthenticationKey
    AuthenticationUrl    string
    Certificates    []*Certificate
    Comment    string
    DataCenter    DataCenter
    Description    string
    Id    string
    Name    string
    Password    string
    Properties    []*Property
    RequiresAuthentication    bool
    TenantName    string
    Url    string
    Username    string
    VolumeTypes    []*OpenStackVolumeType

}

type Template struct {
    VmBase
    Bios    Bios
    Cdroms    []*Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int64
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []*CustomProperty
    DeleteProtected    bool
    Description    string
    DiskAttachments    []*DiskAttachment
    Display    Display
    Domain    Domain
    GraphicsConsoles    []*GraphicsConsole
    HighAvailability    HighAvailability
    Id    string
    Initialization    Initialization
    Io    Io
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int64
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int64
    Name    string
    Nics    []*Nic
    Origin    string
    Os    OperatingSystem
    Permissions    []*Permission
    Quota    Quota
    RngDevice    RngDevice
    SerialNumber    SerialNumber
    SmallIcon    Icon
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    Stateless    bool
    Status    TemplateStatus
    StorageDomain    StorageDomain
    Tags    []*Tag
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    Version    TemplateVersion
    VirtioScsi    VirtioScsi
    Vm    Vm
    Watchdogs    []*Watchdog

}

type Vm struct {
    VmBase
    AffinityLabels    []*AffinityLabel
    Applications    []*Application
    Bios    Bios
    Cdroms    []*Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int64
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []*CustomProperty
    DeleteProtected    bool
    Description    string
    DiskAttachments    []*DiskAttachment
    Display    Display
    Domain    Domain
    ExternalHostProvider    ExternalHostProvider
    Floppies    []*Floppy
    Fqdn    string
    GraphicsConsoles    []*GraphicsConsole
    GuestOperatingSystem    GuestOperatingSystem
    GuestTimeZone    TimeZone
    HighAvailability    HighAvailability
    Host    Host
    HostDevices    []*HostDevice
    Id    string
    Initialization    Initialization
    InstanceType    InstanceType
    Io    Io
    KatelloErrata    []*KatelloErratum
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int64
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int64
    Name    string
    NextRunConfigurationExists    bool
    Nics    []*Nic
    NumaNodes    []*NumaNode
    NumaTuneMode    NumaTuneMode
    Origin    string
    OriginalTemplate    Template
    Os    OperatingSystem
    Payloads    []*Payload
    Permissions    []*Permission
    PlacementPolicy    VmPlacementPolicy
    Quota    Quota
    ReportedDevices    []*ReportedDevice
    RngDevice    RngDevice
    RunOnce    bool
    SerialNumber    SerialNumber
    Sessions    []*Session
    SmallIcon    Icon
    Snapshots    []*Snapshot
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    StartTime    time.Time
    Stateless    bool
    Statistics    []*Statistic
    Status    VmStatus
    StatusDetail    string
    StopReason    string
    StopTime    time.Time
    StorageDomain    StorageDomain
    Tags    []*Tag
    Template    Template
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    UseLatestTemplateVersion    bool
    VirtioScsi    VirtioScsi
    VmPool    VmPool
    Watchdogs    []*Watchdog

}

type Watchdog struct {
    Device
    Action    WatchdogAction
    Comment    string
    Description    string
    Id    string
    InstanceType    InstanceType
    Model    WatchdogModel
    Name    string
    Template    Template
    Vm    Vm
    Vms    []*Vm

}

type Cdrom struct {
    Device
    Comment    string
    Description    string
    File    File
    Id    string
    InstanceType    InstanceType
    Name    string
    Template    Template
    Vm    Vm
    Vms    []*Vm

}

type ExternalHostProvider struct {
    ExternalProvider
    AuthenticationUrl    string
    Certificates    []*Certificate
    Comment    string
    ComputeResources    []*ExternalComputeResource
    Description    string
    DiscoveredHosts    []*ExternalDiscoveredHost
    HostGroups    []*ExternalHostGroup
    Hosts    []*Host
    Id    string
    Name    string
    Password    string
    Properties    []*Property
    RequiresAuthentication    bool
    Url    string
    Username    string

}

type GlusterBrick struct {
    GlusterBrickAdvancedDetails
    BrickDir    string
    Comment    string
    Description    string
    Device    string
    FsName    string
    GlusterClients    []*GlusterClient
    GlusterVolume    GlusterVolume
    Id    string
    InstanceType    InstanceType
    MemoryPools    []*GlusterMemoryPool
    MntOptions    string
    Name    string
    Pid    int64
    Port    int64
    ServerId    string
    Statistics    []*Statistic
    Status    GlusterBrickStatus
    Template    Template
    Vm    Vm
    Vms    []*Vm

}

type InstanceType struct {
    Template
    Bios    Bios
    Cdroms    []*Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int64
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []*CustomProperty
    DeleteProtected    bool
    Description    string
    DiskAttachments    []*DiskAttachment
    Display    Display
    Domain    Domain
    GraphicsConsoles    []*GraphicsConsole
    HighAvailability    HighAvailability
    Id    string
    Initialization    Initialization
    Io    Io
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int64
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int64
    Name    string
    Nics    []*Nic
    Origin    string
    Os    OperatingSystem
    Permissions    []*Permission
    Quota    Quota
    RngDevice    RngDevice
    SerialNumber    SerialNumber
    SmallIcon    Icon
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    Stateless    bool
    Status    TemplateStatus
    StorageDomain    StorageDomain
    Tags    []*Tag
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    Version    TemplateVersion
    VirtioScsi    VirtioScsi
    Vm    Vm
    Watchdogs    []*Watchdog

}

type OpenStackImageProvider struct {
    OpenStackProvider
    AuthenticationUrl    string
    Certificates    []*Certificate
    Comment    string
    Description    string
    Id    string
    Images    []*OpenStackImage
    Name    string
    Password    string
    Properties    []*Property
    RequiresAuthentication    bool
    TenantName    string
    Url    string
    Username    string

}

type OpenStackNetworkProvider struct {
    OpenStackProvider
    AgentConfiguration    AgentConfiguration
    AuthenticationUrl    string
    Certificates    []*Certificate
    Comment    string
    Description    string
    Id    string
    Name    string
    Networks    []*OpenStackNetwork
    Password    string
    PluginType    NetworkPluginType
    Properties    []*Property
    ReadOnly    bool
    RequiresAuthentication    bool
    Subnets    []*OpenStackSubnet
    TenantName    string
    Type    OpenStackNetworkProviderType
    Url    string
    Username    string

}

type Snapshot struct {
    Vm
    AffinityLabels    []*AffinityLabel
    Applications    []*Application
    Bios    Bios
    Cdroms    []*Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int64
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []*CustomProperty
    Date    time.Time
    DeleteProtected    bool
    Description    string
    DiskAttachments    []*DiskAttachment
    Display    Display
    Domain    Domain
    ExternalHostProvider    ExternalHostProvider
    Floppies    []*Floppy
    Fqdn    string
    GraphicsConsoles    []*GraphicsConsole
    GuestOperatingSystem    GuestOperatingSystem
    GuestTimeZone    TimeZone
    HighAvailability    HighAvailability
    Host    Host
    HostDevices    []*HostDevice
    Id    string
    Initialization    Initialization
    InstanceType    InstanceType
    Io    Io
    KatelloErrata    []*KatelloErratum
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int64
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int64
    Name    string
    NextRunConfigurationExists    bool
    Nics    []*Nic
    NumaNodes    []*NumaNode
    NumaTuneMode    NumaTuneMode
    Origin    string
    OriginalTemplate    Template
    Os    OperatingSystem
    Payloads    []*Payload
    Permissions    []*Permission
    PersistMemorystate    bool
    PlacementPolicy    VmPlacementPolicy
    Quota    Quota
    ReportedDevices    []*ReportedDevice
    RngDevice    RngDevice
    RunOnce    bool
    SerialNumber    SerialNumber
    Sessions    []*Session
    SmallIcon    Icon
    SnapshotStatus    SnapshotStatus
    SnapshotType    SnapshotType
    Snapshots    []*Snapshot
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    StartTime    time.Time
    Stateless    bool
    Statistics    []*Statistic
    Status    VmStatus
    StatusDetail    string
    StopReason    string
    StopTime    time.Time
    StorageDomain    StorageDomain
    Tags    []*Tag
    Template    Template
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    UseLatestTemplateVersion    bool
    VirtioScsi    VirtioScsi
    Vm    Vm
    VmPool    VmPool
    Watchdogs    []*Watchdog

}

type AccessProtocol string
const (
    ACCESSPROTOCOL_CIFS AccessProtocol = "cifs"
    ACCESSPROTOCOL_GLUSTER AccessProtocol = "gluster"
    ACCESSPROTOCOL_NFS AccessProtocol = "nfs"
)

type Architecture string
const (
    ARCHITECTURE_PPC64 Architecture = "ppc64"
    ARCHITECTURE_UNDEFINED Architecture = "undefined"
    ARCHITECTURE_X86_64 Architecture = "x86_64"
)

type AutoNumaStatus string
const (
    AUTONUMASTATUS_DISABLE AutoNumaStatus = "disable"
    AUTONUMASTATUS_ENABLE AutoNumaStatus = "enable"
    AUTONUMASTATUS_UNKNOWN AutoNumaStatus = "unknown"
)

type BootDevice string
const (
    BOOTDEVICE_CDROM BootDevice = "cdrom"
    BOOTDEVICE_HD BootDevice = "hd"
    BOOTDEVICE_NETWORK BootDevice = "network"
)

type BootProtocol string
const (
    BOOTPROTOCOL_AUTOCONF BootProtocol = "autoconf"
    BOOTPROTOCOL_DHCP BootProtocol = "dhcp"
    BOOTPROTOCOL_NONE BootProtocol = "none"
    BOOTPROTOCOL_STATIC BootProtocol = "static"
)

type ConfigurationType string
const (
    CONFIGURATIONTYPE_OVF ConfigurationType = "ovf"
)

type CpuMode string
const (
    CPUMODE_CUSTOM CpuMode = "custom"
    CPUMODE_HOST_MODEL CpuMode = "host_model"
    CPUMODE_HOST_PASSTHROUGH CpuMode = "host_passthrough"
)

type CreationStatus string
const (
    CREATIONSTATUS_COMPLETE CreationStatus = "complete"
    CREATIONSTATUS_FAILED CreationStatus = "failed"
    CREATIONSTATUS_IN_PROGRESS CreationStatus = "in_progress"
    CREATIONSTATUS_PENDING CreationStatus = "pending"
)

type DataCenterStatus string
const (
    DATACENTERSTATUS_CONTEND DataCenterStatus = "contend"
    DATACENTERSTATUS_MAINTENANCE DataCenterStatus = "maintenance"
    DATACENTERSTATUS_NOT_OPERATIONAL DataCenterStatus = "not_operational"
    DATACENTERSTATUS_PROBLEMATIC DataCenterStatus = "problematic"
    DATACENTERSTATUS_UNINITIALIZED DataCenterStatus = "uninitialized"
    DATACENTERSTATUS_UP DataCenterStatus = "up"
)

type DiskFormat string
const (
    DISKFORMAT_COW DiskFormat = "cow"
    DISKFORMAT_RAW DiskFormat = "raw"
)

type DiskInterface string
const (
    DISKINTERFACE_IDE DiskInterface = "ide"
    DISKINTERFACE_SPAPR_VSCSI DiskInterface = "spapr_vscsi"
    DISKINTERFACE_VIRTIO DiskInterface = "virtio"
    DISKINTERFACE_VIRTIO_SCSI DiskInterface = "virtio_scsi"
)

type DiskStatus string
const (
    DISKSTATUS_ILLEGAL DiskStatus = "illegal"
    DISKSTATUS_LOCKED DiskStatus = "locked"
    DISKSTATUS_OK DiskStatus = "ok"
)

type DiskStorageType string
const (
    DISKSTORAGETYPE_CINDER DiskStorageType = "cinder"
    DISKSTORAGETYPE_IMAGE DiskStorageType = "image"
    DISKSTORAGETYPE_LUN DiskStorageType = "lun"
)

type DiskType string
const (
    DISKTYPE_DATA DiskType = "data"
    DISKTYPE_SYSTEM DiskType = "system"
)

type DisplayType string
const (
    DISPLAYTYPE_SPICE DisplayType = "spice"
    DISPLAYTYPE_VNC DisplayType = "vnc"
)

type EntityExternalStatus string
const (
    ENTITYEXTERNALSTATUS_ERROR EntityExternalStatus = "error"
    ENTITYEXTERNALSTATUS_FAILURE EntityExternalStatus = "failure"
    ENTITYEXTERNALSTATUS_INFO EntityExternalStatus = "info"
    ENTITYEXTERNALSTATUS_OK EntityExternalStatus = "ok"
    ENTITYEXTERNALSTATUS_WARNING EntityExternalStatus = "warning"
)

type ExternalStatus string
const (
    EXTERNALSTATUS_ERROR ExternalStatus = "error"
    EXTERNALSTATUS_FAILURE ExternalStatus = "failure"
    EXTERNALSTATUS_INFO ExternalStatus = "info"
    EXTERNALSTATUS_OK ExternalStatus = "ok"
    EXTERNALSTATUS_WARNING ExternalStatus = "warning"
)

type ExternalSystemType string
const (
    EXTERNALSYSTEMTYPE_GLUSTER ExternalSystemType = "gluster"
    EXTERNALSYSTEMTYPE_VDSM ExternalSystemType = "vdsm"
)

type ExternalVmProviderType string
const (
    EXTERNALVMPROVIDERTYPE_KVM ExternalVmProviderType = "kvm"
    EXTERNALVMPROVIDERTYPE_VMWARE ExternalVmProviderType = "vmware"
    EXTERNALVMPROVIDERTYPE_XEN ExternalVmProviderType = "xen"
)

type FenceType string
const (
    FENCETYPE_MANUAL FenceType = "manual"
    FENCETYPE_RESTART FenceType = "restart"
    FENCETYPE_START FenceType = "start"
    FENCETYPE_STATUS FenceType = "status"
    FENCETYPE_STOP FenceType = "stop"
)

type GlusterBrickStatus string
const (
    GLUSTERBRICKSTATUS_DOWN GlusterBrickStatus = "down"
    GLUSTERBRICKSTATUS_UNKNOWN GlusterBrickStatus = "unknown"
    GLUSTERBRICKSTATUS_UP GlusterBrickStatus = "up"
)

type GlusterHookStatus string
const (
    GLUSTERHOOKSTATUS_DISABLED GlusterHookStatus = "disabled"
    GLUSTERHOOKSTATUS_ENABLED GlusterHookStatus = "enabled"
    GLUSTERHOOKSTATUS_MISSING GlusterHookStatus = "missing"
)

type GlusterState string
const (
    GLUSTERSTATE_DOWN GlusterState = "down"
    GLUSTERSTATE_UNKNOWN GlusterState = "unknown"
    GLUSTERSTATE_UP GlusterState = "up"
)

type GlusterVolumeStatus string
const (
    GLUSTERVOLUMESTATUS_DOWN GlusterVolumeStatus = "down"
    GLUSTERVOLUMESTATUS_UNKNOWN GlusterVolumeStatus = "unknown"
    GLUSTERVOLUMESTATUS_UP GlusterVolumeStatus = "up"
)

type GlusterVolumeType string
const (
    GLUSTERVOLUMETYPE_DISPERSE GlusterVolumeType = "disperse"
    GLUSTERVOLUMETYPE_DISTRIBUTE GlusterVolumeType = "distribute"
    GLUSTERVOLUMETYPE_DISTRIBUTED_DISPERSE GlusterVolumeType = "distributed_disperse"
    GLUSTERVOLUMETYPE_DISTRIBUTED_REPLICATE GlusterVolumeType = "distributed_replicate"
    GLUSTERVOLUMETYPE_DISTRIBUTED_STRIPE GlusterVolumeType = "distributed_stripe"
    GLUSTERVOLUMETYPE_DISTRIBUTED_STRIPED_REPLICATE GlusterVolumeType = "distributed_striped_replicate"
    GLUSTERVOLUMETYPE_REPLICATE GlusterVolumeType = "replicate"
    GLUSTERVOLUMETYPE_STRIPE GlusterVolumeType = "stripe"
    GLUSTERVOLUMETYPE_STRIPED_REPLICATE GlusterVolumeType = "striped_replicate"
)

type GraphicsType string
const (
    GRAPHICSTYPE_SPICE GraphicsType = "spice"
    GRAPHICSTYPE_VNC GraphicsType = "vnc"
)

type HookContentType string
const (
    HOOKCONTENTTYPE_BINARY HookContentType = "binary"
    HOOKCONTENTTYPE_TEXT HookContentType = "text"
)

type HookStage string
const (
    HOOKSTAGE_POST HookStage = "post"
    HOOKSTAGE_PRE HookStage = "pre"
)

type HookStatus string
const (
    HOOKSTATUS_DISABLED HookStatus = "disabled"
    HOOKSTATUS_ENABLED HookStatus = "enabled"
    HOOKSTATUS_MISSING HookStatus = "missing"
)

type HostProtocol string
const (
    HOSTPROTOCOL_STOMP HostProtocol = "stomp"
    HOSTPROTOCOL_XML HostProtocol = "xml"
)

type HostStatus string
const (
    HOSTSTATUS_CONNECTING HostStatus = "connecting"
    HOSTSTATUS_DOWN HostStatus = "down"
    HOSTSTATUS_ERROR HostStatus = "error"
    HOSTSTATUS_INITIALIZING HostStatus = "initializing"
    HOSTSTATUS_INSTALL_FAILED HostStatus = "install_failed"
    HOSTSTATUS_INSTALLING HostStatus = "installing"
    HOSTSTATUS_INSTALLING_OS HostStatus = "installing_os"
    HOSTSTATUS_KDUMPING HostStatus = "kdumping"
    HOSTSTATUS_MAINTENANCE HostStatus = "maintenance"
    HOSTSTATUS_NON_OPERATIONAL HostStatus = "non_operational"
    HOSTSTATUS_NON_RESPONSIVE HostStatus = "non_responsive"
    HOSTSTATUS_PENDING_APPROVAL HostStatus = "pending_approval"
    HOSTSTATUS_PREPARING_FOR_MAINTENANCE HostStatus = "preparing_for_maintenance"
    HOSTSTATUS_REBOOT HostStatus = "reboot"
    HOSTSTATUS_UNASSIGNED HostStatus = "unassigned"
    HOSTSTATUS_UP HostStatus = "up"
)

type HostType string
const (
    HOSTTYPE_OVIRT_NODE HostType = "ovirt_node"
    HOSTTYPE_RHEL HostType = "rhel"
    HOSTTYPE_RHEV_H HostType = "rhev_h"
)

type ImageTransferDirection string
const (
    IMAGETRANSFERDIRECTION_DOWNLOAD ImageTransferDirection = "download"
    IMAGETRANSFERDIRECTION_UPLOAD ImageTransferDirection = "upload"
)

type ImageTransferPhase string
const (
    IMAGETRANSFERPHASE_CANCELLED ImageTransferPhase = "cancelled"
    IMAGETRANSFERPHASE_FINALIZING_FAILURE ImageTransferPhase = "finalizing_failure"
    IMAGETRANSFERPHASE_FINALIZING_SUCCESS ImageTransferPhase = "finalizing_success"
    IMAGETRANSFERPHASE_FINISHED_FAILURE ImageTransferPhase = "finished_failure"
    IMAGETRANSFERPHASE_FINISHED_SUCCESS ImageTransferPhase = "finished_success"
    IMAGETRANSFERPHASE_INITIALIZING ImageTransferPhase = "initializing"
    IMAGETRANSFERPHASE_PAUSED_SYSTEM ImageTransferPhase = "paused_system"
    IMAGETRANSFERPHASE_PAUSED_USER ImageTransferPhase = "paused_user"
    IMAGETRANSFERPHASE_RESUMING ImageTransferPhase = "resuming"
    IMAGETRANSFERPHASE_TRANSFERRING ImageTransferPhase = "transferring"
    IMAGETRANSFERPHASE_UNKNOWN ImageTransferPhase = "unknown"
)

type InheritableBoolean string
const (
    INHERITABLEBOOLEAN_FALSE InheritableBoolean = "false"
    INHERITABLEBOOLEAN_INHERIT InheritableBoolean = "inherit"
    INHERITABLEBOOLEAN_TRUE InheritableBoolean = "true"
)

type IpVersion string
const (
    IPVERSION_V4 IpVersion = "v4"
    IPVERSION_V6 IpVersion = "v6"
)

type JobStatus string
const (
    JOBSTATUS_ABORTED JobStatus = "aborted"
    JOBSTATUS_FAILED JobStatus = "failed"
    JOBSTATUS_FINISHED JobStatus = "finished"
    JOBSTATUS_STARTED JobStatus = "started"
    JOBSTATUS_UNKNOWN JobStatus = "unknown"
)

type KdumpStatus string
const (
    KDUMPSTATUS_DISABLED KdumpStatus = "disabled"
    KDUMPSTATUS_ENABLED KdumpStatus = "enabled"
    KDUMPSTATUS_UNKNOWN KdumpStatus = "unknown"
)

type LogSeverity string
const (
    LOGSEVERITY_ALERT LogSeverity = "alert"
    LOGSEVERITY_ERROR LogSeverity = "error"
    LOGSEVERITY_NORMAL LogSeverity = "normal"
    LOGSEVERITY_WARNING LogSeverity = "warning"
)

type LunStatus string
const (
    LUNSTATUS_FREE LunStatus = "free"
    LUNSTATUS_UNUSABLE LunStatus = "unusable"
    LUNSTATUS_USED LunStatus = "used"
)

type MessageBrokerType string
const (
    MESSAGEBROKERTYPE_QPID MessageBrokerType = "qpid"
    MESSAGEBROKERTYPE_RABBIT_MQ MessageBrokerType = "rabbit_mq"
)

type MigrateOnError string
const (
    MIGRATEONERROR_DO_NOT_MIGRATE MigrateOnError = "do_not_migrate"
    MIGRATEONERROR_MIGRATE MigrateOnError = "migrate"
    MIGRATEONERROR_MIGRATE_HIGHLY_AVAILABLE MigrateOnError = "migrate_highly_available"
)

type MigrationBandwidthAssignmentMethod string
const (
    MIGRATIONBANDWIDTHASSIGNMENTMETHOD_AUTO MigrationBandwidthAssignmentMethod = "auto"
    MIGRATIONBANDWIDTHASSIGNMENTMETHOD_CUSTOM MigrationBandwidthAssignmentMethod = "custom"
    MIGRATIONBANDWIDTHASSIGNMENTMETHOD_HYPERVISOR_DEFAULT MigrationBandwidthAssignmentMethod = "hypervisor_default"
)

type NetworkPluginType string
const (
    NETWORKPLUGINTYPE_OPEN_VSWITCH NetworkPluginType = "open_vswitch"
)

type NetworkStatus string
const (
    NETWORKSTATUS_NON_OPERATIONAL NetworkStatus = "non_operational"
    NETWORKSTATUS_OPERATIONAL NetworkStatus = "operational"
)

type NetworkUsage string
const (
    NETWORKUSAGE_DEFAULT_ROUTE NetworkUsage = "default_route"
    NETWORKUSAGE_DISPLAY NetworkUsage = "display"
    NETWORKUSAGE_GLUSTER NetworkUsage = "gluster"
    NETWORKUSAGE_MANAGEMENT NetworkUsage = "management"
    NETWORKUSAGE_MIGRATION NetworkUsage = "migration"
    NETWORKUSAGE_VM NetworkUsage = "vm"
)

type NfsVersion string
const (
    NFSVERSION_AUTO NfsVersion = "auto"
    NFSVERSION_V3 NfsVersion = "v3"
    NFSVERSION_V4 NfsVersion = "v4"
    NFSVERSION_V4_1 NfsVersion = "v4_1"
    NFSVERSION_V4_2 NfsVersion = "v4_2"
)

type NicInterface string
const (
    NICINTERFACE_E1000 NicInterface = "e1000"
    NICINTERFACE_PCI_PASSTHROUGH NicInterface = "pci_passthrough"
    NICINTERFACE_RTL8139 NicInterface = "rtl8139"
    NICINTERFACE_RTL8139_VIRTIO NicInterface = "rtl8139_virtio"
    NICINTERFACE_SPAPR_VLAN NicInterface = "spapr_vlan"
    NICINTERFACE_VIRTIO NicInterface = "virtio"
)

type NicStatus string
const (
    NICSTATUS_DOWN NicStatus = "down"
    NICSTATUS_UP NicStatus = "up"
)

type NumaTuneMode string
const (
    NUMATUNEMODE_INTERLEAVE NumaTuneMode = "interleave"
    NUMATUNEMODE_PREFERRED NumaTuneMode = "preferred"
    NUMATUNEMODE_STRICT NumaTuneMode = "strict"
)

type OpenStackNetworkProviderType string
const (
    OPENSTACKNETWORKPROVIDERTYPE_EXTERNAL OpenStackNetworkProviderType = "external"
    OPENSTACKNETWORKPROVIDERTYPE_NEUTRON OpenStackNetworkProviderType = "neutron"
)

type OpenstackVolumeAuthenticationKeyUsageType string
const (
    OPENSTACKVOLUMEAUTHENTICATIONKEYUSAGETYPE_CEPH OpenstackVolumeAuthenticationKeyUsageType = "ceph"
)

type OsType string
const (
    OSTYPE_OTHER OsType = "other"
    OSTYPE_OTHER_LINUX OsType = "other_linux"
    OSTYPE_RHEL_3 OsType = "rhel_3"
    OSTYPE_RHEL_3X64 OsType = "rhel_3x64"
    OSTYPE_RHEL_4 OsType = "rhel_4"
    OSTYPE_RHEL_4X64 OsType = "rhel_4x64"
    OSTYPE_RHEL_5 OsType = "rhel_5"
    OSTYPE_RHEL_5X64 OsType = "rhel_5x64"
    OSTYPE_RHEL_6 OsType = "rhel_6"
    OSTYPE_RHEL_6X64 OsType = "rhel_6x64"
    OSTYPE_UNASSIGNED OsType = "unassigned"
    OSTYPE_WINDOWS_2003 OsType = "windows_2003"
    OSTYPE_WINDOWS_2003X64 OsType = "windows_2003x64"
    OSTYPE_WINDOWS_2008 OsType = "windows_2008"
    OSTYPE_WINDOWS_2008R2X64 OsType = "windows_2008r2x64"
    OSTYPE_WINDOWS_2008X64 OsType = "windows_2008x64"
    OSTYPE_WINDOWS_2012X64 OsType = "windows_2012x64"
    OSTYPE_WINDOWS_7 OsType = "windows_7"
    OSTYPE_WINDOWS_7X64 OsType = "windows_7x64"
    OSTYPE_WINDOWS_8 OsType = "windows_8"
    OSTYPE_WINDOWS_8X64 OsType = "windows_8x64"
    OSTYPE_WINDOWS_XP OsType = "windows_xp"
)

type PayloadEncoding string
const (
    PAYLOADENCODING_BASE64 PayloadEncoding = "base64"
    PAYLOADENCODING_PLAINTEXT PayloadEncoding = "plaintext"
)

type PmProxyType string
const (
    PMPROXYTYPE_CLUSTER PmProxyType = "cluster"
    PMPROXYTYPE_DC PmProxyType = "dc"
    PMPROXYTYPE_OTHER_DC PmProxyType = "other_dc"
)

type PolicyUnitType string
const (
    POLICYUNITTYPE_FILTER PolicyUnitType = "filter"
    POLICYUNITTYPE_LOAD_BALANCING PolicyUnitType = "load_balancing"
    POLICYUNITTYPE_WEIGHT PolicyUnitType = "weight"
)

type PowerManagementStatus string
const (
    POWERMANAGEMENTSTATUS_OFF PowerManagementStatus = "off"
    POWERMANAGEMENTSTATUS_ON PowerManagementStatus = "on"
    POWERMANAGEMENTSTATUS_UNKNOWN PowerManagementStatus = "unknown"
)

type QcowVersion string
const (
    QCOWVERSION_QCOW2_V2 QcowVersion = "qcow2_v2"
    QCOWVERSION_QCOW2_V3 QcowVersion = "qcow2_v3"
)

type QosType string
const (
    QOSTYPE_CPU QosType = "cpu"
    QOSTYPE_HOSTNETWORK QosType = "hostnetwork"
    QOSTYPE_NETWORK QosType = "network"
    QOSTYPE_STORAGE QosType = "storage"
)

type QuotaModeType string
const (
    QUOTAMODETYPE_AUDIT QuotaModeType = "audit"
    QUOTAMODETYPE_DISABLED QuotaModeType = "disabled"
    QUOTAMODETYPE_ENABLED QuotaModeType = "enabled"
)

type ReportedDeviceType string
const (
    REPORTEDDEVICETYPE_NETWORK ReportedDeviceType = "network"
)

type ResolutionType string
const (
    RESOLUTIONTYPE_ADD ResolutionType = "add"
    RESOLUTIONTYPE_COPY ResolutionType = "copy"
)

type RngSource string
const (
    RNGSOURCE_HWRNG RngSource = "hwrng"
    RNGSOURCE_RANDOM RngSource = "random"
    RNGSOURCE_URANDOM RngSource = "urandom"
)

type RoleType string
const (
    ROLETYPE_ADMIN RoleType = "admin"
    ROLETYPE_USER RoleType = "user"
)

type ScsiGenericIO string
const (
    SCSIGENERICIO_FILTERED ScsiGenericIO = "filtered"
    SCSIGENERICIO_UNFILTERED ScsiGenericIO = "unfiltered"
)

type SeLinuxMode string
const (
    SELINUXMODE_DISABLED SeLinuxMode = "disabled"
    SELINUXMODE_ENFORCING SeLinuxMode = "enforcing"
    SELINUXMODE_PERMISSIVE SeLinuxMode = "permissive"
)

type SerialNumberPolicy string
const (
    SERIALNUMBERPOLICY_CUSTOM SerialNumberPolicy = "custom"
    SERIALNUMBERPOLICY_HOST SerialNumberPolicy = "host"
    SERIALNUMBERPOLICY_VM SerialNumberPolicy = "vm"
)

type SnapshotStatus string
const (
    SNAPSHOTSTATUS_IN_PREVIEW SnapshotStatus = "in_preview"
    SNAPSHOTSTATUS_LOCKED SnapshotStatus = "locked"
    SNAPSHOTSTATUS_OK SnapshotStatus = "ok"
)

type SnapshotType string
const (
    SNAPSHOTTYPE_ACTIVE SnapshotType = "active"
    SNAPSHOTTYPE_PREVIEW SnapshotType = "preview"
    SNAPSHOTTYPE_REGULAR SnapshotType = "regular"
    SNAPSHOTTYPE_STATELESS SnapshotType = "stateless"
)

type SpmStatus string
const (
    SPMSTATUS_CONTENDING SpmStatus = "contending"
    SPMSTATUS_NONE SpmStatus = "none"
    SPMSTATUS_SPM SpmStatus = "spm"
)

type SshAuthenticationMethod string
const (
    SSHAUTHENTICATIONMETHOD_PASSWORD SshAuthenticationMethod = "password"
    SSHAUTHENTICATIONMETHOD_PUBLICKEY SshAuthenticationMethod = "publickey"
)

type SsoMethod string
const (
    SSOMETHOD_GUEST_AGENT SsoMethod = "guest_agent"
)

type StatisticKind string
const (
    STATISTICKIND_COUNTER StatisticKind = "counter"
    STATISTICKIND_GAUGE StatisticKind = "gauge"
)

type StatisticUnit string
const (
    STATISTICUNIT_BITS_PER_SECOND StatisticUnit = "bits_per_second"
    STATISTICUNIT_BYTES StatisticUnit = "bytes"
    STATISTICUNIT_BYTES_PER_SECOND StatisticUnit = "bytes_per_second"
    STATISTICUNIT_COUNT_PER_SECOND StatisticUnit = "count_per_second"
    STATISTICUNIT_NONE StatisticUnit = "none"
    STATISTICUNIT_PERCENT StatisticUnit = "percent"
    STATISTICUNIT_SECONDS StatisticUnit = "seconds"
)

type StepEnum string
const (
    STEPENUM_EXECUTING StepEnum = "executing"
    STEPENUM_FINALIZING StepEnum = "finalizing"
    STEPENUM_REBALANCING_VOLUME StepEnum = "rebalancing_volume"
    STEPENUM_REMOVING_BRICKS StepEnum = "removing_bricks"
    STEPENUM_UNKNOWN StepEnum = "unknown"
    STEPENUM_VALIDATING StepEnum = "validating"
)

type StepStatus string
const (
    STEPSTATUS_ABORTED StepStatus = "aborted"
    STEPSTATUS_FAILED StepStatus = "failed"
    STEPSTATUS_FINISHED StepStatus = "finished"
    STEPSTATUS_STARTED StepStatus = "started"
    STEPSTATUS_UNKNOWN StepStatus = "unknown"
)

type StorageDomainStatus string
const (
    STORAGEDOMAINSTATUS_ACTIVATING StorageDomainStatus = "activating"
    STORAGEDOMAINSTATUS_ACTIVE StorageDomainStatus = "active"
    STORAGEDOMAINSTATUS_DETACHING StorageDomainStatus = "detaching"
    STORAGEDOMAINSTATUS_INACTIVE StorageDomainStatus = "inactive"
    STORAGEDOMAINSTATUS_LOCKED StorageDomainStatus = "locked"
    STORAGEDOMAINSTATUS_MAINTENANCE StorageDomainStatus = "maintenance"
    STORAGEDOMAINSTATUS_MIXED StorageDomainStatus = "mixed"
    STORAGEDOMAINSTATUS_PREPARING_FOR_MAINTENANCE StorageDomainStatus = "preparing_for_maintenance"
    STORAGEDOMAINSTATUS_UNATTACHED StorageDomainStatus = "unattached"
    STORAGEDOMAINSTATUS_UNKNOWN StorageDomainStatus = "unknown"
)

type StorageDomainType string
const (
    STORAGEDOMAINTYPE_DATA StorageDomainType = "data"
    STORAGEDOMAINTYPE_EXPORT StorageDomainType = "export"
    STORAGEDOMAINTYPE_IMAGE StorageDomainType = "image"
    STORAGEDOMAINTYPE_ISO StorageDomainType = "iso"
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
    STORAGETYPE_CINDER StorageType = "cinder"
    STORAGETYPE_FCP StorageType = "fcp"
    STORAGETYPE_GLANCE StorageType = "glance"
    STORAGETYPE_GLUSTERFS StorageType = "glusterfs"
    STORAGETYPE_ISCSI StorageType = "iscsi"
    STORAGETYPE_LOCALFS StorageType = "localfs"
    STORAGETYPE_NFS StorageType = "nfs"
    STORAGETYPE_POSIXFS StorageType = "posixfs"
)

type SwitchType string
const (
    SWITCHTYPE_LEGACY SwitchType = "legacy"
    SWITCHTYPE_OVS SwitchType = "ovs"
)

type TemplateStatus string
const (
    TEMPLATESTATUS_ILLEGAL TemplateStatus = "illegal"
    TEMPLATESTATUS_LOCKED TemplateStatus = "locked"
    TEMPLATESTATUS_OK TemplateStatus = "ok"
)

type TransportType string
const (
    TRANSPORTTYPE_RDMA TransportType = "rdma"
    TRANSPORTTYPE_TCP TransportType = "tcp"
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
    VALUETYPE_STRING ValueType = "string"
)

type VmAffinity string
const (
    VMAFFINITY_MIGRATABLE VmAffinity = "migratable"
    VMAFFINITY_PINNED VmAffinity = "pinned"
    VMAFFINITY_USER_MIGRATABLE VmAffinity = "user_migratable"
)

type VmDeviceType string
const (
    VMDEVICETYPE_CDROM VmDeviceType = "cdrom"
    VMDEVICETYPE_FLOPPY VmDeviceType = "floppy"
)

type VmPoolType string
const (
    VMPOOLTYPE_AUTOMATIC VmPoolType = "automatic"
    VMPOOLTYPE_MANUAL VmPoolType = "manual"
)

type VmStatus string
const (
    VMSTATUS_DOWN VmStatus = "down"
    VMSTATUS_IMAGE_LOCKED VmStatus = "image_locked"
    VMSTATUS_MIGRATING VmStatus = "migrating"
    VMSTATUS_NOT_RESPONDING VmStatus = "not_responding"
    VMSTATUS_PAUSED VmStatus = "paused"
    VMSTATUS_POWERING_DOWN VmStatus = "powering_down"
    VMSTATUS_POWERING_UP VmStatus = "powering_up"
    VMSTATUS_REBOOT_IN_PROGRESS VmStatus = "reboot_in_progress"
    VMSTATUS_RESTORING_STATE VmStatus = "restoring_state"
    VMSTATUS_SAVING_STATE VmStatus = "saving_state"
    VMSTATUS_SUSPENDED VmStatus = "suspended"
    VMSTATUS_UNASSIGNED VmStatus = "unassigned"
    VMSTATUS_UNKNOWN VmStatus = "unknown"
    VMSTATUS_UP VmStatus = "up"
    VMSTATUS_WAIT_FOR_LAUNCH VmStatus = "wait_for_launch"
)

type VmType string
const (
    VMTYPE_DESKTOP VmType = "desktop"
    VMTYPE_SERVER VmType = "server"
)

type VnicPassThroughMode string
const (
    VNICPASSTHROUGHMODE_DISABLED VnicPassThroughMode = "disabled"
    VNICPASSTHROUGHMODE_ENABLED VnicPassThroughMode = "enabled"
)

type WatchdogAction string
const (
    WATCHDOGACTION_DUMP WatchdogAction = "dump"
    WATCHDOGACTION_NONE WatchdogAction = "none"
    WATCHDOGACTION_PAUSE WatchdogAction = "pause"
    WATCHDOGACTION_POWEROFF WatchdogAction = "poweroff"
    WATCHDOGACTION_RESET WatchdogAction = "reset"
)

type WatchdogModel string
const (
    WATCHDOGMODEL_I6300ESB WatchdogModel = "i6300esb"
)

