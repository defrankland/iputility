pipeline {
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go test ./...'
            }
        }
    }
}