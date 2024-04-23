package training

type Klaviyo struct {
	Answers []string
}

func NewKlaviyo() *Klaviyo {
	return &Klaviyo{
		Answers: make([]string, 0, 20),
	}
}

func (k *Klaviyo) GetName() string {
	return "Klaviyo"
}

func (k *Klaviyo) GetAns() []string {
	return k.Answers
}

func (k *Klaviyo) GenerateAns() {
	//question 1: Find %[1]s provider API reference/documentation link
	k.Answers = append(k.Answers, "https://developers.klaviyo.com/en/reference/api_overview")

	//question 2: Where can I create account and app for %[1]s provider?
	k.Answers = append(k.Answers, `Create an account at https://www.klaviyo.com/sign-up.\nTo create a client app go to https://www.klaviyo.com/oauth/client and click Create app button in the top right corner.`)

	//question 3: Give me list of all %[1]s provider API scopes
	k.Answers = append(k.Answers, `Each endpoint has 2 associated scopes.\nThere is the list of all scopes: accounts:read, accounts:write, catalogs:read, catalogs:write, conversations:read, conversations:write, coupons:read, coupons:write, coupon-codes:read, coupon-codes:write, campaigns:read, campaigns:write, data-privacy:read, data-privacy:write, events:read, events:write, flows:read, flows:write, images:read, images:write, lists:read, lists:write, metrics:read, metrics:write, profiles:read, profiles:write, push-tokens:read, push-tokens:write, segments:read, segments:write, subscriptions:read, subscriptions:write, tags:read, tags:write, templates:read, templates:write`)

	//question 4: What are the different types of authentications %[1]s provider supports?
	k.Answers = append(k.Answers, "Klaviyo supports OAuth2(https://developers.klaviyo.com/en/docs/set_up_oauth) and API key authentication types.")

	//question 5: For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?
	k.Answers = append(k.Answers, "Access token expires after 1 hour, refresh token expires after 90 days. Refresh token can be called 10 times per minute maximum.")

	//question 6: What entities or objects could be queried via the %[1]s provider API? Provide links to the API documentation for each object or entity:
	k.Answers = append(k.Answers, `Here are some of the main API entities that can be queried via the Klaviyo API:\n1)Profiles - https://developers.klaviyo.com/en/reference/profiles\n2)Lists - https://developers.klaviyo.com/en/reference/lists\n3)Campaigns - https://developers.klaviyo.com/en/reference/campaigns\n4)Flows - https://developers.klaviyo.com/en/reference/flows\n5)Metrics - https://developers.klaviyo.com/en/reference/metrics\n6)Tags - https://developers.klaviyo.com/en/reference/tags`)

	//question 7: Give information about %[1]s provider API rate limits
	k.Answers = append(k.Answers, `OAuth apps receive their own rate limit quota per account per app.\nPrivate key integrations share the same rate limit quota per account.\nIf you have a Private key, and two OAuth apps they all will have different rate limit quotas. (https://developers.klaviyo.com/en/docs/set_up_oauth#rate-limits)\nEach API endpoint has different rate limit. API uses fixed-window rate limiting algorithm with two distinct windows:\nburst window(1-second window)\nsteady window(1-minute window)`)

	//question 8: For %[1]s provider API give me information about what rate limit headers are returned in response:
	k.Answers = append(k.Answers, `Klaviyo API returns rate limit headers in response. Here are the headers:\n1)RateLimit-Limit - Number of requests allowed per time period\n2)RateLimit-Remaining - Approximate number of requests remaining within a window\n3)RateLimit-Reset - Number of seconds remaining before current window resets`)

	//question 9: Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?
	k.Answers = append(k.Answers, `Klaviyo supports Bulk Profile Import (https://developers.klaviyo.com/en/docs/use_klaviyos_bulk_profile_import_api)`)

	//question 10: %[1]s provider in API supports link based or offset based pagination? Tell me about pagination limits
	k.Answers = append(k.Answers, `Klaviyo supports cursor based pagination with query parameter ?page[cursor].\nIn every pagination request you will recieve three links. Using this link you can get to the next page or previous page.`)

	//question 11: What is the error response format in %[1]s provider API?
	k.Answers = append(k.Answers, `Klaviyo API returns error response in JSON format. This is an example of error response: \n{\n  \"errors\": [\n    {\n      \"id\": \"2c76424d-b1ff-4afd-a2a1-728d85dac775\",\n      \"status\": 400,\n      \"code\": \"invalid\",\n      \"title\": \"Invalid input.\",\n      \"detail\": \"'profiless' is not an allowed include parameter for this resource.\",\n      \"source\": {\n        \"parameter\": \"include\"\n      },\n      \"meta\": {}\n    }\n  ]\n}`)

	//question 12: The version of the restful API is usually a part of the endpoint url, e.g. for version 4 the url looks like http://example.com/users/v4/1234/. How many versions of such api is available for %[1]s?:
	k.Answers = append(k.Answers, "The current version of Klaviyo API is v2024-02-15")

	//question 13: Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog
	k.Answers = append(k.Answers, "You can find Klaviyo API changelog here: https://developers.klaviyo.com/en/docs/changelog")
}
