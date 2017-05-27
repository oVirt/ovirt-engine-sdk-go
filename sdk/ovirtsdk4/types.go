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
    Port    int
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
    Active    int
    Total    int

}

type Bios struct {
    OvType
    BootMenu    BootMenu

}

type BlockStatistic struct {
    OvType
    Statistics    []Statistic

}

type Bonding struct {
    OvType
    ActiveSlave    HostNic
    AdPartnerMac    Mac
    Options    []Option
    Slaves    []HostNic

}

type Boot struct {
    OvType
    Devices    []BootDevice

}

type BootMenu struct {
    OvType
    Enabled    bool

}

type CloudInit struct {
    OvType
    AuthorizedKeys    []AuthorizedKey
    Files    []File
    Host    Host
    NetworkConfiguration    NetworkConfiguration
    RegenerateSshKeys    bool
    Timezone    string
    Users    []User

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
    Index    int
    Socket    int

}

type Cpu struct {
    OvType
    Architecture    Architecture
    Cores    []Core
    CpuTune    CpuTune
    Level    int
    Mode    CpuMode
    Name    string
    Speed    float64
    Topology    CpuTopology
    Type    string

}

type CpuTopology struct {
    OvType
    Cores    int
    Sockets    int
    Threads    int

}

type CpuTune struct {
    OvType
    VcpuPins    []VcpuPin

}

type CpuType struct {
    OvType
    Architecture    Architecture
    Level    int
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
    Monitors    int
    Port    int
    Proxy    string
    SecurePort    int
    SingleQxlPci    bool
    SmartcardEnabled    bool
    Type    DisplayType

}

type Dns struct {
    OvType
    SearchDomains    []Host
    Servers    []Host

}

type DnsResolverConfiguration struct {
    OvType
    NameServers    []String

}

type EntityProfileDetail struct {
    OvType
    ProfileDetails    []ProfileDetail

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
    Statistics    []Statistic

}

type GlusterBrickMemoryInfo struct {
    OvType
    MemoryPools    []GlusterMemoryPool

}

type GlusterClient struct {
    OvType
    BytesRead    int
    BytesWritten    int
    ClientPort    int
    HostName    string

}

type GracePeriod struct {
    OvType
    Expiry    int

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
    SupportedRngSources    []RngSource
    Uuid    string
    Version    string

}

type HighAvailability struct {
    OvType
    Enabled    bool
    Priority    int

}

type HostDevicePassthrough struct {
    OvType
    Enabled    bool

}

type HostNicVirtualFunctionsConfiguration struct {
    OvType
    AllNetworksAllowed    bool
    MaxNumberOfVirtualFunctions    int
    NumberOfVirtualFunctions    int

}

type HostedEngine struct {
    OvType
    Active    bool
    Configured    bool
    GlobalMaintenance    bool
    LocalMaintenance    bool
    Score    int

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
    NicConfigurations    []NicConfiguration
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
    Threads    int

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
    Networks    []Network
    StorageConnections    []StorageConnection

}

type IscsiDetails struct {
    OvType
    Address    string
    DiskId    string
    Initiator    string
    LunMapping    int
    Password    string
    Paths    int
    Port    int
    Portal    string
    ProductId    string
    Serial    string
    Size    int
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
    Steps    []Step

}

type KatelloErratum struct {
    Identified
    Comment    string
    Description    string
    Host    Host
    Id    string
    Issued    time.Time
    Name    string
    Packages    []Package
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
    DiscardMaxSize    int
    DiscardZeroesData    bool
    DiskId    string
    Id    string
    LunMapping    int
    Password    string
    Paths    int
    Port    int
    Portal    string
    ProductId    string
    Serial    string
    Size    int
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
    Ranges    []Range

}

type MemoryOverCommit struct {
    OvType
    Percent    int

}

type MemoryPolicy struct {
    OvType
    Ballooning    bool
    Guaranteed    int
    Max    int
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
    CustomValue    int

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
    Mtu    int
    Name    string
    NetworkLabels    []NetworkLabel
    Permissions    []Permission
    ProfileRequired    bool
    Qos    Qos
    Required    bool
    Status    NetworkStatus
    Stp    bool
    Usages    []NetworkUsage
    Vlan    Vlan
    VnicProfiles    []VnicProfile

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
    IpAddressAssignments    []IpAddressAssignment
    Name    string
    Network    Network
    Properties    []Property
    Qos    Qos
    ReportedConfigurations    []ReportedConfiguration

}

type NetworkConfiguration struct {
    OvType
    Dns    Dns
    Nics    []Nic

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
    ProfileDetails    []ProfileDetail

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
    Index    int
    Memory    int
    Name    string
    NodeDistance    string
    Statistics    []Statistic

}

type NumaNodePin struct {
    OvType
    HostNumaNode    NumaNode
    Index    int
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
    DnsServers    []String
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
    Properties    []Property

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
    Files    []File
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
    Agents    []Agent
    AutomaticPmEnabled    bool
    Enabled    bool
    KdumpDetection    bool
    Options    []Option
    Password    string
    PmProxies    []PmProxy
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
    BlockStatistics    []BlockStatistic
    Duration    int
    FopStatistics    []FopStatistic
    ProfileType    string
    Statistics    []Statistic

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
    CpuLimit    int
    DataCenter    DataCenter
    Description    string
    Id    string
    InboundAverage    int
    InboundBurst    int
    InboundPeak    int
    MaxIops    int
    MaxReadIops    int
    MaxReadThroughput    int
    MaxThroughput    int
    MaxWriteIops    int
    MaxWriteThroughput    int
    Name    string
    OutboundAverage    int
    OutboundAverageLinkshare    int
    OutboundAverageRealtime    int
    OutboundAverageUpperlimit    int
    OutboundBurst    int
    OutboundPeak    int
    Type    QosType

}

