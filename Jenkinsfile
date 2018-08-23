node {
    def root = tool name: 'Go1.10.3', type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
        withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
    
            stage('Test'){
                sh 'echo $GOPATH'
                sh 'go version'                   
            }
        }
    }
}