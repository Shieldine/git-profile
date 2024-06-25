# git-profile
Simple git plugin to manage and automatically set git user profiles based on the project's origin.

Okay, but why?<br />
Some developers use their computers for both work-related and private projects.
This usually involves having *at least* two different sets of credentials
for git. If your private project and/or work involve multiple platforms to keep
your source-code on, that increases the number of credential sets you need to manage.
Setting those each time you clone a project can be quite tiresome and - if you forget or misspell something -
lead to the need of amending commits.

With git-profile, you need to type in your attributes exactly *once*.
They get saved in a *profile* along with the project's origin. Upon calling git-profile
in a repository, it will automatically pick a profile based on the origin
and simply set those attributes for you - no need to even remember a profile name!

At the same time, you still get a few little extra commands to have manual control
over your attributes.


## Installation


## Getting started


## Development
