package config

type Config struct {
	Users        []User       `yaml:"users"`
	Teams        []Team       `yaml:"teams"`
	Repositories []Repository `yaml:"repositories"`
}

type User struct {
	Username string     `yaml:"username"`
	Role     string     `yaml:"role"`
	Teams    []UserTeam `yaml:"teams"`
}

type UserTeam struct {
	Name string `yaml:"name"`
	Role string `yaml:"role"`
}

type Team struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Privacy      string `yaml:"privacy"`
	ParentTeamID int    `yaml:"parentTeamId"`
}

type Repository struct {
	AllowAutoMerge      bool               `yaml:"allowAutoMerge"`
	AllowMergeCommit    bool               `yaml:"allowMergeCommit"`
	AllowRebaseMerge    bool               `yaml:"allowRebaseMerge"`
	AllowSquashMerge    bool               `yaml:"allowSquashMerge"`
	Archived            bool               `yaml:"archived"`
	AutoInit            bool               `yaml:"autoInit"`
	DeleteBranchOnMerge bool               `yaml:"deleteBranchOnMerge"`
	HasDownloads        bool               `yaml:"hasDownloads"`
	HasIssues           bool               `yaml:"hasIssues"`
	HasProjects         bool               `yaml:"hasProjects"`
	HasWiki             bool               `yaml:"hasWiki"`
	VulnerabilityAlerts bool               `yaml:"vulnerabilityAlerts"`
	Visibility          string             `yaml:"visibility"`
	Name                string             `yaml:"name"`
	Owner               string             `yaml:"owner"`
	Description         string             `yaml:"description"`
	HomepageUrl         string             `yaml:"homepageUrl"`
	LicenseTemplate     string             `yaml:"licenseTemplate"`
	Topics              []string           `yaml:"topics"`
	Pages               Pages              `yaml:"pages"`
	IsTemplate          bool               `yaml:"isTemplate"`
	Template            Template           `yaml:"template"`
	Collaborators       []Collaborator     `yaml:"collaborators"`
	Teams               []RepoTeam         `yaml:"teams"`
	BranchesProtection  []BranchProtection `yaml:"branchProtection"`
}

type Pages struct {
	CNAME  string `yaml:"cname"`
	Branch string `yaml:"branch"`
	Path   string `yaml:"path"`
}

type Template struct {
	Owner      string `yaml:"owner"`
	Repository string `yaml:"repository"`
}

type Collaborator struct {
	Username   string `yaml:"username"`
	Permission string `yaml:"permission"`
}

type RepoTeam struct {
	Name       string `yaml:"name"`
	ID         string `yaml:"id"`
	Permission string `yaml:"permission"`
}

type BranchProtection struct {
	Pattern                       string   `yaml:"pattern"`
	EnforceAdmins                 bool     `yaml:"enforceAdmins"`
	AllowsDeletions               bool     `yaml:"allowsDeletions"`
	AllowsForcePushes             bool     `yaml:"allowsForcePushes"`
	DismissStaleReviews           bool     `yaml:"dismissStaleReviews"`
	DismissalRestrictions         []string `yaml:"dismissalRestrictions"`
	RestrictDismissals            bool     `yaml:"restrictDismissals"`
	RequiredLinearHistory         bool     `yaml:"requiredLinearHistory"`
	RequireSignedCommits          bool     `yaml:"requireSignedCommits"`
	RequireConversationResolution bool     `yaml:"requireConversationResolution"`
	RequireCodeOwnerReviews       bool     `yaml:"requireCodeOwnerReviews"`
	RequiredApprovingReviewCount  int      `yaml:"requiredApprovingReviewCount"`
	StatusChecks                  []string `yaml:"statusChecks"`
	RequireBranchesUpToDate       bool     `yaml:"requireBranchesUpToDate"`
	PushRestrictions              []string `yaml:"pushRestrictions"`
}
