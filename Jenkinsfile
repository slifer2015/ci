node {
    stage('Build') {
        checkout scm
        sh "make all"
    }
    stage('deploy') {
        checkout scm
        sh "docker build ."
    }
}