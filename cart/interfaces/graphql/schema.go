// Package graphql Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// schema.graphql
package graphql

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xcf\x6f\xdc\xb8\x15\xbe\xfb\xaf\xd0\x20\x97\x29\x10\x14\x68\x8f\x73\xb3\x67\x9c\x60\xb0\xb1\xe3\xd8\xb3\xe9\x21\x08\x0c\x5a\x7a\x33\xc3\x86\x22\xb5\x24\x65\x5b\x58\xec\xff\x5e\x90\x22\x25\xfe\x92\xc4\xb8\xe8\x5e\xba\x3d\x74\x33\xe4\xc7\xc7\x47\xf2\xe3\xe3\xf7\x9e\x2c\xbb\x06\x8a\x2d\xab\x6b\xe0\x25\x3c\xee\xa0\x64\x1c\x49\xa8\xb6\x88\xcb\xe2\xf7\x8b\xa2\x28\x8a\x12\x71\xb9\x19\x21\xaa\x67\xa5\x3b\x2a\x0b\xde\x01\xc1\xcf\xc0\x31\x88\x4d\xf1\xcd\x03\xee\x02\x48\xb7\xfa\xae\x87\x9e\x20\xee\xba\xea\xb6\xac\x82\x75\x65\x7e\xaa\x1f\x9b\xe2\x41\x72\x4c\x4f\xab\xbf\x05\x0e\x44\x83\xad\xd5\x4b\x42\xee\x50\x57\x03\x95\xf7\xf0\x5b\x8b\x39\x54\x7b\x09\xb5\x08\x86\x3f\xde\x71\x5c\x9a\xae\xd5\xb0\xc8\x87\xb6\xae\x11\xef\x42\xac\x69\x5e\x5d\xfc\x71\x71\xe1\xef\x96\xdb\x6d\x36\xab\xc2\xa2\x64\x2d\x95\xe1\x8c\x97\x4d\x43\x30\x54\x3b\xdb\xad\xc1\xa2\xad\xc3\x76\x67\x98\xf6\x31\xc0\x7d\xc4\x47\xb9\x45\xbc\x9a\xc4\x7d\xe4\x88\x56\x07\x26\x11\xf9\x17\x96\xe7\x45\xb8\x46\xda\xc9\xbd\x11\x97\xb5\x6a\x4a\x8e\x3b\x23\x11\xbb\x7d\xc5\x18\x01\x44\x57\x83\x65\xf4\x0a\xd1\xb6\xeb\xc6\xf4\x3e\x9a\xfd\xc3\xd5\xa6\xd8\xef\x7a\x2b\x40\x25\x96\xdd\x7e\x37\xb0\x40\xb7\x3e\x61\x42\x30\x3d\x5d\x56\x15\x07\x11\x6d\x73\xdf\xaa\x81\x4d\xcb\xcb\x33\x12\xc0\x03\xcc\x1d\x70\xc1\xa8\x61\xf0\x34\x71\x3d\xbe\xa2\xaa\xc2\x12\x33\x8a\xc8\x0e\x49\x14\x4f\xea\x74\xf6\x5e\x36\x3d\x09\x1f\x80\x40\xa9\xfa\x22\x02\x06\xfd\xfd\xd2\x80\x30\x7a\x12\x07\x76\xd9\xca\xb3\x5a\x7d\xa9\x28\xfe\xab\x5e\x82\xb7\xbf\x28\xec\x0f\x37\x09\xf5\xe7\xb3\x65\x6d\xc3\xa8\xba\x49\xd1\x02\xc7\x2e\xb3\xc4\x0a\x8e\xa8\x25\x72\xdb\x72\x0e\xb4\xec\x7c\x7b\x52\xf1\x04\xf7\x37\xc9\xb7\x73\xb0\x3d\xc6\x8c\xfa\xe7\xb6\xa7\xce\x9e\x9a\x40\xd1\x70\x56\xb5\xa5\x0c\x9b\xb1\xf0\x76\x01\xaa\x60\x95\xa7\x81\xcb\x21\x0d\x57\x1e\x7f\x0f\xe8\x35\xcd\x56\x0b\x7b\xd2\xb0\x5b\x98\x00\xa0\xe8\x6e\x7d\x4b\xdd\x5d\xdb\xef\x86\xb0\xb7\x44\x2e\x3f\x60\xed\x9c\x41\x6a\x66\x33\xec\xfb\x85\x05\xdc\x20\x4c\x1f\xce\xb8\x69\x30\x3d\x5d\xdf\x20\x4c\xfc\x93\xc1\xe2\xba\x6e\x64\x17\x6c\xdd\x19\x09\x6b\xf8\x03\xe3\xb3\xde\x0d\xe3\xe2\x55\xa9\xf8\xb8\xdf\xad\xb1\xfe\xcf\xe2\x8a\x56\xd6\x40\xee\x40\x85\x1a\x06\xe9\x23\xfa\x22\xbb\x75\x8d\xf8\x0f\x90\x77\x04\x95\xe0\xb9\xfa\xbe\x78\x46\x1c\x23\x2a\xc3\x05\xec\xa9\x1c\x67\xbe\x7e\x95\xc0\x29\x22\xf7\x70\x04\xc5\x63\x58\x73\x38\x2e\x78\x60\x47\x7f\x65\x6d\x79\x06\xfe\x80\x9e\x31\x3d\x45\x21\x73\xf0\x54\xb3\x1e\x12\x81\xe5\xb1\x6f\x35\x06\x45\x5b\xdb\x63\x9b\x64\x9e\x8f\x51\xf1\x77\xf2\x21\x18\xce\xd5\x0e\xd8\x32\x11\xc5\x5d\x44\x88\xed\x3e\x60\x49\x12\x84\xb2\x97\xe1\x23\x67\x62\x62\x0e\x0f\x92\xe1\x93\x73\xbf\xb2\xd0\xfe\xa3\x33\x7f\x73\xeb\x5b\x46\xd5\x21\xdd\x03\xd1\xaf\x7d\xde\xa0\x9f\x1c\x31\xbe\x67\x63\x50\x4c\xdc\x8b\x41\x57\x18\x66\xf9\xf7\xd0\x52\x58\x6b\x8a\xab\xee\xd0\x35\xb0\x56\xaf\x5c\xc8\xd6\xf9\xe8\x39\x86\xbc\xed\x19\xf1\x13\x44\x9b\xf8\x68\xda\x8d\x5b\xa3\xeb\x4e\xf4\x0a\x23\xc1\x3d\xd4\x08\x53\x4c\x4f\x29\x4c\x5a\xd4\x38\xfa\xc8\x51\x81\x46\x4a\x05\x6b\x30\xe0\xe1\x3e\xf5\x2b\x11\x86\x87\xb3\x63\x1e\x1c\x90\x19\x27\x87\x4d\x0c\xf7\xca\x8c\x19\x76\x79\xf5\x7d\xd6\x79\xeb\x8f\xf1\x1f\xcd\x10\x20\x88\x53\xb3\x66\x5d\x97\x33\x4c\xdb\xa8\xbb\xa7\x47\xe6\x51\x61\x76\x92\x61\x8d\x19\x33\x94\x19\x56\x0f\xe8\x35\xc3\x92\x1a\xe8\x93\x5a\x49\xec\x4d\xf1\x81\x30\x24\xa7\x2d\x83\xa5\x48\x52\x1f\x28\xc4\x77\xe7\x69\xe8\x2f\x06\x7a\x3d\x38\x93\x85\x61\x59\x8d\x99\x5c\x8a\x8e\xb1\x66\x46\x5f\x58\x0c\x2f\xc1\x7e\xd4\x20\xfa\xa7\x69\x4e\x3f\xb5\x9a\x45\x98\x4a\xe0\x47\x54\x46\xc7\x11\xc8\x34\x33\xef\x09\x49\x78\x41\x29\x8d\xf4\x15\x91\x16\xe2\xed\x4d\xaf\x65\xd7\x4b\xae\x68\x12\x5c\x37\x04\x54\x93\xf8\x33\xdd\x89\x72\x2a\x9b\xd2\x98\x9f\xb3\xcf\xfe\x90\x0b\x26\xaf\xee\xce\xed\x4d\xa4\x80\xf6\xae\x5e\x75\xfb\x6a\x6d\x52\x80\xc9\x94\x4f\x01\xa7\x56\x90\x74\x5c\xdd\xbd\x09\xe7\x55\xd7\x10\xde\x92\xfc\x9d\x08\x69\x81\x3d\x37\x2a\x2c\xbf\xb3\x0b\xea\x36\x4f\xdc\x2e\x69\xdb\x9f\x78\x6d\xdf\xf2\xd8\xfe\xf4\x5b\xfb\x93\xda\xe2\x0d\xd2\xe2\x8c\x84\x61\xdf\xfc\xeb\xe6\x1e\xbe\x7d\xdd\xbc\x20\xaa\x5a\x5e\x18\xff\x71\x24\xec\xc5\x6f\xad\x41\x9e\x59\xe5\xb7\x95\x88\x73\xac\xc4\xa0\xdb\x68\xb9\xf7\x89\x95\x28\x91\xff\xed\x82\x6e\x33\x46\x60\x0e\xd5\x01\xd7\xb0\x29\xd4\xff\x0f\x45\x0d\x2f\xc1\x5c\xff\x80\xce\x55\x14\x5e\xde\xe7\x21\x7f\x81\xce\x53\x80\x0a\xf1\x2e\x80\x39\x7b\x21\x36\x45\x8d\x9a\x6f\xa2\x0f\x8b\xff\x16\x8c\xfe\xfd\x1e\xbd\xdc\x80\x10\xe8\x04\x19\x83\x6f\x50\x33\xa2\x7c\xb7\x1d\x60\xe8\xfe\x0d\x6a\x22\xdf\x1d\x78\xb8\x86\xd9\x13\xb5\xdb\x59\x98\x63\x8d\x5f\x34\xb4\x58\x36\x68\x05\x5c\x05\x25\x06\x4f\x50\x65\xbc\xb7\x09\x8d\x20\x95\x1c\xf7\x5d\x69\x14\x73\xa7\x2e\xae\x9c\xbd\xf6\x68\xba\x6a\x94\x2a\x36\x39\x0f\x82\x7b\x8d\xf6\xb4\x54\xe1\x65\x42\x0c\x78\x1d\x0b\xaf\x72\x38\xe1\x9c\x20\x08\xb0\x86\x96\x4f\xdd\x16\xd5\x0d\xc2\x27\xad\xbe\xd7\xa5\xf3\xc3\x51\x09\x39\xcb\x7c\xea\x25\xc6\x11\x13\x09\x7c\x4e\x65\xc4\xc3\x73\xd6\x36\xc8\x61\xd7\x41\x3f\x1e\x38\x49\x44\xe1\x77\x11\xf4\x04\xa4\x17\x25\x61\x97\x39\x52\xdb\x39\xad\xcf\x92\xa3\xb1\x70\xe2\x70\x58\x8c\x63\x5c\x7e\xe6\x95\x8a\x50\x46\x0d\x2d\x09\x00\x87\xb7\x38\x7e\xeb\x86\x37\xce\xa8\x2f\x8f\x3f\xba\x25\x6d\xde\xb5\xea\x56\xf9\xc2\x8c\x3d\x88\xb8\xba\x1c\xd0\x44\xe5\x00\xdd\x69\x2a\x02\x37\x13\x25\x03\xd7\xcb\x5b\x54\x07\x1d\x82\xb5\xbc\x84\xb0\x72\xf6\x9b\xec\x9c\x12\xd5\x72\x3c\xf5\x11\x5a\x6f\x45\x98\xcc\x10\x3e\x64\x74\xe1\xa4\x21\xdc\x1c\x6f\xbf\x0a\x4c\x4f\x04\x34\x4b\xe6\x72\xfa\x11\x35\x59\x8c\xe0\xec\x65\xc9\x8c\x85\x2c\x95\xd2\x72\x03\xd3\xf8\x5c\x70\xf6\x12\x56\x8c\xf5\xef\xa9\x4b\xd9\x87\x66\x43\xa7\x67\x24\x9d\x7b\x91\xbe\x21\x47\xcc\x85\xa4\x9a\x04\x93\x18\x82\x92\x10\x9f\x8f\xb8\xaa\x08\xdc\x46\x28\x4f\x7a\xf7\xc1\x7e\xd6\x1f\x81\x48\x2b\x8d\x34\x98\xc4\x48\x0e\x90\x58\x5a\x8c\xb9\xe5\x73\x3e\x8f\x1c\x35\xfb\xf6\x09\xd3\x98\xa5\x25\xab\x1b\x44\xbb\x68\x3a\x2f\xb6\x61\x19\x03\x02\x4c\xc3\x84\x1c\xa2\xdf\xa4\xd7\x3a\xb3\x9c\xb5\xc3\xe1\x84\x9d\x38\x9a\xf6\x47\xd1\x88\x2f\xf8\xdc\x63\x22\x43\xde\x89\x01\x81\xe6\xcc\xe8\x1c\x3b\xa0\xd6\xc5\xd7\x49\x9f\x93\x44\xed\x3f\x36\xd8\xe4\x7b\xf9\x9b\x85\x86\x2b\x05\x24\x11\x26\x21\xf2\xce\xef\xb5\xf1\x13\x0b\x89\xe9\x69\xdb\x0a\xc9\x6a\xe0\x89\x0f\x14\xd7\x09\x48\xda\xdd\x14\x32\x88\xd9\x33\xcb\x1c\x3c\xb3\x09\x18\x92\xf0\xf9\x78\x85\xb9\x3c\x07\x31\x19\x09\xd1\x30\xde\x27\xee\xbc\x4b\x77\xde\xb6\xf5\x53\x28\xab\x29\xea\x79\xac\x69\x38\xbb\xf1\x7e\x10\x35\x0e\xe9\x50\x53\xea\xb5\x5d\x4a\xc9\xf1\x53\x2b\xc1\x11\xae\x1c\x04\xf0\x67\xa8\xf4\x6b\xb9\x58\x10\x1a\x6a\x77\x93\x39\xc4\x94\xe8\xcb\x29\xbf\x24\xa7\x1c\xeb\x93\xc9\x39\xe7\xf4\x8b\xad\xfd\x4d\x3a\x3b\x08\x90\x64\xe0\xb7\x25\xc4\xc9\xcc\xeb\x7e\x44\x2c\xd4\x16\xbf\x22\x82\x2b\x7d\x8e\xf7\x20\x5a\x62\x15\xd5\x19\x09\x85\x63\xf4\x9a\x73\x36\x86\xb3\x40\x7b\x0f\x00\x93\x96\xfc\x02\x01\x7b\xde\x61\x2d\x84\x94\x61\xe1\x5c\xd6\xa0\x28\xa5\xb4\xc8\xe8\x87\x36\xd8\xe7\x16\xef\x12\x0e\x27\xc0\xc5\xef\x17\x9a\x4b\xb6\x72\x18\x05\x03\xdd\x0b\x53\x8e\xbe\x9b\x28\xfe\x29\x05\x63\xd8\x37\xe6\xa6\x4c\xfd\xb6\x97\x21\x88\x49\xfe\x47\x9d\x85\x58\xf4\xe8\xe7\x35\x1f\x18\xb7\xd4\x3d\x32\x5e\xf7\x71\xa3\xff\xdf\xd2\x30\x1d\x40\x0a\xfd\xf2\xda\x6d\x49\x14\x59\x1e\x15\xd4\x39\xee\xa1\xd8\xd2\x70\x56\x82\x10\x8e\x52\x9d\xfa\xb6\x6e\x3e\x0c\x0e\xa5\x77\x87\x2f\xff\xe3\xa9\x13\x06\xec\x66\x61\x20\x95\x66\x41\x5c\xec\xd4\xa3\x3e\x0c\x80\x41\xff\x51\xe0\x88\xcc\x8d\x71\x28\x38\xe5\x8c\x25\x5e\xa1\xeb\x10\x69\xf6\x9b\x8f\xb8\x9f\xfa\x3c\x63\x9e\x0d\x81\xaf\x6f\xb2\x3c\xec\x87\x2f\xaf\x73\xe9\xe7\x84\xe6\xbf\x34\xdc\xa2\x86\xb3\xca\xed\x1f\x9b\x65\xcc\x3f\xa7\x30\xff\xcf\x2a\x4f\x2b\x3c\x27\x92\xa6\x30\x19\x2a\xef\x02\xd3\xa6\x95\x23\xb9\x63\x5e\xef\x35\x20\x87\xd8\x7f\x22\xaf\x33\x68\x9d\xc1\xea\x0c\x52\x67\x70\x3a\x83\xd2\x19\x8c\xce\x20\x74\x06\x9f\x33\xe8\x9c\xc1\xe6\x0c\x32\x67\x70\x39\x83\xca\x19\x4c\xce\x25\x32\xbc\x4a\xa0\x95\x16\x9e\xc5\x97\x76\xfc\x70\xe2\x85\xed\xcd\xc4\x5f\x07\x46\x16\x6e\x0c\x6b\x42\x23\x97\x55\x75\x60\x6a\xc4\x3a\xaa\xe5\xec\x77\xab\xf7\x63\xc5\xe5\x7d\xb1\xf8\x17\x34\x81\x07\xde\x3c\x3b\x20\x20\xc1\x2d\x07\x2f\xff\x45\xce\xb2\x3d\x25\xfe\x86\x3f\x64\xd1\xfe\xfe\x57\x46\x7f\x6d\x54\x52\xa4\x8c\x7e\x91\x5d\x86\x5d\x67\x7b\x72\xa7\xd0\x6f\x6d\x3f\x8f\x1f\x99\xd6\x68\x8c\x50\x9b\xc5\xf0\x15\x7d\x94\x8d\x71\xd3\x13\x07\xc2\x6d\x1d\x7e\xa1\x7c\x1f\x7e\x49\x89\x66\x4b\x4a\xbf\xd4\x84\x2a\xbf\xe9\xc6\xc4\xe8\x33\xb7\x99\xce\xba\xcc\x3a\xa0\x84\xc9\x7b\xa8\xd9\x33\x0c\x76\x4e\xe6\x1f\x79\x07\x3e\x69\x6f\xf4\x71\xed\x96\x88\x17\xed\xfd\x71\xf1\x9f\x00\x00\x00\xff\xff\xb6\x8b\xb3\xc3\xa4\x2b\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"schema.graphql": schemaGraphql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
