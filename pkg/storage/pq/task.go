/*
 * Minimalist Object Storage, (C) 2015 Minio, Inc.
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

package pq

// Task container for any generic tasks
type Task struct {
	job      func() error
	priority int
}

// GetPriority get current task priority
func (t Task) GetPriority() int {
	return t.priority
}

// UpdatePriority update current task priority
func (t Task) UpdatePriority(p int) {
	t.priority = p
}

// Execute execute current task
func (t Task) Execute() error {
	return t.job()
}
