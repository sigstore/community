package config

type Config struct {
	Users        []User       `yaml:"users"`
	Teams        []Team       `yaml:"teams"`
	Repositories []Repository `yaml:"repositories"`
}

type User struct {
	Username string   `yaml:"username"`
	Role     string   `yaml:"role"`
	Teams    []string `yaml:"teams"`
}

type Team struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Privacy     string `yaml:"privacy"`
}

type Repository struct {
	AllowAutoMerge      bool           `yaml:"allowAutoMerge"`
	AllowMergeCommit    bool           `yaml:"allowMergeCommit"`
	AllowRebaseMerge    bool           `yaml:"allowRebaseMerge"`
	AllowSquashMerge    bool           `yaml:"allowSquashMerge"`
	AutoInit            bool           `yaml:"autoInit"`
	DeleteBranchOnMerge bool           `yaml:"deleteBranchOnMerge"`
	HasDownloads        bool           `yaml:"hasDownloads"`
	HasIssues           bool           `yaml:"hasIssues"`
	HasProjects         bool           `yaml:"hasProjects"`
	HasWiki             bool           `yaml:"hasWiki"`
	VulnerabilityAlerts bool           `yaml:"vulnerabilityAlerts"`
	Visibility          string         `yaml:"visibility"`
	Name                string         `yaml:"name"`
	Owner               string         `yaml:"owner"`
	Description         string         `yaml:"description"`
	HomepageUrl         string         `yaml:"homepageUrl"`
	LicenseTemplate     string         `yaml:"licenseTemplate"`
	Topics              []string       `yaml:"topics"`
	Collaborators       []Collaborator `yaml:"collaborators"`
}

type Collaborator struct {
	Username   string `yaml:"username"`
	Permission string `yaml:"permission"`
}