type Quota struct {
    Identified
    ClusterHardLimitPct    int
    ClusterSoftLimitPct    int
    Comment    string
    DataCenter    DataCenter
    Description    string
    Disks    []Disk
    Id    string
    Name    string
    Permissions    []Permission
    QuotaClusterLimits    []QuotaClusterLimit
    QuotaStorageLimits    []QuotaStorageLimit
    StorageHardLimitPct    int
    StorageSoftLimitPct    int
    Users    []User
    Vms    []Vm

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
    VcpuLimit    int
    VcpuUsage    int

}

type QuotaStorageLimit struct {
    Identified
    Comment    string
    Description    string
    Id    string
    Limit    int
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
    Bytes    int
    Period    int

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
    Ips    []Ip
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
    Permits    []Permit
    User    User

}

type SchedulingPolicy struct {
    Identified
    Balances    []Balance
    Comment    string
    DefaultPolicy    bool
    Description    string
    Filters    []Filter
    Id    string
    Locked    bool
    Name    string
    Properties    []Property
    Weight    []Weight

}

type SchedulingPolicyUnit struct {
    Identified
    Comment    string
    Description    string
    Enabled    bool
    Id    string
    Internal    bool
    Name    string
    Properties    []Property
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
    Threshold    int

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
    Priority    int
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
    Port    int
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
    Methods    []Method

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
    Values    []Value
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
    Number    int
    ParentStep    Step
    Progress    int
    StartTime    time.Time
    Statistics    []Statistic
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
    NfsRetrans    int
    NfsTimeo    int
    NfsVersion    NfsVersion
    Password    string
    Path    string
    Port    int
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
    Available    int
    Comment    string
    Committed    int
    CriticalSpaceActionBlocker    int
    DataCenter    DataCenter
    DataCenters    []DataCenter
    Description    string
    DiscardAfterDelete    bool
    DiskProfiles    []DiskProfile
    DiskSnapshots    []DiskSnapshot
    Disks    []Disk
    ExternalStatus    ExternalStatus
    Files    []File
    Host    Host
    Id    string
    Images    []Image
    Import    bool
    Master    bool
    Name    string
    Permissions    []Permission
    Status    StorageDomainStatus
    Storage    HostStorage
    StorageConnections    []StorageConnection
    StorageFormat    StorageFormat
    SupportsDiscard    bool
    SupportsDiscardZeroesData    bool
    Templates    []Template
    Type    StorageDomainType
    Used    int
    Vms    []Vm
    WarningLowSpaceIndicator    int
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
    VersionNumber    int

}

type Ticket struct {
    OvType
    Expiry    int
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
    Groups    []Group
    Id    string
    LastName    string
    LoggedIn    bool
    Name    string
    Namespace    string
    Password    string
    Permissions    []Permission
    Principal    string
    Roles    []Role
    SshPublicKeys    []SshPublicKey
    Tags    []Tag
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
    Vcpu    int

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
    Build    int
    Comment    string
    Description    string
    FullVersion    string
    Id    string
    Major    int
    Minor    int
    Name    string
    Revision    int

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
    Index    int
    Memory    int
    Name    string
    NodeDistance    string
    NumaNodePins    []NumaNodePin
    Statistics    []Statistic
    Vm    Vm

}

type Vlan struct {
    OvType
    Id    int

}

type VmBase struct {
    Identified
    Bios    Bios
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []CustomProperty
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
    Memory    int
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int
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
    Hosts    []Host

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
    MaxUserVms    int
    Name    string
    Permissions    []Permission
    PrestartedVms    int
    RngDevice    RngDevice
    Size    int
    SoundcardEnabled    bool
    Stateful    bool
    Template    Template
    Type    VmPoolType
    UseLatestTemplateVersion    bool
    Vm    Vm

}

type VmSummary struct {
    OvType
    Active    int
    Migrating    int
    Total    int

}

type VnicPassThrough struct {
    OvType
    Mode    VnicPassThroughMode

}

type VnicProfile struct {
    Identified
    Comment    string
    CustomProperties    []CustomProperty
    Description    string
    Id    string
    Migratable    bool
    Name    string
    Network    Network
    NetworkFilter    NetworkFilter
    PassThrough    VnicPassThrough
    Permissions    []Permission
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
    LogicalUnits    []LogicalUnit
    Name    string

}

type Weight struct {
    Identified
    Comment    string
    Description    string
    Factor    int
    Id    string
    Name    string
    SchedulingPolicy    SchedulingPolicy
    SchedulingPolicyUnit    SchedulingPolicyUnit

}

type Action struct {
    Identified
    AllowPartialImport    bool
    Async    bool
    Bricks    []GlusterBrick
    Certificates    []Certificate
    CheckConnectivity    bool
    Clone    bool
    Cluster    Cluster
    CollapseSnapshots    bool
    Comment    string
    ConnectivityTimeout    int
    DataCenter    DataCenter
    DeployHostedEngine    bool
    Description    string
    Details    GlusterVolumeProfileDetails
    DiscardSnapshots    bool
    Disk    Disk
    Disks    []Disk
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
    IscsiTargets    []String
    Job    Job
    LogicalUnits    []LogicalUnit
    MaintenanceEnabled    bool
    ModifiedBonds    []HostNic
    ModifiedLabels    []NetworkLabel
    ModifiedNetworkAttachments    []NetworkAttachment
    Name    string
    Option    Option
    Pause    bool
    PowerManagement    PowerManagement
    ProxyTicket    ProxyTicket
    Reason    string
    ReassignBadMacs    bool
    RemoteViewerConnectionFile    string
    RemovedBonds    []HostNic
    RemovedLabels    []NetworkLabel
    RemovedNetworkAttachments    []NetworkAttachment
    ResolutionType    string
    RestoreMemory    bool
    RootPassword    string
    Snapshot    Snapshot
    Ssh    Ssh
    Status    string
    StopGlusterService    bool
    StorageDomain    StorageDomain
    StorageDomains    []StorageDomain
    Succeeded    bool
    SynchronizedNetworkAttachments    []NetworkAttachment
    Template    Template
    Ticket    Ticket
    UndeployHostedEngine    bool
    UseCloudInit    bool
    UseSysprep    bool
    VirtualFunctionsConfiguration    HostNicVirtualFunctionsConfiguration
    Vm    Vm
    VnicProfileMappings    []VnicProfileMapping

}

