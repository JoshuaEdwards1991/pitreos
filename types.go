package pitreos

import (
	"time"
)

type ChunkMeta struct {
	Start       int64  `yaml:"start"`
	End         int64  `yaml:"end"`
	IsEmpty     bool   `yaml:"is_empty"`
	ContentSHA1 string `yaml:"content_sha1"` // content ?
	URL         string `yaml:"url"`
}

type FileMeta struct {
	Kind        string       `yaml:"kind"`
	MetaVersion string       `yaml:"meta_version"`
	FileName    string       `yaml:"filename"`
	Date        time.Time    `yaml:"date"`
	TotalSize   int64        `yaml:"total_size"`
	Chunks      []*ChunkMeta `yaml:"chunks"`
}

type BackupMeta struct {
	Kind          string    `yaml:"kind"`
	MetaVersion   string    `yaml:"meta_version"`
	Tag           string    `yaml:"tag"`
	Date          time.Time `yaml:"date"`
	MetadataFiles []string  `yaml:"metadata_files"`
}

type PitreosOptions struct {
	BucketName string `short:"n" long:"bucket-name" description:"GS bucket name where backups are stored" default:"eoscanada-playground-pitr"`

	BucketFolder string `short:"f" long:"bucket-folder" description:"Prefixed folder in GS bucket where backups are stored" default:"backups"`

	LocalFolder string `short:"l" long:"local-folder" description:"Folder relative to cwd where files will be backed up or restored" default:"."`

	BackupTag string `short:"t" long:"backup-tag" description:"Tag for the backup, depending on activated plugins like 'history'" default:"linux_ubuntu1604_gcc4_nohistory"`

	BeforeTimestamp int64 `short:"b" long:"before-timestamp" description:"closest timestamp (unix format) before which we want the latest restorable backup" default:"now"`
	Args            struct {
		Command string
	} `positional-args:"yes" required:"yes"`
}
