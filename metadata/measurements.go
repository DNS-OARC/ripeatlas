package ripeatlas

type MeasurementStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Measurement struct {
	Id                    int                    `json:"id"`
	Af                    int                    `json:"af"`
	CreationTime          int                    `json:"creation_time"`
	CurrentProbes         []int                  `json:"current_probes"`
	Description           string                 `json:"description"`
	DestinationOptionSize int                    `json:"destination_option_size"`
	DontFragment          bool                   `json:"dont_fragment"`
	DuplicateTimeout      int                    `json:"duplicate_timeout"`
	FirstHop              int                    `json:"first_hop"`
	IncludeProbeId        bool                   `json:"include_probe_id"`
	Group                 string                 `json:"group"`
	GroupID               int                    `json:"group_id"`
	HopByHopOptionSize    *int                   `json:"hop_by_hop_option_size"`
	InWifiGroup           bool                   `json:"in_wifi_group"`
	Interval              int                    `json:"interval"`
	IsAllScheduled        bool                   `json:"is_all_scheduled"`
	IsOneoff              bool                   `json:"is_oneoff"`
	IsPublic              bool                   `json:"is_public"`
	MaxHops               int                    `json:"max_hops"`
	Packets               int                    `json:"packets"`
	PacketInterval        *int                   `json:"packet_interval"`
	Paris                 int                    `json:"paris"`
	ParticipantCount      int                    `json:"participant_count"`
	ParticipationRequests []ParticipationRequest `json:"participation_requests"`
	Port                  *int                   `json:"port"`
	ProbesRequested       int                    `json:"probes_requested"`
	ProbesScheduled       int                    `json:"probes_scheduled"`
	Protocol              string                 `json:"protocol"`
	ResolveOnProbe        bool                   `json:"resolve_on_probe"`
	ResolvedIPs           []string               `json:"resolved_ips"`
	ResponseTimeout       int                    `json:"response_timeout"`
	Result                string                 `json:"result"`
	Size                  int                    `json:"size"`
	Spread                *int                   `json:"spread"`
	StartTime             int                    `json:"start_time"`
	MeasurementStatus     `json:"status"`
	StopTime              int    `json:"stop_time"`
	Target                string `json:"target"`
	TargetASN             int    `json:"target_asn"`
	TargetIP              string `json:"target_ip"`
	Type                  string `json:"type"`
}

type ParticipationRequest struct {
	Action        string   `json:"action"`
	CreatedAt     int      `json:"created_at"`
	ID            int      `json:"id"`
	MeasurementID int      `json:"measurement_id"`
	Requested     int      `json:"requested"`
	Self          string   `json:"self"`
	TagsExclude   []string `json:"tags_exclude"`
	TagsInclude   []string `json:"tags_include"`
	Type          string   `json:"type"`
	Value         string   `json:"value"`
}
