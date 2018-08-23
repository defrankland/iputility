node {
    def root = tool name: '1.10', type: 'go'
ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/defrankland/iputility") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {

            env.PATH="${GOPATH}/bin:$PATH"

            stage ('Checkout'){
                git branch: 'jenkinsfile', url: 'https://github.com/defrankland/iputility'
            }

            stage('Test'){
                sh 'echo $GOPATH'
                sh 'go version'                   
                sh 'ls -la'
                sh 'go test -cover ./...'
            }
        }
    }
}