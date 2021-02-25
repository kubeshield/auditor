/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"kmodules.xyz/auditor/pkg/cmds"
	"kmodules.xyz/client-go/logs"

	"go.bytebuilders.dev/license-verifier/info"
	_ "go.bytebuilders.dev/license-verifier/info"
	_ "k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	info.LicenseCA = `-----BEGIN CERTIFICATE-----
MIIC6DCCAdCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAlMRYwFAYDVQQKEw1BcHBz
Q29kZSBJbmMuMQswCQYDVQQDEwJjYTAeFw0yMDA4MDkxMTE4MzJaFw0zMDA4MDcx
MTE4MzJaMCUxFjAUBgNVBAoTDUFwcHNDb2RlIEluYy4xCzAJBgNVBAMTAmNhMIIB
IjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtO3NKvg6Jk11RkYqDfkZwajY
/w8bHiq/DV5KjQ7h45BHzLxrd4XupZweRQR1MqMUVxH3sXagO6q7vGMWzuhBmC9e
7B67ZjRGt3z3B49Q6VFIop0NB2DoYWk6FsAK4Fp3jtIgXCMcFApdmPZZ20H3F+mq
KAaS1I6X5VXEr5II9qvncUO2a7O9Tb4H+xZsr1xdv0CuC3FevmfQLbFt5nQJTRNq
ukcugxvImsoF1WZ+d8cz65krnvMlC5uUyGDCqpyIh4Iy1sMssk/7MOVzgqGZ8ISa
f5Jv+3IzZL3rQQ7TGZMBBeMBvIES6bg5FmDvg/6Rgo5K9KifdFkKLFiMUSKvOQID
AQABoyMwITAOBgNVHQ8BAf8EBAMCAqQwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG
9w0BAQsFAAOCAQEAAYumKKA1aPF3MuLRvSnUWijvw5+Rdptc+AdMgNvONrpukZ26
BAGRhA5DkGupBCjkyMCIAcTUX+hgW8QKpucun4VSoMkW2x69Z6xfKxDhGlRk3PD1
jeHcDmdnzC864tQ8rdINd+D7RksdBP9aCWTkXlcBlkYimCkanmFbaz7YmmFvdnZs
LeqjtZmJeARBz7p59AT4Pn2NpdLYKjHP7AEFmVpoF7Z4y1cl0AwsINuVNM5++7El
YMJ9tfGRB4Sgj1BZKcvwmqY3RBOtgpY5P27w46WqxYhRrmP867GEWzeCSzm1jRAI
hqDntdQyIGPXtqiMjPjKUxUMCSsAGL3ZqrMe9Q==
-----END CERTIFICATE-----`
	info.ProductName = "kubedb-community"

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
