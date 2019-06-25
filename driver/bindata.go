// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (5.595kB)
// override/templates/singleton/crdb_upsert.go.tpl (1.48kB)
// override/templates_test/singleton/crdb_main_test.go.tpl (3.929kB)
// override/templates_test/singleton/crdb_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.745kB)

package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5d\x73\xdb\xb6\x12\x7d\x26\x7f\xc5\xc6\x73\x27\x26\xef\xc8\xf4\x7d\xf6\x1d\x3d\xf8\x23\x49\x33\x69\x1c\x35\x8e\x9b\x99\x66\x32\x1e\x88\x5c\x4a\x18\x43\x00\x03\x82\x76\x54\x96\xff\xbd\xb3\x00\x28\x92\xfa\xb0\x95\x34\x69\xd3\x27\x8b\xc0\x62\xf7\xe0\xec\x01\x76\xe1\xba\x3e\x82\xff\x30\xc1\x59\x09\x27\x63\x48\x4e\xe9\x17\x96\xc9\x3b\x36\x15\x08\xee\x4f\x72\xc9\x16\xd8\x34\xa1\x35\x2d\xd3\x39\x2e\x98\x9b\xa6\x05\x9d\x05\xfc\x01\xc9\x55\x37\x6b\x17\xf0\x1c\x92\xd3\x2c\x7b\x21\xd4\x94\x09\x38\x6a\x9a\xf0\xf8\x18\xae\x8b\x12\xb5\x79\x01\xcc\x18\x5c\x14\xa6\x04\x26\x81\x4b\x1a\x1b\x01\x93\x19\x64\x0a\xed\x58\x55\x64\xcc\x20\x28\x0d\x7c\x26\x95\x46\x50\x12\x52\x25\x73\xc1\x53\x93\x84\x79\x25\x53\x88\x14\xfc\xb7\xae\x1d\xfe\xe4\xba\xb8\xe2\x72\x56\x09\xa6\x9b\x26\x6e\xa3\x44\x16\x84\x54\x06\x92\x4b\x75\xae\xa4\xc1\xcf\xa6\x69\x52\xf3\x99\x5c\xd1\x47\xe2\x07\x47\x50\xd7\x28\x33\x02\xe9\x23\xbf\x91\xe7\x3e\x1a\x4c\x95\x12\xa3\x55\xf0\x73\x25\xaa\x85\x2c\xe1\xc3\xc7\xd2\x68\x2e\x67\x23\xbf\xc0\x8f\x8f\xfc\x6e\x5a\xb3\xa9\xe2\x22\xf1\x1f\x31\xa0\xd6\x4a\x43\x1d\x06\x1a\x4d\xa5\x25\xa8\xc4\x21\x75\x40\xfb\x20\xed\xba\x17\x68\x2e\xce\xa2\xb8\xae\x51\x94\x68\x81\x8f\xa0\x9d\xf0\x96\x7e\x5e\x66\x4d\x33\xda\x80\xbe\x81\xfa\x61\xb0\x71\xd8\x84\xe1\x8a\x88\xd0\xa5\x90\x92\xd2\x4b\x23\xfd\x9c\x30\xc9\xd3\xb5\x84\x4e\xfe\x5a\x46\xc1\xfa\x2c\x69\xcc\x72\xb4\x77\x8a\x27\x3f\x5c\x8e\xeb\x30\xe0\x39\xed\x82\x8e\xc8\x0f\x96\xe0\xff\x5b\x5c\x4f\xc6\x20\xb9\x20\xa0\x41\x41\xb4\x47\x36\xe4\x7b\xcd\x8a\x67\x5a\x47\xa8\x75\x1c\x87\x41\xb3\x4d\x0c\x3b\xb2\xbf\x2d\xf9\x50\x95\x5c\xce\xe8\x1b\x3f\x63\x5a\x19\xa5\xbf\xe4\x80\xf7\x5c\x17\x5f\xa7\x8c\xc9\x26\xe5\x04\xc4\xd1\xfb\xcc\x43\xea\x11\xbf\x29\x97\xce\xdc\x0f\xf5\x56\x6d\x4f\xc7\xdf\x24\xa3\x2d\x62\xef\x8b\x9b\x70\xff\xa3\x52\x59\x25\xef\x7b\xc8\xe2\x0a\x71\xc0\x14\x64\x2a\xad\x16\x28\x0d\x33\x5c\x49\xc8\x95\x86\xb9\xba\x07\xa3\xa0\xd0\xaa\x40\x2d\x96\x50\x95\x38\xdc\xab\x8d\x38\xd8\xee\xbe\xaa\xfa\x97\x8b\x6a\x55\x7f\x78\x0e\x0a\xc6\x5d\x72\x7d\x3d\xb2\xf3\x65\x72\x89\xf7\xd1\x41\x5d\x27\x93\xdb\x99\x2b\xff\x27\x20\x15\xd4\xf5\xa0\x25\x20\x7e\xef\x78\x86\x99\xe5\xbc\xb2\xf4\x1c\x58\x35\x84\x01\x75\x0b\x94\x79\x41\xb9\x3c\x30\x7c\x81\xa5\x61\x8b\xe2\xc6\x59\xdd\xcc\x51\x14\xa8\x0f\x20\x81\xc6\x59\x77\xa2\xfe\x49\xa9\xdb\xd2\xca\x68\x20\xff\x4c\x9d\x61\xae\x34\xba\x2c\x58\xa3\xbd\xcf\xc2\xa6\x94\xbb\xdd\x12\x5c\x8b\xd6\x92\x1f\x86\x81\xfc\xfd\x02\x73\x56\x09\x63\x5b\xa2\x4f\x15\x6a\x8e\x65\x72\xa9\xe4\x6f\xa8\x95\x9f\xba\x42\xd2\x81\x57\xc9\x85\xba\x97\x9d\x4e\x3c\xd3\xef\xb9\x99\x7b\xe3\x11\xa8\x38\x0c\x83\xe3\x63\x38\xab\xb8\xc8\x20\x65\xe9\x1c\xe1\x16\x97\xc0\xe5\x91\xe0\x12\xa1\x9a\x09\x2e\x96\x70\x04\x8b\x65\xf9\x49\xc0\x5d\x09\x05\xfd\x2d\xb4\x9a\x0a\x5c\x94\x61\x30\xad\x72\x02\x53\x1a\xbd\x60\x72\x26\x90\xaa\xc3\x59\x95\xe7\xa8\xa3\xd8\xd2\xb4\x21\x19\xda\xe4\xb4\xca\x93\xf7\x9a\x1b\x3c\x5b\x1a\x8c\x0e\xcd\x21\xe5\x06\x48\x9a\xdb\xa6\x73\x3b\x1d\xae\x0f\x27\x34\x4c\xf9\xbd\x19\x41\x4a\x20\x34\x93\x33\xdc\x10\xe3\xc0\xe1\x95\xd5\x65\x94\xee\x76\xb8\x6e\x5a\x1a\x9d\x2a\x79\x97\xbc\x34\x8a\x45\x03\x39\x27\xaf\xb8\xcc\xe2\xad\x18\x86\x76\xe7\x4a\x7c\x5b\x18\xc3\xeb\x61\x37\x8c\xa1\xdd\xd7\xc0\xd8\xf4\xd9\x13\xe1\x03\xbe\x48\x43\x27\x63\xa0\x59\x3f\x11\x87\x41\x27\x92\x49\xd5\x8a\x64\x5a\xe5\xb1\x3d\x66\x5b\x25\xeb\x8e\xd4\x39\xc9\xf2\x75\x65\x92\xb7\x3f\xab\xf4\x96\x3c\x59\xa1\x8e\x9c\x5e\x33\x0a\xf4\xf8\xfa\x0f\xb7\xb8\xfc\xb8\x77\xa0\x6b\x29\x5c\xa8\x30\xb8\x63\xda\x9e\x51\x7b\xff\x84\x56\xd3\x4f\x7c\x60\x22\xa0\x6d\x27\x35\x1a\x02\x32\xa4\xfc\x65\xef\x8b\x4e\x66\x18\x04\xbb\x10\x9c\x0a\xd1\x5e\x93\x0f\x58\x6d\x39\xc3\xfb\x59\xab\xca\xf4\x17\x74\x59\xa4\xcf\x38\x0c\x02\x5f\xdc\x4e\xc6\x6b\xe2\xbd\xee\x7d\x7d\x93\x2d\x4c\x34\x5f\x30\xbd\x7c\x85\xcb\x9e\x31\x11\xbd\xf5\xb6\x78\xfa\x14\x04\x4a\x7f\xf0\x62\x2a\x0b\xff\xb3\xb4\x3f\x5e\x15\x2a\x69\xdf\x82\x46\xf9\xfb\x7f\xbd\x46\x50\xd9\xaa\x44\x66\x6f\xe9\xa9\xbd\xfe\x3c\x05\xa9\x85\x05\x82\x97\xb6\x66\xd8\xa2\x11\xb4\xb7\x0a\x11\xb4\x76\xc3\x38\xe4\x84\xb2\x9d\xe8\xe3\x5c\x2d\x1c\xc3\x82\xdd\x62\xd4\xd5\x46\x5a\xb1\x2f\x47\x74\xbe\xc9\x57\xb1\x5c\x05\x19\xed\x12\xfd\xe6\x62\xbb\x89\xc0\x9d\x9a\x84\xea\xc6\x12\xc6\x6e\xcf\x4e\xf7\xbf\xd0\xd0\xb9\x4a\x6f\xb5\x62\xe9\xfc\xe2\x2c\xca\x38\x13\x48\x21\x0e\xea\xba\xff\xb2\x6e\x9a\x83\x6d\xdd\x9b\x46\xd3\x0e\x77\xcd\x40\x5b\xed\x6d\x6a\x5d\xe8\x3b\x26\x2a\x7c\xcd\x8a\xc2\xee\x9f\x0e\x55\x57\xc6\xce\xb8\xcc\xfc\xd4\x2e\x56\xde\x2d\x0b\xdc\xb9\xeb\x95\xdb\x36\x6a\xd0\x16\xe9\x5e\x71\x1d\x54\x57\xcb\x89\xcf\x9c\x46\x13\x93\x61\x9b\x34\x0b\x57\xa3\xf9\xde\x60\x29\x2e\x05\xdc\x02\x75\x88\xd5\x82\x6d\x5c\x07\x63\x69\xb4\x37\x32\xe6\x94\xa6\xe4\xa5\xcc\xb8\xc6\xd4\x44\xed\xc0\xaf\x64\xf1\x26\x8f\x14\xe9\xe6\x8e\x89\x41\xc3\x60\x27\xcb\xe7\x5a\x2d\xda\x2d\x58\x87\xfe\x3a\x1d\x24\x29\x76\xd7\x9f\x43\x42\x7d\x1d\x97\x06\x75\xce\x52\xac\x5d\x13\x64\x55\xbf\x46\x56\x8f\xc8\x76\x61\x17\x7c\x62\xf4\xee\xd0\x3d\x1f\x6e\xa7\x3c\x77\x4d\xe2\x05\x4e\xab\xd9\x6b\x95\xb9\xf6\x20\x5f\x98\xe4\x79\xa1\xb9\x34\x42\x46\xdd\xbc\x2d\x43\xba\xf5\x65\x65\x1e\x3f\x6e\x4d\xec\x74\xd1\x1e\xd9\xcf\x5a\x87\xed\x7a\xc1\xc0\x69\x83\xda\xb9\xc4\x9e\xa4\xb7\xea\x3e\xea\x81\x70\x31\x92\x24\x89\x93\xab\x94\x59\xad\x11\x29\x34\x60\x5d\xda\xb6\x67\xa7\x27\x1f\x2a\xb2\xcd\xe3\x97\x78\xf5\x2f\x9e\x95\xb6\xc6\x63\x28\x3f\x89\xe4\x99\xd6\x97\xea\xad\xba\x77\xe5\xdb\x47\x24\xd1\x1d\x1f\x03\x5d\x01\xa9\xd2\xf6\x0e\xb0\xaf\x1e\x79\x68\x7c\xf2\x81\xc9\xa5\x99\xd3\xf3\xe8\x7e\x8e\x12\xcc\x1c\x35\x1e\x96\xd4\x7a\xbb\x73\xef\xd5\xd9\xf5\x70\xdb\xa9\xba\x69\x4f\x92\xdd\x23\xbd\x2f\xb6\x33\xb5\x4e\xcc\xe6\xba\xc7\x79\x19\xd2\xd0\x35\xed\x5b\x9b\x6d\x2a\x22\xf4\x74\xa4\x77\xa3\xbd\xf6\xbe\xa4\x94\x1c\x74\x02\xea\xb7\x06\xfb\xf5\x1a\x6d\x4f\xb3\x87\xb9\xed\x61\x60\xec\xb6\xbb\x77\x80\x55\x2f\x13\x3c\xf0\xa0\x59\xfd\xd3\x2f\x53\xa7\xb9\x41\xfd\x55\x8f\x19\xff\x5c\x59\xa5\xcd\x3b\x95\x5c\xf4\x1f\x32\xcd\x9f\x01\x00\x00\xff\xff\x88\xe9\x24\x90\xdb\x15\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x31, 0x37, 0xef, 0xc5, 0x39, 0xc8, 0xb4, 0xd5, 0x1d, 0xb3, 0x81, 0x92, 0xa7, 0xe, 0xc9, 0xc6, 0x9e, 0xcb, 0x97, 0xc7, 0xa6, 0xc4, 0x64, 0x30, 0xfe, 0xa9, 0xcb, 0xc6, 0x68, 0xd2, 0x11, 0x94}}
	return a, nil
}

