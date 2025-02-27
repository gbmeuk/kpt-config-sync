// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parse

import (
	"sort"

	"k8s.io/klog/v2"
	"kpt.dev/configsync/pkg/core"
	"kpt.dev/configsync/pkg/importer/analyzer/ast"
	"kpt.dev/configsync/pkg/metadata"
	"kpt.dev/configsync/pkg/status"
)

// cacheForCommit tracks the progress made by the reconciler for a source commit (a source commit or an oci image digest).
//
// The reconciler resets the whole cache when a new commit is detected.
//
// The reconciler resets the whole cache except for the cached sourceState when:
//   - a force-resync happens, or
//   - one of the watchers noticed a management conflict.
type cacheForCommit struct {
	// source tracks the state of the source repo.
	// This field is only set after the reconciler successfully reads all the source files.
	source sourceState

	// hasParserResult indicates whether the cache includes the parser result.
	hasParserResult bool

	// objsSkipped contains the objects which will not be sent to the applier to apply.
	// For example, the objects whose scope is unknown will not be sent to the applier since
	// the kpt applier cannot handle unknown-scoped objects.
	objsSkipped []ast.FileObject

	// objsToApply contains the objects which will be sent to the applier to apply.
	objsToApply []ast.FileObject

	// parserErrs includes the parser errors.
	parserErrs status.MultiError

	// declaredResourcesUpdated indicates whether the resource declaration set
	// has been updated.
	declaredResourcesUpdated bool

	// applied indicates whether the applier has successfully applied the
	// declared resources.
	applied bool

	// watchesUpdated indicates whether the remediator watches have been updated
	// for the latest declared resources.
	watchesUpdated bool

	// needToRetry indicates whether a retry is needed.
	needToRetry bool

	// errs tracks all the errors encounted during the reconciliation.
	errs status.MultiError
}

func (c *cacheForCommit) setParserResult(objs []ast.FileObject, parserErrs status.MultiError) {
	knownScopeObjs, unknownScopeObjs := splitObjects(objs)
	c.objsSkipped = unknownScopeObjs
	c.objsToApply = knownScopeObjs
	c.parserErrs = parserErrs
	c.hasParserResult = true
}

func (c *cacheForCommit) parserResultUpToDate() bool {
	// If len(c.objsSkipped) > 0, it mean that some objects were skipped to be sent to
	// the kpt applier. For example, the objects whose scope is unknown will not be sent
	// to the applier since the kpt applier cannot handle unknown-scoped objects.
	// Therefore, if len(c.objsSkipped) > 0, we would parse the configs from scratch.
	return c.hasParserResult && len(c.objsSkipped) == 0 && c.parserErrs == nil
}

// splitObjects splits `objs` into two groups: the objects whose scope is known, and the objects whose scope is unknown.
func splitObjects(objs []ast.FileObject) ([]ast.FileObject, []ast.FileObject) {
	var knownScopeObjs, unknownScopeObjs []ast.FileObject
	var unknownScopeIDs []string
	for _, obj := range objs {
		if core.GetAnnotation(obj, metadata.UnknownScopeAnnotationKey) == metadata.UnknownScopeAnnotationValue {
			unknownScopeObjs = append(unknownScopeObjs, obj)
			unknownScopeIDs = append(unknownScopeIDs, core.GKNN(obj.Unstructured))
		} else {
			knownScopeObjs = append(knownScopeObjs, obj)
		}
	}
	if len(unknownScopeIDs) > 0 {
		sort.Strings(unknownScopeIDs)
		klog.Infof("Skip sending %v unknown-scoped objects to the applier: %v", len(unknownScopeIDs), unknownScopeIDs)
	}
	return knownScopeObjs, unknownScopeObjs
}
