# Implementing Group Stages and Knockout Phase

I updated the model structure to make easier to browse the code and extended the structs to hold more information I 
needed to proceed writing the business logic. I moved some reference data code in reference_data, just to simplify 
browsing the code.

## Testing our Drawing Activity

Now that we have the code for the groups, we can correctly implement the DrawGroupStage activity like so

```go
func (d *GroupStageDrawing) DrawGroups(ctx context.Context, input GroupStageDrawInput) (GroupStageDrawResult, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Drawing group stages in %s with a total of %d teams", d.Name, input.Participants.TeamCount())
	logger.Info(msg)
	return GroupStageDrawResult{
		Groups: [8]Group{
			GroupA(),
			GroupB(),
			GroupC(),
			GroupD(),
			GroupE(),
			GroupF(),
			GroupG(),
			GroupH(),
		},
	}, nil
}
```

and we can proceed to testing it as well. Let's compare testing workflows and activities, since they are pretty similar.
We always start from the entry point in Temporal SDK, `testsuite.WorkflowTestSuite`, and we:
 - Invoke `NewTestWorkflowEnvironment` that supports an `ExecuteWorkflow` API, and allows us to mock activities using the 
`OnActivity`
 - Invoke `NewTestActivityEnvironment` that supports `RegisterActivity` and `ExecuteActivity`


```go
func TestGroupStageDraw(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	var drawing = &GroupStageDrawing{Name: "Monte Carlo"}
	env.RegisterActivity(drawing.DrawGroups)
	groups, err := env.ExecuteActivity(drawing.DrawGroups, GroupStageDrawInput{
		Participants: GetParticipants(),
	})
	require.NoError(t, err)
	var res *GroupStageDrawResult
	require.NoError(t, groups.Get(&res))
	require.Equal(t, 8, len(res.Groups))
}
```

Now that our drawing activity, we can use its input to create the Group Stage

## Relationships are important (although sometimes difficult!)

If we really wouldn't value relationships, we would try to stick all the code that implement the group stage logic,
the knockout phase logic, and even the final in the core `ChampionsLeague` workflow. But we are good people, and we also
know that building large modules in software development is a bad idea, so we would structure our system in this way:
- A GroupStage workflow, which needs to execute all the groups and return the qualified teams
- A Group workflow, which can handle all the matches within that group

As I mention, relationships are sometimes difficult and they are not always the right answer. Please refer to
https://docs.temporal.io/workflows#when-to-use-child-workflows to learn more

### And expressing ourselves is important too

That's why we add logging in software. Software is complex, and sometimes it doesn't express feelings or emotions. Over
time, it can get pretty bad. Logging is the way we address this problem: we make our software talk. However, in the case
of resilient distributed applications with Temporal, we know that the work can be distributed among multiple workers and
that the state of a workflow is restored by event history. How does that impact logging? Well, you don't want to 
**log again**, so you need to use a specialized logger, like so:

```go
func GroupStage(ctx workflow.Context, groupStageDrawResult GroupStageDrawResult) (GroupStageResult, error) {
    logger := workflow.GetLogger(ctx)
    logger.Info("Starting the Group Stage of the UEFA Champions League 2022/2023")
    // Let each group play, and decide how to build each pot for the knockout phase drawing
    return GroupStageResult{}, nil
}
```

### Back to relationships 

So what is our Group Stage workflow going to do? Well, simply, it should submit an instance of a group workflow for each
group, wait for it to terminate, and determine which teams are going to get into the knockout phase. Although the logic
 is pretty simple (the two teams that are at the top of the group ranking make into the final), you can imagine this
is the place where we could handle, for example, a disqualified team (let's say a financial fraud, for example).

What I want to stress out here is that the logic is deterministic: once all the groups have played, there is a 
well-defined process to decide which team makes into the knockout phase
