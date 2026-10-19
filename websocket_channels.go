package slack

// ChannelCreatedEvent represents the Channel created event
type ChannelCreatedEvent struct {
	Type           string             `json:"type"`
	Channel        ChannelCreatedInfo `json:"channel"`
	EventTimestamp string             `json:"event_ts"`
}

// ChannelCreatedInfo represents the information associated with the Channel created event
type ChannelCreatedInfo struct {
	ID        string `json:"id"`
	IsChannel bool   `json:"is_channel"`
	Name      string `json:"name"`
	Created   int    `json:"created"`
	Creator   string `json:"creator"`
}

// ChannelJoinedEvent represents the Channel joined event
type ChannelJoinedEvent struct {
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
}

// ChannelInfoEvent represents the Channel info event
type ChannelInfoEvent struct {
	// channel_left
	// channel_deleted
	// channel_archive
	// channel_unarchive
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	User      string `json:"user,omitempty"`
	Timestamp string `json:"ts,omitempty"`
}

// ChannelRenameEvent represents the Channel rename event
type ChannelRenameEvent struct {
	Type      string            `json:"type"`
	Channel   ChannelRenameInfo `json:"channel"`
	Timestamp string            `json:"event_ts"`
}

// ChannelRenameInfo represents the information associated with a Channel rename event
type ChannelRenameInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
}

// ChannelHistoryChangedEvent represents the Channel history changed event
type ChannelHistoryChangedEvent struct {
	Type           string `json:"type"`
	Latest         string `json:"latest"`
	Timestamp      string `json:"ts"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelMarkedEvent represents the Channel marked event
type ChannelMarkedEvent ChannelInfoEvent

// ChannelLeftEvent represents the Channel left event
type ChannelLeftEvent ChannelInfoEvent

// ChannelDeletedEvent represents the Channel deleted event
type ChannelDeletedEvent ChannelInfoEvent

// ChannelArchiveEvent represents the Channel archive event
type ChannelArchiveEvent ChannelInfoEvent

// ChannelUnarchiveEvent represents the Channel unarchive event
type ChannelUnarchiveEvent ChannelInfoEvent

// ChannelUpdateEvent represents the Channel update event
type ChannelUpdateEvent struct {
	Channel        string            `json:"channel"`
	Type           string            `json:"type"`
	Updates        ChannelUpdateInfo `json:"updates"`
	Timestamp      string            `json:"ts"`
	EventTimestamp string            `json:"event_ts"`
}

// ChannelRenameInfo represents the information associated with a Channel rename event
type ChannelUpdateInfo struct {
	IsMpim    bool `json:"is_mpim"`
	IsPrivate bool `json:"is_private"`
}

type ThreadSubscription struct {
	Type            string `json:"type"`
	Channel         string `json:"channel"`
	ThreadTimestamp string `json:"thread_ts"`
	DateCreated     int    `json:"date_created"`
	Active          bool   `json:"active"`
	LastRead        string `json:"last_read"`
}

type ThreadInfoEvent struct {
	Type           string             `json:"type"`
	Subscription   ThreadSubscription `json:"subscription"`
	EventTimestamp string             `json:"event_ts"`
}

type ThreadSubscribedEvent ThreadInfoEvent
type ThreadUnsubscribedEvent ThreadInfoEvent
type ThreadMarkedEvent ThreadInfoEvent
