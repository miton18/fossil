package writer

import (
	log "github.com/Sirupsen/logrus"
	"stash.ovh.net/metrics/fossil/core"
)

// FileWriter write GTS onto files
type FileWriter struct {
	SourceDir string
}

// NewWriter return an instanciated FileWriter
func NewWriter(dir string) *FileWriter {
	return &FileWriter{
		SourceDir: dir,
	}
}

func (fw *FileWriter) write(in chan *core.GTS) {

	for gts := range in {
		log.Debug(gts)
	}
}
