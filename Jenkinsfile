#!/usr/bin/env groovy

node('swarm') {
    String APP_PATH = "/go/src/github.com/elastic/cloud-sdk-go"

    stage('Checkout from GitHub') {
	checkout scm
    }
    docker.image("golang:1.12-stretch").inside("-u root:root -v ${pwd()}:${APP_PATH} -w ${APP_PATH} -e APP_PATH=$APP_PATH") {
        stage("Download dependencies") {
            sh 'cd $APP_PATH && make deps vendor'
        }
        stage("Run linters") {
            sh 'cd $APP_PATH && make lint'
        }
        stage("Compile and Run Unit Tests") {
            sh 'cd $APP_PATH && make unit unit-coverage'
            publishHTML([allowMissing: false, alwaysLinkToLastBuild: false, keepAll: false, reportDir: 'reports', reportFiles: 'coverage.html', reportName: 'Coverage Report', reportTitles: ''])
        }
   }
}
