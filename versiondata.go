package slack

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type VersionData struct {
	VersionHash               string
	VersionTS                 string
	BuildVersionTS            string
	BuildManifestLastModified string
	FrontendBuildType         string
	AppName                   string
	Gantry                    bool
	SlackRoute                string
}

func (vd *VersionData) ToQuery() string {
	if vd == nil {
		return ""
	}
	now := time.Now()
	vals := url.Values{
		"_x_id": {fmt.Sprintf("%s-%.3f", vd.VersionHash[:8], float64(now.UnixMilli())/1000)},
		//"_x_csid": {"TODO"},
		"_x_version_ts":          {vd.VersionTS},
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

func (api *Client) FetchVersionData(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://app.slack.com/client", nil)
	if err != nil {
		return fmt.Errorf("failed to prepare request: %w", err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	addCookies(req, api.cookies)
	resp, err := api.httpclient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	z := html.NewTokenizer(resp.Body)
	var htmlAttrs []html.Attribute
Loop:
	for {
		switch z.Next() {
		case html.StartTagToken:
			tok := z.Token()
			if tok.DataAtom == atom.Html {
				htmlAttrs = tok.Attr
				break Loop
			}
		case html.ErrorToken:
			return fmt.Errorf("failed to parse HTML: %w", z.Err())
		}
	}
	if len(htmlAttrs) == 0 {
		return fmt.Errorf("no HTML attributes found")
	}
	var data VersionData
	for _, attr := range htmlAttrs {
		switch attr.Key {
		case "data-gantry":
			data.Gantry = true
		case "data-version-ts":
			data.VersionTS = attr.Val
		case "data-version-hash":
			data.VersionHash = attr.Val
		case "data-build-version-ts":
			data.BuildVersionTS = attr.Val
		case "data-build-manifest-last-modified":
			data.BuildManifestLastModified = attr.Val
		case "data-frontend-build-type":
			data.FrontendBuildType = attr.Val
		case "data-app-name-override":
			data.AppName = attr.Val
		}
	}
	if data.VersionTS != "" && data.BuildVersionTS != "" && data.BuildManifestLastModified != "" && data.VersionHash != "" && data.FrontendBuildType != "" {
		api.log.Printf("Found version data: %+v", data)
		api.versionData = &data
		return nil
	}
	return fmt.Errorf("missing version data")
}
