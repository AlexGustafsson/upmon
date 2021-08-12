# HTTP

The HTTP monitor performs a HTTP request to determine whether or not a service is up and working.

The monitor has the following options.

| Name | Description | Required |
| :--: | :---------: | :------: |
| `url` | The URL to request | Yes |
| `timeout` | The duration to wait before aborting a ping | No. Defaults to 1s |
| `interval` | The time to wait between each ping | No. Defaults to 1s |
| `method` | The HTTP method such as `GET` to use | No. Defaults to `GET` |
| `followRedirects` | Whether or not to follow redirects | No. Defaults to `false` |
| `maximumRedirects` | The maximum number of redirects to follow before throwing an error | No. Defaults to 10 |
| `expect` | An object containing the matching clauses to determine an alive service | |

The following expect clauses are available.

| Name | Description |
| :--: | :---------: |
| `status` | The expected HTTP status code |
| `regex` | A regular expression the response body is expected to match |
