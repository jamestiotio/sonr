package topic

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	net "github.com/sonr-io/core/internal/network"
	se "github.com/sonr-io/core/internal/session"
	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"
	"google.golang.org/protobuf/proto"
)

const K_MAX_MESSAGES = 128
const K_SERVICE_PID = protocol.ID("/sonr/topic-service/0.1")

type TopicManager struct {
	ctx          context.Context
	host         *net.HostNode
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	eventHandler *pubsub.TopicEventHandler
	Lobby        *md.Lobby

	service      *TopicService
	Messages     chan *md.LobbyEvent
	topicHandler TopicHandler
}

type TopicHandler interface {
	OnEvent(*md.LobbyEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte, session *se.Session)
	OnResponded(inv *md.AuthInvite, p *md.Peer, fs *us.FileSystem)
}

// ^ Create New Contained Topic Manager ^ //
func NewTopic(ctx context.Context, h *net.HostNode, p *md.Peer, name string, isLocal bool, th TopicHandler) (*TopicManager, error) {
	// Join Topic
	topic, sub, handler, err := h.Join(name)
	if err != nil {
		return nil, err
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		topicHandler: th,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		Lobby: &md.Lobby{
			Name:    name[12:],
			Size:    1,
			Count:   0,
			Peers:   make(map[string]*md.Peer),
			IsLocal: isLocal,
			User:    p,
		},
		Messages:     make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(h.Host, K_SERVICE_PID)
	psv := TopicService{
		lobby:  mgr.Lobby,
		peer:   p,
		call:   th,
		respCh: make(chan *md.AuthReply, 1),
	}

	// Register Service
	err = peersvServer.Register(&psv)
	if err != nil {
		return nil, err
	}

	// Set Service
	mgr.service = &psv
	go mgr.handleTopicEvents(p)
	go mgr.handleTopicMessages(p)
	go mgr.processTopicMessages(p)
	return mgr, nil
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (tm *TopicManager) FindPeerInTopic(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	var p *md.Peer
	var i peer.ID

	// Iterate Through Peers, Return Matched Peer
	for _, peer := range tm.Lobby.Peers {
		// If Found Match
		if peer.Id.Peer == q {
			p = peer
		}
	}

	// Validate Peer
	if p == nil {
		return "", nil, errors.New("Peer data was not found in topic.")
	}

	// Iterate through Topic Peers
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			i = id
		}
	}

	// Validate ID
	if i == "" {
		return "", nil, errors.New("Peer ID was not found in topic.")
	}
	return i, p, nil
}

// ^ Helper: ID returns ONE Peer.ID in Topic ^
func (tm *TopicManager) HasPeer(q string) bool {
	// Iterate through PubSub in topic
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// ^ Send message to specific peer in topic ^
func (tm *TopicManager) Send(msg *md.LobbyEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = tm.topic.Publish(tm.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ Leave Current Topic ^
func (tm *TopicManager) LeaveTopic() error {
	return tm.topic.Close()
}