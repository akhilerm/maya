package command

import (
	"testing"

	"github.com/openebs/maya/types/v1"
)

func TestDisplayStats(t *testing.T) {
	validStats := map[string]struct {
		annotation   *Annotations
		cmdOptions   *CmdVolumeOptions
		replicaStats map[int]*ReplicaStats
		initialStats v1.VolumeMetrics
		finalStats   v1.VolumeMetrics
		output       error
		replicaCount int
	}{
		"StatsStdWhenReplicaIs0": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "0",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10",
			},
			replicaStats: nil,
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       0,
				RevisionCounter:      0,
				SectorSize:           "0",
				Size:                 "0",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13162.971420756,
				UsedBlocks:           "0",
				UsedLogicalBlocks:    "0",
				WriteIOPS:            "0",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       0,
				RevisionCounter:      0,
				SectorSize:           "0",
				Size:                 "0",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "0",
				UsedLogicalBlocks:    "0",
				WriteIOPS:            "0",
			},
			output: nil,
		},

		"StatsStdWhenReplicaIs1": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "1",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       2,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       2,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"StatsStdWhenReplicaIs2": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "2",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "Offline",
					DataUpdateIndex: "Unknown",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       2,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       2,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"StatsJSONWhenReplicaIs3": {
			cmdOptions: &CmdVolumeOptions{
				json:    "json",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "3",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,Offline",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "Offline",
					DataUpdateIndex: "Unknown",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       2,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       3,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"StatsStdWhenReplicaIs4": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "4",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,Offline,Running",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil,10.10.10.12",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "Offline",
					DataUpdateIndex: "Unknown",
				},
				3: {
					Replica:         "10.10.10.12",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       4,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       4,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"StatsStdWhenReplicaIs4AndOneErrorPullBack": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "4",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,ErrImagePull,Running",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil,10.10.10.12",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "ErrImagePull",
					DataUpdateIndex: "Unknown",
				},
				3: {
					Replica:         "10.10.10.12",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       4,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       4,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"StatsStdWhenReplicaIs4AndOneCrashLoopBackOff": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "4",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,CrashLoopBackOff,Running",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil,10.10.10.12",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "CrashLoopBackOff",
					DataUpdateIndex: "Unknown",
				},
				3: {
					Replica:         "10.10.10.12",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       4,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       4,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "0",
				TotalReadTime:        "0",
				TotalWriteTime:       "0",
				TotalWriteBlockCount: "0",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"StatsStd": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "3",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,Offline",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "Offline",
					DataUpdateIndex: "Unknown",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       6,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       6,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "4",
				TotalReadTime:        "12",
				TotalWriteTime:       "16",
				TotalWriteBlockCount: "15",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"ReadIOPSIsNotZero": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "3",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,Offline",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "Offline",
					DataUpdateIndex: "Unknown",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       6,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "10",
				ReplicaCounter:       6,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "4",
				TotalReadTime:        "12",
				TotalWriteTime:       "16",
				TotalWriteBlockCount: "15",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "20",
			},
			output: nil,
		},
		"WriteIOPSIsZero": {
			cmdOptions: &CmdVolumeOptions{
				json:    "",
				volName: "vol1",
			},
			annotation: &Annotations{
				TargetPortal:     "10.99.73.74:3260",
				ClusterIP:        "10.99.73.74",
				Iqn:              "iqn.2016-09.com.openebs.jiva:vol1",
				ReplicaCount:     "3",
				ControllerStatus: "Running",
				ReplicaStatus:    "Running,Running,Offline",
				VolSize:          "1G",
				ControllerIP:     "",
				Replicas:         "10.10.10.10,10.10.10.11,nil",
			},
			replicaStats: map[int]*ReplicaStats{
				0: {
					Replica:         "10.10.10.10",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				1: {
					Replica:         "10.10.10.11",
					Status:          "Online",
					DataUpdateIndex: "1",
				},
				2: {
					Replica:         "nil",
					Status:          "Offline",
					DataUpdateIndex: "Unknown",
				},
			},
			initialStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "0",
				ReplicaCounter:       6,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "3",
				TotalReadTime:        "10",
				TotalWriteTime:       "15",
				TotalWriteBlockCount: "10",
				UpTime:               13162.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			finalStats: v1.VolumeMetrics{
				Name:                 "vol1",
				ReadIOPS:             "10",
				ReplicaCounter:       6,
				RevisionCounter:      100,
				SectorSize:           "4096",
				Size:                 "1073741824",
				TotalReadBlockCount:  "4",
				TotalReadTime:        "12",
				TotalWriteTime:       "16",
				TotalWriteBlockCount: "15",
				UpTime:               13170.971420756,
				UsedBlocks:           "1048576",
				UsedLogicalBlocks:    "1048576",
				WriteIOPS:            "15",
			},
			output: nil,
		},
	}
	for name, tt := range validStats {
		t.Run(name, func(t *testing.T) {
			if got := tt.annotation.DisplayStats(tt.cmdOptions, tt.replicaStats, tt.initialStats, tt.finalStats); got != tt.output {
				t.Fatalf("DisplayStats(%v, %v, %v, %v) => %v, want %v", tt.cmdOptions, tt.replicaStats, tt.initialStats, tt.finalStats, got, tt.output)
			}
		})
	}

}