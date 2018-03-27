// Code generated by go-bindata.
// sources:
// internal/app/gateway/assets/banner.txt
// DO NOT EDIT!

package assets

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

var _bannerTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x95\x3d\xb2\xe3\x30\x08\xc7\x7b\x9d\x82\x8e\xbc\xd1\xc2\xbf\xe7\x2c\x99\xa1\xd9\x5e\x85\x5a\x1f\x7e\x07\x7d\x24\x8e\x23\x25\xfb\xb0\xc7\x63\x23\xfd\x00\x21\x90\x69\x08\x33\xb3\x1a\xc0\x69\x2a\xd0\x45\x42\xa1\xf1\x66\x16\x37\x90\x54\x55\x6c\x48\xa8\x24\xd9\x98\x1c\x9f\x66\x22\xd0\xc4\x22\xd2\x67\x01\x1a\x06\x61\x49\xfa\xa4\x18\x10\x1e\xba\x6e\x6a\xcc\xa1\xae\xa4\xaf\xc2\x9c\xc4\x20\xca\xcc\xf4\x2b\x50\x61\x9c\xd4\xc0\x67\x25\xe4\x6c\xda\x46\x70\x17\x8f\x06\x4b\xc4\x72\x0a\x94\xf8\x05\x8c\x41\x2c\xd0\x06\x12\xa9\x19\x3f\xa3\x90\x37\xeb\xd8\x81\xc4\x30\x6e\x1b\xc4\x2c\x57\x70\x61\x2b\xb6\x4e\xd3\x88\xc8\x24\xd2\x2d\x22\xb0\xb7\xc8\xd8\x22\x7f\xaa\x3a\xd3\x28\x7d\xb7\xaf\xc6\x5e\x92\x15\x55\x92\x9f\xd2\xeb\x02\x6f\x18\xe5\x5a\x1f\x9c\x5a\xce\x25\x2f\xc4\x94\xdf\x49\x0e\x74\x2c\xad\x64\x60\x05\x86\x67\x59\xb0\xa4\xa5\x16\x8b\xa8\x4b\xc9\xa2\xc5\x2e\xb4\x09\x32\x3a\xbd\x80\x49\x6b\xcd\xc2\xb9\x94\x98\x29\x72\x0d\x56\x33\x24\xef\x69\xb2\x5a\x73\x83\x97\x6b\x8d\x4d\x18\x39\xd3\x25\x4e\xa8\xa5\xec\xf0\x73\xca\xd7\xde\x89\x4b\xfd\x8e\x63\xeb\x3d\x56\xbf\xc7\x31\x71\xdb\xd0\x84\x5a\x97\x74\xf4\xfc\xb4\xb1\x85\x79\xe5\x1b\xcf\xcd\xc7\x76\xd5\x23\xef\xb2\xad\x13\xac\x2b\xe5\xe1\x37\x5b\x74\xee\x6f\xc1\xd8\xac\xac\xad\x21\xe4\x14\x67\x6b\x88\x5d\x82\x47\x85\xcd\x4e\x64\x56\x31\x8c\x72\x8c\x1e\xdc\x63\xb1\xb5\xad\x82\x54\x5b\xc3\xea\x2c\x25\x53\x8d\x83\x65\x4f\x46\x23\x66\xe9\x27\x7c\x23\x31\xd2\x1c\xdd\xa9\xf8\x44\x5a\xad\xe5\x7c\x2a\xd9\xdf\xb1\x42\x61\xfd\xb4\xc6\x56\x48\xf9\x74\x68\x70\x56\xf4\x3a\x40\xc6\xc7\x60\xa9\x8c\x68\x4f\x22\xf3\xe7\x22\xca\x91\x80\x4f\x38\x9a\xe3\x58\x6a\x3f\x2e\x23\x53\xaa\x81\x8a\xa8\x69\xfc\x78\x54\x78\xca\x27\x4b\x8c\xf0\x75\x15\x5b\xe8\xba\x7c\x09\x6c\xe6\xbf\xf1\xb2\x32\xfd\x50\x4b\x3f\x90\x39\xed\xed\xf9\x78\xfa\x55\xbf\x45\x0e\x3a\xda\xf3\xe6\x3f\x2f\xfa\x9b\xff\xa4\x69\x2c\xee\xb0\x18\x2f\xf1\x7a\xd0\xe1\x1e\x64\xff\x9e\x23\xcd\xbd\x7b\xf7\x75\x10\x3b\x31\x39\xdd\x09\xe4\x1c\x0f\x6f\xba\xfb\xd1\x7c\xf2\x63\xa4\x7d\x23\x66\xa6\x19\xd0\xbc\x6e\x7e\xd0\xdd\x63\xe6\x71\xb9\x6e\x2d\x88\x63\x04\x7e\x74\x72\xea\xfc\xb8\xbb\xff\xf1\xc3\xdd\xd1\xbf\xfb\x35\x47\x82\x69\x6f\xbe\xf8\x3f\xed\x73\xeb\x18\x8e\xfe\x4b\x9a\xf7\xf4\x2f\x00\x00\xff\xff\x5f\x84\x2d\x28\x47\x09\x00\x00")

func bannerTxtBytes() ([]byte, error) {
	return bindataRead(
		_bannerTxt,
		"banner.txt",
	)
}

func bannerTxt() (*asset, error) {
	bytes, err := bannerTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "banner.txt", size: 2375, mode: os.FileMode(420), modTime: time.Unix(1522169541, 0)}
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
	"banner.txt": bannerTxt,
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
	"banner.txt": &bintree{bannerTxt, map[string]*bintree{}},
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
