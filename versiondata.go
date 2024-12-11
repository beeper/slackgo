package slack

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type VersionData struct {
	VersionHash          string
	VersionTS            int64
	BuildManifestLastMod int64
	BuildNumber          int64
	FrontendBuildType    string
	AppName              string
	SlackRoute           string
	Gantry               bool
}

func (vd *VersionData) ToQuery() string {
	if vd == nil {
		return ""
	}
	now := time.Now()
	vals := url.Values{
		"_x_id": {fmt.Sprintf("%s-%.3f", vd.VersionHash[:8], float64(now.UnixMilli())/1000)},
		//"_x_csid": {"TODO"},
		"_x_version_ts":          {strconv.FormatInt(vd.VersionTS, 10)},
		"_x_frontend_build_type": {vd.FrontendBuildType},
		"_x_desktop_ia":          {"4"},
		"fp":                     {"aa"},
		"_x_num_retries":         {"0"},
	}
	if vd.Gantry {
		vals["_x_gantry"] = []string{"true"}
	}
	if vd.SlackRoute != "" {
		vals["slack_route"] = []string{vd.SlackRoute}
	}
	return "?" + vals.Encode()
}
