package peerleecher

import (
	"time"

	"github.com/galaxy-digital/spesbas-base/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}