var _templatesSingletonCrdb_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5d\x6f\xda\x4a\x10\x7d\xde\xfd\x15\x13\x4b\x51\x6c\x69\x65\x6e\x5e\xaf\xc4\x43\x82\x9d\x5c\xae\x90\x09\x60\xda\x4a\x55\x1f\x16\x7b\x0c\xab\x9a\x35\xdd\x0f\x92\xa8\xe1\xbf\x57\xeb\x0f\x4c\x42\xd4\x28\x12\x32\x68\x67\xce\xd9\x33\xc7\x67\x18\x0c\x60\x65\x45\x99\x2f\x77\x1a\x95\x99\x59\x54\xcf\xa3\x2a\xfb\xa9\x2a\x9e\x6d\xa2\xdb\xa6\xa6\x81\xc3\x62\x36\x01\x6d\xb8\xc1\x2d\x4a\x03\xda\x28\x21\xd7\x60\xb5\x7b\x9a\x0d\x82\xad\xe1\x11\x37\x1c\x76\xaa\xda\x8b\x1c\xf3\x90\x16\x56\x66\x7f\x63\xf7\x73\xc1\x21\x57\x62\x8f\x4a\x87\x91\xe0\x25\x66\x86\x81\xe1\xab\x12\x13\xbe\xc5\xf6\x16\x06\x76\x97\x73\x83\x53\x39\xaa\x64\x51\x8a\xcc\xc0\xaa\xaa\x4a\x06\x0a\x4d\x57\x63\x90\xb5\x35\x06\x8f\x1b\x61\xb0\x14\xda\xc0\xf7\x1f\x0d\x43\xd0\xe9\xfd\x4d\x49\xd7\x07\x43\x77\xb8\xe5\x72\x5d\x62\x38\xce\x51\x9a\x99\xad\x0c\x2e\x4a\x91\xa1\xd3\x15\x4e\x66\x0c\xdc\xf7\x7c\xd6\x93\x07\x94\xf4\xec\x9f\x21\x38\xa2\x02\x4a\x14\x7e\x0e\xab\xd0\x04\x94\x92\x95\x2d\xe0\xdf\x53\xdc\x3d\x9a\x5b\x5b\x14\xa8\xfc\x80\x92\x1c\x0b\x54\x27\xc5\x07\xdb\x15\x57\xb6\x70\xf0\xac\x2a\xed\x56\x6a\x47\xe1\x45\xf1\xdd\xcd\x72\x92\xc2\x97\x9b\xc9\x32\x5e\x78\x94\x88\x02\x4a\x94\x7e\xaf\x12\x2e\x86\xf0\x8f\xb3\xeb\x88\x1b\x42\xb1\x35\xe1\x62\xa7\x84\x34\x85\xef\xf9\x97\x3a\x68\xf1\xe0\x7e\x7b\x8c\x12\x42\x1a\x9b\x75\xf8\x7f\x25\x4e\xd8\x18\x78\x0c\xbc\xa0\xeb\xe8\x14\x96\x3c\xc3\x4d\x55\xe6\xa8\x74\x3d\xf3\x52\xe3\x58\xe6\xf8\x74\x5a\x60\x6f\x74\x31\xb8\x66\x70\x1d\x04\x94\x1c\x28\x25\x4e\xd1\x5d\xab\x88\x12\xe7\x90\xbb\xc3\x1b\x27\x8b\x78\x9e\xc2\x38\x49\xa7\x70\xa9\xdd\x67\x9a\xc0\x68\x9a\xdc\x4d\xc6\xa3\x14\x6a\xa5\xc7\x8c\xb1\x7e\x44\x46\x89\x33\x6a\x30\x80\xac\x8b\x28\xe0\xd3\x0e\x33\xa3\xeb\x88\x1f\x93\x83\x7b\x94\x60\x36\x95\x5d\x1b\x78\x44\xe0\x0a\x41\x56\xa6\x49\xa2\x90\xeb\xfa\x55\x85\x5f\x95\x30\x78\xfb\x6c\xd0\xbf\xf2\xaf\x82\x93\xb3\x45\xed\x91\xff\xca\xaa\x3e\xbc\xb5\x53\xef\xb4\x7b\x01\x78\x4e\x9d\x28\xe0\xe2\x6c\x1d\x5e\x5e\x6a\x9b\x9a\xf3\x00\x86\xdd\xbb\x3b\x23\x89\xa6\x90\x4c\xd3\xff\xc6\xc9\xbd\xe7\x2c\x04\x2c\x35\xbe\xee\x3c\x51\xfc\x59\xc9\xef\x69\x8e\xa6\xb0\x7c\x88\x6e\xd2\x18\x16\x71\xda\x4c\x40\x8a\x4a\x81\x60\xb0\x77\x51\x54\x5c\xae\xb1\xdd\xe1\x5a\x88\x1b\x50\xf4\xe9\x3b\x53\xc6\x6a\x65\xe4\xe0\x1e\xbf\xdc\xce\xe4\xaf\x97\xa2\x5f\xa6\xb3\x3d\xda\xd7\xc8\xb7\x22\x1b\x92\x77\x4b\x1e\x0c\x21\xfe\x36\x9a\x2c\xa3\x38\x0a\xbd\x0f\xd0\x87\x26\x92\xed\x26\xb9\x9d\xed\xa7\x38\x27\x9e\xc7\xe9\x72\x9e\x8c\x93\x7b\xf0\x3e\x74\xba\xfe\x9b\xeb\x4c\x76\x77\x28\x34\x56\x49\x70\xa0\xb6\x3f\xa0\x87\x3f\x01\x00\x00\xff\xff\xcb\xf4\xe8\xf0\xc8\x05\x00\x00")

