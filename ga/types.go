package ga

// CommonData includes all necessary data
type CommonData struct {
	// general
	Version    int    `url:"v"`
	TrackingID string `url:"tid"`

	// user
	ClientID string `url:"cid"`

	// t
	HitType string `url:"t"`

	// session
	UserIP    string `url:"uip"`
	UserAgent string `url:"ua"`

	// trafficsources
	DocumentReferer string `url:"dr,omitempty"`

	// system
	ScreenResolution string `url:"sr,omitempty"`
	ViewportSize     string `url:"vp,omitempty"`
	DocumentEncoding string `url:"de,omitempty"`
	ScreenColors     string `url:"sd,omitempty"`
	UserLanguage     string `url:"ul,omitempty"`

	// content
	DocumentLink  string `url:"dl"`
	DocumentTitle string `url:"dt,omitempty"`
}
