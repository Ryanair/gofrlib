@Library('awsCommonTemplates') _

properties([
        parameters([
                string(name: 'RELEASE_VERSION', description: 'Release version'),
        ])
])

app_env = pipelineConfig.application_environments[params.RELEASE_ENVIRONMENT]

checkout_code()
tag(params.RELEASE_VERSION)



