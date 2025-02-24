// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package swift

import (
	"regexp"

	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	engine "github.com/ZupIT/horusec-engine"
	"github.com/ZupIT/horusec-engine/text"
)

func NewSQLiteDatabase() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-1",
			Name:        "SQLite Database",
			Description: "App uses SQLite Database. Sensitive Information should be encrypted.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`sqlite3_exec`),
			regexp.MustCompile(`sqlite3_finalize`),
		},
	}
}

func NewCoreDataDatabase() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-2",
			Name:        "CoreData Database",
			Description: "App uses CoreData Database. Sensitive Information should be encrypted.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`NSManagedObjectContext`),
			regexp.MustCompile(`\.save\(\)`),
		},
	}
}

func NewDTLS12NotUsed() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-3",
			Name:        "DTLS 1.2 not used",
			Description: "DTLS 1.2 should be used. Detected old version - DTLS 1.0.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\.TLSMinimumSupportedProtocolVersion`),
			regexp.MustCompile(`tls_protocol_version_t\.DTLSv10`),
		},
	}
}

func NewTLS13NotUsed() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-4",
			Name:        "TLS 1.3 not used",
			Description: "Older versions of SSL/TLS protocol like \"SSLv3\" have been proven to be insecure. This rule raises an issue when an SSL/TLS context is created with an insecure protocol version (ie: a protocol different from \"TLSv1.2\", \"TLSv1.3\", \"DTLSv1.2\" or \"DTLSv1.3\"). For more information checkout the CWE-326 (https://cwe.mitre.org/data/definitions/326.html) and CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\.TLSMinimumSupportedProtocolVersion`),
			regexp.MustCompile(`tls_protocol_version_t\.TLSv12`),
		},
	}
}

func NewReverseEngineering() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-5",
			Name:        "Reverse engineering",
			Description: "This App may have Reverse engineering detection capabilities.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`"FridaGadget"`),
			regexp.MustCompile(`"cynject"`),
			regexp.MustCompile(`"libcycript"`),
			regexp.MustCompile(`"/usr/sbin/frida-server"`),
		},
	}
}

func NewWeakMD5CryptoCipher() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-6",
			Name:        "Weak MD5 hash using",
			Description: "The MD5 hash algorithm that was used is considered weak. It can also cause hash collisions. It is always recommended to use some CHF (Cryptographic Hash Function), which is mathematically strong and not reversible. SHA512 would be the most recommended hash for storing the password and it is also important to adopt some type of Salt, so that the Hash is more secure. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`import CryptoSwift`),
			regexp.MustCompile(`\.md5()`),
		},
	}
}

func NewWeakCommonDesCryptoCipher() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-7",
			Name:        "Weak DES hash using",
			Description: "DES is considered strong ciphers for modern applications. Currently, NIST recommends the usage of AES block ciphers instead of DES. For more information checkout the CWE-326 (https://cwe.mitre.org/data/definitions/326.html) advisory",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`import CommonCrypto`),
			regexp.MustCompile(`CCAlgorithm\(kCCAlgorithmDES\)`),
		},
	}
}

func NewWeakIDZDesCryptoCipher() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-8",
			Name:        "Weak DES hash using",
			Description: "DES is considered strong ciphers for modern applications. Currently, NIST recommends the usage of AES block ciphers instead of DES. For more information checkout the CWE-326 (https://cwe.mitre.org/data/definitions/326.html) advisory",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`import IDZSwiftCommonCrypto`),
			regexp.MustCompile(`\.des`),
		},
	}
}

func NewWeakBlowfishCryptoCipher() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-9",
			Name:        "Weak Cipher Mode",
			Description: "Cipher algorithms should be robust",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`import CryptoSwift`),
			regexp.MustCompile(`Blowfish\(.*\)`),
		},
	}
}

func NewMD6Collision() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-10",
			Name:        "Weak MD6 hash using",
			Description: "MD6 is a weak hash, which can generate repeated hashes. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.OrMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`(?i)MD6\(`),
			regexp.MustCompile(`CC_MD6\(`),
		},
	}
}

func NewMD5Collision() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-11",
			Name:        "Weak MD5 hash using",
			Description: "MD5 is a weak hash, which can generate repeated hashes. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.OrMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`(?i)MD5\(`),
			regexp.MustCompile(`CC_MD5\(`),
		},
	}
}

func NewSha1Collision() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-12",
			Name:        "Weak SHA1 hash using",
			Description: "SHA1 is a weak hash, which can generate repeated hashes. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.OrMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`(?i)SHA1\(`),
			regexp.MustCompile(`CC_SHA1\(`),
		},
	}
}

