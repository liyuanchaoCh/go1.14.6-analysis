// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !arm
// +build !arm64
// +build !mips64
// +build !mips64le
// +build !mips
// +build !mipsle
// +build !wasm

package runtime

// careful: cputicks is not guaranteed to be monotonic! In particular, we have
// noticed drift between cpus on certain os/arch combinations. See issue 8976.
//
// //注意：cputicks 不能保证是单调的！ 特别是，我们注意到某些 os/arch 组合上cpus之间的漂移。 请参阅问题8976
func cputicks() int64
