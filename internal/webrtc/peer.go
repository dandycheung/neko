package webrtc

import "github.com/pion/webrtc/v3"

type WebRTCPeerCtx struct {
	api            *webrtc.API
	connection     *webrtc.PeerConnection
	dataChannel    *webrtc.DataChannel
	changeVideo    func(videoID string) error
}

func (webrtc_peer *WebRTCPeerCtx) SignalAnswer(sdp string) error {
	return webrtc_peer.connection.SetRemoteDescription(webrtc.SessionDescription{
		SDP: sdp,
		Type: webrtc.SDPTypeAnswer,
	})
}

func (webrtc_peer *WebRTCPeerCtx) SignalCandidate(candidate webrtc.ICECandidateInit) error {
	return webrtc_peer.connection.AddICECandidate(candidate)
}

func (webrtc_peer *WebRTCPeerCtx) SetVideoID(videoID string) error {
	return webrtc_peer.changeVideo(videoID)
}

func (webrtc_peer *WebRTCPeerCtx) Destroy() error {
	if webrtc_peer.connection == nil || webrtc_peer.connection.ConnectionState() != webrtc.PeerConnectionStateConnected {
		return nil
	}

	return webrtc_peer.connection.Close()
}
