// Code generated by go-bindata.
// sources:
// schema.json
// DO NOT EDIT!

package schema

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xc1\x6e\xdb\x3c\x0c\xbe\xf7\x29\x04\xf5\x3f\xfd\x58\x9b\x1d\x76\xca\x79\x40\x77\x2c\x9a\xdd\x86\xa2\x50\x6d\xc6\x51\x61\x4b\x9a\xa4\x74\x0b\x86\xbe\xfb\x60\xcb\x76\xed\x5a\x72\xec\x44\x1e\x14\x40\x3a\xe4\x20\xd2\x24\x45\x91\x1f\x29\x29\x7f\xae\x10\x42\x08\xff\xa7\x92\x1d\x14\x04\xaf\x11\xde\x69\x2d\xd6\xab\xd5\x8b\xe2\xec\xc6\xcc\xde\x72\x99\xad\x52\x49\xb6\xfa\xe6\xf3\x97\x95\x99\xbb\xc6\x9f\xea\x2f\x25\x6c\xcb\xcf\xae\x57\x29\x6c\x29\xa3\x9a\x72\xa6\x56\x1b\x23\xae\xe6\xe9\x50\xf0\x1a\x19\x95\x15\xe1\x2b\x55\x82\xe8\x64\xd7\x9b\xad\x28\x12\x7e\xee\xa9\x84\x14\xaf\xd1\x8f\x1e\xa5\xa2\x32\x52\x40\x2d\xbc\xff\x15\xdf\x6b\x50\xb8\x47\x78\xec\xf3\x61\x21\xb9\x00\xa9\x29\xa8\x81\xd6\x77\xd9\x36\x4a\x45\xd5\x07\x51\x52\xb1\xd2\x92\xb2\x0c\x0f\x98\xde\xdc\x56\x39\x65\x52\x0d\x85\x9b\x8c\xce\xda\x1f\xbb\x34\xeb\x9e\x3d\x94\x66\x0e\x57\xe4\x58\x55\xcf\x1b\x44\x4a\x72\xb0\x38\xe3\x6a\x44\x08\x26\x69\x5a\x69\x26\xf9\x7d\x77\x4b\xb6\x24\x57\xf0\x81\xb5\xd1\xc3\x9f\x5f\x20\xd1\xef\x8a\x3a\x22\xf1\x1d\xd1\xf0\x8b\x1c\x7c\x85\xd2\x2b\x48\x45\x39\xb3\x91\xb4\xa4\x59\x06\x52\xd9\x68\x69\x1d\xd1\xe7\x46\x61\x0a\x2a\x91\x54\x94\xfe\xf1\x1b\x8c\x1d\x03\x43\x0f\xc8\x16\x1d\xbc\xc6\xa4\xc3\x27\x39\x39\x3c\xd1\x82\x64\x9e\x73\xbf\x11\xed\x1f\x55\xfc\x4b\x14\x3c\xa7\x89\x2b\x22\x51\x40\x81\x71\x5f\x1a\x6a\xd9\x5b\xc7\xc2\xd0\x89\x61\xa1\x40\xbe\xd2\xe4\x02\xfc\xb1\x31\x86\x2e\xee\x90\x16\xf8\x42\x77\xc8\x77\x63\xe8\xe2\x0e\x69\x8a\xc4\xe9\x69\xf8\x2f\x2b\xe4\x37\xc2\xd2\x1c\xe4\xdc\x0a\x99\x5a\x61\x78\x76\x41\x73\xb5\x7a\xd3\x7d\x65\x4d\x51\xad\x29\xcb\x46\x22\x52\x10\xad\x41\xb2\xfb\x71\xf3\x5a\xf6\xdb\xff\x47\xe9\x68\x64\x5b\xb4\xdc\x83\x3b\x52\xbb\x4b\x1c\x3a\x79\xa8\xa3\x0a\xc2\x71\x71\x15\xe3\x33\xe7\x39\x10\x5b\x9f\x32\x60\xa5\x4c\x43\x99\x15\x13\x58\xd9\xbe\x78\x9e\xca\x99\xe7\x53\xf8\xea\xc0\x9c\xc0\xe9\xda\xff\xee\x78\x74\x52\xdf\x4e\xca\xf9\x8f\x69\x63\x97\xb6\x70\x7e\xd6\x65\xcd\x53\x03\x5b\x69\x0c\xb3\x07\xf5\xdf\xba\x44\x1c\x88\x38\xd0\x1f\x9e\x71\xc0\x95\x60\x97\x51\xf9\x1f\x40\x09\xce\xd4\xd0\xdc\x23\xd8\x02\x52\x72\x79\x1e\x88\x24\xbc\x10\x39\xfc\x3e\xee\xa8\x26\x80\x27\xf9\xde\x18\xe6\x57\x26\xdd\xfa\xc5\x24\xbe\xd7\x62\xaf\x23\x22\x45\x44\x32\xe3\x82\x3b\x13\x73\x35\x38\x13\x3c\x94\x06\x71\xe6\x2d\x18\x51\x07\x96\x04\x9e\xe6\x5e\x6e\x4d\x8e\x29\x76\x28\x47\x27\x9e\x5b\x65\x5d\x0f\xc2\x3f\xc9\xb7\x95\x6b\x69\x97\x98\x68\x0d\xdd\x1d\x1b\x0d\xe2\x72\xaf\xe8\x37\x8d\x83\x66\xe1\x48\x41\xd4\x2e\xa7\x2c\x7b\x52\xdd\x47\xa5\x1e\x4b\x56\xdf\xfd\x9f\x05\x36\x99\xe3\x01\xa1\x65\xf0\xb6\xc1\x8e\xcd\xbd\xb3\xad\x02\x39\xe2\xf5\xa3\x57\x2e\xa3\x11\x6d\x6e\x2a\xe3\x19\x37\x9e\x71\xdf\x75\xc4\x8e\x72\x8c\x29\x9e\x71\xa7\x41\x4b\x59\x1a\xe7\xb6\xa8\xb6\x87\x93\xb9\x10\xe2\xbb\xa1\xa4\x2c\x1e\x1b\x63\x92\xb7\x63\xf9\x24\x8f\x17\x15\x31\xe2\xba\x63\xf9\x88\x53\x8e\x36\x70\x20\x34\x88\xca\xd2\xbc\x26\xfb\x6c\x5a\x2d\xf3\x3b\xf3\x3c\x1b\xea\x1f\x87\x5a\xf3\x42\x3f\x23\x37\xcf\xdc\x5e\x8f\xc9\xb1\xc3\x77\x2e\x31\x42\xf1\x18\x67\xe8\x50\x1c\x7a\x87\x7f\x65\x7e\xdf\xfe\x06\x00\x00\xff\xff\x42\x3e\x35\x7f\xa9\x2c\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 11433, mode: os.FileMode(420), modTime: time.Unix(1525858263, 0)}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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
