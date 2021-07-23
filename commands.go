/*

Copyright 2021 Andrey Devyatkin.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package main

import (
	"strings"
)

func Push(destinationRepo string, tag string) (ImageId, error) {
	client, err := GetECRClient()
	if err != nil {
		return ImageId{}, err
	}

	authConfig, err := GetDockerAuthConfig(client)
	if err != nil {
		return ImageId{}, err
	}

	cli, err := getDockerClient()
	if err != nil {
		return ImageId{}, err
	}

	return imagePush(cli, authConfig, destinationRepo, tag)
}

func IsScanFailed(destinationRepo string, imageId ImageId, findingScoreThreshold string, cveIgnoreList string) (bool, error) {
	client, err := GetECRClient()
	if err != nil {
		return true, err
	}

	return AreThereCVEsToReport(client, imageId, destinationRepo, strings.Fields(findingScoreThreshold), strings.Fields(cveIgnoreList))
}

func Tag(repo string, imageId string, destinationTag string) error {
	cli, err := getDockerClient()
	if err != nil {
		return err
	}

	return imageTag(cli, repo, imageId, destinationTag)
}