<h1 align="center">♥️ Welcome and thanks for contributing</h1>
<p align="center">
    <strong>
        First please read our <a href="https://github.com/olehan/kek/blob/master/CODE_OF_CONDUCT.md">code of conduct</a> to proceed.
    </strong>
</p>

----

<p align="center">
    <strong><a href="#setup">Setup</a></strong>
    |
    <strong><a href="#linting-testing">Linting/Testing</a></strong>
    |
    <strong><a href="#writing-tests">Writing Tests</a></strong>
    |
    <strong><a href="#internals">Internals</a></strong>
</p>

----

<h2 id="chat">Setup</h2>

We use go modules, so requirements for this project is `go version >= 1.11.0`.

And don't forget to enable `GO111MODULE`:
```sh
$ export GO111MODULE=on
```

***Note:*** pre-requirements for linting is `golang.org/x/lint/golint`, so first install it:
```sh
$ go get -u golang.org/x/lint/golint
```

And then:

```sh
$ git clone https://github.com/olehan/kek
$ cd kek
$ make bootstrap
```


<h2 id="linting-testing">Linting/Testing</h2>

***Lint your code by running:***
```sh
$ make lint
```

***And test it via:***
```sh
$ make test
```

***To lint and test:***
```sh
$ make check
```


<h2 id="writing-tests">Writing Tests</h2>

Tests are written along with site the file it's testing.

For example, if your package `request` has a `service.go` in it
tests would be written in `service_test.go` and so on.


<h2 id="internals">Internals</h2>

* ***Project Architecture*** - [golang clean architecture](https://github.com/bxcodec/go-clean-arch) | ***Struct:***
  ```
  domain
  |   model.go
  |   repo.go
  └── service.go
  ```
  or
  ```
  domain
  |   {sub_domain}_model.go
  |   {sub_domain}_repo.go
  └── {sub_domain}_service.go
  ```

* ***Commit messages convention*** -
[conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0-beta.4/#examples) | ***Examples:***

* ***Branch naming convention*** -
`type/title` | ***Examples:*** `feat/perform-request`, `ci/stale-bot`, `docs/contributing`, `fix/response-typo`

* ***Tar version names*** -
[semver](https://semver.org/) | ***Examples:*** `v1.2.3`, `v2.3.1-rc.3`, `v4.21.2-beta.4`

* ***Why space indention?*** -
[this](https://softwareengineering.stackexchange.com/a/66) answer from stackoverflow describes the idea.