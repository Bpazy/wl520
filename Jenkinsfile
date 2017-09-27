pipeline {
    stages {
        stage('pull source code')
        {
            git url: 'https://github.com/Bpazy/wl520.git'
        }
        stage('build')
        {
            sh 'go build -o wl520 .'
        }
        stage('deploy')
        {
            sh 'mv wl520 $GOPATH/bin/'
        }
    }
}