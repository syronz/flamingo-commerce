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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xcd\x6e\xdb\x38\x17\xdd\xe7\x29\x14\x64\xe3\x0f\x28\x3e\x60\x66\xe9\x9d\x63\xb7\x85\xd1\x26\x4d\xd3\xb4\xb3\x28\x8a\xe0\x46\xbc\xb6\x39\xa5\x48\x0d\x49\x25\x31\x8a\xbe\xfb\x80\x14\x25\xf3\x4f\x3f\xd3\xc1\x74\x33\x93\x45\x51\xeb\x1e\x5e\x5e\x92\x87\x87\x47\x94\x3e\xd6\x58\xac\x45\x55\xa1\x2c\xf1\x7e\x83\xa5\x90\xa0\x91\xac\x41\xea\xe2\xdb\x59\x51\x14\x45\x09\x52\x2f\x4f\x10\x13\x39\xb7\x01\xd2\x81\x37\xc8\xe8\x23\x4a\x8a\x6a\x59\x7c\x0e\x80\x9b\x08\x72\x3c\xff\x62\x9b\xee\x31\x0d\x5d\x1e\xd7\x82\xe0\x82\xb8\x9f\xe6\xc7\xb2\xf8\xa0\x25\xe5\xfb\xf3\xff\x45\x05\x24\x8d\xcf\xbe\x9f\x9d\x85\x43\xf1\x46\x40\xc9\xb2\xd8\x6e\xda\xa2\x91\x6b\xaa\x8f\xdb\x4d\x9f\xda\x3e\x7d\xa0\x8c\x51\xbe\x5f\x11\x22\x51\xa9\xa8\x33\xf7\xd4\x02\xeb\x46\x96\x07\x50\x28\x23\xcc\x0d\x4a\x25\xb8\x9b\x96\xe1\xd9\x08\x26\x01\x08\xa1\x9a\x0a\x0e\x6c\x03\x1a\xd2\x4e\xbd\x60\x5b\x65\x0d\xc7\x0a\xb9\xfe\x80\x0c\x4b\x13\x8b\x9a\xdc\xdf\x44\xf1\x76\x68\xc8\x04\xdf\xab\x3b\xb1\x6a\xf4\xc1\x8c\xbe\x34\xf3\xf6\xd1\x0e\xe1\x52\x08\x86\xc0\xdb\xec\x10\xc7\xe3\x49\x82\xba\x66\x14\xc9\x5a\x34\xb5\xe0\x66\x79\x92\x01\x9e\x42\x6e\x88\x04\x77\xd0\x30\xbd\x6e\xa4\x44\x5e\x1e\xc3\x7c\x5a\x68\x60\x54\x63\x95\xe4\xb9\xeb\x22\x2e\x8d\xf9\xef\x5a\x34\x5c\x2f\x8b\x2d\x77\xec\xab\xa5\x20\x4d\xa9\xe3\xc7\x54\x05\xb3\x80\x24\x1a\xe5\x5e\x02\x27\x36\xbf\x37\x7b\x37\x92\x96\xd8\xc6\x55\x53\xd9\xe8\x1d\x3c\xaf\xaa\x36\x77\x1e\xf6\x60\x61\xd7\x38\x00\x70\x93\xf5\x9a\xee\xf4\x1a\x24\x49\x86\xb8\x0a\xe3\xfe\xbe\xf8\x91\xed\xe0\x76\x41\x94\xa3\x5b\x24\xd7\xec\x8b\xd9\x25\x94\x6b\x94\x3b\x28\x71\x82\x3c\x6e\xef\xec\x41\xe3\x13\xe4\x56\xee\x13\xb0\x06\xd3\xc1\x67\x37\xe2\xfd\xa6\x25\x42\xd2\x09\xad\x6a\x86\xe6\x91\xfa\x99\xe5\x24\xf2\xe1\xb2\x77\x93\x3d\x30\xbd\x91\xec\x6d\x73\xd4\xdd\xf8\xd1\x8c\xda\x19\x8c\x09\x5d\x1e\xb7\x64\xe1\x84\x69\x50\xdd\x0c\x70\x68\x04\xd9\xc2\xb7\x7c\x27\x06\x8a\x37\xa1\x5e\xcf\xb3\xbb\xce\xab\x58\x1d\x68\x5d\x53\xbe\x37\x8f\xa2\x7c\x1f\xbc\x50\xb0\x15\x5e\x4b\x11\x48\xa7\x5d\x80\x89\x3d\x37\x6f\xcb\x4d\xed\xb8\x20\xcb\x86\xaa\xd2\x24\x19\x4d\x55\x5d\x0b\x6e\x06\x70\x8b\xcc\xd2\x60\x56\x9b\xbf\xda\xc0\x9b\x96\xdf\xa8\x3e\x74\x6d\xf2\x73\xe4\x0d\x6f\x1a\x7c\x00\xe5\xc8\xd7\x6b\xdb\x28\x4b\xcc\xda\x77\xa7\xb9\xaf\x23\xf6\xc9\x93\x90\x5f\x77\x4c\x3c\x85\x4f\x2b\xd4\x07\x41\xc2\x67\x25\x48\x49\xcd\xb9\xe1\x3f\xec\xa8\xf7\x56\x94\x90\x39\x94\x36\x51\xd8\xb5\x51\x54\x22\xb9\xa3\x15\x2e\x0b\xf3\x6f\xb7\x4d\xc2\x53\x6f\xf1\x15\x8f\xbe\xe6\x05\x87\x51\x80\x7c\x83\xc7\x40\xe7\x0c\xe2\x22\x82\x79\x73\xa1\x96\x45\x05\xf5\x67\x65\xe1\x5f\x7e\x57\x82\xff\xff\x16\x9e\xae\x50\x29\xd8\xe3\x8c\xc6\x57\x50\x9f\x50\x61\xd9\x1e\x30\x2e\xff\x0a\xea\xa4\x76\x0f\x1e\x8f\x61\x74\x45\xbb\xe9\x2c\xdc\xb2\x1a\x64\x32\x41\x13\x5e\xa6\x51\x78\x19\xf9\x9e\xe0\xa8\x0c\xa9\x92\x2d\xc7\x17\x83\xae\x12\xaa\x59\x54\x4a\x6d\x88\x3b\xb4\x6f\xf5\xe8\xae\x77\xc7\x68\x6e\x3b\x78\xa7\x68\x1f\x1e\x28\x33\x86\x75\xbe\x30\xa7\x82\x11\x76\x68\x21\x22\x58\xef\x95\xab\x1a\xe8\x9e\xb7\xa7\x75\xb8\x79\x7a\x6f\xb4\x34\x3f\xfd\x10\x83\x07\x64\xf6\x69\x11\x87\xdc\xf8\xbb\x60\xce\x69\xb4\x4b\x9f\x6d\x4d\x95\xa7\x59\xd1\xea\x2a\x21\xf5\x3b\x49\xcc\x76\x36\x7f\xd6\x40\x8d\x1f\x96\xde\x22\xd3\xf4\x5c\xe8\xcf\x03\x67\xcc\x82\xe5\xb4\x4f\xf2\xe9\xfd\xac\xbe\x4f\x7f\xd6\x28\x39\xb0\x5b\xdc\xa1\xf1\x8e\x11\xa5\x2a\x90\x5f\x51\xd7\x0c\x4a\x5c\x27\x8a\xf6\x08\x92\x02\xd7\x57\x16\x73\x93\xc7\xb8\x2a\xaf\xa1\x8a\x02\x4a\x34\xb2\xc4\xd8\xfb\xfe\xa1\x8f\x9e\xc9\x9c\x16\x9f\x10\x61\xbd\x49\x82\x99\xa9\x77\x9d\xe2\xaf\xe2\x4e\x63\xb8\x5b\xde\x76\x14\x94\xef\x19\x5a\x96\x64\xcf\xe5\xf3\x18\x35\x68\x64\xa5\x78\x9a\x4a\xd3\x41\x06\x73\x5c\x38\xd4\x1d\x3c\xa3\x9f\xc3\xfe\x1e\xda\xb3\xad\x26\x39\x6a\x3c\x82\xf6\x38\x9e\x67\xfb\x8e\x4a\xa5\xb9\x5d\xd0\x41\x0c\x83\x2c\x24\xe4\x16\x25\x84\xe1\x75\x82\x0a\x2c\x67\xab\x72\xa3\xf5\x28\x60\x8d\x76\x67\xe2\x20\x46\x4b\xc4\xcc\xd0\x52\xcc\xb5\x1c\xab\xf9\xc4\x37\x37\x6f\x6f\x29\x4f\x19\x57\x8a\xaa\x06\x7e\x4c\xba\x0b\x74\x8a\xea\x14\x10\x61\x6a\xa1\x74\xaf\x64\x83\x55\x83\x9e\x9a\x21\x89\x7b\xea\x69\x62\xbe\x1e\x23\xaf\x72\xa2\xe6\x16\x93\x24\x0a\x56\x0c\x19\xd6\x07\xc1\xc7\xd8\x81\x15\x50\x36\x52\x73\x96\xa8\xed\xab\xbf\xe3\xe9\xf4\xa9\x5b\x5b\xb8\x39\xfa\x35\x50\x16\x23\x6f\xc2\x68\xa7\x85\x54\x69\xca\xf7\xeb\x46\x69\x51\xa1\xcc\x5c\x17\xbc\xcc\x40\xf2\xe5\xe6\x90\x91\xfe\x8e\x0c\xb3\xaf\xac\x7b\xf1\x00\x8d\xef\x76\x97\x54\xea\x43\xa4\xaf\xa0\x54\x2d\x64\xfb\x7e\x2e\x8f\xf9\xe0\x75\x53\x3d\xc4\x7e\x92\x43\xcb\x63\x4b\xc3\xd1\x89\x0f\x05\xd1\x15\x64\xa5\xa6\xb4\x63\x5b\x69\x2d\xe9\x43\xa3\xd1\x73\x6c\x12\x15\xca\x47\x24\xf6\xe4\xf3\x45\x3e\xdb\x43\x7f\x0f\x31\x68\x9e\x87\xdc\x4e\x5e\x07\x43\x97\x96\xed\xf2\x74\x85\x92\xed\x73\xcc\x8b\x74\x77\x09\x83\xc5\xf6\x66\x22\x2b\xe2\x86\xfc\x9c\xf2\xfd\xdc\x57\xfa\x4f\xc0\x28\xb1\x8b\x75\x8b\xaa\x61\x9d\x05\x3a\x80\x32\x38\xc1\x5f\x4a\x29\x4e\x9a\x15\x39\xcb\x1e\xe0\x4c\xf7\x1b\x8c\x28\x72\x41\xad\x73\x31\x89\x95\xb7\x23\x43\xaf\x76\x6f\xcc\xc3\xa9\x0e\x9b\xb0\x35\x6c\x17\x99\x82\x33\xe0\xe2\xdb\x99\x25\x8c\xe9\xcc\xb0\x21\xd9\xf1\x36\x8a\x43\x85\x5e\xe4\x67\xc6\x5a\x0e\x47\xb1\xd3\x9b\x97\x30\xbf\x3b\xc6\x47\xc2\x13\xde\xf2\x4c\x08\xce\x7d\xe8\xda\x5f\x09\xd9\xf1\x73\x27\x64\xd5\x8a\x43\xfb\x37\xd5\xcc\xaa\x44\x61\x8f\xd7\x6e\x5a\x32\x37\x08\xf7\x06\xea\x2d\x77\x7f\x93\x50\x4b\x51\xa2\x52\x9e\xb5\x1c\x28\xb8\xbb\x8b\x73\x97\x3a\x01\x5f\xfe\xe1\xae\x33\x09\xba\xc9\xa2\xc8\x88\x65\x41\xf2\x0e\xd0\xb6\x7a\xd5\x03\x7a\xc3\xc6\x51\x02\x1b\x6b\xe3\x51\x70\xa8\x98\x8e\x78\x85\x7d\xcb\xce\xb3\xdf\xdd\x9b\xbe\x6d\x5f\x0c\xc6\xd9\x10\xd5\xfa\x43\x99\xfb\xf9\x08\xfd\xf0\x5c\xfa\x79\xfa\xfb\x9f\x51\x9b\x34\x6a\x9d\x3d\xfb\x65\x39\x8d\xf9\x75\x08\xf3\x6f\xb6\x72\xd6\xc6\x79\x4a\x9a\xc3\xcc\xb0\x72\x67\x94\xd7\x8d\x3e\x91\x3b\xe5\xf5\xd6\x02\xe6\x10\xfb\x27\xf2\x7a\x06\xad\x67\xb0\x7a\x06\xa9\x67\x70\x7a\x06\xa5\x67\x30\x7a\x06\xa1\x67\xf0\x79\x06\x9d\x67\xb0\x79\x06\x99\x67\x70\x79\x06\x95\x67\x30\x79\x2e\x91\xf1\x59\x23\x27\xd6\x5d\x16\xef\x9b\xd3\x57\x81\x40\xb6\x97\x03\x5f\x79\x93\x0c\x57\x8e\x35\x71\x92\x15\x21\x77\xc2\xb4\x58\x24\x97\x2f\xdb\xcd\xf9\x8b\xd3\x15\xc9\x8b\x62\xf2\xa3\x55\x54\x41\xd0\xcf\x06\x19\x6a\xf4\x2f\x3b\xa7\x3f\x82\x4d\xe7\x33\xe6\x6f\xd1\xd9\x3c\x5b\xef\xdf\x4a\xfa\xb1\x36\x6f\x3e\x26\xe9\x7b\x7d\x9c\x91\xd7\x9b\x9e\xb9\x5d\xd8\xb3\xb6\xed\x27\x54\xa6\x05\x9c\x14\x6a\x39\x29\x5f\xf1\xc7\xa5\x0c\x6e\xb8\xe3\xc8\xb8\x2d\xe2\xcf\x6f\x2f\xe2\xef\x04\x49\x6f\x59\xeb\x97\xeb\x70\x45\xbc\x2f\xcb\x0b\xff\xb6\x74\x6a\x69\xce\xbe\xff\x19\x00\x00\xff\xff\xfa\xaa\xd8\x45\xc5\x20\x00\x00")

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
