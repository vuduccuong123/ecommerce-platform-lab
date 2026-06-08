pipeline {
    agent {
        kubernetes {
            yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    command:
    - cat
    tty: true
'''
        }
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/vuduccuong123/ecommerce-platform-lab.git'
            }
        }

        stage('Build and Push Image') {
            steps {
                container('kaniko') {
                    withCredentials([
                        usernamePassword(
                            credentialsId: 'jenkins-devops',
                            usernameVariable: 'DOCKER_USER',
                            passwordVariable: 'DOCKER_PASS'
                        )
                    ]) {
                        sh '''
                        mkdir -p /kaniko/.docker

                        AUTH=$(printf "%s:%s" "$DOCKER_USER" "$DOCKER_PASS" | base64 | tr -d '\\n')

                        cat > /kaniko/.docker/config.json <<EOF
{
  "auths": {
    "https://index.docker.io/v1/": {
      "auth": "${AUTH}"
    }
  }
}
EOF

                        /kaniko/executor \
                          --context $WORKSPACE/app/src/productcatalogservice \
                          --dockerfile $WORKSPACE/app/src/productcatalogservice/Dockerfile \
                          --destination docker.io/cuong2003/productcatalogservice:${BUILD_NUMBER}
                        '''
                    }
                }
            }
        }
    }
}