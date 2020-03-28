
# Introduction

Fluxhub executes jobs in response to [Fluxcd](https://fluxcd.io) events. This
allows you to perform tasks such as running a test suite, or sending an email
in response to a Fluxcd release.

Jobs are matched by the Flux event type and workload. So to trigger a job for
the flux `autorelease` event on the `webservice` workload you would create a
match object:

```
match:
  events: ['autorelease']
  workloads: ['webservice']
```

Both the event and workload values are matched using globbing, so to match
all dev workloads you could use `*-dev` depending on your naming conventions
for cluster resources.


# Defining Jobs

```yaml
---
jobs:
  - name: 'Run Tests'
    match:
      events:
        - release
        - autorelease
      workloads: ['*-dev']
    docker:
      image: curlimages/curl
      command: ['/usr/bin/curl']
      args: ['https://github.com']
```
