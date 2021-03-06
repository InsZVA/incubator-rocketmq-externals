/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constant

const (
	//PERM_PRIORITY PERM_PRIORITY
	PERM_PRIORITY = 0x1 << 3

	//PERM_READ this queue can read
	PERM_READ = 0x1 << 2
	//PERM_WRITE this queue can write
	PERM_WRITE = 0x1 << 1

	//PERM_INHERIT PERM_INHERIT
	PERM_INHERIT = 0x1 << 0
)

//WriteAble this queue can write
func WriteAble(perm int32) (ret bool) {
	ret = (perm & PERM_WRITE) == PERM_WRITE
	return
}

//ReadAble this queue can read
func ReadAble(perm int32) (ret bool) {
	ret = (perm & PERM_READ) == PERM_READ
	return
}
