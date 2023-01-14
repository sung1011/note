# github-flow

```flow
       a---b---c  *feature1
      /         \ PR
 D---E---F---G---H---I---K---L---M  *master
              \     / PR
               x---y  *feature2
```

## branch

### main/master

1. The `main` branch stores the official release history

### feature

1. `Feature` branches are created from `main`
1. When a `feature` is complete it submit a `PR` (pull request) to the `main` branch
1. When the review or discussion is approved, the `PR` is merged into the `main` branch