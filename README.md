# Celebrating FC Internazionale fourth UEFA Champions League with Temporal

In this repo, I will put together two of my passions: software and football (or soccer, if you are american)
The team I support, FC Intenazionale has just won its fourth Champions League, and I thought it was a great occasion to foster my creativity

## Polyglot engineering and workflow orchestration

I have always been a fan of "using the right tool": in any non-trivial systems, it's likely that different components are easier to build, operate and maintain if you pick do not use a single technology stack.

However, if you are building higher-level systems from heterogeneous components, you need to ensure they work together "well". I understand "well" might have multiple meanings, and you'll excuse myself for this "non very technical explanation". I am thinking about the need to have long-running processes that spans days or months. During such a time interval, some components might fail, some other will need to be updated or patched with a bugfix, some others will might need to be rewritten.

Our industry tried for many years to address this problem through **workflow engines**. I worked with a couple of those and I never felt really satisfied for a variety of reasons: they were monolithic, hard to setup and operate, they lacked support for multiple programming languages, good testing facilities and they failed to scale.  As a consequence, I often ended up deliberately choosing to disseminate workflow logic into components, knowing I wasn't doing the right thing by coupling the higher level workflow logic with the components involved in the workflow. 

Sometimes the solution that serves us well in the long term comes with too many drawbacks in the short term, but wouldn't be good if we had a solution that's easy and effective in the short term, that also works well in the long term? 

## Meet Temporal

Temporal is a modern workflow engine that has addressed all the concerns I had in the past about workflow engines. 
Temporal world-class engineering team made several smart decisions in architecting the platform, including support natively multiple programming languages, providing stand-alone executable to start development servers, and imposing some constraints on developers about what is allowed in workflows.

My learning journey with Temporal made me remember a talk from Rúnar Bjarnason to the Scala community. The talk was named: "Freedom constraints, constraints liberate". I think there's something profound in software engineering about how constraints make our life easier and harder at the same time. When you face a new constraint for the first time, it might be frustrating. When you understand it, however, it could really become magic. You realize that accepting that constraint frees you up from so many problems!  

What really passionates me is writing software that has business impact and can change the world, and Temporal has uncovered for me new ways to do so. This is not the first time I encounter ah-ah moment in my learning journey and I am happy to share similar exciting discoveries I did in my career:

- How **functional programming** constrain you to avoid mutability but leads to lower errors and much more composable software
- How mastering Rust **ownership and lifetimes** is daunting upfront, but then leads you to write memory-safe programs with high performance
- How **event sourcing and CQRS** simplify writing applications that do not need strict consistency  


## What does it have to do with UEFA Champions League?

UEFA Champions League is a fairly long-running workflow, it typically takes 8-9 months to complete. The rules of the competition such as its structure, the number of teams involved, how teams are granted participation, has changed over time multiple times. The first time the "workflow was run" was 1955. 

How about modeling UEFA Champions League with Temporal? Ok, we won't have time to model all aspects of the competition, so that's what we are going to do.

- in [Chapter 1](./chapter1) we are going to look at create the competition workflow and understand the basics of Temporal
- in [Chapter 2](./chapter2) we are going to start modeling the structure of the competition, including its different phases
- in [Chapter 3](./chapter3) we are going to support retrieving the current status
- in [Chapter 4](./chapter4) we are going to look at modifying the competition to support future editions. 
- in [Chapter 5](./chapter5) we are going try to list all victories of FC Internazionale

Even if you are already familiar with Temporal, I'll recommend you jump to Chapter 1. It's [Chapter 1](./chapter1/README.md) contains useful information to follow along with the various chapters, so don't skip the readme even if you skip the chapter!

## Getting started

If you haven't got a working installation of Temporal, check out their docs at https://docs.temporal.io to get started.

Once you have it, you should be able to start a local server using

```shell
temporal server start-dev
```

The user interface will be available at `http://127.0.0.1:8233`. Open it and confirm there are no workflows visible

### The general workflow

Temporal is a workflow engine, so the core of what we do will be registering workflows and submitting executions of those workflows.
- Each chapter will contain a README.md that will guide you in your journey.
- We'll interact at the platform and look at code snippets
- We'll use the Temporal Server UI to understand what's happening, but you are free to use the CLI if you think user interfaces are the source of all evil (sometimes I do)


## A final note

GoLang is not my primary language. Please provide your feedback and improvements! 