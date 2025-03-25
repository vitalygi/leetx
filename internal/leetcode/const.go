package leetcode

const (
	leetCodeGraphQLURL  = "https://leetcode.com/graphql/"
	questionDetailQuery = `
        query questionDetail($titleSlug: String!) {
            question(titleSlug: $titleSlug) {
                title
                titleSlug
                content
                questionFrontendId
                difficulty
                questionTitle
                codeSnippets {
                    lang
                    langSlug
                    code
                }
                topicTags {
                    name
                    slug
                }
            }
        }
    `
)
