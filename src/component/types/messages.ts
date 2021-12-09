import { ICEServer } from '../internal/webrtc'

/////////////////////////////
// System
/////////////////////////////

export interface SystemWebRTC {
  videos: string[]
}

export interface SystemInit {
  session_id: string
  control_host: ControlHost
  screen_size: ScreenSize
  sessions: Record<string, SessionData>
  implicit_hosting: boolean
  inactive_cursors: boolean
  screencast_enabled: boolean
  webrtc: SystemWebRTC
}

export interface SystemAdmin {
  screen_sizes_list: ScreenSize[]
  broadcast_status: BroadcastStatus
}

export type SystemLogs = SystemLog[]

export interface SystemLog {
  level: 'debug' | 'info' | 'warn' | 'error'
  fields: Record<string, string>
  message: string
}

export interface SystemDisconnect {
  message: string
}

/////////////////////////////
// Signal
/////////////////////////////

export interface SignalProvide {
  sdp: string
  iceservers: ICEServer[]
  video: string
}

export type SignalCandidate = RTCIceCandidateInit

export interface SignalDescription {
  sdp: string
}

export interface SignalVideo {
  video: string
}

/////////////////////////////
// Session
/////////////////////////////

export interface SessionID {
  id: string
}

export interface MemberProfile {
  id: string
  name: string
  is_admin: boolean
  can_login: boolean
  can_connect: boolean
  can_watch: boolean
  can_host: boolean
  can_share_media: boolean
  can_access_clipboard: boolean
  sends_inactive_cursor: boolean
  can_see_inactive_cursors: boolean
}

export interface SessionState {
  id: string
  is_connected: boolean
  is_watching: boolean
}

export interface SessionData {
  id: string
  profile: MemberProfile
  is_connected: boolean
  is_watching: boolean
}

export interface SessionCursor {
  id: string
  x: number
  y: number
}

/////////////////////////////
// Control
/////////////////////////////

export interface ControlHost {
  has_host: boolean
  host_id: string | undefined
}

// TODO: New.
export interface ControlMove {
  x: number
  y: number
}

// TODO: New.
export interface ControlScroll {
  x: number
  y: number
}

// TODO: New.
export interface ControlKey {
  key: number
}

/////////////////////////////
// Screen
/////////////////////////////

export interface ScreenSize {
  width: number
  height: number
  rate: number
}

/////////////////////////////
// Clipboard
/////////////////////////////

export interface ClipboardData {
  text: string
}

/////////////////////////////
// Keyboard
/////////////////////////////

export interface KeyboardModifiers {
  caps_lock: boolean
  num_lock: boolean
  scroll_lock: boolean
}

export interface KeyboardMap {
  layout: string
  variant: string
}

/////////////////////////////
// Broadcast
/////////////////////////////

export interface BroadcastStatus {
  is_active: boolean
  url: string | undefined
}

/////////////////////////////
// Send
/////////////////////////////

export interface SendMessage {
  sender: string
  subject: string
  body: string
}
