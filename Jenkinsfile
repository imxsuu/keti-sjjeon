podTemplate(label: 'podman-argocd',
  containers: [
    containerTemplate(
      name: 'podman',
      image: 'mgoltzsche/podman',
      command: 'cat',
      ttyEnabled: true,
      privileged: true
    ),
    containerTemplate(
      name: 'argocd',
      image: 'argoproj/argo-cd-ci-builder:latest',
      command: 'cat',
      ttyEnabled: true
    ),
  ],
) {
    node('podman-argocd') {
        
        stage('Checkout'){
            container('argocd'){
                checkout scm
            }
        }
        
        stage('Build'){
            container('podman'){
                sh("""
                    #!/bin/sh

                    # Construct Image Name
                    IMAGE=10.0.1.150:5000/openfx/openfx-gateway:${env.BUILD_NUMBER}
                    
                    podman build -t \${IMAGE} .
                """)
            }
        }

        stage('Push'){
            container('podman'){
                sh("""
                    #!/bin/sh

                    # Construct Image Name
                    IMAGE=10.0.1.150:5000/openfx/openfx-gateway:${env.BUILD_NUMBER}
                   
                    podman login -u admin -p Ketilinux11 10.0.1.150:5000 --tls-verify=false

                    podman push \${IMAGE} --tls-verify=false
                """)
            }
        }
        

        stage('Deploy'){
            container('argocd'){
                checkout([$class: 'GitSCM',
                        branches: [[name: '*/develop' ]],
                        extensions: scm.extensions,
                        userRemoteConfigs: [[
                            url: 'http://github.com/ketiops/openfx.git',
                            credentialsId: 'jenkins_agent_ssh',
                        ]]
                ])
                sshagent(credentials: ['jenkins_agent_ssh']){
                    sh("""
                        #!/usr/bin/env bash
                        set +x
                        export GIT_SSH_COMMAND="ssh -oStrictHostKeyChecking=no"
                        git config --global user.email "imxsuu@gmail.com"
                        git checkout develop
                        cd overlay/dev && kustomize edit set image 10.0.1.150:5000/openfx/openfx-gateway:${BUILD_NUMBER}
                        git commit -a -m "[UPDATE] change the image versioning ${currentBuild.number}"
                        git push
                    """)
                }
            }
        }
    } 
}