type AffinityGroup struct {
    Identified
    Cluster    Cluster
    Comment    string
    Description    string
    Enforcing    bool
    Hosts    []Host
    HostsRule    AffinityRule
    Id    string
    Name    string
    Positive    bool
    Vms    []Vm
    VmsRule    AffinityRule

}

type AffinityLabel struct {
    Identified
    Comment    string
    Description    string
    Hosts    []Host
    Id    string
    Name    string
    ReadOnly    bool
    Vms    []Vm

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
    Options    []Option
    Order    int
    Password    string
    Port    int
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
    ProfileDetails    []ProfileDetail

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
    AffinityGroups    []AffinityGroup
    BallooningEnabled    bool
    Comment    string
    Cpu    Cpu
    CpuProfiles    []CpuProfile
    CustomSchedulingPolicyProperties    []Property
    DataCenter    DataCenter
    Description    string
    Display    Display
    ErrorHandling    ErrorHandling
    FencingPolicy    FencingPolicy
    GlusterHooks    []GlusterHook
    GlusterService    bool
    GlusterTunedProfile    string
    GlusterVolumes    []GlusterVolume
    HaReservation    bool
    Id    string
    Ksm    Ksm
    MacPool    MacPool
    MaintenanceReasonRequired    bool
    ManagementNetwork    Network
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    Name    string
    NetworkFilters    []NetworkFilter
    Networks    []Network
    OptionalReason    bool
    Permissions    []Permission
    RequiredRngSources    []RngSource
    SchedulingPolicy    SchedulingPolicy
    SerialNumber    SerialNumber
    SupportedVersions    []Version
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
    CpuTypes    []CpuType
    Description    string
    Id    string
    Name    string
    Permits    []Permit

}

type CpuProfile struct {
    Identified
    Cluster    Cluster
    Comment    string
    Description    string
    Id    string
    Name    string
    Permissions    []Permission
    Qos    Qos

}

type DataCenter struct {
    Identified
    Clusters    []Cluster
    Comment    string
    Description    string
    Id    string
    IscsiBonds    []IscsiBond
    Local    bool
    MacPool    MacPool
    Name    string
    Networks    []Network
    Permissions    []Permission
    Qoss    []Qos
    QuotaMode    QuotaModeType
    Quotas    []Quota
    Status    DataCenterStatus
    StorageDomains    []StorageDomain
    StorageFormat    StorageFormat
    SupportedVersions    []Version
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
    Vms    []Vm

}

type Disk struct {
    Device
    Active    bool
    ActualSize    int
    Alias    string
    Bootable    bool
    Comment    string
    Description    string
    DiskProfile    DiskProfile
    Format    DiskFormat
    Id    string
    ImageId    string
    InitialSize    int
    InstanceType    InstanceType
    Interface    DiskInterface
    LogicalName    string
    LunStorage    HostStorage
    Name    string
    OpenstackVolumeType    OpenStackVolumeType
    Permissions    []Permission
    PropagateErrors    bool
    ProvisionedSize    int
    QcowVersion    QcowVersion
    Quota    Quota
    ReadOnly    bool
    Sgio    ScsiGenericIO
    Shareable    bool
    Snapshot    Snapshot
    Sparse    bool
    Statistics    []Statistic
    Status    DiskStatus
    StorageDomain    StorageDomain
    StorageDomains    []StorageDomain
    StorageType    DiskStorageType
    Template    Template
    UsesScsiReservation    bool
    Vm    Vm
    Vms    []Vm
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
    Permissions    []Permission
    Qos    Qos
    StorageDomain    StorageDomain

}

type DiskSnapshot struct {
    Disk
    Active    bool
    ActualSize    int
    Alias    string
    Bootable    bool
    Comment    string
    Description    string
    Disk    Disk
    DiskProfile    DiskProfile
    Format    DiskFormat
    Id    string
    ImageId    string
    InitialSize    int
    InstanceType    InstanceType
    Interface    DiskInterface
    LogicalName    string
    LunStorage    HostStorage
    Name    string
    OpenstackVolumeType    OpenStackVolumeType
    Permissions    []Permission
    PropagateErrors    bool
    ProvisionedSize    int
    QcowVersion    QcowVersion
    Quota    Quota
    ReadOnly    bool
    Sgio    ScsiGenericIO
    Shareable    bool
    Snapshot    Snapshot
    Sparse    bool
    Statistics    []Statistic
    Status    DiskStatus
    StorageDomain    StorageDomain
    StorageDomains    []StorageDomain
    StorageType    DiskStorageType
    Template    Template
    UsesScsiReservation    bool
    Vm    Vm
    Vms    []Vm
    WipeAfterDelete    bool

}

type Domain struct {
    Identified
    Comment    string
    Description    string
    Groups    []Group
    Id    string
    Name    string
    User    User
    Users    []User

}

type Event struct {
    Identified
    Cluster    Cluster
    Code    int
    Comment    string
    CorrelationId    string
    CustomData    string
    CustomId    int
    DataCenter    DataCenter
    Description    string
    FloodRate    int
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
    Properties    []Property
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
    Position    int
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
    Vms    []Vm

}

type GlusterBrickAdvancedDetails struct {
    Device
    Comment    string
    Description    string
    Device    string
    FsName    string
    GlusterClients    []GlusterClient
    Id    string
    InstanceType    InstanceType
    MemoryPools    []GlusterMemoryPool
    MntOptions    string
    Name    string
    Pid    int
    Port    int
    Template    Template
    Vm    Vm
    Vms    []Vm

}

