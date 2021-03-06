// Code generated by go-bindata.
// sources:
// static/index.html
// DO NOT EDIT!

package static

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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x6d\x6f\xdb\xc8\x11\xfe\xee\x5f\x31\x67\xb4\xc7\xd5\x85\xa1\x14\x17\x49\xaf\x96\xe8\x02\x75\x84\x22\xad\x1d\x19\x96\x81\x44\x28\x8a\xbb\x35\x39\x12\x79\xa6\x76\x89\xdd\x95\x65\x41\xd1\x7f\x2f\x76\x97\xa4\x96\x6f\x8a\x51\xe4\x8b\x49\xee\xcc\xce\x3c\xf3\x3e\xd6\xe4\xa7\x8f\xb3\xeb\x87\xc5\xdd\x14\x12\xb5\xce\xae\xce\x26\x09\xd2\xf8\xea\x0c\x60\xb2\x46\x45\x21\x4a\xa8\x90\xa8\xc2\xf3\x8d\x5a\xbe\xfd\xf5\xdc\x10\xa4\xda\x65\x08\x6b\x8c\x53\x1a\x9e\xcb\x48\x20\x32\x43\x00\x88\x28\x7b\xa6\x12\xf6\xe6\x03\x60\x9b\xc6\x2a\xb9\x84\xf7\xa3\x51\xfe\x32\x2e\xce\x12\x4c\x57\x89\x6a\x1c\xae\xa9\x58\xa5\x4c\x1f\xe6\x2f\x40\x37\x8a\x97\x84\x38\x95\x79\x46\x77\x97\xf0\x98\xf1\xe8\xc9\x9e\x1e\x34\x86\xa1\x01\x71\x75\x36\x19\x5a\xbc\x93\x47\x1e\xef\x0c\xba\x02\x43\x1a\x87\xe7\xf6\xf5\xdc\xe2\x08\xcf\xdf\x8f\x46\xe7\x85\x7e\xfb\x71\x35\x19\x5a\x16\x6b\x56\x24\xd2\x5c\x81\xda\xe5\x18\x9e\x2b\x7c\x51\xc3\x3f\xe8\x33\xb5\xa7\xa5\x7d\x19\x95\x12\xae\x53\x11\x65\x58\x59\x19\x71\x26\x95\xd8\x44\x8a\x0b\x12\x21\x53\x28\xbe\xfa\x60\x5f\x16\x3e\x08\x1a\xa7\x1b\x39\xa8\xb8\x01\x54\x92\xca\xa0\x60\x84\xb0\xe0\xfc\x3a\xee\xa2\x2f\x2a\xfa\xa2\x41\xb7\x62\x21\x2c\xe4\x97\xd4\xc3\x59\xe9\x38\x41\xb7\x24\xe2\x4c\xdb\xe1\x2a\x2f\x8e\x82\x47\x5c\xa5\xec\x8e\xaa\x84\x0c\xc6\x2d\x22\x15\x11\x71\x51\xfa\x35\x4c\xbe\x8b\xc0\x87\x91\x0f\x17\xf0\x0b\xdc\x52\x95\x04\x77\x9f\x7c\x58\xd2\x4c\x62\x87\xd0\x65\x9a\x65\x73\x93\x39\x21\x78\x2b\x9d\x34\x5e\x37\xd3\x11\xd1\xe1\xcc\xb1\xc9\xfa\xfe\x26\x65\xdd\x9e\x97\x8a\x0a\xf5\xd5\x07\xf3\x5c\xf8\x80\x2c\xfe\x6a\xfe\x2e\x5a\xbe\xb7\xac\x10\x5a\xde\xa6\xe7\xad\x80\x92\xda\xf4\xbb\x16\x0b\xa1\x91\xde\xa6\x2c\x2c\x65\xf1\x63\xa3\xb1\xe6\xcf\xf8\xc0\x89\x03\xdd\x77\x91\x76\xdc\xc8\x52\x56\xdd\xb0\x7e\xa8\x10\xf6\x70\x7f\xd1\x15\x02\x21\xbc\x6b\x93\xa5\x12\xfc\x09\xab\xc8\x6d\x93\x54\x61\x47\xe4\x2c\xdb\xc9\xd8\x5d\xf3\x75\x4e\xa5\xfc\x3f\xca\xc7\x07\xca\x56\x46\xff\xa8\x15\x4d\x91\xb2\x15\x84\xc0\x70\x5b\x14\x66\x7f\x11\x36\x22\x56\xca\x34\xcf\xf1\x59\x45\xcc\x50\xc1\xfe\xc5\x87\xdd\x01\x42\xcb\xc9\x10\xe3\x0c\xef\xb8\x4c\x55\xca\x99\x36\xb2\x2e\xc9\xd2\x0b\x14\x3a\x45\x3b\x30\x8c\x7c\x18\x0d\x5a\x99\xd1\x94\xec\x58\x27\x50\x6d\x04\x73\x0e\x00\x5e\x2e\x6d\x9d\x45\x5c\x12\x52\x54\x1c\xfc\xe2\x58\x33\x80\x21\xbc\xfb\x75\x34\x28\x0f\xb5\x73\xca\x5e\xf1\xc6\x39\x2a\xe1\x39\xb2\x77\x85\x6c\x99\xb2\x1f\x21\x7b\x51\x89\x3e\xbc\xba\x1c\x8e\x32\x6a\x1c\x3f\x20\x34\x65\xdd\xbe\x8c\xfb\xa8\xba\x76\x77\x3d\x97\x1b\x68\xfa\x13\x3c\x4b\x91\xa9\xce\xcc\xde\x9b\x19\xe4\x17\xe3\xe7\xe0\xc3\x3e\xf5\x21\x3f\xf4\xce\x87\x14\x42\x48\x1b\x60\x73\x08\x21\x6f\x42\x14\x7c\x0b\xa1\x8d\xdc\x32\xe3\x5c\x10\x92\xc2\x5b\x78\xa7\x83\xf5\x21\x18\x35\x53\x3e\xe2\xd9\x66\xcd\x20\x84\x92\xeb\xcf\xf0\xa1\x29\xb1\x6a\x80\xc4\x82\x1d\x7e\x38\x86\x9c\x6f\xe1\x4d\x63\xe4\x34\x9b\x2a\x31\x96\x3a\x97\x0a\x9d\xc7\x7b\x0d\x48\xb6\x1d\x14\x15\xec\xb6\x88\xde\x96\x77\x6c\x0b\xd6\x2f\xed\xb2\x3a\x99\x61\x56\x63\x77\x54\x7b\xe6\xd5\x63\x46\xa3\xa7\xae\x79\xc5\x99\xd2\xf4\x0b\xbd\xb8\x48\x14\xe9\xb2\x67\xa8\x3d\xe0\x8b\xb2\xf6\xa4\x7e\x0d\x85\x5b\x8e\xf0\x16\xde\xf7\x52\x17\xf0\x46\x53\xdf\x8d\x4e\x65\xe0\x3f\xe9\xba\xa7\xb3\xf6\xf9\xc2\x9e\xeb\x3d\xc3\xbe\x35\x13\xc6\x6e\x53\x61\xb1\xda\xb5\x83\xae\xb4\x7f\xf6\xed\x1a\x8f\x32\xa4\xe2\xda\x5c\x22\x7d\x5a\x03\xc3\x74\x8f\x91\x22\x23\xd3\x20\x1d\x95\x41\x51\x30\xee\x91\xcd\xc7\x76\xb4\x25\xaa\xb9\x46\x42\x0c\x9e\xae\x79\x6f\x50\x9a\x67\x47\xb6\x2b\x0c\x22\x53\xb9\xda\xce\xd9\xe3\x1f\x18\xa9\xe0\x09\x77\x92\xb4\x19\x06\xc1\x9a\xe6\x84\xd8\xaf\x4f\x1f\x07\x10\x5e\xd5\x1a\x74\xd1\xb3\x4d\x2a\x1b\x1e\xe2\xe0\xf7\x3b\x34\xfe\xa7\x14\xf5\x5f\x1f\x2e\x46\x83\x63\xd3\xec\xc9\xe9\xb6\x2b\x5d\x3f\x9f\x34\x2e\x58\x72\x31\xa5\x51\x52\xc2\x6f\x81\xb7\xc7\xb6\x2e\xdc\x30\x39\xc5\x71\x18\xf4\xe7\xde\x17\x2e\xb2\xb8\x12\x68\x2a\x95\x24\x5c\x76\xe4\x5c\x99\x53\x31\x8f\x36\x6b\xad\x72\x85\x6a\x9a\xa1\x7e\xfd\xc7\xee\x53\x4c\x3c\xcb\xe1\xb5\xbb\x57\x95\xac\x36\x21\x56\xa8\xae\xed\x19\xf1\x2e\xe2\x16\xff\x4a\x57\x83\x6d\x2c\xba\x30\x6a\x46\xd5\x32\xab\xe5\xb8\xdf\x22\xbe\x5e\x53\x16\x6b\x94\xfb\x0c\x97\xea\xd2\xee\xb5\x3e\x08\xfb\xef\x8b\xf9\x3a\x34\xf4\xfd\x96\x51\xa9\x6e\x51\x4a\xba\x32\x5d\xc3\xeb\xa2\xdf\x63\x8e\x54\xdd\xce\x1b\x5d\x5b\x83\xfc\xa8\x73\x78\x30\xd0\x66\x3d\xa4\x6b\xfd\xea\xe0\xda\xa6\x2c\xe6\xdb\x80\xc6\xf1\xf4\x19\x99\xba\x49\xa5\x42\x86\x82\x78\x4f\xb8\xdb\xe4\x9e\x0f\x04\x9f\xcb\xa8\x1a\x6d\x09\x65\x71\x86\xff\xc6\x1d\x29\xa0\x5b\x7a\x7b\x43\x3f\x25\x39\xe6\x5b\x76\x52\xb6\x12\x9b\x0e\xd1\x75\xbb\xb7\xb2\x88\xc2\x17\x7c\x9c\xf3\xe8\x09\x15\xf9\x7d\x2b\x2f\x87\xc3\x3f\xed\x75\x82\x1c\x86\x5b\xf9\x7b\xc7\xa5\x80\x33\x9e\xa3\x99\x56\xf8\xdc\x91\xae\x9c\x49\x9e\x61\x90\xf1\x15\xf1\x66\x77\xd3\xcf\x9e\x2b\x03\x60\x38\x84\xb9\x4e\x42\xd8\xe4\xa0\x12\xb4\x18\x21\xe3\x3c\x07\xc5\x41\x22\x8b\x81\x2f\x97\x50\x45\x5a\x71\xc3\x56\x54\x81\x23\xc8\x96\x12\xb2\xf8\xba\x60\x75\xb7\xf4\x43\x27\xee\x28\xe3\x12\x5f\x05\xfc\xfa\x66\x36\x9f\x36\x90\x3b\x4e\xdb\x64\xd9\xf7\x74\xad\xab\x84\xfb\xbe\xb6\xfb\xe9\xfc\x6e\xf6\x79\x3e\xbd\x04\x0f\xde\x00\x3e\xab\x20\xa6\x8a\xd6\xb5\xeb\x1d\xab\x6c\x99\xff\x9a\xcf\x3e\x07\x39\x15\x12\x89\xc3\xdc\x40\xaa\xcb\x2c\x68\xf4\xe0\x6e\x26\xdb\xc2\xbe\x67\x10\x0a\xc1\xc5\xab\xcc\x99\xde\xdf\xcf\xee\x5b\xb6\x38\xd2\x8b\x97\x52\xcb\x31\x6d\x73\x81\x52\x96\x79\xeb\x28\x90\xdb\x54\x45\x49\x91\xef\x7a\x0a\x5c\xf3\xb8\x36\x51\x00\x22\x2a\x11\xfe\xf2\xd7\x4b\xe7\xa8\xd9\x35\x02\xdd\x32\xf4\xb6\xa6\xb5\x8c\x6b\x8c\x8f\x02\xe9\x53\xcd\xdf\x56\xde\xdf\x4e\xca\x33\x5d\xe7\xa4\xc0\xca\xe8\xf6\x70\x64\x31\x29\x72\xa4\xf1\xbf\xe7\xd1\x91\xf3\xe9\xe7\x07\xeb\xc7\x92\x73\xdc\x0a\x4c\x4d\x50\xc7\x0c\x76\xeb\x03\xf6\xf5\x9d\x9d\xf1\xed\xab\x3b\x9e\x7b\xcf\xc9\xed\x2e\x0f\xff\x1d\xbc\xdb\xf1\x8d\x07\x97\x2d\xb2\x75\x98\xa1\xdf\x6b\xba\x77\x3b\xf6\x6a\xed\x34\x5d\x42\x69\x0c\xfc\xfc\x33\x10\xd2\xee\xe0\x3f\x85\x95\x37\xe0\xdb\x37\x20\xd6\x88\xb7\x5d\xbd\xfc\x0a\x2e\x46\xa3\xc1\xa0\x9e\x28\x55\xdf\x68\x79\x0d\x7a\x26\x46\xc1\xd7\xcd\xe6\x0c\x0e\x83\xc4\x8d\xf9\x31\x7f\xad\x23\xf9\x46\x11\x72\x6c\xd7\xf5\xe0\xf8\xf0\xbe\x67\x89\xac\x06\x72\x7b\x16\x7c\x9c\xdd\x9a\x51\xcb\xd4\x0d\xa7\x31\xc6\x7a\x28\xd4\x8a\xd3\x6c\x9b\xb0\x35\x4b\x40\xd1\xed\xf5\xbb\xb3\x94\x2c\x51\x45\x09\xf1\x86\xba\xe7\x7b\xc7\x55\x27\x50\x09\x32\x42\x04\x4a\x23\x4f\xa0\x0c\xcc\x44\x1f\xb4\x58\xec\x36\xd1\x68\x08\x46\x63\xe0\xac\x1b\xe3\xd6\x12\x65\x9f\x93\xa1\xfd\x2d\xef\xea\x6c\x32\xb4\xbf\x16\x4e\x86\xf6\x37\xcf\xff\x05\x00\x00\xff\xff\x2f\x75\x18\x57\x04\x15\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 5380, mode: os.FileMode(420), modTime: time.Unix(1484197120, 0)}
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
	"static/index.html": staticIndexHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
	}},
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

