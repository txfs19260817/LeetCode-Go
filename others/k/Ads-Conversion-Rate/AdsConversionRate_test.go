package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_adsConversionRate(t *testing.T) {
	type args struct {
		completedPurchaseUserIds []string
		adClicks                 []string
		allUserIps               []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				completedPurchaseUserIds: []string{"3123122444", "234111110", "8321125440", "99911063"},
				adClicks:                 []string{"122.121.0.1,2016-11-03 11:41:19,Buy wool coats for your pets", "96.3.199.11,2016-10-15 20:18:31,2017 Pet Mittens", "122.121.0.250,2016-11-01 06:13:13,The Best Hollywood Coats", "82.1.106.8,2016-11-12 23:05:14,Buy wool coats for your pets", "92.130.6.144,2017-01-01 03:18:55,Buy wool coats for your pets", "92.130.6.145,2017-01-01 03:18:55,2017 Pet Mittens"},
				allUserIps:               []string{"2339985511,122.121.0.155", "234111110,122.121.0.1", "3123122444,92.130.6.145", "39471289472,2001:0db8:ac10:fe01:0000:0000:0000:0000", "8321125440,82.1.106.8", "99911063,92.130.6.144"},
			},
			want: []string{"Bought Clicked Ad Text", "1 of 2  2017 Pet Mittens", "0 of 1  The Best Hollywood Coats", "3 of 3  Buy wool coats for your pets"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, adsConversionRate(tt.args.completedPurchaseUserIds, tt.args.adClicks, tt.args.allUserIps), tt.want)
		})
	}
}
