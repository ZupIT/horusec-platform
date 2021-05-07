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

import { useCallback, useState } from 'react';

type Operation<T> = (obj: T) => void;

export function useListState<T>(
  initialState: T[],
  identificationFn: (a: T, b: T) => boolean
): [T[], Operation<T | T[]>, Operation<T>, Operation<T>] {
  const [state, setState] = useState(initialState);

  const add: Operation<T | T[]> = useCallback((data) => {
    if (Array.isArray(data)) {
      setState(data);
      return;
    }

    setState((state) => [...state, data]);
  }, []);

  const update: Operation<T> = useCallback(
    (data) => {
      setState((state) => [
        ...state.map((item) => {
          if (identificationFn(item, data)) {
            return data;
          }

          return item;
        }),
      ]);
    },
    [identificationFn]
  );

  const remove: Operation<T> = useCallback(
    (data) => {
      setState((state) =>
        state.filter((item) => !identificationFn(item, data))
      );
    },
    [identificationFn]
  );

  return [state, add, update, remove];
}
