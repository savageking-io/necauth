pipeline {

    agent {
        node {
            label 'Go Builder'
        }
    }

    stages {
        stage('Build') {
            steps {
                script {
                    sh 'go version'
                }
            }
        }
    }

}