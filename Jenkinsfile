pipeline {
    agent any
    tools {
        dockerTool 'docker'
    }
    stages {        
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
                sh 'docker login -u ${dockerhub_USR} -p ${dockerhub_PSW}'
                sh 'docker push donatmolnar/guesspass:1.0.0'
            }
        }
    }
}