// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package general

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshTcMarkDecryptBuf struct{ Data [40]int8 }

type KmeshTcMarkDecryptLpmKey struct {
	TrieKey struct {
		Prefixlen uint32
		Data      [0]uint8
	}
	Ip struct {
		Ip4 uint32
		_   [12]byte
	}
}

// LoadKmeshTcMarkDecrypt returns the embedded CollectionSpec for KmeshTcMarkDecrypt.
func LoadKmeshTcMarkDecrypt() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshTcMarkDecryptBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshTcMarkDecrypt: %w", err)
	}

	return spec, err
}

// LoadKmeshTcMarkDecryptObjects loads KmeshTcMarkDecrypt and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshTcMarkDecryptObjects
//	*KmeshTcMarkDecryptPrograms
//	*KmeshTcMarkDecryptMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshTcMarkDecryptObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshTcMarkDecrypt()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshTcMarkDecryptSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTcMarkDecryptSpecs struct {
	KmeshTcMarkDecryptProgramSpecs
	KmeshTcMarkDecryptMapSpecs
	KmeshTcMarkDecryptVariableSpecs
}

// KmeshTcMarkDecryptProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTcMarkDecryptProgramSpecs struct {
	TcMarkDecrypt *ebpf.ProgramSpec `ebpf:"tc_mark_decrypt"`
}

// KmeshTcMarkDecryptMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTcMarkDecryptMapSpecs struct {
	KmLogEvent *ebpf.MapSpec `ebpf:"km_log_event"`
	KmNodeinfo *ebpf.MapSpec `ebpf:"km_nodeinfo"`
	KmTmpbuf   *ebpf.MapSpec `ebpf:"km_tmpbuf"`
}

// KmeshTcMarkDecryptVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTcMarkDecryptVariableSpecs struct {
	BpfLogLevel *ebpf.VariableSpec `ebpf:"bpf_log_level"`
}

// KmeshTcMarkDecryptObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTcMarkDecryptObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTcMarkDecryptObjects struct {
	KmeshTcMarkDecryptPrograms
	KmeshTcMarkDecryptMaps
	KmeshTcMarkDecryptVariables
}

func (o *KmeshTcMarkDecryptObjects) Close() error {
	return _KmeshTcMarkDecryptClose(
		&o.KmeshTcMarkDecryptPrograms,
		&o.KmeshTcMarkDecryptMaps,
	)
}

// KmeshTcMarkDecryptMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTcMarkDecryptObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTcMarkDecryptMaps struct {
	KmLogEvent *ebpf.Map `ebpf:"km_log_event"`
	KmNodeinfo *ebpf.Map `ebpf:"km_nodeinfo"`
	KmTmpbuf   *ebpf.Map `ebpf:"km_tmpbuf"`
}

func (m *KmeshTcMarkDecryptMaps) Close() error {
	return _KmeshTcMarkDecryptClose(
		m.KmLogEvent,
		m.KmNodeinfo,
		m.KmTmpbuf,
	)
}

// KmeshTcMarkDecryptVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTcMarkDecryptObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTcMarkDecryptVariables struct {
	BpfLogLevel *ebpf.Variable `ebpf:"bpf_log_level"`
}

// KmeshTcMarkDecryptPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTcMarkDecryptObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTcMarkDecryptPrograms struct {
	TcMarkDecrypt *ebpf.Program `ebpf:"tc_mark_decrypt"`
}

func (p *KmeshTcMarkDecryptPrograms) Close() error {
	return _KmeshTcMarkDecryptClose(
		p.TcMarkDecrypt,
	)
}

func _KmeshTcMarkDecryptClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshtcmarkdecrypt_bpfel.o
var _KmeshTcMarkDecryptBytes []byte
