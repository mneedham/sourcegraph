// Code generated by stringdata. DO NOT EDIT.

package schema

// CoreSchemaJSON is the content of the file "core.schema.json".
const CoreSchemaJSON = `{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"$id": "core.schema.json#",
	"title": "Core site configuration",
	"description": "Core configuration for a Sourcegraph site.",
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"auth.userOrgMap": {
			"description": "Ensure that matching users are members of the specified orgs (auto-joining users to the orgs if they are not already a member). Provide a JSON object of the form ` + "`" + `{\"*\": [\"org1\", \"org2\"]}` + "`" + `, where org1 and org2 are orgs that all users are automatically joined to. Currently the only supported key is ` + "`" + `\"*\"` + "`" + `.",
			"type": "object",
			"additionalProperties": {
				"type": "array",
				"items": {
					"type": "string"
				}
			}
		},
		"log": {
			"description": "Configuration for logging and alerting, including to external services.",
			"type": "object",
			"additionalProperties": false,
			"properties": {
				"sentry": {
					"description": "Configuration for Sentry",
					"type": "object",
					"additionalProperties": false,
					"properties": {
						"dsn": {
							"description": "Sentry Data Source Name (DSN). Per the Sentry docs (https://docs.sentry.io/quickstart/#about-the-dsn), it should match the following pattern: '{PROTOCOL}://{PUBLIC_KEY}@{HOST}/{PATH}{PROJECT_ID}'.",
							"type": "string",
							"pattern": "^https?://"
						}
					}
				}
			}
		},
		"appURL": {
			"description": "Publicly accessible URL to web app (e.g., what you type into your browser).",
			"type": "string"
		},
		"tls.letsencrypt": {
			"description": "Toggles ACME functionality for automatically using a TLS certificate issued by the Let's Encrypt Certificate Authority.\nThe default value is auto, which uses the following conditions to switch on:\n - tlsCert and tlsKey are unset.\n - appURL is a https:// URL\n - Can successfully bind to port 443",
			"default": "auto",
			"enum": [ "auto", "on", "off" ],
			"type": "string"
		},
		"tlsCert": {
			"description": "The contents of the PEM-encoded TLS certificate for the web server (for the web app and API).\n\nSee https://about.sourcegraph.com/docs/config/tlsssl/ for more information.",
			"type": "string",
			"pattern": "^-----BEGIN CERTIFICATE-----\n"
		},
		"tlsKey": {
			"description": "The contents of the PEM-encoded TLS key for the web server (for the web app and API).\n\nSee https://about.sourcegraph.com/docs/config/tlsssl/ for more information.",
			"type": "string",
			"pattern": "^-----BEGIN "
		},
		"httpToHttpsRedirect": {
			"description": "Redirect users from HTTP to HTTPS. Accepted values are \"on\", \"off\", and \"load-balanced\" (boolean values true and false are also accepted and equivalent to \"on\" and \"off\" respectively). If \"load-balanced\" then additionally we use \"X-Forwarded-Proto\" to determine if on HTTP.",
			"anyOf": [ { "type": "string", "enum": [ "on", "off", "load-balanced" ] }, { "type": "boolean" } ],
			"default": "off"
		},
		"httpStrictTransportSecurity": {
			"description": "The value of the HTTP Strict-Transport-Security (HSTS) header sent by Sourcegraph, if non-empty",
			"anyOf": [ { "type": "string" }, { "type": "boolean" } ],
			"default": "max-age=31536000"
		},
		"lightstepAccessToken": {
			"description": "Access token for sending traces to LightStep.",
			"type": "string"
		},
		"lightstepProject": {
			"description": "The project ID on LightStep that corresponds to the ` + "`" + `lightstepAccessToken` + "`" + `, only for generating links to traces. For example, if ` + "`" + `lightstepProject` + "`" + ` is ` + "`" + `mycompany-prod` + "`" + `, all HTTP responses from Sourcegraph will include an X-Trace header with the URL to the trace on LightStep, of the form ` + "`" + `https://app.lightstep.com/mycompany-prod/trace?span_guid=...&at_micros=...` + "`" + `.",
			"type": "string"
		},
		"useJaeger": {
			"description": "Use local Jaeger instance for tracing. Kubernetes cluster deployments only.\n\nAfter enabling Jaeger and updating your Kubernetes cluster, ` + "`" + `kubectl get pods` + "`" + `\nshould display pods prefixed with ` + "`" + `jaeger-cassandra` + "`" + `,\n` + "`" + `jaeger-collector` + "`" + `, and ` + "`" + `jaeger-query` + "`" + `. ` + "`" + `jaeger-collector` + "`" + ` will start\ncrashing until you initialize the Cassandra DB. To do so, do the\nfollowing:\n\n1. Install [` + "`" + `cqlsh` + "`" + `](https://pypi.python.org/pypi/cqlsh).\n1. ` + "`" + `kubectl port-forward $(kubectl get pods | grep jaeger-cassandra | awk '{ print $1 }') 9042` + "`" + `\n1. ` + "`" + `git clone https://github.com/uber/jaeger && cd jaeger && MODE=test ./plugin/storage/cassandra/schema/create.sh | cqlsh` + "`" + `\n1. ` + "`" + `kubectl port-forward $(kubectl get pods | grep jaeger-query | awk '{ print $1 }') 16686` + "`" + `\n1. Go to http://localhost:16686 to view the Jaeger dashboard.",
			"type": "boolean"
		},
		"htmlHeadTop": {
			"description": "HTML to inject at the top of the ` + "`" + `<head>` + "`" + ` element on each page, for analytics scripts",
			"type": "string"
		},
		"htmlHeadBottom": {
			"description": "HTML to inject at the bottom of the ` + "`" + `<head>` + "`" + ` element on each page, for analytics scripts",
			"type": "string"
		},
		"htmlBodyTop": {
			"description": "HTML to inject at the top of the ` + "`" + `<body>` + "`" + ` element on each page, for analytics scripts",
			"type": "string"
		},
		"htmlBodyBottom": {
			"description": "HTML to inject at the bottom of the ` + "`" + `<body>` + "`" + ` element on each page, for analytics scripts",
			"type": "string"
		},
		"licenseKey": {
			"description": "The license key associated with a Sourcegraph product subscription, which is necessary to activate Sourcegraph Enterprise functionality. To obtain this value, contact Sourcegraph to purchase a subscription.",
			"type": "string"
		},
		"auth.providers": {
			"description": "The authentication providers to use for identifying and signing in users.",
			"type": "array",
			"items": {
				"required": [ "type" ],
				"properties": {
					"type": {
						"type": "string",
						"enum": [ "builtin", "saml", "openidconnect", "http-header" ]
					}
				},
				"oneOf": [
					{ "$ref": "#/definitions/BuiltinAuthProvider" },
					{ "$ref": "#/definitions/SAMLAuthProvider" },
					{ "$ref": "#/definitions/OpenIDConnectAuthProvider" },
					{ "$ref": "#/definitions/HTTPHeaderAuthProvider" }
				],
				"!go": {
					"taggedUnionType": true
				}
			}
		},
		"auth.public": {
			"description": "Allows anonymous visitors full read access to repositories, code files, search, and other data (except site configuration).\n\nSECURITY WARNING: If you enable this, you must ensure that only authorized users can access the server (using firewall rules or an external proxy, for example).\n\nRequires usage of the builtin authentication provider.",
			"type": "boolean",
			"default": false
		},
		"update.channel": {
			"description": "The channel on which to automatically check for Sourcegraph updates.",
			"type": [ "string" ],
			"enum": [ "release", "none" ],
			"default": "release"
		}
	},
	"definitions": {
		"BuiltinAuthProvider": {
			"description": "Configures the builtin username-password authentication provider.",
			"type": "object",
			"additionalProperties": false,
			"required": [ "type" ],
			"properties": {
				"type": {
					"type": "string",
					"const": "builtin"
				},
				"allowSignup": {
					"description": "Allows new visitors to sign up for accounts. The sign-up page will be enabled and accessible to all visitors.\n\nSECURITY: If the site has no users (i.e., during initial setup), it will always allow the first user to sign up and become site admin **without any approval** (first user to sign up becomes the admin).",
					"type": "boolean",
					"default": false
				}
			}
		},
		"OpenIDConnectAuthProvider": {
			"description": "Configures the OpenID Connect authentication provider for SSO.",
			"type": "object",
			"additionalProperties": false,
			"required": [ "type", "issuer", "clientID", "clientSecret" ],
			"properties": {
				"type": {
					"type": "string",
					"const": "openidconnect"
				},
				"displayName": { "$ref": "#/definitions/AuthProviderCommon/properties/displayName" },
				"configID": {
					"description": "An identifier that can be used to reference this authentication provider in other parts of the config. For example, in configuration for a code host, you may want to designate this authentication provider as the identity provider for the code host.",
					"type": "string"
				},
				"issuer": {
					"description": "The URL of the OpenID Connect issuer.\n\nFor Google Apps: https://accounts.google.com",
					"type": "string",
					"format": "uri",
					"pattern": "^https?://"
				},
				"clientID": {
					"description": "The client ID for the OpenID Connect client for this site.\n\nFor Google Apps: obtain this value from the API console (https://console.developers.google.com), as described at https://developers.google.com/identity/protocols/OpenIDConnect#getcredentials",
					"type": "string",
					"pattern": "^[^<]"
				},
				"clientSecret": {
					"description": "The client secret for the OpenID Connect client for this site.\n\nFor Google Apps: obtain this value from the API console (https://console.developers.google.com), as described at https://developers.google.com/identity/protocols/OpenIDConnect#getcredentials",
					"type": "string",
					"pattern": "^[^<]"
				},
				"requireEmailDomain": {
					"description": "Only allow users to authenticate if their email domain is equal to this value (example: mycompany.com). Do not include a leading \"@\". If not set, all users on this OpenID Connect provider can authenticate to Sourcegraph.",
					"type": "string",
					"pattern": "^[^<@]"
				}
			}
		},
		"SAMLAuthProvider": {
			"description": "Configures the SAML authentication provider for SSO.\n\nNote: if you are using IdP-initiated login, you must have *at most one* SAMLAuthProvider in the ` + "`" + `auth.providers` + "`" + ` array.",
			"type": "object",
			"additionalProperties": false,
			"required": [ "type" ],
			"dependencies": {
				"serviceProviderCertificate": [ "serviceProviderPrivateKey" ],
				"serviceProviderPrivateKey": [ "serviceProviderCertificate" ],
				"signRequests": [ "serviceProviderCertificate", "serviceProviderPrivateKey" ]
			},
			"properties": {
				"type": {
					"type": "string",
					"const": "saml"
				},
				"configID": {
					"description": "An identifier that can be used to reference this authentication provider in other parts of the config. For example, in configuration for a code host, you may want to designate this authentication provider as the identity provider for the code host.",
					"type": "string"
				},
				"displayName": { "$ref": "#/definitions/AuthProviderCommon/properties/displayName" },
				"serviceProviderIssuer": {
					"description": "The name of this SAML Service Provider, which is used by the Identity Provider to identify this Service Provider. It defaults to https://sourcegraph.example.com/.auth/saml/metadata (where https://sourcegraph.example.com is replaced with this Sourcegraph instance's \"appURL\"). It is only necessary to explicitly set the issuer if you are using multiple SAML authentication providers.",
					"type": "string"
				},
				"identityProviderMetadataURL": {
					"description": "The SAML Identity Provider metadata URL (for dynamic configuration of the SAML Service Provider).",
					"type": "string",
					"format": "uri",
					"pattern": "^https?://"
				},
				"identityProviderMetadata": {
					"description": "The SAML Identity Provider metadata XML contents (for static configuration of the SAML Service Provider). The value of this field should be an XML document whose root element is ` + "`" + `<EntityDescriptor>` + "`" + ` or ` + "`" + `<EntityDescriptors>` + "`" + `.",
					"type": "string"
				},
				"serviceProviderCertificate": {
					"description": "The SAML Service Provider certificate in X.509 encoding (begins with \"-----BEGIN CERTIFICATE-----\"). This certificate is used by the Identity Provider to validate the Service Provider's AuthnRequests and LogoutRequests. It corresponds to the Service Provider's private key (` + "`" + `serviceProviderPrivateKey` + "`" + `).",
					"type": "string",
					"$comment": "The pattern matches either X.509 encoding or an env var.",
					"pattern": "^(-----BEGIN CERTIFICATE-----\n|\\$)",
					"minLength": 1
				},
				"serviceProviderPrivateKey": {
					"description": "The SAML Service Provider private key in PKCS#8 encoding (begins with \"-----BEGIN PRIVATE KEY-----\"). This private key is used to sign AuthnRequests and LogoutRequests. It corresponds to the Service Provider's certificate (` + "`" + `serviceProviderCertificate` + "`" + `).",
					"type": "string",
					"$comment": "The pattern matches either PKCS#8 encoding or an env var.",
					"pattern": "^(-----BEGIN PRIVATE KEY-----\n|\\$)",
					"minLength": 1
				},
				"nameIDFormat": {
					"description": "The SAML NameID format to use when performing user authentication.",
					"type": "string",
					"pattern": "^urn:",
					"default": "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
					"examples": [
						"urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
						"urn:oasis:names:tc:SAML:1.1:nameid-format:persistent",
						"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified",
						"urn:oasis:names:tc:SAML:2.0:nameid-format:emailAddress",
						"urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
						"urn:oasis:names:tc:SAML:2.0:nameid-format:transient",
						"urn:oasis:names:tc:SAML:2.0:nameid-format:unspecified"
					]
				},
				"signRequests": {
					"description": "Sign AuthnRequests and LogoutRequests sent to the Identity Provider using the Service Provider's private key (` + "`" + `serviceProviderPrivateKey` + "`" + `). It defaults to true if the ` + "`" + `serviceProviderPrivateKey` + "`" + ` and ` + "`" + `serviceProviderCertificate` + "`" + ` are set, and false otherwise.",
					"type": "boolean",
					"!go": { "pointer": true }
				},
				"insecureSkipAssertionSignatureValidation": {
					"description": "Whether the Service Provider should (insecurely) accept assertions from the Identity Provider without a valid signature.",
					"type": "boolean",
					"default": false
				}
			}
		},
		"HTTPHeaderAuthProvider": {
			"description": "Configures the HTTP header authentication provider (which authenticates users by consulting an HTTP request header set by an authentication proxy such as https://github.com/bitly/oauth2_proxy).",
			"type": "object",
			"additionalProperties": false,
			"required": [ "type", "usernameHeader" ],
			"properties": {
				"type": {
					"type": "string",
					"const": "http-header"
				},
				"usernameHeader": {
					"description": "The name (case-insensitive) of an HTTP header whose value is taken to be the username of the client requesting the page. Set this value when using an HTTP proxy that authenticates requests, and you don't want the extra configurability of the other authentication methods.",
					"type": "string"
				}
			}
		},
		"AuthProviderCommon": {
			"$comment": "This schema is not used directly. The *AuthProvider schemas refer to its properties directly.",
			"description": "Common properties for authentication providers.",
			"type": "object",
			"properties": {
				"displayName": {
					"description": "The name to use when displaying this authentication provider in the UI. Defaults to an auto-generated name with the type of authentication provider and other relevant identifiers (such as a hostname).",
					"type": "string"
				}
			}
		}
	}
}
`