func templatesSingletonCrdb_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonCrdb_upsertGoTpl,
		"templates/singleton/crdb_upsert.go.tpl",
	)
}

func templatesSingletonCrdb_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonCrdb_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/crdb_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0x98, 0x7d, 0xd0, 0x22, 0xf5, 0xe7, 0x3b, 0x44, 0x9c, 0xfd, 0x2b, 0xd8, 0x4c, 0x42, 0xc0, 0x39, 0xf2, 0x5a, 0xef, 0xb7, 0x3a, 0xa3, 0xc9, 0x86, 0x70, 0x6a, 0x35, 0xd5, 0xfe, 0xb7, 0x5c}}
	return a, nil
}

var _templates_testSingletonCrdb_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x6d\x6f\xdb\x36\x10\xfe\x1c\xff\x8a\xab\x80\xb4\x52\xa0\xd0\xfb\x9c\x41\xc8\x1c\xdb\x29\x82\xb6\x49\x67\xa7\xeb\x5e\x3a\xac\xb4\x74\xb2\x89\x48\xa4\x42\x52\x75\xdd\x2e\xff\x7d\xe0\x8b\x2c\xc9\x71\xda\x6c\x1d\xb6\x4f\x89\xef\x9e\x7b\x79\xee\x8e\xe4\xe9\x03\x95\x20\x97\x1f\xc7\x93\xb3\xf3\x1b\xdc\x40\x02\x12\x97\xf8\xb1\x22\xaf\x6a\xa5\xc7\xa2\xac\x58\x81\xe1\xfb\xf0\xb4\x8c\xc2\xf0\x1d\x8f\x4e\xc9\xd1\xf8\xea\x72\x7e\x3d\x1b\x5d\x5c\x5e\x93\xa3\xd3\xf3\xab\xd9\xf4\xe2\xf9\x25\xbc\x98\xfe\x42\x8e\x4e\xdf\xf1\x3f\x2d\xe8\x37\x7a\xfc\x69\x74\xfc\x2b\xfc\xf1\xfb\xd1\x4f\xa3\x97\x17\x93\xd1\xf5\x14\x7a\x66\x16\x1b\xbd\x8f\x06\x03\xbd\xa9\x10\x52\x99\x2d\xae\x51\x69\x94\xa0\xb4\xac\x53\x0d\x9f\x07\x00\xd9\x62\x2c\x38\x87\x23\x75\x5b\x90\xc9\xd9\xc0\x4a\x2e\x69\x89\x60\x40\x8c\x2f\x07\x00\x2b\xa1\x34\x40\x47\x50\x2b\x94\x3d\x41\x45\x95\xea\x09\x94\x2a\x4a\x91\x61\x07\x21\xa4\xf5\xc1\xb8\x36\x31\x34\x2a\x3d\x39\xb3\x71\x3c\x64\x2b\x7c\x33\x7b\xd9\x9a\x65\x0b\xf3\xb3\xe3\xfa\x6e\x30\xc8\x6b\x9e\x02\xe3\x4c\x87\x91\x67\xf0\x8a\x32\x0e\x09\x3c\x6d\x19\x7e\xbe\x33\xc8\xe1\x10\x14\xea\xba\x82\xac\x2e\x2b\x05\x7a\x85\x90\x51\x4d\x17\x54\x21\xa8\x74\x85\x25\x05\xca\x33\x60\xa5\xc9\x4e\x01\xd3\x26\x3d\x01\x14\x34\x1a\x11\x95\x1b\x90\x94\x67\xa2\x2c\x36\xc6\xd7\x12\x39\x4a\xaa\x31\xb3\x99\x76\x5c\x09\xd0\x2b\xaa\xad\x54\x41\x4a\x39\x2c\x10\x64\xcd\x81\x2e\x29\xe3\x4a\x1b\xc7\xb5\x62\x7c\x69\x32\xe8\x3b\x52\xb7\xc5\x42\xb0\x02\x25\x5c\xcd\x5e\x41\x45\xd3\x1b\xba\x44\xe2\x28\x86\x29\x1c\xb5\x8c\x22\x47\x25\x8c\x00\xa5\x14\xd2\x32\x37\x53\x85\x52\x3a\x89\xa9\x60\x4a\x7c\xf3\x12\xf8\xc0\x2a\x94\xe4\x39\xea\xb9\x2d\x5c\x18\x18\x57\x24\x5b\x70\x5a\x62\x10\x59\xac\xed\xeb\x43\x48\xa3\xf4\x38\xdb\xee\x87\x70\x46\xe9\x71\x76\x0a\x1e\xc2\x19\x65\x83\x33\xb3\xd0\xc1\x5d\x70\xdd\x80\x84\x6c\x82\x36\x13\xf4\x90\x3f\xaf\xb7\xe8\xe1\x10\xc6\x12\xa9\x46\xa0\xbe\x61\xec\x13\x66\x90\x2d\xc0\xb0\x25\xd6\x5f\x67\xe2\x92\x16\x44\xe6\x9a\x2e\x0a\x74\x8a\xb0\x29\x5f\x34\xe8\x98\x98\x01\x4c\x60\x51\xb3\x22\xfb\xb1\x46\xb9\xf1\x59\xb8\xb2\xc4\x9e\x76\xdc\x8b\x10\xfb\xe2\xc6\x9e\x6c\xdc\xf2\x89\x7c\x97\x1e\xed\xd5\x65\xf4\x45\x8f\x03\x00\x96\xdb\x41\x48\x8c\x81\x14\xd5\xb5\x4d\x25\x8c\xbe\xb7\xd2\x27\x09\x70\x56\xd8\x89\x01\x90\xa8\x6b\xc9\x8d\x7c\x00\x70\xd7\x37\x4d\x6d\x11\x1f\x6f\x6c\x0e\x5f\x5d\x56\xe3\x32\x83\x93\x04\xf0\x23\xa6\x64\x2c\xca\x92\xf2\x2c\x0c\x52\x91\xde\x48\x41\xd3\x55\x10\x43\x60\x50\x41\x97\x4d\x70\x7c\x5c\xcb\xc2\x8b\xde\xcc\x5e\x5a\x09\xe3\x0a\xd3\x5a\x62\x60\x7f\x19\x9b\x63\x43\x30\x71\x47\xd5\xcd\x85\xcd\xf0\xeb\x01\xd5\x6d\x11\xf4\xa2\x6c\x7b\xe9\x7c\xfb\xa3\x1b\xec\xf6\xad\x9b\x85\x2d\xac\x8c\x61\x6d\x62\x31\x41\x5e\xb3\x0a\xc3\xa8\xe5\x4c\xe6\x3a\x13\xb5\x99\xe4\x75\x37\x33\x23\xb6\xd7\x11\xc7\xf5\xf9\x0b\xdc\x4c\x50\x69\x29\x36\x28\xc3\xf6\x01\x88\x41\xf6\xdb\xd6\xba\xa4\x52\xef\x2d\x7d\xb7\xf8\x42\x2a\xf2\x56\xd2\x2a\x44\x29\x63\x08\x72\xca\x0a\x73\x29\x09\x50\xc6\x1c\x9e\x6d\x4b\x61\xfd\x3e\x83\xd4\xd5\xc8\x56\xb0\xdf\xf2\x4e\xd2\xff\x6e\x60\x75\x5b\xec\xc6\xdd\xc7\xf7\x2d\x65\xdf\x10\x75\x4d\x99\x86\x5c\xc8\xaf\x07\x1e\x0e\x61\x94\x9b\x97\xaf\x99\x57\xa6\x20\x13\xdc\x9c\xab\x42\x28\xb4\x4f\xc3\x5a\x32\x8d\x80\x3c\x03\x91\x5b\x41\xc5\x2a\xec\x26\xbd\x26\x63\x03\xfe\xe7\xf9\xba\x58\xc6\xed\x9e\x92\xb4\xbd\xf8\xaf\x8a\xe2\x9d\x71\x56\x6c\x9f\xd5\xdd\x37\x47\xd6\x7c\x5c\x66\xa1\x32\x23\x1d\x37\x1e\xfc\x5b\x1c\x03\x95\x4b\x05\x84\x10\xf7\xbb\xfb\x32\xa5\x7b\x4e\xa8\xb7\x76\x66\x84\x90\xc8\xe1\xc8\x94\x7f\x80\x04\x84\x32\xff\x30\x29\x78\xd8\x1c\x8d\x02\xb9\x8b\x1c\x99\x32\x7c\xe7\x8b\x90\x76\x8e\x98\x0b\xac\xc8\x25\xae\x67\x48\x33\x94\x1e\xdf\x30\x54\xee\x84\x9e\x24\xf0\x74\xb1\xd1\xa8\xc8\x59\x9d\xe7\x76\x3b\xb0\x3a\x53\xe0\xbd\xba\xb4\x7b\xba\x9d\x93\x56\xea\xda\xe5\xcc\xdb\x06\x9e\x24\x56\x3f\xab\xf9\xde\xd6\xe5\xa5\x26\xaf\x25\xe3\xba\xe0\x61\xd3\x2c\x59\x73\xce\xf8\xf2\x24\xd8\x56\xd6\xd5\x26\xba\x67\xe1\x52\x20\xfe\x91\x88\xf6\x02\x50\xca\x1d\xc0\xfd\xdb\xfa\x11\x1d\xef\xbf\x01\x9d\x96\x2a\x5d\xda\x4a\x9a\xb0\xf3\xca\xc4\xcd\xc3\x60\x3c\x9b\x9a\xb5\x73\x32\xba\x1e\x9d\x8d\xe6\x53\x38\x54\x3b\x57\x6a\xd4\x46\x4d\x89\x9f\xa6\xc0\x5c\xcd\x7f\xe7\xba\xee\x3f\x0c\x66\xa8\x6a\x6d\x7e\x98\x94\xa2\x07\x99\x74\x1f\xc2\xaf\xf1\x98\xcc\xae\x5e\xb7\x2c\x2e\xce\x61\xfa\xf3\xc5\xfc\x7a\x0e\x87\x0a\xc6\xa3\xf9\x78\x34\x99\xfe\x2f\xbc\x86\x43\xd0\x48\x65\x26\xd6\x1c\xbc\x5a\x41\x5a\x20\xe5\x75\x05\x9a\xaa\x1b\x05\xeb\x15\x72\x7b\x5b\xb9\xed\x33\x67\x9c\xa9\x55\x33\x5a\xfb\x0b\xd3\xb8\xfc\xc2\x2e\xd9\xdf\x28\xec\x27\xc2\x17\x2e\xbf\x7b\x3b\x45\x63\x03\x16\xf5\x4d\x0b\xca\x23\xa7\x56\x70\x43\x27\xf4\x9f\x31\xb1\xa3\xe1\xbe\x0f\x58\xde\xe6\xb3\x2f\x52\xa3\x8c\x6d\x04\x1f\x73\xb7\x1c\x2d\xc8\x1f\xff\xdb\x82\x5c\x55\xc8\xc3\xa0\x12\x4a\x2f\x25\xaa\x7e\x7f\xa3\x96\xf3\xbe\x98\x9c\x15\xf1\x1e\x8a\xfd\x54\x1a\xb2\xf7\xb6\x44\xb7\x23\xba\x0d\xd1\x6d\xf4\xb1\xfb\x48\x6b\x2e\x66\xbb\x61\x33\xae\xe3\x9d\x0f\xb1\xc8\xff\x85\xcf\x83\x03\xc3\xb0\xae\x9a\xef\xaa\x03\x96\xbb\xcf\xba\x27\x09\x04\x81\xd1\x1f\xd4\x15\x24\x56\x36\x38\xb8\xb3\x7a\xbb\xdf\xef\xe8\x7b\xa7\xe8\x50\x9d\xd8\x0b\xa0\xae\x5c\x76\x91\xb1\x1c\x1c\x78\x72\x3d\x68\x53\xb5\xdb\xe2\x64\x38\x3c\x54\x3f\x18\xd3\x6c\x78\xa8\x4e\x7d\xc2\xc9\xd6\x91\x5b\x7b\xdd\xd2\xdb\x90\xdd\xae\xbe\x77\x7f\x05\x00\x00\xff\xff\xaf\x06\x6e\x3c\x59\x0f\x00\x00")

