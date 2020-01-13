pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo build'
                sh 'cd cli && go build client.go'
                sh 'cd services && go build main.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests...'
                sh 'cd services/cinemahall && go test'
                sh 'cd services/movie && go test'
                sh 'cd services/reservation && go test'
                sh 'cd services/show && go test'
                sh 'cd services/user && go test'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo no lint'
                //sh 'golangci-lint run --deadline 20m --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "echo build docker"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s client -f client.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s services -f services.dockerfile"
            }
        }
    }
}
