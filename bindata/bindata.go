// Code generated by go-bindata.
// sources:
// template/ccd.file.tmpl
// template/client.ovpn.tmpl
// template/dh4096.pem.tmpl
// template/iptables.tmpl
// template/server.conf.tmpl
// DO NOT EDIT!

package bindata

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

var _templateCcdFileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x6a\xc3\x30\x0c\x86\xef\x7e\x0a\x91\x9d\x23\xec\x40\xc6\xf2\x04\x63\x87\x8d\x91\x4b\xcf\x6e\x2c\x27\xa6\xe0\x04\xdb\xa1\x0d\xc2\xef\x5e\x62\x1a\xda\x1e\x04\xfa\x05\xff\xa7\xcf\xd9\x61\xf6\xd6\x8d\xf5\xb2\xc6\x09\x98\x01\x7f\xfe\x21\xe7\xb2\xfd\x51\xfa\xd5\xf1\x02\x39\x8b\x0f\x17\xe6\x35\x11\xa8\xae\x41\xf5\xf9\x85\x9d\x44\x09\x4d\xdb\xe2\x31\x52\x30\x3b\x0b\xd8\x93\x71\x81\x86\xf4\x7d\xda\x6b\x05\x5a\x85\xc7\xad\x1e\x75\xa2\xab\xde\xc0\x90\x55\x70\xde\x16\x1d\x63\x6d\xa6\x61\xa9\x04\x33\x90\x37\x7b\x45\x30\x07\xed\x47\x02\xec\xf7\x8f\xf1\x49\x29\x02\xcc\xce\x1b\xba\x01\x82\x2c\x96\x47\x52\x6f\xa9\xc9\xf9\x95\x79\x0f\x00\x00\xff\xff\x5e\x42\x13\x8e\xe5\x00\x00\x00")

func templateCcdFileTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateCcdFileTmpl,
		"template/ccd.file.tmpl",
	)
}

func templateCcdFileTmpl() (*asset, error) {
	bytes, err := templateCcdFileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/ccd.file.tmpl", size: 229, mode: os.FileMode(420), modTime: time.Unix(1503828686, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateClientOvpnTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xce\x31\x4b\x04\x31\x10\x05\xe0\x7e\x7e\xc5\x80\x8d\x16\xb9\x2d\xec\x64\x11\xc4\x42\x41\xd4\xab\xb4\x10\x8b\x6c\x76\xee\x36\x5c\x76\x26\x24\x93\x85\x78\xec\x7f\x97\xec\x89\xdd\xcc\xf7\x8a\xf7\xae\x50\x27\x9f\x51\x96\xc8\x78\xf0\x81\xd0\x67\xb4\x45\x65\xb6\xea\x9d\x0d\xa1\xe2\x91\x98\x92\x55\x1a\x71\xa8\xf8\xf5\xfe\xb1\x7f\xfd\xbe\x9e\x54\x63\xbe\xeb\xba\xa3\xd7\xa9\x0c\x3b\x27\x73\xe7\xec\xd8\xc9\x12\xe7\x1b\x00\x17\x3c\xb1\xc2\x48\x0b\x6a\x61\x88\x49\x54\xb0\x8c\x11\x12\xcd\xa2\x84\xe7\x33\xee\x9e\x25\x2b\xdb\x99\x70\x5d\xb7\x7f\x2f\x49\x71\x5d\x21\x51\x96\xb0\x98\x44\x9a\x2a\x7a\x3e\x78\xf6\x4a\xc0\xd9\x38\x4a\x6a\xb4\x46\xc2\x4c\x69\xa1\x04\x2c\x83\xe7\x11\x22\xa5\xec\xb3\x9a\x13\xd5\xff\xbb\xb5\x3a\x99\xa3\x09\x3f\x02\x0b\xa5\x01\x6f\xc1\x16\x9d\x0c\x8b\xb3\x6e\x22\x80\xde\xd9\x7b\x68\xbd\x8f\x0f\xb8\xae\x7d\xd7\xde\xbe\x55\xfc\x29\x6d\x6b\xfa\xee\x42\xfd\x89\xea\x25\x78\xa1\xba\xf9\x06\x4d\xfc\x01\x77\x6f\xf2\xf4\xb9\x6d\x97\xa2\x64\x58\x62\x09\xa1\x65\xc4\x63\xe3\xdf\x00\x00\x00\xff\xff\x1b\xa9\x43\x1f\x63\x01\x00\x00")

func templateClientOvpnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateClientOvpnTmpl,
		"template/client.ovpn.tmpl",
	)
}

