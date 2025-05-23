/*
Copyright 2016 The TensorFlow Authors. All Rights Reserved.

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

package tensorflow

// #cgo windows CFLAGS: -I${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/windows/tensorflow/include
// #cgo windows LDFLAGS: -L${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/windows/tensorflow/lib -ltensorflow
// #cgo darwin CFLAGS: -I${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/osx/tensorflow/include
// #cgo darwin LDFLAGS: -L${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/osx/tensorflow/lib -Wl,-undefined,dynamic_lookup
// #cgo linux CFLAGS: -I${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/linux/tensorflow/include
// #cgo linux LDFLAGS: -L${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/linux/tensorflow/lib -Wl,-rpath,${SRCDIR}/../../../../../../../src/github.com/kiteco/kiteco/linux/tensorflow/lib,-z,undefs
// #cgo CFLAGS: -I${SRCDIR}/../../
import "C"
