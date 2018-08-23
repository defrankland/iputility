pipeline {
    agent none
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go test ./...'
            }
        }
    }
}