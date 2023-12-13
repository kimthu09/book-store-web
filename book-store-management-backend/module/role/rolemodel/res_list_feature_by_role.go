package rolemodel

import "book-store-management-backend/module/feature/featuremodel"

type ResListFeatureByRole struct {
	// Data contains list of feature of role.
	Data []featuremodel.ResFeatureDetail `json:"data"`
}
