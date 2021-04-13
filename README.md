# BlueMantis   
   
   
BlueMantis is a Go package in development that aim to make the process of sending issues and bugs in Go applications to the Open Source Bug Tracking software [MantisBT](http://mantisbt.org/).

This package came from my need to have a system that is more rubust than simply storing the logs locally in the server and simpler than using something like Sentry or DataDog. Also, it is always good to have the option to have the bug tracking software as part of your infrastructure instead of having it in a third party (specially if they are paid, like Sentry and DataDog are).

This is (at the moment at least) a one person small project. But I would love to have more people interested in participating in the development. First because I'm far from an expert in the subject, second because we  all have a limited amount of hours in a day.


### How to Contribute
Well, the first thing you need to know is that I manage the development of this using the [Projects tab](https://github.com/gustavohmsilva/bluemantis/projects/6) of this repository. Use it to know where to start focusing your development. If you wanto to have a more specific pain point to focus your help, [the issues tab is a good point](https://github.com/gustavohmsilva/bluemantis/issues). I add things I notice that I might need or want help.   
   
It is also a good point to check out the [wiki tab](https://github.com/gustavohmsilva/bluemantis/wiki) to know how to setup your environment, because currently you will need to use Docker to run a local version of [MantisBT](http://mantisbt.org/) to test BlueMantis.


### Usage

   
```
...
    btClient := bluemantis.NewClient(   
        "http://localhost:8989",
        "AAaaBBbbCCccDDddEEeeFFffGGgg--__",
    )

    btProject, err := bt.GetProject("Empty Project Alpha")
    if err != nil {
        panic(err)
    }

    issue, ok := bt.NewIssue(
        &bluemantis.BaseIssue{
            Summary:        "Empty Project",
            Description:    "This is an empty project, have you noticed that this is the only code in main?",
            Category:       "Critical",
            Project:        btProject,
        },
    ).Send().Retry().Delay()
...
```   
*currently! This is in VERY early development, and no function naming or style has been set on stone yet.