func templates_testSingletonCrdb_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonCrdb_main_testGoTpl,
		"templates_test/singleton/crdb_main_test.go.tpl",
	)
}

func templates_testSingletonCrdb_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonCrdb_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/crdb_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x78, 0x2d, 0x5b, 0x41, 0xf5, 0x39, 0x92, 0xf1, 0xc8, 0x65, 0xf5, 0x2, 0xb, 0xea, 0xa8, 0x4b, 0x94, 0xc9, 0xaa, 0xab, 0xf9, 0x34, 0xde, 0x35, 0xb5, 0x0, 0x16, 0x24, 0x3b, 0x9e, 0xda, 0xe2}}
	return a, nil
}

var _templates_testSingletonCrdb_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xc1\xca\x83\x30\x10\x84\xef\xff\x53\x0c\xe2\x41\x7f\x34\x0f\xf0\xc3\x7f\xea\xa9\x3d\xf4\x50\xec\x03\xa4\x75\x95\x40\xba\x15\xb3\x42\x21\xe6\xdd\x8b\x31\x58\x7b\xcb\x64\xe6\xdb\x9d\xed\x26\xbe\xa3\x21\x27\xd7\xc1\xd1\x28\x85\xe0\x57\xc8\x89\xe1\x5e\x35\x25\xfc\x0f\xe0\x7d\x8d\x51\x73\x4f\xc8\x0d\xb7\xf4\xaa\x90\x8b\xbe\x59\xc2\xdf\x3f\x54\xb3\xbc\x5c\x08\x29\x67\xba\x64\xaa\xa3\x3b\x3d\x0d\x47\x1b\xf5\xe6\x93\x75\x7b\xb9\x66\xcf\xfa\x11\x87\x25\x32\xca\x19\x83\x9d\x46\x6d\x31\x43\x8c\x58\x3a\xe8\x0d\x14\x75\x99\xb8\xc8\xbc\xff\xd0\x21\x64\x15\x96\xda\xdf\x9f\xeb\x49\x65\x5c\x46\xdc\xee\x7b\x24\x15\xde\x01\x00\x00\xff\xff\x91\x6e\x4f\x2d\xff\x00\x00\x00")