func templateClientOvpnTmpl() (*asset, error) {
	bytes, err := templateClientOvpnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/client.ovpn.tmpl", size: 355, mode: os.FileMode(420), modTime: time.Unix(1503776353, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDh4096PemTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd4\x35\x12\xa4\x08\x00\x40\xd1\x9c\x53\x4c\x4e\x6d\xe1\x16\xe2\xee\xde\x19\xd0\x38\x8d\xfb\xe9\xb7\x76\xe2\xfd\xe1\x3f\xc0\xfb\xe7\xbf\x38\x51\x56\xad\x3f\x82\xf2\xc7\x61\x3d\xd6\x14\x03\xd1\xf3\xff\x7e\xc0\x54\x55\x51\x70\x75\x9e\xad\x45\xb6\x2f\xb3\xc9\xd2\xe7\x37\x56\x43\x1c\xa4\x1b\x2d\xac\xa4\x52\xd1\xb4\xbe\xa5\x86\xf7\xd6\x99\x21\xd2\xde\x29\x40\x6c\xec\xf4\x04\x64\x3e\x11\x13\xee\x7e\x38\xf0\xc5\x28\x74\x5d\xa8\x62\x55\x98\x61\x1a\xea\x16\x8e\x9e\xaf\x9b\xd5\x15\x48\x63\x88\xf6\xd1\xa8\x1f\x03\x39\x69\x95\x7a\x63\x88\xe6\x15\x81\x75\xb8\x4a\xdb\x34\x5f\x4b\x1c\xdf\xc5\x06\x57\x17\xc0\xf7\x75\x64\x4a\xc5\x2d\x99\x2d\x90\x54\x89\x30\x0f\x94\x8f\x86\xce\xa5\xda\x3e\x8b\xaa\x8d\x2f\x7c\x14\x4d\xab\xeb\x97\x43\x29\x3e\xd4\xd9\x9e\x62\xf5\x10\x0d\x6a\xcb\xbf\x7c\x5f\xfa\xb5\x77\x01\x52\x8d\x90\x43\x92\xb3\xfa\x3a\xdd\x57\xdd\x33\x6c\xbb\x15\x03\x5d\xe9\xe4\x5e\xb0\x56\x9e\x92\x31\x26\xb3\x57\x2e\x0a\xf2\x73\x82\xf0\x0c\xd7\x9f\xd8\xef\xa7\xfc\x40\xe5\xf0\x9b\xa9\x9b\x42\x02\xdd\xf1\x83\xa0\xe0\x73\x4e\xc7\xd5\x0b\x62\x43\xc2\xc7\x2e\x11\x97\x73\x4d\x8b\xf9\x94\x4c\xd6\x84\x76\x79\xbf\x20\xb5\xab\xd4\x51\x22\xd0\xb0\x96\x60\x63\x64\x61\xec\xb6\xf5\x2b\x96\x04\x99\x03\x95\x47\xc4\xe1\xc7\x6b\x2e\xab\x66\xde\x60\xac\x67\x29\x30\x92\x96\x55\x31\x46\xa3\x6f\x25\x0f\xdb\x2f\xd6\x9f\x10\xef\x0b\x0e\xbf\x90\x49\x13\x9a\xf1\xc3\x3f\x06\xb7\xf9\xfe\xb7\x40\xd4\x6e\x01\x66\x8f\xa4\x08\x8e\x5b\x8c\xcb\x9f\xf0\xc9\xb6\xa0\x09\x09\x6a\x33\x15\x30\x08\xbb\xfd\xac\x4f\xfc\xd0\x38\x7d\x5b\x48\x07\xe3\x3b\x2a\x63\x6b\xd0\x7a\x03\x72\x20\x5a\xd6\xec\xa7\xcc\x9e\x30\x07\xe6\x25\x22\x07\x29\x4c\xa4\x09\x49\xba\x7e\x88\xdd\xad\x58\xd0\xfa\x3a\xf4\xf3\xa2\x12\xfe\x0b\x3b\x63\x9a\x7d\xe1\xf5\x36\xa2\x1c\xca\x8c\x27\xd5\xdf\x13\x2f\x56\x83\xf0\xbc\x0d\xba\xed\xa2\x02\x6c\x54\xc0\xa0\x09\x1f\xbe\xd8\x18\xbb\x44\xcb\x11\xad\x82\x53\xbc\xcc\x0e\x63\xe6\x41\x70\x27\x10\xb5\x3c\xda\x12\xfd\xab\x2d\xf9\x26\x71\xc7\xe4\xc0\x98\x70\xcb\xfe\xfd\x40\x59\xf6\xdb\x18\x1d\x30\xe1\x62\x4f\x14\xea\x8c\xbe\xee\x59\x80\x2b\x4d\xea\xc2\x0e\xa9\x6e\x32\xcd\xc7\x34\x25\x88\x8e\xf7\xbf\x66\x09\xbb\x3e\xd6\xd6\x6e\x87\xc3\x33\xf2\x66\x74\xeb\xd5\xd4\xb7\xc4\x72\xe8\x58\x18\xb0\x8c\x92\x87\x7a\x38\xe5\xe0\xd5\x8a\x3f\x37\x16\x7a\x35\x96\xf8\xd0\x09\x4f\x9a\x0d\x3a\x3b\x64\x13\x21\xda\x1e\xd3\xac\x7c\x66\xf4\x88\x1b\xc6\x28\x5d\x70\x69\x37\xbe\xae\x55\x56\xdb\x7f\x27\x80\x2b\x14\x4a\xd9\x50\x26\x14\xfb\x83\x14\xb4\x8e\xa1\xb3\x2e\x3f\x3f\xd0\x90\xb9\xfb\x66\xc6\xe3\x30\x83\x6e\xce\x08\x0b\x8a\x2f\xe8\xda\x95\xc2\x8d\xf5\x68\x42\x94\x6f\x3f\xca\x13\x67\x27\x05\xc0\x64\xa4\x4a\x11\x14\x5e\x69\xd2\x71\x89\xaf\x5b\xe2\x39\x8e\x47\x69\x0f\xee\x20\x33\x1b\xf7\xc6\x05\xd0\x8a\x91\x6e\x8b\x94\xd2\xb3\xa6\xb5\xd9\x4c\xad\x51\x77\x55\xad\x21\x7b\xb6\x26\x22\x03\x28\x1a\xd9\x19\xda\x0f\x12\x54\x7d\xba\xaf\x83\x48\x3c\x25\xbc\x90\x54\xa3\x74\xa1\x5b\x2e\x44\xfd\xc8\xa5\x29\xc6\x4d\x78\x6f\xcf\xe9\x3c\xf2\xd9\x3e\x7c\xab\x38\xf9\x6f\x3d\x2d\xed\xd5\x46\x07\xa0\xda\x1a\x19\xcd\x3c\x57\x5a\x88\x1d\x40\x87\xd1\x91\x6a\x5a\x3d\x96\x00\x23\xbb\x22\xaa\x33\x70\xf6\x85\xfc\xdd\x93\x4b\xd0\x44\xc1\xaf\x46\x8a\x89\x8b\x5b\x82\x95\x33\x84\x6f\xc1\x19\x50\x0e\x54\x66\x34\x5e\x0c\x9c\x53\x9b\x1b\xa9\x8c\x28\x26\xf7\x3c\x13\xe8\xa1\xa6\xcc\xbd\x07\xef\x56\x65\x3a\x26\x9f\xd8\xfc\x0b\x4a\x49\x2c\x8b\xe7\x9b\xa1\x90\x7d\xd1\x93\x9c\xbe\x53\x3e\x12\x3e\x0b\xd8\xcc\x2d\xce\x05\xfa\xb3\xc2\xc1\xb6\x02\x7c\x6f\x38\xa7\x81\x30\x2b\x75\xe4\xcf\xa8\xb6\x42\x6a\xd1\xa7\xd2\xe9\xad\x3f\xa6\x45\xea\x6d\x79\x3c\xdf\x67\x20\x7e\x3c\x4b\xec\x55\xb2\x21\xbe\x11\x40\x63\xe2\xb9\x59\xa6\x3d\xc9\xf4\xf7\xd0\xdc\x66\x38\xdc\xbb\xcf\x60\xbb\x91\x6a\x38\x0c\xe4\xa6\x2c\x10\x24\xd9\xe0\x36\xd8\x5b\x9a\x3f\x47\x7b\xd6\x97\xa7\x46\xcf\x64\xb0\x47\x58\x65\x70\x07\x50\xa3\x82\xbd\x64\x78\x9e\x8d\x91\x76\x1e\x6c\xe3\xd6\xd4\x26\x14\x89\x98\xfd\x7d\x9a\x45\x27\x28\xca\x67\xe0\xfd\xbe\xb7\x03\x0c\x19\x22\x5d\x64\x0f\x3c\x30\xb9\x55\xfc\x34\xcc\x8d\x8e\x6a\x04\xb0\x8a\x7c\xbb\x51\xd0\x97\xcb\x4d\x1f\x54\xcf\x28\x35\x8c\xee\x09\xa3\xfc\xc8\x8b\xce\x24\x89\x8a\x79\xc7\xd0\x34\xa3\x04\xcd\x58\x19\xec\x7d\xc6\x79\xb2\x9c\x21\x47\x1a\x0b\x52\x10\xfd\x89\x01\xea\x98\x67\x1d\x43\x35\xa9\xbd\xc8\x99\xd6\xf7\x63\xee\x77\xd4\x7e\xbf\x44\xd3\xc0\xd6\x46\x9e\x40\x01\x43\xf8\xf7\x78\x8e\x78\x71\x15\x4d\xbd\xd7\xdb\x64\x47\xe8\x60\x85\x3b\xcd\x70\x74\x5f\x01\xe9\x6e\xca\xfc\x55\xbd\x3a\x07\xc7\x1b\xe6\xc7\x18\x92\x3d\x07\xd3\xbd\x42\xfd\xa8\x2b\xd8\x6e\xdd\xd2\x23\x19\x65\xfe\x45\x17\xf8\xeb\xb0\x68\x09\xff\xa7\xf3\xbf\x01\x00\x00\xff\xff\x2c\x8e\x5a\x0b\xbc\x05\x00\x00")

