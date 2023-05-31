# Dealing with a scandal

The team in charge of the UEFA Champions League gets dismissed because of a corruption scandal. Ok, I know this hasn't happened in this competition, but recent events suggest this is not an impossible event. 

In complex software systems, this happen all the time: a process carrying out some work might fail, the underlying hardware might fail, the underlying availability zone might fail. Everything fails all the time. For long-running processes, it will be a real pity if we had to start processing all over again (and it might be **very expensive**)

## Introducing some suspense

We now FC Internazionale has been the winner, but for this chapter, we are going to introduce some suspense. We want he result to be returned after one hour from the beginning of the competition (so to allow bookmarkers to collect some bets, obviously.)

Now what happen if you have a Java process like so that fails in the middle and you run it again?

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

Well, you lost track of how much time the previous process was waiting when it crashed, and you will wait again for one hour. This is not acceptable.

### Workflow executions

(Most of this chapter is a copy-paste of Temporal documentation)

A Temporal Workflow Execution is a durable, reliable, and scalable function execution. I'll quote the documentation here:
- Durability is the absence of an imposed time limit.
- Reliability is responsiveness in the presence of failure.
- Scalability is responsiveness in the presence of load.

The use of Workflow APIs in a Workflow Function generates Commands. Commands tell the Cluster which Events to create and add to the Workflow Execution's Event History. When a Workflow Function executes, the Commands that are emitted are compared with the existing Event History. If a corresponding Event already exists within the Event History that maps to the generation of that Command in the same sequence, and some specific metadata of that Command matches with some specific metadata of the Event, then the Function Execution progresses

A special treatment 


