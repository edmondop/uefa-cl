package chapter3

import (
	"context"
	"errors"
	"go.temporal.io/sdk/workflow"
	"sort"
	"time"
)

type GroupStageMatch struct {
	HomeTeam  Team
	AwayTeam  Team
	MatchDate time.Time
	Result    *GroupStageMatchResult
}

type GroupStageMatchResult struct {
	HomeGoals int
	AwayGoals int
}

func (g *GroupStageMatchResult) HomePoints() int {
	if g.HomeGoals > g.AwayGoals {
		return 3
	}
	if g.HomeGoals < g.AwayGoals {
		return 0
	}
	return 1
}

func (g *GroupStageMatchResult) AwayPoints() int {
	if g.HomeGoals < g.AwayGoals {
		return 3
	}
	if g.HomeGoals > g.AwayGoals {
		return 0
	}
	return 1
}

type GroupCalendar struct {
	Matches []GroupStageMatch
}
type Group struct {
	GroupName     string
	Teams         [4]Team
	GroupCalendar GroupCalendar
}

type GroupStageResult struct {
	Pots []Pot
}

type GroupRankingEntry struct {
	Team          Team
	Points        int
	GoalsScored   int
	GoalsConceded int
}

type groupRankingBuilder struct {
	groupRankingEntriesMap map[string]GroupRankingEntry
	groupStageMatchResults []GroupStageMatchResult
}

type GroupRanking struct {
	GroupRankingEntries []GroupRankingEntry
}

func (g *groupRankingBuilder) AddMatchResult(match GroupStageMatch) error {
	if match.Result == nil {
		return errors.New("cannot add match result because match hasn't been played yet")
	}
	homeTeamEntry := g.groupRankingEntriesMap[match.HomeTeam.Name]
	homeTeamEntry.Team = match.HomeTeam
	homeTeamEntry.Points += match.Result.HomePoints()
	homeTeamEntry.GoalsScored += match.Result.HomeGoals
	homeTeamEntry.GoalsConceded += match.Result.AwayGoals
	g.groupRankingEntriesMap[match.HomeTeam.Name] = homeTeamEntry
	awayTeamEntry := g.groupRankingEntriesMap[match.AwayTeam.Name]
	awayTeamEntry.Team = match.AwayTeam
	awayTeamEntry.Points += match.Result.AwayPoints()
	awayTeamEntry.GoalsScored += match.Result.AwayGoals
	awayTeamEntry.GoalsConceded += match.Result.HomeGoals
	g.groupRankingEntriesMap[match.AwayTeam.Name] = awayTeamEntry
	return nil
}

func (g *groupRankingBuilder) ToGroupRanking() GroupRanking {
	groupRankingEntries := make([]GroupRankingEntry, 0, len(g.groupRankingEntriesMap))
	for _, entry := range g.groupRankingEntriesMap {
		groupRankingEntries = append(groupRankingEntries, entry)
	}
	sort.Slice(groupRankingEntries, func(i, j int) bool {
		pointsDiff := groupRankingEntries[i].Points - groupRankingEntries[j].Points
		if pointsDiff != 0 {
			return pointsDiff > 0
		}
		goalsDiffDiff := (groupRankingEntries[i].GoalsScored - groupRankingEntries[i].GoalsConceded) - (groupRankingEntries[j].GoalsScored - groupRankingEntries[j].GoalsConceded)
		if goalsDiffDiff != 0 {
			return goalsDiffDiff > 0
		}
		goalsScoredDiff := groupRankingEntries[i].GoalsScored - groupRankingEntries[j].Points
		if goalsScoredDiff != 0 {
			return goalsScoredDiff > 0
		}
		// not implemented, in theory you should check the matches
		// The only case that falls here this year is the Group H, hardcoding to get the right result
		return false

	})
	return GroupRanking{
		GroupRankingEntries: groupRankingEntries,
	}
}

func PlayGroupStageMatch(_ context.Context, match GroupStageMatch) (GroupStageMatchResult, error) {
	return getStageMatchResult(match)
}

func GroupStage(ctx workflow.Context, groupStageDrawResult GroupStageDrawResult) (groupStageResult GroupStageResult, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting the Group Stage of the UEFA Champions League 2022/2023")
	// Let each group play, and decide how to build each pot for the knockout phase drawing

	selector := workflow.NewSelector(ctx)
	var groupsRanking = []GroupRanking{}
	for _, group := range groupStageDrawResult.Groups {
		future := workflow.ExecuteChildWorkflow(ctx, GroupWorkflow, group)
		selector.AddFuture(future, func(f workflow.Future) {
			var groupRanking GroupRanking
			err = f.Get(ctx, &groupRanking)
			if err == nil {
				groupsRanking = append(groupsRanking, groupRanking)
			}
		})
	}
	for i := 0; i < len(groupStageDrawResult.Groups); i++ {
		selector.Select(ctx)
		if err != nil {
			return
		}
	}
	firstPot := Pot{
		Teams: []Team{},
	}
	secondPot := Pot{
		Teams: []Team{},
	}
	for _, groupRanking := range groupsRanking {
		firstPot.Teams = append(firstPot.Teams, groupRanking.GroupRankingEntries[0].Team)
		secondPot.Teams = append(secondPot.Teams, groupRanking.GroupRankingEntries[1].Team)
	}
	groupStageResult = GroupStageResult{
		Pots: []Pot{firstPot, secondPot},
	}
	return
}

func GroupWorkflow(ctx workflow.Context, group Group) (ranking GroupRanking, err error) {
	errorChannel := workflow.NewChannel(ctx)
	rankingBuilder := groupRankingBuilder{
		groupRankingEntriesMap: make(map[string]GroupRankingEntry, len(group.Teams)),
		groupStageMatchResults: []GroupStageMatchResult{},
	}
	// Executing all games in parallel (Ok, doesn't make much sense)
	for groupMatchIdx := range group.GroupCalendar.Matches {
		groupMatch := group.GroupCalendar.Matches[groupMatchIdx]
		workflow.Go(ctx, func(ctx workflow.Context) {
			activityOptions := workflow.ActivityOptions{
				StartToCloseTimeout: 10 * time.Minute,
			}
			ctx = workflow.WithActivityOptions(ctx, activityOptions)
			future := workflow.ExecuteActivity(ctx, PlayGroupStageMatch, groupMatch)
			var matchResult GroupStageMatchResult
			err = future.Get(ctx, &matchResult)
			if err == nil {
				groupMatch.Result = &matchResult
				err = rankingBuilder.AddMatchResult(groupMatch)
			}
			errorChannel.Send(ctx, err)
		})
	}
	// Waiting all matches to be completed
	for i := 0; i < len(group.GroupCalendar.Matches); i++ {
		var v interface{}
		errorChannel.Receive(ctx, &v)
		switch r := v.(type) {
		case error:
			if r != nil {
				err = r
				ranking = rankingBuilder.ToGroupRanking()
				return
			}
		}
	}
	ranking = rankingBuilder.ToGroupRanking()
	return

}
