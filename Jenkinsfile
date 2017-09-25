node {
    stage('git')
    {
        git url: 'https://github.com/Bpazy/wl520.git'
    }
    stage('go get')
    {
        sh 'go build -o wl520 .'
    }
}