type GlusterHook struct {
    Identified
    Checksum    string
    Cluster    Cluster
    Comment    string
    ConflictStatus    int
    Conflicts    string
    Content    string
    ContentType    HookContentType
    Description    string
    GlusterCommand    string
    Id    string
    Name    string
    ServerHooks    []GlusterServerHook
    Stage    HookStage
    Status    GlusterHookStatus

}

type GlusterMemoryPool struct {
    Identified
    AllocCount    int
    ColdCount    int
    Comment    string
    Description    string
    HotCount    int
    Id    string
    MaxAlloc    int
    MaxStdalloc    int
    Name    string
    PaddedSize    int
    PoolMisses    int
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
    Bricks    []GlusterBrick
    Cluster    Cluster
    Comment    string
    Description    string
    DisperseCount    int
    Id    string
    Name    string
    Options    []Option
    RedundancyCount    int
    ReplicaCount    int
    Statistics    []Statistic
    Status    GlusterVolumeStatus
    StripeCount    int
    TransportTypes    []TransportType
    VolumeType    GlusterVolumeType

}

type GlusterVolumeProfileDetails struct {
    Identified
    BrickProfileDetails    []BrickProfileDetail
    Comment    string
    Description    string
    Id    string
    Name    string
    NfsProfileDetails    []NfsProfileDetail

}

type GraphicsConsole struct {
    Identified
    Address    string
    Comment    string
    Description    string
    Id    string
    InstanceType    InstanceType
    Name    string
    Port    int
    Protocol    GraphicsType
    Template    Template
    TlsPort    int
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
    Permissions    []Permission
    Roles    []Role
    Tags    []Tag

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
    AffinityLabels    []AffinityLabel
    Agents    []Agent
    AutoNumaStatus    AutoNumaStatus
    Certificate    Certificate
    Cluster    Cluster
    Comment    string
    Cpu    Cpu
    Description    string
    DevicePassthrough    HostDevicePassthrough
    Devices    []Device
    Display    Display
    ExternalHostProvider    ExternalHostProvider
    ExternalStatus    ExternalStatus
    HardwareInformation    HardwareInformation
    Hooks    []Hook
    HostedEngine    HostedEngine
    Id    string
    Iscsi    IscsiDetails
    KatelloErrata    []KatelloErratum
    KdumpStatus    KdumpStatus
    Ksm    Ksm
    LibvirtVersion    Version
    MaxSchedulingMemory    int
    Memory    int
    Name    string
    NetworkAttachments    []NetworkAttachment
    Nics    []Nic
    NumaNodes    []NumaNode
    NumaSupported    bool
    Os    OperatingSystem
    OverrideIptables    bool
    Permissions    []Permission
    Port    int
    PowerManagement    PowerManagement
    Protocol    HostProtocol
    RootPassword    string
    SeLinux    SeLinux
    Spm    Spm
    Ssh    Ssh
    Statistics    []Statistic
    Status    HostStatus
    StatusDetail    string
    StorageConnectionExtensions    []StorageConnectionExtension
    Storages    []HostStorage
    Summary    VmSummary
    Tags    []Tag
    TransparentHugePages    TransparentHugePages
    Type    HostType
    UnmanagedNetworks    []UnmanagedNetwork
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
    IommuGroup    int
    Name    string
    ParentDevice    HostDevice
    PhysicalFunction    HostDevice
    Placeholder    bool
    Product    Product
    Vendor    Vendor
    VirtualFunctions    int
    Vm    Vm

}

type HostNic struct {
    Identified
    AdAggregatorId    int
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
    Mtu    int
    Name    string
    Network    Network
    NetworkLabels    []NetworkLabel
    OverrideConfiguration    bool
    PhysicalFunction    HostNic
    Properties    []Property
    Qos    Qos
    Speed    int
    Statistics    []Statistic
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
    LogicalUnits    []LogicalUnit
    MountOptions    string
    Name    string
    NfsRetrans    int
    NfsTimeo    int
    NfsVersion    NfsVersion
    OverrideLuns    bool
    Password    string
    Path    string
    Port    int
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
    NetworkAttachments    []NetworkAttachment
    NetworkFilterParameters    []NetworkFilterParameter
    NetworkLabels    []NetworkLabel
    OnBoot    bool
    Plugged    bool
    ReportedDevices    []ReportedDevice
    Statistics    []Statistic
    Template    Template
    VirtualFunctionAllowedLabels    []NetworkLabel
    VirtualFunctionAllowedNetworks    []Network
    Vm    Vm
    Vms    []Vm
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
    Properties    []Property
    RequiresAuthentication    bool
    TenantName    string
    Url    string
    Username    string

}

type OpenStackVolumeProvider struct {
    OpenStackProvider
    AuthenticationKeys    []OpenstackVolumeAuthenticationKey
    AuthenticationUrl    string
    Certificates    []Certificate
    Comment    string
    DataCenter    DataCenter
    Description    string
    Id    string
    Name    string
    Password    string
    Properties    []Property
    RequiresAuthentication    bool
    TenantName    string
    Url    string
    Username    string
    VolumeTypes    []OpenStackVolumeType

}

type Template struct {
    VmBase
    Bios    Bios
    Cdroms    []Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []CustomProperty
    DeleteProtected    bool
    Description    string
    DiskAttachments    []DiskAttachment
    Display    Display
    Domain    Domain
    GraphicsConsoles    []GraphicsConsole
    HighAvailability    HighAvailability
    Id    string
    Initialization    Initialization
    Io    Io
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int
    Name    string
    Nics    []Nic
    Origin    string
    Os    OperatingSystem
    Permissions    []Permission
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
    Tags    []Tag
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    Version    TemplateVersion
    VirtioScsi    VirtioScsi
    Vm    Vm
    Watchdogs    []Watchdog

}

