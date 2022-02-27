# tango_romeo

    tl;dr: _tango_romeo_ is a easy way to maintain your team's tech stack decisions.

_tango_romeo_ is my implementation of tech radar. It basically has three things any team needs to maintain in terms of deciding between technologies: list of technologies, their types and their level/status of acceptance across the team. I have a few things in mind to add to the idea, few of them are following but not limited to:

- managing tech stack for multiple teams 
- integrate Slack.
- polling on tech adoption


_[Note: features are still being explored, feel free to suggest your idea through the issues]_


## Features

### Visualization
- Have a visual representation of the technologies used in different teams, as well as their adoption status.

### API
- Fetching tech stack info through API is helpful when it comes to integrating your decision making process with different systems.
### Suggestions and Team Poll 
- Collect suggestions of technology preference from teammates and experts. 
- Make decision on which to pick or avoid through team-wide polls on Tango Romeo dashboard. Generate links/invite teammates to cast vote.
- Keep a log of past decisions.

### Slack integration
- Get summarized version of tech decisions and adoption status easily via Slack bot.

## Inspiration

**Different implementations**
- [ThoughtWork - Build Your Own Radar](https://www.thoughtworks.com/radar/)
- [Zalando Tech Radar](https://opensource.zalando.com/tech-radar/)

**Tech stack sharing**
- [Stackshare](https://stackshare.io/)

## Why am I doing this?

It's a personal project to practice Go and VueJS. During a discussion at a workplace, we realized the way we had been exploring different tech has been scattered and without any proper vetting, resulting in fractured usage across different dev teams. At the end we ended up doing nothing regarding this but for me it was motivation enough to create something useful :) However I do hope it becomes something others can contribute to and benefit from, it would be cool to see a community build around the concept.

## What's with the name?

"Tango Romeo" is NATO spelling for TR, which in turn means Tech Radar. I Honestly I couldn't find a good name for the project and didn't want to waste more than an hour. As the old adage goes, if there are two hard things in Computer Science, they are writing git commit messages and naming things.
