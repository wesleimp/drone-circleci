Use this plugin for trigger cirlceci builds.

## Settings

- `token` - Circleci access token
- `user` - Repository user
- `repo` - Repository name
- `branch` - Git branch for current commit

## Example

```yml
steps:
  - name: Circleci build
    image: wesleimp/drone-circleci
    settings:
      token: key-1234567890123567490
      user: octocat
      repo: hello-world
      branch: master
```