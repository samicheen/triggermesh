/*
Copyright 2020 TriggerMesh Inc.

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
	"knative.dev/pkg/injection/sharedmain"

	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awscloudwatchlogssource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awscloudwatchsource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awscodecommitsource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awscognitoidentitysource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awscognitouserpoolsource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awsdynamodbsource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awskinesissource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awsperformanceinsightssource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awss3source"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awssnssource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awssqssource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/httppollersource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/slacksource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/webhooksource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/zendesksource"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/alibabaosstarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/awscomprehendtarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/awstarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/confluenttarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/datadogtarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/elasticsearchtarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/googlecloudfirestoretarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/googlecloudstoragetarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/googlecloudworkflowstarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/googlesheettarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/hasuratarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/httptarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/infratarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/jiratarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/logztarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/oracletarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/salesforcetarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/sendgridtarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/slacktarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/splunktarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/tektontarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/twiliotarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/uipathtarget"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/zendesktarget"
)

func main() {
	sharedmain.Main("triggermesh-controller",
		// sources
		awscloudwatchlogssource.NewController,
		awscloudwatchsource.NewController,
		awscodecommitsource.NewController,
		awscognitoidentitysource.NewController,
		awscognitouserpoolsource.NewController,
		awsdynamodbsource.NewController,
		awskinesissource.NewController,
		awsperformanceinsightssource.NewController,
		awss3source.NewController,
		awssnssource.NewController,
		awssqssource.NewController,
		httppollersource.NewController,
		slacksource.NewController,
		webhooksource.NewController,
		zendesksource.NewController,
		// targets
		alibabaosstarget.NewController,
		awstarget.NewDynamoDBController,
		awstarget.NewLambdaController,
		awstarget.NewS3Controller,
		awstarget.NewSNSController,
		awstarget.NewSQSController,
		awstarget.NewKinesisController,
		awscomprehendtarget.NewController,
		confluenttarget.NewController,
		elasticsearchtarget.NewController,
		googlecloudstoragetarget.NewController,
		googlecloudfirestoretarget.NewController,
		googlecloudworkflowstarget.NewController,
		googlesheettarget.NewController,
		hasuratarget.NewController,
		httptarget.NewController,
		datadogtarget.NewController,
		infratarget.NewController,
		jiratarget.NewController,
		logztarget.NewController,
		oracletarget.NewController,
		salesforcetarget.NewController,
		sendgridtarget.NewController,
		slacktarget.NewController,
		splunktarget.NewController,
		tektontarget.NewController,
		twiliotarget.NewController,
		uipathtarget.NewController,
		zendesktarget.NewController,
	)
}