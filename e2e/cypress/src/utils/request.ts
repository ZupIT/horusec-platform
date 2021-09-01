/**
 * Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

import axios, { AxiosInstance, AxiosResponse } from "axios";

export interface IServices {
    Auth: string;
    Api: string;
    Analytic: string;
    Account: string;
    Manager: string;
    Keycloak: string;
}

export class Requests {
    public baseURL = "http://127.0.0.1";
    public services: IServices = {
        Auth: ":8006",
        Api: ":8000",
        Analytic: ":8005",
        Account: ":8003",
        Manager: ":8043",
        Keycloak: ":8080",
    };
    private _axiosInstance: AxiosInstance;

    constructor() {
        this._axiosInstance = this._axios();
    }

    public setHeadersAllRequests(headers?: any): Requests {
        this._axiosInstance = this._axios(headers);
        return this;
    }

    public get(url: string, headers?: any): Promise<AxiosResponse> {
        return this._axiosInstance.get(url, headers);
    }

    public post(url: string, body?: any, headers?: any): Promise<AxiosResponse> {
        return this._axiosInstance.post(url, body, {"content-type": "*", ...headers});
    }

    public put(url: string, body?: any, headers?: any): Promise<AxiosResponse> {
        return this._axiosInstance.put(url, body, {"content-type": "*", ...headers});
    }

    public patch(url: string, body?: any, headers?: any): Promise<AxiosResponse> {
        return this._axiosInstance.patch(url, body, {"content-type": "*", ...headers});
    }

    public delete(url: string, headers?: any): Promise<AxiosResponse> {
        return this._axiosInstance.delete(url, headers);
    }

    private _axios(headers?: any): AxiosInstance {
        return axios.create({
            timeout: 15000,
            headers,
        });
    }
}
