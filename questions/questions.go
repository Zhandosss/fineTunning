package questions

func New() []string {
	questions := make([]string, 0, 16)

	//question 1
	questions = append(questions, "Find %[1]s provider API reference/documentation link")

	//question 2
	questions = append(questions, "Where can I create account and app for %[1]s provider?")

	//question 3
	questions = append(questions, "Give me list of all %[1]s provider API scopes")

	//question 4
	questions = append(questions, "What are the different types of authentications %[1]s provider supports?")

	//question 5
	questions = append(questions, "For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?")

	//question 6
	questions = append(questions, "What entities or objects could be queried via the %[1]s provider API? Provide links to the API documentation for each object or entity:")

	//question 7
	questions = append(questions, "Give information about %[1]s provider API rate limits")

	//question 8
	questions = append(questions, "For %[1]s provider API give me information about what rate limit headers are returned in response:")

	//question 9
	questions = append(questions, "Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?")

	//question 10
	questions = append(questions, "%[1]s provider in API supports link based or offset based pagination? Tell me about pagination limits")

	//question 11
	questions = append(questions, "What is the error response format in %[1]s provider API?")

	//question 12
	questions = append(questions, "The version of the restful API is usually a part of the endpoint url, e.g. for version 4 the url looks like http://example.com/users/v4/1234/. How many versions of such api is available for %[1]s?:")

	//question 13
	questions = append(questions, "Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog")
	return questions
}
