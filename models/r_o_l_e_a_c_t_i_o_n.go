// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ROLEACTION r o l e a c t i o n
//
// swagger:model ROLE_ACTION
type ROLEACTION string

func NewROLEACTION(value ROLEACTION) *ROLEACTION {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ROLEACTION.
func (m ROLEACTION) Pointer() *ROLEACTION {
	return &m
}

const (

	// ROLEACTIONAsterisk captures enum value "*"
	ROLEACTIONAsterisk ROLEACTION = "*"

	// ROLEACTIONMANAGEDATACENTER captures enum value "MANAGE_DATA_CENTER"
	ROLEACTIONMANAGEDATACENTER ROLEACTION = "MANAGE_DATA_CENTER"

	// ROLEACTIONMANAGECLUSTERCONNECTION captures enum value "MANAGE_CLUSTER_CONNECTION"
	ROLEACTIONMANAGECLUSTERCONNECTION ROLEACTION = "MANAGE_CLUSTER_CONNECTION"

	// ROLEACTIONMANAGESTORAGECLUSTERCONNECTION captures enum value "MANAGE_STORAGE_CLUSTER_CONNECTION"
	ROLEACTIONMANAGESTORAGECLUSTERCONNECTION ROLEACTION = "MANAGE_STORAGE_CLUSTER_CONNECTION"

	// ROLEACTIONMANAGEHOST captures enum value "MANAGE_HOST"
	ROLEACTIONMANAGEHOST ROLEACTION = "MANAGE_HOST"

	// ROLEACTIONMANAGENICMTU captures enum value "MANAGE_NIC_MTU"
	ROLEACTIONMANAGENICMTU ROLEACTION = "MANAGE_NIC_MTU"

	// ROLEACTIONMANAGEDISK captures enum value "MANAGE_DISK"
	ROLEACTIONMANAGEDISK ROLEACTION = "MANAGE_DISK"

	// ROLEACTIONMANAGEHARDWARETOPO captures enum value "MANAGE_HARDWARE_TOPO"
	ROLEACTIONMANAGEHARDWARETOPO ROLEACTION = "MANAGE_HARDWARE_TOPO"

	// ROLEACTIONMANAGEUSBDEVICE captures enum value "MANAGE_USB_DEVICE"
	ROLEACTIONMANAGEUSBDEVICE ROLEACTION = "MANAGE_USB_DEVICE"

	// ROLEACTIONMANAGEGPUDEVICE captures enum value "MANAGE_GPU_DEVICE"
	ROLEACTIONMANAGEGPUDEVICE ROLEACTION = "MANAGE_GPU_DEVICE"

	// ROLEACTIONMANAGEVDS captures enum value "MANAGE_VDS"
	ROLEACTIONMANAGEVDS ROLEACTION = "MANAGE_VDS"

	// ROLEACTIONMANAGEVLAN captures enum value "MANAGE_VLAN"
	ROLEACTIONMANAGEVLAN ROLEACTION = "MANAGE_VLAN"

	// ROLEACTIONMANAGESYSTEMVLAN captures enum value "MANAGE_SYSTEM_VLAN"
	ROLEACTIONMANAGESYSTEMVLAN ROLEACTION = "MANAGE_SYSTEM_VLAN"

	// ROLEACTIONMANAGEISCSIDATASTORE captures enum value "MANAGE_ISCSI_DATA_STORE"
	ROLEACTIONMANAGEISCSIDATASTORE ROLEACTION = "MANAGE_ISCSI_DATA_STORE"

	// ROLEACTIONMANAGENFSDATASTORE captures enum value "MANAGE_NFS_DATA_STORE"
	ROLEACTIONMANAGENFSDATASTORE ROLEACTION = "MANAGE_NFS_DATA_STORE"

	// ROLEACTIONMANAGENVMFDATASTORE captures enum value "MANAGE_NVMF_DATA_STORE"
	ROLEACTIONMANAGENVMFDATASTORE ROLEACTION = "MANAGE_NVMF_DATA_STORE"

	// ROLEACTIONCREATEVM captures enum value "CREATE_VM"
	ROLEACTIONCREATEVM ROLEACTION = "CREATE_VM"

	// ROLEACTIONUPDATEVM captures enum value "UPDATE_VM"
	ROLEACTIONUPDATEVM ROLEACTION = "UPDATE_VM"

	// ROLEACTIONDELETEVM captures enum value "DELETE_VM"
	ROLEACTIONDELETEVM ROLEACTION = "DELETE_VM"

	// ROLEACTIONUPDATEVMADVANCEDSETTING captures enum value "UPDATE_VM_ADVANCED_SETTING"
	ROLEACTIONUPDATEVMADVANCEDSETTING ROLEACTION = "UPDATE_VM_ADVANCED_SETTING"

	// ROLEACTIONUPDATEVMGUEST captures enum value "UPDATE_VM_GUEST"
	ROLEACTIONUPDATEVMGUEST ROLEACTION = "UPDATE_VM_GUEST"

	// ROLEACTIONVMOPERATIONOPENTERMINAL captures enum value "VM_OPERATION_OPEN_TERMINAL"
	ROLEACTIONVMOPERATIONOPENTERMINAL ROLEACTION = "VM_OPERATION_OPEN_TERMINAL"

	// ROLEACTIONVMOPERATIONMIGRATE captures enum value "VM_OPERATION_MIGRATE"
	ROLEACTIONVMOPERATIONMIGRATE ROLEACTION = "VM_OPERATION_MIGRATE"

	// ROLEACTIONVMOPERATIONVMFOLDER captures enum value "VM_OPERATION_VM_FOLDER"
	ROLEACTIONVMOPERATIONVMFOLDER ROLEACTION = "VM_OPERATION_VM_FOLDER"

	// ROLEACTIONVMOPERATIONVMPOWER captures enum value "VM_OPERATION_VM_POWER"
	ROLEACTIONVMOPERATIONVMPOWER ROLEACTION = "VM_OPERATION_VM_POWER"

	// ROLEACTIONVMOPERATIONCLONE captures enum value "VM_OPERATION_CLONE"
	ROLEACTIONVMOPERATIONCLONE ROLEACTION = "VM_OPERATION_CLONE"

	// ROLEACTIONVMOPERATIONINSTALLTOOLS captures enum value "VM_OPERATION_INSTALL_TOOLS"
	ROLEACTIONVMOPERATIONINSTALLTOOLS ROLEACTION = "VM_OPERATION_INSTALL_TOOLS"

	// ROLEACTIONVMIMPORTEXPORT captures enum value "VM_IMPORT_EXPORT"
	ROLEACTIONVMIMPORTEXPORT ROLEACTION = "VM_IMPORT_EXPORT"

	// ROLEACTIONCREATEVMTEMPLATE captures enum value "CREATE_VM_TEMPLATE"
	ROLEACTIONCREATEVMTEMPLATE ROLEACTION = "CREATE_VM_TEMPLATE"

	// ROLEACTIONMANAGEVMTEMPLATE captures enum value "MANAGE_VM_TEMPLATE"
	ROLEACTIONMANAGEVMTEMPLATE ROLEACTION = "MANAGE_VM_TEMPLATE"

	// ROLEACTIONVMTEMPLATEIMPORTEXPORT captures enum value "VM_TEMPLATE_IMPORT_EXPORT"
	ROLEACTIONVMTEMPLATEIMPORTEXPORT ROLEACTION = "VM_TEMPLATE_IMPORT_EXPORT"

	// ROLEACTIONMANAGEVMSNAPSHOT captures enum value "MANAGE_VM_SNAPSHOT"
	ROLEACTIONMANAGEVMSNAPSHOT ROLEACTION = "MANAGE_VM_SNAPSHOT"

	// ROLEACTIONMANAGEVMVOLUME captures enum value "MANAGE_VM_VOLUME"
	ROLEACTIONMANAGEVMVOLUME ROLEACTION = "MANAGE_VM_VOLUME"

	// ROLEACTIONVMVOLUMEIMPORTEXPORT captures enum value "VM_VOLUME_IMPORT_EXPORT"
	ROLEACTIONVMVOLUMEIMPORTEXPORT ROLEACTION = "VM_VOLUME_IMPORT_EXPORT"

	// ROLEACTIONMANAGEISO captures enum value "MANAGE_ISO"
	ROLEACTIONMANAGEISO ROLEACTION = "MANAGE_ISO"

	// ROLEACTIONDOWNLOADISO captures enum value "DOWNLOAD_ISO"
	ROLEACTIONDOWNLOADISO ROLEACTION = "DOWNLOAD_ISO"

	// ROLEACTIONQUERYSENSITIVERESOURCELIST captures enum value "QUERY_SENSITIVE_RESOURCE_LIST"
	ROLEACTIONQUERYSENSITIVERESOURCELIST ROLEACTION = "QUERY_SENSITIVE_RESOURCE_LIST"

	// ROLEACTIONQUERYSENSITIVERESOURCE captures enum value "QUERY_SENSITIVE_RESOURCE"
	ROLEACTIONQUERYSENSITIVERESOURCE ROLEACTION = "QUERY_SENSITIVE_RESOURCE"

	// ROLEACTIONMANAGESENSITIVERESOURCE captures enum value "MANAGE_SENSITIVE_RESOURCE"
	ROLEACTIONMANAGESENSITIVERESOURCE ROLEACTION = "MANAGE_SENSITIVE_RESOURCE"

	// ROLEACTIONMANAGEVMPLACEMENTGROUP captures enum value "MANAGE_VM_PLACEMENT_GROUP"
	ROLEACTIONMANAGEVMPLACEMENTGROUP ROLEACTION = "MANAGE_VM_PLACEMENT_GROUP"

	// ROLEACTIONMANAGESNAPSHOTPLAN captures enum value "MANAGE_SNAPSHOT_PLAN"
	ROLEACTIONMANAGESNAPSHOTPLAN ROLEACTION = "MANAGE_SNAPSHOT_PLAN"

	// ROLEACTIONMANAGEALERT captures enum value "MANAGE_ALERT"
	ROLEACTIONMANAGEALERT ROLEACTION = "MANAGE_ALERT"

	// ROLEACTIONMANAGEMONITORVIEW captures enum value "MANAGE_MONITOR_VIEW"
	ROLEACTIONMANAGEMONITORVIEW ROLEACTION = "MANAGE_MONITOR_VIEW"

	// ROLEACTIONMANAGEENTITYFILTER captures enum value "MANAGE_ENTITY_FILTER"
	ROLEACTIONMANAGEENTITYFILTER ROLEACTION = "MANAGE_ENTITY_FILTER"

	// ROLEACTIONMANAGECLUSTERBASICINFO captures enum value "MANAGE_CLUSTER_BASIC_INFO"
	ROLEACTIONMANAGECLUSTERBASICINFO ROLEACTION = "MANAGE_CLUSTER_BASIC_INFO"

	// ROLEACTIONMANAGECLUSTERLICENCE captures enum value "MANAGE_CLUSTER_LICENCE"
	ROLEACTIONMANAGECLUSTERLICENCE ROLEACTION = "MANAGE_CLUSTER_LICENCE"

	// ROLEACTIONMANAGECLUSTERSNMPTRANSPORT captures enum value "MANAGE_CLUSTER_SNMP_TRANSPORT"
	ROLEACTIONMANAGECLUSTERSNMPTRANSPORT ROLEACTION = "MANAGE_CLUSTER_SNMP_TRANSPORT"

	// ROLEACTIONMANAGESNMPTRAP captures enum value "MANAGE_SNMP_TRAP"
	ROLEACTIONMANAGESNMPTRAP ROLEACTION = "MANAGE_SNMP_TRAP"

	// ROLEACTIONMANAGECLUSTERVIP captures enum value "MANAGE_CLUSTER_VIP"
	ROLEACTIONMANAGECLUSTERVIP ROLEACTION = "MANAGE_CLUSTER_VIP"

	// ROLEACTIONMANAGECLUSTERMANAGEMENTIP captures enum value "MANAGE_CLUSTER_MANAGEMENT_IP"
	ROLEACTIONMANAGECLUSTERMANAGEMENTIP ROLEACTION = "MANAGE_CLUSTER_MANAGEMENT_IP"

	// ROLEACTIONMANAGEDNSSERVER captures enum value "MANAGE_DNS_SERVER"
	ROLEACTIONMANAGEDNSSERVER ROLEACTION = "MANAGE_DNS_SERVER"

	// ROLEACTIONMANAGENTPSERVER captures enum value "MANAGE_NTP_SERVER"
	ROLEACTIONMANAGENTPSERVER ROLEACTION = "MANAGE_NTP_SERVER"

	// ROLEACTIONMANAGEIPMI captures enum value "MANAGE_IPMI"
	ROLEACTIONMANAGEIPMI ROLEACTION = "MANAGE_IPMI"

	// ROLEACTIONMANAGECLUSTERVMCPUMODEL captures enum value "MANAGE_CLUSTER_VM_CPU_MODEL"
	ROLEACTIONMANAGECLUSTERVMCPUMODEL ROLEACTION = "MANAGE_CLUSTER_VM_CPU_MODEL"

	// ROLEACTIONMANAGECLUSTERVMTOOLS captures enum value "MANAGE_CLUSTER_VM_TOOLS"
	ROLEACTIONMANAGECLUSTERVMTOOLS ROLEACTION = "MANAGE_CLUSTER_VM_TOOLS"

	// ROLEACTIONMANAGECLUSTERHOTMIGRATION captures enum value "MANAGE_CLUSTER_HOT_MIGRATION"
	ROLEACTIONMANAGECLUSTERHOTMIGRATION ROLEACTION = "MANAGE_CLUSTER_HOT_MIGRATION"

	// ROLEACTIONMANAGECLUSTERHA captures enum value "MANAGE_CLUSTER_HA"
	ROLEACTIONMANAGECLUSTERHA ROLEACTION = "MANAGE_CLUSTER_HA"

	// ROLEACTIONMANAGESSLCERTIFICATE captures enum value "MANAGE_SSL_CERTIFICATE"
	ROLEACTIONMANAGESSLCERTIFICATE ROLEACTION = "MANAGE_SSL_CERTIFICATE"

	// ROLEACTIONMANAGELOGCOLLECTION captures enum value "MANAGE_LOG_COLLECTION"
	ROLEACTIONMANAGELOGCOLLECTION ROLEACTION = "MANAGE_LOG_COLLECTION"

	// ROLEACTIONMANAGESYSLOG captures enum value "MANAGE_SYSLOG"
	ROLEACTIONMANAGESYSLOG ROLEACTION = "MANAGE_SYSLOG"

	// ROLEACTIONMANAGELOGFIND captures enum value "MANAGE_LOG_FIND"
	ROLEACTIONMANAGELOGFIND ROLEACTION = "MANAGE_LOG_FIND"

	// ROLEACTIONMANAGELABEL captures enum value "MANAGE_LABEL"
	ROLEACTIONMANAGELABEL ROLEACTION = "MANAGE_LABEL"

	// ROLEACTIONMANAGEUSERANDROLE captures enum value "MANAGE_USER_AND_ROLE"
	ROLEACTIONMANAGEUSERANDROLE ROLEACTION = "MANAGE_USER_AND_ROLE"

	// ROLEACTIONMANAGEPASSWORDSETTINGS captures enum value "MANAGE_PASSWORD_SETTINGS"
	ROLEACTIONMANAGEPASSWORDSETTINGS ROLEACTION = "MANAGE_PASSWORD_SETTINGS"

	// ROLEACTIONMANAGEACCESSCONTROL captures enum value "MANAGE_ACCESS_CONTROL"
	ROLEACTIONMANAGEACCESSCONTROL ROLEACTION = "MANAGE_ACCESS_CONTROL"

	// ROLEACTIONMANAGESESSIONEXPIRATION captures enum value "MANAGE_SESSION_EXPIRATION"
	ROLEACTIONMANAGESESSIONEXPIRATION ROLEACTION = "MANAGE_SESSION_EXPIRATION"

	// ROLEACTIONMANAGEVCENTERASSOCIATION captures enum value "MANAGE_VCENTER_ASSOCIATION"
	ROLEACTIONMANAGEVCENTERASSOCIATION ROLEACTION = "MANAGE_VCENTER_ASSOCIATION"

	// ROLEACTIONMANAGEESXIASSOCIATION captures enum value "MANAGE_ESXI_ASSOCIATION"
	ROLEACTIONMANAGEESXIASSOCIATION ROLEACTION = "MANAGE_ESXI_ASSOCIATION"

	// ROLEACTIONMANAGEAUDITLOG captures enum value "MANAGE_AUDIT_LOG"
	ROLEACTIONMANAGEAUDITLOG ROLEACTION = "MANAGE_AUDIT_LOG"

	// ROLEACTIONMANAGEALERTEMAILSETTING captures enum value "MANAGE_ALERT_EMAIL_SETTING"
	ROLEACTIONMANAGEALERTEMAILSETTING ROLEACTION = "MANAGE_ALERT_EMAIL_SETTING"

	// ROLEACTIONMANAGESMTPSERVER captures enum value "MANAGE_SMTP_SERVER"
	ROLEACTIONMANAGESMTPSERVER ROLEACTION = "MANAGE_SMTP_SERVER"

	// ROLEACTIONMANAGECLUSTERUPGRADE captures enum value "MANAGE_CLUSTER_UPGRADE"
	ROLEACTIONMANAGECLUSTERUPGRADE ROLEACTION = "MANAGE_CLUSTER_UPGRADE"

	// ROLEACTIONMANAGEVMRECYCLEBINSETTING captures enum value "MANAGE_VM_RECYCLE_BIN_SETTING"
	ROLEACTIONMANAGEVMRECYCLEBINSETTING ROLEACTION = "MANAGE_VM_RECYCLE_BIN_SETTING"

	// ROLEACTIONMANAGEREPORT captures enum value "MANAGE_REPORT"
	ROLEACTIONMANAGEREPORT ROLEACTION = "MANAGE_REPORT"

	// ROLEACTIONMANAGESHARINGVMTOOLS captures enum value "MANAGE_SHARING_VM_TOOLS"
	ROLEACTIONMANAGESHARINGVMTOOLS ROLEACTION = "MANAGE_SHARING_VM_TOOLS"

	// ROLEACTIONMANAGEADVANCEDMONITOR captures enum value "MANAGE_ADVANCED_MONITOR"
	ROLEACTIONMANAGEADVANCEDMONITOR ROLEACTION = "MANAGE_ADVANCED_MONITOR"

	// ROLEACTIONMANAGETHIRDPARTYDRIVER captures enum value "MANAGE_THIRD_PARTY_DRIVER"
	ROLEACTIONMANAGETHIRDPARTYDRIVER ROLEACTION = "MANAGE_THIRD_PARTY_DRIVER"

	// ROLEACTIONMANAGEORGANIZATIONNAME captures enum value "MANAGE_ORGANIZATION_NAME"
	ROLEACTIONMANAGEORGANIZATIONNAME ROLEACTION = "MANAGE_ORGANIZATION_NAME"

	// ROLEACTIONMANAGECLOUDTOWERLICENSE captures enum value "MANAGE_CLOUD_TOWER_LICENSE"
	ROLEACTIONMANAGECLOUDTOWERLICENSE ROLEACTION = "MANAGE_CLOUD_TOWER_LICENSE"

	// ROLEACTIONMANAGECONSISTENCYGROUP captures enum value "MANAGE_CONSISTENCY_GROUP"
	ROLEACTIONMANAGECONSISTENCYGROUP ROLEACTION = "MANAGE_CONSISTENCY_GROUP"

	// ROLEACTIONMANAGENIC captures enum value "MANAGE_NIC"
	ROLEACTIONMANAGENIC ROLEACTION = "MANAGE_NIC"

	// ROLEACTIONMANAGECLUSTERISCSI captures enum value "MANAGE_CLUSTER_ISCSI"
	ROLEACTIONMANAGECLUSTERISCSI ROLEACTION = "MANAGE_CLUSTER_ISCSI"

	// ROLEACTIONMANAGEBACKUPLICENSE captures enum value "MANAGE_BACKUP_LICENSE"
	ROLEACTIONMANAGEBACKUPLICENSE ROLEACTION = "MANAGE_BACKUP_LICENSE"

	// ROLEACTIONMANAGEBACKUPPACKAGE captures enum value "MANAGE_BACKUP_PACKAGE"
	ROLEACTIONMANAGEBACKUPPACKAGE ROLEACTION = "MANAGE_BACKUP_PACKAGE"

	// ROLEACTIONMANAGEBACKUPSERVICE captures enum value "MANAGE_BACKUP_SERVICE"
	ROLEACTIONMANAGEBACKUPSERVICE ROLEACTION = "MANAGE_BACKUP_SERVICE"

	// ROLEACTIONMANAGEBACKUPSTOREREPOSITORY captures enum value "MANAGE_BACKUP_STORE_REPOSITORY"
	ROLEACTIONMANAGEBACKUPSTOREREPOSITORY ROLEACTION = "MANAGE_BACKUP_STORE_REPOSITORY"

	// ROLEACTIONMANAGEBACKUPPLAN captures enum value "MANAGE_BACKUP_PLAN"
	ROLEACTIONMANAGEBACKUPPLAN ROLEACTION = "MANAGE_BACKUP_PLAN"

	// ROLEACTIONMANAGEBACKUPTASK captures enum value "MANAGE_BACKUP_TASK"
	ROLEACTIONMANAGEBACKUPTASK ROLEACTION = "MANAGE_BACKUP_TASK"

	// ROLEACTIONMANAGEBACKUPRESTOREPOINT captures enum value "MANAGE_BACKUP_RESTORE_POINT"
	ROLEACTIONMANAGEBACKUPRESTOREPOINT ROLEACTION = "MANAGE_BACKUP_RESTORE_POINT"

	// ROLEACTIONMANAGEBACKUPRESTOREPOINTTASK captures enum value "MANAGE_BACKUP_RESTORE_POINT_TASK"
	ROLEACTIONMANAGEBACKUPRESTOREPOINTTASK ROLEACTION = "MANAGE_BACKUP_RESTORE_POINT_TASK"

	// ROLEACTIONMANAGESECURITYPOLICY captures enum value "MANAGE_SECURITY_POLICY"
	ROLEACTIONMANAGESECURITYPOLICY ROLEACTION = "MANAGE_SECURITY_POLICY"

	// ROLEACTIONMANAGESECURITYGROUP captures enum value "MANAGE_SECURITY_GROUP"
	ROLEACTIONMANAGESECURITYGROUP ROLEACTION = "MANAGE_SECURITY_GROUP"

	// ROLEACTIONISOLATEVM captures enum value "ISOLATE_VM"
	ROLEACTIONISOLATEVM ROLEACTION = "ISOLATE_VM"

	// ROLEACTIONMANAGEEVEROUTELICENSE captures enum value "MANAGE_EVEROUTE_LICENSE"
	ROLEACTIONMANAGEEVEROUTELICENSE ROLEACTION = "MANAGE_EVEROUTE_LICENSE"

	// ROLEACTIONMANAGEEVEROUTEPACKAGE captures enum value "MANAGE_EVEROUTE_PACKAGE"
	ROLEACTIONMANAGEEVEROUTEPACKAGE ROLEACTION = "MANAGE_EVEROUTE_PACKAGE"

	// ROLEACTIONDEPLOYEVEROUTECLUSTER captures enum value "DEPLOY_EVEROUTE_CLUSTER"
	ROLEACTIONDEPLOYEVEROUTECLUSTER ROLEACTION = "DEPLOY_EVEROUTE_CLUSTER"

	// ROLEACTIONUNDEPLOYEVEROUTECLUSTER captures enum value "UNDEPLOY_EVEROUTE_CLUSTER"
	ROLEACTIONUNDEPLOYEVEROUTECLUSTER ROLEACTION = "UNDEPLOY_EVEROUTE_CLUSTER"

	// ROLEACTIONUPDATEEVEROUTECLUSTER captures enum value "UPDATE_EVEROUTE_CLUSTER"
	ROLEACTIONUPDATEEVEROUTECLUSTER ROLEACTION = "UPDATE_EVEROUTE_CLUSTER"

	// ROLEACTIONUPGRADEEVEROUTECLUSTER captures enum value "UPGRADE_EVEROUTE_CLUSTER"
	ROLEACTIONUPGRADEEVEROUTECLUSTER ROLEACTION = "UPGRADE_EVEROUTE_CLUSTER"

	// ROLEACTIONMANAGEEVEROUTENETWORKPOLICYRULESERVICE captures enum value "MANAGE_EVEROUTE_NETWORK_POLICY_RULE_SERVICE"
	ROLEACTIONMANAGEEVEROUTENETWORKPOLICYRULESERVICE ROLEACTION = "MANAGE_EVEROUTE_NETWORK_POLICY_RULE_SERVICE"

	// ROLEACTIONMANAGEEVEROUTECLUSTERASSOCIATION captures enum value "MANAGE_EVEROUTE_CLUSTER_ASSOCIATION"
	ROLEACTIONMANAGEEVEROUTECLUSTERASSOCIATION ROLEACTION = "MANAGE_EVEROUTE_CLUSTER_ASSOCIATION"

	// ROLEACTIONMANAGEEVEROUTECLUSTERGLOBALPOLICY captures enum value "MANAGE_EVEROUTE_CLUSTER_GLOBAL_POLICY"
	ROLEACTIONMANAGEEVEROUTECLUSTERGLOBALPOLICY ROLEACTION = "MANAGE_EVEROUTE_CLUSTER_GLOBAL_POLICY"

	// ROLEACTIONMANAGEMICROSEGMENTATION captures enum value "MANAGE_MICRO_SEGMENTATION"
	ROLEACTIONMANAGEMICROSEGMENTATION ROLEACTION = "MANAGE_MICRO_SEGMENTATION"

	// ROLEACTIONMANAGELOADBALANCERRESOURCE captures enum value "MANAGE_LOAD_BALANCER_RESOURCE"
	ROLEACTIONMANAGELOADBALANCERRESOURCE ROLEACTION = "MANAGE_LOAD_BALANCER_RESOURCE"

	// ROLEACTIONMANAGELOADBALANCER captures enum value "MANAGE_LOAD_BALANCER"
	ROLEACTIONMANAGELOADBALANCER ROLEACTION = "MANAGE_LOAD_BALANCER"

	// ROLEACTIONMANAGELOADBALANCERVNETBOND captures enum value "MANAGE_LOAD_BALANCER_VNET_BOND"
	ROLEACTIONMANAGELOADBALANCERVNETBOND ROLEACTION = "MANAGE_LOAD_BALANCER_VNET_BOND"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDSERVICE captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_SERVICE"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDSERVICE ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_SERVICE"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDCLUSTERBINDING captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_CLUSTER_BINDING"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDCLUSTERBINDING ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_CLUSTER_BINDING"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDEDGEGATEWAY captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_EDGE_GATEWAY"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDEDGEGATEWAY ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_EDGE_GATEWAY"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDEXTERNALSUBNET captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_EXTERNAL_SUBNET"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDEXTERNALSUBNET ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_EXTERNAL_SUBNET"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDBASICRESOURCE captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_BASIC_RESOURCE"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDBASICRESOURCE ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_BASIC_RESOURCE"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDSECURITYGROUP captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_SECURITY_GROUP"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDSECURITYGROUP ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_SECURITY_GROUP"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDSECURITYPOLICY captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_SECURITY_POLICY"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDSECURITYPOLICY ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_SECURITY_POLICY"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDISOLATIONPOLICY captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_ISOLATION_POLICY"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDISOLATIONPOLICY ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_ISOLATION_POLICY"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDFLOATINGIP captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_FLOATING_IP"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDFLOATINGIP ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_FLOATING_IP"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDROUTERGATEWAY captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_ROUTER_GATEWAY"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDROUTERGATEWAY ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_ROUTER_GATEWAY"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDNATGATEWAY captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_NAT_GATEWAY"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDNATGATEWAY ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_NAT_GATEWAY"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDLAYER2GATEWAY captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_LAYER2_GATEWAY"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDLAYER2GATEWAY ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_LAYER2_GATEWAY"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDPEERING captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_PEERING"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDPEERING ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_PEERING"

	// ROLEACTIONMANAGEVIRTUALPRIVATECLOUDLOADBALANCERRESOURCE captures enum value "MANAGE_VIRTUAL_PRIVATE_CLOUD_LOAD_BALANCER_RESOURCE"
	ROLEACTIONMANAGEVIRTUALPRIVATECLOUDLOADBALANCERRESOURCE ROLEACTION = "MANAGE_VIRTUAL_PRIVATE_CLOUD_LOAD_BALANCER_RESOURCE"

	// ROLEACTIONMANAGELDAPADCONFIG captures enum value "MANAGE_LDAP_AD_CONFIG"
	ROLEACTIONMANAGELDAPADCONFIG ROLEACTION = "MANAGE_LDAP_AD_CONFIG"

	// ROLEACTIONMANAGEMFACONFIG captures enum value "MANAGE_MFA_CONFIG"
	ROLEACTIONMANAGEMFACONFIG ROLEACTION = "MANAGE_MFA_CONFIG"

	// ROLEACTIONMANAGEDEFAULTLOGINOPTION captures enum value "MANAGE_DEFAULT_LOGIN_OPTION"
	ROLEACTIONMANAGEDEFAULTLOGINOPTION ROLEACTION = "MANAGE_DEFAULT_LOGIN_OPTION"

	// ROLEACTIONMANAGECLUSTERSTORAGEPOLICY captures enum value "MANAGE_CLUSTER_STORAGE_POLICY"
	ROLEACTIONMANAGECLUSTERSTORAGEPOLICY ROLEACTION = "MANAGE_CLUSTER_STORAGE_POLICY"

	// ROLEACTIONMANAGESKSSERVICE captures enum value "MANAGE_SKS_SERVICE"
	ROLEACTIONMANAGESKSSERVICE ROLEACTION = "MANAGE_SKS_SERVICE"

	// ROLEACTIONMANAGESKSLICENSE captures enum value "MANAGE_SKS_LICENSE"
	ROLEACTIONMANAGESKSLICENSE ROLEACTION = "MANAGE_SKS_LICENSE"

	// ROLEACTIONCONFIGURESKSSERVICE captures enum value "CONFIGURE_SKS_SERVICE"
	ROLEACTIONCONFIGURESKSSERVICE ROLEACTION = "CONFIGURE_SKS_SERVICE"

	// ROLEACTIONCREATESKSWORKLOADCLUSTER captures enum value "CREATE_SKS_WORKLOAD_CLUSTER"
	ROLEACTIONCREATESKSWORKLOADCLUSTER ROLEACTION = "CREATE_SKS_WORKLOAD_CLUSTER"

	// ROLEACTIONDELETESKSWORKLOADCLUSTER captures enum value "DELETE_SKS_WORKLOAD_CLUSTER"
	ROLEACTIONDELETESKSWORKLOADCLUSTER ROLEACTION = "DELETE_SKS_WORKLOAD_CLUSTER"

	// ROLEACTIONUPDATESKSWORKLOADCLUSTER captures enum value "UPDATE_SKS_WORKLOAD_CLUSTER"
	ROLEACTIONUPDATESKSWORKLOADCLUSTER ROLEACTION = "UPDATE_SKS_WORKLOAD_CLUSTER"

	// ROLEACTIONMANAGECONTAINERREGISTRY captures enum value "MANAGE_CONTAINER_REGISTRY"
	ROLEACTIONMANAGECONTAINERREGISTRY ROLEACTION = "MANAGE_CONTAINER_REGISTRY"

	// ROLEACTIONDOWNLOADSKSWORKLOADCLUSTERKUBECONFIG captures enum value "DOWNLOAD_SKS_WORKLOAD_CLUSTER_KUBECONFIG"
	ROLEACTIONDOWNLOADSKSWORKLOADCLUSTERKUBECONFIG ROLEACTION = "DOWNLOAD_SKS_WORKLOAD_CLUSTER_KUBECONFIG"

	// ROLEACTIONMANAGESKSWORKLOADCLUSTERRECONCILE captures enum value "MANAGE_SKS_WORKLOAD_CLUSTER_RECONCILE"
	ROLEACTIONMANAGESKSWORKLOADCLUSTERRECONCILE ROLEACTION = "MANAGE_SKS_WORKLOAD_CLUSTER_RECONCILE"

	// ROLEACTIONMANAGEOBSERVABILITYPACKAGE captures enum value "MANAGE_OBSERVABILITY_PACKAGE"
	ROLEACTIONMANAGEOBSERVABILITYPACKAGE ROLEACTION = "MANAGE_OBSERVABILITY_PACKAGE"

	// ROLEACTIONMANAGEOBSERVABILITYSERVICE captures enum value "MANAGE_OBSERVABILITY_SERVICE"
	ROLEACTIONMANAGEOBSERVABILITYSERVICE ROLEACTION = "MANAGE_OBSERVABILITY_SERVICE"

	// ROLEACTIONMANAGEALERTNOTIFIER captures enum value "MANAGE_ALERT_NOTIFIER"
	ROLEACTIONMANAGEALERTNOTIFIER ROLEACTION = "MANAGE_ALERT_NOTIFIER"

	// ROLEACTIONCONFIGDYNAMICRESOURCESCHEDULE captures enum value "CONFIG_DYNAMIC_RESOURCE_SCHEDULE"
	ROLEACTIONCONFIGDYNAMICRESOURCESCHEDULE ROLEACTION = "CONFIG_DYNAMIC_RESOURCE_SCHEDULE"

	// ROLEACTIONGENERATEDRSPROPOSALS captures enum value "GENERATE_DRS_PROPOSALS"
	ROLEACTIONGENERATEDRSPROPOSALS ROLEACTION = "GENERATE_DRS_PROPOSALS"

	// ROLEACTIONAPPLYDRSPROPOSAL captures enum value "APPLY_DRS_PROPOSAL"
	ROLEACTIONAPPLYDRSPROPOSAL ROLEACTION = "APPLY_DRS_PROPOSAL"

	// ROLEACTIONMANAGEAGENTMESH captures enum value "MANAGE_AGENT_MESH"
	ROLEACTIONMANAGEAGENTMESH ROLEACTION = "MANAGE_AGENT_MESH"

	// ROLEACTIONMANAGEREPLICATIONSERVICE captures enum value "MANAGE_REPLICATION_SERVICE"
	ROLEACTIONMANAGEREPLICATIONSERVICE ROLEACTION = "MANAGE_REPLICATION_SERVICE"

	// ROLEACTIONMANAGEREPLICATIONRESTOREPOINT captures enum value "MANAGE_REPLICATION_RESTORE_POINT"
	ROLEACTIONMANAGEREPLICATIONRESTOREPOINT ROLEACTION = "MANAGE_REPLICATION_RESTORE_POINT"

	// ROLEACTIONMANAGEREPLICATIONPLAN captures enum value "MANAGE_REPLICATION_PLAN"
	ROLEACTIONMANAGEREPLICATIONPLAN ROLEACTION = "MANAGE_REPLICATION_PLAN"

	// ROLEACTIONMANAGEREPLICATIONTASK captures enum value "MANAGE_REPLICATION_TASK"
	ROLEACTIONMANAGEREPLICATIONTASK ROLEACTION = "MANAGE_REPLICATION_TASK"

	// ROLEACTIONMANAGEREPLICATIONFAULTTASK captures enum value "MANAGE_REPLICATION_FAULT_TASK"
	ROLEACTIONMANAGEREPLICATIONFAULTTASK ROLEACTION = "MANAGE_REPLICATION_FAULT_TASK"

	// ROLEACTIONMANAGECLUSTERPRIORITIZED captures enum value "MANAGE_CLUSTER_PRIORITIZED"
	ROLEACTIONMANAGECLUSTERPRIORITIZED ROLEACTION = "MANAGE_CLUSTER_PRIORITIZED"

	// ROLEACTIONSMTXINSPECTOR captures enum value "SMTX_INSPECTOR"
	ROLEACTIONSMTXINSPECTOR ROLEACTION = "SMTX_INSPECTOR"

	// ROLEACTIONMANAGESFSLICENSE captures enum value "MANAGE_SFS_LICENSE"
	ROLEACTIONMANAGESFSLICENSE ROLEACTION = "MANAGE_SFS_LICENSE"

	// ROLEACTIONMANAGESFSIMAGE captures enum value "MANAGE_SFS_IMAGE"
	ROLEACTIONMANAGESFSIMAGE ROLEACTION = "MANAGE_SFS_IMAGE"

	// ROLEACTIONMANAGESFSFILESTORAGECLUSTER captures enum value "MANAGE_SFS_FILE_STORAGE_CLUSTER"
	ROLEACTIONMANAGESFSFILESTORAGECLUSTER ROLEACTION = "MANAGE_SFS_FILE_STORAGE_CLUSTER"

	// ROLEACTIONMANAGESFSFILESYSTEMCONFIG captures enum value "MANAGE_SFS_FILE_SYSTEM_CONFIG"
	ROLEACTIONMANAGESFSFILESYSTEMCONFIG ROLEACTION = "MANAGE_SFS_FILE_SYSTEM_CONFIG"

	// ROLEACTIONMANAGESFSFILESYSTEMACCESSIBILITY captures enum value "MANAGE_SFS_FILE_SYSTEM_ACCESSIBILITY"
	ROLEACTIONMANAGESFSFILESYSTEMACCESSIBILITY ROLEACTION = "MANAGE_SFS_FILE_SYSTEM_ACCESSIBILITY"

	// ROLEACTIONMANAGESFSSNAPSHOT captures enum value "MANAGE_SFS_SNAPSHOT"
	ROLEACTIONMANAGESFSSNAPSHOT ROLEACTION = "MANAGE_SFS_SNAPSHOT"

	// ROLEACTIONMANAGECLOUDTOWERSNMPTRANSPORT captures enum value "MANAGE_CLOUDTOWER_SNMP_TRANSPORT"
	ROLEACTIONMANAGECLOUDTOWERSNMPTRANSPORT ROLEACTION = "MANAGE_CLOUDTOWER_SNMP_TRANSPORT"

	// ROLEACTIONMANAGECLOUDTOWERNTP captures enum value "MANAGE_CLOUD_TOWER_NTP"
	ROLEACTIONMANAGECLOUDTOWERNTP ROLEACTION = "MANAGE_CLOUD_TOWER_NTP"
)

// for schema
var rOLEACTIONEnum []interface{}

func init() {
	var res []ROLEACTION
	if err := json.Unmarshal([]byte(`["*","MANAGE_DATA_CENTER","MANAGE_CLUSTER_CONNECTION","MANAGE_STORAGE_CLUSTER_CONNECTION","MANAGE_HOST","MANAGE_NIC_MTU","MANAGE_DISK","MANAGE_HARDWARE_TOPO","MANAGE_USB_DEVICE","MANAGE_GPU_DEVICE","MANAGE_VDS","MANAGE_VLAN","MANAGE_SYSTEM_VLAN","MANAGE_ISCSI_DATA_STORE","MANAGE_NFS_DATA_STORE","MANAGE_NVMF_DATA_STORE","CREATE_VM","UPDATE_VM","DELETE_VM","UPDATE_VM_ADVANCED_SETTING","UPDATE_VM_GUEST","VM_OPERATION_OPEN_TERMINAL","VM_OPERATION_MIGRATE","VM_OPERATION_VM_FOLDER","VM_OPERATION_VM_POWER","VM_OPERATION_CLONE","VM_OPERATION_INSTALL_TOOLS","VM_IMPORT_EXPORT","CREATE_VM_TEMPLATE","MANAGE_VM_TEMPLATE","VM_TEMPLATE_IMPORT_EXPORT","MANAGE_VM_SNAPSHOT","MANAGE_VM_VOLUME","VM_VOLUME_IMPORT_EXPORT","MANAGE_ISO","DOWNLOAD_ISO","QUERY_SENSITIVE_RESOURCE_LIST","QUERY_SENSITIVE_RESOURCE","MANAGE_SENSITIVE_RESOURCE","MANAGE_VM_PLACEMENT_GROUP","MANAGE_SNAPSHOT_PLAN","MANAGE_ALERT","MANAGE_MONITOR_VIEW","MANAGE_ENTITY_FILTER","MANAGE_CLUSTER_BASIC_INFO","MANAGE_CLUSTER_LICENCE","MANAGE_CLUSTER_SNMP_TRANSPORT","MANAGE_SNMP_TRAP","MANAGE_CLUSTER_VIP","MANAGE_CLUSTER_MANAGEMENT_IP","MANAGE_DNS_SERVER","MANAGE_NTP_SERVER","MANAGE_IPMI","MANAGE_CLUSTER_VM_CPU_MODEL","MANAGE_CLUSTER_VM_TOOLS","MANAGE_CLUSTER_HOT_MIGRATION","MANAGE_CLUSTER_HA","MANAGE_SSL_CERTIFICATE","MANAGE_LOG_COLLECTION","MANAGE_SYSLOG","MANAGE_LOG_FIND","MANAGE_LABEL","MANAGE_USER_AND_ROLE","MANAGE_PASSWORD_SETTINGS","MANAGE_ACCESS_CONTROL","MANAGE_SESSION_EXPIRATION","MANAGE_VCENTER_ASSOCIATION","MANAGE_ESXI_ASSOCIATION","MANAGE_AUDIT_LOG","MANAGE_ALERT_EMAIL_SETTING","MANAGE_SMTP_SERVER","MANAGE_CLUSTER_UPGRADE","MANAGE_VM_RECYCLE_BIN_SETTING","MANAGE_REPORT","MANAGE_SHARING_VM_TOOLS","MANAGE_ADVANCED_MONITOR","MANAGE_THIRD_PARTY_DRIVER","MANAGE_ORGANIZATION_NAME","MANAGE_CLOUD_TOWER_LICENSE","MANAGE_CONSISTENCY_GROUP","MANAGE_NIC","MANAGE_CLUSTER_ISCSI","MANAGE_BACKUP_LICENSE","MANAGE_BACKUP_PACKAGE","MANAGE_BACKUP_SERVICE","MANAGE_BACKUP_STORE_REPOSITORY","MANAGE_BACKUP_PLAN","MANAGE_BACKUP_TASK","MANAGE_BACKUP_RESTORE_POINT","MANAGE_BACKUP_RESTORE_POINT_TASK","MANAGE_SECURITY_POLICY","MANAGE_SECURITY_GROUP","ISOLATE_VM","MANAGE_EVEROUTE_LICENSE","MANAGE_EVEROUTE_PACKAGE","DEPLOY_EVEROUTE_CLUSTER","UNDEPLOY_EVEROUTE_CLUSTER","UPDATE_EVEROUTE_CLUSTER","UPGRADE_EVEROUTE_CLUSTER","MANAGE_EVEROUTE_NETWORK_POLICY_RULE_SERVICE","MANAGE_EVEROUTE_CLUSTER_ASSOCIATION","MANAGE_EVEROUTE_CLUSTER_GLOBAL_POLICY","MANAGE_MICRO_SEGMENTATION","MANAGE_LOAD_BALANCER_RESOURCE","MANAGE_LOAD_BALANCER","MANAGE_LOAD_BALANCER_VNET_BOND","MANAGE_VIRTUAL_PRIVATE_CLOUD_SERVICE","MANAGE_VIRTUAL_PRIVATE_CLOUD_CLUSTER_BINDING","MANAGE_VIRTUAL_PRIVATE_CLOUD_EDGE_GATEWAY","MANAGE_VIRTUAL_PRIVATE_CLOUD_EXTERNAL_SUBNET","MANAGE_VIRTUAL_PRIVATE_CLOUD_BASIC_RESOURCE","MANAGE_VIRTUAL_PRIVATE_CLOUD_SECURITY_GROUP","MANAGE_VIRTUAL_PRIVATE_CLOUD_SECURITY_POLICY","MANAGE_VIRTUAL_PRIVATE_CLOUD_ISOLATION_POLICY","MANAGE_VIRTUAL_PRIVATE_CLOUD_FLOATING_IP","MANAGE_VIRTUAL_PRIVATE_CLOUD_ROUTER_GATEWAY","MANAGE_VIRTUAL_PRIVATE_CLOUD_NAT_GATEWAY","MANAGE_VIRTUAL_PRIVATE_CLOUD_LAYER2_GATEWAY","MANAGE_VIRTUAL_PRIVATE_CLOUD_PEERING","MANAGE_VIRTUAL_PRIVATE_CLOUD_LOAD_BALANCER_RESOURCE","MANAGE_LDAP_AD_CONFIG","MANAGE_MFA_CONFIG","MANAGE_DEFAULT_LOGIN_OPTION","MANAGE_CLUSTER_STORAGE_POLICY","MANAGE_SKS_SERVICE","MANAGE_SKS_LICENSE","CONFIGURE_SKS_SERVICE","CREATE_SKS_WORKLOAD_CLUSTER","DELETE_SKS_WORKLOAD_CLUSTER","UPDATE_SKS_WORKLOAD_CLUSTER","MANAGE_CONTAINER_REGISTRY","DOWNLOAD_SKS_WORKLOAD_CLUSTER_KUBECONFIG","MANAGE_SKS_WORKLOAD_CLUSTER_RECONCILE","MANAGE_OBSERVABILITY_PACKAGE","MANAGE_OBSERVABILITY_SERVICE","MANAGE_ALERT_NOTIFIER","CONFIG_DYNAMIC_RESOURCE_SCHEDULE","GENERATE_DRS_PROPOSALS","APPLY_DRS_PROPOSAL","MANAGE_AGENT_MESH","MANAGE_REPLICATION_SERVICE","MANAGE_REPLICATION_RESTORE_POINT","MANAGE_REPLICATION_PLAN","MANAGE_REPLICATION_TASK","MANAGE_REPLICATION_FAULT_TASK","MANAGE_CLUSTER_PRIORITIZED","SMTX_INSPECTOR","MANAGE_SFS_LICENSE","MANAGE_SFS_IMAGE","MANAGE_SFS_FILE_STORAGE_CLUSTER","MANAGE_SFS_FILE_SYSTEM_CONFIG","MANAGE_SFS_FILE_SYSTEM_ACCESSIBILITY","MANAGE_SFS_SNAPSHOT","MANAGE_CLOUDTOWER_SNMP_TRANSPORT","MANAGE_CLOUD_TOWER_NTP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rOLEACTIONEnum = append(rOLEACTIONEnum, v)
	}
}

func (m ROLEACTION) validateROLEACTIONEnum(path, location string, value ROLEACTION) error {
	if err := validate.EnumCase(path, location, value, rOLEACTIONEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this r o l e a c t i o n
func (m ROLEACTION) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateROLEACTIONEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this r o l e a c t i o n based on context it is used
func (m ROLEACTION) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
