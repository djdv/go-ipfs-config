package config

type (
	// "Fuse": {...}, "P9P": {...}
	fsHostAPIKey map[string]fsNodeAPIKey
	// "IPFS": [...], "IPNS": [...], ...
	fsNodeAPIKey map[string][]fsBinding
	// "/mnt/ipfs":[], "/ip4/127.0.0.1/tcp/564":[], ...
	fsBinding map[fsTarget][]fsArguments

	fsTarget    = string
	fsArguments = []string
)

// FileSystem defines configuration options to use when interfacing with the host file system APIs.
/*
"Filesystem": {
	"Fuse": {
		"IPFS": [
			"I:": [
				"allowOther=true",
				"uid=80"
			],
			"\\localhost\ipfs": []
		],
		"IPNS": [
			"N:": [],
			"\\localhost\ipns": []
		]
	},
	"P9P": {
		"PinFS": [
			"/mnt/just/the/mountpoint": [],
			"/mnt/ipfs": [
				"/unix/tmp/with.specified.socketname.uds"
			],
			"/unix/tmp/socket/only/no/mountpoint.uds": []
		],
		"Files": [
			"/ip4/127.0.0.1/tcp/564": []
		]
	}
}
*/
type FileSystem fsHostAPIKey