type Vm struct {
    VmBase
    AffinityLabels    []AffinityLabel
    Applications    []Application
    Bios    Bios
    Cdroms    []Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []CustomProperty
    DeleteProtected    bool
    Description    string
    DiskAttachments    []DiskAttachment
    Display    Display
    Domain    Domain
    ExternalHostProvider    ExternalHostProvider
    Floppies    []Floppy
    Fqdn    string
    GraphicsConsoles    []GraphicsConsole
    GuestOperatingSystem    GuestOperatingSystem
    GuestTimeZone    TimeZone
    HighAvailability    HighAvailability
    Host    Host
    HostDevices    []HostDevice
    Id    string
    Initialization    Initialization
    InstanceType    InstanceType
    Io    Io
    KatelloErrata    []KatelloErratum
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int
    Name    string
    NextRunConfigurationExists    bool
    Nics    []Nic
    NumaNodes    []NumaNode
    NumaTuneMode    NumaTuneMode
    Origin    string
    OriginalTemplate    Template
    Os    OperatingSystem
    Payloads    []Payload
    Permissions    []Permission
    PlacementPolicy    VmPlacementPolicy
    Quota    Quota
    ReportedDevices    []ReportedDevice
    RngDevice    RngDevice
    RunOnce    bool
    SerialNumber    SerialNumber
    Sessions    []Session
    SmallIcon    Icon
    Snapshots    []Snapshot
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    StartTime    time.Time
    Stateless    bool
    Statistics    []Statistic
    Status    VmStatus
    StatusDetail    string
    StopReason    string
    StopTime    time.Time
    StorageDomain    StorageDomain
    Tags    []Tag
    Template    Template
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    UseLatestTemplateVersion    bool
    VirtioScsi    VirtioScsi
    VmPool    VmPool
    Watchdogs    []Watchdog

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
    Vms    []Vm

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
    Vms    []Vm

}

type ExternalHostProvider struct {
    ExternalProvider
    AuthenticationUrl    string
    Certificates    []Certificate
    Comment    string
    ComputeResources    []ExternalComputeResource
    Description    string
    DiscoveredHosts    []ExternalDiscoveredHost
    HostGroups    []ExternalHostGroup
    Hosts    []Host
    Id    string
    Name    string
    Password    string
    Properties    []Property
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
    GlusterClients    []GlusterClient
    GlusterVolume    GlusterVolume
    Id    string
    InstanceType    InstanceType
    MemoryPools    []GlusterMemoryPool
    MntOptions    string
    Name    string
    Pid    int
    Port    int
    ServerId    string
    Statistics    []Statistic
    Status    GlusterBrickStatus
    Template    Template
    Vm    Vm
    Vms    []Vm

}

type InstanceType struct {
    Template
    Bios    Bios
    Cdroms    []Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []CustomProperty
    DeleteProtected    bool
    Description    string
    DiskAttachments    []DiskAttachment
    Display    Display
    Domain    Domain
    GraphicsConsoles    []GraphicsConsole
    HighAvailability    HighAvailability
    Id    string
    Initialization    Initialization
    Io    Io
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int
    Name    string
    Nics    []Nic
    Origin    string
    Os    OperatingSystem
    Permissions    []Permission
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
    Tags    []Tag
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    Version    TemplateVersion
    VirtioScsi    VirtioScsi
    Vm    Vm
    Watchdogs    []Watchdog

}

type OpenStackImageProvider struct {
    OpenStackProvider
    AuthenticationUrl    string
    Certificates    []Certificate
    Comment    string
    Description    string
    Id    string
    Images    []OpenStackImage
    Name    string
    Password    string
    Properties    []Property
    RequiresAuthentication    bool
    TenantName    string
    Url    string
    Username    string

}

type OpenStackNetworkProvider struct {
    OpenStackProvider
    AgentConfiguration    AgentConfiguration
    AuthenticationUrl    string
    Certificates    []Certificate
    Comment    string
    Description    string
    Id    string
    Name    string
    Networks    []OpenStackNetwork
    Password    string
    PluginType    NetworkPluginType
    Properties    []Property
    ReadOnly    bool
    RequiresAuthentication    bool
    Subnets    []OpenStackSubnet
    TenantName    string
    Type    OpenStackNetworkProviderType
    Url    string
    Username    string

}

type Snapshot struct {
    Vm
    AffinityLabels    []AffinityLabel
    Applications    []Application
    Bios    Bios
    Cdroms    []Cdrom
    Cluster    Cluster
    Comment    string
    Console    Console
    Cpu    Cpu
    CpuProfile    CpuProfile
    CpuShares    int
    CreationTime    time.Time
    CustomCompatibilityVersion    Version
    CustomCpuModel    string
    CustomEmulatedMachine    string
    CustomProperties    []CustomProperty
    Date    time.Time
    DeleteProtected    bool
    Description    string
    DiskAttachments    []DiskAttachment
    Display    Display
    Domain    Domain
    ExternalHostProvider    ExternalHostProvider
    Floppies    []Floppy
    Fqdn    string
    GraphicsConsoles    []GraphicsConsole
    GuestOperatingSystem    GuestOperatingSystem
    GuestTimeZone    TimeZone
    HighAvailability    HighAvailability
    Host    Host
    HostDevices    []HostDevice
    Id    string
    Initialization    Initialization
    InstanceType    InstanceType
    Io    Io
    KatelloErrata    []KatelloErratum
    LargeIcon    Icon
    Lease    StorageDomainLease
    Memory    int
    MemoryPolicy    MemoryPolicy
    Migration    MigrationOptions
    MigrationDowntime    int
    Name    string
    NextRunConfigurationExists    bool
    Nics    []Nic
    NumaNodes    []NumaNode
    NumaTuneMode    NumaTuneMode
    Origin    string
    OriginalTemplate    Template
    Os    OperatingSystem
    Payloads    []Payload
    Permissions    []Permission
    PersistMemorystate    bool
    PlacementPolicy    VmPlacementPolicy
    Quota    Quota
    ReportedDevices    []ReportedDevice
    RngDevice    RngDevice
    RunOnce    bool
    SerialNumber    SerialNumber
    Sessions    []Session
    SmallIcon    Icon
    SnapshotStatus    SnapshotStatus
    SnapshotType    SnapshotType
    Snapshots    []Snapshot
    SoundcardEnabled    bool
    Sso    Sso
    StartPaused    bool
    StartTime    time.Time
    Stateless    bool
    Statistics    []Statistic
    Status    VmStatus
    StatusDetail    string
    StopReason    string
    StopTime    time.Time
    StorageDomain    StorageDomain
    Tags    []Tag
    Template    Template
    TimeZone    TimeZone
    TunnelMigration    bool
    Type    VmType
    Usb    Usb
    UseLatestTemplateVersion    bool
    VirtioScsi    VirtioScsi
    Vm    Vm
    VmPool    VmPool
    Watchdogs    []Watchdog

}

