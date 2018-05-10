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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xcd\x6e\xdb\x30\x0c\xbe\xe7\x29\x04\xb7\xc7\xb6\xda\x61\xa7\x1c\xb7\xd3\x4e\x2d\xb0\xdd\x86\x22\x50\x6c\xda\x56\x66\x4b\x9e\xa4\x34\x08\x8a\xbc\xfb\x60\xf9\x27\x56\x22\xc9\xce\xa2\x04\x29\xd0\x1e\x7a\x20\x29\x92\x1f\xf3\x91\x94\xfc\x3e\x43\x08\xa1\xe8\x5e\xc6\x39\x94\x24\x9a\xa3\x28\x57\xaa\x9a\x63\xbc\x92\x9c\x3d\x36\xd2\x27\x2e\x32\x9c\x08\x92\xaa\xc7\x2f\x5f\x71\x23\xbb\x8b\x1e\x9a\x93\x6a\x5b\x41\x7d\x8c\x2f\x57\x10\xab\x4e\x2a\xe0\xef\x9a\x0a\x48\xa2\x39\xfa\xad\x25\x5a\x5a\x12\x99\x17\x94\x65\x8b\x36\xda\xc3\x5e\x95\x11\x05\x1b\xb2\x8d\xb4\xe4\xb5\xf5\x52\x09\x5e\x81\x50\x14\x64\x34\x47\xef\x6e\x3f\x43\xa5\x91\x94\x54\x82\xb2\x2c\xea\x95\x3b\x4b\x44\xe7\x61\x03\x51\xaf\xb5\x22\xeb\xb5\x8c\x94\x70\x70\x42\xcb\xdf\x40\x48\xca\x99\x4d\x15\x73\x96\xd2\x6c\x2d\x88\xa2\x9c\x49\x9b\x85\x12\x34\xcb\x40\x58\x75\xf0\x06\x4c\x2d\x72\xc2\x92\xc2\x6b\x51\x50\xf6\x47\x46\x86\xf6\xf5\x00\x99\xa3\xda\x26\x36\x9b\x06\xf9\x4a\xde\xfd\xed\x3c\x55\x09\xea\x34\xa1\xb2\x2a\xc8\x76\x11\x3e\xdd\xce\x33\x2d\x49\x16\xda\x35\xc8\x58\xd0\x4a\x05\xaf\xc6\x01\xbb\x9c\xbe\xa9\x82\xd2\xad\xd6\x26\xf7\x02\xd2\x3a\xfc\x1d\x4e\x20\xa5\x8c\x6a\x8f\xd8\x08\x70\x9c\x93\x23\x2f\x03\x0f\x11\xa2\xeb\xfc\x31\x38\x07\x74\x0f\x0f\xc7\x08\x70\x25\x38\x4d\x6f\x5e\x0a\x4b\xed\xfd\xe2\x40\xfa\x11\x15\x1e\x45\xeb\x3a\x2c\x84\x99\xc7\x49\x44\x92\x44\xc7\x26\xc5\xcb\x70\x24\xa6\xa4\x90\x30\x33\x5d\xb4\x47\xa3\x41\xc2\xe6\xaa\x32\xdb\xe3\x2a\xbb\x46\xfb\xb4\xc8\x25\x28\x45\x59\x76\x73\x5b\xa0\x3d\xf3\x11\xa6\x69\x5f\xc2\x51\xaf\xd6\x1f\xb4\xb7\x72\x30\xcc\xd3\x1b\x63\xd9\xa2\x23\x56\x1f\x4b\xfe\x8f\xe7\x83\x2b\x53\xd7\x8a\x9f\x34\xfe\xa4\xf1\x07\xa6\xb1\xb9\xe3\xc3\x93\xd9\x4b\x4c\xce\xe0\x39\xb5\x1e\x76\xfc\x14\xde\x80\x03\xab\x14\x04\xb0\x18\xec\x35\x7d\x9d\xc2\x8b\xb3\x12\xd8\x6f\xc0\xc9\x19\xdc\x58\x03\x5f\xac\xdd\x2a\x22\x88\xe7\xee\x73\x6b\xcd\x66\x41\xb0\x67\x57\xe0\x8a\xf7\x9c\x39\xb3\x38\x13\x39\x4a\x14\x71\x78\xd0\x7a\x9a\xf8\xb4\xf5\xed\xd4\x4e\x6d\x47\x52\x23\x14\x36\xd3\xf2\x59\x20\x4b\x21\x9c\xc6\x8e\x2b\x31\x6a\x01\x4e\x8e\xe3\x63\xcd\x58\x9c\xe6\x22\x7f\x5a\x20\xb7\x3b\xd4\x72\x85\xac\x0b\xd5\x7d\xa1\x92\x73\x8c\x33\xaa\xf2\xf5\xf2\x29\xe6\x25\xfe\xf5\xe3\xdb\xf7\xe7\x9f\x3c\x55\x1b\x22\x00\xa7\x05\xcf\xf8\x63\xcc\x99\x12\x74\x89\x97\x05\x5f\xe2\x92\x48\x05\x02\x93\xb8\x66\x5b\x6d\xb0\x19\x3c\x32\xda\x6f\x5a\x4f\x2b\xe9\x9a\x5e\xc8\xda\x2a\x9e\x32\x4c\xdb\x4a\x76\xd7\x81\x16\x9c\x7e\xf8\x05\xdd\x6e\xbe\x0f\x51\x09\x95\x15\x51\x71\x0e\x67\x5e\xcc\xc6\x5f\x92\xe6\xdb\xce\x51\xfe\x09\xcf\xcd\x70\xe3\x71\x80\xfd\x6a\x59\x7b\xe7\x21\x9a\x3c\x13\xb5\xa5\xf7\x73\x07\x72\x4f\x38\x74\xc2\x94\xd3\xb6\x74\x7c\x2c\xa0\x53\x66\x10\xf2\xcf\x21\x34\x04\x77\xf5\xc0\x94\x55\x6b\xf5\xe2\xdf\xfd\x96\xe0\xe3\xbf\x6c\x7f\xe0\xe4\x2b\x81\x2d\xdc\x14\xa8\xc8\x39\xfe\xc6\xb5\x8e\xb1\x79\x91\x19\x38\x6b\xfe\xef\x66\xff\x02\x00\x00\xff\xff\x96\x2a\x42\x06\xd7\x18\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 6359, mode: os.FileMode(420), modTime: time.Unix(1525272079, 0)}
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