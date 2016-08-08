// Code generated by go-bindata.
// sources:
// numbers.png
// DO NOT EDIT!

package main

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

var _numbersPng = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x02\xd2\x01\x40\xcc\xc1\xc1\x06\x24\xf3\xce\x54\x98\x03\x29\xc6\xe2\x20\x77\x27\x86\x75\xe7\x64\x5e\x02\x39\x2c\xe9\x8e\xbe\x8e\x0c\x0c\x1b\xfb\xb9\xff\x24\xb2\x02\xf9\x9c\x05\x1e\x91\xc5\x0c\x0c\x7c\x87\x41\x98\xf1\x78\xfe\x8a\x14\xa0\xa0\x64\x89\x6b\x44\x49\x70\x7e\x5a\x49\x79\x62\x51\x2a\x43\x41\x62\x66\x5e\x89\x5e\x5e\x6a\x89\x82\x89\x9e\x81\x9e\xa1\xc1\x5a\x2e\xae\x03\x40\x45\x27\x3c\x5d\x1c\x43\x3c\xbc\xd7\xf6\x32\xf2\x1d\x36\xe0\x60\xfe\xfc\xdf\xbf\x9c\xa9\xf7\xd4\x67\x86\xdd\xb7\xfc\x94\xbc\x2c\x9c\xcf\x7d\x7d\x36\xe7\x73\xe0\x99\x2b\xfa\x05\xdb\x4e\x6b\xca\xfb\xc8\xfd\x8e\x38\x21\x1c\xde\x9a\xa3\x37\xa3\xe6\xde\x32\xc6\x65\xaf\x59\xb5\x5e\xbf\xb9\xf5\xe9\xae\xcd\xa3\x37\xe9\xdb\xfe\x33\x5d\x7f\xbe\x52\x4c\x24\x76\xfd\xe1\x93\x89\x6d\x9f\x5f\x9d\xb5\xf8\xaf\x78\xae\xee\x4d\xae\x9d\xff\xd6\x05\x2f\xaa\x7a\x4e\xb5\x3c\xef\xf9\xfc\xcc\x69\x4a\xba\xd5\xfa\x67\x37\xc3\xd5\xce\x57\x29\xe4\xf1\x07\xca\x1d\xcc\xb9\xb3\xfb\x96\xc3\xfd\xd4\xe5\x95\xe6\xd6\x67\xc5\x9f\x7e\xfe\x73\x3c\xe8\xab\xe2\xe7\xd7\x7a\x67\x4a\x2a\xb7\x8b\xdb\x9c\x14\x7b\xfa\xb9\x44\xdc\xc7\xfb\xee\xf9\x09\xb9\x6e\x6f\xb9\x0d\x1e\xbf\x9b\xff\x6d\xd5\x9b\x67\xf9\xc2\xbc\x3a\xc7\x4a\x23\x4e\x3c\x93\x30\x3d\x6d\x57\xf0\xa7\xcd\xf8\xff\xdb\xf5\xaf\x59\x57\x1d\xbb\xd6\xe3\x57\x26\x6f\x09\xf4\x13\x83\xa7\xab\x9f\xcb\x3a\xa7\x84\x26\x40\x00\x00\x00\xff\xff\x2c\xc5\x37\xe3\x58\x01\x00\x00")

func numbersPngBytes() ([]byte, error) {
	return bindataRead(
		_numbersPng,
		"numbers.png",
	)
}

func numbersPng() (*asset, error) {
	bytes, err := numbersPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "numbers.png", size: 344, mode: os.FileMode(438), modTime: time.Unix(1470656180, 0)}
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
	"numbers.png": numbersPng,
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
	"numbers.png": &bintree{numbersPng, map[string]*bintree{}},
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

