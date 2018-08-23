node {
    def root = tool name: '1.10', type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
        withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
    
            stage('Test'){
                sh 'echo $GOPATH'
                sh 'go version'                   
            }
        }
    }
}