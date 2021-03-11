pipeline {
    agent {
        docker { image 'golang:1.16.0-alpine3.13' }
    }
    environment {
        GO116MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Install dependencies') {
            steps {
                echo 'Installing dependencies'
                sh ```go version
go get -d ./...
```
            }
        }
        stage('Building the app') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }
        stage('Building Docker image') {
            steps {
                echo 'Compiling and building'
                sh ```
docker build -t donatmolnar/guesspass:1.0.0 .
                ```
            }
        }
        stage('Pushing image to dockerhub') {
            environment {
                dockerhub = credentials(dockerhub)
            }
            steps {
                echo 'Pushing to dockerhub'
                sh ```
sudo docker login
${dockerhub_USR} 
${dockerhub_PSW}
sudo docker push donatmolnar/guesspass:1.0.0
                ```
            }
        }
    }
    post {
        always {}
    }  
}