func templateDh4096PemTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDh4096PemTmpl,
		"template/dh4096.pem.tmpl",
	)
}

func templateDh4096PemTmpl() (*asset, error) {
	bytes, err := templateDh4096PemTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dh4096.pem.tmpl", size: 1468, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIptablesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templateIptablesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateIptablesTmpl,
		"template/iptables.tmpl",
	)
}

func templateIptablesTmpl() (*asset, error) {
	bytes, err := templateIptablesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/iptables.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServerConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x7a\x6d\x6f\xdb\xb8\xb2\xff\x7b\x7d\x8a\x81\xfd\xa2\x09\x60\x2b\x8e\xdb\xee\x69\x37\xff\xff\xbd\xf0\x3a\xee\xd6\xd8\xd6\xf5\x6d\x9c\x7d\x00\xf6\x0d\x2d\x8d\x2d\xc2\x12\xa9\x25\x29\xbb\x3e\x07\xe7\xbb\x5f\xcc\x90\x94\xe4\x24\xdd\x8b\xb3\x8b\xa0\xb6\x34\x24\xe7\x79\x7e\x33\xf4\x5d\xad\x8d\x83\xdb\xdb\xf7\x6f\x12\xfe\xf4\xaf\x7f\x41\xba\xa6\x0f\xff\xfe\x77\x92\x0c\x61\x33\x5f\x83\x36\xf0\x78\xbf\x06\x8b\xe6\x88\xe6\xbf\x93\xbb\xda\x68\xa7\xc1\x65\x75\xe2\x3f\x35\x79\x4d\xa4\x83\x1c\x8f\xe0\x1a\x35\x80\x93\x2c\x4b\xc8\x0c\x0a\x87\x20\xc0\xe8\xc6\x61\x0e\xcb\x35\xbd\x54\x58\x8e\x5a\x5a\x51\x3f\xa1\x55\x80\xae\x40\xa3\xd0\x05\xda\x34\x19\xc2\xa3\xc5\x96\x7e\x32\x00\xb9\x83\xb3\x6e\x40\x18\xec\x88\xb7\x46\xe6\x7b\xa9\xf6\xc9\x10\x84\xca\xa1\x10\x47\x84\xda\xa0\xdf\x36\x07\xc1\x4b\xe1\x28\x8d\x6b\x44\x09\x52\x39\x34\x3b\x91\x61\x20\xe7\xd5\x98\x83\x74\x70\x92\xae\xa0\xed\x4d\xb7\x77\x4b\x4d\xbc\x2c\xfd\xe1\x27\xa1\x1c\x38\x0d\x99\x56\xce\xe8\x12\x44\x96\xa1\xb5\x50\xeb\x52\x66\x12\x6d\x32\x04\x7d\x44\x03\xae\x40\xf8\x75\xbd\x1a\xf1\x9a\xaa\xb1\x2e\x0a\xba\x93\x06\x4f\xa2\x2c\x93\x21\x98\xa6\x44\x0b\x3b\xed\xa9\xe9\x6f\xf3\xb8\xba\xd9\xcc\xd6\x97\x07\x7f\x51\xa0\xb4\x1a\xff\x26\x55\xae\x4f\x16\xec\xd9\x3a\xac\xac\xdf\x39\x13\x0a\xf6\xf2\xe8\xc5\x01\xfc\x56\x13\x17\x0e\x1a\x25\x1d\xa8\xa6\xda\xa2\x19\x81\x6d\xb2\x02\x84\x25\xb5\x4e\xc2\x7e\x61\xaf\x11\x34\x41\xc1\x63\xa5\x73\x1c\x04\x5e\xa4\x0d\x64\x95\xb6\xae\x3b\x2f\x88\xe4\xcd\xa6\xb4\x83\x5d\xa3\x32\x27\xb5\x4a\x86\xd0\xa8\x92\xb4\x40\x2c\xd5\xc2\x38\x29\xca\xf2\x4c\xce\xb3\x6b\xe8\x43\x2e\xad\xd8\x96\xc4\x24\xed\x11\x35\xd0\x4a\xfe\x82\xd4\x77\xc1\xe8\x49\x70\x2c\x72\xb2\xa8\x00\x85\x98\x5b\xbf\x70\xb6\x26\xb5\xbc\x9e\x82\xc8\x45\xed\xd0\x80\x12\x15\x1d\xb3\x33\xba\x62\x8a\x15\xba\x93\x36\x07\x98\x6b\xa5\x90\x99\xb5\x50\x0b\x85\x65\xf0\xa5\x64\xe8\x3d\xa6\xd2\x86\x2c\x20\x14\x68\x85\x29\x90\xf0\xbf\xaf\xe1\x61\x3d\x25\x21\x0a\xb9\x2f\xd0\x90\xef\xb2\x31\xc5\x99\x59\x20\x27\xb0\x58\xd2\xae\x47\xec\x84\xa4\x63\x7b\xcc\x3e\x17\x76\xb6\x8e\xdc\x92\x96\x57\xcf\x2d\x0b\x8d\x6d\x58\x7f\xb9\x56\xaf\x5c\x38\x8b\x8d\x72\x17\x2d\x05\x9f\xcf\x1b\xc1\x91\xf7\xf0\xf0\xe9\x66\xf3\xe9\x01\x8c\xd6\x0e\x32\x34\x4e\xee\x64\x46\x7e\x76\x95\x89\xeb\x51\xff\x49\x32\x84\x2b\xfa\x7a\x3d\x62\xcf\xaf\x8d\x3c\x12\xdd\x01\xcf\x70\x75\xc0\xf3\x75\x0a\xb0\x10\x59\x01\x59\x29\x51\xb9\x10\x1f\xc4\xb1\x0f\x7e\xef\xc4\xac\x2c\x57\xa0\x34\xa0\x4f\x8a\xb7\x27\xba\x64\xc8\xfb\xec\x64\x49\xca\xdb\x74\x8b\x68\x0f\x92\xde\x6f\x6a\xd9\x77\xc8\x5d\xac\xf7\x77\x2b\x2a\x84\x4c\xf8\x85\xc9\x90\xe4\x41\xff\x66\x80\xc2\x9e\xc7\xc6\x8a\x01\xe4\xd2\x60\xe6\xb4\x39\xb3\x0e\x05\x6d\x1d\x02\x6d\x07\x36\x33\xb2\x76\x3e\x88\xf6\xa8\xd0\x08\x27\xd5\x1e\xbe\x3e\xcc\xfa\xa2\xdb\x20\x4d\x4f\x66\x9b\x02\x7c\xc5\x0a\x29\x44\xc8\x92\x8d\xe5\x18\xa2\xc8\xf9\xab\x41\x98\xeb\xaa\xd2\x0a\x56\xc4\x5f\xb4\x9c\x17\x29\x6c\x85\xa4\x2a\xbd\xe3\x17\x5e\xb8\x8b\x03\xbd\x30\x33\x75\x86\xdf\xdf\x4e\xde\xb3\x72\x2a\xa1\xc4\x1e\x2b\xa2\xf4\x86\xe6\xe0\xdd\x22\x1d\x9d\x73\xc0\xd5\xa8\x28\xbe\xe8\xb1\x28\x2d\xf3\x04\x02\xd6\xbf\xcc\x1f\x60\x78\x3b\x25\x3e\x2a\xe1\x28\xa9\x45\x5d\x93\x49\x2d\x22\x0c\xea\x43\x66\x6f\xa7\x51\x55\xf2\x88\x20\x15\x1d\x08\xb5\xd8\xe3\x75\x9a\xdc\x65\x02\xa2\x42\x6f\x48\xfa\x9b\x4c\xa4\x99\x71\xc9\x1d\x5b\xf0\xf2\x95\x97\xd3\xbf\xa6\x93\x5e\x7c\x4b\x2f\x60\x08\x9b\x42\x5a\x66\x05\x6c\xa1\x9b\x32\x27\x79\x0e\x58\x3b\xb0\x94\x81\x5d\x92\x64\x82\xcb\xca\x7c\x36\x47\xe3\xd6\xc2\x15\x54\x5c\xf8\x50\x7e\xdc\x7b\x48\x3b\xd2\xb3\x5f\xf0\x1c\x1f\x25\x43\xb8\x97\xbb\x9d\x44\x28\xb0\x2c\xbd\x3c\x46\x54\xe8\xd0\x70\x8a\xfa\xd9\x5b\x1c\x7d\xda\x26\x87\xa4\x24\xfe\x63\x32\x04\x00\x5d\xa3\xb2\xb6\x84\xbc\xe0\x35\x30\xd6\x8d\x83\xbc\xb8\x9d\x4c\xdf\xa4\x35\x56\x40\x1f\xc8\xdf\x9a\xad\x75\xd2\x35\x0e\x61\x3a\x79\xf3\x8e\x8d\x4d\xaf\xfa\xb5\xa6\xb1\xbe\xbe\x30\xc1\x56\x3a\xef\x3f\xc9\x30\x2f\x7a\x1b\x26\x77\x79\xf1\x44\x55\x79\x41\x2b\xf8\x65\x5e\xb0\x6c\xf7\x1f\xd7\xc4\x8c\xed\x09\x18\x33\x94\xd3\xb5\x2e\xf5\xfe\x4c\x3c\xb5\xaa\xb4\xcd\x96\xea\xd0\x95\xc8\x73\x83\x96\xd8\x80\xa3\x14\xb0\x5c\x5f\x77\x39\x37\x66\x8f\x18\x62\xc7\x69\x3a\x49\xdf\xb3\x93\x96\xfa\x84\x26\xc4\xac\x4e\x86\x7e\xc7\x9a\x6a\x3d\xe6\x70\xe5\x0a\x54\xa0\xd0\xbd\x9e\x8c\x40\xa6\x98\x82\x80\x9b\xd7\x13\xa8\xd1\x84\xbd\xe8\x90\x7b\xdc\x89\xa6\x74\x96\x62\x84\x69\xe1\x8a\x52\xbf\xc1\x4c\x57\x15\xaa\x1c\xf3\xeb\x24\xb2\x1e\xd8\x25\xa9\xe6\x5a\xed\xe4\xbe\x31\x5d\xf6\xa0\xac\x45\x3c\x11\x03\xe5\x19\x04\x97\x92\xb0\x60\xc8\x6a\x8f\xfe\xef\x34\xe4\x46\x9c\x62\x58\x05\xd9\xa9\x52\x1a\x5d\x91\xd9\x7b\xf9\x85\x6b\x91\x13\x07\x84\xdb\x49\xfa\x2e\x9d\xa4\xb7\xbc\x95\x74\x16\xcb\xdd\x28\x14\x1c\x83\xd6\x79\xca\x2d\x42\x25\x88\x91\xa3\x90\xa5\xcf\xd6\x3a\x2a\x8e\x76\xee\x25\xc0\x76\x41\x24\x33\x1c\xf2\x17\x89\x40\xab\xf6\xd8\x94\x53\x06\xad\xa3\x54\x0d\xa5\x54\x08\xe4\x71\x9d\x1b\x25\xc3\xe7\xa0\x25\x6d\xf3\x5d\x0c\x56\x66\x9f\xab\x91\x54\x3b\x9d\x26\x77\x41\xd0\x70\xce\x04\xa6\x6f\xdf\xa6\xf1\x6f\x92\x84\xb7\xe4\x5b\x2b\x24\xe0\xc6\x1f\x3f\x0b\x7b\x08\xee\xf5\x59\x48\xe5\x84\x54\x04\xc6\x30\xd3\x26\xa7\xa4\x15\x24\xfc\x7f\xe3\xff\x6a\x81\xd1\x72\x1d\x15\x4d\x09\xce\x5a\x9d\x49\xe1\xcb\xa5\x54\x5e\xa4\x90\xdc\x97\xbb\xd6\x50\x7b\x8d\x16\x72\x0a\x3b\x4d\xda\x90\x96\x35\x2d\xc8\xbb\x46\x7c\x9a\xaf\xb9\x6a\xdf\xfa\x66\xc8\x78\xc2\x5a\xb9\x57\x98\x07\xfb\x70\x15\x78\xce\x48\x57\xc5\x6b\xad\x4b\xaa\xce\x0e\x4e\x82\xf8\xab\x0d\x1e\xa5\x6e\x2c\xf9\x51\xd8\x2a\x4d\xe4\x2e\x63\xa7\x1b\x13\xf5\xb8\x46\x63\xa5\x75\x20\xeb\x3a\x75\xdf\xfe\xc6\x27\x49\xdf\xcf\xed\x92\x0c\xe1\x8f\x08\xdc\x76\xd2\x58\xc7\xc9\x98\xd3\xcc\x97\x87\x57\xb6\x25\x84\x4c\xd4\x62\x2b\x4b\xe9\x28\x70\x9d\x0e\x80\xb2\xad\xf4\x2d\xa4\xf1\xf0\x92\x1e\xc7\xc3\x28\xf6\x97\xf3\x1e\xe8\xe1\xc2\xa9\x3a\xc0\x58\x09\xe5\x71\x80\x25\x4c\xcc\xa8\x62\xb9\xbe\x51\xe8\x2a\x32\xaf\x56\xbc\x5b\x38\xaf\xdd\x65\x04\x05\x1a\x84\x13\x7a\x33\x36\x55\x1b\x19\x6f\x6e\x2e\x5c\x27\x05\xf8\x20\x15\xef\xcf\xc4\x7c\x24\x9d\x24\xac\xe4\x50\x25\x4b\x18\xa1\x78\x73\xef\x01\x6d\xb4\x5e\xb1\x95\xff\x7f\xd8\xf8\xed\x04\x50\xe5\xf1\xdb\xed\x64\x72\x4d\x8a\x10\x65\xa9\x03\xec\xf0\x68\xf9\x89\x2f\xa4\x00\x9f\xd0\xa3\x89\x18\x2f\x3e\xa5\x38\xf6\x0b\x8a\x9d\x1e\xa8\x7c\x11\xf2\xb7\xd1\x31\x0e\x5a\x88\x92\x5e\x06\x09\x74\x7c\x76\x3c\xfe\x87\x0e\xc1\x98\x85\xb8\x17\x70\xff\x71\xbe\x1e\xd7\x46\x7f\x3b\x8f\xe0\xc4\xca\x8e\xde\xed\x44\x79\xf0\xe2\x92\x65\x62\x94\x04\x16\x59\xab\xb4\xb6\x4b\x1f\x9c\x54\x32\x94\x2d\xa4\xea\xf9\x7e\x50\x9f\x47\xd7\x94\x37\xef\x57\x0f\x2d\xa8\x8a\x19\x31\x85\x17\xbc\xd4\xc3\xd4\xef\xfb\xe9\x73\x2f\xa5\xd8\xfd\xbe\x9f\x3e\xf1\x52\xc6\xac\x0e\x7f\xf4\x76\x63\x8d\x69\x45\x4e\xa4\xcd\xc1\x92\x57\x46\x6d\x5c\x85\xbe\xa3\x03\xc3\xd7\x51\x61\x1d\x68\xf2\x6a\xe9\x61\x62\x90\xb4\x60\xab\x1b\xc5\x00\xdb\xab\x3b\x10\x3f\xb5\x37\xd9\x70\xdd\xd8\xc2\xb7\x99\x36\x2a\x3e\x64\xb7\xe0\x83\x27\x90\x2e\xea\x9a\x31\x1b\x89\xd5\x22\x41\xef\xd2\x16\xb6\x58\x48\xd5\xe6\x23\x0f\x71\xfa\x00\x91\x72\x8f\x2b\x90\xb5\xfb\x74\x2d\x17\x0a\x06\x6c\x04\xd5\xfd\x59\x07\xa5\x4f\x7c\x26\x71\x76\xe1\x0e\x1d\xbc\x0e\x96\xe6\xe4\x76\x15\xf3\xfb\x65\x90\x52\x05\xde\x8a\xec\xf0\xb2\x4f\xa5\xc9\xd0\xef\x7f\xfb\x7e\x9a\xde\xfe\xf0\x2e\x7d\xff\xac\x40\x44\x8f\x9f\x3e\x23\xbd\xfd\x1e\xe9\xdb\x67\xa4\xd3\xef\x91\xbe\x7e\x46\xfa\xfa\x7b\xa4\x6f\x92\x61\x4d\xa6\x1a\x84\x05\xff\xa0\x05\x14\x88\x4f\x17\x0c\x9e\x10\x86\x9d\x7f\x78\x89\x10\x36\x3a\x14\x00\xb0\x35\x66\x04\xbe\x7b\x31\xe4\x3d\x22\xbe\x48\x86\xad\x67\x12\x40\xd8\x81\x78\x9e\x94\xa0\x10\x16\x44\x34\x70\x32\x8c\xf8\xcb\x7b\x07\x48\xe7\x1d\x21\x40\x5d\x36\x39\x83\x2b\x32\x89\x9f\x03\x8c\xfa\xcd\x4d\xb3\xed\x5a\x97\x41\x96\xe5\xbe\xbf\x8e\x8e\xdf\xe3\x2b\xe4\x21\x0e\x77\xae\xb3\xd6\xe3\xfa\x97\x41\xc1\x75\x4a\x8e\xbf\xf8\x7d\xf6\x79\xfd\x69\xf1\x23\x3c\x10\xa2\xb3\xfd\xa0\xf2\x4d\x2d\x89\xc5\xcf\x7a\x4d\x61\xe6\xdb\x1a\xea\x92\x61\xb0\x29\xb0\xd4\x8a\xea\x28\xa9\x32\x48\x43\xf2\xdb\x8a\x5a\xb6\x4b\xd9\x29\xd8\x3b\x7d\x51\xbd\x10\x59\x21\x15\x76\xe3\x85\x68\xa9\x37\x93\xf4\x76\xfa\xee\xc2\x8f\xa7\x6f\xde\x51\xde\xf8\x40\x29\x6a\x04\x8d\x0a\x89\x9e\x93\x3c\x47\x15\xe7\x7f\xfb\x63\x32\x0c\xca\x09\x95\x3c\x97\x06\xb2\x2c\xff\xbf\xdd\xdc\x83\x42\xd5\x0d\x9e\xb8\x25\xc9\xb2\xfc\xa6\x13\x32\x66\xb6\x50\x6d\x7c\x97\x20\x2f\x77\xf6\xbc\xc3\x13\xde\x93\xd0\xe7\x84\x40\xa7\x9c\xd2\x6d\xfb\xea\x49\x3e\xf0\x48\x3b\x4c\x85\xc2\xc8\x84\x4b\xbb\xb4\x80\xdf\x44\x55\x97\xe8\x37\x6a\xb3\x26\xa5\xdf\xae\xd9\x20\x86\xa4\xda\x8f\x78\xc2\x12\xd3\x77\x00\xe8\x1d\x92\xf4\xc5\xa8\x1b\xbb\x51\x85\x18\xf8\xac\xd0\xeb\x02\xed\xcb\x9e\xd2\x1f\x63\x85\xe1\x51\x4f\x4d\xa4\xbc\x6f\x98\xb3\x53\xf7\x0a\x92\xde\x51\x28\xbf\x67\xac\x1b\x6d\xf9\x77\xa6\xbc\x7b\xd9\x94\xcf\x9f\x72\x1f\x38\xbf\x8f\x2d\x51\xb0\xa4\xc8\xf3\x1e\x32\x20\x04\x71\x61\xcc\x60\xbd\x16\xf2\x51\xc6\x88\xec\xc5\x0f\xd3\x24\xc9\x4c\x39\x3e\xa2\x91\x3b\xdf\x5a\xce\xbf\x7e\xea\x8e\xe9\xe2\x46\xb8\x0b\x95\xa0\x62\xc8\x9f\xcb\xdd\x0e\x8d\x0f\xa7\x76\x8a\xf3\x64\xda\xc7\x81\xd9\x12\xc2\xde\xe8\xa6\x0e\x93\x89\x0e\xe6\x6c\xb8\xe6\x91\x6d\xdd\x49\x43\x85\xae\xd0\x39\x0b\x70\x75\x7b\x0d\x5f\x1b\x05\x55\x53\x3a\x49\x7e\x11\xd3\x7b\x2e\xb0\xd2\xca\x8e\x40\xab\x80\x47\x44\x56\xb0\xc4\xe0\xcf\xf0\x23\x9c\x96\xab\x17\x87\x68\x81\x3e\x2e\xf7\x0b\x6f\xfc\xd6\x20\xea\xda\xe8\xda\x48\xe1\xb0\x3c\x93\x3d\xaf\xa6\xd7\x70\x35\xcb\x8f\x42\x65\x98\x5f\xc3\x3c\x06\x92\x1f\xb0\x70\x4f\x76\x56\xa2\x92\x19\x21\xc6\xb0\x75\xa5\x73\xd2\xec\xc5\x4c\x4f\x2a\x6a\x02\x6a\xad\x2c\x5b\xcd\xeb\x2b\xb2\x42\x78\xbe\xd3\x56\xa7\xa1\x07\x9f\xeb\x02\xd9\xf3\x8c\x47\xb0\xa2\x44\x61\xd4\x38\x7a\xa3\x67\x2b\x4d\xee\x2e\x1f\xa7\x37\xfe\x45\xe2\xe7\xb5\xde\x92\xf9\xc8\x7b\x52\x37\x1c\xf1\x83\xe7\x88\xfe\x38\xf9\x75\x23\x2a\x46\x0a\x9e\x36\xc0\xb2\xdc\xf7\xc0\xc9\x90\x3a\x60\x6e\xd6\xf7\xc2\xe1\x49\x90\xe8\x46\x37\xfb\xa2\x9b\xf6\x66\x22\x0e\x0b\x68\xcb\xe5\x1a\x9c\x11\x3b\x2a\x4b\x31\x4b\x9e\x70\x0b\x5b\xa3\x4f\x1e\x4c\x32\xe4\x88\xe8\xae\xd4\xfa\xd0\xd4\xcc\xc0\x5e\x3f\xdd\x9a\x4c\xb4\x79\x86\x00\x62\x16\xbe\x18\x4a\xae\x66\x1b\x72\x41\x73\x81\xf2\x9e\x7a\x47\x44\x14\xfc\xc0\xa3\x7a\x49\x7d\x5c\x8e\xa6\x9d\x02\x13\x0d\x8b\x4b\xbe\x82\xa6\x3c\x5f\xa7\xc9\x5d\x28\xce\x41\x43\xe3\xa8\x89\x1c\x77\xb7\xb0\x3d\xd7\xc2\xda\x71\x5e\x64\xf5\xe0\x3f\xa0\x24\x28\x8e\x86\xbb\xd4\x80\x14\xdb\xea\xd8\x6a\xdc\xa2\xa3\xb4\x48\xae\x14\x9a\x48\xda\xde\x0b\x1c\x2c\xd7\x95\xa2\xfb\xd5\x83\xd7\xc0\x6f\xcb\x97\x51\xf3\x7c\xf6\xeb\x62\xb6\xa1\x18\x2c\x9c\xab\x7f\xbc\xb9\xd1\x35\xaa\x63\xad\x52\x85\xee\x66\x27\xfe\x4a\x0b\x57\x95\x43\xe2\x2e\x13\x47\x14\xce\x86\xa9\x43\x87\x2c\xb6\x48\x65\xc0\xe0\xce\x8f\x0f\xb9\x4b\x6d\xb6\x25\xd7\xf3\x0e\xaa\x5b\x52\xdd\x51\xe6\x98\xc3\xf6\xcc\x23\xa9\x5c\xd9\x34\xd3\x55\xab\x48\x3a\x63\xac\x6b\xae\xfc\xb4\x6e\x3a\x79\x97\xfe\xf0\x8f\x74\x3a\x9d\xd2\xdf\x20\x79\x99\xea\x5d\xca\xff\xb3\xea\x1e\xdb\x1c\xfc\xc4\xcb\x5b\x00\xdc\xcf\x64\x3d\x2f\xef\xcd\x36\x06\x16\x71\x00\x1d\x3e\xa6\x8c\xf0\xd3\x39\x3a\xfe\xe8\x62\x7c\xeb\x0b\x97\x0d\x73\x8b\x16\x89\x12\x16\xdb\x69\x93\x61\xff\x88\x17\x48\x47\x61\xec\x7e\x89\x9a\x99\xd9\x7e\x4e\xba\x48\x6e\x84\xc4\x78\xf1\x2b\xfb\xd2\x65\x41\xa8\x29\x4e\x8f\x03\xfe\xf9\x3b\xad\xc8\x5d\x97\x6f\x03\xa7\x84\x66\xe4\xbe\x70\x11\xe0\x74\x3d\x90\x9f\x4f\x77\x00\xea\xe6\x80\x67\x2e\x07\x04\xd1\x08\xcb\x75\x78\xca\xc6\x22\xcf\x63\x90\x76\x3a\xc6\x53\xa2\xd2\x8f\xae\x1d\x5a\x06\x9b\x75\x63\xa8\xf2\xd0\x8a\x0f\x9a\x7a\x11\x9d\x37\x7c\x37\x41\xd0\x91\x10\x24\xf6\xc6\x50\x01\x6c\x32\xce\x94\x04\x5e\xc3\xd8\xfd\x92\xa3\x5a\x48\xe3\x67\xce\xcb\x0f\xf0\xc7\x97\x47\xf8\x38\xfb\x75\x01\xab\x2f\x1b\xf8\x79\xb1\x5a\x7c\x9d\x6d\x16\xf7\xb0\x5c\xdd\x2f\x7f\x5d\xde\x3f\xce\x3e\x51\xc0\x2d\xbe\x6e\x96\x1f\x96\xf3\xd9\x66\x71\xf3\xcb\xe2\x0f\x58\xcf\x96\x5f\x1f\xe0\xc3\x97\xaf\xb0\x98\xcd\x3f\xc2\xfc\xd3\x72\xb1\xda\x10\x2f\xfc\xf5\xe3\xec\xd7\xe5\xea\x67\x58\x6e\x1e\xe0\xcb\x6f\x2b\x78\x5c\x2d\xff\xe7\x71\x01\x83\xf9\x97\xcf\x9f\xbf\xac\x60\x35\xfb\xbc\x18\x10\xed\xe3\x8a\x9e\x2c\x56\x1b\xd8\x7c\x5c\x3e\xc0\xa7\xe5\x6a\x01\x5f\x1e\x37\x69\x72\x97\x37\x75\xc9\xfc\x8e\x33\xbe\xcb\xa1\x60\x3a\x20\xd6\xa2\x24\x9b\x74\xd6\xa1\xfc\x89\x16\x6a\xa9\xf6\xe3\x52\x1e\x78\x30\x81\xd6\x8a\x3d\x46\x97\xb5\xa4\x14\xee\x89\xb8\x04\x6a\xe3\x0a\xbe\x71\x0b\xbd\x5b\x29\xd5\x01\xac\xf6\x55\x9d\xf5\xc8\x1d\x26\xf5\x63\x96\xba\x4f\x15\xe8\x7c\x1b\xc8\xef\x08\xf4\xee\xa9\xd2\xe6\xfa\xa4\xc8\x93\xd7\x64\x24\x3c\xa2\x39\xc3\xed\x04\x2c\x66\x5a\xe5\x76\x14\xc7\x29\xbc\xb3\xc1\x4a\x73\x87\x50\x23\xf7\xac\x7e\x0c\x26\x77\xa0\x34\x33\x1f\x1b\xfb\x1c\xf2\xc6\x84\x82\x00\xb7\xd3\xb8\x1b\x38\x59\x21\xd4\x68\xa4\xce\xd3\xa4\x53\xc4\xed\x84\x88\x48\x41\xe4\x18\xf8\xcd\x19\x41\x2b\x1a\x43\x7d\xfb\x16\xcf\xbc\x94\xce\x8f\x69\x85\xfa\xc3\x73\xbc\x04\x1a\xf5\x2e\x53\x07\x1f\x3f\xcf\xe6\x6d\x10\x0d\x7c\x57\x5a\x60\x59\xc3\xb6\xd4\xd9\x01\xee\xf5\x03\x08\xe7\x44\x76\xb0\xac\xc8\xc7\xfb\x35\xf0\x25\xf0\xae\xd4\x3a\xf7\x23\xb0\xfe\x84\xfd\x72\xb0\x7e\xac\x15\x8c\xc7\x7b\x54\x07\x3c\xc3\x78\xec\x87\xfd\xe0\x44\xca\xde\x78\x39\xa2\x6d\x6f\x4b\x82\x43\xb7\xf7\x48\xac\x93\x4c\xd7\x67\x7f\x8f\x22\x2d\x1c\xf0\xdc\xcd\x77\x59\x4f\xed\xbc\xbf\x77\xbf\xf0\x6a\xf2\xca\xcf\x5e\xdd\xe5\x21\xaf\x6e\x5f\xc5\xa7\xed\x48\xf7\xce\x95\x76\x2c\x1a\x0a\x67\x66\x0e\x26\x17\x97\x16\xd2\xb6\x17\x15\x43\x78\xe0\xcb\x3c\xe2\xc9\x9c\x6b\xa7\xf7\x46\xd4\x85\xcc\x20\x93\x75\xc8\x88\x9b\xd0\x04\xed\xe4\x1e\xa4\xc3\xca\x8b\xb2\xa5\xb6\xaa\x96\x9c\xc3\x82\x73\xc5\x0b\x21\x4f\xca\x27\x71\xe9\x2f\xcb\x34\xb9\xf3\xfb\xc1\x4f\x1f\xc6\xf3\x9f\xe6\x10\xfe\x1b\xc2\x4f\xa5\x3e\xed\xa4\x2d\xe0\x2a\x64\xdd\xeb\x96\x74\xb6\x78\x18\xdf\x4e\xdf\x05\xfa\x21\x7d\x6f\xdf\xdd\x2f\x1e\xc6\x8b\xfb\xc5\x6b\xff\x72\x08\x1b\x43\x99\x6d\x7c\xbf\x78\x60\xe8\xef\xa1\x6c\xa6\xab\x9a\xef\x0f\xb4\x8a\x1a\x22\x0c\x41\xb1\xd2\xbb\xd9\x0e\xb8\x57\x3a\x1e\x21\xf6\x6e\xae\x29\x4b\x53\x56\x6a\xdf\x4b\xf5\x1d\x31\xd3\x84\x8e\x1a\x97\xff\xd4\x31\xc8\x2b\xf1\x4d\x56\x4d\x15\x2e\xa3\x19\x18\x6b\x95\x35\x86\xca\x51\x79\x8e\x19\x97\x1d\xb9\x2d\x32\xd8\x42\x71\xae\x5f\x69\x72\x57\x89\x6f\xe3\xf8\x3a\x0c\xef\x96\xee\x15\xf5\x2a\x7b\xad\x73\x90\x39\x8a\x80\xdb\x9a\xec\x62\xdc\x92\x0c\x03\x98\x7e\x65\xb9\x4f\x93\x25\x52\x2a\x11\x3b\x1e\x35\x29\xe9\xa4\x28\xe5\x3f\xb9\xf1\xf6\x0e\xff\x47\xb8\x54\x6f\x2e\xeb\x07\xb5\x37\x3c\x8a\x7b\xe1\x26\x3e\x4d\x1a\x8b\x06\x94\xde\xea\xfc\x9c\x30\xcc\x8e\x5f\x82\x12\xe2\x30\xda\x57\xf0\x50\x46\x9d\xe1\x51\x9c\x38\x6a\x99\xb7\x1d\x23\xcf\x22\x02\x16\x32\x68\x75\x63\x32\xe4\xa9\x5a\x98\xac\xb3\x83\x09\xe7\xf1\x9e\x86\x52\xab\x3d\x1a\xae\xe4\x7e\x39\x19\x68\x8b\x9c\x45\x7d\x1b\xc2\xc0\x24\x0a\xce\x59\x6a\x6f\x44\x8e\x69\x12\x78\x1a\x53\xc4\xc6\xcf\xe1\xa2\xfd\x4b\xe3\xea\x86\x02\xc1\x16\x94\x11\xac\x13\xae\xe9\x6e\xf9\x4e\x3e\x9d\x05\x1b\xb6\x43\x01\xee\x57\x9c\x69\x14\x65\xf9\x08\x6b\x0d\x9e\x8c\x74\x0e\x55\x48\xa7\x95\x54\x8d\xc3\x34\x09\x5b\x86\x54\x32\xf6\x5f\xd3\x52\xef\x93\x27\xc0\xa3\xd4\xfb\xae\x00\xb0\xda\x08\x1a\x7b\xbc\x65\xcf\x96\x5e\x5f\xe9\x70\x15\xd3\xfe\xa8\x41\xee\xc0\x34\x4a\x31\xc2\xb6\xfe\xc2\xf8\x28\x33\xe4\x5f\x30\x9c\x7b\xbb\x84\x68\x1d\xfc\xb9\x36\x14\xed\x15\x7c\xa0\xba\xfe\x67\x70\x9d\x3f\x4b\xbd\xef\x5d\x3e\x5f\xc7\xdf\xa3\xd0\x99\xda\xd0\x3f\x63\x51\x13\xa2\x63\xa0\x73\x44\x63\xa8\x98\x78\xb4\xe1\xd9\xa7\x15\x03\xde\x25\x18\xdc\x2b\xc7\xd7\x29\x1d\x12\x83\x56\x1d\xa6\x27\x0b\x37\x35\x55\xd2\x53\x41\xef\x06\xdd\x21\x61\x8f\xee\x44\xe9\x52\x60\x7e\xb4\x42\x0f\x76\xbb\xb2\x76\xb5\x6d\x9c\x9f\x1f\x68\x57\x10\x68\xa7\xd3\xe2\x7f\x11\xea\x92\xb6\xef\x7a\x52\x5c\xbc\xe0\x74\xc8\xfd\x4f\x1f\x96\x41\x89\x47\x2c\xc9\xad\x88\xc6\xe3\x20\x38\xa2\xd9\x6a\x2b\xdd\xd9\x07\xd0\x84\x93\xaa\x2c\x51\xb9\x11\xe0\xb7\x0c\x6b\xc7\xe0\x67\x27\x9c\x28\x01\x8d\xd1\x86\x00\xd7\x1b\x8f\x93\x84\xd5\x3e\xab\x74\x17\xf8\x25\x34\x64\xee\x64\x08\x6f\xd9\x87\x7e\xe0\x78\xe4\xd2\x45\x5d\x28\x6e\x9b\x7d\xcf\xe9\xa8\x12\x6e\x4b\xac\x68\xcf\xf7\xc0\x33\x16\x67\xb0\x22\x04\xe9\x19\xc3\x84\xfe\x85\xd7\x2c\x12\xb1\x95\x21\x18\xac\xd1\xff\x50\x20\xfa\x56\x0a\x30\x73\xfe\xf7\x2e\xd3\x09\xc3\xcd\xbf\x1a\x54\x94\x1d\x3a\xf7\x0b\xd1\xc4\x98\x30\x3c\xe4\x6e\xc4\xe1\x5e\x9b\x73\x7b\x5f\xa8\x7d\xf4\x04\x27\x2d\xf5\x9e\xf2\x97\xbf\x5f\x4e\xfe\x37\x00\x00\xff\xff\xe6\x4d\xee\x30\x72\x25\x00\x00")

func templateServerConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateServerConfTmpl,
		"template/server.conf.tmpl",
	)
}

func templateServerConfTmpl() (*asset, error) {
	bytes, err := templateServerConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/server.conf.tmpl", size: 9586, mode: os.FileMode(420), modTime: time.Unix(1503828686, 0)}
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
	"template/ccd.file.tmpl": templateCcdFileTmpl,
	"template/client.ovpn.tmpl": templateClientOvpnTmpl,
	"template/dh4096.pem.tmpl": templateDh4096PemTmpl,
	"template/iptables.tmpl": templateIptablesTmpl,
	"template/server.conf.tmpl": templateServerConfTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"ccd.file.tmpl": &bintree{templateCcdFileTmpl, map[string]*bintree{}},
		"client.ovpn.tmpl": &bintree{templateClientOvpnTmpl, map[string]*bintree{}},
		"dh4096.pem.tmpl": &bintree{templateDh4096PemTmpl, map[string]*bintree{}},
		"iptables.tmpl": &bintree{templateIptablesTmpl, map[string]*bintree{}},
		"server.conf.tmpl": &bintree{templateServerConfTmpl, map[string]*bintree{}},
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

