pipeline {
    agent any

    stages {
        stage('Checkout'){
            steps {
                checkout scm
            }
        }
        stage('Run tests') {

            steps {
                withEnv(['PATH+EXTRA=/usr/local/go/bin']) {
                    sh '''
                        make testallv
                        '''
                }
            }
        }
    }

    post {
        always {
            sh 'docker logout'
        }
    }
}