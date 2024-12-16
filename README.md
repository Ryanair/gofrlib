### Common library for Go projects

This project was cloned from github.
To be able to use it in your project, you need to execute the following steps:

1. Set bitbucket as a private repository:\
   `go env -w GOPRIVATE='bitbucket.org/ryanair/*'`
2. Use ssh instead of https:\
 `git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"`

### Codebuild

Add these to the buildspec.yml file:

To env:

```
 env:
  git-credential-helper: yes
```

To commands:

`
      - go env -w GOPRIVATE='bitbucket.org/ryanair/*'
`


Example of a buildspec.yml file:

```
version: 0.2
env:
  git-credential-helper: yes
phases:
  install:
    runtime-versions:
      golang: 1.22
    commands:
      - go env -w GOPRIVATE='bitbucket.org/ryanair/*'
      - aws ecr get-login-password --region $AWS_DEFAULT_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com
      - nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay2 &
      - timeout 15 sh -c "until docker info; do echo .; sleep 1; done"
      - go get ./...

  build:
    on-failure: ABORT
    commands:
      - make build

artifacts:
  files:
    - 'deployments/**/*'
    - 'Makefile'
    - 'deployspec.yaml'

cache:
  paths:
    - '/go/pkg/mod/**/*'
```