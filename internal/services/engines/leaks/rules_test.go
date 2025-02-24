// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leaks

import (
	"testing"

	engine "github.com/ZupIT/horusec-engine"
	"github.com/ZupIT/horusec/internal/utils/testutil"
)

func TestRulesVulnerableCode(t *testing.T) {
	testcases := []*testutil.RuleTestCase{
		{
			Name: "HS-LEAKS-1",
			Rule: NewAWSManagerID(),
			Src:  SampleVulnerableHSLEAKS1,
			Findings: []engine.Finding{
				{
					CodeSample: "ACCESS_KEY: 'AKIAJSIE27KKMHXI3BJQ'",
					SourceLocation: engine.Location{
						Line:   7,
						Column: 18,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-2",
			Rule: NewAWSSecretKey(),
			Src:  SampleVulnerableHSLEAKS2,
			Findings: []engine.Finding{
				{
					CodeSample: `AWS_SECRET_KEY: 'doc5eRXFpsWllGC5yKJV/Ymm5KwF+IRZo95EudOm'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-3",
			Rule: NewAWSMWSKey(),
			Src:  SampleVulnerableHSLEAKS3,
			Findings: []engine.Finding{
				{
					CodeSample: `AWS_WMS_KEY: 'amzn.mws.986478f0-9775-eabc-2af4-e499a8496828'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 20,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-4",
			Rule: NewFacebookSecretKey(),
			Src:  SampleVulnerableHSLEAKS4,
			Findings: []engine.Finding{
				{
					CodeSample: `FB_SECRET_KEY: 'cb6f53505911332d30867f44a1c1b9b5'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-5",
			Rule: NewFacebookClientID(),
			Src:  SampleVulnerableHSLEAKS5,
			Findings: []engine.Finding{
				{
					CodeSample: `FB_CLIENT_ID: '148695999071979'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-6",
			Rule: NewTwitterSecretKey(),
			Src:  SampleVulnerableHSLEAKS6,
			Findings: []engine.Finding{
				{
					CodeSample: `TWITTER_SECRET_KEY: 'ej64cqk9k8px9ae3e47ip89l7if58tqhpxi1r'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-7",
			Rule: NewTwitterClientID(),
			Src:  SampleVulnerableHSLEAKS7,
			Findings: []engine.Finding{
				{
					CodeSample: `TWITTER_CLIENT_ID: '1h6433fsvygnyre5a40'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-8",
			Rule: NewGithub(),
			Src:  SampleVulnerableHSLEAKS8,
			Findings: []engine.Finding{
				{
					CodeSample: `GITHUB_SECRET_KEY: 'edzvPbU3SYUc7pFc9le20lzIRErTOaxCABQ1'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-9",
			Rule: NewLinkedInClientID(),
			Src:  SampleVulnerableHSLEAKS9,
			Findings: []engine.Finding{
				{
					CodeSample: `LINKEDIN_CLIENT_ID: 'g309xttlaw25'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-10",
			Rule: NewLinkedInSecretKey(),
			Src:  SampleVulnerableHSLEAKS10,
			Findings: []engine.Finding{
				{
					CodeSample: `LINKEDIN_SECRET_KEY: '0d16kcnjyfzmcmjp'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-11",
			Rule: NewSlack(),
			Src:  SampleVulnerableHSLEAKS11,
			Findings: []engine.Finding{
				{
					CodeSample: `SLACK_WEBHOOK: 'https://hooks.slack.com/services/TNeqvYPeO/BncTJ74Hf/NlvFFKKAKPkd6h7FlQCz1Blu'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 22,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-12",
			Rule: NewAsymmetricPrivateKey(),
			Src:  SampleVulnerableHSLEAKS12,
			Findings: []engine.Finding{
				{
					CodeSample: `SSH_PRIVATE_KEY: '-----BEGIN PRIVATE KEY-----MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDBj08sp5++4anGcmQxJjAkBgNVBAoTHVByb2dyZXNzIFNvZnR3YXJlIENvcnBvcmF0aW9uMSAwHgYDVQQDDBcqLmF3cy10ZXN0LnByb2dyZXNzLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD...bml6YXRpb252YWxzaGEyZzIuY3JsMIGgBggrBgEFBQcBAQSBkzCBkDBNBggrBgEFBQcwAoZBaHR0cDovL3NlY3VyZS5nbG9iYWxzaWduLmNvbS9jYWNlcnQvZ3Nvcmdhz3P668YfhUbKdRF6S42Cg6zn-----END PRIVATE KEY-----'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 24,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-13",
			Rule: NewGoogleAPIKey(),
			Src:  SampleVulnerableHSLEAKS13,
			Findings: []engine.Finding{
				{
					CodeSample: `GCP_API_KEY: 'AIzaMPZHYiu1RdzE1nG2SaVyOoz244TuacQIR6m'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 20,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-14",
			Rule: NewGoogleGCPServiceAccount(),
			Src:  SampleVulnerableHSLEAKS14,
			Findings: []engine.Finding{
				{
					CodeSample: `GCP_SERVICE_ACCOUNT: '18256698220617903267772185514630273595-oy8_uzouz8tyy46y84ckrwei9_6rq_pb.apps.googleusercontent.com'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-15",
			Rule: NewHerokuAPIKey(),
			Src:  SampleVulnerableHSLEAKS15,
			Findings: []engine.Finding{
				{
					CodeSample: `HEROKU_API_KEY: '3623f8e9-2d05-c9bb-2209082d6b5c'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-16",
			Rule: NewMailChimpAPIKey(),
			Src:  SampleVulnerableHSLEAKS16,
			Findings: []engine.Finding{
				{
					CodeSample: `MAILCHIMP_API_KEY: 'f7e9c13c10d0b19c3bb003a9f635d488-us72'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-17",
			Rule: NewMailgunAPIKey(),
			Src:  SampleVulnerableHSLEAKS17,
			Findings: []engine.Finding{
				{
					CodeSample: `MAILGUN_API_KEY: 'key-xke9nbc2i5po5cjw3ngyxiz450zxpapu'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-18",
			Rule: NewPayPalBraintreeAccessToken(),
			Src:  SampleVulnerableHSLEAKS18,
			Findings: []engine.Finding{
				{
					CodeSample: `PAY_PAL_ACCESS_TOKEN: 'access_token$production$mk0sech2v7qqsol3$db651af2221c22b4ca2f0f583798135e'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 29,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-19",
			Rule: NewPicaticAPIKey(),
			Src:  SampleVulnerableHSLEAKS19,
			Findings: []engine.Finding{
				{
					CodeSample: `PICATIC_API_KEY: 'sk_live_voy1p9k7r9g9j8ezmif488nk2p8310nl'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 24,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-20",
			Rule: NewSendGridAPIKey(),
			Src:  SampleVulnerableHSLEAKS20,
			Findings: []engine.Finding{
				{
					CodeSample: `SEND_GRID_API_KEY: 'SG.44b7kq3FurdH0bSHBGjPSWhE8vJ.1evu4Un0TXFIb1_6zW4YOdjTMeE'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 26,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-21",
			Rule: NewStripeAPIKey(),
			Src:  SampleVulnerableHSLEAKS21,
			Findings: []engine.Finding{
				{
					CodeSample: `STRIPE_API_KEY: 'rk_live_8qSZpoI9t0BOGkOLVzvesc6K'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-22",
			Rule: NewSquareAccessToken(),
			Src:  SampleVulnerableHSLEAKS22,
			Findings: []engine.Finding{
				{
					CodeSample: `SQUARE_ACCESS_TOKEN: 'sq0atp-clYRBSht6oefa7w_2R56ra'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 28,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-23",
			Rule: NewSquareOAuthSecret(),
			Src:  SampleVulnerableHSLEAKS23,
			Findings: []engine.Finding{
				{
					CodeSample: `SQUARE_SECRET: 'sq0csp-LsEBYQNja]OgT3hRxjJV5cWX^XjpT12n3QkRY_vep2z'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 22,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-24",
			Rule: NewTwilioAPIKey(),
			Src:  SampleVulnerableHSLEAKS24,
			Findings: []engine.Finding{
				{
					CodeSample: `TWILIO_API_KEY: '^SK9ae6bd84ccd091eb6bfad8e2a474af95'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 6,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-25",
			Rule: NewHardCodedCredentialGeneric(),
			Src:  SampleVulnerableHSLEAKS25,
			Findings: []engine.Finding{
				{
					CodeSample: `POSTGRES_DBPASSWD: 'Ch@ng3m3'`,
					SourceLocation: engine.Location{
						Line:   7,
						Column: 15,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-26",
			Rule: NewHardCodedPassword(),
			Src:  SampleVulnerableHSLEAKS26,
			Findings: []engine.Finding{
				{
					CodeSample: `DB_PASSWORD="gorm"`,
					SourceLocation: engine.Location{
						Line:   12,
						Column: 4,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-27",
			Rule: NewPasswordExposedInHardcodedURL(),
			Src:  SampleVulnerableHSLEAKS27,
			Findings: []engine.Finding{
				{
					CodeSample: `dsn := "postgresql://gorm:gorm@127.0.0.1:5432/gorm?sslmode=disable"`,
					SourceLocation: engine.Location{
						Line:   10,
						Column: 9,
					},
				},
			},
		},
		{
			Name: "HS-LEAKS-28",
			Rule: NewWPConfig(),
			Src:  SampleVulnerableHSLEAKS28,
			Findings: []engine.Finding{
				{
					CodeSample: `define('AUTH_KEY', 'put your unique phrase here');`,
					SourceLocation: engine.Location{
						Line:   3,
						Column: 0,
					},
				},
				{
					CodeSample: `define('DB_PASSWORD', 'wen0221!');`,
					SourceLocation: engine.Location{
						Line:   4,
						Column: 0,
					},
				},
			},
		},
	}

	testutil.TestVulnerableCode(t, testcases)
}
func TestRulesSafeCode(t *testing.T) {
	testcases := []*testutil.RuleTestCase{
		{
			Name: "HS-LEAKS-1",
			Rule: NewAWSManagerID(),
			Src:  SampleSafeHSLEAKS1,
		},
		{
			Name: "HS-LEAKS-2",
			Rule: NewAWSSecretKey(),
			Src:  SampleSafeHSLEAKS2,
		},
		{
			Name: "HS-LEAKS-3",
			Rule: NewAWSMWSKey(),
			Src:  SampleSafeHSLEAKS3,
		},
		{
			Name: "HS-LEAKS-4",
			Rule: NewFacebookSecretKey(),
			Src:  SampleSafeHSLEAKS4,
		},
		{
			Name: "HS-LEAKS-5",
			Rule: NewFacebookClientID(),
			Src:  SampleSafeHSLEAKS5,
		},
		{
			Name: "HS-LEAKS-6",
			Rule: NewTwitterSecretKey(),
			Src:  SampleSafeHSLEAKS6,
		},
		{
			Name: "HS-LEAKS-7",
			Rule: NewTwitterClientID(),
			Src:  SampleSafeHSLEAKS7,
		},
		{
			Name: "HS-LEAKS-8",
			Rule: NewGithub(),
			Src:  SampleSafeHSLEAKS8,
		},
		{
			Name: "HS-LEAKS-9",
			Rule: NewLinkedInClientID(),
			Src:  SampleSafeHSLEAKS9,
		},
		{
			Name: "HS-LEAKS-10",
			Rule: NewLinkedInSecretKey(),
			Src:  SampleSafeHSLEAKS10,
		},
		{
			Name: "HS-LEAKS-11",
			Rule: NewSlack(),
			Src:  SampleSafeHSLEAKS11,
		},
		{
			Name: "HS-LEAKS-12",
			Rule: NewAsymmetricPrivateKey(),
			Src:  SampleSafeHSLEAKS12,
		},
		{
			Name: "HS-LEAKS-13",
			Rule: NewGoogleAPIKey(),
			Src:  SampleSafeHSLEAKS13,
		},
		{
			Name: "HS-LEAKS-14",
			Rule: NewGoogleGCPServiceAccount(),
			Src:  SampleSafeHSLEAKS14,
		},
		{
			Name: "HS-LEAKS-15",
			Rule: NewHerokuAPIKey(),
			Src:  SampleSafeHSLEAKS15,
		},
		{
			Name: "HS-LEAKS-16",
			Rule: NewMailChimpAPIKey(),
			Src:  SampleSafeHSLEAKS16,
		},
		{
			Name: "HS-LEAKS-17",
			Rule: NewMailgunAPIKey(),
			Src:  SampleSafeHSLEAKS17,
		},
		{
			Name: "HS-LEAKS-18",
			Rule: NewPayPalBraintreeAccessToken(),
			Src:  SampleSafeHSLEAKS18,
		},
		{
			Name: "HS-LEAKS-19",
			Rule: NewPicaticAPIKey(),
			Src:  SampleSafeHSLEAKS19,
		},
		{
			Name: "HS-LEAKS-20",
			Rule: NewSendGridAPIKey(),
			Src:  SampleSafeHSLEAKS20,
		},
		{
			Name: "HS-LEAKS-21",
			Rule: NewStripeAPIKey(),
			Src:  SampleSafeHSLEAKS21,
		},
		{
			Name: "HS-LEAKS-22",
			Rule: NewSquareAccessToken(),
			Src:  SampleSafeHSLEAKS22,
		},
		{
			Name: "HS-LEAKS-23",
			Rule: NewSquareOAuthSecret(),
			Src:  SampleSafeHSLEAKS23,
		},
		{
			Name: "HS-LEAKS-24",
			Rule: NewTwilioAPIKey(),
			Src:  SampleSafeHSLEAKS24,
		},
		{
			Name: "HS-LEAKS-25",
			Rule: NewHardCodedCredentialGeneric(),
			Src:  SampleSafeHSLEAKS25,
		},
		{
			Name: "HS-LEAKS-26",
			Rule: NewHardCodedPassword(),
			Src:  SampleSafeHSLEAKS26,
		},
		{
			Name: "HS-LEAKS-27",
			Rule: NewPasswordExposedInHardcodedURL(),
			Src:  SampleSafeHSLEAKS27,
		},
		{
			Name: "HS-LEAKS-28",
			Rule: NewWPConfig(),
			Src:  SampleSafeHSLEAKS28,
		},
	}
	testutil.TestSafeCode(t, testcases)
}
