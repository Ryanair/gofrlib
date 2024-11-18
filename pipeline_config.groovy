libraries{
    git
    codebuild {
        projectGroup = 'RDT'
        sourceType = 'bitbucket'
        singleEnvPipeline = true
        scmHttpsUrl = "https://bitbucket.org/ryanair/gofrlib.git"
    }
}