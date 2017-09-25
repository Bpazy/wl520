node {
    stage('git')
    {
        git url: 'https://github.com/Bpazy/welove520.git'
    }
    stage('go get')
    {
        sh 'go build github.com/Bpazy/welove520'
    }
}