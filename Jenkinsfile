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
                sh 'cd services/cinemahallservice && go test'
                sh 'cd services/movieservice && go test'
                sh 'cd services/reservationservice && go test'
                sh 'cd services/showservice && go test'
                sh 'cd services/userservice && go test'
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
                sh "docker-build-and-push -b ${BRANCH_NAME} -s userservice -f services/userservice/Dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s movieservice -f services/movieservice/Dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s cinemahallservice -f services/cinemahallservice/Dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s showservice -f services/showservice/Dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservationservice -f services/reservationservice/Dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s client -f client.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s services -f services.dockerfile"

            }
        }
    }
}
