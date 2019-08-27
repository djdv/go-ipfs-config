package config

// FileSystem defines configuration options for the file system server(s)
// NOTE: Highly unstable; experimental
type FileSystem struct {
	// multi-addresses for the server and client
	//e.g. "9p":"/ip4/localhost/tcp/564", "fuse":"/mountpoint", ...
	Service map[string]string
}
