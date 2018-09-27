// Code generated by go-bindata.
// sources:
// assets/router/cluster-role-binding.yaml
// assets/router/cluster-role.yaml
// assets/router/daemonset.yaml
// assets/router/namespace.yaml
// assets/router/service-account.yaml
// assets/router/service-cloud.yaml
// DO NOT EDIT!

package manifests

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

var _assetsRouterClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\x31\x4f\xc4\x30\x0c\x85\xf7\xfc\x0a\xeb\x98\x5b\xc4\x86\xb2\x01\x3b\xc3\x21\xb1\xfb\x52\xdf\xd5\x5c\x1b\x57\xb6\x73\x12\xfc\x7a\x14\x25\x62\x80\xa1\x62\xb4\xf4\xfc\xde\xf7\xdd\xc1\x33\xe7\xc9\xc0\x67\x02\x95\xe2\xa4\xa0\xb2\x10\xb8\x00\xbb\xc1\x1b\xe9\x8d\x13\xc1\x53\x4a\x52\xb2\x8f\xe1\xca\x79\x8a\xf0\xb2\x14\x73\xd2\xa3\x2c\x54\xdf\x39\x5f\x02\x6e\xfc\x4e\x6a\x2c\x39\x82\x9e\x30\x8d\x58\x7c\x16\xe5\x2f\x74\x96\x3c\x5e\x1f\x6d\x64\xb9\xbf\x3d\x84\x95\x1c\x27\x74\x8c\x01\x20\xe3\x4a\x11\x52\x6b\x1b\x38\x5f\x94\xcc\x62\xe3\x08\x56\x4e\x1f\x94\xdc\x62\x18\xa0\xcd\x76\x9a\x0e\xf3\xf3\xdf\xf3\xed\xb4\x0d\x13\x45\x90\x8d\xb2\xcd\x7c\xf6\xe1\x57\xfb\xd0\xd3\xd5\xf2\x48\xe7\x4a\xf1\xc7\x69\x8f\xec\x7f\x4b\xc5\x48\x5f\x6b\xbe\x8a\x1c\xec\xd3\x9c\xd6\x68\x4d\x05\x9b\x4a\xdc\x2b\xe9\xcb\x87\xf0\x1d\x00\x00\xff\xff\x23\x98\xf9\xf1\xb0\x01\x00\x00")

func assetsRouterClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterClusterRoleBindingYaml,
		"assets/router/cluster-role-binding.yaml",
	)
}

func assetsRouterClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsRouterClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/cluster-role-binding.yaml", size: 432, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRouterClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x22\x37\xdb\xc1\x6d\x07\xad\x37\xdc\x1e\x1c\xba\xd3\x32\x53\xb3\x51\x44\x81\xa4\x1c\xa0\xbf\xbe\xb0\x9d\xa0\x41\x8b\x0e\x29\xba\xe9\x41\x78\x8f\xdf\x23\x7f\xc1\xdf\xdc\xcc\x49\xc1\x92\x54\x1a\x41\x25\x13\x1c\x45\x41\xa5\x39\xa9\xf5\xf0\x7f\x62\x03\x9b\xa4\xe5\x11\x06\x02\x34\x50\x32\x57\x4e\xce\xf3\x2a\xab\x98\xf1\x90\xa9\x0f\x27\x2e\x63\xbc\x25\x1e\x24\x53\xc0\xca\x4f\xa4\xc6\x52\x22\xe8\x80\xa9\xc7\xe6\x93\x28\xbf\xa2\xb3\x94\xfe\xf4\xc7\x7a\x96\xfd\xfc\x3b\x9c\xc9\x71\x44\xc7\x18\x00\x0a\x9e\x29\x42\xda\x62\x3a\x2e\xcf\x4a\x66\x71\x03\x0a\xda\x32\x59\x0c\x1d\x60\xe5\x7f\x2a\xad\xda\x62\xe9\x60\xb7\x0b\x00\xe8\xae\x3c\x34\xa7\xc3\x0d\x51\x8a\x45\x28\x2d\xe7\x00\x0b\xb7\x34\x4d\x74\x75\x50\x19\xab\x70\x71\x5b\xd5\x32\xd4\x2a\x26\xda\xa4\x91\xce\xbc\x89\x99\x74\xb8\x5a\x32\x9b\xaf\x8f\x0b\x7a\x9a\x3e\x43\x2c\xed\xa8\x38\xa7\xfb\x7a\x8f\x72\xb9\x9c\xa8\x28\xcd\x4c\x97\x2b\x4b\x1b\x5e\x28\x39\xa6\x44\x66\xef\x1f\x77\x5c\x49\x09\x9d\xbe\x58\x4a\xb7\xdd\xb2\x97\x4a\xc5\x26\x3e\xfa\x37\x98\xd6\x84\x07\x97\xf1\xc3\xc3\xf7\xe6\xe8\xed\x03\x43\xab\xe3\x52\xfc\x2d\x00\x00\xff\xff\x14\x2b\x2f\xcf\xc7\x02\x00\x00")

func assetsRouterClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterClusterRoleYaml,
		"assets/router/cluster-role.yaml",
	)
}

func assetsRouterClusterRoleYaml() (*asset, error) {
	bytes, err := assetsRouterClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/cluster-role.yaml", size: 711, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRouterDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x93\x41\x6f\xdb\x30\x0c\x85\xef\xf9\x15\x84\x73\x4e\xd2\xa0\xc5\xb0\xe9\x16\xa4\xc1\x50\x60\x6d\x8d\x38\xeb\xb5\x60\x65\x26\x26\x2a\x4b\x02\x45\x67\xeb\x7e\xfd\xa0\x36\x5d\xed\x20\x2b\xba\xd3\x78\x32\xc4\xf7\x3e\x3e\x53\xd0\x18\x2e\x91\xda\xe0\x2b\x52\xf8\xc1\xda\x40\x4d\x5b\xec\x9c\xc2\x1e\x5d\x47\x69\x34\x86\xa5\xeb\x92\x92\xc0\x95\xdf\x09\xa5\x04\x29\x92\xe5\x2d\xdb\x83\x02\x50\x08\x30\x46\xc7\x54\x03\x2a\x48\xe7\x95\x5b\x9a\x8e\x1e\xd9\xd7\xe6\x8d\x3e\xc2\xc8\x77\x24\x89\x83\x37\x59\x9f\x66\xfb\xf9\xa8\x25\xc5\x1a\x15\xcd\x08\x60\x0c\x37\xd8\x12\x70\x82\x44\x3a\x20\x01\x78\x6c\x29\x45\xb4\x64\x20\x44\xf2\xa9\xe1\xad\x4e\xec\x4b\xb0\x09\xbf\x04\x9b\x48\xe8\x94\x64\x04\xe0\xf0\x81\x5c\xca\x4c\xc8\x93\x0c\x1c\x3a\x39\x79\x3e\x4d\xe4\xc8\x6a\x90\x17\x45\x8b\x6a\x9b\x6f\x3d\xcb\xd0\x04\xa0\xd4\x46\x87\x4a\x07\x79\x2f\x72\x2e\x37\x70\x1e\x7b\x01\x5e\x87\x3e\x7f\x93\xec\xd9\xd2\xc2\xda\xd0\x79\xcd\xbf\x3b\x90\x02\xf8\x50\x53\x35\x48\x97\xab\xc8\xc7\x13\x09\x8e\xa6\x8f\xdd\x03\x89\x27\xa5\x34\xe5\x30\x63\xbf\x15\x2c\x0c\x14\x2a\x1d\x15\x07\xbd\x0d\x5e\x91\x3d\x49\x2f\xd4\xe4\x79\x83\x47\xc3\x72\x71\x8b\x3b\x32\x50\xd4\xc1\x3e\x92\x64\xe6\x9f\xfd\xce\x82\xf0\x8e\xfd\xa4\xc1\x28\xe1\xe7\xd3\x61\xbd\x66\x7f\x3e\x9d\xcf\xa7\x67\xc5\x31\xa4\xec\x9c\x2b\x83\x63\xfb\x64\xe0\x6a\x7b\x13\xb4\x14\x4a\xe4\xb5\xa7\x8b\x41\xb4\x17\xea\x2d\x56\xa3\x1a\x7b\xc7\xbd\x7f\x28\x83\xa8\x81\xcf\x67\x83\x6e\x94\xa0\xc1\x06\x67\x60\xb3\x2c\xff\x82\x4b\xef\xf1\x2e\x2e\xce\xff\x09\x98\x14\xf5\x5d\xe0\xfc\xcb\xf9\xa7\x0f\x11\xc7\x70\x4d\xb2\x3b\x7a\x29\x6f\x6d\xf2\xfb\x53\xfb\xa9\x36\x8b\x4d\x75\x5f\xde\xae\x37\x83\x21\xcf\x0f\xd0\x40\x91\xa7\x17\x27\x6c\xeb\xdb\xef\x9b\xd5\xfa\xbe\x5a\xad\xef\xae\x96\xab\xfb\x9b\xc5\xf5\xaa\x2a\x17\xcb\xd5\x29\xc8\x07\x9e\xd5\x6b\x39\xde\x93\xa7\x94\x4a\x09\x0f\x64\x06\x30\xf6\xac\x8c\xee\x92\x1c\x3e\x55\x64\x83\xaf\x93\x81\xf9\xf0\xf6\xf2\xed\x7c\x25\x1d\x1a\x01\x22\x6a\x63\x60\xd6\x10\x3a\x6d\x7e\x1d\x37\x4f\x6d\x59\x08\x6b\xfe\x7f\x41\x7e\x07\x00\x00\xff\xff\x7a\x7b\xfb\x95\x3a\x05\x00\x00")

func assetsRouterDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterDaemonsetYaml,
		"assets/router/daemonset.yaml",
	)
}

func assetsRouterDaemonsetYaml() (*asset, error) {
	bytes, err := assetsRouterDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/daemonset.yaml", size: 1338, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRouterNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcb\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x58\xd7\x07\x44\x9b\x21\x28\xe9\xad\xe4\x03\x16\x17\x3b\xb2\x7d\xcc\x8f\xa8\x6e\x80\xf7\x11\x1d\x8d\xee\x3c\x11\x8b\x3b\x0a\x2f\x79\xc0\x43\x4c\x1b\x7d\x6f\x65\x22\x79\x70\x72\x2b\x44\xca\x13\x8d\x6c\x41\xe3\x2d\xcf\xac\x7d\x3f\x22\xe1\x55\xf4\xe5\x88\xa8\x6e\x47\xc2\x0b\x11\xab\x5a\x72\x8a\x69\xfc\x1d\x9d\xe6\x22\x76\x55\x1b\xa8\x81\x1d\x3d\xcd\x1b\x6d\x5b\xf9\x05\x00\x00\xff\xff\x68\xe9\x4e\xf7\x84\x00\x00\x00")

func assetsRouterNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterNamespaceYaml,
		"assets/router/namespace.yaml",
	)
}

func assetsRouterNamespaceYaml() (*asset, error) {
	bytes, err := assetsRouterNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/namespace.yaml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRouterServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\xb1\x6a\x04\x31\x0c\x84\xe1\xde\x4f\x31\x70\xf5\x1e\xa4\x75\x97\x32\x6d\x02\xe9\x1d\xef\xdc\x9d\xc8\xae\xe5\x48\xf2\x86\xbc\x7d\xd8\xc3\xa5\x40\xf3\xf3\x5d\xf0\x5a\xab\x8e\x16\xb8\xa9\xc1\x74\x04\xcd\x51\x8d\x25\xb8\xe2\xeb\x0f\xf1\x20\xb4\xd3\x4a\xa8\x5d\xf1\x16\xf8\x95\x6d\x83\xf1\x67\x88\x11\x75\x1b\x1e\x34\x78\xd5\xce\x35\x5d\xd0\x69\xbb\xb8\x8b\x36\x87\x71\x7b\x56\x42\xf1\x7e\x86\xd1\x4d\x2b\xdd\xa5\xdd\xaf\xe9\x5b\xda\x9a\xf1\x41\x3b\xa4\x72\x1a\x52\xe9\xf2\x49\x3b\xd7\x19\xc7\x4b\xda\x19\x65\x2d\x51\x72\x02\x5a\xd9\x99\x27\x70\x9e\xde\x4b\x65\x3e\x75\xcd\x1f\x72\x8b\x65\x6a\x16\x69\x77\xa3\xfb\x32\xbf\xff\x03\x00\x00\xff\xff\x46\x51\x70\x49\xe4\x00\x00\x00")

func assetsRouterServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceAccountYaml,
		"assets/router/service-account.yaml",
	)
}

