// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"github.com/fleetdm/fleet/server/fleet"
)

var _ fleet.TeamStore = (*TeamStore)(nil)

type NewTeamFunc func(team *fleet.Team) (*fleet.Team, error)

type SaveTeamFunc func(team *fleet.Team) (*fleet.Team, error)

type DeleteTeamFunc func(tid uint) error

type TeamFunc func(id uint) (*fleet.Team, error)

type TeamByNameFunc func(name string) (*fleet.Team, error)

type ListTeamsFunc func(filter fleet.TeamFilter, opt fleet.ListOptions) ([]*fleet.Team, error)

type SearchTeamsFunc func(filter fleet.TeamFilter, matchQuery string, omit ...uint) ([]*fleet.Team, error)

type TeamEnrollSecretsFunc func(teamID uint) ([]*fleet.EnrollSecret, error)

type TeamStore struct {
	NewTeamFunc        NewTeamFunc
	NewTeamFuncInvoked bool

	SaveTeamFunc        SaveTeamFunc
	SaveTeamFuncInvoked bool

	DeleteTeamFunc        DeleteTeamFunc
	DeleteTeamFuncInvoked bool

	TeamFunc        TeamFunc
	TeamFuncInvoked bool

	TeamByNameFunc        TeamByNameFunc
	TeamByNameFuncInvoked bool

	ListTeamsFunc        ListTeamsFunc
	ListTeamsFuncInvoked bool

	SearchTeamsFunc        SearchTeamsFunc
	SearchTeamsFuncInvoked bool

	TeamEnrollSecretsFunc        TeamEnrollSecretsFunc
	TeamEnrollSecretsFuncInvoked bool
}

func (s *TeamStore) NewTeam(team *fleet.Team) (*fleet.Team, error) {
	s.NewTeamFuncInvoked = true
	return s.NewTeamFunc(team)
}

func (s *TeamStore) SaveTeam(team *fleet.Team) (*fleet.Team, error) {
	s.SaveTeamFuncInvoked = true
	return s.SaveTeamFunc(team)
}

func (s *TeamStore) DeleteTeam(tid uint) error {
	s.DeleteTeamFuncInvoked = true
	return s.DeleteTeamFunc(tid)
}

func (s *TeamStore) Team(id uint) (*fleet.Team, error) {
	s.TeamFuncInvoked = true
	return s.TeamFunc(id)
}

func (s *TeamStore) TeamByName(identifier string) (*fleet.Team, error) {
	s.TeamByNameFuncInvoked = true
	return s.TeamByNameFunc(identifier)
}

func (s *TeamStore) ListTeams(filter fleet.TeamFilter, opt fleet.ListOptions) ([]*fleet.Team, error) {
	s.ListTeamsFuncInvoked = true
	return s.ListTeamsFunc(filter, opt)
}

func (s *TeamStore) SearchTeams(filter fleet.TeamFilter, matchQuery string, omit ...uint) ([]*fleet.Team, error) {
	s.SearchTeamsFuncInvoked = true
	return s.SearchTeamsFunc(filter, matchQuery, omit...)
}

func (s *TeamStore) TeamEnrollSecrets(teamID uint) ([]*fleet.EnrollSecret, error) {
	s.TeamEnrollSecretsFuncInvoked = true
	return s.TeamEnrollSecretsFunc(teamID)
}
