pipeline {
  agent any
  stages {
    stage('Pull') {
      steps {
        git 'https://github.com/Bpazy/wl520.git'
      }
    }
    stage('Build') {
      steps {
        parallel(
          "Test": {
            sh 'go test ./welove'
            
          },
          "Build": {
            sh 'go build -o wl520 .'
            
          }
        )
      }
    }
    stage('Deploy') {
      steps {
        sh 'mv wl520 $GOPATH/bin/'
      }
    }
  }
}