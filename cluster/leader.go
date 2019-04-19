package cluster

import (
	"github.com/2se/dolphin/config"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/rpc"
	"time"
)

// Failover config
type clusterFailover struct {
	// Current leader
	leader string
	// Current election term
	term int
	// Hearbeat interval
	heartBeat time.Duration
	// Vote timeout: the number of missed heartbeats before a new election is initiated.
	voteTimeout int

	// The list of peers the leader considers active
	activeNodes []string
	// The number of heartbeats a node can fail before being declared dead
	nodeFailCountLimit int

	// Channel for processing leader pings
	leaderPing chan *PingRequest
	// Channel for processing election votes
	electionVote chan *voteReqResp
	// Channel for stopping the failover runner
	done chan bool
}

func (c *Cluster) failoverInit(config *config.ClusterFailoverConfig) {
	log.Debug("cluster: initing failover")
	if config == nil || !config.Enabled {
		log.Debug("cluster: not enable failover")
		return
	}

	if len(c.peers) < 2 {
		log.Warnf("cluster: failover disabled; need at least 3 peers, got %d", len(c.peers)+1)
		return
	}

	// Generate ring hash on the assumption that all peers are alive and well.
	// This minimizes rehashing during normal operations.
	var activeNodes []string
	for _, node := range c.peers {
		activeNodes = append(activeNodes, node.name)
	}
	activeNodes = append(activeNodes, c.thisName)

	log.WithField("config_heartbeat", config.Heartbeat.Get()).Debug("heartbeat value from config")

	// Random heartbeat ticker: 0.75 * config.HeartBeat + random(0, 0.5 * config.HeartBeat)
	rand.Seed(time.Now().UnixNano())
	hb := config.Heartbeat.Get()
	hb = (hb >> 1) + (hb >> 2) + time.Duration(rand.Intn(int(hb>>1)))

	voteAfter := config.VoteAfter
	failAfter := config.NodeFailAfter

	if voteAfter <= 0 {
		voteAfter = defaultVoteAfter
	}

	if failAfter <= 0 {
		failAfter = defaultFailAfter
	}

	c.fo = &clusterFailover{
		activeNodes:        activeNodes,
		heartBeat:          hb,
		voteTimeout:        voteAfter,
		nodeFailCountLimit: failAfter,
		leaderPing:         make(chan *PingRequest, voteAfter),
		electionVote:       make(chan *voteReqResp, len(c.peers)),
		done:               make(chan bool, 1)}

	log.WithField("heartBeat", hb).Info("cluster: failover mode enabled.")
}

func (c *Cluster) sendPings() {
	rehash := false

	for _, node := range c.peers {
		unused := false
		err := node.call("Cluster.Ping", &PingRequest{
			Leader: c.thisName,
			Term:   c.fo.term,
			Nodes:  c.fo.activeNodes}, &unused)

		if err != nil {
			node.failCount++
			if node.failCount == c.fo.nodeFailCountLimit {
				// peer failed too many times
				rehash = true
			}
		} else {
			if node.failCount >= c.fo.nodeFailCountLimit {
				// peer has recovered
				rehash = true
			}
			node.failCount = 0
		}
	}

	if rehash {
		var activeNodes []string
		for _, node := range c.peers {
			if node.failCount < c.fo.nodeFailCountLimit {
				activeNodes = append(activeNodes, node.name)
			}
		}
		activeNodes = append(activeNodes, c.thisName)
		c.fo.activeNodes = activeNodes
		log.Info("cluster: initiating failover rehash for peers. active nodes: ", activeNodes)
	}
}

func (c *Cluster) electLeader() {
	// Increment the term (voting for myself in this term) and clear the leader
	c.fo.term++
	c.fo.leader = ""

	log.Info("cluster: leading new election for term ", c.fo.term)

	nodeCount := len(c.peers)
	// Number of votes needed to elect the leader
	expectVotes := (nodeCount+1)>>1 + 1
	done := make(chan *rpc.Call, nodeCount)

	// Send async requests for votes to other peers
	for _, node := range c.peers {
		response := VoteResponse{}
		node.callAsync("Cluster.Vote", &VoteRequest{
			Node: c.thisName,
			Term: c.fo.term}, &response, done)
	}

	// Number of votes received (1 vote for self)
	voteCount := 1
	timeout := time.NewTimer(c.fo.heartBeat>>1 + c.fo.heartBeat)
	log.WithField("timeout", c.fo.heartBeat>>1+c.fo.heartBeat).
		WithField("expectVotes", expectVotes).
		WithField("failover term", c.fo.term).
		Debug("cluster: start another leader election")

	// Wait for one of the following
	// 1. More than half of the peers voting in favor
	// 2. All peers responded.
	// 3. Timeout.
	for i := 0; i < nodeCount && voteCount < expectVotes; {
		select {
		case call := <-done:
			if call.Error == nil {
				if call.Reply.(*VoteResponse).Result {
					// Vote in my favor
					voteCount++
				} else if c.fo.term < call.Reply.(*VoteResponse).Term {
					// Vote against me. Abandon vote: this node's term is behind the cluster
					i = nodeCount
					voteCount = 0
				}
			}

			i++
		case <-timeout.C:
			// break the loop
			i = nodeCount
		}
	}

	if voteCount >= expectVotes {
		// Current node elected as the leader
		c.fo.leader = c.thisName
		log.WithField("fo.term", c.fo.term).Info("Elected myself as a new leader")
	}
}
