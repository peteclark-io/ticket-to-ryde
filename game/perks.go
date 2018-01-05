package game

import log "github.com/sirupsen/logrus"

const (
	apPerkType   = "ap"
	costPerkType = "discount"

	districtDiscountType   = "district"
	foodCouponDiscountType = "food-coupon"
	activityDiscountType   = "activity"
)

type Perk struct {
	Type               string  `json:"type"`
	ExtraAP            int     `json:"extra_ap"`
	Discount           float32 `json:"discount"`
	DiscountType       string  `json:"discount_type"`
	DiscountActivityID string  `json:"discount_activity_id"`
	DiscountDistrictID string  `json:"discount_district_id"`
}

func (p *Perk) applyToTurn(t *Turn) {
	if p.Type == costPerkType {
		return
	}

	t.AP += p.ExtraAP
	log.WithField("extra_ap", p.ExtraAP).Info("Applying extra AP to turn")
}

func (p *Perk) applyToActivity(a *Activity) {
	if p.Type == apPerkType {
		return
	}
}
