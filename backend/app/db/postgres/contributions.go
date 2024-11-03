package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertContributions(githubId int64, login string, fork bool, repoId int64, fullName string, cons float64, talent float64) error {
	con := models.ContributionsStored{
		DeveloperGithubId: githubId,
		DeveloperLogin:    login,
		Fork:              fork,
		RepoGithubId:      repoId,
		RepoFullName:      fullName,
		Contributions:     cons,
		TalentScore:       talent,
	}
	res := pdb.Create(&con)
	if res.Error != nil {
		zap.L().Error("insert contributions failed", zap.Error(res.Error))
		zap.L().Debug("insert contributions failed", zap.Error(res.Error),
			zap.String("github_login", login),
			zap.String("repo_full_name", fullName))
		return res.Error
	}
	return nil
}

func GetContributionsByDeveloper(githubLogin string) ([]models.ContributionsStored, error) {
	var cons []models.ContributionsStored
	res := pdb.Order("talent_score desc").Find(&cons, "developer_login = ?", githubLogin)
	if res.Error != nil {
		zap.L().Error("get contributions by developer failed", zap.Error(res.Error))
		return []models.ContributionsStored{}, res.Error
	}
	return cons, nil
}
