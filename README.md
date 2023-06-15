# OpenAPI request validator

Package to validate if a Golang request is valid given an OpenAPI definition.

It validates parameters either in the URL and in the payload as well as other stuff like content-type, security schema, etc.

You could also validate the response if necessary, so you can be sure that your response is always valid according to the OpenAPI definition.
