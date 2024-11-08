package models

type ContributionsStored struct {
	Id                int64   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	DeveloperGithubId int64   `json:"developer_github_id" gorm:"not null"`
	DeveloperLogin    string  `json:"developer_login" gorm:"not null;index:idx_github_login,type:hash"`
	Fork              bool    `json:"fork" gorm:"not null"`
	RepoGithubId      int64   `json:"repo_github_id" gorm:"not null"`
	RepoFullName      string  `json:"repo_full_name" gorm:"not null;index:idx_repo_full_name,type:hash"`
	StargazersCount   int64   `json:"stargazers_count" gorm:"not null;default:0"`
	Contributions     float64 `json:"contributions,string" gorm:"not null;index:idx_contributions,type:btree"`
	TalentScore       float64 `json:"talent_score,string" gorm:"not null;default:0"`
}
