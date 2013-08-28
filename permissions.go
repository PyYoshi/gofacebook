package gofacebook

type Permissions struct {
	// Extended Profile Properties
	// https://developers.facebook.com/docs/reference/login/extended-profile-properties/
	UserAboutMe                bool
	FriendsAboutMe             bool
	UserActivities             bool
	FriendsActivities          bool
	UserBirthday               bool
	FriendsBirthday            bool
	UserCheckIns               bool
	FriendsCheckIns            bool
	UserEducationHistory       bool
	FriendsEducationHistory    bool
	UserEvents                 bool
	FriendsEvents              bool
	UserGroups                 bool
	FriendsGroups              bool
	UserHometown               bool
	FriendHometown             bool
	UserInterests              bool
	FriendsInterests           bool
	UserLikes                  bool
	FriendsLikes               bool
	UserLocation               bool
	FriendsLocation            bool
	UserNotes                  bool
	FriendsNotes               bool
	UserPhotos                 bool
	FriendsPhotos              bool
	UserQuestions              bool
	FriendsQuestions           bool
	UserRelationships          bool
	FriendsRelationships       bool
	UserRelationshipDetails    bool
	FriendsRelationShipDetails bool
	UserReligionPolitics       bool
	FriendsReligionPolitics    bool
	UserStatus                 bool
	FriendsStatus              bool
	UserSubscriptions          bool
	FriendsSubscriptions       bool
	UserVideos                 bool
	FriendsVideos              bool
	UserWebSite                bool
	FriendsWebSite             bool
	UserWorkHistroy            bool
	FriendsWorkHistroy         bool

	// Email Permissions
	// https://developers.facebook.com/docs/reference/login/email-permissions/
	Email bool
}

func (p *Permissions) _getUserAboutMeKeyString() string {
	return "user_about_me"
}

func (p *Permissions) _getFriendsAboutMeKeyString() string {
	return "friends_about_me"
}

func (p *Permissions) _getUserActivitiesKeyString() string {
	return "user_activities"
}

func (p *Permissions) _getFriendsActivitiesKeyString() string {
	return "friends_activities"
}

func (p *Permissions) _getUserBirthdayKeyString() string {
	return "user_birthday"
}

func (p *Permissions) _getFriendsBirthdayKeyString() string {
	return "friends_birthday"
}

func (p *Permissions) _getUserCheckInsKeyString() string {
	return "user_checkins"
}

func (p *Permissions) _getFriendsCheckInsKeyString() string {
	return "friends_checkins"
}

func (p *Permissions) _getUserEducationHistoryKeyString() string {
	return "user_education_history"
}

func (p *Permissions) _getFriendsEducationHistoryKeyString() string {
	return "friends_education_history"
}

func (p *Permissions) _getUserEventsKeyString() string {
	return "user_events"
}

func (p *Permissions) _getFriendsEventsKeyString() string {
	return "friends_events"
}

func (p *Permissions) _getUserGroupsKeyString() string {
	return "user_groups"
}

func (p *Permissions) _getFriendsGroupsKeyString() string {
	return "friends_groups"
}

func (p *Permissions) _getUserHometownKeyString() string {
	return "user_hometown"
}

func (p *Permissions) _getFriendsHometownKeyString() string {
	return "friends_hometown"
}

func (p *Permissions) _getUserInterestsKeyString() string {
	return "user_interests"
}

func (p *Permissions) _getFriendsInterestsKeyString() string {
	return "friends_interests"
}

func (p *Permissions) _getUserLikesKeyString() string {
	return "user_likes"
}

func (p *Permissions) _getFriendsLikesKeyString() string {
	return "friends_likes"
}

func (p *Permissions) _getUserLocationKeyString() string {
	return "user_location"
}

func (p *Permissions) _getFriendsLocationKeyString() string {
	return "friends_location"
}

func (p *Permissions) _getUserNotesKeyString() string {
	return "user_notes"
}

func (p *Permissions) _getFriendsNotesKeyString() string {
	return "friends_notes"
}

func (p *Permissions) _getUserPhotosKeyString() string {
	return "user_photos"
}

func (p *Permissions) _getFriendsPhotosKeyString() string {
	return "friends_photos"
}

func (p *Permissions) _getUserQuestionsKeyString() string {
	return "user_questions"
}

func (p *Permissions) _getFriendsQuestionsKeyString() string {
	return "friends_questions"
}

func (p *Permissions) _getUserRelationshipsKeyString() string {
	return "user_relationships"
}

func (p *Permissions) _getFriendsFriendsRelationShipsKeyString() string {
	return "friends_relationships"
}

func (p *Permissions) _getUserRelationShipDetailsKeyString() string {
	return "user_relationship_details"
}

func (p *Permissions) _getFriendsRelationShipDetailsKeyString() string {
	return "friends_relationship_details"
}

func (p *Permissions) _getUserReligionPoliticsKeyString() string {
	return "user_religion_politics"
}

func (p *Permissions) _getFriendsReligionPoliticsKeyString() string {
	return "friends_religion_politics"
}

func (p *Permissions) _getUserStatusKeyString() string {
	return "user_status"
}

func (p *Permissions) _getFriendsStatusKeyString() string {
	return "friends_status"
}

func (p *Permissions) _getUserSubscriptionsKeyString() string {
	return "user_subscriptions"
}

func (p *Permissions) _getFriendsSubscriptionsKeyString() string {
	return "friends_subscriptions"
}

func (p *Permissions) _getUserVideosKeyString() string {
	return "user_videos"
}

func (p *Permissions) _getFriendsVideosKeyString() string {
	return "friends_videos"
}

func (p *Permissions) _getUserWebSiteKeyString() string {
	return "user_website"
}

func (p *Permissions) _getFriendsWebSiteKeyString() string {
	return "friends_website"
}

func (p *Permissions) _getUserWorkHistroyKeyString() string {
	return "user_work_history"
}

func (p *Permissions) _getFriendsWorkHistroyKeyString() string {
	return "friends_work_history"
}

func (p *Permissions) _getEmailKeyString() string {
	return "email"
}

func (p *Permissions) BuildScope() string {
	var q string
	// TODO:
	return q
}
