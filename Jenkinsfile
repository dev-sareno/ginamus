pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'cat /etc/os-release'
                script {
                    // TODO: Fix "docker not found"
                    def customImage = docker.build("devsareno/testimage:latest", "-f ./golang")
                    customImage.push()
                }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
