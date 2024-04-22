package training

type Wrike struct {
	Answers []string
}

func NewWrike() *Wrike {
	return &Wrike{
		Answers: make([]string, 0, 20),
	}
}

func (w *Wrike) GetName() string {
	return "Wrike"
}

func (w *Wrike) GetAns() []string {
	return w.Answers
}

func (w *Wrike) GenerateAns() {
	//question 1: Find %[1]s provider API reference/documentation link
	w.Answers = append(w.Answers, "https://developers.wrike.com/overview/")

	//question 2: Where can I create account and app for %[1]s provider?
	w.Answers = append(w.Answers, "Create an account at: https://www.wrike.com/vak/. After this confirm your email.\\nCreate an app here: https://www.wrike.com/frontend/apps/index.html#/api.")

	//question 3: Give me list of all %[1]s provider API scopes
	w.Answers = append(w.Answers, "In API Reference you can find required scopes for every endpoint. For example, https://developers.wrike.com/api/v4/contacts/#query-contacts endpoint needs to have one of three scopes: Default, wsReadOnly, wsReadWrite.\\n\\nThere is the list of all scopes: Default, wsReadOnly, wsReadWrite, amReadOnlyUser, amReadWriteUser, amReadOnlyGroup, amReadWriteGroup, amReadOnlyInvitation, amReadWriteInvitation, amReadOnlyWorkflow, amReadWriteWorkflow, amReadOnlyTimelogCategory, amReadWriteTimelogCategory, dataExportFull, amReadOnlyAuditLog, amReadOnlyAccessRole, amReadOnlyWorkSchedule, amReadWriteWorkSchedule")

	//question 4: What are the different types of authentications %[1]s provider supports?
	w.Answers = append(w.Answers, "Wrike supports OAuth2 and permanent token(API key) authentication types.")

	//question 5: For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?
	w.Answers = append(w.Answers, "Access token expires after 1 hour. I didn't find information about refresh token expiration time.")

	//question 6: What entities or objects could be queried via the %[1]s provider API? Provide links to the API documentation for each object or entity:
	w.Answers = append(w.Answers, "Here are some of the main API entities that can be queried via the Attio API:\\n1)Contacts - https://developers.wrike.com/api/v4/contacts/\\n2)Users - https://developers.wrike.com/api/v4/users/\\n3)Groups - https://developers.wrike.com/api/v4/groups/\\n4)Tasks - https://developers.wrike.com/api/v4/tasks/\\n5)Workflows - https://developers.wrike.com/api/workflows\\n6)Invitations - https://developers.wrike.com/api/v4/invitations/")

	//question 7: Give information about %[1]s provider API rate limits
	w.Answers = append(w.Answers, "I didn't find information about rate limits in Wrike API documentation.")

	//question 8: For %[1]s provider API give me information about what rate limit headers are returned in response:
	w.Answers = append(w.Answers, "I didn't find information about rate limit headers in Wrike API documentation.")

	//question 9: Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?
	w.Answers = append(w.Answers, "Batch export endpoint:\\nhttps://developers.wrike.com/api/v4/data-export/\\n\\nGroups have update in bulk endpoint.\\nhttps://developers.wrike.com/api/v4/groups/")

	//question 10: %[1]s provider in API supports link based or offset based pagination? Tell me about pagination limits
	w.Answers = append(w.Answers, "Wrike API supports offset-based pagination. The maximum number of items per page is 1000.")

	//question 11: What is the error response format in %[1]s provider API?
	w.Answers = append(w.Answers, `Wrike API returns error response in JSON format. This is an example of error response:\\n{\\n  \"error\": \"invalid_request\",\\n  \"errorDescription\": \"Invalid request\"\\n}`)

	//question 12: The version of the restful API is usually a part of the endpoint url, e.g. for version 4 the url looks like http://example.com/users/v4/1234/. How many versions of such api is available for %[1]s?:
	w.Answers = append(w.Answers, "The current version of Wrike API is v4")

	//question 13: Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog
	w.Answers = append(w.Answers, "You can find Wrike API changelog here: https://developers.wrike.com/change-log/")
}
