

# Defining Actions

```yaml
---
jobs:
  - name: 'Run Tests'
    match:
      service: '*-dev'
      namespace: '*-dev'
      events:
        - release
        - autorelease
    docker:
      image: curlimages/curl
      command: ['/usr/bin/curl']
      args: ['https://github.com']
```
