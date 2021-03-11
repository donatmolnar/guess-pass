pipeline {
    agent any
    tools {
        go 'go1.16'
    }
    environment {
        GO116MODULE = 'on'
        CGO_ENABLED = 0
        //GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        //GOROOT="${root}"
        //PATH+GO="${root}/bin"
        //GOBIN="${root}/bin"
        //GOPATH="${root}/go"
    }
    stages {        
        stage('Install dependencies') {
            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin:${HOME}/go/bin"]) {
                    echo 'Installing dependencies'
                    sh 'go version'
                    sh 'go get -d ./...'
                }
            }
        }
        stage('Building the app') {
            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin:${HOME}/go/bin"]) {
                    echo 'Compiling and building'
                    sh 'go build main.go'
                }
            }
        }
        stage('Building Docker image') {
            steps {
                echo 'Building Docker image'
                sh 'docker build -t donatmolnar/guesspass:1.0.0 .'
            }
        }
        stage('Pushing image to dockerhub') {
            environment {
                dockerhub = credentials('dockerhub')
            }
            steps {
                echo 'Pushing to dockerhub'
                sh 'sudo docker login'
                sh '${dockerhub_USR}'
                sh '${dockerhub_PSW}'
                sh 'sudo docker push donatmolnar/guesspass:1.0.0'
            }
        }
    }
}