@unique
class AccessProtocol(Enum):
    CIFS = 'cifs'
    GLUSTER = 'gluster'
    NFS = 'nfs'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class Architecture(Enum):
    PPC64 = 'ppc64'
    UNDEFINED = 'undefined'
    X86_64 = 'x86_64'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class AutoNumaStatus(Enum):
    DISABLE = 'disable'
    ENABLE = 'enable'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class BootDevice(Enum):
    CDROM = 'cdrom'
    HD = 'hd'
    NETWORK = 'network'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class BootProtocol(Enum):
    AUTOCONF = 'autoconf'
    DHCP = 'dhcp'
    NONE = 'none'
    STATIC = 'static'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ConfigurationType(Enum):
    OVF = 'ovf'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class CpuMode(Enum):
    CUSTOM = 'custom'
    HOST_MODEL = 'host_model'
    HOST_PASSTHROUGH = 'host_passthrough'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class CreationStatus(Enum):
    COMPLETE = 'complete'
    FAILED = 'failed'
    IN_PROGRESS = 'in_progress'
    PENDING = 'pending'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DataCenterStatus(Enum):
    CONTEND = 'contend'
    MAINTENANCE = 'maintenance'
    NOT_OPERATIONAL = 'not_operational'
    PROBLEMATIC = 'problematic'
    UNINITIALIZED = 'uninitialized'
    UP = 'up'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DiskFormat(Enum):
    COW = 'cow'
    RAW = 'raw'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DiskInterface(Enum):
    IDE = 'ide'
    SPAPR_VSCSI = 'spapr_vscsi'
    VIRTIO = 'virtio'
    VIRTIO_SCSI = 'virtio_scsi'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DiskStatus(Enum):
    ILLEGAL = 'illegal'
    LOCKED = 'locked'
    OK = 'ok'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DiskStorageType(Enum):
    CINDER = 'cinder'
    IMAGE = 'image'
    LUN = 'lun'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DiskType(Enum):
    DATA = 'data'
    SYSTEM = 'system'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class DisplayType(Enum):
    SPICE = 'spice'
    VNC = 'vnc'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class EntityExternalStatus(Enum):
    ERROR = 'error'
    FAILURE = 'failure'
    INFO = 'info'
    OK = 'ok'
    WARNING = 'warning'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ExternalStatus(Enum):
    ERROR = 'error'
    FAILURE = 'failure'
    INFO = 'info'
    OK = 'ok'
    WARNING = 'warning'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ExternalSystemType(Enum):
    GLUSTER = 'gluster'
    VDSM = 'vdsm'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ExternalVmProviderType(Enum):
    KVM = 'kvm'
    VMWARE = 'vmware'
    XEN = 'xen'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class FenceType(Enum):
    MANUAL = 'manual'
    RESTART = 'restart'
    START = 'start'
    STATUS = 'status'
    STOP = 'stop'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class GlusterBrickStatus(Enum):
    DOWN = 'down'
    UNKNOWN = 'unknown'
    UP = 'up'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class GlusterHookStatus(Enum):
    DISABLED = 'disabled'
    ENABLED = 'enabled'
    MISSING = 'missing'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class GlusterState(Enum):
    DOWN = 'down'
    UNKNOWN = 'unknown'
    UP = 'up'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class GlusterVolumeStatus(Enum):
    DOWN = 'down'
    UNKNOWN = 'unknown'
    UP = 'up'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class GlusterVolumeType(Enum):
    DISPERSE = 'disperse'
    DISTRIBUTE = 'distribute'
    DISTRIBUTED_DISPERSE = 'distributed_disperse'
    DISTRIBUTED_REPLICATE = 'distributed_replicate'
    DISTRIBUTED_STRIPE = 'distributed_stripe'
    DISTRIBUTED_STRIPED_REPLICATE = 'distributed_striped_replicate'
    REPLICATE = 'replicate'
    STRIPE = 'stripe'
    STRIPED_REPLICATE = 'striped_replicate'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class GraphicsType(Enum):
    SPICE = 'spice'
    VNC = 'vnc'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class HookContentType(Enum):
    BINARY = 'binary'
    TEXT = 'text'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class HookStage(Enum):
    POST = 'post'
    PRE = 'pre'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class HookStatus(Enum):
    DISABLED = 'disabled'
    ENABLED = 'enabled'
    MISSING = 'missing'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class HostProtocol(Enum):
    STOMP = 'stomp'
    XML = 'xml'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class HostStatus(Enum):
    CONNECTING = 'connecting'
    DOWN = 'down'
    ERROR = 'error'
    INITIALIZING = 'initializing'
    INSTALL_FAILED = 'install_failed'
    INSTALLING = 'installing'
    INSTALLING_OS = 'installing_os'
    KDUMPING = 'kdumping'
    MAINTENANCE = 'maintenance'
    NON_OPERATIONAL = 'non_operational'
    NON_RESPONSIVE = 'non_responsive'
    PENDING_APPROVAL = 'pending_approval'
    PREPARING_FOR_MAINTENANCE = 'preparing_for_maintenance'
    REBOOT = 'reboot'
    UNASSIGNED = 'unassigned'
    UP = 'up'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class HostType(Enum):
    OVIRT_NODE = 'ovirt_node'
    RHEL = 'rhel'
    RHEV_H = 'rhev_h'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ImageTransferDirection(Enum):
    DOWNLOAD = 'download'
    UPLOAD = 'upload'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ImageTransferPhase(Enum):
    CANCELLED = 'cancelled'
    FINALIZING_FAILURE = 'finalizing_failure'
    FINALIZING_SUCCESS = 'finalizing_success'
    FINISHED_FAILURE = 'finished_failure'
    FINISHED_SUCCESS = 'finished_success'
    INITIALIZING = 'initializing'
    PAUSED_SYSTEM = 'paused_system'
    PAUSED_USER = 'paused_user'
    RESUMING = 'resuming'
    TRANSFERRING = 'transferring'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class InheritableBoolean(Enum):
    FALSE = 'false'
    INHERIT = 'inherit'
    TRUE = 'true'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class IpVersion(Enum):
    V4 = 'v4'
    V6 = 'v6'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class JobStatus(Enum):
    ABORTED = 'aborted'
    FAILED = 'failed'
    FINISHED = 'finished'
    STARTED = 'started'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class KdumpStatus(Enum):
    DISABLED = 'disabled'
    ENABLED = 'enabled'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class LogSeverity(Enum):
    ALERT = 'alert'
    ERROR = 'error'
    NORMAL = 'normal'
    WARNING = 'warning'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class LunStatus(Enum):
    FREE = 'free'
    UNUSABLE = 'unusable'
    USED = 'used'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class MessageBrokerType(Enum):
    QPID = 'qpid'
    RABBIT_MQ = 'rabbit_mq'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class MigrateOnError(Enum):
    DO_NOT_MIGRATE = 'do_not_migrate'
    MIGRATE = 'migrate'
    MIGRATE_HIGHLY_AVAILABLE = 'migrate_highly_available'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class MigrationBandwidthAssignmentMethod(Enum):
    AUTO = 'auto'
    CUSTOM = 'custom'
    HYPERVISOR_DEFAULT = 'hypervisor_default'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NetworkPluginType(Enum):
    OPEN_VSWITCH = 'open_vswitch'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NetworkStatus(Enum):
    NON_OPERATIONAL = 'non_operational'
    OPERATIONAL = 'operational'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NetworkUsage(Enum):
    DEFAULT_ROUTE = 'default_route'
    DISPLAY = 'display'
    GLUSTER = 'gluster'
    MANAGEMENT = 'management'
    MIGRATION = 'migration'
    VM = 'vm'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NfsVersion(Enum):
    AUTO = 'auto'
    V3 = 'v3'
    V4 = 'v4'
    V4_1 = 'v4_1'
    V4_2 = 'v4_2'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NicInterface(Enum):
    E1000 = 'e1000'
    PCI_PASSTHROUGH = 'pci_passthrough'
    RTL8139 = 'rtl8139'
    RTL8139_VIRTIO = 'rtl8139_virtio'
    SPAPR_VLAN = 'spapr_vlan'
    VIRTIO = 'virtio'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NicStatus(Enum):
    DOWN = 'down'
    UP = 'up'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class NumaTuneMode(Enum):
    INTERLEAVE = 'interleave'
    PREFERRED = 'preferred'
    STRICT = 'strict'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class OpenStackNetworkProviderType(Enum):
    EXTERNAL = 'external'
    NEUTRON = 'neutron'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class OpenstackVolumeAuthenticationKeyUsageType(Enum):
    CEPH = 'ceph'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class OsType(Enum):
    OTHER = 'other'
    OTHER_LINUX = 'other_linux'
    RHEL_3 = 'rhel_3'
    RHEL_3X64 = 'rhel_3x64'
    RHEL_4 = 'rhel_4'
    RHEL_4X64 = 'rhel_4x64'
    RHEL_5 = 'rhel_5'
    RHEL_5X64 = 'rhel_5x64'
    RHEL_6 = 'rhel_6'
    RHEL_6X64 = 'rhel_6x64'
    UNASSIGNED = 'unassigned'
    WINDOWS_2003 = 'windows_2003'
    WINDOWS_2003X64 = 'windows_2003x64'
    WINDOWS_2008 = 'windows_2008'
    WINDOWS_2008R2X64 = 'windows_2008r2x64'
    WINDOWS_2008X64 = 'windows_2008x64'
    WINDOWS_2012X64 = 'windows_2012x64'
    WINDOWS_7 = 'windows_7'
    WINDOWS_7X64 = 'windows_7x64'
    WINDOWS_8 = 'windows_8'
    WINDOWS_8X64 = 'windows_8x64'
    WINDOWS_XP = 'windows_xp'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class PayloadEncoding(Enum):
    BASE64 = 'base64'
    PLAINTEXT = 'plaintext'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class PmProxyType(Enum):
    CLUSTER = 'cluster'
    DC = 'dc'
    OTHER_DC = 'other_dc'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class PolicyUnitType(Enum):
    FILTER = 'filter'
    LOAD_BALANCING = 'load_balancing'
    WEIGHT = 'weight'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class PowerManagementStatus(Enum):
    OFF = 'off'
    ON = 'on'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class QcowVersion(Enum):
    QCOW2_V2 = 'qcow2_v2'
    QCOW2_V3 = 'qcow2_v3'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class QosType(Enum):
    CPU = 'cpu'
    HOSTNETWORK = 'hostnetwork'
    NETWORK = 'network'
    STORAGE = 'storage'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class QuotaModeType(Enum):
    AUDIT = 'audit'
    DISABLED = 'disabled'
    ENABLED = 'enabled'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ReportedDeviceType(Enum):
    NETWORK = 'network'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ResolutionType(Enum):
    ADD = 'add'
    COPY = 'copy'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class RngSource(Enum):
    HWRNG = 'hwrng'
    RANDOM = 'random'
    URANDOM = 'urandom'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class RoleType(Enum):
    ADMIN = 'admin'
    USER = 'user'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ScsiGenericIO(Enum):
    FILTERED = 'filtered'
    UNFILTERED = 'unfiltered'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SeLinuxMode(Enum):
    DISABLED = 'disabled'
    ENFORCING = 'enforcing'
    PERMISSIVE = 'permissive'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SerialNumberPolicy(Enum):
    CUSTOM = 'custom'
    HOST = 'host'
    VM = 'vm'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SnapshotStatus(Enum):
    IN_PREVIEW = 'in_preview'
    LOCKED = 'locked'
    OK = 'ok'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SnapshotType(Enum):
    ACTIVE = 'active'
    PREVIEW = 'preview'
    REGULAR = 'regular'
    STATELESS = 'stateless'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SpmStatus(Enum):
    CONTENDING = 'contending'
    NONE = 'none'
    SPM = 'spm'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SshAuthenticationMethod(Enum):
    PASSWORD = 'password'
    PUBLICKEY = 'publickey'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SsoMethod(Enum):
    GUEST_AGENT = 'guest_agent'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StatisticKind(Enum):
    COUNTER = 'counter'
    GAUGE = 'gauge'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StatisticUnit(Enum):
    BITS_PER_SECOND = 'bits_per_second'
    BYTES = 'bytes'
    BYTES_PER_SECOND = 'bytes_per_second'
    COUNT_PER_SECOND = 'count_per_second'
    NONE = 'none'
    PERCENT = 'percent'
    SECONDS = 'seconds'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StepEnum(Enum):
    EXECUTING = 'executing'
    FINALIZING = 'finalizing'
    REBALANCING_VOLUME = 'rebalancing_volume'
    REMOVING_BRICKS = 'removing_bricks'
    UNKNOWN = 'unknown'
    VALIDATING = 'validating'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StepStatus(Enum):
    ABORTED = 'aborted'
    FAILED = 'failed'
    FINISHED = 'finished'
    STARTED = 'started'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StorageDomainStatus(Enum):
    ACTIVATING = 'activating'
    ACTIVE = 'active'
    DETACHING = 'detaching'
    INACTIVE = 'inactive'
    LOCKED = 'locked'
    MAINTENANCE = 'maintenance'
    MIXED = 'mixed'
    PREPARING_FOR_MAINTENANCE = 'preparing_for_maintenance'
    UNATTACHED = 'unattached'
    UNKNOWN = 'unknown'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StorageDomainType(Enum):
    DATA = 'data'
    EXPORT = 'export'
    IMAGE = 'image'
    ISO = 'iso'
    VOLUME = 'volume'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StorageFormat(Enum):
    V1 = 'v1'
    V2 = 'v2'
    V3 = 'v3'
    V4 = 'v4'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class StorageType(Enum):
    CINDER = 'cinder'
    FCP = 'fcp'
    GLANCE = 'glance'
    GLUSTERFS = 'glusterfs'
    ISCSI = 'iscsi'
    LOCALFS = 'localfs'
    NFS = 'nfs'
    POSIXFS = 'posixfs'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class SwitchType(Enum):
    LEGACY = 'legacy'
    OVS = 'ovs'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class TemplateStatus(Enum):
    ILLEGAL = 'illegal'
    LOCKED = 'locked'
    OK = 'ok'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class TransportType(Enum):
    RDMA = 'rdma'
    TCP = 'tcp'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class UsbType(Enum):
    LEGACY = 'legacy'
    NATIVE = 'native'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class ValueType(Enum):
    DECIMAL = 'decimal'
    INTEGER = 'integer'
    STRING = 'string'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class VmAffinity(Enum):
    MIGRATABLE = 'migratable'
    PINNED = 'pinned'
    USER_MIGRATABLE = 'user_migratable'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class VmDeviceType(Enum):
    CDROM = 'cdrom'
    FLOPPY = 'floppy'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class VmPoolType(Enum):
    AUTOMATIC = 'automatic'
    MANUAL = 'manual'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class VmStatus(Enum):
    DOWN = 'down'
    IMAGE_LOCKED = 'image_locked'
    MIGRATING = 'migrating'
    NOT_RESPONDING = 'not_responding'
    PAUSED = 'paused'
    POWERING_DOWN = 'powering_down'
    POWERING_UP = 'powering_up'
    REBOOT_IN_PROGRESS = 'reboot_in_progress'
    RESTORING_STATE = 'restoring_state'
    SAVING_STATE = 'saving_state'
    SUSPENDED = 'suspended'
    UNASSIGNED = 'unassigned'
    UNKNOWN = 'unknown'
    UP = 'up'
    WAIT_FOR_LAUNCH = 'wait_for_launch'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class VmType(Enum):
    DESKTOP = 'desktop'
    SERVER = 'server'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class VnicPassThroughMode(Enum):
    DISABLED = 'disabled'
    ENABLED = 'enabled'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class WatchdogAction(Enum):
    DUMP = 'dump'
    NONE = 'none'
    PAUSE = 'pause'
    POWEROFF = 'poweroff'
    RESET = 'reset'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


@unique
class WatchdogModel(Enum):
    I6300ESB = 'i6300esb'

    def __init__(self, image):
        self._image = image

    def __str__(self):
        return self._image


