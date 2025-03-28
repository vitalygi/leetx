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
	problemsetQuestionListQuery = `
		query problemsetQuestionList(
		  $categorySlug: String, 
		  $limit: Int, 
		  $skip: Int, 
		  $filters: QuestionListFilterInput
		) {
		  problemsetQuestionList: questionList(
			categorySlug: $categorySlug
			limit: $limit
			skip: $skip
			filters: $filters
		  ) {
			total: totalNum
			questions: data {
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
		}
	`
)