func NewJailbreakDetect() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-13",
			Name:        "Jailbreak detection",
			Description: "This App may have Jailbreak detection capabilities.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.OrMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`/Applications/Cydia\.app`),
			regexp.MustCompile(`/Library/MobileSubstrate/MobileSubstrate\.dylib`),
			regexp.MustCompile(`/usr/sbin/sshd`),
			regexp.MustCompile(`/etc/apt`),
			regexp.MustCompile(`cydia://`),
			regexp.MustCompile(`/var/lib/cydia`),
			regexp.MustCompile(`/Applications/FakeCarrier\.app`),
			regexp.MustCompile(`/Applications/Icy\.app`),
			regexp.MustCompile(`/Applications/IntelliScreen\.app`),
			regexp.MustCompile(`/Applications/SBSettings\.app`),
			regexp.MustCompile(`/Library/MobileSubstrate/DynamicLibraries/LiveClock\.plist`),
			regexp.MustCompile(`/System/Library/LaunchDaemons/com\.ikey\.bbot\.plist`),
			regexp.MustCompile(`/System/Library/LaunchDaemons/com\.saurik\.Cydia\.Startup\.plist`),
			regexp.MustCompile(`/etc/ssh/sshd_config`),
			regexp.MustCompile(`/private/var/tmp/cydia\.log`),
			regexp.MustCompile(`/usr/libexec/ssh-keysign`),
			regexp.MustCompile(`/Applications/MxTube\.app`),
			regexp.MustCompile(`/Applications/RockApp\.app`),
			regexp.MustCompile(`/Applications/WinterBoard\.app`),
			regexp.MustCompile(`/Applications/blackra1n\.app`),
			regexp.MustCompile(`/Library/MobileSubstrate/DynamicLibraries/Veency\.plist`),
			regexp.MustCompile(`/private/var/lib/apt`),
			regexp.MustCompile(`/private/var/lib/cydia`),
			regexp.MustCompile(`/private/var/mobile/Library/SBSettings/Themes`),
			regexp.MustCompile(`/private/var/stash`),
			regexp.MustCompile(`/usr/bin/sshd`),
			regexp.MustCompile(`/usr/libexec/sftp-server`),
			regexp.MustCompile(`/var/cache/apt`),
			regexp.MustCompile(`/var/lib/apt`),
			regexp.MustCompile(`/usr/sbin/frida-server`),
			regexp.MustCompile(`/usr/bin/cycript`),
			regexp.MustCompile(`/usr/local/bin/cycript`),
			regexp.MustCompile(`/usr/lib/libcycript.dylib`),
			regexp.MustCompile(`frida-server`),
			regexp.MustCompile(`/etc/apt/sources\.list\.d/electra\.list`),
			regexp.MustCompile(`/etc/apt/sources\.list\.d/sileo\.sources`),
			regexp.MustCompile(`/.bootstrapped_electra`),
			regexp.MustCompile(`/usr/lib/libjailbreak\.dylib`),
			regexp.MustCompile(`/jb/lzma`),
			regexp.MustCompile(`/\.cydia_no_stash`),
			regexp.MustCompile(`/\.installed_unc0ver`),
			regexp.MustCompile(`/jb/offsets\.plist`),
			regexp.MustCompile(`/usr/share/jailbreak/injectme\.plist`),
			regexp.MustCompile(`/Library/MobileSubstrate/MobileSubstrate\.dylib`),
			regexp.MustCompile(`/usr/libexec/cydia/firmware\.sh`),
			regexp.MustCompile(`/private/var/cache/apt/`),
			regexp.MustCompile(`/Library/MobileSubstrate/CydiaSubstrate\.dylib`),
		},
	}
}

func NewLoadHTMLString() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-14",
			Name:        "Javascript injection",
			Description: "User input not sanitized in \"loadHTMLString\" can result in an injection of JavaScript in the context of your application, allowing access to private data. For more information checkout the CWE-95 (https://cwe.mitre.org/data/definitions/95.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.OrMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`loadHTMLString`),
			regexp.MustCompile(`webView`),
		},
	}
}

func NewWeakDesCryptoCipher() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-15",
			Name:        "Weak Cipher Mode",
			Description: "DES is considered strong ciphers for modern applications. Currently, NIST recommends the usage of AES block ciphers instead of DES. For more information checkout the CWE-326 (https://cwe.mitre.org/data/definitions/326.html) advisory",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.OrMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`Cryptor\((.*algorithm: \.des)`),
			regexp.MustCompile(`\.CryptAlgorithm((\s+=)|=)+((\s)|)+\"3des"`),
		},
	}
}

func NewRealmDatabase() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-16",
			Name:        "Realm Database",
			Description: "App uses Realm Database. Sensitive Information should be encrypted.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`realm\.write`),
		},
	}
}

func NewTLSMinimum() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-17",
			Name:        "Deperected tls property",
			Description: "Use of deprecated property tlsMinimumSupportedProtocol. To avoid potential security risks, use tlsMinimumSupportedProtocolVersion",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\.tlsMinimumSupportedProtocol`),
		},
	}
}

func NewUIPasteboard() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-18",
			Name:        "UIPasteboard",
			Description: "This application uses UIPasteboard, improper use of this class can lead to security issues.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`UIPasteboard`),
		},
	}
}

func NewFileProtection() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-19",
			Name:        "File protection",
			Description: "The file has no special protections associated with it.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\.noFileProtection`),
		},
	}
}

func NewWebViewSafari() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-20",
			Name:        "WebView Safari",
			Description: "It is recommended to use WKWebView instead of SFSafariViewController or UIWebView to prevent navigating to arbitrary URLs.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`UIWebView|SFSafariViewController`),
		},
	}
}

func NewKeyboardCache() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-21",
			Name:        "Keyboard cache",
			Description: "Keyboard cache should be disabled for all sensitive data inputs.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`.autocorrectionType = .no`),
		},
	}
}

func NewMD4Collision() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-22",
			Name:        "Weak MD4 hash using",
			Description: "MD4 is a weak hash, which can generate repeated hashes. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`CC_MD4\(`),
		},
	}
}

func NewMD2Collision() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "HS-SWIFT-23",
			Name:        "Weak MD2 hash using",
			Description: "MD2 is a weak hash, which can generate repeated hashes. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    severities.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.Regular,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`CC_MD2\(`),
		},
	}
}