func templates_testSingletonCrdb_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonCrdb_suites_testGoTpl,
		"templates_test/singleton/crdb_suites_test.go.tpl",
	)
}

func templates_testSingletonCrdb_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonCrdb_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/crdb_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3c, 0xee, 0x2e, 0x96, 0xb6, 0x68, 0x8e, 0xe6, 0x13, 0x5b, 0x19, 0xc, 0x3f, 0x8d, 0xbd, 0x45, 0xf6, 0x91, 0xd7, 0x45, 0x15, 0xdd, 0x12, 0x15, 0x2a, 0x56, 0x1f, 0x17, 0xb9, 0xb8, 0xb3, 0x99}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x41\xaf\x29\x7c\x70\x7e\x0e\x41\x51\xc3\x88\xe5\x73\xc1\x48\x2b\x87\x30\x4d\x0a\xe4\xaa\xb6\x2b\xf0\xdd\x0b\x52\x4e\xe2\xfc\x15\x46\xd1\xa2\xe8\xc1\x96\x48\xcc\xce\xec\xce\xee\xaa\xef\x4f\xe0\x7f\xa9\x95\xf4\x70\x36\x06\x31\x89\x6f\xe8\x45\x29\x6f\x35\xc2\xf0\x10\x53\xb9\xc6\x10\x58\xd3\x99\x0a\x08\x3d\xf5\xfd\x10\x21\x16\xed\x4c\x77\x4e\xea\x10\x16\xad\x47\x47\x9c\xe0\x43\x04\x28\xb3\x14\x65\x0e\x3d\xcb\x48\xcc\xa4\x93\x5a\xa3\xe6\x39\x63\x99\x6a\x40\xa3\xe1\x0f\x04\x97\x76\x63\xe6\xca\x2c\x3b\x2d\x5d\x08\x13\xad\x2f\xac\xee\xd6\xc6\xe7\x30\x1e\xff\x0c\x39\x73\x6a\x2d\xdd\xee\x33\xee\x1e\x02\x7a\x96\x65\x24\xe6\x2b\xd5\xf2\x51\xfc\x6f\x95\x59\x02\xa5\x32\x36\x8a\xee\xc0\x1a\xbd\x83\x76\x88\x83\x15\xee\xa0\x1a\x22\x47\x39\xcb\x02\x63\x99\x47\xac\xa3\x05\x4e\x9a\xda\xae\xd5\x77\x14\x53\xdc\xcc\x11\x6b\x9e\xb3\xec\x9b\x74\x80\x2e\xfd\xac\x63\xd9\xe9\x29\x4c\x88\x70\xdd\x12\xd0\x1d\xc2\xf5\x74\x7e\x75\x53\x82\x57\x35\x82\x6d\x40\x1a\x58\xcc\xe2\x0d\xcb\x6c\x64\x3c\xb0\xeb\xb1\x82\x3e\x24\x37\x22\xe9\xa1\xe6\x9c\x5c\x57\x11\x8f\xc9\x14\xf0\xde\x16\xf0\x86\x01\x97\xe7\xe5\xae\x45\x5f\x00\xb9\x0e\xf3\x4f\x89\xe7\xbf\x31\x18\xa5\xf7\x46\x5c\xc5\x4c\x1b\x3e\x5a\x98\x64\x01\xd9\x47\x91\xd7\x13\x02\x9f\xa4\xcf\xe0\x9d\x1f\x15\x91\x6f\xef\x4b\xdf\xab\x06\x8c\x25\x10\x53\x7b\x61\x0d\xe1\x96\x42\xa8\x68\x1b\x2b\xab\x86\xb3\x38\x97\xd5\x6a\xe9\x6c\x67\x6a\x9e\xf7\x3d\x9a\x3a\x04\x96\x0d\x90\x2f\x9d\xa7\x72\xcb\x13\xcb\x21\xc3\x8b\x8b\x5b\xab\xb4\x38\xc7\xa5\x32\x89\x43\x7b\x3c\xbc\x2b\xb7\xbc\xa2\x6d\x11\x0b\xbc\x57\x38\x0a\x94\xb3\xac\xc6\x06\x1d\xc4\xe1\xe5\x39\xf4\xf0\x15\xc6\x40\x5b\x71\x63\xb5\xbe\x95\xd5\x8a\xe7\x10\x62\x87\x1f\x7a\x61\xc5\x7e\x96\xdf\x2a\x3c\xf6\x04\x4d\x0d\x27\x21\x40\x3c\x35\x52\x7b\x4c\xa2\x05\xa4\x5c\xae\x4d\x83\x8e\xe7\x4f\x4f\xc7\xf5\xa8\x4b\xd2\xaf\x37\xe8\x45\x67\x2a\xdb\x19\x4a\x17\xcf\xa6\xec\x7e\x29\x79\x2e\x2e\x22\xe6\xc8\x52\x1e\x5d\x78\x99\x25\xbf\x97\x8d\x90\x24\x1c\x41\x1f\x9f\x40\x46\x1b\x69\x08\xac\x41\x70\x58\x59\x57\x17\xb0\xb4\x74\x36\x2a\x06\xfc\x3e\xe9\x67\xab\xb3\x98\x5d\x4e\xca\xab\xd7\x56\xe7\x77\x2c\xc7\xbe\x35\xc7\x7e\x44\x84\x10\x7f\x74\x95\x7e\x7d\xc6\xe2\x96\xff\xe5\x11\xfb\x47\x26\x2c\xfc\x08\x00\x00\xff\xff\xdc\x3f\x97\x97\xd1\x06\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x58, 0x5e, 0xcb, 0x8a, 0xd4, 0x31, 0x25, 0x32, 0x7a, 0x7f, 0x6, 0x25, 0x4b, 0x8d, 0x87, 0x3b, 0x70, 0x44, 0x98, 0x4d, 0x7, 0x12, 0xce, 0x74, 0x38, 0x36, 0x58, 0x21, 0x78, 0x5c, 0xd6, 0x41}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,

	"templates/singleton/crdb_upsert.go.tpl": templatesSingletonCrdb_upsertGoTpl,

	"templates_test/singleton/crdb_main_test.go.tpl": templates_testSingletonCrdb_main_testGoTpl,

	"templates_test/singleton/crdb_suites_test.go.tpl": templates_testSingletonCrdb_suites_testGoTpl,

	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"crdb_upsert.go.tpl": &bintree{templatesSingletonCrdb_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"crdb_main_test.go.tpl":   &bintree{templates_testSingletonCrdb_main_testGoTpl, map[string]*bintree{}},
			"crdb_suites_test.go.tpl": &bintree{templates_testSingletonCrdb_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
