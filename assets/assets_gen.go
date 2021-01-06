// Code generated by vfsgen; DO NOT EDIT.

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 1, 6, 20, 7, 1, 85429995, time.UTC),
		},
		"/schemas": &vfsgen۰DirInfo{
			name:    "schemas",
			modTime: time.Date(2021, 1, 6, 20, 52, 3, 459726156, time.UTC),
		},
		"/schemas/manifest.json": &vfsgen۰CompressedFileInfo{
			name:             "manifest.json",
			modTime:          time.Date(2021, 1, 6, 20, 52, 3, 459571737, time.UTC),
			uncompressedSize: 23404,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xdb\x8e\x1c\xb7\xb1\xef\xfa\x0a\xa2\x61\xc0\xf6\x39\xd2\x4a\xf2\x39\x76\x60\x01\x79\x50\x6c\x05\x59\x43\xb1\x14\x69\x65\x3d\xc8\x9b\x01\xa7\xbb\x7a\x9a\x1a\x36\xd9\x26\xd9\x33\x3b\x31\x04\xf8\x43\x92\x9f\xf3\x97\x04\xbc\xf4\x9d\xec\xdb\x8e\x64\x25\xf0\x02\x82\x76\x79\xa9\x2a\x16\xeb\x4a\x16\xfb\xe7\x3b\x08\x45\x9f\xc8\x38\x83\x1c\x47\x8f\x50\x94\x29\x55\x3c\xba\x7f\xff\xad\xe4\xec\x9e\x6d\xbd\xe0\x62\x77\x3f\x11\x38\x55\xf7\x1e\xfc\xe1\xbe\x1b\x79\x57\x4f\x53\xa7\x02\xf4\x1c\xbe\x7d\x0b\xb1\x72\x6d\x44\x51\xd3\xf8\x57\x2c\xf6\xa0\x0a\x8a\x63\x40\xb8\x28\x50\x7b\xa2\x80\x9f\x4a\x22\x20\x89\x1e\xa1\x37\x77\x10\x42\x28\x62\x38\x07\xd3\x87\x50\x74\x00\x21\x09\x67\xd1\x1d\x84\xae\xcd\xf0\x04\x0a\x60\x09\xb0\x98\x80\x8c\x1e\xa1\x9f\xed\x30\x0d\x04\xa4\x7a\x25\x41\x3c\x2e\x55\xa6\x61\x45\xa5\x04\xb1\xc1\xa5\xca\x36\x31\xa6\x74\x8b\xe3\xfd\xa6\x14\x34\xba\xbe\x83\xd0\x3b\x03\x0a\x27\x09\x51\x84\x33\x4c\x9f\x0b\x5e\x80\x50\x16\x64\x8a\xa9\x04\x33\xa0\x68\x37\x3b\x4c\x24\xa9\x7f\x6f\xad\x5a\x2a\x41\xd8\xce\x11\xdd\x5e\xf9\xe3\xa2\x40\x97\xdf\x36\x1d\x09\xa4\xb8\xa4\x4a\x77\x45\xa6\xed\x9d\x5b\x28\x17\xbb\x45\x80\x9f\x89\x1d\x66\xe4\x1f\x58\x2f\x60\x06\x78\xc3\xd3\xa5\x84\x9b\x49\x1d\x30\xd5\x76\x2c\x85\x24\x21\xc7\x4c\x91\x18\x55\x00\xea\x61\x05\x56\x0a\x84\x86\x18\xfd\xfd\xb3\x37\x0f\xee\x7d\x7d\xfd\xbf\x9f\xff\xf8\xe3\x85\xe7\xd7\x4f\xba\xa4\x14\x82\xe4\x58\x9c\xbe\xc1\x0a\x76\x5c\x9c\x3c\x24\xbd\xa9\x69\x42\x11\x2b\x29\x8d\xae\x87\xc4\xb5\x25\x53\x13\xea\xa0\xa2\xb8\x02\x5b\x4f\x01\x56\xe6\xb5\x8c\x56\x14\x1c\x70\xdc\x0c\xd1\x22\x25\x25\xa8\xc8\xfd\x7d\xdd\xa1\x57\x42\xcc\x59\x72\x7e\x8a\x6b\xb8\xef\x81\x66\x51\xd2\x96\xec\xb7\xe9\xac\xd4\x3c\x4c\xa7\xa6\xcd\xce\xaf\xfb\xa6\x14\xce\x51\x38\x50\x3a\x2b\x7a\x04\x8e\x9d\x96\xb0\xec\x75\x09\xd1\x44\x20\xb8\x29\x04\x48\x49\xd8\x0e\x1d\x33\x60\x48\x65\xd6\x12\x11\x89\x0e\x44\x92\x2d\x85\xee\xec\x81\x22\xb5\xd8\x62\x06\x10\x26\x15\xa6\xf4\xf6\x14\x71\xd6\xa5\x28\xc6\x0c\x6d\x01\x39\xf8\x90\xcc\xa1\xeb\x4e\xfb\xff\x6a\xf3\x62\x5c\xe0\x2d\xa1\xa4\xc7\xca\x9a\x44\x2c\x04\x3e\x05\xf4\xb5\x33\xb7\x1e\x42\x14\xe4\xbd\x5d\x09\xae\x37\x00\xaf\x23\x79\x03\xf1\xd4\xdb\x0f\x22\x27\x4a\xcb\x07\xde\x75\x8d\x5b\xab\xfb\x92\xed\x40\x0e\x3b\x05\xd9\x65\xe1\xa9\xa6\xb7\x99\x59\x77\x5d\x7b\xb9\x27\xcb\xa2\xe0\x42\x41\xf2\x14\xb3\x5d\x89\x77\x01\x3d\x70\x4c\x1c\x57\x83\x1a\x16\xa2\x35\xb0\x75\x4c\xf5\x71\x0c\x26\xd7\x52\x94\xa2\xe0\xf2\x36\x2b\xa8\x21\x9c\x91\x6c\x29\x81\x29\x82\xe9\x46\x82\x38\x90\xb8\x05\x1c\x39\x3b\x90\x6c\x80\x65\x98\xc5\x90\x03\x53\xdd\x5e\xcc\x30\x3d\x29\x12\xf7\x26\x6d\x21\xc3\x07\xc2\x05\xa6\x1b\x9c\x1c\xb4\x19\x91\x03\x35\x2c\x40\x48\x6d\x83\xfa\xce\xd3\xae\x01\x28\x6c\xf2\xbd\xea\xcd\x91\xb9\xf4\xb4\x42\x8e\x09\xf5\xb4\x07\x9a\x13\xac\xf0\x66\x2b\xf8\x5e\x93\x34\xb1\x65\x46\x58\x6f\xb1\x61\x6e\xfe\xf9\xb6\x2b\x01\x0a\xaa\x67\x24\x71\x1c\x83\xec\xed\x40\xcc\x85\x80\xd8\xa3\xb8\x5c\xa8\xca\x04\x4c\xac\x9d\x24\x5a\x30\xd4\xe9\x65\x81\xe3\xdb\x08\x6d\x05\x07\x49\x0b\x68\x0e\x33\x5a\xe1\xeb\x34\xc8\xf6\x30\xce\xe0\x59\xda\x63\x5a\xdb\x33\xf8\x22\xdd\x0e\xdb\x92\x2e\x73\x51\x2f\x12\x6e\xb5\xa6\x84\x2a\x10\x32\xea\xb4\x5f\x77\x87\x85\xdc\xa8\xeb\xcd\x31\xc3\x3b\x48\x3c\x5d\x2d\x5e\x6c\x39\xa7\x80\xd9\x80\x80\x36\x53\x5e\x67\xa0\x32\x10\x48\x65\x44\xf6\xb8\xa3\x7d\xab\xc3\x83\xb8\x40\x39\x2e\x8a\x9e\x37\xab\x48\x6d\x42\x40\x13\x0a\x44\xbd\x21\xef\x06\x1c\x30\xdc\x1a\xa5\xdd\xe3\x7f\xfb\xa4\x0f\xf7\x13\x19\xb8\x93\xd8\x7b\xb1\xf4\xd9\xb0\x37\xe1\xf6\x18\xf6\x04\x64\x2c\x48\xa1\xba\x61\xf8\xd9\x88\x68\x83\xf7\x4c\xf5\xc6\x1f\x21\x52\x2b\x41\x1d\x25\xb3\x1b\x86\xcc\xa4\xb2\x82\xec\x99\x36\x54\x6e\x0f\xd2\x81\x9a\x7b\xf1\x3a\x34\xfe\x61\x23\xda\xdc\x46\xe6\x9b\xab\x59\xc3\x45\x8e\x55\x9f\x85\x68\xa8\xc8\x6e\xfc\x9c\xd0\xb9\x33\x61\x54\xff\xbb\xdc\xf0\xf7\xce\x91\xa4\x6a\xdc\xe4\x4e\xa1\x11\x5e\xf8\x1d\x4e\x7f\x84\xf6\xa9\x41\x00\xc6\x24\x30\x25\x38\xa5\x20\x36\x71\x29\x15\xcf\x41\x6c\x88\xcf\xdc\xd4\x33\x30\x4b\x04\x27\x49\x3b\x50\x98\x39\x63\x7c\x54\x4a\x04\x2c\x00\x4a\xb8\x5c\x38\xfa\x00\x2c\xe1\x53\xab\xcb\x49\x2c\xb8\xe4\xa9\x5a\x00\xbc\x99\x53\x94\x5b\x4a\x64\x36\xc9\x43\xc1\xf7\xe5\xb2\xc1\xf3\xa9\x29\x19\xf9\xa9\x04\x1d\x30\x1d\x65\x40\x09\xdd\x48\x1d\xb4\x15\x19\x67\x03\xeb\x59\xfd\x5c\x7b\xdb\x07\x06\xcb\x81\x73\xba\xf9\x61\x14\xc3\x21\xbb\x85\x6a\x08\x7c\x1c\x63\x8e\xcc\xf0\xc3\x89\xfe\x2f\xbe\xfc\x6a\x62\xc4\x97\x0f\xbf\x58\xc8\x5b\x4f\xeb\xb0\xad\xdf\xd2\xfd\xbb\xfd\x57\x67\xaf\xde\x53\x5c\xe5\xb5\x51\xd1\x01\x0b\x82\xb7\x14\xfe\x1b\xc3\x2d\x25\xca\xdf\xa3\xad\x51\xec\xff\x41\xd1\x56\xd0\x9b\xdf\x92\xc6\x71\x0f\x3e\x66\xa0\x46\xbc\xf6\x12\x7f\xbd\xc4\x53\x4f\xf9\xe8\xd9\xde\x79\xae\x5f\x9e\xe1\x91\x97\xf9\xe2\x25\x5e\x78\xa6\xff\x9d\xed\x79\xe7\xf8\xdc\x31\x6f\xdb\xf7\x05\x9e\x94\x20\xec\x5b\xcf\x23\xa6\x41\x7f\xda\x56\x26\xbf\xcb\x1c\x15\xe5\x90\x97\x0d\xfb\xd7\x31\xcf\x1a\xf4\xa9\xd3\x1c\xac\xfd\xd1\x28\x0f\x83\x09\x4e\xc3\xc3\x1f\x1c\x20\x94\x40\x4a\x18\x09\xd9\xa0\x89\x3c\x27\xa2\x3c\x1e\x9e\xa3\xd5\xbd\x3e\x2b\xeb\xcd\x71\x16\x66\x38\xd3\xf9\x4d\xd8\x71\xa0\xd9\x21\x5c\xc3\xac\xef\x71\x0e\x88\xa7\xe6\x94\xde\x1f\x11\xd8\x1f\x6f\x4c\xd9\xf0\xe8\x5c\xe4\x3c\x75\x00\x07\x24\x05\x26\x4e\x44\x91\xd1\x91\xb0\x84\x87\xc3\xc8\x28\xe6\x7c\x4f\x46\x12\xb8\x04\x2b\xfc\x14\x9f\x40\xf8\xa3\x44\x5f\x8c\xe8\xe7\xd4\xdb\x63\x38\xf0\x9e\xbe\x7d\xea\x8f\xaf\xd8\xf5\xdd\xeb\x2b\x94\x80\xc2\x84\xfa\x8e\x0c\xd0\xac\x74\x7e\x4a\xd4\xcd\x88\x3d\x9c\x02\x1c\x08\xa0\x5d\x9a\xd7\xcf\xcb\xec\xe7\x88\x1c\x5a\x90\xc8\xb4\x79\x59\x01\x46\x8a\xa3\x94\xb0\xa4\x23\x7f\xe8\x48\x54\x46\xec\x65\xd6\x77\xaf\xaf\xc6\x00\x4e\x26\x36\x08\x45\x19\xe0\x64\x34\xf5\xd3\xa2\x49\x31\xc9\x65\x28\x3b\x09\xe5\x27\xc1\xec\xcf\x6d\xe2\xb9\xb9\xb6\x87\x53\x5f\x57\xfb\xbc\x0a\xd1\xf9\x61\xf2\xab\xfa\xf7\xc0\x75\x58\xc6\x85\xfa\xd6\x1b\x0c\x4f\x57\x00\x98\xc9\xed\x58\x17\xa9\x0c\x2b\xa4\x33\x14\x2c\x24\x22\x0c\x49\xc0\x22\xce\x90\x00\x59\x52\x25\x11\x66\x09\xda\x61\x1d\x1c\x76\xee\x1f\x83\xd5\x0d\x56\xb7\x21\x59\x4b\x60\x35\x3f\x4c\x23\x6f\x2e\x68\x9d\x21\x41\x32\x16\x00\x73\x6a\x2f\xcc\x9d\xa5\xb9\xe9\xfd\x9e\xab\xe5\x55\x18\xcd\x74\xc4\xb8\x02\x24\x33\x7e\x34\xda\xa7\x09\x2a\x25\x08\x7b\x7f\x8c\xe5\x9e\xb0\x9d\x8e\x7f\x10\x2e\x55\xc6\xc5\xfc\xe2\x90\x06\xc3\x6d\xae\x08\x5b\x40\xd6\x5d\x3b\xd5\xe0\x9e\xd7\xa0\x22\xbf\x30\x82\x2a\x8b\x57\x2f\x9e\xae\x28\x43\x51\x65\x81\xf4\xcc\x0e\xbc\x8c\xe7\x50\xe0\x1d\xac\x01\x59\xcd\x1d\x42\xad\xd2\x9a\x97\xf6\xf6\x77\x0d\xf0\x0a\x44\x75\x83\x3c\x44\xe2\xea\x39\x9e\x73\x4a\xe2\xd3\x1a\x14\x0e\x00\x2a\x0c\x84\x21\x02\xa9\xb0\x2a\xe5\x2a\x66\x9b\x99\x43\x88\x8a\xaf\x02\xa7\x40\xe4\x52\xdb\x50\x77\x5d\x3c\x04\x9c\xf0\x78\x0d\xe0\x84\xc7\x65\x0e\x4c\x59\xaf\x36\x80\x4a\xf9\x8e\xaf\x2c\x81\xe9\x97\xea\x18\x50\xf5\x30\x6f\xc4\x11\x51\xc2\xf6\x75\x31\x4e\x3d\x76\x56\xe5\x4c\x55\x1a\xa7\x5e\x34\x90\x5b\x8a\x77\x24\x89\xab\x8c\xcb\x80\xec\x32\xd5\x0e\x9b\xaa\x26\xdd\x6b\xc7\xd5\x6e\xa0\x06\x1f\x2a\xcc\xa9\x96\xbb\xbc\x0e\x46\xf3\x03\xd9\xbf\x3a\x03\x26\x8b\x6f\x0c\x93\xd6\x60\x4c\x09\x85\x6e\x4a\xd2\x06\x5b\x71\xc8\x0b\x97\x30\x05\xbb\x5e\x30\xd2\x5b\x8a\x9d\x1f\x58\xca\x03\x2f\xca\x9a\xef\x6b\x71\x3a\x00\xd3\x48\xbd\x96\xb4\x10\x70\x20\x70\x5c\x69\xf6\x87\xc5\x73\x0e\xda\x1c\xfb\x3f\x72\xd3\x7e\x99\xe3\x5d\xe7\x72\x3d\x10\x9c\x77\x94\xa5\x13\x63\xcf\x8c\xab\x27\x54\x66\x42\x69\xc6\xd4\xa6\xbb\xcb\xc1\x90\xdd\xab\x3c\x53\x11\x66\x8f\x51\x1e\x0d\x0a\x1f\x1d\x76\xa2\x5e\x9f\x1e\x2d\x45\x3e\xd4\xa8\x3e\x12\x9f\x56\x4d\xc8\xf8\x10\xcd\x50\xb5\x02\x72\xde\xc7\xee\x55\xb0\xc5\xe8\x3d\x5a\x16\xc6\x3f\xa1\x73\x31\x67\x0a\xc7\xb7\x29\xee\xa9\x21\xac\xd3\xb3\x90\x36\x0d\x4e\x78\xdd\xf9\xed\x6d\x14\x6c\x44\xf2\x7d\xc7\xd5\x73\x65\xaf\xc5\x05\xdf\xc1\x74\x28\xb9\x8c\x14\xc4\x19\x23\x31\x1e\x1e\x4a\x47\xb9\x31\x66\xbe\x7c\x2e\x4a\x09\xc3\x2c\xf6\xdc\x0d\x49\x88\x4b\xd1\xa9\x68\x42\xbd\x8c\xf3\x9d\x87\x9f\xe7\x58\xb4\xef\x64\xbd\x39\x5c\xed\x6f\xdc\xb4\x50\xc2\x4d\x41\x04\xbc\x92\x20\xae\xf8\x1e\xbc\x79\xc0\xe0\x82\xaa\xa1\xec\x89\x99\x6d\x93\x11\x65\xe7\x7b\xb2\x0e\x23\x1b\x1d\xac\x02\x52\x01\x32\xbb\x64\x0a\xc4\x01\xd3\xc5\xa1\x9b\x9b\x8f\x48\x05\xa0\x09\x56\x9a\x4b\xae\xcf\xde\x3c\xbc\xf7\xf5\xb5\x29\x22\xff\x9f\xcf\xb3\x19\xe9\xd0\xf0\x31\xc1\x02\x56\xbc\xb0\x93\x2d\x2f\x74\x22\x36\x8f\x13\xa5\x43\xf6\x8d\x7b\xaa\x30\x33\x90\x6d\x76\xbc\x14\x24\xc0\xa4\x67\x1a\x2e\xaa\xde\x40\x98\x00\x77\x0e\x0f\x12\x22\x20\x56\xcf\xd8\xab\x22\xc1\xfe\xac\x75\x8c\x09\x76\xb6\x4e\x9b\x4b\x3b\x7f\x16\x17\x8e\xb0\xcd\x38\xdf\xdf\xa2\xe2\xbc\x82\x30\x11\x65\x97\xa2\xad\x3c\x5a\x8f\x05\x74\x6c\x24\x1c\x80\x29\xb9\x32\x10\x0f\x45\xca\x1a\xeb\x82\xa8\xd5\xbf\xb9\xdd\x35\xbf\xb6\xeb\x45\xaf\x5e\x5c\xfa\x63\x5a\xb7\xb4\x35\xc1\xf2\x9f\xb0\x84\xaf\xfe\xff\x1e\xb0\x98\x27\x90\x54\xac\x45\x16\xa2\x16\xf0\x04\x29\x8e\x24\xd9\x31\xe4\xd8\xd5\x81\xa3\x4d\x15\x30\xf5\x44\x4f\xd7\x48\xb4\xc4\x18\x88\x0b\x83\xfd\xee\x69\xc6\x9a\x85\x54\x4c\xea\x40\x42\xf6\x70\xb1\x5e\x87\xee\x04\xa6\x48\x8c\x15\x54\xa7\x2b\x7d\x69\x9a\x47\xb0\x63\x86\x9f\xd2\xa0\x8b\xf7\x93\xec\x63\xac\xbf\xc2\x6d\xc2\x8b\x84\x3c\x62\xe1\xf5\x78\xae\x56\xfe\xe0\xce\x7a\x2f\x62\x01\x58\x79\xca\x09\x06\x03\xd3\x92\xa6\x64\xf0\x6c\xc1\xee\x63\x51\x5c\xf8\x5f\x35\xd4\xdd\xce\x9e\xfb\x3b\x4b\x36\x32\xdb\xbe\x0b\x08\xd3\xe9\xfa\x6d\x15\x73\x32\xe2\xaf\xbd\x3b\x9a\xe3\x9b\xbf\x3d\x7f\xb9\x28\x4b\x6b\x94\x97\x30\xf5\x7f\x5f\x84\x44\xe8\xe1\x03\x7f\x5e\xa8\x68\x58\x80\xc6\x6e\x3f\x1a\x09\xba\x7a\xfa\x52\x47\x0b\x29\xd9\x95\xc2\xf7\x90\x22\x7c\x8b\x10\x11\x66\x02\x1b\xdf\xed\xd9\xc0\xf8\x77\xfa\xfd\x27\xca\xfe\x88\xe3\xc8\xc5\x3e\xa5\xfc\xb8\x2e\x0c\x6e\xee\x2c\x11\x4f\x51\x05\x0b\x59\xd6\x9c\x3b\x30\x1e\x14\x11\x0d\x0b\x88\x22\x12\xf7\x39\x9c\xb2\xf7\x14\x39\x7b\xeb\x71\xe6\x06\x91\x59\x99\x63\x76\x4f\x00\x4e\xcc\x1d\x84\xbd\x44\x4f\x89\xbd\x39\x37\x46\xcf\xf1\x25\x18\xc4\x7a\xef\x55\xe7\xa2\x67\xad\x7b\xd4\xc7\xb1\x22\x07\xa2\x4e\x08\xcb\xe1\xa1\x76\x18\xbf\xe1\xf4\x5a\xfc\x55\xba\x5a\xd1\xa0\x81\x19\x2f\x96\xf1\x63\x18\xe5\xad\x72\x15\x3d\xac\x42\xe7\xbd\x96\x6f\x7b\x14\xec\x78\x32\xd7\x7c\x07\xc6\xb7\xf5\x6b\x66\x6e\x92\xae\x67\xaa\x5e\x59\x5a\xb2\xb8\xba\x16\xdc\x02\x82\x1b\x88\xcb\x9e\xa5\xed\xe2\xe3\xe6\x96\x25\xec\xc7\xa6\x6e\x79\xbb\xe8\x1d\xb4\x5a\x88\x7b\x16\xa1\x4f\xfa\xfc\x6b\xd7\x89\xca\x3e\x85\xe5\x7e\xf3\x53\x09\xe5\xea\xd2\xab\x45\x35\x5e\xd5\xb2\x36\x96\xbf\x84\xb3\x8d\x22\x39\xf0\x72\xa2\xa4\xc6\x7f\xca\x81\x06\x8e\xaa\x17\x9c\x0d\x48\x7c\x30\x9f\x42\x51\x7e\xbc\xb4\x99\x5d\xfb\xd8\x88\xc3\x44\x6d\x52\x2e\x36\x31\x66\x31\x50\x1a\xbe\xbd\x9f\x53\x31\xea\xcb\xb2\xc6\xb0\xcb\x38\x83\xa4\xa4\xb0\x51\x7c\x13\x53\x2e\xe1\x63\x63\x4f\x9b\x40\xa9\xb0\x50\x1f\x1d\x81\x96\xa8\x8f\x94\x7d\x19\x60\xa1\xb6\x80\x7f\x6b\xb6\x85\xc3\xc5\x6e\x74\x8a\x05\x1e\xc9\x71\x46\x93\xa8\xb6\x6b\x30\x70\xc0\xf3\x54\x28\xf8\x4c\x68\xaa\x82\x6e\xbc\x50\xc8\x5f\x6f\x1e\xaa\x38\xaf\xb0\xf5\x9a\x07\xb5\x42\x8b\xaa\x84\xa6\xea\x83\xc2\xe5\xdc\xf3\x2a\x5b\x82\xa1\xa4\xad\x1d\x6d\x47\x94\x2e\xe8\xa9\x37\x61\x58\xe2\xe2\x29\xc2\x19\xab\xdb\x5b\x46\x9f\x0e\x23\xef\xa5\x82\x00\x4b\xe8\x09\xb5\xa3\xbe\x85\x14\x8d\xbc\x93\x5a\x46\x51\x3b\x14\x5c\x48\x43\xa3\x53\xb7\x27\xc3\xc1\x42\x07\x4c\xcb\x59\xf4\xac\xd4\x62\x5e\xaa\xa2\x1c\x9c\x84\xac\x50\x63\x0b\xc8\xd2\xfb\xbb\x26\x77\x49\xfc\x90\x9a\x6c\xf7\xe1\xa3\x51\xe3\x25\xe4\xbc\x17\x1d\x5e\x42\xc0\x7b\x57\xe0\x20\x31\x2b\xb5\xd7\x1e\x1e\x85\x93\xc2\x45\x99\x95\x12\x65\x3b\x08\x9d\x3a\x1c\xd2\x41\xc6\xa2\x83\xa1\xb3\x9c\xf3\x0c\x4f\x75\xba\x4a\x7d\xd6\x53\x9c\x35\xc7\x28\xdd\xd7\x48\xd1\x55\x06\x68\x0f\x27\x94\x12\xa0\x09\x22\x12\x1d\x33\x10\x60\x84\x21\xc1\x0a\xa3\x23\xa1\x54\x27\xe3\x12\x1f\xcc\x19\xf7\x05\xba\xca\x88\x44\x79\x29\x95\x6e\x76\x6a\x5e\x80\xb0\x10\x2e\x7e\x64\xe8\xcf\x5c\x20\xb8\xc1\x79\x41\xe1\x2e\x22\xa9\x01\xff\x47\xf4\x69\x55\x88\x26\x3f\xd5\xd0\x99\xab\x63\x6d\x09\x9f\x25\xa1\x8b\x90\x30\x43\xc6\x45\x3d\x79\xe4\x80\x65\x4d\x19\x84\x87\x1b\xd6\x3a\x08\x5b\x12\xe1\x0e\x03\x88\x44\x31\xcf\x0b\xce\x80\x8d\x9c\x6a\xad\x39\xe2\xf1\x10\x60\xac\x83\xdb\xf4\x53\xcd\x8f\xfa\xae\x04\x28\xc4\x0a\x1d\x33\x12\x67\x0d\x51\xba\x47\x00\x4b\xb4\x9d\xb5\xac\x4d\x85\xb9\x3b\x49\x2e\xd0\xa5\x42\x31\x66\x8c\x9b\x0d\xc3\x0c\xc1\x0d\x91\xca\xd4\x79\x1a\x86\x2f\xb9\xfb\xf6\x3f\xb4\xaf\x96\xd6\x69\x0e\x9f\x12\xc9\x8c\x1f\x9f\xb1\x1f\x86\xdf\x62\x9a\x4c\x4b\xfb\xdc\xba\x4c\x8d\x51\xb8\xdb\x17\x1f\x5b\x71\x5b\x71\x22\xe6\x8c\xd9\xcf\x89\x20\x7e\x00\x61\x3e\x02\x15\x3c\x3d\xeb\xa7\xba\x43\xd2\x7f\xd0\x32\xfb\x8c\x3d\x49\xc8\x48\x49\xc8\x6c\xfa\xad\xb2\xb7\x16\x60\x55\xc2\x2c\xc3\x6d\x99\x3d\xd7\x74\x8b\x81\x84\xa8\x7e\xd9\xf0\x60\x11\x7d\x4b\xd9\x59\x83\x86\xe0\x7d\xfc\x73\x06\xe2\x1b\x39\x8b\x33\xcc\x76\x83\x5b\x94\xb9\x34\x0a\x7f\x29\xd5\xd9\x04\x64\x0b\xa8\x42\xa1\x35\x27\xc3\x07\x40\xd8\xb2\x3e\xac\xde\x39\x61\x4f\x81\xed\x56\x14\x22\x8d\xe7\xbb\xb3\x6a\x90\x72\x7c\xf3\xdb\x21\x6f\x2a\x21\x16\xda\xb6\x19\xf5\x63\xa6\xf4\x2f\xe3\x34\x01\x71\x06\xdb\xd9\x82\x86\x14\xdc\x28\x5b\x81\xdf\x36\x0a\xee\x43\x69\x44\x36\x2e\x0f\xf2\x62\x78\x5a\x3e\x8b\xf8\x50\x40\xd6\x04\x1b\xa1\x8f\xf1\xf9\x97\xf0\x6d\x27\x28\xab\x44\xb5\xf1\x94\xb5\x37\x32\xb4\xdf\x45\x5b\x48\x79\x55\x34\x63\x0a\x58\xb0\xb1\x73\x17\xe8\x2f\xf8\xa0\x2d\x3c\x46\xc9\x10\xa2\x36\x82\x82\x24\x16\x6e\x9f\x61\xe1\xb5\xca\x02\x28\x8d\x33\x88\x47\xca\xfd\xe6\x69\xa5\x89\x20\x24\x28\xe3\x84\x0c\x49\xc0\x4c\xd2\x60\x50\x20\x83\xa3\x76\x63\x26\xaa\x58\x69\x46\xf2\x92\x2a\x52\x9c\xc5\xd4\x59\x33\x52\x01\x74\x69\x64\xf5\xa1\x3d\x9c\x24\x36\x56\x69\xf6\x46\x07\x42\x57\x59\x3d\xb0\x2d\x81\x58\x6a\x2f\x6c\x5c\x29\x72\x6f\x77\x1e\x3f\xbf\x34\x0f\x56\x30\x43\xbf\xfe\xf2\xcf\xc7\x49\x82\x1e\x33\xae\x32\x10\xbf\xfe\xf2\x2f\xb4\x2d\x95\xe2\xac\x96\x06\xf7\xad\xc1\x0e\x83\x10\xa6\x94\x1f\x35\x3b\x8d\xc3\x13\x50\xbf\xec\x6b\x62\x4c\x3b\x52\xfa\x22\x9a\x95\xde\x50\x47\x66\x4b\x63\xfb\xf1\xab\x90\xea\x0b\x58\x9e\x44\xd5\x25\xee\xe3\xa9\xcf\xc4\x69\x80\x7f\x77\xaf\x32\x90\x80\xb0\x68\x29\x9a\x6c\xd9\x0d\xed\x81\x4d\xcc\x05\x89\xe5\x79\xb5\xc5\xfe\x04\x6b\xec\xdb\x41\xd3\x07\x0a\x6e\xd4\xac\x97\x84\x78\x0b\xe1\xcf\xd9\x44\x03\x67\xd6\xfc\x7c\xf0\x87\x84\x96\x96\x25\xef\xe1\x02\x03\xc3\x0f\xee\x2c\x37\xce\x81\xc2\xdb\x3e\xf3\xad\x9c\x27\x67\x7f\x2b\x83\x0f\x28\x67\x25\xec\x3d\x61\x7d\xa2\xad\x3c\xc2\xe8\xbb\x97\xcf\xbe\x47\x8f\x8d\x0d\x51\x5c\x3b\x00\x13\xe7\xcb\x8c\x97\x34\xd1\x22\x6b\xbd\xbe\x96\xd9\xb6\xb9\xe1\x69\x55\xde\x60\x5e\x0e\x9a\x5c\x27\x41\xad\x4d\x0c\x90\x10\xbc\x5e\xf4\xaf\x79\x58\xa4\x76\x8e\x25\x97\x82\x5a\xaa\xb1\x4d\x4d\x25\x2f\x45\x0c\xda\x84\xb6\x98\x61\xd7\xed\xb2\x54\x67\x9e\xab\xd4\xa9\xe0\x45\x49\xb1\xd2\xa1\xde\x4b\x9b\x45\x51\x22\x95\x85\x09\x37\x3a\xc2\xc1\xd4\xc2\xf2\x9e\x12\x3a\xea\xc2\x25\x75\x33\x19\x36\xff\x2c\x25\x74\xc2\x71\x47\xff\x7b\x77\xe7\xdf\x01\x00\x00\xff\xff\x93\x57\x5e\x5e\x6c\x5b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas"].(os.FileInfo),
	}
	fs["/schemas"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/manifest.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
