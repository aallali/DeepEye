# Contributing to DeepEye

Thank you so much for contributing to **DeepEye**. We appreciate your time and help.
Here are some guidelines to help you get started.

## Code of Conduct

Be kind and respectful to the members of the community. Take time to educate
others who are seeking help. Harassment of any kind will not be tolerated.

## Questions

If you have questions regarding DeepEye, feel free to send me : hi+deepeye@allali.me

## Filing a bug or feature

1. Before filing an issue, please check the existing issues to see if a
   similar one was already opened. If there is one already opened, feel free
   to comment on it.
1. If you believe you've found a bug, please provide detailed steps of
   reproduction, the version of DeepEye and anything else you believe will be
   useful to help troubleshoot it (e.g. OS environment, environment variables,
   etc...). Also state the current behavior vs. the expected behavior.
1. If you'd like to see a feature or an enhancement please open an issue with
   a clear title and description of what the feature is and why it would be
   beneficial to the project and its users.

## Submitting changes

(currently there is no tests written so far, since i just tried to realise the idea i had in mind quickly, i am still working on them, but if you like to add something new to the app, please write some unit tests that cover its use case + the bugs that it fixes. THANKS!)
1. Tests: If you are submitting code, please ensure you have adequate tests
   for the feature. Tests can be run via `go test ./...` or `make test`.
1. Since this is golang project, ensure the new code is properly formatted to
   ensure code consistency. Run `make all`.

### Quick steps to contribute

1. Fork the project.
1. Download your fork to your PC (`git clone https://github.com/your_username/deepeye && cd deepeye`)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Make changes and run tests (`make test`)
1. Add them to staging (`git add .`)
1. Commit your changes (`git commit -m 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new pull request

<!-- Links -->
