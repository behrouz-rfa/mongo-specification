package filters

type UserBy struct {
	ID             *string
	Email          *string
	FirstName      *string
	LastName       *string
	MiddleName     *string
	ProfilePicture *string
	PhoneNumber    *string
	EmailVerified  *bool
	Active         *bool
}

type UserFilter struct {
	Email          *StringFilter
	FirstName      *StringFilter
	LastName       *StringFilter
	MiddleName     *StringFilter
	ProfilePicture *StringFilter
	PhoneNumber    *StringFilter
	EmailVerified  *bool
	Active         *bool
	CreatedAt      *TimeRange
	UpdatedAt      *TimeRange
}
