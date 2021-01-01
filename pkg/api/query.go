package api

const query = `{
	viewer {
	repositories(
		first: 20
		orderBy: { field: PUSHED_AT, direction: DESC }
		ownerAffiliations: [COLLABORATOR, OWNER]
	) {
		edges {
		node {
			defaultBranchRef {
			target {
				... on Commit {
				checkSuites(first: 20) {
					nodes {
					conclusion
					status
					}
				}
				}
			}
			}
		}
		}
	}
	}
}`

type QueryOutline struct {
	Data struct {
		Viewer struct {
			Repositories struct {
				Edges []struct {
					Node struct {
						DefaultBranchRef struct {
							Target struct {
								CheckSuites struct {
									Nodes []struct {
										Conclusion interface{} `json:"conclusion"`
										Status     string      `json:"status"`
									} `json:"nodes"`
								} `json:"checkSuites"`
							} `json:"target"`
						} `json:"defaultBranchRef"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"repositories"`
		} `json:"viewer"`
	} `json:"data"`
}
