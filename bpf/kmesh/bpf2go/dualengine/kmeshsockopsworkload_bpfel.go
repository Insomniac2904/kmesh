// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package dualengine

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshSockopsWorkloadBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshSockopsWorkloadBuf struct{ Data [40]int8 }

type KmeshSockopsWorkloadKmeshConfig struct {
	BpfLogLevel uint32
	NodeIp      [4]uint32
	PodGateway  [4]uint32
}

type KmeshSockopsWorkloadLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

type KmeshSockopsWorkloadManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshSockopsWorkloadOperationUsageData struct {
	StartTime     uint64
	EndTime       uint64
	PidTgid       uint64
	OperationType uint32
	_             [4]byte
}

type KmeshSockopsWorkloadOperationUsageKey struct {
	SocketCookie  uint64
	OperationType uint32
	_             [4]byte
}

type KmeshSockopsWorkloadSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshSockopsWorkload returns the embedded CollectionSpec for KmeshSockopsWorkload.
func LoadKmeshSockopsWorkload() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshSockopsWorkloadBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshSockopsWorkload: %w", err)
	}

	return spec, err
}

// LoadKmeshSockopsWorkloadObjects loads KmeshSockopsWorkload and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshSockopsWorkloadObjects
//	*KmeshSockopsWorkloadPrograms
//	*KmeshSockopsWorkloadMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshSockopsWorkloadObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshSockopsWorkload()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshSockopsWorkloadSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSockopsWorkloadSpecs struct {
	KmeshSockopsWorkloadProgramSpecs
	KmeshSockopsWorkloadMapSpecs
}

// KmeshSockopsWorkloadSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSockopsWorkloadProgramSpecs struct {
	SockopsProg *ebpf.ProgramSpec `ebpf:"sockops_prog"`
}

// KmeshSockopsWorkloadMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSockopsWorkloadMapSpecs struct {
	KmeshBackend     *ebpf.MapSpec `ebpf:"kmesh_backend"`
	KmeshConfigMap   *ebpf.MapSpec `ebpf:"kmesh_config_map"`
	KmeshEndpoint    *ebpf.MapSpec `ebpf:"kmesh_endpoint"`
	KmeshEvents      *ebpf.MapSpec `ebpf:"kmesh_events"`
	KmeshFrontend    *ebpf.MapSpec `ebpf:"kmesh_frontend"`
	KmeshManage      *ebpf.MapSpec `ebpf:"kmesh_manage"`
	KmeshMap1600     *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192      *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296      *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64       *ebpf.MapSpec `ebpf:"kmesh_map64"`
	KmeshPerfInfo    *ebpf.MapSpec `ebpf:"kmesh_perf_info"`
	KmeshPerfMap     *ebpf.MapSpec `ebpf:"kmesh_perf_map"`
	KmeshService     *ebpf.MapSpec `ebpf:"kmesh_service"`
	MapOfAuth        *ebpf.MapSpec `ebpf:"map_of_auth"`
	MapOfDstInfo     *ebpf.MapSpec `ebpf:"map_of_dst_info"`
	MapOfKmeshSocket *ebpf.MapSpec `ebpf:"map_of_kmesh_socket"`
	MapOfSockStorage *ebpf.MapSpec `ebpf:"map_of_sock_storage"`
	MapOfTcpInfo     *ebpf.MapSpec `ebpf:"map_of_tcp_info"`
	MapOfTuple       *ebpf.MapSpec `ebpf:"map_of_tuple"`
	MapOfWlPolicy    *ebpf.MapSpec `ebpf:"map_of_wl_policy"`
	TmpBuf           *ebpf.MapSpec `ebpf:"tmp_buf"`
	TmpLogBuf        *ebpf.MapSpec `ebpf:"tmp_log_buf"`
}

// KmeshSockopsWorkloadObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSockopsWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSockopsWorkloadObjects struct {
	KmeshSockopsWorkloadPrograms
	KmeshSockopsWorkloadMaps
}

func (o *KmeshSockopsWorkloadObjects) Close() error {
	return _KmeshSockopsWorkloadClose(
		&o.KmeshSockopsWorkloadPrograms,
		&o.KmeshSockopsWorkloadMaps,
	)
}

// KmeshSockopsWorkloadMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSockopsWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSockopsWorkloadMaps struct {
	KmeshBackend     *ebpf.Map `ebpf:"kmesh_backend"`
	KmeshConfigMap   *ebpf.Map `ebpf:"kmesh_config_map"`
	KmeshEndpoint    *ebpf.Map `ebpf:"kmesh_endpoint"`
	KmeshEvents      *ebpf.Map `ebpf:"kmesh_events"`
	KmeshFrontend    *ebpf.Map `ebpf:"kmesh_frontend"`
	KmeshManage      *ebpf.Map `ebpf:"kmesh_manage"`
	KmeshMap1600     *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192      *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296      *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64       *ebpf.Map `ebpf:"kmesh_map64"`
	KmeshPerfInfo    *ebpf.Map `ebpf:"kmesh_perf_info"`
	KmeshPerfMap     *ebpf.Map `ebpf:"kmesh_perf_map"`
	KmeshService     *ebpf.Map `ebpf:"kmesh_service"`
	MapOfAuth        *ebpf.Map `ebpf:"map_of_auth"`
	MapOfDstInfo     *ebpf.Map `ebpf:"map_of_dst_info"`
	MapOfKmeshSocket *ebpf.Map `ebpf:"map_of_kmesh_socket"`
	MapOfSockStorage *ebpf.Map `ebpf:"map_of_sock_storage"`
	MapOfTcpInfo     *ebpf.Map `ebpf:"map_of_tcp_info"`
	MapOfTuple       *ebpf.Map `ebpf:"map_of_tuple"`
	MapOfWlPolicy    *ebpf.Map `ebpf:"map_of_wl_policy"`
	TmpBuf           *ebpf.Map `ebpf:"tmp_buf"`
	TmpLogBuf        *ebpf.Map `ebpf:"tmp_log_buf"`
}

func (m *KmeshSockopsWorkloadMaps) Close() error {
	return _KmeshSockopsWorkloadClose(
		m.KmeshBackend,
		m.KmeshConfigMap,
		m.KmeshEndpoint,
		m.KmeshEvents,
		m.KmeshFrontend,
		m.KmeshManage,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
		m.KmeshPerfInfo,
		m.KmeshPerfMap,
		m.KmeshService,
		m.MapOfAuth,
		m.MapOfDstInfo,
		m.MapOfKmeshSocket,
		m.MapOfSockStorage,
		m.MapOfTcpInfo,
		m.MapOfTuple,
		m.MapOfWlPolicy,
		m.TmpBuf,
		m.TmpLogBuf,
	)
}

// KmeshSockopsWorkloadPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSockopsWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSockopsWorkloadPrograms struct {
	SockopsProg *ebpf.Program `ebpf:"sockops_prog"`
}

func (p *KmeshSockopsWorkloadPrograms) Close() error {
	return _KmeshSockopsWorkloadClose(
		p.SockopsProg,
	)
}

func _KmeshSockopsWorkloadClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshsockopsworkload_bpfel.o
var _KmeshSockopsWorkloadBytes []byte
