package raft

import (
	"io"
	"time"
)

// Config provides any necessary configuraiton to
// the Raft server
type Config struct {
	// Time in follower state without a leader before we attempt an election
	HeartbeatTimeout time.Duration

	// Time in candidate state without a leader before we attempt an election
	ElectionTimeout time.Duration

	// Time without an Apply() operation before we heartbeat to ensure
	// a timely commit. Should be far less than HeartbeatTimeout to ensure
	// we don't lose leadership.
	CommitTimeout time.Duration

	// MaxAppendEntries controls the maximum number of append entries
	// to send at once. We want to strike a balance between efficiency
	// and avoiding waste if the follower is going to reject because of
	// an inconsistent log
	MaxAppendEntries int

	// If we are a member of a cluster, and RemovePeer is invoked for the
	// local node, then we forget all peers and transition into the follower state.
	// If ShutdownOnRemove is is set, we additional shutdown Raft. Otherwise,
	// we can become a leader of a cluster containing only this node.
	ShutdownOnRemove bool

	// TrailingLogs controls how many logs we leave after a snapshot. This is
	// used so that we can quickly replay logs on a follower instead of being
	// forced to send an entire snapshot.
	TrailingLogs uint64

	// SnapshotInterval controls how often we check if we should perform a snapshot.
	// We randomly stagger between this value and 2x this value to avoid the entire
	// cluster from performing a snapshot at once
	SnapshotInterval time.Duration

	// SnapshotThreshold controls how many outstanding logs there must be before
	// we perform a snapshot. This is to prevent excessive snapshots when we can
	// just replay a small set of logs.
	SnapshotThreshold uint64

	// EnableSingleMode allows for a single node mode of operation. This
	// is false by default, which prevents a lone node from electing itself
	// leader.
	EnableSingleNode bool

	// LogOutput is used as a sink for logs. Defaults to os.Stderr.
	LogOutput io.Writer
}

func DefaultConfig() *Config {
	return &Config{
		HeartbeatTimeout:  200 * time.Millisecond,
		ElectionTimeout:   250 * time.Millisecond,
		CommitTimeout:     10 * time.Millisecond,
		MaxAppendEntries:  64,
		ShutdownOnRemove:  true,
		TrailingLogs:      1024,
		SnapshotInterval:  150 * time.Second,
		SnapshotThreshold: 8192,
	}
}
