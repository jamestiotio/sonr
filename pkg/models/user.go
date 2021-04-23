package models

import (
	"hash/fnv"
	"log"
	"math"
	"time"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ^ Create New Peer from Connection Request and Host ID ^ //
func NewPeer(cr *ConnectionRequest, id peer.ID) (*Peer, *SonrError) {
	// Initialize
	deviceID := cr.Device.GetId()
	profile := Profile{
		Username:  cr.GetUsername(),
		FirstName: cr.Contact.GetFirstName(),
		LastName:  cr.Contact.GetLastName(),
		Picture:   cr.Contact.GetPicture(),
		Platform:  cr.Device.GetPlatform(),
	}

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return nil, NewError(err, ErrorMessage_HOST_KEY)
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := mid.ProtectedID("Sonr"); err != nil {
			log.Println(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	// Set Peer
	return &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  &profile,
		Platform: cr.Device.Platform,
		Model:    cr.Device.Model,
	}, nil
}

// ^ Returns Peer as Buffer ^ //
func (p *Peer) Buffer() ([]byte, error) {
	buf, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// ^ Checks for Host Peer ID is Same ^ //
func (p *Peer) IsPeerID(pid peer.ID) bool {
	return p.Id.Peer == pid.String()
}

// ^ Checks for Host Peer ID String is Same ^ //
func (p *Peer) IsPeerIDString(pid string) bool {
	return p.Id.Peer == pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (p *Peer) IsNotPeerIDString(pid string) bool {
	return p.Id.Peer != pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// ^ SignMessage Creates Lobby Event with Message ^
func (p *Peer) SignMessage(m string, to *Peer) *LobbyEvent {
	return &LobbyEvent{
		Event:   LobbyEvent_MESSAGE,
		From:    p,
		Id:      p.Id.Peer,
		Message: m,
		To:      to.Id.Peer,
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithContact(c *Contact, flat bool, req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		Payload: Payload_CONTACT,
		IsFlat:  flat,
		Remote:  req.GetRemote(),
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),

			// Transfer Properties
			Status: TransferCard_INVITE,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,
			Owner:     p.Profile,

			// Data Properties
			Contact: c,
		},
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithFile(tc *TransferCard, req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		Payload: tc.Payload,
		Card:    tc,
		Remote:  req.GetRemote(),
	}
}

// ^ Generate AuthInvite with URL Payload from Request and User Peer Data ^ //
func (p *Peer) SignInviteWithLink(req *InviteRequest) AuthInvite {
	// Get URL Data
	urlInfo, err := GetPageInfoFromUrl(req.Url)
	if err != nil {
		urlInfo = &URLLink{
			Link: req.Url,
		}
	}

	// Create Invite
	return AuthInvite{
		From:    p,
		Payload: Payload_CONTACT,
		Remote:  req.GetRemote(),
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),

			// Transfer Properties
			Status: TransferCard_INVITE,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,
			Owner:     p.Profile,

			// Data Properties
			Url: urlInfo,
		},
	}
}

// ^ SignReply Creates AuthReply ^
func (p *Peer) SignReply(d bool, req *RespondRequest) *AuthReply {
	return &AuthReply{
		From:     p,
		Type:     AuthReply_Transfer,
		Decision: d,
		Remote:   req.GetRemote(),
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_UNDEFINED,
			Received: int32(time.Now().Unix()),
			Preview:  p.Profile.Picture,

			// Transfer Properties
			Status: TransferCard_REPLY,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,
			Owner:     p.Profile,
		},
	}
}

// ^ SignReply Creates AuthReply with Contact  ^
func (p *Peer) SignReplyWithContact(c *Contact, flat bool, req *RespondRequest) *AuthReply {
	// Set Reply Type
	var kind AuthReply_Type
	if flat {
		kind = AuthReply_FlatContact
	} else {
		kind = AuthReply_Contact
	}

	// Check if Request Provided
	if req != nil {
		// Build Reply
		return &AuthReply{
			From:   p,
			Type:   kind,
			Remote: req.GetRemote(),
			Card: &TransferCard{
				// SQL Properties
				Payload:  Payload_CONTACT,
				Received: int32(time.Now().Unix()),
				Preview:  p.Profile.Picture,

				// Transfer Properties
				Status: TransferCard_REPLY,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,
				Owner:     p.Profile,

				// Data Properties
				Contact: c,
			},
		}
	} else {
		// Build Reply
		return &AuthReply{
			From: p,
			Type: kind,
			Card: &TransferCard{
				// SQL Properties
				Payload:  Payload_CONTACT,
				Received: int32(time.Now().Unix()),
				Preview:  p.Profile.Picture,

				// Transfer Properties
				Status: TransferCard_REPLY,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,
				Owner:     p.Profile,

				// Data Properties
				Contact: c,
			},
		}
	}

}

// ^ SignUpdate Creates Lobby Event with Peer Data ^
func (p *Peer) SignUpdate() *LobbyEvent {
	return &LobbyEvent{
		Event: LobbyEvent_UPDATE,
		From:  p,
		Id:    p.Id.Peer,
	}
}

// ^ Processes Update Request ^ //
func (p *Peer) Update(u *UpdateRequest) {
	if u.Type == UpdateRequest_Position {
		// Extract Data
		facing := u.Position.Facing
		heading := u.Position.Heading

		// Update User Values
		var faceDir float64
		var faceAnpd float64
		var headDir float64
		var headAnpd float64
		faceDir = math.Round(facing*100) / 100
		headDir = math.Round(heading*100) / 100
		desg := int((facing / 11.25) + 0.25)

		// Find Antipodal
		if facing > 180 {
			faceAnpd = math.Round((facing-180)*100) / 100
		} else {
			faceAnpd = math.Round((facing+180)*100) / 100
		}

		// Find Antipodal
		if heading > 180 {
			headAnpd = math.Round((heading-180)*100) / 100
		} else {
			headAnpd = math.Round((heading+180)*100) / 100
		}

		// Set Position
		p.Position = &Position{
			Facing:           faceDir,
			FacingAntipodal:  faceAnpd,
			Heading:          headDir,
			HeadingAntipodal: headAnpd,
			Designation:      Position_Designation(desg % 32),
			Accelerometer:    u.Position.GetAccelerometer(),
			Gyroscope:        u.Position.GetGyroscope(),
			Magnometer:       u.Position.GetMagnometer(),
			Orientation:      u.Position.GetOrientation(),
		}
	}

	// Set Properties
	if u.Type == UpdateRequest_Properties {
		p.Properties = u.Properties
	}

	// Check for New Contact, Update Peer Profile
	if u.Type == UpdateRequest_Contact {
		p.Profile = &Profile{
			FirstName: u.Contact.GetFirstName(),
			LastName:  u.Contact.GetLastName(),
			Picture:   u.Contact.GetPicture(),
		}
	}
}
