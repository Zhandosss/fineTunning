package training

type Mural struct {
	answers     []string
	searchSites []string
}

func NewMural() *Mural {
	return &Mural{
		answers:     make([]string, 0, 20),
		searchSites: []string{"https://developers.mural.co/public/docs/getting-started", "https://developers.mural.co/public/reference/intro"},
	}
}

func (m *Mural) GetSearchSites() []string {
	return m.searchSites
}

func (m *Mural) GetName() string {
	return "Mural"
}

func (m *Mural) GetAns() []string {
	return m.answers
}

func (m *Mural) GenerateAns() {
	//question 1: Find %[1]s provider API reference/documentation link
	//TODO: Delete enterprise API link?
	m.answers = append(m.answers, "Public Mural API docs: https://developers.mural.co/public/docs/getting-started\\nEnterprise API docs: https://developers.mural.co/enterprise/docs/welcome")

	//question 2: Where can I create account and app for %[1]s provider?
	m.answers = append(m.answers, "Sing up via this link: https://app.mural.co/signup\\nTo create a new app go here and click the New app button: https://app.mural.co/me/apps. First, provide the app name and redirect URI. Then you will get the client ID and client secret.")

	//question 3: Give me list of all %[1]s provider API scopes
	m.answers = append(m.answers, "There is the list of all scopes: rooms:read, users:read, workspaces:read, murals:read, identity:read, templates:read, rooms:write, murals:write, templates:write")

	//question 4: What are the different types of authentications %[1]s provider supports?
	m.answers = append(m.answers, "Public mural API uses OAuth2 Authentication. You can read more about OAuath2 authentication here: https://developers.mural.co/public/docs/oauth ")

	//question 5: For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?
	m.answers = append(m.answers, " Access token expires in 15 minutes. I didn't find information about refresh token expiration time.")

	//question 6: What entities or objects could be queried via the %[1]s provider API? Provide links to the API documentation for each object or entity:
	m.answers = append(m.answers, "Here are some of the main API entities that can be queried via the Mural API:\\n1)Rooms - https://developers.mural.co/public/reference/rooms\\n2)Users - https://developers.mural.co/public/reference/users\\n3)Workspaces - https://developers.mural.co/public/reference/workspaces\\n4)Murals - https://developers.mural.co/public/reference/murals\\n5)Templates - https://developers.mural.co/public/reference/templates\\n6)Mural Contents - https://developers.mural.co/public/reference/mural-contents")

	//question 7: Give information about %[1]s provider API rate limits
	m.answers = append(m.answers, "Mural has a User rate limit and a Client application rate limit.\\nUser rate limit: 25 requests/user/second\\nClient application rate limit: 10,000/app/min")

	//question 8: For %[1]s provider API give me information about what rate limit headers are returned in response:
	m.answers = append(m.answers, "Mural API returns rate limit headers in response. Here are the headers:\\n1)X-RateLimit-Limit - User rate limit\\n2)X-RateLimit-Remaining - How many requests are available for the user until the rate limit reset\\n3)X-RateLimit-Reset - Timestamp when the user rate limit will reset\\n4)X-RateLimit-App-Limit - Client application rate limit\\n5)X-RateLimit-App-Remaining - How many requests are available for the client application until the rate limit reset\\n6)X-RateLimit-App-Reset - Timestamp when the client application rate limit will reset")

	//question 9: Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?
	//TODO: Bad answer? Try to provide another answer
	m.answers = append(m.answers, "According to this changelog: https://developers.mural.co/public/changelog/mural-api-changelog-2022-11-15 create sticky note endpoint(https://developers.mural.co/public/reference/createstickynote), create text box endpoint(https://developers.mural.co/public/reference/createtextbox) and create title endpoint(https://developers.mural.co/public/reference/createtitle) supports bulk operations.")

	//question 10: %[1]s provider in API supports link based or offset based pagination? Tell me about pagination limits
	m.answers = append(m.answers, "Mural uses token pagination. Each resource type has its own default page size of 100 or fewer. \\nMural uses two pagination query params limit and next. With the next key use the token to get access to the next page")

	//question 11: What is the error response format in %[1]s provider API?
	m.answers = append(m.answers, `The error response body has the following format:\\n{\\n    \"code\": \"<error code>\",\\n    \"message\": \"<error message>\"\\n}`)

	//question 12: The version of the restful API is usually a part of the endpoint url, e.g. for version 4 the url looks like http://example.com/users/v4/1234/. How many versions of such api is available for %[1]s?:
	m.answers = append(m.answers, "The current version of Mural API is v1")

	//question 13: Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog
	m.answers = append(m.answers, "You can find Mural API changelog here: https://developers.mural.co/public/changelog")
}
