# Sysinfo server

A simple server for retrieving system bootup duration. The returned
duration is expressed in seconds.

Usage example:

	$ ./sysinfo_server &
	[1] 30177

	$ curl http://localhost:8080
	Welcome to sysinfo_server, version 0.0.1
	Available endpoints: /version, /duration

	$ curl http://localhost:8080/version
	0.0.1

	$ curl http://localhost:8080/duration
	170.94

## Implementation note.

The response of `duration` does not report the unit in order to
be easier to parse, and the choosen unit, seconds, is the
one that satisfy better the principle of less surprise.

Duration are retrieved using `systemd-analyze` and parsing its
output. The parsing is done using regular expressions and is
currently the weakest part of the code.
In order to keep it simply, several assumption are made about
the output format of `systemd-analyze`.

## Possible improvements

Providing the output in json would be possible in several ways.
The most easy way would be by using the `Gin` framework, that
would offer other benefits as the application grows in complexity.

Without using a framework it would still be possible to use
`encoding/json` for structuring more complex data.
