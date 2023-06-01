# Dealing with a scandal

The team in charge of the UEFA Champions League gets dismissed because of a corruption scandal. Ok, I know this hasn't happened in this competition, but recent events suggest this is not an impossible event. 

In complex software systems, this happens all the time: a process carrying out some work might fail, the underlying hardware might fail, the underlying availability zone might fail. Everything fails all the time. For long-running processes, it will be a real pity if we had to start processing all over again (and it might be **very expensive**)

## Introducing some suspense

We now FC Internazionale has been the winner, but for this chapter, we are going to introduce some suspense. We want the result to be returned after one hour from the beginning of the competition (so to allow bookmakers to collect some bets, obviously.)

Now what happen if you have a Java process like so that fails in the middle, and you run it again?

```java
import java.time.Duration;
import java.time.temporal.ChronoUnit;

public class ChampionsLeague {
    
    public static void main(String[] args) throws Exception{
        Duration oneHour = Duration.of(1, ChronoUnit.HOURS);
        Thread.sleep(oneHour.toMillis());
        System.out.println("FC Internazionale is the winner!");
    }
}
```

Well, you lost track of how much time the previous process was waiting when it crashed, and you will wait again for one 
hour.  

### Understanding Workflow Executions

(Please refer to Temporal documentation for a less "layman" explanation of this topic)

If you want to make your workflow durable and reliable, you can't rely on local state. The store need to be saved
remotely, so it can be restored on a new worker if required. In particular, in Temporal you do that by invoking the
Workflow API, that allows you to send **Commands** to the server. The server interpret those commands and store
**event**.

When a worker dies and a different worker gets assigned the difficult job of completing the work that someone
else has started (we all know how annoying that could be, couldn't you have included me from the beginning, so I could 
understand what this is about?), it is necessary to restore the state of the workflow. The way Temporal achieve this is
by executing the Workflow code and re-issuing the commands. The Commands that are emitted are
compared with the existing Event History. If a corresponding Event already exists within the Event History that maps to
the generation of that Command in the same sequence, and some specific metadata of that Command matches with some
specific metadata of the Event, then the Function Execution progresses

### Replacing native sleep with Temporal sleep

Using Temporal makes software engineer sleep better. They do not have to care about long-running transaction getting
lost, so let's proceed in this direction

```go

package chapter2

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func ChampionsLeague(ctx workflow.Context, participants Participants) (Result, error) {
	workflow.Sleep(ctx, time.Second*60)
	return Result{
		Winner: Team{Name: "FC Internazionale"},
	}, nil
}


```

First of all, interestingly enough, we can see that our Unit Test terminates immediately. How is that possible? The 
unit test log says:

> DEBUG Auto fire timer TimerID 1 TimerDuration 1m0s TimeSkipped 1m0s

So the SDK test support makes easier to test asynchronous code by firing the timer immediately. Let's now try to run the
workflow against a Temporal cluster. If you are still running the old instance, issue a new `temportal server start-dev`
after killing the previous process, then issue the usual registration of the workflow and its submission

```shell

cd management_team && go run  main.go ## in one terminal
cd main && go run  edmondo.go ## in another terminal

```



### Introducing some suspense correctly




