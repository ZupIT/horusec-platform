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

package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
)

func TestRemoveHorusecIDFromDetails(t *testing.T) {
	migration := &Migration{}

	t.Run("should success remove csharp horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-CSHARP-1: Command Injection
				If a malicious user controls either the FileName or Arguments, he might be able to execute unwanted 
				commands or add unwanted argument. This behavior would not be possible if input parameter are 
				validate against a white-list of characters. 
				For more information access: (https://security-code-scan.github.io/#SCS0001).
		`

		expected := `
				Command Injection
				If a malicious user controls either the FileName or Arguments, he might be able to execute unwanted 
				commands or add unwanted argument. This behavior would not be possible if input parameter are 
				validate against a white-list of characters. 
				For more information access: (https://security-code-scan.github.io/#SCS0001).
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove dart horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-DART-17: Using shell interpreter when executing OS commands
				Arbitrary OS command injection vulnerabilities are more likely when a shell is spawned rather 
				than a new process, indeed shell meta-chars can be used (when parameters are user-controlled 
				for instance) to inject OS commands. For more information checkout the CWE-78 
				(https://cwe.mitre.org/data/definitions/78.html) advisory.
		`

		expected := `
				Using shell interpreter when executing OS commands
				Arbitrary OS command injection vulnerabilities are more likely when a shell is spawned rather 
				than a new process, indeed shell meta-chars can be used (when parameters are user-controlled 
				for instance) to inject OS commands. For more information checkout the CWE-78 
				(https://cwe.mitre.org/data/definitions/78.html) advisory.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove java horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-JAVA-149: Databases should be password-protected
				Databases should always be password protected. The use of a database connection with an empty 
				password is a clear indication of a database that is not protected. For more information checkout 
				the CWE-521 (https://cwe.mitre.org/data/definitions/521.html) advisory.
		`

		expected := `
				Databases should be password-protected
				Databases should always be password protected. The use of a database connection with an empty 
				password is a clear indication of a database that is not protected. For more information checkout 
				the CWE-521 (https://cwe.mitre.org/data/definitions/521.html) advisory.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove jvm horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-JVM-36: Super User Privileges
				This App may request root (Super User) privileges. For more information 
				checkout the CWE-250 (https://cwe.mitre.org/data/definitions/250.html) advisory.
		`

		expected := `
				Super User Privileges
				This App may request root (Super User) privileges. For more information 
				checkout the CWE-250 (https://cwe.mitre.org/data/definitions/250.html) advisory.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove kotlin horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-KOTLIN-123456789: Example Title
				This is a example of a vulnerability. For more information checkout the CWE-000 
				(https://cwe.mitre.org/data/definitions/000.html) advisory.
		`

		expected := `
				Example Title
				This is a example of a vulnerability. For more information checkout the CWE-000 
				(https://cwe.mitre.org/data/definitions/000.html) advisory.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove kubernetes horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-KUBERNETES-7: Host IPC
				Sharing the host's IPC namespace allows container processes to communicate with processes on the host.
		`

		expected := `
				Host IPC
				Sharing the host's IPC namespace allows container processes to communicate with processes on the host.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove leaks horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-LEAKS-28: Wordpress configuration file disclosure
				Wordpress configuration file exposed, this can lead to the leak of admin passwords, database 
				credentials and a lot of sensitive data about the system. 
				Check CWE-200 (https://cwe.mitre.org/data/definitions/200.html) for more details.
		`

		expected := `
				Wordpress configuration file disclosure
				Wordpress configuration file exposed, this can lead to the leak of admin passwords, database 
				credentials and a lot of sensitive data about the system. 
				Check CWE-200 (https://cwe.mitre.org/data/definitions/200.html) for more details.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove nginx horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-NGINX-3: Missing Content-Security-Policy header
				A Content Security Policy (also named CSP) requires careful tuning and precise definition of 
				the policy. If enabled, CSP has significant impact on the way browsers render pages 
				(e.g., inline JavaScript is disabled by default and must be explicitly allowed in the policy). 
				CSP prevents a wide range of attacks, including cross-site scripting and other cross-site injections. 
				For more information checkout https://owasp.org/www-project-secure-headers/#content-security-policy.
		`

		expected := `
				Missing Content-Security-Policy header
				A Content Security Policy (also named CSP) requires careful tuning and precise definition of 
				the policy. If enabled, CSP has significant impact on the way browsers render pages 
				(e.g., inline JavaScript is disabled by default and must be explicitly allowed in the policy). 
				CSP prevents a wide range of attacks, including cross-site scripting and other cross-site injections. 
				For more information checkout https://owasp.org/www-project-secure-headers/#content-security-policy.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove javascript horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-JAVASCRIPT-53: SQL Injection
				SQL queries often need to use a hardcoded SQL string with a dynamic parameter coming from a user 
				request. Formatting a string to add those parameters to the request is a bad practice as it can result 
				in an SQL injection. The safe way to add parameters to a SQL query is to use SQL binding mechanisms. 
				For more information checkout the CWE-564 (https://cwe.mitre.org/data/definitions/564.html) 
				and OWASP A1:2017 (https://owasp.org/www-project-top-ten/2017/A1_2017-Injection.html) advisory.
		`

		expected := `
				SQL Injection
				SQL queries often need to use a hardcoded SQL string with a dynamic parameter coming from a user 
				request. Formatting a string to add those parameters to the request is a bad practice as it can result 
				in an SQL injection. The safe way to add parameters to a SQL query is to use SQL binding mechanisms. 
				For more information checkout the CWE-564 (https://cwe.mitre.org/data/definitions/564.html) 
				and OWASP A1:2017 (https://owasp.org/www-project-top-ten/2017/A1_2017-Injection.html) advisory.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})

	t.Run("should success remove swift horusec vulnerability id", func(t *testing.T) {
		input := `
				HS-SWIFT-11: Weak MD5 hash using
				MD5 is a weak hash, which can generate repeated hashes. For more information checkout the 
				CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.
		`

		expected := `
				Weak MD5 hash using
				MD5 is a weak hash, which can generate repeated hashes. For more information checkout the 
				CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.
		`

		assert.Equal(t, expected, migration.removeHorusecIDFromDetails(input))
	})
}

func TestGenerateExpectedHash(t *testing.T) {
	migration := &Migration{}

	t.Run("should success generate hash without horusec vulnerability id", func(t *testing.T) {
		const expected = "b402258a378914b3198b01b78fa500b2dbf984288cc7978795f30ed13809d8cf"

		details := "HS-JAVA-114: Insecure Random Number Generator\nThe App uses an insecure Random Number Generator. " +
			"For more information checkout the CWE-330 (https://cwe.mitre.org/data/definitions/330.html) advisory."

		vuln := &Vulnerability{
			Vulnerability: &vulnerability.Vulnerability{
				Line:        "18",
				Code:        "import java.util.Random;",
				Details:     details,
				File:        "src/main/java/com/mycompany/app/App.java",
				CommitEmail: "-",
			},
		}

		assert.Equal(t, expected, migration.generateExpectedHash(vuln))
	})

	t.Run("should success generate hash without horusec vulnerability id", func(t *testing.T) {
		const expected = "90798dc88885e0ca8b119ca7b6b17f5f2a3a0b4d2dbe92131862f3cecd8248f8"

		details := "HS-NGINX-2: Missing X-Content-Type-Options header\nSetting this header will prevent the browser " +
			"from interpreting files as a different MIME type to what is specified in the Content-Type HTTP header " +
			"(e.g. treating text/plain as text/css). For more information checkout " +
			"https://owasp.org/www-project-secure-headers/#x-content-type-options"

		vuln := &Vulnerability{
			Vulnerability: &vulnerability.Vulnerability{
				Line:        "0",
				Code:        "",
				Details:     details,
				File:        "server.nginx",
				CommitEmail: "-",
			},
		}

		assert.Equal(t, expected, migration.generateExpectedHash(vuln))
	})

	t.Run("should success generate hash without horusec vulnerability id", func(t *testing.T) {
		const expected = "6e7e72a6a6812672083de2a8046fb5e62c9c267511595934dcb89109da52acc7"

		details := "HS-LEAKS-12: Asymmetric Private Key\nFound SSH and/or x.509 Cerficates among the files of your " +
			"project, make sure you want this kind of information inside your Git repo, since it can be missused " +
			"by someone with access to any kind of copy.  For more information checkout the CWE-312 " +
			"(https://cwe.mitre.org/data/definitions/312.html) advisory."

		vuln := &Vulnerability{
			Vulnerability: &vulnerability.Vulnerability{
				Line:        "1",
				Code:        "-----BEGIN CERTIFICATE-----",
				Details:     details,
				File:        "example1/deployments/certificates/client-cert.txt",
				CommitEmail: "-",
			},
		}

		assert.Equal(t, expected, migration.generateExpectedHash(vuln))
	})

	t.Run("should success generate hash without horusec vulnerability id", func(t *testing.T) {
		const expected = "0a26b4688bf072db883d5917cd7b1eb8c0b321ca90a55cf0e42afe3926fdaedc"

		details := "HS-CSHARP-57: Weak hashing function md5 or sha1\nMD5 or SHA1 have known collision weaknesses " +
			"and are no longer considered strong hashing algorithms. For more information checkout the CWE-326 " +
			"(https://cwe.mitre.org/data/definitions/326.html) advisory."

		vuln := &Vulnerability{
			Vulnerability: &vulnerability.Vulnerability{
				Line:        "29",
				Code:        "var hashProvider = new SHA1CryptoServiceProvider();",
				Details:     details,
				File:        "Vulnerabilities.cs",
				CommitEmail: "-",
			},
		}

		assert.Equal(t, expected, migration.generateExpectedHash(vuln))
	})

	t.Run("should success generate hash without horusec vulnerability id", func(t *testing.T) {
		const expected = "88f070d97ae258fc4470b1f3783414f1997a1ecd5990451de28e83b7a5a30243"

		details := "HS-KUBERNETES-4: Capability System Admin\nCAP_SYS_ADMIN is the most privileged capability " +
			"and should always be avoided."

		vuln := &Vulnerability{
			Vulnerability: &vulnerability.Vulnerability{
				Line:        "0",
				Code:        "",
				Details:     details,
				File:        "example1/example.yaml",
				CommitEmail: "-",
			},
		}

		assert.Equal(t, expected, migration.generateExpectedHash(vuln))
	})
}
