pipeline {
    agent { node { label 'jenkins-agent'} }
    stages('Checkout'){
        container('podman'){
            checkout scm
        }
    }    
    stages('Build'){
        container('podman'){
            sh("""
               #!/bin/sh

               # Construct Image Name
               IMAGE=10.0.1.150:5000/openfx/openfx-gateway:${env.BUILD_NUMBER}
                  
               podman build -t \${IMAGE} .
               """)
        }
    }
    stages('Push'){
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
    stages('Deploy'){
        container('argocd'){
            checkout([$class: 'GitSCM',
                     branches: [[name: '*/main' ]],
                     extensions: scm.extensions,
                     userRemoteConfigs: [[
                        url: 'https://github.com/ketiops/openfx.git',
                        credentialsId: 'github-jenkins',
                     ]]             
            ])
            sshagent(credentials: ['github-jenkins']){
                sh("""
                   #!/usr/bin/env bash
                   set +x
                   export GIT_SSH_COMMAND="ssh -oStrictHostKeyChecking=no"
                   git config --global user.email "imxsuu@gmail.com"
                   git checkout develop
                   cd overlay/dev && kustomize edit set image 10.0.1.150:5000/openfx/openfx-gateway:${BUILD_NUMBER}
                   """)
            }
        }
    } 
}
