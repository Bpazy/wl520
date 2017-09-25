node {
    stage('git')
    {
        git url: 'https://github.com/Bpazy/welove520.git'
    }
    stage('go get')
    {
        sh 'go install github.com/Bpazy/welove520'
    }
}