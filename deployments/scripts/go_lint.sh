#!/bin/sh
# Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GOLANG_CI_LINT=$(pwd)/bin/golangci-lint
GOLANG_CI_LINT_FILE=$(pwd)/.golangci.yml

installGolangCILint() {
    if [ ! -f "$GOLANG_CI_LINT" ]
    then
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
    fi
}

runLint() {
    for PROJECT in api core analytic messages webhook auth; do
        cd $PROJECT || echo "not found"
        echo "==================== Running on project ($PROJECT) ===================="
        if ! ${GOLANG_CI_LINT} run -v --timeout=5m -c "$GOLANG_CI_LINT_FILE" ./...
        then
            break
        else
            echo "[FINISH] REPOSITORY ALREADY IS OK!"
        fi
        cd ..
    done
}

installGolangCILint

runLint