func assetsRouterServiceAccountYaml() (*asset, error) {
	bytes, err := assetsRouterServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-account.yaml", size: 228, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRouterServiceCloudYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\x31\x8b\xdc\x40\x0c\x85\x7b\xff\x8a\x07\x5b\xef\x91\x74\xc1\xe5\x5d\x15\x38\xc2\x42\x96\xf4\xba\xb1\xbc\x33\xdc\x58\x1a\x24\x79\x93\xfd\xf7\x61\xc6\x77\x45\x48\xe9\x87\x25\x7d\xef\x9b\x13\x5e\x95\x16\x3c\x53\x25\x49\x6c\xf8\xc9\x76\x2f\x89\x11\x8a\x56\x29\x31\x8a\x60\x35\x95\x80\xae\x88\xcc\x30\xdd\x83\xad\xc7\xa9\xea\xbe\x80\xe5\x5e\x4c\x65\x63\x09\x7f\x9a\x4e\x78\xa9\xbb\xf7\x1f\xbe\xcb\xcd\xd8\x1d\xde\x38\x95\xb5\x24\xdc\xa9\xee\xec\x20\x63\x50\x6b\xb5\xf0\x02\x0a\xd8\x2e\x51\x36\x7e\x9a\xde\x8b\x2c\xf3\xe7\xf9\x89\x5a\xf9\xc5\xe6\x45\x65\xc6\xfd\xeb\xb4\x71\xd0\x42\x41\xf3\x04\x9c\xf0\x83\x36\x46\x71\x38\xc7\x3f\x2b\x00\xa1\x8d\xbd\x51\xe2\x19\xda\x58\x3c\x97\x35\xce\xe9\x20\x3a\x97\x83\xe8\x7c\x34\x98\x80\x4a\x6f\x5c\xbd\xef\x44\x47\x9a\x3f\xba\x4d\x1d\xb9\xa7\xf1\x68\x3c\x0f\x3f\x9f\x7a\x26\xc0\xb9\x72\x0a\xb5\xff\xc7\x3a\xda\x35\x17\x07\x55\x57\x64\xf2\xe1\x8b\xd7\x95\xd3\xb0\xb7\x91\xbd\x17\xb9\xe1\xf5\x19\x4d\xb5\x22\xc8\x6e\x1c\x0e\x72\xec\x92\x99\x6a\xe4\x07\x7e\x67\x16\x88\x8e\x65\x1f\xaa\x9b\x2e\x87\xb6\x66\xec\xdc\x5f\x42\x40\x10\x5d\x18\x6f\x9c\x8b\x2c\xe3\x8e\x1f\xe6\xba\x05\xfe\x13\x6c\x42\xf5\x6a\xb4\xae\x25\x5d\xb4\x96\xf4\xe8\x45\x12\xd5\x09\x68\x6a\x31\x5a\x9f\x87\xaf\x19\x39\xa2\x8d\x36\xcd\x34\x34\x69\x9d\x71\x7d\xb9\x1c\x89\x5a\xcc\xf8\xf6\x65\x7c\x1c\xc0\x97\x11\x8d\x99\xbf\x01\x00\x00\xff\xff\x17\xe2\x05\xe9\x3d\x02\x00\x00")

func assetsRouterServiceCloudYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceCloudYaml,
		"assets/router/service-cloud.yaml",
	)
}

func assetsRouterServiceCloudYaml() (*asset, error) {
	bytes, err := assetsRouterServiceCloudYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-cloud.yaml", size: 573, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"assets/router/cluster-role-binding.yaml": assetsRouterClusterRoleBindingYaml,
	"assets/router/cluster-role.yaml": assetsRouterClusterRoleYaml,
	"assets/router/daemonset.yaml": assetsRouterDaemonsetYaml,
	"assets/router/namespace.yaml": assetsRouterNamespaceYaml,
	"assets/router/service-account.yaml": assetsRouterServiceAccountYaml,
	"assets/router/service-cloud.yaml": assetsRouterServiceCloudYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"router": &bintree{nil, map[string]*bintree{
			"cluster-role-binding.yaml": &bintree{assetsRouterClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml": &bintree{assetsRouterClusterRoleYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{assetsRouterDaemonsetYaml, map[string]*bintree{}},
			"namespace.yaml": &bintree{assetsRouterNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": &bintree{assetsRouterServiceAccountYaml, map[string]*bintree{}},
			"service-cloud.yaml": &bintree{assetsRouterServiceCloudYaml, map[string]*bintree{}},
		}},
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

