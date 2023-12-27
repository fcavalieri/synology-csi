// Copyright 2021 Synology Inc.

package models

import (
	"fmt"
	"github.com/SynologyOpenSource/synology-csi/pkg/utils"
)

const (
	K8sCsiName = "Kubernetes CSI"

	// ISCSI definitions
	FsTypeExt4       = "ext4"
	FsTypeBtrfs      = "btrfs"
	LunTypeFile      = "FILE"
	LunTypeThin      = "THIN"
	LunTypeAdv       = "ADV"
	LunTypeBlun      = "BLUN"       // thin provision, mapped to type 263
	LunTypeBlunThick = "BLUN_THICK" // thick provision, mapped to type 259
	MaxIqnLen        = 128

	// Share definitions
	MaxShareLen             = 32
	MaxShareDescLen         = 64
	UserGroupTypeLocalUser  = "local_user"
	UserGroupTypeLocalGroup = "local_group"
	UserGroupTypeSystem     = "system"

	// CSI definitions
	TargetPrefix            = "k8s-csi"
	LunPrefix               = "k8s-csi"
	IqnPrefix               = "iqn.2000-01.com.synology:"
	SharePrefixISCSI        = "k8s-csi-iscsi"
	SharePrefixSMB          = "k8s-csi-smb"
	SharePrefixNFS          = "k8s-csi-nfs"
	ShareSnapshotDescPrefix = "(Do not change)"
)

func GenLunName(volName string) string {
	return fmt.Sprintf("%s-%s", LunPrefix, volName)
}

func GenTargetName(volName string) string {
	return fmt.Sprintf("%s-%s", TargetPrefix, volName)
}

func GenShareName(volName string, protocol string) string {
	shareName := ""
	if protocol == utils.ProtocolIscsi {
		shareName = fmt.Sprintf("%s-%s", SharePrefixISCSI, volName)
	} else if protocol == utils.ProtocolSmb {
		shareName = fmt.Sprintf("%s-%s", SharePrefixSMB, volName)
	} else if protocol == utils.ProtocolNfs {
		shareName = fmt.Sprintf("%s-%s", SharePrefixNFS, volName)
	} else {
		return ""
	}

	if len(shareName) > MaxShareLen {
		return shareName[:MaxShareLen]
	}
	return shareName
}
