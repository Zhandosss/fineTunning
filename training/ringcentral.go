package training

type RingCentral struct {
	answers     []string
	SearchSites []string
}

func NewRingCentral() *RingCentral {
	return &RingCentral{
		answers:     make([]string, 0, 20),
		SearchSites: []string{"https://developers.ringcentral.com/", "https://developers.ringcentral.com/api-reference", "https://developers.ringcentral.com/guide"},
	}
}

func (r *RingCentral) GetName() string {
	return "RingCentral"
}

func (r *RingCentral) GetAns() []string {
	return r.answers
}

func (r *RingCentral) GenerateAns() {
	//question 1: Find %[1]s provider API reference/documentation link
	r.answers = append(r.answers, "https://developers.ringcentral.com/api-reference")

	//question 2: Where can I create account and app for %[1]s provider?
	r.answers = append(r.answers, `Create account on this link https://developers.ringcentral.com/sign-up. After this activate your account via message on email.\nTo create oauth2 app go to: https://developers.ringcentral.com/my-account.html#/applications`)

	//question 3: Give me list of all %[1]s provider API scopes
	r.answers = append(r.answers, `Required scopes are generally declared at the stage of application registration and confirmed by the user during the authorization stage of connecting to that application.\n\nAdding and removing scopes is accomplished in the Developer Console by editing the application's settings.`)

	//question 4: What are the different types of authentications %[1]s provider supports?
	r.answers = append(r.answers, `OAuth2: https://developers.ringcentral.com/guide/authentication/auth-code-flow\nOAuth2 with PKCE: https://developers.ringcentral.com/guide/authentication/auth-code-pkce-flow\nJWT: https://developers.ringcentral.com/guide/authentication/jwt-flow\nPassword auth(deprecated): https://developers.ringcentral.com/guide/authentication/password-flo`)

	//question 5: For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?
	r.answers = append(r.answers, `Access token expires after 1 hour, refresh token expires after 7 days.`)

	//question 6: What entities or objects could be queried via the %[1]s provider API? Provide links to the API documentation for each object or entity:
	r.answers = append(r.answers, `Here are some of the main API entities that can be queried via the RingCentral API:\n1)Account - https://developers.ringcentral.com/api-reference/account\n2)Voice - https://developers.ringcentral.com/api-reference/voice\n3)SMS and Fax - https://developers.ringcentral.com/api-reference/sms-and-fax\n4)Social Messaging - https://developers.ringcentral.com/api-reference/social-messaging\n5)Team Messaging - https://developers.ringcentral.com/api-reference/team-messaging\n6)Video - https://developers.ringcentral.com/api-reference/video`)

	//question 7: Give information about %[1]s provider API rate limits
	r.answers = append(r.answers, `Every API endpoint assigned to different API Group with different Rate Limit. \nLight - 50 requests/minute\nMedium - 40 requests/minute\nHeavy - 10 requests/minute\nAuth - 5 requests/minute\n\nTo see to which group the endpoint belongs go to API reference(https://developers.ringcentral.com/api-reference/) select any API group from left menu and select any endpoint. For example https://developers.ringcentral.com/api-reference/Adaptive-Cards/createGlipAdaptiveCardNew endpoint. In the description you can see Usage Plan Group`)

	//question 8: For %[1]s provider API give me information about what rate limit headers are returned in response:
	r.answers = append(r.answers, `RingCentral API returns rate limit headers in response. Here are the headers:\n1)X-Rate-Limit-Limit - current rate limit for the given request\n2)X-Rate-Limit-Remaining - current rate limit for the given request\n3)X-Rate-Limit-Window - current rate limit for the given request\n4)X-Rate-Limit-Group - current rate limit for the given request`)

	//question 9: Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?
	r.answers = append(r.answers, `You can find information about batch operations here: Link: https://developers.ringcentral.com/guide/basics/batch-requests`)

	//question 10: %[1]s provider in API supports link based or offset based pagination? Tell me about pagination limits
	r.answers = append(r.answers, `It seems like RingEX API supports different types of pagination. For example https://developers.ringcentral.com/api-reference/Calendar-Events/readGlipEventsNew endpoint supports token(cursor) pagination. Maximum recordCount is 250. Default - 30. And https://developers.ringcentral.com/api-reference/Call-Blocking/listBlockedAllowedNumbers endpoint supports link pagination. It seems like  maximum page limit for this endpoint is 1000. Default - 100.`)

	//question 11: What is the error response format in %[1]s provider API?
	r.answers = append(r.answers, `RingEX supports specific error codes. If an error occurred, the response body returns a errorCode, message and sometimes additionalInfo\n{\n    \"errors\": [\n    {\n        \"errorCode\": \"ABC-123\",    \n        \"message\": \"Error message\",\n        \"additionalInfo\": \"value\",  // additional attribute for this error, e.g \"parameterName\": \"extensionId\"\n    },\n    {\n        \"errorCode\": \"XYZ-321\",\n        \"message\": \"Second error message\"\n    }\n    ]\n}`)

	//question 12: The version of the restful API is usually a part of the endpoint url, e.g. for version 4 the url looks like http://example.com/users/v4/1234/. How many versions of such api is available for %[1]s?:
	r.answers = append(r.answers, "The current version of RingCentral API is v2")

	//question 13: Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog
	//TODO: add release notes from https://support.ringcentral.com/release-notes/ringex/ringex-developers/api-sdk.html
	r.answers = append(r.answers, "I didn't find information about RingCentral API changelog.")
}
