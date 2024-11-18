// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package normal

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshCgroupSockBuf struct{ Data [40]int8 }

type KmeshCgroupSockClusterSockData struct{ ClusterId uint32 }

type KmeshCgroupSockKmeshConfig struct {
	BpfLogLevel uint32
	NodeIp      [4]uint32
	PodGateway  [4]uint32
}

type KmeshCgroupSockLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

type KmeshCgroupSockManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockRatelimitKey struct {
	Key struct {
		SkSkb struct {
			Netns  uint64
			Ipv4   uint32
			Port   uint32
			Family uint32
			_      [4]byte
		}
	}
}

type KmeshCgroupSockRatelimitValue struct {
	LastTopup uint64
	Tokens    uint64
}

type KmeshCgroupSockSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSock returns the embedded CollectionSpec for KmeshCgroupSock.
func LoadKmeshCgroupSock() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSock: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockObjects loads KmeshCgroupSock and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockObjects
//	*KmeshCgroupSockPrograms
//	*KmeshCgroupSockMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSock()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockSpecs struct {
	KmeshCgroupSockProgramSpecs
	KmeshCgroupSockMapSpecs
}

// KmeshCgroupSockSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.ProgramSpec `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.ProgramSpec `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.ProgramSpec `ebpf:"filter_manager"`
}

// KmeshCgroupSockMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockMapSpecs struct {
	KmeshCluster        *ebpf.MapSpec `ebpf:"kmesh_cluster"`
	KmeshClusterStats   *ebpf.MapSpec `ebpf:"kmesh_cluster_stats"`
	KmeshConfigMap      *ebpf.MapSpec `ebpf:"kmesh_config_map"`
	KmeshEvents         *ebpf.MapSpec `ebpf:"kmesh_events"`
	KmeshListener       *ebpf.MapSpec `ebpf:"kmesh_listener"`
	KmeshManage         *ebpf.MapSpec `ebpf:"kmesh_manage"`
	KmeshMap1600        *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192         *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296         *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64          *ebpf.MapSpec `ebpf:"kmesh_map64"`
	KmeshRatelimit      *ebpf.MapSpec `ebpf:"kmesh_ratelimit"`
	KmeshTailCallCtx    *ebpf.MapSpec `ebpf:"kmesh_tail_call_ctx"`
	KmeshTailCallProg   *ebpf.MapSpec `ebpf:"kmesh_tail_call_prog"`
	MapOfClusterEps     *ebpf.MapSpec `ebpf:"map_of_cluster_eps"`
	MapOfClusterEpsData *ebpf.MapSpec `ebpf:"map_of_cluster_eps_data"`
	MapOfClusterSock    *ebpf.MapSpec `ebpf:"map_of_cluster_sock"`
	MapOfSockStorage    *ebpf.MapSpec `ebpf:"map_of_sock_storage"`
	OuterOfMaglev       *ebpf.MapSpec `ebpf:"outer_of_maglev"`
	TmpBuf              *ebpf.MapSpec `ebpf:"tmp_buf"`
	TmpLogBuf           *ebpf.MapSpec `ebpf:"tmp_log_buf"`
}

// KmeshCgroupSockObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockObjects struct {
	KmeshCgroupSockPrograms
	KmeshCgroupSockMaps
}

func (o *KmeshCgroupSockObjects) Close() error {
	return _KmeshCgroupSockClose(
		&o.KmeshCgroupSockPrograms,
		&o.KmeshCgroupSockMaps,
	)
}

// KmeshCgroupSockMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockMaps struct {
	KmeshCluster        *ebpf.Map `ebpf:"kmesh_cluster"`
	KmeshClusterStats   *ebpf.Map `ebpf:"kmesh_cluster_stats"`
	KmeshConfigMap      *ebpf.Map `ebpf:"kmesh_config_map"`
	KmeshEvents         *ebpf.Map `ebpf:"kmesh_events"`
	KmeshListener       *ebpf.Map `ebpf:"kmesh_listener"`
	KmeshManage         *ebpf.Map `ebpf:"kmesh_manage"`
	KmeshMap1600        *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192         *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296         *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64          *ebpf.Map `ebpf:"kmesh_map64"`
	KmeshRatelimit      *ebpf.Map `ebpf:"kmesh_ratelimit"`
	KmeshTailCallCtx    *ebpf.Map `ebpf:"kmesh_tail_call_ctx"`
	KmeshTailCallProg   *ebpf.Map `ebpf:"kmesh_tail_call_prog"`
	MapOfClusterEps     *ebpf.Map `ebpf:"map_of_cluster_eps"`
	MapOfClusterEpsData *ebpf.Map `ebpf:"map_of_cluster_eps_data"`
	MapOfClusterSock    *ebpf.Map `ebpf:"map_of_cluster_sock"`
	MapOfSockStorage    *ebpf.Map `ebpf:"map_of_sock_storage"`
	OuterOfMaglev       *ebpf.Map `ebpf:"outer_of_maglev"`
	TmpBuf              *ebpf.Map `ebpf:"tmp_buf"`
	TmpLogBuf           *ebpf.Map `ebpf:"tmp_log_buf"`
}

func (m *KmeshCgroupSockMaps) Close() error {
	return _KmeshCgroupSockClose(
		m.KmeshCluster,
		m.KmeshClusterStats,
		m.KmeshConfigMap,
		m.KmeshEvents,
		m.KmeshListener,
		m.KmeshManage,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
		m.KmeshRatelimit,
		m.KmeshTailCallCtx,
		m.KmeshTailCallProg,
		m.MapOfClusterEps,
		m.MapOfClusterEpsData,
		m.MapOfClusterSock,
		m.MapOfSockStorage,
		m.OuterOfMaglev,
		m.TmpBuf,
		m.TmpLogBuf,
	)
}

// KmeshCgroupSockPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.Program `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.Program `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.Program `ebpf:"filter_manager"`
}

func (p *KmeshCgroupSockPrograms) Close() error {
	return _KmeshCgroupSockClose(
		p.CgroupConnect4Prog,
		p.ClusterManager,
		p.FilterChainManager,
		p.FilterManager,
	)
}

func _KmeshCgroupSockClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsock_bpfeb.o
var _KmeshCgroupSockBytes []byte
