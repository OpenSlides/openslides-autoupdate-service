package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestMeeting(t *testing.T) {
	s := new(projector.SlideStore)
	slide.WiFiAccessData(s)

	wifiSlide := s.GetSlider("wifi_access_data")
	assert.NotNilf(t, wifiSlide, "Slide with name `wifi_access_data` not found.")

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"All data filled in",
			dsmock.YAMLData(`
			meeting/1:
				users_pdf_wlan_encryption: WPA
				users_pdf_wlan_password: Super&StrongP455Word
				users_pdf_wlan_ssid: RandomWiWi
			`),
			`{
				"users_pdf_wlan_encryption": "WPA",
				"users_pdf_wlan_password": "Super&StrongP455Word",
				"users_pdf_wlan_ssid": "RandomWiWi"
			  }
			`,
		},
		{
			"Password missing",
			dsmock.YAMLData(`
				meeting/1:
					users_pdf_wlan_encryption: WPA
					users_pdf_wlan_ssid: RandomWiWi
			`),
			`{
				"users_pdf_wlan_encryption": "WPA",
				"users_pdf_wlan_ssid": "RandomWiWi"
			  }
			`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "meeting/1",
				MeetingID:       1,
			}

			bs, err := wifiSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.NoError(t, fetch.Err())
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
