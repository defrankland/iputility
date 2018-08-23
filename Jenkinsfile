node {
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
        withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
            env.PATH="${GOPATH}/bin:$PATH"
    
            stage('Test'){
                sh 'echo $GOPATH'
                sh 'go version'                   
            }
        }
    }
}