/**
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

interface DashboardDataTypes {
  corrected: number;
  falsePositive: number;
  riskAccepted: number;
  vulnerability: number;
}

interface DashboardDataItem {
  count: number;
  types: DashboardDataTypes;
}

interface DashboardCriticality {
  critical: DashboardDataItem;
  high: DashboardDataItem;
  info: DashboardDataItem;
  low: DashboardDataItem;
  medium: DashboardDataItem;
  unknown: DashboardDataItem;
}

export interface VulnerabilitiesByAuthor extends DashboardCriticality {
  author: string;
}

export interface VulnerabilitiesByLanguageData extends DashboardCriticality {
  language: string;
}

export interface VulnerabilitiesByRepository extends DashboardCriticality {
  repositoryName: string;
}

export type VulnerabilityBySeverity = DashboardCriticality;

export interface VulnerabilityByTime extends DashboardCriticality {
  time: string;
}

export interface DashboardData {
  totalAuthors: number;
  totalRepositories: number;
  vulnerabilitiesByAuthor: VulnerabilitiesByAuthor[];
  vulnerabilitiesByLanguage: VulnerabilitiesByLanguageData[];
  vulnerabilitiesByRepository: VulnerabilitiesByRepository[];
  vulnerabilityBySeverity: VulnerabilityBySeverity;
  vulnerabilityByTime: VulnerabilityByTime[];
}
