pipeline {
    agent any

    options {
        disableConcurrentBuilds()
    }

    stages {
        stage('Checkout'){
            steps {
                checkout scm
            }
        }
        stage('Build binaries') {
            steps {
                withEnv(['PATH+EXTRA=/usr/local/go/bin']) {
                    sh 'make'
                }
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
        stage('Run Benchmark on real input') {
            steps {
                withCredentials([string(credentialsId: 'aoc_token', variable: 'AOC_AUTH_TOKEN')]) {
                    withEnv(['PATH+EXTRA=/usr/local/go/bin']) {
                        sh 'make benchmark TOKEN=$AOC_AUTH_TOKEN'
                    }
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