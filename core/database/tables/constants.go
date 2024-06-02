package tables

const (
	prefix = "_tb"

	userTableName        = "user" + prefix
	sessionTableName     = "session" + prefix
	interestTableName    = "interest" + prefix
	userIntrestTableName = "user_interest" + prefix

	postTableName    = "post" + prefix
	commentTableName = "comment" + prefix
	likeTableName    = "like" + prefix

	tripTableName  = "trip" + prefix
	venueTableName = "venue" + prefix

	tripInterestTableName = "trip_interest" + prefix
	tripVenueTableName    = "trip_venue" + prefix

	postViewTableName = "post_view" + prefix
)
