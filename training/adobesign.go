package training

type AdobeSign struct {
	answers     []string
	SearchSites []string
}

func NewAdobeSign() *AdobeSign {
	return &AdobeSign{
		answers:     make([]string, 0, 20),
		SearchSites: []string{"https://opensource.adobe.com/acrobat-sign/", "https://secure.na1.adobesign.com/public/docs/restapi/v6"},
	}
}

func (a *AdobeSign) GetName() string {
	return "AdobeSign"
}

func (a *AdobeSign) GetAns() []string {
	return a.answers
}

func (a *AdobeSign) GenerateAns() {
	//question 1: Find %[1]s provider API reference/documentation link
	a.answers = append(a.answers, "https://secure.na1.adobesign.com/public/docs/restapi/v6#://")

	//question 2: Where can I create account and app for %[1]s provider?
	a.answers = append(a.answers, `Create an account at https://www.adobe.com/sign/developer-form.html\nCreate an app at https://secure.na4.adobesign.com/account/accountSettingsPage#pageId::API_APPLICATIONS`)

	//question 3: Give me list of all %[1]s provider API scopes
	a.answers = append(a.answers, `There are 3 scopes: user_read, user_write, user_login, agreement_read, agreement_write, agreement_send, widget_read, widget_write, library_read, library_write, workflow_read, workflow_write, webhook_read, webhook_write, webhook_retention, application_read, application_write.\nEach scope has one of three modifiers: self, group, account.`)

	//question 4: What are the different types of authentications %[1]s provider supports?
	a.answers = append(a.answers, `Adobe sign supports OAuth2 (https://secure.na1.adobesign.com/public/static/oauthDoc.jsp) and integration key (https://helpx.adobe.com/sign/kb/how-to-create-an-integration-key.html).`)

	//question 5: For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?
	a.answers = append(a.answers, `Access token expires after 1 hour, refresh token expires after 60 days.`)

	//question 6: What entities or objects could be queried via the %[1]s provider API? Provide links to the API documentation for each object or entity:
	a.answers = append(a.answers, `Here are some of the main API entities that can be queried via the AdobeSign API:\n1)Agreements - https://secure.na1.adobesign.com/public/docs/restapi/v6#!/Agreements\n2)Widgets - https://secure.na1.adobesign.com/public/docs/restapi/v6#!/Widgets\n3)Users - https://secure.na1.adobesign.com/public/docs/restapi/v6#!/Users\n4)Webhooks - https://secure.na1.adobesign.com/public/docs/restapi/v6#!/Webhooks\n5)Groups - https://secure.na1.adobesign.com/public/docs/restapi/v6#!/Groups\n6)Workflows - https://secure.na1.adobesign.com/public/docs/restapi/v6#!/Workflows`)

	//question 7: Give information about %[1]s provider API rate limits
	a.answers = append(a.answers, `General info about throttling from https://helpx.adobe.com/sign/using/transaction-limits.html\n\"Each request to Acrobat Sign is evaluated based on the amount of system resources it consumes. Different parameters passed to the same endpoint might contribute to a different amount of resource consumption.\nFurthermore, some requests may trigger lengthy background processes that are also considered in the throttling evaluation algorithm.\nTherefore, the rate of requests can not be described as simple numbers of requests in a certain period of time. Throttling policies are based on historical daily request data, including legitimate usages that stress the system.\nCustomers can be assured that the Acrobat Sign policies are generous enough that regular daily workflow volume will not be affected.\"`)

	//question 8: For %[1]s provider API give me information about what rate limit headers are returned in response:
	a.answers = append(a.answers, `AdobeSign provider don't return rate limit headers in response.`)

	//question 9: Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?
	a.answers = append(a.answers, `You need to use megaSigns API recourse to get access to Bulk send. \nLink: https://secure.na1.adobesign.com/public/docs/restapi/v6#!/megaSigns/createMegaSign`)

	//question 10: %[1]s provider in API supports link based or offset based pagination? Tell me about pagination limits
	a.answers = append(a.answers, `Adobe sign use cursor based pagination. Page size limited by pageSize query param. If you don't set page size you will receive all list of resources (https://opensource.adobe.com/acrobat-sign/releasenotes/v6releasenotes.html#pagination-support). In pagination request if you have next page you will receive \"nextCursor\" key in response body.`)

	//question 11: What is the error response format in %[1]s provider API?
	a.answers = append(a.answers, `If the request was unsuccessful, you will receive code(not HTTP error code) and message in response body. Example of how this response looks:{\n    \"code\": \"INVALID_PAGE_SIZE\",\n    \"message\": \"Page size is either invalid or not within permissible range.\"\n}`)

	//question 12: The version of the restful API is usually a part of the endpoint url, e.g. for version 4 the url looks like http://example.com/users/v4/1234/. How many versions of such api is available for %[1]s?:
	a.answers = append(a.answers, "The current version of AdobeSign API is v6")

	//question 13: Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog
	a.answers = append(a.answers, "You can find AdopeSign API changelog here https://opensource.adobe.com/acrobat-sign/releasenotes/v6releasenotes.html")

}
