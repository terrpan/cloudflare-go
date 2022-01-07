package cloudflare

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// UserSubscriptionsResponse is the response from the UserSubscriptions endpoint.
type UserSubscriptions struct {
	ID                 string                             `json:"id,omitempty"`
	Prodcut            UserSubscriptionsProduct           `json:"product,omitempty"`
	RatePlan           UserSubscriptionsRatePlan          `json:"rate_plan,omitempty"`
	ComponentValues    []UserSubscriptionsComponentValues `json:"component_values,omitempty"`
	Zone               UserSubscriptionsZone              `json:"zone,omitempty"`
	Frequency          string                             `json:"frequency,omitempty"`
	CurrentPeriodStart *time.Time                         `json:"current_period_start,omitempty"`
	CurrentPeriodEnd   *time.Time                         `json:"current_period_end,omitempty"`
	State              string                             `json:"state,omitempty"`
	Currency           string                             `json:"currency,omitempty"`
	AppliedCredits     UserSubscriptionsApp               `json:"app,omitempty"`
	Price              float64                            `json:"price,omitempty"`
	Entitled           bool                               `json:"entitled,omitempty"`
	CancelAtPeriodEnd  bool                               `json:"cancel_at_period_end,omitempty"`
}
// UserSubscriptionsProduct contains information about the product of a subscription.
type UserSubscriptionsProduct struct {
	Name       string `json:"name,omitempty"`
	Period     string `json:"period,omitempty"`
	Billing    string `json:"billing,omitempty"`
	PublicName string `json:"public_name,omitempty"`
	Duration   int    `json:"duration,omitempty"`
}

// UserSubscriptionsRatePlan contains information about the rate plan of a subscription.
type UserSubscriptionsRatePlan struct {
	ID                string   `json:"id,omitempty"`
	PublicName        string   `json:"public_name,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Scope             string   `json:"scope,omitempty"`
	ExternallyManaged bool     `json:"externally_managed,omitempty"`
	Sets              []string `json:"sets,omitempty"`
	IsContract        bool     `json:"is_contract,omitempty"`
}

// UserSubscriptionsComponentValues contains information about the component values of a subscription.
type UserSubscriptionsComponentValues struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

// UserSubscriptionsZone contains information about the zone of a subscription.
type UserSubscriptionsZone struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// UserSubscriptionsApp contains information about the app of a subscription.
type UserSubscriptionsApp struct {
	Install_ID string `json:"install_id,omitempty"`
}

// // UserSubscriptionsResponse wraps a response containing user subscriptions.
type UserSubscriptionsResponse struct {
	Response
	Result []UserSubscriptions `json:"result"`
}

// UserSubscriptions returns a list of all subscriptions for the current user.
//
// API reference: https://api.cloudflare.com/#user-subscription-get-user-subscriptions
func (api *API) UserSubscriptions(ctx context.Context) ([]UserSubscriptions, error) {
	var r UserSubscriptionsResponse
	res, err := api.makeRequestContext(ctx, http.MethodGet, "/user/subscriptions", nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}
