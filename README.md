# dummy-ms

[![CircleCI](https://circleci.com/gh/danielpsf/go-dummy-ms.svg?style=svg)](https://circleci.com/gh/danielpsf/go-dummy-ms) [![codecov](https://codecov.io/gh/danielpsf/go-dummy-ms/branch/master/graph/badge.svg)](https://codecov.io/gh/danielpsf/go-dummy-ms) 

dummy-ms aims to be a simple micro-service which can simulate a variety of scenarios such as:
- Timeouts
- Computer misbehave, like:
  - Excessive memory consumption
  - Excessive CPU consumption
  - Excessive network consumption
- etc

By having a middleware capable of such things we could easily test our infrastructure knowing how it will react when a component of a business case misbehave.

Well, of course this is also a pet project so we could learn Golang best practices :wink:

## How to build and develop
Please follow the below instructions to safely prepare your environment for development

### How to build your local environment
- [Install Go](https://golang.org/doc/install)
- Checkout the repository
- Have fun

### Running the unit tests and coverage
To run the unit tests and generate the coverage report (html), navigate to the project's root folder and execute `go test -cover -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`

### How to debug locally
- Configure your [Visual Studio Code](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code) so you can be able to debug the code locally.

### Before submitting a Pull Request or making a `git push`, please
- Check if your code passes on the checks of [Effective Go](https://golang.org/doc/effective_go.html)
    - Configure your [Visual Studio Code](https://github.com/Microsoft/vscode-go/wiki/On-Save-features) to execute `go build`, `go fmt` and `golint` on save

## FAQ

### Why ginko and gomega instead of native go unit testing
Well, I rather allow the tests to be our live documentation, so BDD makes it easier by having a better description of what is happening. [ginko](http://onsi.github.io/ginkgo/) and [gomega](http://onsi.github.io/gomega/) allow us to have better description through human readable test description and matchers (evaluations) and deliver great development capabilities with amazing assertions and nice tooling for TDD, such as `$ ginko watch`.

## References
- https://github.com/bxcodec/go-clean-arch
- https://github.com/rosenhouse/counter-demo
- https://medium.com/@zach_4342/dependency-injection-in-golang-e587c69478a8
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- https://github.com/cevatbarisyilmaz/lossy (network degradation tool)
